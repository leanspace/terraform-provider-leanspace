package service_accounts

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var serviceAccountSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"policy_ids": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Required:    true,
		Validators:  []validator.Set{setvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"credentials": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: credentialSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
})

var credentialSchema = map[string]resourceschema.Attribute{
	"client_id": resourceschema.StringAttribute{
		Computed: true,
	},
	"client_secret": resourceschema.StringAttribute{
		Computed: true,
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"policy_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
