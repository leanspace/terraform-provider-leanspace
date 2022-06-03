package asset

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

var sortSchema = &schema.Schema{
	Type:     schema.TypeList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"property": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_case": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"null_handling": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ascending": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"descending": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	},
}

var pageableSchema = &schema.Schema{
	Type:     schema.TypeList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sort": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"property": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ignore_case": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"null_handling": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ascending": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"descending": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"offset": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page_number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"paged": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"unpaged": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	},
}

var tagsSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}

var propertyFieldSchema = map[string]*schema.Schema{
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
	"value": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
}

var geoPointFieldsSchema = map[string]*schema.Schema{
	"elevation": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: propertyFieldSchema,
		},
	},
	"latitude": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: propertyFieldSchema,
		},
	},
	"longitude": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: propertyFieldSchema,
		},
	},
}

var propertySchema = map[string]*schema.Schema{
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
	"node_id": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
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
	"tags": tagsSchema,
	"min_length": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	},
	"max_length": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	},
	"pattern": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"fields": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: geoPointFieldsSchema,
		},
	},
	"options": &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
	},
	"min": &schema.Schema{
		Type:     schema.TypeFloat,
		Optional: true,
	},
	"max": &schema.Schema{
		Type:     schema.TypeFloat,
		Optional: true,
	},
	"scale": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	},
	"precision": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	},
	"unit_id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"value": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
}

var commandDefinitionSchema = map[string]*schema.Schema{
	"id": &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	},
	"node_id": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
	"description": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"identifier": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"metadata": &schema.Schema{
		Type:     schema.TypeList,
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
				"unit_id": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"value": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"required": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"type": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
	"arguments": &schema.Schema{
		Type:     schema.TypeList,
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
				"identifier": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"description": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"min_length": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"max_length": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"pattern": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"before": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"after": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"options": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
				},
				"min": &schema.Schema{
					Type:     schema.TypeFloat,
					Optional: true,
				},
				"max": &schema.Schema{
					Type:     schema.TypeFloat,
					Optional: true,
				},
				"scale": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"precision": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"unit_id": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"default_value": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"type": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"required": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
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
}

func tagsInterfaceToStruct(tags interface{}) []Tag {
	tagsStruct := []Tag{}
	if tags != nil {
		for _, tag := range tags.([]interface{}) {
			newTag := Tag{Key: tag.(map[string]interface{})["key"].(string), Value: tag.(map[string]interface{})["value"].(string)}
			tagsStruct = append(tagsStruct, newTag)
		}
	}
	return tagsStruct
}

func tagsStructToInterface(tags []Tag) interface{} {
	if tags != nil {
		tagsList := make([]interface{}, len(tags))
		for i, tag := range tags {
			tagMap := make(map[string]interface{})
			tagMap["key"] = tag.Key
			tagMap["value"] = tag.Value
			tagsList[i] = tagMap
		}

		return tagsList
	}
	return make([]interface{}, 0)
}

func sortStructToInterface(sort []Sort) interface{} {
	sortList := make([]map[string]interface{}, len(sort))
	for i, sortItem := range sort {
		sortMap := make(map[string]interface{})
		sortMap["direction"] = sortItem.Direction
		sortMap["property"] = sortItem.Property
		sortMap["ignore_case"] = sortItem.IgnoreCase
		sortMap["null_handling"] = sortItem.NullHandling
		sortMap["ascending"] = sortItem.Ascending
		sortMap["descending"] = sortItem.Descending
		sortList[i] = sortMap
	}

	return sortList
}

func pageableStructToInterface(pageable Pageable, sort interface{}) interface{} {
	pageableList := make([]map[string]interface{}, 1)
	pageableMap := make(map[string]interface{})
	pageableMap["sort"] = sort
	pageableMap["offset"] = pageable.Offset
	pageableMap["page_number"] = pageable.PageNumber
	pageableMap["page_size"] = pageable.PageSize
	pageableMap["paged"] = pageable.Paged
	pageableMap["unpaged"] = pageable.Unpaged
	pageableList[0] = pageableMap

	return pageableList
}

func fieldStructToInterface(field *Field[interface{}]) map[string]interface{} {
	fieldMap := make(map[string]interface{})
	fieldMap["id"] = field.ID
	fieldMap["name"] = field.Name
	fieldMap["description"] = field.Description
	fieldMap["created_at"] = field.CreatedAt
	fieldMap["created_by"] = field.CreatedBy
	fieldMap["last_modified_at"] = field.LastModifiedAt
	fieldMap["last_modified_by"] = field.LastModifiedBy
	fieldMap["type"] = field.Type
	fieldMap["value"] = strconv.FormatFloat(field.Value.(float64), 'g', -1, 64)

	return fieldMap
}

func propertyStructToInterface(property Property[interface{}]) map[string]interface{} {
	propertyMap := make(map[string]interface{})
	propertyMap["id"] = property.ID
	propertyMap["name"] = property.Name
	propertyMap["description"] = property.Description
	propertyMap["node_id"] = property.NodeId
	propertyMap["created_at"] = property.CreatedAt
	propertyMap["created_by"] = property.CreatedBy
	propertyMap["last_modified_at"] = property.LastModifiedAt
	propertyMap["last_modified_by"] = property.LastModifiedBy
	propertyMap["type"] = property.Type
	propertyMap["tags"] = tagsStructToInterface(property.Tags)
	switch property.Type {
	case "NUMERIC":
		if property.Value != nil {
			propertyMap["value"] = strconv.FormatFloat(property.Value.(float64), 'g', -1, 64)
		}
		propertyMap["min"] = property.Min
		propertyMap["max"] = property.Max
		propertyMap["scale"] = property.Scale
		propertyMap["precision"] = property.Precision
		propertyMap["unit_id"] = property.UnitId
	case "ENUM":
		if property.Value != nil {
			propertyMap["value"] = strconv.FormatFloat(property.Value.(float64), 'g', -1, 64)
		}
		if property.Options != nil {
			propertyMap["options"] = *property.Options
		}
	case "TEXT":
		if property.Value != nil {
			propertyMap["value"] = property.Value.(string)
		}
		propertyMap["min_length"] = property.MinLength
		propertyMap["max_length"] = property.MaxLength
		propertyMap["pattern"] = property.Pattern
	case "TIMESTAMP", "DATE", "TIME":
		if property.Value != nil {
			propertyMap["value"] = property.Value.(string)
		}
		propertyMap["before"] = property.Before
		propertyMap["after"] = property.After
	case "BOOLEAN":
		if property.Value != nil {
			propertyMap["value"] = strconv.FormatBool(property.Value.(bool))
		}
	case "GEOPOINT":
		if property.Fields != nil {
			fieldList := make([]map[string]interface{}, 1)
			fieldMap := make(map[string]interface{})
			elevationList := make([]map[string]interface{}, 1)
			elevationList[0] = fieldStructToInterface(&property.Fields.Elevation)
			fieldMap["elevation"] = elevationList
			latitudeList := make([]map[string]interface{}, 1)
			latitudeList[0] = fieldStructToInterface(&property.Fields.Latitude)
			fieldMap["latitude"] = latitudeList
			longitudeList := make([]map[string]interface{}, 1)
			longitudeList[0] = fieldStructToInterface(&property.Fields.Longitude)
			fieldMap["longitude"] = longitudeList
			fieldList[0] = fieldMap
			propertyMap["fields"] = fieldList
		}
	}
	return propertyMap
}

func commandDefinitionStructToInterface(commandDefinition CommandDefinition) map[string]interface{} {
	commandDefinitionMap := make(map[string]interface{})
	commandDefinitionMap["id"] = commandDefinition.ID
	commandDefinitionMap["node_id"] = commandDefinition.NodeId
	commandDefinitionMap["name"] = commandDefinition.Name
	commandDefinitionMap["description"] = commandDefinition.Description
	commandDefinitionMap["identifier"] = commandDefinition.Identifier
	commandDefinitionMap["created_at"] = commandDefinition.CreatedAt
	commandDefinitionMap["created_by"] = commandDefinition.CreatedBy
	commandDefinitionMap["last_modified_at"] = commandDefinition.LastModifiedAt
	commandDefinitionMap["last_modified_by"] = commandDefinition.LastModifiedBy
	if commandDefinition.Metadata != nil {
		commandDefinitionMap["metadata"] = make([]interface{}, len(commandDefinition.Metadata))
		for i, metadata := range commandDefinition.Metadata {
			metadataMap := make(map[string]interface{})
			metadataMap["id"] = metadata.ID
			metadataMap["name"] = metadata.Name
			metadataMap["description"] = metadata.Description
			metadataMap["unit_id"] = metadata.UnitId
			metadataMap["value"] = metadata.Value
			metadataMap["required"] = metadata.Required
			metadataMap["type"] = metadata.Type
			commandDefinitionMap["metadata"].([]interface{})[i] = metadataMap
		}
	}

	if commandDefinition.Arguments != nil {
		commandDefinitionMap["arguments"] = make([]interface{}, len(commandDefinition.Arguments))
		for i, argument := range commandDefinition.Arguments {
			argumentMap := make(map[string]interface{})
			argumentMap["id"] = argument.ID
			argumentMap["name"] = argument.Name
			argumentMap["identifier"] = argument.Identifier
			argumentMap["description"] = argument.Description
			argumentMap["type"] = argument.Type
			argumentMap["required"] = argument.Required
			switch argument.Type {
			case "NUMERIC":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = strconv.FormatFloat(argument.DefaultValue.(float64), 'g', -1, 64)
				}
				argumentMap["min"] = argument.Min
				argumentMap["max"] = argument.Max
				argumentMap["scale"] = argument.Scale
				argumentMap["precision"] = argument.Precision
				argumentMap["unit_id"] = argument.UnitId
			case "ENUM":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = strconv.FormatFloat(argument.DefaultValue.(float64), 'g', -1, 64)
				}
				if argument.Options != nil {
					argumentMap["options"] = *argument.Options
				}
			case "TEXT":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = argument.DefaultValue.(string)
				}
				argumentMap["min_length"] = argument.MinLength
				argumentMap["max_length"] = argument.MaxLength
				argumentMap["pattern"] = argument.Pattern
			case "TIMESTAMP", "DATE", "TIME":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = argument.DefaultValue.(string)
				}
				argumentMap["before"] = argument.Before
				argumentMap["after"] = argument.After
			case "BOOLEAN":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = strconv.FormatBool(argument.DefaultValue.(bool))
				}
			}
			commandDefinitionMap["arguments"].([]interface{})[i] = argumentMap
		}
	}

	return commandDefinitionMap
}
