package general_objects

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

// serverManagedTimestampModifier marks a Computed-only timestamp field as unknown
// whenever the resource is being updated (i.e. any user-configurable attribute has changed).
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
	// Check for a no-op plan using a comparison that treats unknown plan values as equal
	// to their state counterparts. Unknown values appear for Optional+Computed attributes
	// whose UseStateForUnknown modifier hasn't run yet — they are not user-initiated changes.
	isNoOp := noOpPlan(req.Plan.Raw, req.State.Raw)
	if isNoOp {
		// Restore the known state value even if something earlier already set the plan to
		// unknown (e.g. another plan modifier upstream). On a no-op plan the server will not
		// touch this field, so we can safely keep the prior known value.
		resp.PlanValue = req.StateValue
		return
	}
	// Something is changing — the server will update the timestamp, so mark it
	// as unknown to accept whatever value comes back after apply.
	resp.PlanValue = types.StringUnknown()
}

// noOpPlan returns true if plan and state are effectively equal, treating unknown
// plan values as equal to whatever the state has. Unknown values in the plan arise
// for Optional+Computed attributes whose UseStateForUnknown modifier hasn't run yet;
// they do not represent a user-driven change to the resource.
// NOTE: do NOT use serverManagedTimestampModifier on resources whose Computed-only
// attributes depend on recreated external resources — use alwaysUnknownAfterApplyModifier
// instead to keep phase-1 and phase-2 planning consistent.
func noOpPlan(plan, state tftypes.Value) bool {
	if plan.Equal(state) {
		return true
	}
	// Unknown in plan → will be preserved from state by UseStateForUnknown → treat as equal.
	if !plan.IsKnown() {
		return true
	}
	if plan.IsNull() != state.IsNull() {
		return false
	}
	if plan.IsNull() {
		return true // both null
	}
	// Both known and non-null but not bitwise equal: recurse into compound types.
	typ := plan.Type()
	switch {
	case typ.Is(tftypes.Object{}):
		planAttrs := map[string]tftypes.Value{}
		stateAttrs := map[string]tftypes.Value{}
		_ = plan.As(&planAttrs)
		_ = state.As(&stateAttrs)
		for k, pv := range planAttrs {
			sv, ok := stateAttrs[k]
			if !ok {
				return false
			}
			if !noOpPlan(pv, sv) {
				return false
			}
		}
		return true
	case typ.Is(tftypes.List{}) || typ.Is(tftypes.Set{}) || typ.Is(tftypes.Tuple{}):
		planElems := []tftypes.Value{}
		stateElems := []tftypes.Value{}
		_ = plan.As(&planElems)
		_ = state.As(&stateElems)
		if len(planElems) != len(stateElems) {
			return false
		}
		for i := range planElems {
			if !noOpPlan(planElems[i], stateElems[i]) {
				return false
			}
		}
		return true
	case typ.Is(tftypes.Map{}):
		planMap := map[string]tftypes.Value{}
		stateMap := map[string]tftypes.Value{}
		_ = plan.As(&planMap)
		_ = state.As(&stateMap)
		if len(planMap) != len(stateMap) {
			return false
		}
		for k, pv := range planMap {
			sv, ok := stateMap[k]
			if !ok {
				return false
			}
			if !noOpPlan(pv, sv) {
				return false
			}
		}
		return true
	default:
		// Primitive types: plan is known and non-null but not equal to state.
		return false
	}
}

// isComputedOnly returns true when the attribute is Computed but neither Optional nor Required.
// Such attributes are server-managed; unknown plan values for them arise from UseStateForUnknown
// running lazily and do NOT indicate a user-driven change.
func isComputedOnly(attr resourceschema.Attribute) bool {
	switch a := attr.(type) {
	case resourceschema.StringAttribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.Int64Attribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.Float64Attribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.BoolAttribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.NumberAttribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.ListAttribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.SetAttribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.ListNestedAttribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.SetNestedAttribute:
		return a.Computed && !a.Optional && !a.Required
	case resourceschema.SingleNestedAttribute:
		return a.Computed && !a.Optional && !a.Required
	default:
		return false
	}
}

// isOptionalComputed returns true when the attribute is both Optional and Computed.
// Such attributes may be server-managed when the user omits them.
func isOptionalComputed(attr resourceschema.Attribute) bool {
	switch a := attr.(type) {
	case resourceschema.StringAttribute:
		return a.Computed && a.Optional
	case resourceschema.Int64Attribute:
		return a.Computed && a.Optional
	case resourceschema.Float64Attribute:
		return a.Computed && a.Optional
	case resourceschema.BoolAttribute:
		return a.Computed && a.Optional
	case resourceschema.NumberAttribute:
		return a.Computed && a.Optional
	case resourceschema.ListAttribute:
		return a.Computed && a.Optional
	case resourceschema.SetAttribute:
		return a.Computed && a.Optional
	case resourceschema.ListNestedAttribute:
		return a.Computed && a.Optional
	case resourceschema.SetNestedAttribute:
		return a.Computed && a.Optional
	case resourceschema.SingleNestedAttribute:
		return a.Computed && a.Optional
	default:
		return false
	}
}

// getNestedAttrSchema returns the child attribute schema for a nested attribute, or nil for scalars.
func getNestedAttrSchema(schema map[string]resourceschema.Attribute, key string) map[string]resourceschema.Attribute {
	if schema == nil {
		return nil
	}
	attr, ok := schema[key]
	if !ok {
		return nil
	}
	switch a := attr.(type) {
	case resourceschema.SingleNestedAttribute:
		return a.Attributes
	case resourceschema.ListNestedAttribute:
		return a.NestedObject.Attributes
	case resourceschema.SetNestedAttribute:
		return a.NestedObject.Attributes
	default:
		return nil
	}
}

// noOpPlanSchema is like noOpPlan but schema-aware: when a plan value is unknown,
// it checks whether the corresponding schema attribute is Computed-only.
//   - Computed-only unknown → UseStateForUnknown will resolve it → treat as no-op.
//   - Required/Optional unknown → dependency being recreated → treat as potential change.
//
// When schema is nil the function falls back to treating all unknowns as no-ops (safe default).
func noOpPlanSchema(plan, state tftypes.Value, schema map[string]resourceschema.Attribute) bool {
	if plan.Equal(state) {
		return true
	}
	if !plan.IsKnown() {
		return true // whole-value unknown: very unlikely at top level, be conservative
	}
	if plan.IsNull() != state.IsNull() {
		return false
	}
	if plan.IsNull() {
		return true // both null
	}
	typ := plan.Type()
	switch {
	case typ.Is(tftypes.Object{}):
		planAttrs := map[string]tftypes.Value{}
		stateAttrs := map[string]tftypes.Value{}
		_ = plan.As(&planAttrs)
		_ = state.As(&stateAttrs)
		for k, pv := range planAttrs {
			sv, ok := stateAttrs[k]
			if !ok {
				return false
			}
			if !pv.IsKnown() {
				// Unknown plan value: check schema to decide if it's harmless.
				if schema != nil {
					if attr, exists := schema[k]; exists && isComputedOnly(attr) {
						continue // Computed-only — UseStateForUnknown will handle it
					}
					// Optional+Computed with null state: user omitted it and the server determines
					// the value. UseStateForUnknown won't fire (state is null), so this unknown
					// is not a user-driven change.
					if attr, exists := schema[k]; exists && isOptionalComputed(attr) && (!sv.IsKnown() || sv.IsNull()) {
						continue
					}
				}
				return false // Required/Optional unknown → dependency recreation → potential change
			}
			nestedSchema := getNestedAttrSchema(schema, k)
			if !noOpPlanSchema(pv, sv, nestedSchema) {
				return false
			}
		}
		return true
	case typ.Is(tftypes.List{}) || typ.Is(tftypes.Set{}) || typ.Is(tftypes.Tuple{}):
		planElems := []tftypes.Value{}
		stateElems := []tftypes.Value{}
		_ = plan.As(&planElems)
		_ = state.As(&stateElems)
		if len(planElems) != len(stateElems) {
			return false
		}
		// schema here is already the element-level schema (passed from the parent Object case)
		for i := range planElems {
			if !noOpPlanSchema(planElems[i], stateElems[i], schema) {
				return false
			}
		}
		return true
	case typ.Is(tftypes.Map{}):
		planMap := map[string]tftypes.Value{}
		stateMap := map[string]tftypes.Value{}
		_ = plan.As(&planMap)
		_ = state.As(&stateMap)
		if len(planMap) != len(stateMap) {
			return false
		}
		for k, pv := range planMap {
			sv, ok := stateMap[k]
			if !ok {
				return false
			}
			if !noOpPlanSchema(pv, sv, schema) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// schemaAwareTimestampModifier behaves like serverManagedTimestampModifier but uses
// noOpPlanSchema to distinguish Computed-only unknowns (harmless UseStateForUnknown)
// from Required/Optional unknowns (dependency being recreated → potential change).
// This makes it safe to use on resources like plan_templates where Computed-only
// sub-attributes reference external resources that may be recreated in the same apply.
type schemaAwareTimestampModifier struct {
	schema map[string]resourceschema.Attribute
}

func (m schemaAwareTimestampModifier) Description(_ context.Context) string {
	return "Marks the field as (known after apply) when the resource is being updated."
}
func (m schemaAwareTimestampModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}
func (m schemaAwareTimestampModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return // Creating — already unknown
	}
	if noOpPlanSchema(req.Plan.Raw, req.State.Raw, m.schema) {
		resp.PlanValue = req.StateValue
		return
	}
	resp.PlanValue = types.StringUnknown()
}

func PaginatedListSchemaDS(content, filters map[string]datasourceschema.Attribute) map[string]datasourceschema.Attribute {
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

func baseAttributeFieldSchema(isValueField, isGeoPoint bool) map[string]resourceschema.Attribute {
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
		Optional:    true,
		Computed:    true,
		Description: "Array only: No duplicated elements are allowed",
		Default:     booldefault.StaticBool(false),
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

func DefinitionAttributeArrayConstraintSchema(excludeTypes, excludeFields []string) map[string]resourceschema.Attribute {
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

// ServerManagedTimestampAttribute returns a Computed string attribute whose plan
// uses serverManagedTimestampModifier. Use this for most resources.
// For resources whose Computed-only sub-attributes depend on recreated external
// resources, use AlwaysUnknownAfterApplyAttribute instead.
func ServerManagedTimestampAttribute(description string) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		Computed:      true,
		Description:   description,
		PlanModifiers: []planmodifier.String{serverManagedTimestampModifier{}},
	}
}

// AlwaysUnknownAfterApplyAttribute returns a Computed string attribute that is
// always shown as (known after apply) in plans after creation. This is the safe
// choice when serverManagedTimestampModifier would produce inconsistent results
// between planning phase 1 and phase 2 (e.g. when Computed-only sub-attributes
// depend on external resources being recreated in the same apply).
func AlwaysUnknownAfterApplyAttribute(description string) resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		Computed:      true,
		Description:   description,
		PlanModifiers: []planmodifier.String{alwaysUnknownAfterApplyModifier{}},
	}
}

// alwaysUnknownAfterApplyModifier always marks the field as (known after apply)
// once the resource exists. Unlike serverManagedTimestampModifier it does not
// attempt to keep the old state value on no-op plans, which makes it safe to use
// even when the surrounding plan contains unknowns from recreated dependencies.
type alwaysUnknownAfterApplyModifier struct{}

func (m alwaysUnknownAfterApplyModifier) Description(_ context.Context) string {
	return "Always marks the field as (known after apply) after the resource is created."
}
func (m alwaysUnknownAfterApplyModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}
func (m alwaysUnknownAfterApplyModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return // Creating — already unknown
	}
	resp.PlanValue = types.StringUnknown()
}

func ResourceSchemaWith(fields map[string]resourceschema.Attribute) map[string]resourceschema.Attribute {
	result := make(map[string]resourceschema.Attribute, len(fields)+5)
	result["id"] = resourceschema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}}
	result["created_at"] = resourceschema.StringAttribute{Computed: true, Description: "When it was created", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}}
	result["created_by"] = resourceschema.StringAttribute{Computed: true, Description: "Who created it", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}}
	for k, v := range fields {
		result[k] = v
	}
	// Use a schema-aware modifier that holds a reference to the (now-fully-populated) result map.
	// This lets it distinguish Computed-only unknowns (UseStateForUnknown, harmless) from
	// Required/Optional unknowns (dependency recreation, potential change).
	mod := schemaAwareTimestampModifier{schema: result}
	result["last_modified_at"] = resourceschema.StringAttribute{Computed: true, Description: "When it was last modified", PlanModifiers: []planmodifier.String{mod}}
	result["last_modified_by"] = resourceschema.StringAttribute{Computed: true, Description: "Who modified it the last", PlanModifiers: []planmodifier.String{mod}}
	return result
}
