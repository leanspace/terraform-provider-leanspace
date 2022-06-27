package asset

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var nodeSchema = map[string]*schema.Schema{
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
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"type": &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringInSlice([]string{"ASSET", "GROUP", "COMPONENT"}, false),
	},
	"kind": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
	},
	"tags": tagsSchema,
	"norad_id": &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d{5}$`), "It must be 5 digits"),
	},
	"international_designator": &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d{4}-|\d{2})[0-9]{3}[A-Za-z]{0,3}$`), ""),
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
}

var rootNodeSchema = map[string]*schema.Schema{
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
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d{5}$`), "It must be 5 digits"),
	},
	"international_designator": &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d{4}-|\d{2})[0-9]{3}[A-Za-z]{0,3}$`), ""),
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
}

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
				Required: true,
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
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice([]string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN", "GEOPOINT"}, false),
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
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
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
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
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
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsRFC3339Time,
	},
	"after": &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsRFC3339Time,
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
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"value": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringInSlice([]string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN", "GEOPOINT"}, false),
	},
}

var commandDefinitionSchema = map[string]*schema.Schema{
	"id": &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	},
	"node_id": &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
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
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.IsUUID,
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
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice([]string{"NUMERIC", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN"}, false),
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
					Type:         schema.TypeInt,
					Optional:     true,
					ValidateFunc: validation.IntAtLeast(1),
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
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.IsRFC3339Time,
				},
				"after": &schema.Schema{
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.IsRFC3339Time,
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
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.IsUUID,
				},
				"default_value": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"type": &schema.Schema{
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice([]string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN"}, false),
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
			metadataMap["type"] = metadata.Type
			switch metadata.Type {
			case "NUMERIC":
				if metadata.Value != nil {
					metadataMap["value"] = strconv.FormatFloat(metadata.Value.(float64), 'g', -1, 64)
				}
			case "TEXT":
				if metadata.Value != nil {
					metadataMap["value"] = metadata.Value.(string)
				}
			case "TIMESTAMP", "DATE", "TIME":
				if metadata.Value != nil {
					metadataMap["value"] = metadata.Value.(string)
				}
			case "BOOLEAN":
				if metadata.Value != nil {
					metadataMap["value"] = strconv.FormatBool(metadata.Value.(bool))
				}
			}
			metadataMap["required"] = metadata.Required
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

func nodeStructToInterfaceBase(node Node) map[string]interface{} {
	return nodeStructToInterface(&node, 0)
}

func nodeStructToInterface(node *Node, level int) map[string]interface{} {
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
	// Here we seemed to lock going to depth > 1?
	// Any reason for this ? Would need to ask @Gerome
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

var tle1stLine = `^1 (?P<noradId>[ 0-9]{5})[A-Z] [ 0-9]{5}[ A-Z]{3} [ 0-9]{5}[.][ 0-9]{8} (?:(?:[ 0+-][.][ 0-9]{8})|(?: [ +-][.][ 0-9]{7})) [ +-][ 0-9]{5}[+-][ 0-9] [ +-][ 0-9]{5}[+-][ 0-9] [ 0-9] [ 0-9]{4}[ 0-9]$`
var tle2ndLine = `^2 (?P<noradId>[ 0-9]{5}) [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{7} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{2}[.][ 0-9]{13}[ 0-9]$`

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
		matched, _ := regexp.MatchString(tle1stLine, node["tle"].([]interface{})[0].(string))
		if !matched {
			return nodeStruct, fmt.Errorf("TLE first line mutch match %q, got: %q", tle1stLine, node["tle"].([]interface{})[0].(string))
		}
		matched, _ = regexp.MatchString(tle2ndLine, node["tle"].([]interface{})[1].(string))
		if !matched {
			return nodeStruct, fmt.Errorf("TLE second line mutch match %q, got: %q", tle2ndLine, node["tle"].([]interface{})[1].(string))
		}
		for i, tle := range node["tle"].([]interface{}) {
			nodeStruct.Tle[i] = tle.(string)
		}

	}

	return nodeStruct, nil
}

func fieldInterfaceToStruct(fieldList []interface{}) Field[interface{}] {
	field := fieldList[0].(map[string]interface{})
	fieldStruct := Field[interface{}]{}
	fieldStruct.ID = field["id"].(string)
	fieldStruct.Name = field["name"].(string)
	fieldStruct.Description = field["description"].(string)
	fieldStruct.CreatedAt = field["created_at"].(string)
	fieldStruct.CreatedBy = field["created_by"].(string)
	fieldStruct.LastModifiedAt = field["last_modified_at"].(string)
	fieldStruct.LastModifiedBy = field["last_modified_by"].(string)
	fieldStruct.Type = field["type"].(string)
	fieldStruct.Value = field["value"]

	return fieldStruct
}

func getPropertyData(property map[string]interface{}) (Property[interface{}], error) {
	propertyMap := Property[interface{}]{}

	propertyMap.Name = property["name"].(string)
	propertyMap.Description = property["description"].(string)
	propertyMap.NodeId = property["node_id"].(string)
	propertyMap.CreatedAt = property["created_at"].(string)
	propertyMap.CreatedBy = property["created_by"].(string)
	propertyMap.LastModifiedAt = property["last_modified_at"].(string)
	propertyMap.LastModifiedBy = property["last_modified_by"].(string)
	propertyMap.Value = property["value"]
	propertyMap.Type = property["type"].(string)
	propertyMap.Tags = tagsInterfaceToStruct(property["tags"])
	switch propertyMap.Type {
	case "NUMERIC":
		propertyMap.Min = property["min"].(float64)
		propertyMap.Max = property["max"].(float64)
		propertyMap.Scale = property["scale"].(int)
		propertyMap.Precision = property["precision"].(int)
		propertyMap.UnitId = property["unit_id"].(string)
	case "ENUM":
		if property["options"] != nil {
			option := property["options"].(map[string]interface{})
			propertyMap.Options = &option
		}
	case "TEXT":
		propertyMap.MinLength = property["min_length"].(int)
		propertyMap.MaxLength = property["max_length"].(int)
		propertyMap.Pattern = property["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		propertyMap.Before = property["before"].(string)
		propertyMap.After = property["after"].(string)
	case "BOOLEAN":
		// no extra property for booleans
	case "GEOPOINT":
		if property["fields"] != nil {
			propertyMap.Fields = &Fields{}
			propertyMap.Fields.Elevation = fieldInterfaceToStruct(property["fields"].([]interface{})[0].(map[string]interface{})["elevation"].([]interface{}))
			propertyMap.Fields.Latitude = fieldInterfaceToStruct(property["fields"].([]interface{})[0].(map[string]interface{})["latitude"].([]interface{}))
			propertyMap.Fields.Longitude = fieldInterfaceToStruct(property["fields"].([]interface{})[0].(map[string]interface{})["longitude"].([]interface{}))
		}
	}

	return propertyMap, nil
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
	case "BOOLEAN":
		// no extra field
	}
	argumentStruct.DefaultValue = argument["default_value"]

	return argumentStruct
}

func getCommandDefinitionData(commandDefinition map[string]interface{}) (CommandDefinition, error) {
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

	return commandDefinitionMap, nil
}
