package connections

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var leafSpaceConnectionSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"domain_url": {
		Type:     schema.TypeString,
		Required: true,
	},
	"password": {
		Type:      schema.TypeString,
		Optional:  true,
		Computed:  true,
		Sensitive: true,
	},
	"username": {
		Type:      schema.TypeString,
		Optional:  true,
		Computed:  true,
		Sensitive: true,
	},
	"authentication_token": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"status": {
		Type:     schema.TypeString,
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
}

var leafSpaceConnectionFilterSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
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
	"domain_url": {
		Type:     schema.TypeString,
		Computed: true,
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
	"status": {
		Type:     schema.TypeString,
		Computed: true,
	},
}
