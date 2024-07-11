package plan_templates

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type PlanTemplate struct {
	ID                         string                      `json:"id"`
	AssetId                    string                      `json:"assetId"`
	Name                       string                      `json:"name"`
	Description                string                      `json:"description,omitempty"`
	IntegrityStatus            string                      `json:"integrityStatus"`
	ActivityConfigs            []ActivityConfigResult      `json:"activityConfigs,omitempty"`
	EstimatedDurationInSeconds int                         `json:"estimatedDurationInSeconds"`
	InvalidPlanTemplateReasons []InvalidPlanTemplateReason `json:"invalidPlanTemplateReasons"`
	CreatedAt                  string                      `json:"createdAt"`
	CreatedBy                  string                      `json:"createdBy"`
	LastModifiedAt             string                      `json:"lastModifiedAt"`
	LastModifiedBy             string                      `json:"lastModifiedBy"`
}

func (template *PlanTemplate) GetID() string { return template.ID }

type ActivityConfigResult struct {
	ActivityDefinitionId        string                          `json:"activityDefinitionId"`
	DelayReferenceOnPredecessor string                          `json:"delayReferenceOnPredecessor"`
	Position                    int                             `json:"position"`
	DelayInSeconds              int                             `json:"delayInSeconds"`
	EstimatedDurationInSeconds  int                             `json:"estimatedDurationInSeconds"`
	Name                        string                          `json:"name"`
	Arguments                   []Argument                      `json:"arguments"`
	ResourceFunctionFormulas    ResourceFunctionFormulaOverload `json:"resourceFunctionFormulas,omitempty"`
	Tags                        []general_objects.KeyValue      `json:"tags,omitempty"`

	DefinitionLinkStatus         string                        `json:"definitionLinkStatus"`
	InvalidDefinitionLinkReasons []InvalidDefinitionLinkReason `json:"invalidDefinitionLinkReasons"`
}

type InvalidPlanTemplateReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Argument struct {
	Name       string                                `json:"name"`
	Attributes []general_objects.ValueAttribute[any] `json:"attributes"`
}

type ResourceFunctionFormulaOverload struct {
	ResourceFunctionId string                  `json:"resourceFunctionId"`
	Formula            ResourceFunctionFormula `json:"formula"`
}

type ResourceFunctionFormula struct {
	Type string `json:"type"`
}

type InvalidDefinitionLinkReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
