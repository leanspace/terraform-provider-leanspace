package general_objects

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var (
	typeAuditModel   = reflect.TypeOf(AuditModel{})
	typeAuditModelTF = reflect.TypeOf(AuditModelTF{})
	typeKeyValues    = reflect.TypeOf([]KeyValue{})
	typeKeyValueTFs  = reflect.TypeOf([]KeyValueTF{})
	typeString       = reflect.TypeOf("")
	typePtrString    = reflect.TypeOf((*string)(nil))
	typeBool         = reflect.TypeOf(false)
	typePtrBool      = reflect.TypeOf((*bool)(nil))
	typeInt          = reflect.TypeOf(int(0))
	typePtrInt       = reflect.TypeOf((*int)(nil))
	typeFloat64      = reflect.TypeOf(float64(0))
	typePtrFloat64   = reflect.TypeOf((*float64)(nil))
	typeStrings      = reflect.TypeOf([]string{})
	typeTFString     = reflect.TypeOf(types.String{})
	typeTFBool       = reflect.TypeOf(types.Bool{})
	typeTFInt64      = reflect.TypeOf(types.Int64{})
	typeTFFloat64    = reflect.TypeOf(types.Float64{})
	typeTFStrings    = reflect.TypeOf([]types.String{})
)

// ReflectToTF maps a flat API struct (or pointer) to a new TF struct.
//
// Supported API → TF field type pairs:
//
//	string /*string        → types.String
//	bool   /*bool          → types.Bool
//	int    /*int           → types.Int64
//	float64/*float64       → types.Float64
//	[]string               → []types.String
//	[]KeyValue             → []KeyValueTF
//	AuditModel (embedded)  → AuditModelTF (embedded)
//
// Fields in TF with no matching API field are left as their zero/null value.
// Use this for flat structs with purely mechanical conversions; write a hand-coded
// ToTF when custom logic (conditional branches, types.List, etc.) is needed.
func ReflectToTF[TF any](apiPtr any) *TF {
	tf := new(TF)
	apiVal := reflect.ValueOf(apiPtr)
	if apiVal.Kind() == reflect.Pointer {
		apiVal = apiVal.Elem()
	}
	tfVal := reflect.ValueOf(tf).Elem()
	tfType := tfVal.Type()

	for i := 0; i < tfType.NumField(); i++ {
		tfField := tfType.Field(i)
		tfFieldVal := tfVal.Field(i)

		// Embedded AuditModelTF ← AuditModel
		if tfField.Anonymous && tfField.Type == typeAuditModelTF {
			apiField := apiVal.FieldByName("AuditModel")
			if apiField.IsValid() && apiField.Type() == typeAuditModel {
				audit := apiField.Interface().(AuditModel)
				tfFieldVal.Set(reflect.ValueOf(AuditModelToTF(&audit)))
			}
			continue
		}

		// General field: match by Go field name
		apiField := apiVal.FieldByName(tfField.Name)
		if !apiField.IsValid() {
			continue
		}
		converted := toTFValue(apiField)
		if converted.IsValid() && converted.Type().AssignableTo(tfField.Type) {
			tfFieldVal.Set(converted)
		}
	}
	return tf
}

// ReflectFromTF maps a flat TF struct (or pointer) back to a new API struct.
// See ReflectToTF for the list of supported field type pairs.
func ReflectFromTF[API any](tfPtr any) *API {
	api := new(API)
	tfVal := reflect.ValueOf(tfPtr)
	if tfVal.Kind() == reflect.Pointer {
		tfVal = tfVal.Elem()
	}
	apiVal := reflect.ValueOf(api).Elem()
	apiType := apiVal.Type()

	for i := 0; i < apiType.NumField(); i++ {
		apiField := apiType.Field(i)
		apiFieldVal := apiVal.Field(i)

		// Embedded AuditModel ← AuditModelTF
		if apiField.Anonymous && apiField.Type == typeAuditModel {
			tfField := tfVal.FieldByName("AuditModelTF")
			if tfField.IsValid() && tfField.Type() == typeAuditModelTF {
				auditTF := tfField.Interface().(AuditModelTF)
				apiFieldVal.Set(reflect.ValueOf(AuditModelFromTF(auditTF)))
			}
			continue
		}

		// General field: match by Go field name
		tfField := tfVal.FieldByName(apiField.Name)
		if !tfField.IsValid() {
			continue
		}
		converted := fromTFValue(tfField, apiField.Type)
		if converted.IsValid() && converted.Type().AssignableTo(apiField.Type) {
			apiFieldVal.Set(converted)
		}
	}
	return api
}

func toTFValue(v reflect.Value) reflect.Value {
	switch v.Type() {
	case typeString:
		return reflect.ValueOf(types.StringValue(v.String()))
	case typePtrString:
		if v.IsNil() {
			return reflect.ValueOf(types.StringNull())
		}
		return reflect.ValueOf(types.StringValue(v.Elem().String()))
	case typeBool:
		return reflect.ValueOf(types.BoolValue(v.Bool()))
	case typePtrBool:
		if v.IsNil() {
			return reflect.ValueOf(types.BoolNull())
		}
		return reflect.ValueOf(types.BoolValue(v.Elem().Bool()))
	case typeInt:
		return reflect.ValueOf(types.Int64Value(v.Int()))
	case typePtrInt:
		if v.IsNil() {
			return reflect.ValueOf(types.Int64Null())
		}
		return reflect.ValueOf(types.Int64Value(v.Elem().Int()))
	case typeFloat64:
		return reflect.ValueOf(types.Float64Value(v.Float()))
	case typePtrFloat64:
		if v.IsNil() {
			return reflect.ValueOf(types.Float64Null())
		}
		return reflect.ValueOf(types.Float64Value(v.Elem().Float()))
	case typeStrings:
		return reflect.ValueOf(helper.TFStringsValue(v.Interface().([]string)))
	case typeKeyValues:
		return reflect.ValueOf(KeyValuesToTF(v.Interface().([]KeyValue)))
	}
	return reflect.Value{}
}

func fromTFValue(v reflect.Value, destType reflect.Type) reflect.Value {
	switch v.Type() {
	case typeTFString:
		s := v.Interface().(types.String)
		switch destType {
		case typeString:
			return reflect.ValueOf(helper.FromTFString(s))
		case typePtrString:
			return reflect.ValueOf(helper.FromTFStringPtr(s))
		}
	case typeTFBool:
		b := v.Interface().(types.Bool)
		switch destType {
		case typeBool:
			return reflect.ValueOf(helper.FromTFBool(b))
		case typePtrBool:
			return reflect.ValueOf(helper.FromTFBoolPtr(b))
		}
	case typeTFInt64:
		i := v.Interface().(types.Int64)
		switch destType {
		case typeInt:
			return reflect.ValueOf(helper.FromTFInt64(i))
		case typePtrInt:
			return reflect.ValueOf(helper.FromTFIntPtr(i))
		}
	case typeTFFloat64:
		f := v.Interface().(types.Float64)
		switch destType {
		case typeFloat64:
			return reflect.ValueOf(helper.FromTFFloat64(f))
		case typePtrFloat64:
			return reflect.ValueOf(helper.FromTFFloat64Ptr(f))
		}
	case typeTFStrings:
		return reflect.ValueOf(helper.FromTFStrings(v.Interface().([]types.String)))
	case typeKeyValueTFs:
		return reflect.ValueOf(KeyValuesFromTF(v.Interface().([]KeyValueTF)))
	}
	return reflect.Value{}
}
