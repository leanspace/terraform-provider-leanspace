package teams

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-asset/asset"
)

func contains(slice []string, element string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func (team *Team) currentPolicies(client *asset.Client) ([]string, error) {
	path := fmt.Sprintf("%s/%s/%s/access-policies", client.HostURL, TeamDataType.Path, team.ID)
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

func (team *Team) currentMembers(client *asset.Client) ([]string, error) {
	path := fmt.Sprintf("%s/teams-repository/members?teamIds=%s", client.HostURL, team.ID)
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
	rawMembers := dataMap["content"].([]any)
	currentMembers := make([]string, len(rawMembers))
	for i, member := range rawMembers {
		currentMembers[i] = member.(map[string]any)["id"].(string)
	}
	return currentMembers, nil
}

type apiValidPolicies struct {
	PolicyIds []string `json:"policyIds"`
}

func (team *Team) policyChange(action string, policies []string, client *asset.Client) error {
	policyData, err := json.Marshal(apiValidPolicies{PolicyIds: policies})
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/%s/access-policies", client.HostURL, TeamDataType.Path, team.ID)
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

func (team *Team) addPolicies(policies []string, client *asset.Client) error {
	return team.policyChange("POST", policies, client)
}

func (team *Team) removePolicies(policies []string, client *asset.Client) error {
	return team.policyChange("DELETE", policies, client)
}

func (team *Team) memberChange(action string, memberId string, client *asset.Client) error {
	path := fmt.Sprintf("%s/%s/%s/members/%s", client.HostURL, TeamDataType.Path, team.ID, memberId)
	req, err := http.NewRequest(action, path, nil)
	if err != nil {
		return err
	}
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func (team *Team) addMember(memberId string, client *asset.Client) error {
	return team.memberChange("POST", memberId, client)
}

func (team *Team) removeMember(memberId string, client *asset.Client) error {
	return team.memberChange("DELETE", memberId, client)
}

func (team *Team) PostReadProcess(client *asset.Client) error {
	if policies, err := team.currentPolicies(client); err != nil {
		return err
	} else {
		team.PolicyIds = policies
	}
	if members, err := team.currentMembers(client); err != nil {
		return err
	} else {
		team.Members = members
	}
	return nil
}

// Although policies can be set on creation, members can't, so we just need to add members here.
func (team *Team) PostCreateProcess(client *asset.Client, teamRaw any) error {
	createdTeam := teamRaw.(*Team)
	expectedMembers := team.Members

	// Add all members directly
	for _, member := range expectedMembers {
		err := createdTeam.addMember(member, client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (team *Team) PostUpdateProcess(client *asset.Client, teamRaw any) error {
	teamCurrent := teamRaw.(*Team)
	currentPolicies, err := teamCurrent.currentPolicies(client)
	if err != nil {
		return err
	}
	expectedPolicies := team.PolicyIds

	policiesToRemove := []string{}
	policiesToAdd := []string{}

	// Diff policies to see what to add/remove
	for _, policy := range currentPolicies {
		if !contains(expectedPolicies, policy) {
			policiesToRemove = append(policiesToRemove, policy)
		}
	}
	for _, policy := range expectedPolicies {
		if !contains(currentPolicies, policy) {
			policiesToAdd = append(policiesToAdd, policy)
		}
	}

	// Apply diff
	if len(policiesToRemove) > 0 {
		err = team.removePolicies(policiesToRemove, client)
		if err != nil {
			return err
		}
	}
	if len(policiesToAdd) > 0 {
		err = team.addPolicies(policiesToAdd, client)
		if err != nil {
			return err
		}
	}

	currentMembers, err := teamCurrent.currentMembers(client)
	if err != nil {
		return err
	}
	expectedMembers := team.Members

	// Diff members, and add/remove directly (we're limited to 1 member per request)
	for _, member := range currentMembers {
		if !contains(expectedMembers, member) {
			err = team.removeMember(member, client)
			if err != nil {
				return err
			}
		}
	}
	for _, member := range expectedMembers {
		if !contains(currentMembers, member) {
			err = team.addMember(member, client)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
