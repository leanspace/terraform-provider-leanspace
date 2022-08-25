package provider

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Env          string `json:"env"`
	Tenant       string `json:"tenant"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// AuthResponse -
type AuthResponse struct {
	TokenType string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
	Token     string `json:"access_token"`
}

func NewClient(host, env, tenant, clientId, clientSecret *string) (*Client, error) {
	if env == nil || *env == "" {
		environment := "prod"
		env = &environment
	}
	hostUrl := "https://api.leanspace.io"
	switch *env {
	case "develop", "demo":
		hostUrl = fmt.Sprintf("https://api.%s.leanspace.io", *env)
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 120 * time.Second},
		HostURL:    hostUrl,
	}

	if host != nil {
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
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.ClientId == "" || c.Auth.ClientSecret == "" {
		return nil, fmt.Errorf("define client id and client secret")
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s-%s.auth.eu-central-1.amazoncognito.com/oauth2/token?scope=https://api.leanspace.io/READ&grant_type=client_credentials", c.Auth.Tenant, c.Auth.Env), strings.NewReader("Content-Type=application%2Fx-www-form-urlencoded"))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Auth.ClientId, c.Auth.ClientSecret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
