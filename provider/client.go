package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client -
type Client struct {
	HostURL      string
	HTTPClient   *http.Client
	Token        string
	Auth         AuthStruct
	RetryTimeout time.Duration
}

// AuthStruct -
type AuthStruct struct {
	Env          string `json:"env"`
	Tenant       string `json:"tenant"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Region       string `json:"region"`
}

// AuthResponse -
type AuthResponse struct {
	TokenType string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
	Token     string `json:"access_token"`
}

func NewClient(host, env, tenant, clientId, clientSecret, region *string) (*Client, error) {
	hostUrl := "https://api.leanspace.io"
	switch *env {
	case "prod":
		hostUrl = "https://api.leanspace.io"
	default:
		hostUrl = fmt.Sprintf("https://api.%s.leanspace.io", *env)
	}
	timeout := 300 * time.Second

	c := Client{
		HTTPClient:   &http.Client{Timeout: timeout},
		HostURL:      hostUrl,
		RetryTimeout: timeout,
	}

	if host != nil && *host != "" {
		c.HostURL = *host
	}

	// If tenant, clientId or clientSecret not provided, return empty client
	if tenant == nil || clientId == nil || clientSecret == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Env:          *env,
		Tenant:       *tenant,
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
		Region:       *region,
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token
	scheduledTime := time.Duration(ar.ExpiresIn)*time.Second - 2*time.Minute // schedule token refresh before expiration to avoid 401 while the token refreshes
	scheduledTimeIfError := 5 * time.Second
	errorCount := 0

	go func(scheduledTime time.Duration) {
		ticker := time.NewTicker(scheduledTime)
		defer ticker.Stop()  // should never be called as we never exit the function
		for range ticker.C { // enters at each scheduledTime
			authResponse, err := c.SignIn()
			if err != nil {
				errorCount++
				if errorCount > 5 {
					panic("Token couldn't be refreshed") // avoid infinite loop
				}
				ticker.Reset(scheduledTimeIfError) // call sign in earlier in case of error
				continue
			} else {
				errorCount = 0
				ticker.Reset(scheduledTime) // reset to original scheduled time in case it was changed
			}
			c.Token = authResponse.Token
		}
	}(scheduledTime)

	return &c, nil
}

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.ClientId == "" || c.Auth.ClientSecret == "" {
		return nil, fmt.Errorf("define client id and client secret")
	}

	tokenUrl := fmt.Sprintf("%s/teams-repository/oauth2/token?tenant=%s", c.HostURL, c.Auth.Tenant)

	req, err := http.NewRequest("POST", tokenUrl, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Auth.ClientId, c.Auth.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body, err, _ := c.DoRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) DoRequest(req *http.Request, authToken *string) ([]byte, error, int) {

	if authToken != nil {
		token := *authToken
		req.Header.Set("Authorization", "Bearer "+token)
	}

	req.Header.Set("X-LS-Application", "terraform-provider-leanspace")

	// We want to be able to print the request body in the error message,
	// in case it goes wrong.
	// Because req.Body is a ReadCloser, reading it twice isn't possible. We thus
	// need to make a copy of the content, and replace the current request body.
	bodyOriginal := &bytes.Buffer{}
	if req.Body != nil {
		_, err := io.Copy(bodyOriginal, req.Body)
		if err != nil {
			return nil, err, 0
		}
		bodyCopy := bytes.NewReader(bodyOriginal.Bytes())
		req.Body = io.NopCloser(bodyCopy)
		bodyCopy.Seek(0, 0)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		status := 0
		if res != nil {
			status = res.StatusCode
		}
		return nil, err, status
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err, res.StatusCode
	}

	if res.StatusCode != http.StatusOK {
		var prettyRequestJSON bytes.Buffer
		json.Indent(&prettyRequestJSON, bodyOriginal.Bytes(), "", "    ")

		var prettyResponseJSON bytes.Buffer
		json.Indent(&prettyResponseJSON, body, "", "    ")

		extra := ""
		if res.StatusCode == 409 {
			extra = "Hint: This seems to be an error caused by a name collision.\n" +
				"Try renaming your resource or deleting the resource with the same name on Leanspace."
		}

		return nil, fmt.Errorf(
			"status %d when performing the request.\n"+
				"Sent %s to %s\n"+
				"Request body: %s\n"+
				"Response body: %s\n"+
				"%s",
			res.StatusCode,
			req.Method,
			req.URL,
			&prettyRequestJSON,
			&prettyResponseJSON,
			extra,
		), res.StatusCode
	}

	return body, err, res.StatusCode
}
