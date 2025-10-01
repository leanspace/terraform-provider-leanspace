package access_policies

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (policy *AccessPolicy) ToMap() map[string]any {
	policyMap := make(map[string]any)
	policyMap["id"] = policy.ID
	policyMap["name"] = policy.Name
	policyMap["description"] = policy.Description
	policyMap["read_only"] = policy.ReadOnly
	policyMap["statements"] = helper.ParseToMaps(policy.Statements)
	policyMap["tags"] = helper.ParseToMaps(policy.Tags)
	policyMap["created_at"] = policy.CreatedAt
	policyMap["created_by"] = policy.CreatedBy
	policyMap["last_modified_at"] = policy.LastModifiedAt
	policyMap["last_modified_by"] = policy.LastModifiedBy

	return policyMap
}

func (statement *Statement) ToMap() map[string]any {
	statementMap := make(map[string]any)
	statementMap["name"] = statement.Name
	statementMap["actions"] = statement.Actions
	return statementMap
}

func (policy *AccessPolicy) FromMap(policyMap map[string]any) error {
	policy.ID = policyMap["id"].(string)
	policy.Name = policyMap["name"].(string)
	policy.Description = policyMap["description"].(string)
	policy.ReadOnly = policyMap["read_only"].(bool)
	if statements, err := helper.ParseFromMaps[Statement](policyMap["statements"].(*schema.Set).List()); err != nil {
		return err
	} else {
		policy.Statements = statements
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](policyMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		policy.Tags = tags
	}
	policy.CreatedAt = policyMap["created_at"].(string)
	policy.CreatedBy = policyMap["created_by"].(string)
	policy.LastModifiedAt = policyMap["last_modified_at"].(string)
	policy.LastModifiedBy = policyMap["last_modified_by"].(string)

	return nil
}

func (statement *Statement) FromMap(statementMap map[string]any) error {
	statement.Name = statementMap["name"].(string)
	statement.Actions = make([]string, len(statementMap["actions"].(*schema.Set).List()))
	for i, action := range statementMap["actions"].(*schema.Set).List() {
		statement.Actions[i] = action.(string)
	}
	return nil
}
