package service_accounts

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var serviceAccountSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"policy_ids": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"credentials": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: credentialSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
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

var credentialSchema = map[string]*schema.Schema{
	"client_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"client_secret": {
		Type:     schema.TypeString,
		Computed: true,
	},
}
