package general_objects

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func PaginatedListSchema(content map[string]*schema.Schema, filters map[string]*schema.Schema) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: content,
			},
		},
		"total_elements": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of elements in total",
		},
		"total_pages": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of pages in total",
		},
		"number_of_elements": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of elements fetched in this page",
		},
		"number": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Page number",
		},
		"size": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Size of this page",
		},
		"sort": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: SortSchema,
			},
		},
		"first": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True if this is the first page",
		},
		"last": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True if this is the last page",
		},
		"empty": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True if the content is empty",
		},
		"pageable": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: PageableSchema,
			},
		},
		"filters": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: FilterSchema(filters),
			},
		},
	}
}

func FilterSchema(filters map[string]*schema.Schema) map[string]*schema.Schema {
	baseFilter := map[string]*schema.Schema{
		"ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"query": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"page": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"sort": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	for key, value := range filters {
		baseFilter[key] = value
	}

	return baseFilter
}

var SortSchema = map[string]*schema.Schema{
	"direction": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Direction of the sorting, either DESC or ASC",
	},
	"property": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Property used to sort by",
	},
	"ignore_case": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "True if the search ignores case",
	},
	"null_handling": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "How null values are handled",
	},
	"ascending": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "True if the direction of the sorting is ascending",
	},
	"descending": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "True if the direction of the sorting is descending",
	},
}

var PageableSchema = map[string]*schema.Schema{
	"sort": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: SortSchema,
		},
	},
	"offset": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Number of elements in previous pages",
	},
	"page_number": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Page number",
	},
	"page_size": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Size of this page",
	},
	"paged": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "True if this query is paged",
	},
	"unpaged": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "True if this query is unpaged",
	},
}

var geoPointFieldsDefSchema = map[string]*schema.Schema{
	"latitude": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: attributeFieldDefSchema,
		},
	},
	"longitude": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: attributeFieldDefSchema,
		},
	},
	"elevation": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: attributeFieldDefSchema,
		},
	},
}

var attributeFieldDefSchema = map[string]*schema.Schema{
	"default_value": {
		Type:     schema.TypeString,
		Optional: true,
	},

	// Numeric only
	"scale": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Property field with numeric type only: the scale required.",
	},
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Property field with numeric type only",
	},
	"min": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Property field with numeric type only: the minimum value allowed.",
	},
	"precision": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Property field with numeric type only: How many values after the comma should be accepted",
	},
	"max": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Property field with numeric type only: the maximum value allowed.",
	},
}

var geoPointFieldsSchema = map[string]*schema.Schema{
	"latitude": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: attributeFieldSchema,
		},
	},
	"longitude": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: attributeFieldSchema,
		},
	},
	"elevation": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: attributeFieldSchema,
		},
	},
}

var attributeFieldSchema = map[string]*schema.Schema{
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},

	// Numeric only
	"scale": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Property field with numeric type only: the scale required.",
	},
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Property field with numeric type only",
	},
	"min": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Property field with numeric type only: the minimum value allowed.",
	},
	"precision": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Property field with numeric type only: How many values after the comma should be accepted",
	},
	"max": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Property field with numeric type only: the maximum value allowed.",
	},
}

var KeyValuesSchema = &schema.Schema{
	Type:     schema.TypeSet,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}

var ValidAttributeSchemaTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM", "BINARY", "ARRAY", "TLE", "GEOPOINT", "STRUCTURE",
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func DefinitionAttributeSchema(excludeTypes []string, excludeFields []string, forceNew bool) map[string]*schema.Schema {
	validTypes := []string{}
	for _, value := range ValidAttributeSchemaTypes {
		if contains(excludeTypes, value) {
			continue
		}
		validTypes = append(validTypes, value)
	}

	schema := map[string]*schema.Schema{
		// Common fields
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(validTypes, false),
			Description:  helper.AllowedValuesToDescription(validTypes),
		},
		"required": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"default_value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The default value can be of any type. In case of an array type, please surround the list values with double quotes and use the comma separator.",
		},
		// Text & Binary
		"min_length": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntAtLeast(1),
			Description:  "Text only: Minimum length of this text (at least 1)",
		},
		"max_length": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntAtLeast(1),
			Description:  "Text only: Maximum length of this text (at least 1)",
		},

		// Text only
		"pattern": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Text only: Regex defined the allowed pattern of this text",
		},

		// Numeric only
		"min": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Numeric only",
		},
		"max": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Numeric only",
		},
		"scale": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Numeric only",
		},
		"precision": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Numeric only: How many values after the comma should be accepted",
		},
		"unit_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsUUID,
			Description:  "Numeric only",
		},

		// Time, date, timestamp only
		"before": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: helper.IsValidTimeDateOrTimestamp,
			Description:  "Time/date/timestamp only: Maximum date allowed",
		},
		"after": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: helper.IsValidTimeDateOrTimestamp,
			Description:  "Time/date/timestamp only: Minimum date allowed",
		},

		// Enum only
		"options": {
			Type: schema.TypeMap,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "Enum only: The allowed values for the enum in the format 1 = \"value\"",
		},

		// Geopoint only
		"fields": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: geoPointFieldsDefSchema,
			},
			Description: "Geopoint only",
		},

		// Array
		"min_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Array only: The minimum number of elements allowed",
		},
		"max_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Array only: The maximum number of elements allowed",
		},
		"unique": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Array only: No duplicated elements are allowed",
		},
		"constraint": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Array only: Constraint applied to all elements in the array",
			MaxItems:    1,
			MinItems:    1,
			Elem: &schema.Resource{
				Schema: DefinitionAttributeArrayConstraintSchema(
					[]string{"ARRAY", "STRUCTURE", "GEOPOINT", "TLE"}, // element types not allowed in array
					[]string{"default_value"},                         // Field unused as only the default value of the array is taken into account
				),
			},
		},
	}

	if forceNew {
		schema["type"].ForceNew = true
	}

	for _, field := range excludeFields {
		delete(schema, field)
	}

	return schema
}

func DefinitionAttributeArrayConstraintSchema(excludeTypes []string, excludeFields []string) map[string]*schema.Schema {
	validTypes := []string{}
	for _, value := range ValidAttributeSchemaTypes {
		if contains(excludeTypes, value) {
			continue
		}
		validTypes = append(validTypes, value)
	}

	schema := map[string]*schema.Schema{
		// Common fields
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(validTypes, false),
			Description:  helper.AllowedValuesToDescription(validTypes),
		},
		"required": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		// Text & Binary
		"max_length": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntAtLeast(1),
			Description:  "Only array elements with text type: Maximum length of this text (at least 1)",
		},
		"min_length": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntAtLeast(1),
			Description:  "Only array elements with text type: Minimum length of this text (at least 1)",
		},

		// Text only
		"pattern": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Only array elements with text type: Regex defined the allowed pattern of this text",
		},

		// Numeric only
		"max": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Only array elements with numeric type : maximum value allowed",
		},
		"precision": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Only array elements with numeric type : how many values after the comma should be accepted",
		},
		"min": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Only array elements with numeric type : minimum value allowed",
		},
		"unit_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsUUID,
			Description:  "Only array elements with numeric type",
		},
		"scale": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Only array elements with numeric type",
		},

		// Enum only
		"options": {
			Type: schema.TypeMap,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "Only array elements with enum type : The allowed values for the enum in the format 1 = \"value\"",
		},

		// Time, date, timestamp only
		"after": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: helper.IsValidTimeDateOrTimestamp,
			Description:  "Only array elements with time/date/timestamp type : Minimum date allowed",
		},
		"before": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: helper.IsValidTimeDateOrTimestamp,
			Description:  "Only array elements with time/date/timestamp type : Maximum date allowed",
		},
	}

	for _, field := range excludeFields {
		delete(schema, field)
	}

	return schema
}

var validMetadataTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM", "BINARY", "ARRAY", "TLE", "GEOPOINT", "STRUCTURE",
}

var validArraydataTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM", "BINARY",
}

func ValueAttributeSchema(excludeTypes []string) map[string]*schema.Schema {
	validTypes := []string{}
	for _, value := range validMetadataTypes {
		if contains(excludeTypes, value) {
			continue
		}
		validTypes = append(validTypes, value)
	}

	schema := map[string]*schema.Schema{
		"value": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(validTypes, false),
			Description:  helper.AllowedValuesToDescription(validTypes),
		},
		"data_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(validArraydataTypes, false),
			Description:  helper.AllowedValuesToDescription(validArraydataTypes),
		},
		// Numeric only
		"unit_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsUUID,
		},
		// Geopoint only
		"fields": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: geoPointFieldsSchema,
			},
			Description: "Geopoint only",
		},
	}

	return schema
}
