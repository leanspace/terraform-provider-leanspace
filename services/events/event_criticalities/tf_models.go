package event_criticalities

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type EventCriticalitiesTF struct {
	general_objects.AuditModelTF
	Name     types.String                `tfsdk:"name"`
	ReadOnly types.Bool                  `tfsdk:"read_only"`
	Tags     []general_objects.KeyValueTF `tfsdk:"tags"`
}

func (s *EventCriticalities) ToTF() any {
	return &EventCriticalitiesTF{
		AuditModelTF: general_objects.AuditModelToTF(&s.AuditModel),
		Name:         helper.TFStringValue(s.Name),
		ReadOnly:     helper.TFBoolValue(s.ReadOnly),
		Tags:         general_objects.KeyValuesToTF(s.Tags),
	}
}

func (tf *EventCriticalitiesTF) ToAPI() any {
	return &EventCriticalities{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		ReadOnly:   helper.FromTFBool(tf.ReadOnly),
		Tags:       general_objects.KeyValuesFromTF(tf.Tags),
	}
}
