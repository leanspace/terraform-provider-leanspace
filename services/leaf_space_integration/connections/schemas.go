package connections

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var leafSpaceConnectionSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"domain_url": resourceschema.StringAttribute{
		Required: true,
	},
	"password": resourceschema.StringAttribute{
		Optional:  true,
		Computed:  true,
		Sensitive: true,
	},
	"username": resourceschema.StringAttribute{
		Optional: true,
		Computed: true,
	},
	"authentication_token": resourceschema.StringAttribute{
		Optional: true,
	},
	"status": resourceschema.StringAttribute{
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
}

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
