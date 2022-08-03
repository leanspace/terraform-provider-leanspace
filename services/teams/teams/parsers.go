package teams

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func (team *Team) ToMap() map[string]any {
	teamMap := make(map[string]any)
	teamMap["id"] = team.ID
	teamMap["name"] = team.Name
	teamMap["policy_ids"] = team.PolicyIds
	teamMap["members"] = team.Members
	teamMap["created_at"] = team.CreatedAt
	teamMap["created_by"] = team.CreatedBy
	teamMap["last_modified_at"] = team.LastModifiedAt
	teamMap["last_modified_by"] = team.LastModifiedBy

	return teamMap
}

func (team *Team) FromMap(teamMap map[string]any) error {
	team.ID = teamMap["id"].(string)
	team.Name = teamMap["name"].(string)
	team.PolicyIds = make([]string, teamMap["policy_ids"].(*schema.Set).Len())
	for i, value := range teamMap["policy_ids"].(*schema.Set).List() {
		team.PolicyIds[i] = value.(string)
	}
	team.Members = make([]string, teamMap["members"].(*schema.Set).Len())
	for i, value := range teamMap["members"].(*schema.Set).List() {
		team.Members[i] = value.(string)
	}
	team.CreatedAt = teamMap["created_at"].(string)
	team.CreatedBy = teamMap["created_by"].(string)
	team.LastModifiedAt = teamMap["last_modified_at"].(string)
	team.LastModifiedBy = teamMap["last_modified_by"].(string)

	return nil
}
