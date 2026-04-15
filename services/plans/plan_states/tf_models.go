package plan_states

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type PlanStateTF struct {
	general_objects.AuditModelTF
	Name     types.String `tfsdk:"name"`
	ReadOnly types.Bool   `tfsdk:"read_only"`
}

func (s *PlanState) ToTF() any {
	return general_objects.ReflectToTF[PlanStateTF](s)
}

func (tf *PlanStateTF) ToAPI() any {
	return general_objects.ReflectFromTF[PlanState](tf)
}
