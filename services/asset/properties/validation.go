package properties

import (
	. "github.com/leanspace/terraform-provider-leanspace/helper"
)

var propertyValidators = Validators{
	If(
		Or(IsSet("min"), IsSet("max"), IsSet("scale"), IsSet("precision"), IsSet("unit_id")),
		Equals("type", "NUMERIC"),
	),
	If(
		And(IsSet("options"), Not(IsEmpty("options"))),
		Equals("type", "ENUM"),
	),
	If(
		Or(IsSet("min_length"), IsSet("max_length"), IsSet("pattern")),
		Equals("type", "TEXT"),
	),
	If(
		Or(IsSet("before"), IsSet("after")),
		Or(Equals("type", "TIMESTAMP"), Equals("type", "DATE"), Equals("type", "TIME")),
	),
	If(
		And(IsSet("fields"), Not(IsEmpty("fields"))),
		Equals("type", "GEOPOINT"),
	),
}

func (property *Property[T]) Validate() error {
	obj := map[string]any{
		"type":       property.Attributes.Type,
		"min":        property.Attributes.Min,
		"max":        property.Attributes.Max,
		"scale":      property.Attributes.Scale,
		"precision":  property.Attributes.Precision,
		"unit_id":    property.Attributes.UnitId,
		"options":    property.Attributes.Options,
		"min_length": property.Attributes.MinLength,
		"max_length": property.Attributes.MaxLength,
		"pattern":    property.Attributes.Pattern,
		"before":     property.Attributes.Before,
		"after":      property.Attributes.After,
		"fields":     property.Attributes.Fields,
	}
	return propertyValidators.Check(obj)
}
