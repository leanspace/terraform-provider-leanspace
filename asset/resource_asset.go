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
				Optional: true,
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
						},
						"kind": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"tags": tagsSchema,
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

func setAssetData(asset *Asset, d *schema.ResourceData) {
	assetList := make([]map[string]interface{}, 1)
	newAsset := make(map[string]interface{})

	newAsset["id"] = asset.ID
	newAsset["name"] = asset.Name
	newAsset["description"] = asset.Description
	newAsset["created_at"] = asset.CreatedAt
	newAsset["created_by"] = asset.CreatedBy
	newAsset["parent_node_id"] = asset.ParentNodeId
	newAsset["last_modified_at"] = asset.LastModifiedAt
	newAsset["last_modified_by"] = asset.LastModifiedBy
	newAsset["type"] = asset.Type
	newAsset["kind"] = asset.Kind
	newAsset["tags"] = tagsStructToInterface(asset.Tags)

	assetList[0] = newAsset
	d.Set("asset", assetList)
}

func getAssetData(assetList []interface{}) Asset {
	asset := assetList[0].(map[string]interface{})
	newAsset := Asset{}

	newAsset.Name = asset["name"].(string)
	newAsset.Description = asset["description"].(string)
	newAsset.CreatedAt = asset["created_at"].(string)
	newAsset.CreatedBy = asset["created_by"].(string)
	newAsset.ParentNodeId = asset["parent_node_id"].(string)
	newAsset.LastModifiedAt = asset["last_modified_at"].(string)
	newAsset.LastModifiedBy = asset["last_modified_by"].(string)
	newAsset.Type = asset["type"].(string)
	newAsset.Kind = asset["kind"].(string)
	newAsset.Tags = tagsInterfaceToStruct(asset["tags"])

	return newAsset
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
