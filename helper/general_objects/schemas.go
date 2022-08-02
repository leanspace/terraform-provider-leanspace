package general_objects

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func PaginatedListSchema(content map[string]*schema.Schema) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: content,
			},
		},
		"total_elements": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"total_pages": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"number_of_elements": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"number": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"size": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"sort": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: SortSchema,
			},
		},
		"first": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"last": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"empty": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"pageable": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: PageableSchema,
			},
		},
		"filters": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

var SortSchema = map[string]*schema.Schema{
	"direction": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"property": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"ignore_case": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"null_handling": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"ascending": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"descending": {
		Type:     schema.TypeBool,
		Computed: true,
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
		Type:     schema.TypeInt,
		Computed: true,
	},
	"page_number": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"page_size": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"paged": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"unpaged": {
		Type:     schema.TypeBool,
		Computed: true,
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

var validAttributeSchemaTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM",
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
	for _, value := range validAttributeSchemaTypes {
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
		},
		"max_length": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntAtLeast(1),
		},
		"pattern": {
			Type:     schema.TypeString,
			Optional: true,
		},

		// Numeric only
		"min": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"max": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"scale": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"precision": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"unit_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsUUID,
		},

		// Time, date, timestamp only
		"before": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsRFC3339Time,
		},
		"after": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsRFC3339Time,
		},

		// Enum only
		"options": {
			Type:     schema.TypeMap,
			Optional: true,
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
	},
	// Numeric only
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
}
