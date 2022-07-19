package service_accounts

import (
	"encoding/json"
	"fmt"
	"leanspace-terraform-provider/helper"
	"leanspace-terraform-provider/provider"
	"net/http"
	"strings"
)

func (serviceAccount *ServiceAccount) currentPolicies(client *provider.Client) ([]string, error) {
	path := fmt.Sprintf("%s/%s/%s/access-policies", client.HostURL, ServiceAccountDataType.Path, serviceAccount.ID)
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

func (serviceAccount *ServiceAccount) policyChange(action string, policies []string, client *provider.Client) error {
	policyData, err := json.Marshal(apiValidPolicies{PolicyIds: policies})
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/%s/access-policies", client.HostURL, ServiceAccountDataType.Path, serviceAccount.ID)
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

func (serviceAccount *ServiceAccount) addPolicies(policies []string, client *provider.Client) error {
	return serviceAccount.policyChange("POST", policies, client)
}

func (serviceAccount *ServiceAccount) removePolicies(policies []string, client *provider.Client) error {
	return serviceAccount.policyChange("DELETE", policies, client)
}

func (serviceAccount *ServiceAccount) PostReadProcess(client *provider.Client, rawServiceAccount any) error {
	currentServiceAccount := rawServiceAccount.(*ServiceAccount)
	if policies, err := currentServiceAccount.currentPolicies(client); err != nil {
		return err
	} else {
		currentServiceAccount.PolicyIds = policies
	}
	return nil
}

func (serviceAccount *ServiceAccount) PostUpdateProcess(client *provider.Client, serviceAccountRaw any) error {
	serviceAccountCurrent := serviceAccountRaw.(*ServiceAccount)
	currentPolicies, err := serviceAccountCurrent.currentPolicies(client)
	if err != nil {
		return err
	}
	expectedPolicies := serviceAccount.PolicyIds

	policiesToRemove := []string{}
	policiesToAdd := []string{}

	// Diff policies to see what to add/remove
	for _, policy := range currentPolicies {
		if !helper.Contains(expectedPolicies, policy) {
			policiesToRemove = append(policiesToRemove, policy)
		}
	}
	for _, policy := range expectedPolicies {
		if !helper.Contains(currentPolicies, policy) {
			policiesToAdd = append(policiesToAdd, policy)
		}
	}

	// Apply diff
	if len(policiesToRemove) > 0 {
		err = serviceAccount.removePolicies(policiesToRemove, client)
		if err != nil {
			return err
		}
	}
	if len(policiesToAdd) > 0 {
		err = serviceAccount.addPolicies(policiesToAdd, client)
		if err != nil {
			return err
		}
	}

	return nil
}
