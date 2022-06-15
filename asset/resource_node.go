package asset

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNodeCreate,
		ReadContext:   resourceNodeRead,
		UpdateContext: resourceNodeUpdate,
		DeleteContext: resourceNodeDelete,
		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
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
								Schema: nodeSchema,
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

func resourceNodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	node := d.Get("node").([]interface{})
	createNode, err := client.CreateNode(getNodeData(node))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(createNode.ID)
	for _, diag := range resourceNodeRead(ctx, d, m) {
		diags = append(diags, diag)
	}

	return diags
}

func resourceNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	nodeId := d.Id()
	node, err := client.GetNode(nodeId)
	if err != nil {
		return diag.FromErr(err)
	}
	setNodeData(node, d)

	return diags
}

func nodeStructToInterface(node *Node, level int) interface{} {
	nodeMap := make(map[string]interface{})

	nodeMap["id"] = node.ID
	nodeMap["name"] = node.Name
	nodeMap["description"] = node.Description
	nodeMap["created_at"] = node.CreatedAt
	nodeMap["created_by"] = node.CreatedBy
	nodeMap["parent_node_id"] = node.ParentNodeId
	nodeMap["last_modified_at"] = node.LastModifiedAt
	nodeMap["last_modified_by"] = node.LastModifiedBy
	nodeMap["type"] = node.Type
	nodeMap["kind"] = node.Kind
	nodeMap["tags"] = tagsStructToInterface(node.Tags)
	if node.Nodes != nil && level == 0 {
		nodes := make([]interface{}, len(node.Nodes))
		for i, node := range node.Nodes {
			nodes[i] = nodeStructToInterface(&node, level+1)
		}
		nodeMap["nodes"] = nodes
	}
	if len(node.NoradId) != 0 {
		nodeMap["norad_id"] = node.NoradId
	}
	if len(node.InternationalDesignator) != 0 {
		nodeMap["international_designator"] = node.InternationalDesignator
	}
	if len(node.Tle) != 2 {
		nodeMap["tle"] = node.Tle
	}

	return nodeMap
}

func setNodeData(node *Node, d *schema.ResourceData) {
	nodeList := make([]map[string]interface{}, 1)

	nodeList[0] = nodeStructToInterface(node, 0).(map[string]interface{})
	d.Set("node", nodeList)
}

func nodeInterfaceToStruct(node map[string]interface{}) Node {
	nodeStruct := Node{}

	nodeStruct.Name = node["name"].(string)
	nodeStruct.Description = node["description"].(string)
	nodeStruct.CreatedAt = node["created_at"].(string)
	nodeStruct.CreatedBy = node["created_by"].(string)
	nodeStruct.ParentNodeId = node["parent_node_id"].(string)
	nodeStruct.LastModifiedAt = node["last_modified_at"].(string)
	nodeStruct.LastModifiedBy = node["last_modified_by"].(string)
	nodeStruct.Type = node["type"].(string)
	nodeStruct.Kind = node["kind"].(string)
	nodeStruct.Tags = tagsInterfaceToStruct(node["tags"])
	if node["nodes"] != nil {
		nodeStruct.Nodes = make([]Node, len(node["nodes"].([]interface{})))
		for i, node := range node["nodes"].([]interface{}) {
			nodeStruct.Nodes[i] = nodeInterfaceToStruct(node.(map[string]interface{}))
		}
	}
	nodeStruct.NoradId = node["norad_id"].(string)
	nodeStruct.InternationalDesignator = node["international_designator"].(string)
	if node["tle"] != nil && len(node["tle"].([]interface{})) == 2 {
		nodeStruct.Tle = make([]string, 2)
		for i, tle := range node["tle"].([]interface{}) {
			nodeStruct.Tle[i] = tle.(string)
		}

	}

	return nodeStruct
}

func getNodeData(nodeList []interface{}) Node {
	return nodeInterfaceToStruct(nodeList[0].(map[string]interface{}))
}

func resourceNodeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("node") {
		nodeId := d.Id()
		node := d.Get("node").([]interface{})
		_, err := client.UpdateNode(nodeId, getNodeData(node))
		if err != nil {
			return diag.FromErr(err)
		}
		for _, diag := range resourceNodeRead(ctx, d, m) {
			diags = append(diags, diag)
		}

		return diags
	}

	return resourceNodeCreate(ctx, d, m)

}

func resourceNodeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	nodeId := d.Id()
	err := client.DeleteNode(nodeId)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
