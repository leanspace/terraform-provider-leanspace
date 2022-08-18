package analysis_definitions

import (
	"encoding/json"
	"fmt"
	"leanspace-terraform-provider/helper"
)

func (analysisDefinition *AnalysisDefinition) Validate() error {
	data, _ := json.Marshal(analysisDefinition)
	helper.Logger.Printf("Validating, %v", string(data))
	if err := analysisDefinition.Inputs.Validate(); err != nil {
		return err
	}
	return nil
}

func (field *Field) Validate() error {
	if field.Type != "ARRAY" && field.Items != nil && len(field.Items) != 0 {
		return fmt.Errorf("items must only be set if type is ARRAY (is %q), got %v", field.Type, field.Items)
	}

	if field.Type != "STRUCTURE" && field.Fields != nil && len(field.Fields) != 0 {
		return fmt.Errorf("fields must only be set if type is STRUCTURE (is %q), got %v", field.Type, field.Items)
	}

	if field.Type != "STRUCTURE" && field.Type != "ARRAY" {
		if field.Source == "" {
			return fmt.Errorf("source must be set for field of type %q, but got %v", field.Type, field.Source)
		}
		if field.Source == "STATIC" && field.Value == "" {
			return fmt.Errorf("value must be set if field source is STATIC")
		}
		if field.Source == "STATIC" && field.Ref != "" {
			return fmt.Errorf("ref mustn't be set if field source is STATIC, got %q", field.Ref)
		}
		if field.Source == "REFERENCE" && field.Ref == "" {
			return fmt.Errorf("ref must be set if field source is REFERENCE")
		}
		if field.Source == "REFERENCE" && field.Value != nil && field.Value != "" {
			return fmt.Errorf("value mustn't be set if field source is REFERENCE, got %q", field.Value)
		}
	}

	if field.Type == "ARRAY" {
		for _, subfield := range field.Items {
			if err := subfield.Validate(); err != nil {
				return err
			}
		}
	}

	if field.Type == "STRUCTURE" {
		for _, subfield := range field.Fields {
			if err := subfield.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
