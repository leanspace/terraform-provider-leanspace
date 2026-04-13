package properties

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (property *Property[T]) ToMap() map[string]any {
	propertyMap := property.ToAuditMap()
	propertyMap["name"] = helper.NilIfEmpty(property.Name)
	propertyMap["description"] = helper.NilIfEmpty(property.Description)
	propertyMap["built_in"] = helper.NilIfEmpty(property.IsBuiltIn)
	propertyMap["node_id"] = helper.NilIfEmpty(property.NodeId)
	propertyMap["tags"] = helper.ParseToMaps(property.Tags)
	propertyMap["type"] = helper.NilIfEmpty(property.Attributes.Type)
	switch property.Attributes.Type {
	case "NUMERIC":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = helper.ParseFloat(any(property.Attributes.Value).(float64))
		}
		propertyMap["min"] = helper.Float64PtrToAny(property.Attributes.Min)
		propertyMap["max"] = helper.Float64PtrToAny(property.Attributes.Max)
		propertyMap["scale"] = helper.IntPtrToAny(property.Attributes.Scale)
		propertyMap["precision"] = helper.IntPtrToAny(property.Attributes.Precision)
		propertyMap["unit_id"] = helper.NilIfEmpty(property.Attributes.UnitId)
	case "ENUM":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = helper.ParseFloat(any(property.Attributes.Value).(float64))
		}
		if property.Attributes.Options != nil {
			propertyMap["options"] = *property.Attributes.Options
		}
	case "TEXT":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = helper.NilIfEmpty(property.Attributes.Value)
		}
		propertyMap["min_length"] = helper.IntPtrToAny(property.Attributes.MinLength)
		propertyMap["max_length"] = helper.IntPtrToAny(property.Attributes.MaxLength)
		propertyMap["pattern"] = helper.NilIfEmpty(property.Attributes.Pattern)
	case "TIMESTAMP", "DATE", "TIME":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = helper.NilIfEmpty(property.Attributes.Value)
		}
		propertyMap["before"] = helper.NilIfEmpty(property.Attributes.Before)
		propertyMap["after"] = helper.NilIfEmpty(property.Attributes.After)
	case "BOOLEAN":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = strconv.FormatBool(any(property.Attributes.Value).(bool))
		}
	case "GEOPOINT":
		if property.Attributes.Fields != nil {
			fieldMap := make(map[string]any)
			elevationMap := (&property.Attributes.Fields.Elevation).ToMap()
			fieldMap["elevation"] = elevationMap
			latitudeMap := (&property.Attributes.Fields.Latitude).ToMap()
			fieldMap["latitude"] = latitudeMap
			longitudeMap := (&property.Attributes.Fields.Longitude).ToMap()
			fieldMap["longitude"] = longitudeMap
			propertyMap["fields"] = fieldMap
		}
	case "TLE":
		if any(property.Attributes.Value) != nil {
			var tleValue string
			var tleValues []interface{} = any(property.Attributes.Value).([]interface{})
			for _, value := range tleValues {
				tleValue = tleValue + "," + fmt.Sprint(value)
			}
			propertyMap["value"] = strings.TrimPrefix(tleValue, ",")
		}
	}
	return propertyMap
}

func (property *Property[T]) FromMap(propertyMap map[string]any) error {
	property.FromAuditMap(propertyMap)
	property.Name = helper.CastString(propertyMap, "name")
	property.IsBuiltIn = helper.CastBool(propertyMap, "built_in")
	property.NodeId = helper.CastString(propertyMap, "node_id")
	property.Description = helper.CastString(propertyMap, "description")
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(propertyMap, "tags")); err != nil {
		return err
	} else {
		property.Tags = tags
	}
	property.Attributes.Type = helper.CastString(propertyMap, "type")
	switch property.Attributes.Type {
	case "NUMERIC":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
		property.Attributes.Min = helper.CastFloat64Ptr(propertyMap, "min")
		property.Attributes.Max = helper.CastFloat64Ptr(propertyMap, "max")
		property.Attributes.Scale = helper.CastIntPtr(propertyMap, "scale")
		property.Attributes.Precision = helper.CastIntPtr(propertyMap, "precision")
		property.Attributes.UnitId = helper.CastString(propertyMap, "unit_id")
	case "ENUM":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
		if propertyMap["options"] != nil {
			option := helper.CastMapAny(propertyMap, "options")
			property.Attributes.Options = &option
		}
	case "TEXT":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
		property.Attributes.MinLength = helper.CastIntPtr(propertyMap, "min_length")
		property.Attributes.MaxLength = helper.CastIntPtr(propertyMap, "max_length")
		property.Attributes.Pattern = helper.CastString(propertyMap, "pattern")
	case "TIMESTAMP", "DATE", "TIME":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
		property.Attributes.Before = helper.CastString(propertyMap, "before")
		property.Attributes.After = helper.CastString(propertyMap, "after")
	case "BOOLEAN":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
	case "TLE":
		if tleValue, ok := propertyMap["value"]; ok {
			var stringTleValues []string = strings.Split(tleValue.(string), ",")
			if len(stringTleValues) == 2 {
				var interfaceOfTleValues []interface{}
				for _, str := range stringTleValues {
					var stringValue = strings.TrimSpace(str)
					interfaceOfTleValues = append(interfaceOfTleValues, stringValue)
				}
				property.Attributes.Value = any(interfaceOfTleValues).(T)
			}
		}
	case "GEOPOINT":
		if propertyMap["fields"] != nil {
			fields := helper.CastMapAny(propertyMap, "fields")
			property.Attributes.Fields = &general_objects.Fields{}
			property.Attributes.Fields.Elevation.FromMap(helper.CastMapAny(fields, "elevation"))
			property.Attributes.Fields.Latitude.FromMap(helper.CastMapAny(fields, "latitude"))
			property.Attributes.Fields.Longitude.FromMap(helper.CastMapAny(fields, "longitude"))
		}
	}

	return nil
}
