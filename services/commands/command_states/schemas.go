package command_states

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var reservedLeanspaceStates = []string{"DRAFT", "STAGED", "QUEUED", "POPPED", "ABORTED"}

func validateStateName(value interface{}, key string) ([]string, []error) {
	stateNameWarnings, stateNameErrors := helper.IsValidStateName(value, key)
	reservedStateWarnings, reservedStateErrors := validation.StringNotInSlice(reservedLeanspaceStates, false)(value, key)
	return append(stateNameWarnings, reservedStateWarnings...), append(stateNameErrors, reservedStateErrors...)
}

var commandStateSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validateStateName,
	},
	"read_only": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"created_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who created it",
	},
	"last_modified_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who modified it the last",
	},
	"tags": general_objects.KeyValuesSchema,
}
