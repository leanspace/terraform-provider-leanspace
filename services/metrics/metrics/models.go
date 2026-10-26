package metrics

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Metric[T any] struct {
	ID             string                                 `json:"id"`
	Name           string                                 `json:"name"`
	Description    string                                 `json:"description,omitempty"`
	NodeId         string                                 `json:"nodeId"`
	CreatedAt      string                                 `json:"createdAt"`
	CreatedBy      string                                 `json:"createdBy"`
	LastModifiedAt string                                 `json:"lastModifiedAt"`
	LastModifiedBy string                                 `json:"lastModifiedBy"`
	Tags           []general_objects.KeyValue             `json:"tags,omitempty"`
	Attributes     general_objects.DefinitionAttribute[T] `json:"attributes"`
	Calibration    *Calibration                           `json:"calibration"`
}

func (metric *Metric[T]) GetID() string { return metric.ID }

type CalibrationType string

const (
	CalibrationTypePolynomial  CalibrationType = "POLYNOMIAL"
	CalibrationTypeLog         CalibrationType = "LOG"
	CalibrationTypeInverseLog  CalibrationType = "INVERSE_LOG"
	CalibrationTypeCapping     CalibrationType = "CAPPING"
	CalibrationTypeRange       CalibrationType = "RANGE"
	CalibrationTypeTextMapping CalibrationType = "TEXT_MAPPING"
)

var ValidCalibrationTypes = []string{
	string(CalibrationTypePolynomial),
	string(CalibrationTypeLog),
	string(CalibrationTypeInverseLog),
	string(CalibrationTypeCapping),
	string(CalibrationTypeRange),
	string(CalibrationTypeTextMapping),
}

type Calibration struct {
	Type         CalibrationType    `json:"type"`
	Description  string             `json:"description,omitempty"`
	Coefficients []float64          `json:"coefficients,omitempty"`
	Base         float64            `json:"base,omitempty"`
	Capping      *CalibrationBounds `json:"capping,omitempty"`
	Min          *float64           `json:"min,omitempty"`
	Max          *float64           `json:"max,omitempty"`
	Ranges       []CalibrationRange `json:"ranges,omitempty"`
	Mappings     map[string]string  `json:"mappings,omitempty"`
}

type CalibrationBounds struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

type CalibrationRange struct {
	Min   *float64 `json:"min,omitempty"`
	Max   *float64 `json:"max,omitempty"`
	Value float64  `json:"value"`
}
