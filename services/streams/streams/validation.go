package streams

import (
	. "leanspace-terraform-provider/helper"
)

var streamComponentValidators = Validators{
	If(
		Or(IsSet("length_in_bits"), IsSet("processor"), IsSet("data_type"), IsSet("endianness"), IsSet("unit_id")),
		Equals("type", "FIELD"),
	),
	If(
		And(IsSet("elements"), Not(IsEmpty("elements"))),
		Or(Equals("type", "SWITCH"), Equals("type", "CONTAINER")),
	),
	If(
		And(IsSet("expression"), Not(IsEmpty("expression"))),
		Equals("type", "SWITCH"),
	),
	If(
		Or(Equals("data_type", "INTEGER"), Equals("data_type", "UINTEGER")),
		LessThanEq("length_in_bits", 32),
	),
	If(
		Equals("data_type", "DECIMAL"),
		Or(Equals("length_in_bits", 0), Equals("length_in_bits", 32), Equals("length_in_bits", 64)),
	),
}

func (stream *Stream) Validate(obj map[string]any) error {
	configMap := obj["configuration"].([]any)[0].(map[string]any)
	structureMap := configMap["structure"].([]any)[0].(map[string]any)
	for _, elem := range structureMap["elements"].([]any) {
		if err := validateStreamComponent(elem.(map[string]any)); err != nil {
			return err
		}
	}

	return streamComponentValidators.Check(obj)
}

func validateStreamComponent(obj map[string]any) error {
	if err := streamComponentValidators.Check(obj); err != nil {
		return err
	}
	if obj["type"] == "SWITCH" || obj["type"] == "CONTAINER" {
		for _, subComponent := range obj["elements"].([]any) {
			if err := validateStreamComponent(subComponent.(map[string]any)); err != nil {
				return err
			}
		}
	}
	return nil
}
