package properties

import (
	. "leanspace-terraform-provider/helper"
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

func (property *Property[T]) Validate(obj map[string]any) error {
	return propertyValidators.Check(obj)
}
