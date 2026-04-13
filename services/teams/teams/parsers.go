package teams

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (team *Team) ToMap() map[string]any {
	teamMap := team.ToAuditMap()
	teamMap["name"] = helper.NilIfEmpty(team.Name)
	teamMap["policy_ids"] = helper.NilIfEmpty(team.PolicyIds)
	teamMap["members"] = helper.NilIfEmpty(team.Members)
	teamMap["tags"] = helper.ParseToMaps(team.Tags)
	return teamMap
}

func (team *Team) FromMap(teamMap map[string]any) error {
	team.FromAuditMap(teamMap)
	team.Name = helper.CastString(teamMap, "name")
	team.PolicyIds = make([]string, len(helper.CastSlice(teamMap, "policy_ids")))
	for i, value := range helper.CastSlice(teamMap, "policy_ids") {
		team.PolicyIds[i] = value.(string)
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(teamMap, "tags")); err != nil {
		return err
	} else {
		team.Tags = tags
	}
	team.Members = make([]string, len(helper.CastSlice(teamMap, "members")))
	for i, value := range helper.CastSlice(teamMap, "members") {
		team.Members[i] = value.(string)
	}
	return nil
}
