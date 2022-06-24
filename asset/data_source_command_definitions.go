package asset

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCommandDefinitions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCommandDefinitionsRead,
		Schema: map[string]*schema.Schema{
			"command_definitions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: commandDefinitionSchema,
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

func dataSourceCommandDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	commandDefinitions, err := client.forCommandDefinitions().GetAll()
	if err != nil {
		return diag.FromErr(err)
	}
	setCommandDefinitionsData(commandDefinitions, d)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func setCommandDefinitionsData(commandDefinitions *PaginatedList[CommandDefinition], d *schema.ResourceData) {
	commandDefinitionList := make([]map[string]interface{}, len(commandDefinitions.Content))

	for i, commandDefinition := range commandDefinitions.Content {
		commandDefinitionList[i] = commandDefinitionStructToInterface(commandDefinition)
	}
	d.Set("command_definitions", commandDefinitionList)

	d.Set("total_elements", commandDefinitions.TotalElements)
	d.Set("total_pages", commandDefinitions.TotalPages)
	d.Set("number_of_elements", commandDefinitions.NumberOfElements)
	d.Set("number", commandDefinitions.Number)
	d.Set("size", commandDefinitions.Size)
	d.Set("first", commandDefinitions.First)
	d.Set("last", commandDefinitions.Last)
	d.Set("empty", commandDefinitions.Empty)

	sort := sortStructToInterface(commandDefinitions.Sort)
	d.Set("sort", sort)
	d.Set("pageable", pageableStructToInterface(commandDefinitions.Pageable, sort))
}
