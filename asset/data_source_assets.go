package asset

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAssetsRead,
		Schema: map[string]*schema.Schema{
			"assets": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"total_elements": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_pages": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_elements": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sort": sortSchema,
			"first": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"empty": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"pageable": pageableSchema,
		},
	}
}

func dataSourceAssetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	assets, err := client.GetAllAssets()
	if err != nil {
		return diag.FromErr(err)
	}
	setAssetsData(assets, d)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func setAssetsData(assets *PaginatedList[Asset], d *schema.ResourceData) {
	assetList := make([]map[string]interface{}, len(assets.Content))

	for i, asset := range assets.Content {
		assetMap := make(map[string]interface{})
		assetMap["id"] = asset.ID
		assetMap["type"] = asset.Type
		assetMap["kind"] = asset.Kind
		assetMap["name"] = asset.Name
		assetMap["description"] = asset.Description
		assetList[i] = assetMap
	}
	d.Set("assets", assetList)

	d.Set("total_elements", assets.TotalElements)
	d.Set("total_pages", assets.TotalPages)
	d.Set("number_of_elements", assets.NumberOfElements)
	d.Set("number", assets.Number)
	d.Set("size", assets.Size)
	d.Set("first", assets.First)
	d.Set("last", assets.Last)
	d.Set("empty", assets.Empty)

	sort := sortStructToInterface(assets.Sort)
	d.Set("sort", sort)
	d.Set("pageable", pageableStructToInterface(assets.Pageable, sort))
}
