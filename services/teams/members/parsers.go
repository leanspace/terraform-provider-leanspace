package members

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (member *Member) ToMap() map[string]any {
	memberMap := member.ToAuditMap()
	memberMap["name"] = helper.NilIfEmpty(member.Name)
	memberMap["email"] = helper.NilIfEmpty(member.Email)
	memberMap["status"] = helper.NilIfEmpty(member.Status)
	memberMap["policy_ids"] = helper.NilIfEmpty(member.PolicyIds)
	return memberMap
}

func (member *Member) FromMap(memberMap map[string]any) error {
	member.FromAuditMap(memberMap)
	member.Name = helper.CastString(memberMap, "name")
	member.Email = helper.CastString(memberMap, "email")
	member.Status = helper.CastString(memberMap, "status")
	member.PolicyIds = make([]string, len(helper.CastSlice(memberMap, "policy_ids")))
	for i, value := range helper.CastSlice(memberMap, "policy_ids") {
		member.PolicyIds[i] = value.(string)
	}
	return nil
}
