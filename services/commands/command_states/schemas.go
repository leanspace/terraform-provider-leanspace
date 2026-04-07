package command_states

import (
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var commandStateSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidStateName(),
	},
	"read_only": resourceschema.BoolAttribute{
		Computed: true,
	},
	"created_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who created it",
	},
	"last_modified_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who modified it the last",
	},
	"tags": general_objects.KeyValuesSchema,
}
