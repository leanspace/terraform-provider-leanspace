package activity_definitions

import (
	"fmt"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (activityDefinition *ActivityDefinition) PreMarshallProcess() error {
	for i := range activityDefinition.CommandMappings {
		activityDefinition.CommandMappings[i].Position = i
	}
	return nil
}

func (activityDefinition *ActivityDefinition) PostReadProcess(_ *provider.Client, newValue any) error {
	newDef, ok := newValue.(*ActivityDefinition)
	if !ok || newDef == nil {
		return nil
	}
	newDef.Tags = general_objects.ReorderKeyValues(activityDefinition.Tags, newDef.Tags)
	newDef.Metadata = helper.ReorderByKey(activityDefinition.Metadata, newDef.Metadata,
		func(m Metadata[any]) string { return m.Name })
	newDef.ArgumentDefinitions = helper.ReorderByKey(activityDefinition.ArgumentDefinitions, newDef.ArgumentDefinitions,
		func(a ArgumentDefinition[any]) string { return a.Name })
	newDef.CommandMappings = helper.ReorderByKey(activityDefinition.CommandMappings, newDef.CommandMappings,
		func(c CommandMapping) string { return fmt.Sprintf("%s:%d", c.CommandDefinitionId, c.Position) })
	for i, stateCmd := range activityDefinition.CommandMappings {
		if i >= len(newDef.CommandMappings) {
			break
		}
		newDef.CommandMappings[i].ArgumentMappings = helper.ReorderByKey(
			stateCmd.ArgumentMappings, newDef.CommandMappings[i].ArgumentMappings,
			func(m ArgumentMapping) string { return m.ActivityDefinitionArgumentName },
		)
		newDef.CommandMappings[i].MetadataMappings = helper.ReorderByKey(
			stateCmd.MetadataMappings, newDef.CommandMappings[i].MetadataMappings,
			func(m MetadataMapping) string { return m.ActivityDefinitionMetadataName },
		)
	}
	return nil
}
