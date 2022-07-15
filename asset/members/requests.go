package members

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-asset/asset"
)

func (member *Member) currentPolicies(client *asset.Client) ([]string, error) {
	path := fmt.Sprintf("%s/%s/%s/access-policies", client.HostURL, MemberDataType.Path, member.ID)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	data, err, _ := client.DoRequest(req, &(client).Token)
	if err != nil {
		return nil, err
	}
	dataMap := make(map[string]any)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		return nil, err
	}
	rawPolicies := dataMap["content"].([]any)
	currentPolicies := make([]string, len(rawPolicies))
	for i, policy := range rawPolicies {
		currentPolicies[i] = policy.(map[string]any)["id"].(string)
	}
	return currentPolicies, nil
}

type apiValidPolicies struct {
	PolicyIds []string `json:"policyIds"`
}

func (member *Member) policyChange(action string, policies []string, client *asset.Client) error {
	policyData, err := json.Marshal(apiValidPolicies{PolicyIds: policies})
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/%s/access-policies", client.HostURL, MemberDataType.Path, member.ID)
	req, err := http.NewRequest(action, path, strings.NewReader(string(policyData)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func (member *Member) addPolicies(policies []string, client *asset.Client) error {
	return member.policyChange("POST", policies, client)
}

func (member *Member) removePolicies(policies []string, client *asset.Client) error {
	return member.policyChange("DELETE", policies, client)
}

func (member *Member) PostReadProcess(client *asset.Client) error {
	if policies, err := member.currentPolicies(client); err != nil {
		return err
	} else {
		member.PolicyIds = policies
	}
	return nil
}

func (member *Member) PostUpdateProcess(client *asset.Client, memberRaw any) error {
	memberCurrent := memberRaw.(*Member)
	currentPolicies, err := memberCurrent.currentPolicies(client)
	if err != nil {
		return err
	}
	expectedPolicies := member.PolicyIds

	policiesToRemove := []string{}
	policiesToAdd := []string{}

	// Diff policies to see what to add/remove
	for _, policy := range currentPolicies {
		if !asset.Contains(expectedPolicies, policy) {
			policiesToRemove = append(policiesToRemove, policy)
		}
	}
	for _, policy := range expectedPolicies {
		if !asset.Contains(currentPolicies, policy) {
			policiesToAdd = append(policiesToAdd, policy)
		}
	}

	// Apply diff
	if len(policiesToRemove) > 0 {
		err = member.removePolicies(policiesToRemove, client)
		if err != nil {
			return err
		}
	}
	if len(policiesToAdd) > 0 {
		err = member.addPolicies(policiesToAdd, client)
		if err != nil {
			return err
		}
	}

	return nil
}
