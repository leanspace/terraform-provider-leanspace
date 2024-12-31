package request_states

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var requestStateSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: helper.IsValidStateName,
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

var requestStateFilterSchema = map[string]*schema.Schema{
	"created_bys": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"to_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
	},
	"last_modified_bys": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"from_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
	},
	"to_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
	},
}
