package members

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func (member *Member) ToMap() map[string]any {
	memberMap := make(map[string]any)
	memberMap["id"] = member.ID
	memberMap["name"] = member.Name
	memberMap["email"] = member.Email
	memberMap["status"] = member.Status
	memberMap["policy_ids"] = member.PolicyIds
	memberMap["created_at"] = member.CreatedAt
	memberMap["created_by"] = member.CreatedBy
	memberMap["last_modified_at"] = member.LastModifiedAt
	memberMap["last_modified_by"] = member.LastModifiedBy

	return memberMap
}

func (member *Member) FromMap(memberMap map[string]any) error {
	member.ID = memberMap["id"].(string)
	member.Name = memberMap["name"].(string)
	member.Email = memberMap["email"].(string)
	member.Status = memberMap["status"].(string)
	member.PolicyIds = make([]string, memberMap["policy_ids"].(*schema.Set).Len())
	for i, value := range memberMap["policy_ids"].(*schema.Set).List() {
		member.PolicyIds[i] = value.(string)
	}
	member.CreatedAt = memberMap["created_at"].(string)
	member.CreatedBy = memberMap["created_by"].(string)
	member.LastModifiedAt = memberMap["last_modified_at"].(string)
	member.LastModifiedBy = memberMap["last_modified_by"].(string)

	return nil
}
