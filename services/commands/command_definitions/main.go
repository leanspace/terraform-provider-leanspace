package command_definitions

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

const commandDefinitionAPIPath = "commands-repository/command-definitions"

var CommandDataType = provider.DataSourceType[CommandDefinition, *CommandDefinition]{
	ResourceIdentifier: "leanspace_command_definitions",
	Path:               commandDefinitionAPIPath,
	Schema:             commandDefinitionSchema,
	FilterSchema:       dataSourceFilterSchema,
	TFModelFactory:     func() any { return &CommandDefinitionTF{} },
	SchemaVersion:      1,
	// Use StateUpgraderFactory (instead of the generic StateUpgraders) so that the
	// upgrader can call the API and get arguments/metadata in their canonical API order.
	// This avoids Set→List ordering mismatches that would cause "Provider produced
	// inconsistent result after apply" errors on resources with set-typed list fields.
	StateUpgraderFactory: commandDefinitionUpgraderFactory,
}

// commandDefinitionUpgraderFactory is a named function (not a lambda) to avoid
// Go initialization cycles.  It deliberately does NOT reference CommandDataType —
// it builds a GenericClient directly from the hard-coded API path.
func commandDefinitionUpgraderFactory(client *provider.Client) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var partial struct {
					ID string `json:"id"`
				}
				if err := json.Unmarshal(req.RawState.JSON, &partial); err != nil {
					resp.Diagnostics.AddError("State upgrade: failed to parse resource ID from v0 state", err.Error())
					return
				}
				apiClient := provider.GenericClient[CommandDefinition, *CommandDefinition]{
					Client: client,
					Path:   commandDefinitionAPIPath,
				}
				var zero CommandDefinition
				value, err := apiClient.Get(partial.ID, &zero)
				if err != nil {
					resp.Diagnostics.AddError("State upgrade: failed to re-read command definition from API", err.Error())
					return
				}
				if value == nil {
					resp.Diagnostics.AddError("State upgrade: command definition not found in API", partial.ID)
					return
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, (*value).ToTF())...)
			},
		},
	}
}
