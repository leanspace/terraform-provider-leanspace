package general_objects

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

// AuditModelTF is the Terraform model for AuditModel.
// Embed this (without a tfsdk tag) in service TF models that use AuditModel.
type AuditModelTF struct {
	ID             types.String `tfsdk:"id"`
	CreatedAt      types.String `tfsdk:"created_at"`
	CreatedBy      types.String `tfsdk:"created_by"`
	LastModifiedAt types.String `tfsdk:"last_modified_at"`
	LastModifiedBy types.String `tfsdk:"last_modified_by"`
}

func AuditModelToTF(a *AuditModel) AuditModelTF {
	return AuditModelTF{
		ID:             helper.TFStringValue(a.ID),
		CreatedAt:      helper.TFStringValue(a.CreatedAt),
		CreatedBy:      helper.TFStringValue(a.CreatedBy),
		LastModifiedAt: helper.TFStringValue(a.LastModifiedAt),
		LastModifiedBy: helper.TFStringValue(a.LastModifiedBy),
	}
}

func AuditModelFromTF(tf AuditModelTF) AuditModel {
	return AuditModel{
		ID:             helper.FromTFString(tf.ID),
		CreatedAt:      helper.FromTFString(tf.CreatedAt),
		CreatedBy:      helper.FromTFString(tf.CreatedBy),
		LastModifiedAt: helper.FromTFString(tf.LastModifiedAt),
		LastModifiedBy: helper.FromTFString(tf.LastModifiedBy),
	}
}

// KeyValueTF is the Terraform model for KeyValue.
type KeyValueTF struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func KeyValuesToTF(kvs []KeyValue) []KeyValueTF {
	if kvs == nil {
		return nil
	}
	result := make([]KeyValueTF, len(kvs))
	for i, kv := range kvs {
		result[i] = KeyValueTF{
			Key:   helper.TFStringValue(kv.Key),
			Value: helper.TFStringValue(kv.Value),
		}
	}
	return result
}

func KeyValuesFromTF(tfs []KeyValueTF) []KeyValue {
	if tfs == nil {
		return nil
	}
	result := make([]KeyValue, len(tfs))
	for i, tf := range tfs {
		result[i] = KeyValue{
			Key:   helper.FromTFString(tf.Key),
			Value: helper.FromTFString(tf.Value),
		}
	}
	return result
}

// --- Geopoint Field TF models ---

// FieldDefTF is the TF model for FieldDef (geopoint definition fields).
type FieldDefTF struct {
	DefaultValue types.String  `tfsdk:"default_value"`
	Min          types.Float64 `tfsdk:"min"`
	Max          types.Float64 `tfsdk:"max"`
	Scale        types.Int64   `tfsdk:"scale"`
	Precision    types.Int64   `tfsdk:"precision"`
	UnitId       types.String  `tfsdk:"unit_id"`
}

// FieldTF is the TF model for Field (geopoint value fields).
type FieldTF struct {
	Value     types.String  `tfsdk:"value"`
	Min       types.Float64 `tfsdk:"min"`
	Max       types.Float64 `tfsdk:"max"`
	Scale     types.Int64   `tfsdk:"scale"`
	Precision types.Int64   `tfsdk:"precision"`
	UnitId    types.String  `tfsdk:"unit_id"`
}

type FieldsDefTF struct {
	Elevation *FieldDefTF `tfsdk:"elevation"`
	Latitude  *FieldDefTF `tfsdk:"latitude"`
	Longitude *FieldDefTF `tfsdk:"longitude"`
}

type FieldsTF struct {
	Elevation *FieldTF `tfsdk:"elevation"`
	Latitude  *FieldTF `tfsdk:"latitude"`
	Longitude *FieldTF `tfsdk:"longitude"`
}

func FieldDefToTF(f *FieldDef[any]) FieldDefTF {
	var dv types.String
	if any(f.DefaultValue) != nil {
		dv = helper.TFStringValue(fmt.Sprint(f.DefaultValue))
	} else {
		dv = types.StringNull()
	}
	return FieldDefTF{
		DefaultValue: dv,
		Min:          helper.TFFloat64PtrValue(f.Min),
		Max:          helper.TFFloat64PtrValue(f.Max),
		Scale:        helper.TFIntPtrValue(f.Scale),
		Precision:    helper.TFIntPtrValue(f.Precision),
		UnitId:       helper.TFStringValue(f.UnitId),
	}
}

func FieldDefFromTF(tf FieldDefTF) FieldDef[any] {
	var dv any
	if !tf.DefaultValue.IsNull() && !tf.DefaultValue.IsUnknown() {
		dv = tf.DefaultValue.ValueString()
	}
	return FieldDef[any]{
		DefaultValue: dv,
		Min:          helper.FromTFFloat64Ptr(tf.Min),
		Max:          helper.FromTFFloat64Ptr(tf.Max),
		Scale:        helper.FromTFIntPtr(tf.Scale),
		Precision:    helper.FromTFIntPtr(tf.Precision),
		UnitId:       helper.FromTFString(tf.UnitId),
	}
}

func FieldToTF(f *Field[any]) FieldTF {
	var v types.String
	if any(f.Value) != nil {
		v = helper.TFStringValue(fmt.Sprint(f.Value))
	} else {
		v = types.StringNull()
	}
	return FieldTF{
		Value:     v,
		Min:       helper.TFFloat64PtrValue(f.Min),
		Max:       helper.TFFloat64PtrValue(f.Max),
		Scale:     helper.TFIntPtrValue(f.Scale),
		Precision: helper.TFIntPtrValue(f.Precision),
		UnitId:    helper.TFStringValue(f.UnitId),
	}
}

func FieldFromTF(tf FieldTF) Field[any] {
	var v any
	if !tf.Value.IsNull() && !tf.Value.IsUnknown() {
		v = tf.Value.ValueString()
	}
	return Field[any]{
		Value:     v,
		Min:       helper.FromTFFloat64Ptr(tf.Min),
		Max:       helper.FromTFFloat64Ptr(tf.Max),
		Scale:     helper.FromTFIntPtr(tf.Scale),
		Precision: helper.FromTFIntPtr(tf.Precision),
		UnitId:    helper.FromTFString(tf.UnitId),
	}
}

func FieldsDefToTF(f *FieldsDef) *FieldsDefTF {
	if f == nil {
		return nil
	}
	elev := FieldDefToTF(&f.Elevation)
	lat := FieldDefToTF(&f.Latitude)
	lon := FieldDefToTF(&f.Longitude)
	return &FieldsDefTF{Elevation: &elev, Latitude: &lat, Longitude: &lon}
}

func FieldsDefFromTF(tf *FieldsDefTF) *FieldsDef {
	if tf == nil {
		return nil
	}
	f := &FieldsDef{}
	if tf.Elevation != nil {
		f.Elevation = FieldDefFromTF(*tf.Elevation)
	}
	if tf.Latitude != nil {
		f.Latitude = FieldDefFromTF(*tf.Latitude)
	}
	if tf.Longitude != nil {
		f.Longitude = FieldDefFromTF(*tf.Longitude)
	}
	return f
}

func FieldsToTF(f *Fields) *FieldsTF {
	if f == nil {
		return nil
	}
	elev := FieldToTF(&f.Elevation)
	lat := FieldToTF(&f.Latitude)
	lon := FieldToTF(&f.Longitude)
	return &FieldsTF{Elevation: &elev, Latitude: &lat, Longitude: &lon}
}

func FieldsFromTF(tf *FieldsTF) *Fields {
	if tf == nil {
		return nil
	}
	f := &Fields{}
	if tf.Elevation != nil {
		f.Elevation = FieldFromTF(*tf.Elevation)
	}
	if tf.Latitude != nil {
		f.Latitude = FieldFromTF(*tf.Latitude)
	}
	if tf.Longitude != nil {
		f.Longitude = FieldFromTF(*tf.Longitude)
	}
	return f
}

// --- ArrayConstraint TF model ---

type ArrayConstraintTF struct {
	Type      types.String            `tfsdk:"type"`
	Required  types.Bool              `tfsdk:"required"`
	MinLength types.Int64             `tfsdk:"min_length"`
	MaxLength types.Int64             `tfsdk:"max_length"`
	Pattern   types.String            `tfsdk:"pattern"`
	Min       types.Float64           `tfsdk:"min"`
	Max       types.Float64           `tfsdk:"max"`
	Scale     types.Int64             `tfsdk:"scale"`
	Precision types.Int64             `tfsdk:"precision"`
	UnitId    types.String            `tfsdk:"unit_id"`
	Before    types.String            `tfsdk:"before"`
	After     types.String            `tfsdk:"after"`
	Options   map[string]types.String `tfsdk:"options"`
}

func ArrayConstraintToTF(c *ArrayConstraint[any]) *ArrayConstraintTF {
	if c == nil || c.Type == "" {
		return nil
	}
	tf := &ArrayConstraintTF{
		Type:      helper.TFStringValue(c.Type),
		Required:  helper.TFBoolPtrValue(c.Required),
		MinLength: helper.TFIntPtrValue(c.MinLength),
		MaxLength: helper.TFIntPtrValue(c.MaxLength),
		Pattern:   helper.TFStringValue(c.Pattern),
		Min:       helper.TFFloat64PtrValue(c.Min),
		Max:       helper.TFFloat64PtrValue(c.Max),
		Scale:     helper.TFIntPtrValue(c.Scale),
		Precision: helper.TFIntPtrValue(c.Precision),
		UnitId:    helper.TFStringValue(c.UnitId),
		Before:    helper.TFStringValue(c.Before),
		After:     helper.TFStringValue(c.After),
	}
	if c.Options != nil {
		tf.Options = make(map[string]types.String, len(*c.Options))
		for k, v := range *c.Options {
			tf.Options[k] = helper.TFStringValue(fmt.Sprint(v))
		}
	}
	return tf
}

func ArrayConstraintFromTF(tf *ArrayConstraintTF) ArrayConstraint[any] {
	if tf == nil {
		return ArrayConstraint[any]{}
	}
	c := ArrayConstraint[any]{
		Type:      helper.FromTFString(tf.Type),
		Required:  helper.FromTFBoolPtr(tf.Required),
		MinLength: helper.FromTFIntPtr(tf.MinLength),
		MaxLength: helper.FromTFIntPtr(tf.MaxLength),
		Pattern:   helper.FromTFString(tf.Pattern),
		Min:       helper.FromTFFloat64Ptr(tf.Min),
		Max:       helper.FromTFFloat64Ptr(tf.Max),
		Scale:     helper.FromTFIntPtr(tf.Scale),
		Precision: helper.FromTFIntPtr(tf.Precision),
		UnitId:    helper.FromTFString(tf.UnitId),
		Before:    helper.FromTFString(tf.Before),
		After:     helper.FromTFString(tf.After),
	}
	if tf.Options != nil {
		opts := make(map[string]any, len(tf.Options))
		for k, v := range tf.Options {
			opts[k] = helper.FromTFString(v)
		}
		c.Options = &opts
	}
	return c
}

// --- DefinitionAttribute TF model (full version with all fields) ---

type DefinitionAttributeTF struct {
	Type         types.String            `tfsdk:"type"`
	Required     types.Bool              `tfsdk:"required"`
	DefaultValue types.String            `tfsdk:"default_value"`
	MinLength    types.Int64             `tfsdk:"min_length"`
	MaxLength    types.Int64             `tfsdk:"max_length"`
	Pattern      types.String            `tfsdk:"pattern"`
	Min          types.Float64           `tfsdk:"min"`
	Max          types.Float64           `tfsdk:"max"`
	Scale        types.Int64             `tfsdk:"scale"`
	Precision    types.Int64             `tfsdk:"precision"`
	UnitId       types.String            `tfsdk:"unit_id"`
	Before       types.String            `tfsdk:"before"`
	After        types.String            `tfsdk:"after"`
	Options      map[string]types.String `tfsdk:"options"`
	Fields       *FieldsDefTF            `tfsdk:"fields"`
	MinSize      types.Int64             `tfsdk:"min_size"`
	MaxSize      types.Int64             `tfsdk:"max_size"`
	Unique       types.Bool              `tfsdk:"unique"`
	Constraint   *ArrayConstraintTF      `tfsdk:"constraint"`
}

func DefinitionAttributeToTF(a *DefinitionAttribute[any]) DefinitionAttributeTF {
	tf := DefinitionAttributeTF{
		Type:      helper.TFStringValue(a.Type),
		Required:  helper.TFBoolPtrValue(a.Required),
		MinLength: helper.TFIntPtrValue(a.MinLength),
		MaxLength: helper.TFIntPtrValue(a.MaxLength),
		Pattern:   helper.TFStringValue(a.Pattern),
		Min:       helper.TFFloat64PtrValue(a.Min),
		Max:       helper.TFFloat64PtrValue(a.Max),
		Scale:     helper.TFIntPtrValue(a.Scale),
		Precision: helper.TFIntPtrValue(a.Precision),
		UnitId:    helper.TFStringValue(a.UnitId),
		Before:    helper.TFStringValue(a.Before),
		After:     helper.TFStringValue(a.After),
		Fields:    FieldsDefToTF(a.Fields),
		MinSize:   helper.TFIntPtrValue(a.MinSize),
		MaxSize:   helper.TFIntPtrValue(a.MaxSize),
		Unique:    helper.TFBoolValue(a.Unique),
	}
	if any(a.DefaultValue) != nil {
		switch v := a.DefaultValue.(type) {
		case []interface{}:
			parts := make([]string, len(v))
			for i, elem := range v {
				if f, ok := elem.(float64); ok && f == float64(int64(f)) {
					parts[i] = strconv.FormatInt(int64(f), 10)
				} else {
					parts[i] = fmt.Sprint(elem)
				}
			}
			tf.DefaultValue = helper.TFStringValue(strings.Join(parts, ","))
		default:
			tf.DefaultValue = helper.TFStringValue(fmt.Sprint(a.DefaultValue))
		}
	} else {
		tf.DefaultValue = types.StringNull()
	}
	if a.Options != nil {
		tf.Options = make(map[string]types.String, len(*a.Options))
		for k, v := range *a.Options {
			tf.Options[k] = helper.TFStringValue(fmt.Sprint(v))
		}
	}
	tf.Constraint = ArrayConstraintToTF(&a.Constraint)
	return tf
}

func DefinitionAttributeFromTF(tf DefinitionAttributeTF) DefinitionAttribute[any] {
	a := DefinitionAttribute[any]{
		Type:      helper.FromTFString(tf.Type),
		Required:  helper.FromTFBoolPtr(tf.Required),
		MinLength: helper.FromTFIntPtr(tf.MinLength),
		MaxLength: helper.FromTFIntPtr(tf.MaxLength),
		Pattern:   helper.FromTFString(tf.Pattern),
		Min:       helper.FromTFFloat64Ptr(tf.Min),
		Max:       helper.FromTFFloat64Ptr(tf.Max),
		Scale:     helper.FromTFIntPtr(tf.Scale),
		Precision: helper.FromTFIntPtr(tf.Precision),
		UnitId:    helper.FromTFString(tf.UnitId),
		Before:    helper.FromTFString(tf.Before),
		After:     helper.FromTFString(tf.After),
		Fields:    FieldsDefFromTF(tf.Fields),
		MinSize:   helper.FromTFIntPtr(tf.MinSize),
		MaxSize:   helper.FromTFIntPtr(tf.MaxSize),
		Unique:    helper.FromTFBool(tf.Unique),
	}
	if !tf.DefaultValue.IsNull() && !tf.DefaultValue.IsUnknown() {
		rawDefault := tf.DefaultValue.ValueString()
		if helper.FromTFString(tf.Type) == "ARRAY" && rawDefault != "" {
			parts := strings.Split(rawDefault, ",")
			arr := make([]any, len(parts))
			constraintType := ""
			if tf.Constraint != nil {
				constraintType = helper.FromTFString(tf.Constraint.Type)
			}
			for i, p := range parts {
				p = strings.TrimSpace(p)
				switch constraintType {
				case "NUMERIC":
					if n, err := strconv.ParseFloat(p, 64); err == nil {
						arr[i] = n
					} else {
						arr[i] = p
					}
				case "BOOLEAN":
					if b, err := strconv.ParseBool(p); err == nil {
						arr[i] = b
					} else {
						arr[i] = p
					}
				default:
					arr[i] = p
				}
			}
			a.DefaultValue = arr
		} else {
			a.DefaultValue = rawDefault
		}
	}
	if tf.Options != nil {
		opts := make(map[string]any, len(tf.Options))
		for k, v := range tf.Options {
			opts[k] = helper.FromTFString(v)
		}
		a.Options = &opts
	}
	a.Constraint = ArrayConstraintFromTF(tf.Constraint)
	return a
}

// --- ValueAttribute TF model ---

type ValueAttributeTF struct {
	Value    types.String `tfsdk:"value"`
	Type     types.String `tfsdk:"type"`
	DataType types.String `tfsdk:"data_type"`
	UnitId   types.String `tfsdk:"unit_id"`
	Fields   *FieldsTF    `tfsdk:"fields"`
}

func ValueAttributeToTF(a *ValueAttribute[any]) ValueAttributeTF {
	tf := ValueAttributeTF{
		Type:     helper.TFStringValue(a.Type),
		DataType: helper.TFStringValue(a.DataType),
		UnitId:   helper.TFStringValue(a.UnitId),
		Fields:   FieldsToTF(a.Fields),
	}
	if any(a.Value) != nil {
		switch v := a.Value.(type) {
		case []interface{}:
			parts := make([]string, len(v))
			for i, elem := range v {
				if f, ok := elem.(float64); ok && f == float64(int64(f)) {
					parts[i] = strconv.FormatInt(int64(f), 10)
				} else {
					parts[i] = fmt.Sprint(elem)
				}
			}
			tf.Value = helper.TFStringValue(strings.Join(parts, ","))
		default:
			tf.Value = helper.TFStringValue(fmt.Sprint(a.Value))
		}
	} else {
		tf.Value = types.StringNull()
	}
	return tf
}

func ValueAttributeFromTF(tf ValueAttributeTF) ValueAttribute[any] {
	a := ValueAttribute[any]{
		Type:     helper.FromTFString(tf.Type),
		DataType: helper.FromTFString(tf.DataType),
		UnitId:   helper.FromTFString(tf.UnitId),
		Fields:   FieldsFromTF(tf.Fields),
	}
	if !tf.Value.IsNull() && !tf.Value.IsUnknown() {
		rawValue := tf.Value.ValueString()
		if helper.FromTFString(tf.Type) == "ARRAY" && rawValue != "" {
			parts := strings.Split(rawValue, ",")
			arr := make([]any, len(parts))
			for i, p := range parts {
				p = strings.TrimSpace(p)
				if n, err := strconv.ParseFloat(p, 64); err == nil {
					arr[i] = n
				} else if b, err := strconv.ParseBool(p); err == nil {
					arr[i] = b
				} else {
					arr[i] = p
				}
			}
			a.Value = arr
		} else {
			a.Value = rawValue
		}
	}
	return a
}

// --- Data Source TF models ---

// SortTF is the TF model for Sort (used in paginated data source responses).
type SortTF struct {
	Direction    types.String `tfsdk:"direction"`
	Property     types.String `tfsdk:"property"`
	IgnoreCase   types.Bool   `tfsdk:"ignore_case"`
	NullHandling types.String `tfsdk:"null_handling"`
	Ascending    types.Bool   `tfsdk:"ascending"`
	Descending   types.Bool   `tfsdk:"descending"`
}

func SortToTF(s *Sort) SortTF {
	return SortTF{
		Direction:    helper.TFStringValue(s.Direction),
		Property:     helper.TFStringValue(s.Property),
		IgnoreCase:   helper.TFBoolValue(s.IgnoreCase),
		NullHandling: helper.TFStringValue(s.NullHandling),
		Ascending:    helper.TFBoolValue(s.Ascending),
		Descending:   helper.TFBoolValue(s.Descending),
	}
}

func SortsToTF(sorts []Sort) []SortTF {
	result := make([]SortTF, len(sorts))
	for i := range sorts {
		result[i] = SortToTF(&sorts[i])
	}
	return result
}

// PageableTF is the TF model for Pageable.
type PageableTF struct {
	Sort       []SortTF    `tfsdk:"sort"`
	Offset     types.Int64 `tfsdk:"offset"`
	PageNumber types.Int64 `tfsdk:"page_number"`
	PageSize   types.Int64 `tfsdk:"page_size"`
	Paged      types.Bool  `tfsdk:"paged"`
	Unpaged    types.Bool  `tfsdk:"unpaged"`
}

func PageableToTF(p *Pageable) *PageableTF {
	return &PageableTF{
		Sort:       SortsToTF(p.Sort),
		Offset:     helper.TFInt64Value(p.Offset),
		PageNumber: helper.TFInt64Value(p.PageNumber),
		PageSize:   helper.TFInt64Value(p.PageSize),
		Paged:      helper.TFBoolValue(p.Paged),
		Unpaged:    helper.TFBoolValue(p.Unpaged),
	}
}

// PaginatedDataSourceTF is the TF model for paginated data source responses.
// It is the same for ALL paginated data sources (content items are empty objects
// since DataSourceSchema is never set by any service).
type PaginatedDataSourceTF struct {
	ID               types.String `tfsdk:"id"`
	Content          types.List   `tfsdk:"content"`
	TotalElements    types.Int64  `tfsdk:"total_elements"`
	TotalPages       types.Int64  `tfsdk:"total_pages"`
	NumberOfElements types.Int64  `tfsdk:"number_of_elements"`
	Number           types.Int64  `tfsdk:"number"`
	Size             types.Int64  `tfsdk:"size"`
	Sort             []SortTF     `tfsdk:"sort"`
	First            types.Bool   `tfsdk:"first"`
	Last             types.Bool   `tfsdk:"last"`
	Empty            types.Bool   `tfsdk:"empty"`
	Pageable         *PageableTF  `tfsdk:"pageable"`
	Filters          types.Object `tfsdk:"filters"`
}

// ToDataSourceTF converts a PaginatedList API response into a PaginatedDataSourceTF
// suitable for State.Set(). Content items become empty objects (matching the nil
// DataSourceSchema convention). Filters are passed through from the config.
func (pl *PaginatedList[T, PT]) ToDataSourceTF(filters types.Object) PaginatedDataSourceTF {
	// Build content as a list of empty objects
	emptyObjType := types.ObjectType{AttrTypes: map[string]attr.Type{}}
	contentElems := make([]attr.Value, len(pl.Content))
	for i := range contentElems {
		contentElems[i], _ = types.ObjectValue(map[string]attr.Type{}, map[string]attr.Value{})
	}
	contentList, _ := types.ListValue(emptyObjType, contentElems)

	return PaginatedDataSourceTF{
		ID:               types.StringValue(strconv.FormatInt(time.Now().Unix(), 10)),
		Content:          contentList,
		TotalElements:    types.Int64Value(int64(pl.TotalElements)),
		TotalPages:       types.Int64Value(int64(pl.TotalPages)),
		NumberOfElements: types.Int64Value(int64(pl.NumberOfElements)),
		Number:           types.Int64Value(int64(pl.Number)),
		Size:             types.Int64Value(int64(pl.Size)),
		Sort:             SortsToTF(pl.Sort),
		First:            types.BoolValue(pl.First),
		Last:             types.BoolValue(pl.Last),
		Empty:            types.BoolValue(pl.Empty),
		Pageable:         PageableToTF(&pl.Pageable),
		Filters:          filters,
	}
}

// FilterObjectToMap converts a types.Object (from data source config filters)
// into a map[string]any suitable for building API query parameters.
// Filter values are always simple types: strings, ints, bools, or lists of strings.
func FilterObjectToMap(filtersObj types.Object) map[string]any {
	if filtersObj.IsNull() || filtersObj.IsUnknown() {
		return nil
	}
	result := make(map[string]any)
	for key, val := range filtersObj.Attributes() {
		if val == nil || val.IsNull() || val.IsUnknown() {
			continue
		}
		switch v := val.(type) {
		case types.String:
			result[key] = v.ValueString()
		case types.Int64:
			result[key] = int(v.ValueInt64())
		case types.Bool:
			result[key] = v.ValueBool()
		case types.Float64:
			result[key] = v.ValueFloat64()
		case types.List:
			elems := v.Elements()
			list := make([]any, len(elems))
			for i, e := range elems {
				if s, ok := e.(types.String); ok {
					list[i] = s.ValueString()
				}
			}
			result[key] = list
		}
	}
	return result
}
