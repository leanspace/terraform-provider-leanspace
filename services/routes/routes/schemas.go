package routes

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validLogLevels = []string{"INFO", "DEBUG", "TRACE", "WARN", "ERROR"}

var routeSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"tags": general_objects.KeyValuesSchema,

	"definition": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: definitionSchema,
	},

	"route_instances": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: routeInstanceSchema,
		},
	},

	"processor_ids": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.Set{setvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
})

var definitionSchema = map[string]resourceschema.Attribute{
	"configuration": resourceschema.StringAttribute{
		Optional: true,
	},
	"log_level": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validLogLevels),
		Validators:  []validator.String{stringvalidator.OneOf(validLogLevels...)},
	},
	"valid": resourceschema.BoolAttribute{
		Computed: true,
	},
	"service_account_id": resourceschema.StringAttribute{
		Computed:   true,
		Optional:   true,
		Validators: helper.ValidUUID(),
	},
	"errors": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: errorSchema,
		},
	},
}

var errorSchema = map[string]resourceschema.Attribute{
	"code": resourceschema.StringAttribute{
		Computed: true,
	},
	"message": resourceschema.StringAttribute{
		Computed: true,
	},
}

var routeInstanceSchema = map[string]resourceschema.Attribute{
	"status": resourceschema.StringAttribute{
		Required: true,
	},
	"last_status_at": resourceschema.StringAttribute{
		Optional: true,
	},
	"container_id": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"last_message_start_process_at": resourceschema.StringAttribute{
		Optional: true,
	},
	"last_message_end_process_at": resourceschema.StringAttribute{
		Optional: true,
	},
	"number_of_messages_processed": resourceschema.Int64Attribute{
		Optional: true,
	},
	"camel_route_id": resourceschema.StringAttribute{
		Optional:   true,
		Validators: helper.ValidUUID(),
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"last_modified_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who last modified the entry. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
