package asset

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCommandDefinition() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCommandDefinitionCreate,
		ReadContext:   resourceCommandDefinitionRead,
		UpdateContext: resourceCommandDefinitionUpdate,
		DeleteContext: resourceCommandDefinitionDelete,
		Schema: map[string]*schema.Schema{
			"command_definition": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: commandDefinitionSchema,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCommandDefinitionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	commandDefinition := d.Get("command_definition").([]interface{})
	commandDefinitionData := getCommandDefinitionData(commandDefinition)
	createCommandDefinition, err := client.CreateCommandDefinition(commandDefinitionData.NodeId, commandDefinitionData)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(createCommandDefinition.ID)
	for _, diag := range resourceCommandDefinitionRead(ctx, d, m) {
		diags = append(diags, diag)
	}

	return diags
}

func resourceCommandDefinitionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	commandDefinitionId := d.Id()
	commandDefinition, err := client.GetCommandDefinition(commandDefinitionId)
	if err != nil {
		return diag.FromErr(err)
	}
	setCommandDefinitionData(commandDefinition, d)

	return diags
}

func setCommandDefinitionData(commandDefinition *CommandDefinition, d *schema.ResourceData) {
	commandDefinitionList := make([]map[string]interface{}, 1)

	commandDefinitionList[0] = commandDefinitionStructToInterface(*commandDefinition)
	d.Set("command_definition", commandDefinitionList)
}

func getCommandDefinitionData(commandDefinitionList []interface{}) CommandDefinition {
	commandDefinition := commandDefinitionList[0].(map[string]interface{})
	commandDefinitionMap := CommandDefinition{}

	commandDefinitionMap.NodeId = commandDefinition["node_id"].(string)
	commandDefinitionMap.Name = commandDefinition["name"].(string)
	commandDefinitionMap.Description = commandDefinition["description"].(string)
	commandDefinitionMap.Identifier = commandDefinition["identifier"].(string)
	commandDefinitionMap.CreatedAt = commandDefinition["created_at"].(string)
	commandDefinitionMap.CreatedBy = commandDefinition["created_by"].(string)
	commandDefinitionMap.LastModifiedAt = commandDefinition["last_modified_at"].(string)
	commandDefinitionMap.LastModifiedBy = commandDefinition["last_modified_by"].(string)
	if commandDefinition["metadata"] != nil {
		commandDefinitionMap.Metadata = []Metadata[interface{}]{}
		for _, metadata := range commandDefinition["metadata"].([]interface{}) {
			commandDefinitionMap.Metadata = append(commandDefinitionMap.Metadata, metadataInterfaceToStruct(metadata.(map[string]interface{})))
		}
	}
	if commandDefinition["arguments"] != nil {
		commandDefinitionMap.Arguments = []Argument[interface{}]{}
		for _, argument := range commandDefinition["arguments"].([]interface{}) {
			commandDefinitionMap.Arguments = append(commandDefinitionMap.Arguments, argumentInterfaceToStruct(argument.(map[string]interface{})))
		}
	}

	return commandDefinitionMap
}

func metadataInterfaceToStruct(metadata map[string]interface{}) Metadata[interface{}] {
	metadataStruct := Metadata[interface{}]{}
	metadataStruct.ID = metadata["id"].(string)
	metadataStruct.Name = metadata["name"].(string)
	metadataStruct.Description = metadata["description"].(string)
	metadataStruct.UnitId = metadata["unit_id"].(string)
	metadataStruct.Value = metadata["value"]
	metadataStruct.Required = metadata["required"].(bool)
	metadataStruct.Type = metadata["type"].(string)

	return metadataStruct
}

func argumentInterfaceToStruct(argument map[string]interface{}) Argument[interface{}] {
	argumentStruct := Argument[interface{}]{}
	argumentStruct.ID = argument["id"].(string)
	argumentStruct.Name = argument["name"].(string)
	argumentStruct.Identifier = argument["identifier"].(string)
	argumentStruct.Description = argument["description"].(string)
	argumentStruct.Type = argument["type"].(string)
	argumentStruct.Required = argument["required"].(bool)
	switch argumentStruct.Type {
	case "NUMERIC":
		argumentStruct.Min = argument["min"].(float64)
		argumentStruct.Max = argument["max"].(float64)
		argumentStruct.Scale = argument["scale"].(int)
		argumentStruct.Precision = argument["precision"].(int)
		argumentStruct.UnitId = argument["unit_id"].(string)
	case "ENUM":
		if argument["options"] != nil {
			option := argument["options"].(map[string]interface{})
			argumentStruct.Options = &option
		}
	case "TEXT":
		argumentStruct.MinLength = argument["min_length"].(int)
		argumentStruct.MaxLength = argument["max_length"].(int)
		argumentStruct.Pattern = argument["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		argumentStruct.Before = argument["before"].(string)
		argumentStruct.After = argument["after"].(string)
	}
	argumentStruct.DefaultValue = argument["default_value"]

	return argumentStruct
}

func resourceCommandDefinitionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("command_definition") {
		commandDefinitionId := d.Id()
		commandDefinition := d.Get("command_definition").([]interface{})
		_, err := client.UpdateCommandDefinition(commandDefinitionId, getCommandDefinitionData(commandDefinition))
		if err != nil {
			return diag.FromErr(err)
		}
		for _, diag := range resourceCommandDefinitionRead(ctx, d, m) {
			diags = append(diags, diag)
		}

		return diags
	}

	return resourceCommandDefinitionCreate(ctx, d, m)

}

func resourceCommandDefinitionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	commandDefinitionId := d.Id()
	err := client.DeleteCommandDefinition(commandDefinitionId)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
