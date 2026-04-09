package general_objects

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"

	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
)

func PaginatedListSchemaDS(content map[string]datasourceschema.Attribute, filters map[string]datasourceschema.Attribute) map[string]datasourceschema.Attribute {
	return map[string]datasourceschema.Attribute{
		"id": datasourceschema.StringAttribute{
			Computed: true,
		},
		"content": datasourceschema.ListNestedAttribute{
			Computed: true,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: content,
			},
		},
		"total_elements": datasourceschema.Int64Attribute{
			Computed:    true,
			Description: "Number of elements in total",
		},
		"total_pages": datasourceschema.Int64Attribute{
			Computed:    true,
			Description: "Number of pages in total",
		},
		"number_of_elements": datasourceschema.Int64Attribute{
			Computed:    true,
			Description: "Number of elements fetched in this page",
		},
		"number": datasourceschema.Int64Attribute{
			Computed:    true,
			Description: "Page number",
		},
		"size": datasourceschema.Int64Attribute{
			Computed:    true,
			Description: "Size of this page",
		},
		"sort": datasourceschema.ListNestedAttribute{
			Computed: true,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: SortSchemaDS,
			},
		},
		"first": datasourceschema.BoolAttribute{
			Computed:    true,
			Description: "True if this is the first page",
		},
		"last": datasourceschema.BoolAttribute{
			Computed:    true,
			Description: "True if this is the last page",
		},
		"empty": datasourceschema.BoolAttribute{
			Computed:    true,
			Description: "True if the content is empty",
		},
		"pageable": datasourceschema.SingleNestedAttribute{
			Computed:   true,
			Attributes: PageableSchemaDS,
		},
		"filters": datasourceschema.SingleNestedAttribute{
			Optional:   true,
			Attributes: FilterSchemaDS(filters),
		},
	}
}

func FilterSchemaDS(filters map[string]datasourceschema.Attribute) map[string]datasourceschema.Attribute {
	baseFilter := map[string]datasourceschema.Attribute{
		"ids": datasourceschema.ListAttribute{
			Optional:    true,
			ElementType: types.StringType,
		},
		"query": datasourceschema.StringAttribute{
			Optional: true,
		},
		"page": datasourceschema.Int64Attribute{
			Optional: true,
		},
		"size": datasourceschema.Int64Attribute{
			Optional: true,
		},
		"sort": datasourceschema.ListAttribute{
			Optional:    true,
			ElementType: types.StringType,
		},
	}

	for key, value := range filters {
		baseFilter[key] = value
	}

	return baseFilter
}

func AuditFilterFieldsWithTagsAndSingularBy(filters map[string]datasourceschema.Attribute) map[string]datasourceschema.Attribute {
	return AuditFilterFields(filters, true, true)
}

func AuditFilterFieldsWithTags(filters map[string]datasourceschema.Attribute) map[string]datasourceschema.Attribute {
	return AuditFilterFields(filters, true, false)
}

// AuditFilterFields returns the standard audit filter attributes: created_bys, last_modified_bys,
// from_created_at, to_created_at, from_last_modified_at, to_last_modified_at, and optionally tags. If singularBy is true, it returns created_by and last_modified_by instead of their plural version.
func AuditFilterFields(filters map[string]datasourceschema.Attribute, includeTags bool, singularBy bool) map[string]datasourceschema.Attribute {
	baseFilter := map[string]datasourceschema.Attribute{
		"from_created_at": datasourceschema.StringAttribute{
			Optional:    true,
			Description: "Filter on the creation date. Entries with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
		},
		"to_created_at": datasourceschema.StringAttribute{
			Optional:    true,
			Description: "Filter on the creation date. Entries with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
		},
		"from_last_modified_at": datasourceschema.StringAttribute{
			Optional:    true,
			Description: "Filter on the last modification date. Entries with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
		},
		"to_last_modified_at": datasourceschema.StringAttribute{
			Optional:    true,
			Description: "Filter on the last modification date. Entries with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
		},
	}
	if singularBy {
		baseFilter["last_modified_by"] = datasourceschema.StringAttribute{
			Optional:    true,
			Description: "Filter on the user who last modified the entry. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
		}
		baseFilter["created_by"] = datasourceschema.StringAttribute{
			Optional:    true,
			Description: "Filter on the user who created the Node. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
		}
	} else {
		baseFilter["last_modified_bys"] = datasourceschema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "Filter on the user who last modified the entry. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
		}
		baseFilter["created_bys"] = datasourceschema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "Filter on the user who created the entry. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
		}
	}

	if includeTags {
		baseFilter["tags"] = datasourceschema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
		}
	}
	for k, v := range filters {
		baseFilter[k] = v
	}
	return baseFilter
}

var SortSchemaDS = map[string]datasourceschema.Attribute{
	"direction": datasourceschema.StringAttribute{
		Computed:    true,
		Description: "Direction of the sorting, either DESC or ASC",
	},
	"property": datasourceschema.StringAttribute{
		Computed:    true,
		Description: "Property used to sort by",
	},
	"ignore_case": datasourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if the search ignores case",
	},
	"null_handling": datasourceschema.StringAttribute{
		Computed:    true,
		Description: "How null values are handled",
	},
	"ascending": datasourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if the direction of the sorting is ascending",
	},
	"descending": datasourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if the direction of the sorting is descending",
	},
}

var PageableSchemaDS = map[string]datasourceschema.Attribute{
	"sort": datasourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: datasourceschema.NestedAttributeObject{
			Attributes: SortSchemaDS,
		},
	},
	"offset": datasourceschema.Int64Attribute{
		Computed:    true,
		Description: "Number of elements in previous pages",
	},
	"page_number": datasourceschema.Int64Attribute{
		Computed:    true,
		Description: "Page number",
	},
	"page_size": datasourceschema.Int64Attribute{
		Computed:    true,
		Description: "Size of this page",
	},
	"paged": datasourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if this query is paged",
	},
	"unpaged": datasourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if this query is unpaged",
	},
}

var SortSchemaR = map[string]resourceschema.Attribute{
	"direction": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Direction of the sorting, either DESC or ASC",
	},
	"property": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Property used to sort by",
	},
	"ignore_case": resourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if the search ignores case",
	},
	"null_handling": resourceschema.StringAttribute{
		Computed:    true,
		Description: "How null values are handled",
	},
	"ascending": resourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if the direction of the sorting is ascending",
	},
	"descending": resourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if the direction of the sorting is descending",
	},
}

var PageableSchemaR = map[string]resourceschema.Attribute{
	"sort": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: SortSchemaR,
		},
	},
	"offset": resourceschema.Int64Attribute{
		Computed:    true,
		Description: "Number of elements in previous pages",
	},
	"page_number": resourceschema.Int64Attribute{
		Computed:    true,
		Description: "Page number",
	},
	"page_size": resourceschema.Int64Attribute{
		Computed:    true,
		Description: "Size of this page",
	},
	"paged": resourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if this query is paged",
	},
	"unpaged": resourceschema.BoolAttribute{
		Computed:    true,
		Description: "True if this query is unpaged",
	},
}

func CreateGeoPointFieldsSchema(isValueField bool) map[string]resourceschema.Attribute {
	return map[string]resourceschema.Attribute{
		"latitude": resourceschema.SingleNestedAttribute{
			Optional:   true,
			Attributes: baseAttributeFieldSchema(isValueField, true),
		},
		"longitude": resourceschema.SingleNestedAttribute{
			Optional:   true,
			Attributes: baseAttributeFieldSchema(isValueField, true),
		},
		"elevation": resourceschema.SingleNestedAttribute{
			Optional:   true,
			Attributes: baseAttributeFieldSchema(isValueField, true),
		},
	}
}

func createGeoPointFieldsSchemaDS(isValueField bool) map[string]datasourceschema.Attribute {
	return map[string]datasourceschema.Attribute{
		"latitude": datasourceschema.SingleNestedAttribute{
			Computed:   true,
			Attributes: baseAttributeFieldSchemaDS(isValueField),
		},
		"longitude": datasourceschema.SingleNestedAttribute{
			Computed:   true,
			Attributes: baseAttributeFieldSchemaDS(isValueField),
		},
		"elevation": datasourceschema.SingleNestedAttribute{
			Computed:   true,
			Attributes: baseAttributeFieldSchemaDS(isValueField),
		},
	}
}

var geoPointFieldsDefSchema = CreateGeoPointFieldsSchema(false)
var geoPointFieldsSchema = CreateGeoPointFieldsSchema(true)
var geoPointFieldsDefSchemaDS = createGeoPointFieldsSchemaDS(false)
var geoPointFieldsSchemaDS = createGeoPointFieldsSchemaDS(true)

func baseAttributeFieldSchema(isValueField bool, isGeoPoint bool) map[string]resourceschema.Attribute {
	baseSchema := map[string]resourceschema.Attribute{
		"scale": resourceschema.Int64Attribute{
			Computed:    isGeoPoint,
			Optional:    !isGeoPoint,
			Description: "Property field with numeric type only: the scale required.",
		},
		"unit_id": resourceschema.StringAttribute{
			Computed:    isGeoPoint,
			Optional:    !isGeoPoint,
			Description: "Property field with numeric type only",
			Validators:  helper.ValidUUID(),
		},
		"min": resourceschema.Float64Attribute{
			Computed:    isGeoPoint,
			Optional:    !isGeoPoint,
			Description: "Property field with numeric type only: the minimum value allowed.",
		},
		"precision": resourceschema.Int64Attribute{
			Computed:    isGeoPoint,
			Optional:    !isGeoPoint,
			Description: "Property field with numeric type only: How many values after the comma should be accepted",
		},
		"max": resourceschema.Float64Attribute{
			Computed:    isGeoPoint,
			Optional:    !isGeoPoint,
			Description: "Property field with numeric type only: the maximum value allowed.",
		},
	}

	if isValueField {
		baseSchema["value"] = resourceschema.StringAttribute{
			Optional: true,
		}
	} else {
		baseSchema["default_value"] = resourceschema.StringAttribute{
			Optional: true,
		}
	}

	return baseSchema
}

func baseAttributeFieldSchemaDS(isValueField bool) map[string]datasourceschema.Attribute {
	baseSchema := map[string]datasourceschema.Attribute{
		"scale": datasourceschema.Int64Attribute{
			Computed:    true,
			Description: "Property field with numeric type only: the scale required.",
		},
		"unit_id": datasourceschema.StringAttribute{
			Computed:    true,
			Description: "Property field with numeric type only",
		},
		"precision": datasourceschema.Int64Attribute{
			Computed:    true,
			Description: "Property field with numeric type only: How many values after the comma should be accepted",
		},
		"min": datasourceschema.Float64Attribute{
			Computed:    true,
			Description: "Property field with numeric type only: the minimum value allowed.",
		},
		"max": datasourceschema.Float64Attribute{
			Computed:    true,
			Description: "Property field with numeric type only: the maximum value allowed.",
		},
	}

	if isValueField {
		baseSchema["value"] = datasourceschema.StringAttribute{
			Computed: true,
		}
	} else {
		baseSchema["default_value"] = datasourceschema.StringAttribute{
			Computed: true,
		}
	}

	return baseSchema
}

// KeyValuesSchema returns a SetNestedAttribute for key-value tags (resource schema).
var KeyValuesSchema = resourceschema.SetNestedAttribute{
	Optional: true,
	NestedObject: resourceschema.NestedAttributeObject{
		Attributes: map[string]resourceschema.Attribute{
			"key": resourceschema.StringAttribute{
				Required: true,
			},
			"value": resourceschema.StringAttribute{
				Optional: true,
			},
		},
	},
}

// KeyValuesSchemaDS returns a SetNestedAttribute for key-value tags (data source schema).
var KeyValuesSchemaDS = datasourceschema.SetNestedAttribute{
	Computed: true,
	NestedObject: datasourceschema.NestedAttributeObject{
		Attributes: map[string]datasourceschema.Attribute{
			"key": datasourceschema.StringAttribute{
				Computed: true,
			},
			"value": datasourceschema.StringAttribute{
				Computed: true,
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

func DefinitionAttributeSchema(excludeTypes []string, excludeFields []string, forceNew bool) map[string]resourceschema.Attribute {
	validTypes := []string{}
	for _, value := range ValidAttributeSchemaTypes {
		if contains(excludeTypes, value) {
			continue
		}
		validTypes = append(validTypes, value)
	}

	var typePlanModifiers []planmodifier.String
	if forceNew {
		typePlanModifiers = []planmodifier.String{stringplanmodifier.RequiresReplace()}
	}

	attribute := map[string]resourceschema.Attribute{
		// Common fields
		"type": resourceschema.StringAttribute{
			Required:      true,
			Description:   helper.AllowedValuesToDescription(validTypes),
			Validators:    []validator.String{stringvalidator.OneOf(validTypes...)},
			PlanModifiers: typePlanModifiers,
		},
		"required": resourceschema.BoolAttribute{
			Optional: true,
			Computed: true,
			Default:  booldefault.StaticBool(false),
		},
		"default_value": resourceschema.StringAttribute{
			Optional:    true,
			Description: "The default value can be of any type. In case of an array type, please surround the list values with double quotes and use the comma separator.",
		},
		// Text & Binary
		"min_length": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Text only: Minimum length of this text (at least 1)",
			Validators:  []validator.Int64{int64validator.AtLeast(1)},
		},
		"max_length": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Text only: Maximum length of this text (at least 1)",
			Validators:  []validator.Int64{int64validator.AtLeast(1)},
		},
		// Text only
		"pattern": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Text only: Regex defined the allowed pattern of this text",
		},
		// Numeric only
		"min": resourceschema.Float64Attribute{
			Optional:    true,
			Description: "Numeric only",
		},
		"max": resourceschema.Float64Attribute{
			Optional:    true,
			Description: "Numeric only",
		},
		"scale": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Numeric only",
		},
		"precision": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Numeric only: How many values after the comma should be accepted",
		},
		"unit_id": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Numeric only",
			Validators:  helper.ValidUUID(),
		},
		// Time, date, timestamp only
		"before": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Time/date/timestamp only: Maximum date allowed",
			Validators:  helper.IsValidTimeDateOrTimestamp(),
		},
		"after": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Time/date/timestamp only: Minimum date allowed",
			Validators:  helper.IsValidTimeDateOrTimestamp(),
		},
		// Enum only
		"options": resourceschema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "Enum only: The allowed values for the enum in the format 1 = \"value\"",
		},
		// Geopoint only
		"fields": resourceschema.SingleNestedAttribute{
			Optional:    true,
			Attributes:  geoPointFieldsDefSchema,
			Description: "Geopoint only",
		},
		// Array
		"min_size": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Array only: The minimum number of elements allowed",
		},
		"max_size": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Array only: The maximum number of elements allowed",
		},
		"unique": resourceschema.BoolAttribute{
			Optional:    true,
			Description: "Array only: No duplicated elements are allowed",
		},
		"constraint": resourceschema.SingleNestedAttribute{
			Optional:    true,
			Description: "Array only: Constraint applied to all elements in the array",
			Attributes: DefinitionAttributeArrayConstraintSchema(
				[]string{"ARRAY", "STRUCTURE", "GEOPOINT", "TLE"},
				[]string{"default_value"},
			),
		},
	}

	for _, field := range excludeFields {
		delete(attribute, field)
	}

	return attribute
}

// DefinitionAttributeSchemaDS generates the datasource schema for typed definition attributes (all computed).
func DefinitionAttributeSchemaDS(excludeTypes []string, excludeFields []string) map[string]datasourceschema.Attribute {
	s := map[string]datasourceschema.Attribute{
		"type":          datasourceschema.StringAttribute{Computed: true},
		"required":      datasourceschema.BoolAttribute{Computed: true},
		"default_value": datasourceschema.StringAttribute{Computed: true},
		"min_length":    datasourceschema.Int64Attribute{Computed: true},
		"max_length":    datasourceschema.Int64Attribute{Computed: true},
		"pattern":       datasourceschema.StringAttribute{Computed: true},
		"min":           datasourceschema.Float64Attribute{Computed: true},
		"max":           datasourceschema.Float64Attribute{Computed: true},
		"scale":         datasourceschema.Int64Attribute{Computed: true},
		"precision":     datasourceschema.Int64Attribute{Computed: true},
		"unit_id":       datasourceschema.StringAttribute{Computed: true},
		"before":        datasourceschema.StringAttribute{Computed: true},
		"after":         datasourceschema.StringAttribute{Computed: true},
		"options": datasourceschema.MapAttribute{
			ElementType: types.StringType,
			Computed:    true,
		},
		"fields": datasourceschema.SingleNestedAttribute{
			Computed:   true,
			Attributes: geoPointFieldsDefSchemaDS,
		},
		"min_size": datasourceschema.Int64Attribute{Computed: true},
		"max_size": datasourceschema.Int64Attribute{Computed: true},
		"unique":   datasourceschema.BoolAttribute{Computed: true},
		"constraint": datasourceschema.SingleNestedAttribute{
			Computed: true,
			Attributes: DefinitionAttributeArrayConstraintSchemaDS(
				[]string{"ARRAY", "STRUCTURE", "GEOPOINT", "TLE"},
				[]string{"default_value"},
			),
		},
	}

	for _, field := range excludeFields {
		delete(s, field)
	}

	return s
}

func DefinitionAttributeArrayConstraintSchema(excludeTypes []string, excludeFields []string) map[string]resourceschema.Attribute {
	validTypes := []string{}
	for _, value := range ValidAttributeSchemaTypes {
		if contains(excludeTypes, value) {
			continue
		}
		validTypes = append(validTypes, value)
	}

	attribute := map[string]resourceschema.Attribute{
		"type": resourceschema.StringAttribute{
			Optional:    true,
			Description: helper.AllowedValuesToDescription(validTypes),
			Validators: []validator.String{
				stringvalidator.OneOf(validTypes...),
				helper.RequiredIfParentConfigured(),
			},
		},
		"required": resourceschema.BoolAttribute{
			Optional: true,
			Computed: true,
			Default:  booldefault.StaticBool(false),
		},
		"max_length": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Only array elements with text type: Maximum length of this text (at least 1)",
			Validators:  []validator.Int64{int64validator.AtLeast(1)},
		},
		"min_length": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Only array elements with text type: Minimum length of this text (at least 1)",
			Validators:  []validator.Int64{int64validator.AtLeast(1)},
		},
		"pattern": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Only array elements with text type: Regex defined the allowed pattern of this text",
		},
		"max": resourceschema.Float64Attribute{
			Optional:    true,
			Description: "Only array elements with numeric type : maximum value allowed",
		},
		"precision": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Only array elements with numeric type : how many values after the comma should be accepted",
		},
		"min": resourceschema.Float64Attribute{
			Optional:    true,
			Description: "Only array elements with numeric type : minimum value allowed",
		},
		"unit_id": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Only array elements with numeric type",
			Validators:  helper.ValidUUID(),
		},
		"scale": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Only array elements with numeric type",
		},
		"options": resourceschema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "Only array elements with enum type : The allowed values for the enum in the format 1 = \"value\"",
		},
		"after": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Only array elements with time/date/timestamp type : Minimum date allowed",
			Validators:  helper.IsValidTimeDateOrTimestamp(),
		},
		"before": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Only array elements with time/date/timestamp type : Maximum date allowed",
			Validators:  helper.IsValidTimeDateOrTimestamp(),
		},
	}

	for _, field := range excludeFields {
		delete(attribute, field)
	}

	return attribute
}

func DefinitionAttributeArrayConstraintSchemaDS(excludeTypes []string, excludeFields []string) map[string]datasourceschema.Attribute {
	s := map[string]datasourceschema.Attribute{
		"type":       datasourceschema.StringAttribute{Computed: true},
		"required":   datasourceschema.BoolAttribute{Computed: true},
		"max_length": datasourceschema.Int64Attribute{Computed: true},
		"min_length": datasourceschema.Int64Attribute{Computed: true},
		"pattern":    datasourceschema.StringAttribute{Computed: true},
		"max":        datasourceschema.Float64Attribute{Computed: true},
		"precision":  datasourceschema.Int64Attribute{Computed: true},
		"min":        datasourceschema.Float64Attribute{Computed: true},
		"unit_id":    datasourceschema.StringAttribute{Computed: true},
		"scale":      datasourceschema.Int64Attribute{Computed: true},
		"options": datasourceschema.MapAttribute{
			ElementType: types.StringType,
			Computed:    true,
		},
		"after":  datasourceschema.StringAttribute{Computed: true},
		"before": datasourceschema.StringAttribute{Computed: true},
	}

	for _, field := range excludeFields {
		delete(s, field)
	}

	return s
}

var validMetadataTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM", "BINARY", "ARRAY", "TLE", "GEOPOINT", "STRUCTURE",
}

var validArraydataTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM", "BINARY",
}

func ValueAttributeSchema(excludeTypes []string) map[string]resourceschema.Attribute {
	validTypes := []string{}
	for _, value := range validMetadataTypes {
		if contains(excludeTypes, value) {
			continue
		}
		validTypes = append(validTypes, value)
	}

	return map[string]resourceschema.Attribute{
		"value": resourceschema.StringAttribute{
			Optional: true,
		},
		"type": resourceschema.StringAttribute{
			Required:    true,
			Description: helper.AllowedValuesToDescription(validTypes),
			Validators:  []validator.String{stringvalidator.OneOf(validTypes...)},
		},
		"data_type": resourceschema.StringAttribute{
			Optional:    true,
			Description: helper.AllowedValuesToDescription(validArraydataTypes),
			Validators:  []validator.String{stringvalidator.OneOf(validArraydataTypes...)},
		},
		"unit_id": resourceschema.StringAttribute{
			Optional:   true,
			Validators: helper.ValidUUID(),
		},
		"fields": resourceschema.SingleNestedAttribute{
			Optional:    true,
			Attributes:  geoPointFieldsSchema,
			Description: "Geopoint only",
		},
	}
}

func ValueAttributeSchemaDS(excludeTypes []string) map[string]datasourceschema.Attribute {
	return map[string]datasourceschema.Attribute{
		"value":     datasourceschema.StringAttribute{Computed: true},
		"type":      datasourceschema.StringAttribute{Computed: true},
		"data_type": datasourceschema.StringAttribute{Computed: true},
		"unit_id":   datasourceschema.StringAttribute{Computed: true},
		"fields": datasourceschema.SingleNestedAttribute{
			Computed:   true,
			Attributes: geoPointFieldsSchemaDS,
		},
	}
}

func ResourceSchemaWith(fields map[string]resourceschema.Attribute) map[string]resourceschema.Attribute {
	result := make(map[string]resourceschema.Attribute, len(fields)+5)
	result["id"] = resourceschema.StringAttribute{Computed: true}
	result["created_at"] = resourceschema.StringAttribute{Computed: true, Description: "When it was created"}
	result["created_by"] = resourceschema.StringAttribute{Computed: true, Description: "Who created it"}
	result["last_modified_at"] = resourceschema.StringAttribute{Computed: true, Description: "When it was last modified"}
	result["last_modified_by"] = resourceschema.StringAttribute{Computed: true, Description: "Who modified it the last"}
	for k, v := range fields {
		result[k] = v
	}
	return result
}
