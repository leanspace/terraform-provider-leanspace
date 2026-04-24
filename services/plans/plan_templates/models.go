package plan_templates

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type PlanTemplate struct {
	general_objects.AuditModel
	AssetId                    string                      `json:"assetId"`
	Name                       string                      `json:"name"`
	Description                *string                     `json:"description,omitempty"`
	IntegrityStatus            string                      `json:"integrityStatus"`
	ActivityConfigs            []ActivityConfigResult      `json:"activityConfigs,omitempty"`
	EstimatedDurationInSeconds int                         `json:"estimatedDurationInSeconds"`
	InvalidPlanTemplateReasons []InvalidPlanTemplateReason `json:"invalidPlanTemplateReasons,omitempty"`
}

type ActivityConfigResult struct {
	ActivityDefinitionId         string                            `json:"activityDefinitionId"`
	DelayReferenceOnPredecessor  *string                           `json:"delayReferenceOnPredecessor,omitempty"`
	Position                     int                               `json:"position"`
	DelayInSeconds               int                               `json:"delayInSeconds"`
	EstimatedDurationInSeconds   *int                              `json:"estimatedDurationInSeconds,omitempty"`
	Name                         *string                           `json:"name,omitempty"`
	Arguments                    []Argument                        `json:"arguments,omitempty"`
	ResourceFunctionFormulas     []ResourceFunctionFormulaOverload `json:"resourceFunctionFormulas,omitempty"`
	Tags                         []general_objects.KeyValue        `json:"tags,omitempty"`
	DefinitionLinkStatus         *string                           `json:"definitionLinkStatus,omitempty"`
	InvalidDefinitionLinkReasons []InvalidDefinitionLinkReason     `json:"invalidDefinitionLinkReasons,omitempty"`
}

type InvalidPlanTemplateReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Argument struct {
	Name       string                              `json:"name"`
	Attributes general_objects.ValueAttribute[any] `json:"attributes"`
}

type ResourceFunctionFormulaOverload struct {
	ResourceFunctionId *string                  `json:"resourceFunctionId,omitempty"`
	Formula            *ResourceFunctionFormula `json:"formula,omitempty"`
}

type ResourceFunctionFormula struct {
	Type      string   `json:"type"`
	Amplitude *float64 `json:"amplitude,omitempty"`
	Constant  *float64 `json:"constant,omitempty"`
	Rate      *float64 `json:"rate,omitempty"`
	TimeUnit  *string  `json:"timeUnit,omitempty"`
}

type InvalidDefinitionLinkReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
