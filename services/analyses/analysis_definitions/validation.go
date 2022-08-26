package analysis_definitions

import (
	. "leanspace-terraform-provider/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var fieldValidators = Validators{
	Equivalence(
		Equals("type", "ARRAY"),
		And(IsSet("items"), Not(IsEmpty("items"))),
	),
	Equivalence(
		Equals("type", "STRUCTURE"),
		And(IsSet("fields"), Not(IsEmpty("fields"))),
	),
	Equivalence(
		Not(Or(Equals("type", "ARRAY"), Equals("type", "STRUCTURE"))),
		IsSet("source"),
	),
	Equivalence(
		Equals("source", "STATIC"),
		// value could be "unset" (eg. 0, "") and be valid
		And( /* IsSet("value"), */ Not(IsSet("ref"))),
	),
	Equivalence(
		Equals("source", "REFERENCE"),
		And(IsSet("ref"), Not(IsSet("value"))),
	),
}

func (analysisDefinition *AnalysisDefinition) Validate(data map[string]any) error {
	if err := validateField(data["inputs"].([]any)[0].(map[string]any)); err != nil {
		return err
	}
	return nil
}

func validateField(data map[string]any) error {
	err := fieldValidators.Check(data)
	if err != nil {
		return err
	}

	if data["type"] == "ARRAY" {
		for _, subfield := range data["items"].([]any) {
			if err := validateField(subfield.(map[string]any)); err != nil {
				return err
			}
		}
	}

	if data["type"] == "STRUCTURE" {
		for _, subfield := range data["fields"].(*schema.Set).List() {
			if err := validateField(subfield.(map[string]any)); err != nil {
				return err
			}
		}
	}

	return nil
}
