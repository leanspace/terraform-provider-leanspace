package asset

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAsset() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAssetCreate,
		ReadContext:   resourceAssetRead,
		UpdateContext: resourceAssetUpdate,
		DeleteContext: resourceAssetDelete,
		Schema: map[string]*schema.Schema{
			"asset": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"created_at": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified_at": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified_by": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_node_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"kind": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"tags": tagsSchema,
						"nodes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: assetSchema,
							},
						},
						"norad_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"international_designator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tle": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 2,
							MinItems: 2,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAssetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	asset := d.Get("asset").([]interface{})
	createAsset, err := client.CreateOrder(getAssetData(asset))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(createAsset.ID)
	for _, diag := range resourceAssetRead(ctx, d, m) {
		diags = append(diags, diag)
	}

	return diags
}

func resourceAssetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	assetId := d.Id()
	asset, err := client.GetAsset(assetId)
	if err != nil {
		return diag.FromErr(err)
	}
	setAssetData(asset, d)

	return diags
}

func assetStructToInterface(asset *Asset, level int) interface{} {
	assetMap := make(map[string]interface{})

	assetMap["id"] = asset.ID
	assetMap["name"] = asset.Name
	assetMap["description"] = asset.Description
	assetMap["created_at"] = asset.CreatedAt
	assetMap["created_by"] = asset.CreatedBy
	assetMap["parent_node_id"] = asset.ParentNodeId
	assetMap["last_modified_at"] = asset.LastModifiedAt
	assetMap["last_modified_by"] = asset.LastModifiedBy
	assetMap["type"] = asset.Type
	assetMap["kind"] = asset.Kind
	assetMap["tags"] = tagsStructToInterface(asset.Tags)
	if asset.Nodes != nil && level == 0 {
		nodes := make([]interface{}, len(asset.Nodes))
		for i, node := range asset.Nodes {
			nodes[i] = assetStructToInterface(&node, level+1)
		}
		assetMap["nodes"] = nodes
	}
	if len(asset.NoradId) != 0 {
		assetMap["norad_id"] = asset.NoradId
	}
	if len(asset.InternationalDesignator) != 0 {
		assetMap["international_designator"] = asset.InternationalDesignator
	}
	if len(asset.Tle) != 2 {
		assetMap["tle"] = asset.Tle
	}

	return assetMap
}

func setAssetData(asset *Asset, d *schema.ResourceData) {
	assetList := make([]map[string]interface{}, 1)

	assetList[0] = assetStructToInterface(asset, 0).(map[string]interface{})
	d.Set("asset", assetList)
}

func assetInterfaceToStruct(asset map[string]interface{}) Asset {
	assetStruct := Asset{}

	assetStruct.Name = asset["name"].(string)
	assetStruct.Description = asset["description"].(string)
	assetStruct.CreatedAt = asset["created_at"].(string)
	assetStruct.CreatedBy = asset["created_by"].(string)
	assetStruct.ParentNodeId = asset["parent_node_id"].(string)
	assetStruct.LastModifiedAt = asset["last_modified_at"].(string)
	assetStruct.LastModifiedBy = asset["last_modified_by"].(string)
	assetStruct.Type = asset["type"].(string)
	assetStruct.Kind = asset["kind"].(string)
	assetStruct.Tags = tagsInterfaceToStruct(asset["tags"])
	if asset["nodes"] != nil {
		assetStruct.Nodes = make([]Asset, len(asset["nodes"].([]interface{})))
		for i, node := range asset["nodes"].([]interface{}) {
			assetStruct.Nodes[i] = assetInterfaceToStruct(node.(map[string]interface{}))
		}
	}
	assetStruct.NoradId = asset["norad_id"].(string)
	assetStruct.InternationalDesignator = asset["international_designator"].(string)
	if asset["tle"] != nil && len(asset["tle"].([]interface{})) == 2 {
		assetStruct.Tle = make([]string, 2)
		for i, tle := range asset["tle"].([]interface{}) {
			assetStruct.Tle[i] = tle.(string)
		}

	}

	return assetStruct
}

func getAssetData(assetList []interface{}) Asset {
	return assetInterfaceToStruct(assetList[0].(map[string]interface{}))
}

func resourceAssetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("asset") {
		assetId := d.Id()
		asset := d.Get("asset").([]interface{})
		_, err := client.UpdateAsset(assetId, getAssetData(asset))
		if err != nil {
			return diag.FromErr(err)
		}
		for _, diag := range resourceAssetRead(ctx, d, m) {
			diags = append(diags, diag)
		}

		return diags
	}

	return resourceAssetCreate(ctx, d, m)

}

func resourceAssetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	assetId := d.Id()
	err := client.DeleteAsset(assetId)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
