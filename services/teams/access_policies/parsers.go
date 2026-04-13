package access_policies

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (policy *AccessPolicy) ToMap() map[string]any {
	policyMap := policy.ToAuditMap()
	policyMap["name"] = helper.NilIfEmpty(policy.Name)
	policyMap["description"] = helper.NilIfEmpty(policy.Description)
	policyMap["read_only"] = helper.NilIfEmpty(policy.ReadOnly)
	policyMap["statements"] = helper.ParseToMaps(policy.Statements)
	policyMap["tags"] = helper.ParseToMaps(policy.Tags)
	return policyMap
}

func (statement *Statement) ToMap() map[string]any {
	statementMap := make(map[string]any)
	statementMap["name"] = helper.NilIfEmpty(statement.Name)
	statementMap["actions"] = helper.NilIfEmpty(statement.Actions)
	return statementMap
}

func (policy *AccessPolicy) FromMap(policyMap map[string]any) error {
	policy.FromAuditMap(policyMap)
	policy.Name = helper.CastString(policyMap, "name")
	policy.Description = helper.CastString(policyMap, "description")
	policy.ReadOnly = helper.CastBool(policyMap, "read_only")
	if statements, err := helper.ParseFromMaps[Statement](helper.CastSlice(policyMap, "statements")); err != nil {
		return err
	} else {
		policy.Statements = statements
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(policyMap, "tags")); err != nil {
		return err
	} else {
		policy.Tags = tags
	}
	return nil
}

func (statement *Statement) FromMap(statementMap map[string]any) error {
	statement.Name = helper.CastString(statementMap, "name")
	statement.Actions = make([]string, len(helper.CastSlice(statementMap, "actions")))
	for i, action := range helper.CastSlice(statementMap, "actions") {
		statement.Actions[i] = action.(string)
	}
	return nil
}
