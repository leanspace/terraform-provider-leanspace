package action_templates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ActionTemplateTF struct {
	general_objects.AuditModelTF
	Name    types.String            `tfsdk:"name"`
	Type    types.String            `tfsdk:"type"`
	URL     types.String            `tfsdk:"url"`
	Payload types.String            `tfsdk:"payload"`
	Content types.String            `tfsdk:"content"`
	Headers map[string]types.String `tfsdk:"headers"`
}

func (x *ActionTemplate) ToTF() any {
	var headers map[string]types.String
	if x.Headers != nil {
		headers = make(map[string]types.String, len(x.Headers))
		for k, v := range x.Headers {
			headers[k] = helper.TFStringValue(v)
		}
	}
	return &ActionTemplateTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Type:         helper.TFStringValue(x.Type),
		URL:          helper.TFStringValue(x.URL),
		Payload:      helper.TFStringValue(x.Payload),
		Content:      helper.TFStringValue(x.Content),
		Headers:      headers,
	}
}

func (tf *ActionTemplateTF) ToAPI() any {
	var headers map[string]string
	if tf.Headers != nil {
		headers = make(map[string]string, len(tf.Headers))
		for k, v := range tf.Headers {
			headers[k] = helper.FromTFString(v)
		}
	}
	return &ActionTemplate{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		Type:       helper.FromTFString(tf.Type),
		URL:        helper.FromTFString(tf.URL),
		Payload:    helper.FromTFString(tf.Payload),
		Content:    helper.FromTFString(tf.Content),
		Headers:    headers,
	}
}
