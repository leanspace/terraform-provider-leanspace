package remote_agents

import (
	"encoding/json"
	"fmt"
	"leanspace-terraform-provider/provider"
	"net/http"
)

// Cleanup function needed because an access policy and a service account are created
// with each remote agent and aren't remove automatically.
func (agent *RemoteAgent) PostDeleteProcess(client *provider.Client) error {
	// Fetch matching acess policies
	path := fmt.Sprintf("%s/teams-repository/service-accounts/%s/access-policies", client.HostURL, agent.ServiceAccountId)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}
	data, err, code := client.DoRequest(req, &(client).Token)
	if code == http.StatusNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	dataMap := make(map[string]any)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		return err
	}

	// Delete policies
	for _, policy := range dataMap["content"].([]any) {
		policyId := policy.(map[string]any)["id"]
		path = fmt.Sprintf("%s/teams-repository/access-policies/%s", client.HostURL, policyId)
		req, err := http.NewRequest("DELETE", path, nil)
		if err != nil {
			return err
		}
		_, err, code = client.DoRequest(req, &(client).Token)
		if code != http.StatusNotFound && err != nil {
			return err
		}
	}

	// Delete service account
	path = fmt.Sprintf("%s/teams-repository/service-accounts/%s", client.HostURL, agent.ServiceAccountId)
	req, err = http.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	_, err, code = client.DoRequest(req, &(client).Token)
	if code != http.StatusNotFound && err != nil {
		return err
	}

	return nil
}
