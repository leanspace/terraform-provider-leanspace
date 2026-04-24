package command_states

import (
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var commandStateSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidStateName(),
	},
	"read_only": resourceschema.BoolAttribute{
		Computed: true,
	},
	"tags": general_objects.KeyValuesSchema,
})
