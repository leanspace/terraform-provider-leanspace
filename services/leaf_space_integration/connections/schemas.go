package connections

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var leafSpaceConnectionSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"domain_url": resourceschema.StringAttribute{
		Required: true,
	},
	"password": resourceschema.StringAttribute{
		Optional:  true,
		Sensitive: true,
	},
	"username": resourceschema.StringAttribute{
		Optional: true,
	},
	"authentication_token": resourceschema.StringAttribute{
		Optional: true,
	},
	"status": resourceschema.StringAttribute{
		Computed: true,
	},
})

var leafSpaceConnectionFilterSchema = map[string]datasourceschema.Attribute{
	"id": datasourceschema.StringAttribute{
		Computed: true,
	},
	"created_at": datasourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": datasourceschema.StringAttribute{
		Computed:    true,
		Description: "Who created it",
	},
	"domain_url": datasourceschema.StringAttribute{
		Computed: true,
	},
	"last_modified_at": datasourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": datasourceschema.StringAttribute{
		Computed:    true,
		Description: "Who modified it the last",
	},
	"status": datasourceschema.StringAttribute{
		Computed: true,
	},
}
