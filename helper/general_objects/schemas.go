package general_objects

import (
	"context"

	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
)

// serverManagedTimestampModifier marks a Computed-only timestamp field as unknown
// whenever the resource is being updated (i.e. any other attribute has changed).
// This prevents the "inconsistent result after apply" error caused by the server
// updating the timestamp on every write while the plan kept the old known value.
type serverManagedTimestampModifier struct{}

func (m serverManagedTimestampModifier) Description(_ context.Context) string {
	return "Marks the field as (known after apply) when the resource is being updated."
}
func (m serverManagedTimestampModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}
func (m serverManagedTimestampModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// During create the value is already unknown — nothing to do.
	if req.StateValue.IsNull() {
		return
	}
	// If the entire plan equals the current state, this is a no-op plan.
	// Keep the known value so no spurious "(known after apply)" diff is shown.
	if req.Plan.Raw.Equal(req.State.Raw) {
		return
	}
	// Something is changing — the server will update the timestamp, so mark it
	// as unknown to accept whatever value comes back after apply.
	resp.PlanValue = types.StringUnknown()
}

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
			Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
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

func AuditFilterFieldsWithoutTags(filters map[string]datasourceschema.Attribute) map[string]datasourceschema.Attribute {
	return AuditFilterFields(filters, false, false)
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

var geoPointFieldsDefSchema = CreateGeoPointFieldsSchema(false)
var geoPointFieldsSchema = CreateGeoPointFieldsSchema(true)

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

// filterDefinitionTypes returns the slice of valid attribute types with the given types removed.
func filterDefinitionTypes(excludeTypes []string) []string {
	validTypes := make([]string, 0, len(ValidAttributeSchemaTypes))
	for _, v := range ValidAttributeSchemaTypes {
		if !contains(excludeTypes, v) {
			validTypes = append(validTypes, v)
		}
	}
	return validTypes
}

// sharedDefinitionConstraintFields returns the attribute constraint fields common to both
// DefinitionAttributeSchema and DefinitionAttributeArrayConstraintSchema: the required bool
// plus all per-type constraint fields (text, numeric, time/date, enum).
func sharedDefinitionConstraintFields() map[string]resourceschema.Attribute {
	return map[string]resourceschema.Attribute{
		"required": resourceschema.BoolAttribute{
			Optional: true,
			Computed: true,
			Default:  booldefault.StaticBool(false),
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
	}
}

func DefinitionAttributeSchema(excludeTypes []string, excludeFields []string, forceNew bool) map[string]resourceschema.Attribute {
	validTypes := filterDefinitionTypes(excludeTypes)

	var typePlanModifiers []planmodifier.String
	if forceNew {
		typePlanModifiers = []planmodifier.String{stringplanmodifier.RequiresReplace()}
	}

	attribute := sharedDefinitionConstraintFields()
	attribute["type"] = resourceschema.StringAttribute{
		Required:      true,
		Description:   helper.AllowedValuesToDescription(validTypes),
		Validators:    []validator.String{stringvalidator.OneOf(validTypes...)},
		PlanModifiers: typePlanModifiers,
	}
	attribute["default_value"] = resourceschema.StringAttribute{
		Optional:    true,
		Description: "The default value can be of any type. In case of an array type, please surround the list values with double quotes and use the comma separator.",
	}
	// Geopoint only
	attribute["fields"] = resourceschema.SingleNestedAttribute{
		Optional:    true,
		Attributes:  geoPointFieldsDefSchema,
		Description: "Geopoint only",
	}
	// Array
	attribute["min_size"] = resourceschema.Int64Attribute{
		Optional:    true,
		Description: "Array only: The minimum number of elements allowed",
	}
	attribute["max_size"] = resourceschema.Int64Attribute{
		Optional:    true,
		Description: "Array only: The maximum number of elements allowed",
	}
	attribute["unique"] = resourceschema.BoolAttribute{
		Optional:      true,
		Computed:      true,
		Description:   "Array only: No duplicated elements are allowed",
		PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
	}
	attribute["constraint"] = resourceschema.SingleNestedAttribute{
		Optional:    true,
		Description: "Array only: Constraint applied to all elements in the array",
		Attributes: DefinitionAttributeArrayConstraintSchema(
			[]string{"ARRAY", "STRUCTURE", "GEOPOINT", "TLE"},
			nil,
		),
	}

	for _, field := range excludeFields {
		delete(attribute, field)
	}

	return attribute
}

func DefinitionAttributeArrayConstraintSchema(excludeTypes []string, excludeFields []string) map[string]resourceschema.Attribute {
	validTypes := filterDefinitionTypes(excludeTypes)

	attribute := sharedDefinitionConstraintFields()
	attribute["type"] = resourceschema.StringAttribute{
		Optional:    true,
		Description: helper.AllowedValuesToDescription(validTypes),
		Validators: []validator.String{
			stringvalidator.OneOf(validTypes...),
			helper.RequiredIfParentConfigured(),
		},
	}

	for _, field := range excludeFields {
		delete(attribute, field)
	}

	return attribute
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

func ResourceSchemaWith(fields map[string]resourceschema.Attribute) map[string]resourceschema.Attribute {
	result := make(map[string]resourceschema.Attribute, len(fields)+5)
	result["id"] = resourceschema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}}
	result["created_at"] = resourceschema.StringAttribute{Computed: true, Description: "When it was created", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}}
	result["created_by"] = resourceschema.StringAttribute{Computed: true, Description: "Who created it", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}}
	result["last_modified_at"] = resourceschema.StringAttribute{Computed: true, Description: "When it was last modified", PlanModifiers: []planmodifier.String{serverManagedTimestampModifier{}}}
	result["last_modified_by"] = resourceschema.StringAttribute{Computed: true, Description: "Who modified it the last", PlanModifiers: []planmodifier.String{serverManagedTimestampModifier{}}}
	for k, v := range fields {
		result[k] = v
	}
	return result
}
