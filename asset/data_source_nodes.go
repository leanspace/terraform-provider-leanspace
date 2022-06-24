package asset

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNodesRead,
		Schema: map[string]*schema.Schema{
			"nodes": &schema.Schema{
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

func dataSourceNodesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	nodes, err := client.forNodes().GetAll()
	if err != nil {
		return diag.FromErr(err)
	}
	setNodesData(nodes, d)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func setNodesData(nodes *PaginatedList[Node], d *schema.ResourceData) {
	nodeList := make([]map[string]interface{}, len(nodes.Content))

	for i, node := range nodes.Content {
		nodeMap := make(map[string]interface{})
		nodeMap["id"] = node.ID
		nodeMap["type"] = node.Type
		nodeMap["kind"] = node.Kind
		nodeMap["name"] = node.Name
		nodeMap["description"] = node.Description
		nodeList[i] = nodeMap
	}
	d.Set("nodes", nodeList)

	d.Set("total_elements", nodes.TotalElements)
	d.Set("total_pages", nodes.TotalPages)
	d.Set("number_of_elements", nodes.NumberOfElements)
	d.Set("number", nodes.Number)
	d.Set("size", nodes.Size)
	d.Set("first", nodes.First)
	d.Set("last", nodes.Last)
	d.Set("empty", nodes.Empty)

	sort := sortStructToInterface(nodes.Sort)
	d.Set("sort", sort)
	d.Set("pageable", pageableStructToInterface(nodes.Pageable, sort))
}
