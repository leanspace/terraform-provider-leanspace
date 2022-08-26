package general_objects

import (
	"leanspace-terraform-provider/helper"

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
	if filters == nil {
		filters = map[string]*schema.Schema{}
	}
	filters["ids"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	filters["query"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	filters["page"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  0,
	}
	filters["size"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  100,
	}
	filters["sort"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	return filters
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

var TagsSchema = &schema.Schema{
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
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM", "BINARY",
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func DefinitionAttributeSchema(excludeTypes []string, excludeFields []string) map[string]*schema.Schema {
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
			Type:     schema.TypeString,
			Optional: true,
		},

		// Text only
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
			ValidateFunc: validation.IsRFC3339Time,
			Description:  "Time/date/timestamp only: Maximum date allowed",
		},
		"after": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsRFC3339Time,
			Description:  "Time/date/timestamp only: Minimum date allowed",
		},

		// Enum only
		"options": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "Enum only: The allowed values for the enum in the format 1 = \"value\"",
		},
	}

	for _, field := range excludeFields {
		delete(schema, field)
	}

	return schema
}

var validMetadataTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM",
}

var ValueAttributeSchema = map[string]*schema.Schema{
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validMetadataTypes, false),
		Description:  helper.AllowedValuesToDescription(validMetadataTypes),
	},
	// Numeric only
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
}
