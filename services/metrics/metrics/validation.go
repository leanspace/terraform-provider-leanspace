package metrics

import (
	"fmt"

	. "github.com/leanspace/terraform-provider-leanspace/helper"
)

var isCoefficientsCalibration = Or(
	Equals("type", string(CalibrationTypePolynomial)),
	Equals("type", string(CalibrationTypeLog)),
	Equals("type", string(CalibrationTypeInverseLog)),
)

// The API ignores fields that do not belong to the calibration type, which would cause a permanent diff.
var calibrationValidators = Validators{
	Equivalence(
		isCoefficientsCalibration,
		Not(IsEmpty("coefficients")),
	),
	Equivalence(
		Equals("type", string(CalibrationTypeRange)),
		Not(IsEmpty("ranges")),
	),
	Equivalence(
		Equals("type", string(CalibrationTypeTextMapping)),
		Not(IsEmpty("mappings")),
	),
	If(
		IsSet("base"),
		Or(Equals("type", string(CalibrationTypeLog)), Equals("type", string(CalibrationTypeInverseLog))),
	),
	If(
		Not(IsEmpty("capping")),
		Or(isCoefficientsCalibration, Equals("type", string(CalibrationTypeRange))),
	),
	If(
		Or(Not(IsEmpty("min")), Not(IsEmpty("max"))),
		Equals("type", string(CalibrationTypeCapping)),
	),
}

func (metric *Metric[T]) Validate(obj map[string]any) error {
	if len(obj["calibration"].([]any)) == 0 || obj["calibration"].([]any)[0] == nil {
		return nil
	}
	calibration := obj["calibration"].([]any)[0].(map[string]any)
	if err := calibrationValidators.Check(calibration); err != nil {
		return err
	}

	calibrationType := calibration["type"].(string)
	expectedAttributeType := "NUMERIC"
	if calibrationType == string(CalibrationTypeTextMapping) {
		expectedAttributeType = "TEXT"
	}
	if len(obj["attributes"].([]any)) > 0 && obj["attributes"].([]any)[0] != nil {
		attributeType := obj["attributes"].([]any)[0].(map[string]any)["type"].(string)
		if attributeType != expectedAttributeType {
			return fmt.Errorf("a %s calibration can only be used on %s metrics, got %s", calibrationType, expectedAttributeType, attributeType)
		}
	}

	return nil
}
