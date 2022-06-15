package asset

import (
	"context"
	"fmt"
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var tle1stLine = `^1 (?P<noradId>[ 0-9]{5})[A-Z] [ 0-9]{5}[ A-Z]{3} [ 0-9]{5}[.][ 0-9]{8} (?:(?:[ 0+-][.][ 0-9]{8})|(?: [ +-][.][ 0-9]{7})) [ +-][ 0-9]{5}[+-][ 0-9] [ +-][ 0-9]{5}[+-][ 0-9] [ 0-9] [ 0-9]{4}[ 0-9]$`
var tle2ndLine = `^2 (?P<noradId>[ 0-9]{5}) [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{7} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{2}[.][ 0-9]{13}[ 0-9]$`

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
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								value := val.(string)
								if !(value == "ASSET" || value == "GROUP" || value == "COMPONENT") {
								  errs = append(errs, fmt.Errorf("%q must be either ASSET, GROUP ou COMPONENT, got: %q", key, value))
								}
								return
							  },
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
							ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d{5}$`),"It must be 5 digits"),
						},
						"international_designator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d{4}-|\d{2})[0-9]{3}[A-Za-z]{0,3}$`),""),
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
	nodeData, err := getNodeData(node)
	if err != nil {
		return diag.FromErr(err)
	}
	createNode, err := client.CreateNode(nodeData)
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

func nodeInterfaceToStruct(node map[string]interface{}) (Node, error) {
	nodeStruct := Node{}

	nodeStruct.Name = node["name"].(string)
	nodeStruct.Description = node["description"].(string)
	nodeStruct.CreatedAt = node["created_at"].(string)
	nodeStruct.CreatedBy = node["created_by"].(string)
	nodeStruct.ParentNodeId = node["parent_node_id"].(string)
	nodeStruct.LastModifiedAt = node["last_modified_at"].(string)
	nodeStruct.LastModifiedBy = node["last_modified_by"].(string)
	nodeStruct.Type = node["type"].(string)
	if nodeStruct.Type == "ASSET" && !(node["kind"] == "GENERIC" || node["kind"] == "SATELLITE" || node["kind"] == "GROUND_STATION") {
		return nodeStruct, fmt.Errorf("kind must be either GENERIC, SATELLITE ou GROUND_STATION, got: %q", node["kind"])
	}
	nodeStruct.Kind = node["kind"].(string)
	nodeStruct.Tags = tagsInterfaceToStruct(node["tags"])
	if node["nodes"] != nil {
		nodeStruct.Nodes = make([]Node, len(node["nodes"].([]interface{})))
		for i, node := range node["nodes"].([]interface{}) {
			childNodeStruct, err := nodeInterfaceToStruct(node.(map[string]interface{}))
			if err != nil {
				return nodeStruct, err
			}
			nodeStruct.Nodes[i] = childNodeStruct
		}
	}
	nodeStruct.NoradId = node["norad_id"].(string)
	nodeStruct.InternationalDesignator = node["international_designator"].(string)
	if node["tle"] != nil && len(node["tle"].([]interface{})) == 2 {
		nodeStruct.Tle = make([]string, 2)
		matched, _ := regexp.MatchString(tle1stLine,node["tle"].([]interface{})[0].(string))
		if !matched {
			return nodeStruct, fmt.Errorf("TLE first line mutch match %q, got: %q", tle1stLine, node["tle"].([]interface{})[0].(string))
		}
		matched, _ =regexp.MatchString(tle2ndLine,node["tle"].([]interface{})[1].(string))
		if !matched {
			return nodeStruct, fmt.Errorf("TLE second line mutch match %q, got: %q", tle2ndLine, node["tle"].([]interface{})[1].(string))
		}
		for i, tle := range node["tle"].([]interface{}) {
			nodeStruct.Tle[i] = tle.(string)
		}

	}

	return nodeStruct, nil
}

func getNodeData(nodeList []interface{}) (Node, error) {
	return nodeInterfaceToStruct(nodeList[0].(map[string]interface{}))
}

func resourceNodeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("node") {
		nodeId := d.Id()
		node := d.Get("node").([]interface{})
		nodeData, err := getNodeData(node)
		if err != nil {
			return diag.FromErr(err)
		}
		_, err = client.UpdateNode(nodeId, nodeData)
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
