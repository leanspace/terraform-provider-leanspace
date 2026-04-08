package request_states

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var requestStateSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidStateName(),
	},
})

var requestStateFilterSchema = map[string]datasourceschema.Attribute{
	"created_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"to_created_at": datasourceschema.StringAttribute{
		Optional: true,
	},
	"last_modified_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"from_last_modified_at": datasourceschema.StringAttribute{
		Optional: true,
	},
	"to_last_modified_at": datasourceschema.StringAttribute{
		Optional: true,
	},
}
