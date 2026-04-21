package action_templates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	headers := make(map[string]types.String)
	for k, v := range x.Headers {
		headers[k] = types.StringValue(v)
	}
	return &ActionTemplateTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         types.StringValue(x.Name),
		Type:         types.StringValue(x.Type),
		URL:          types.StringPointerValue(x.URL),
		Payload:      types.StringPointerValue(x.Payload),
		Content:      types.StringPointerValue(x.Content),
		Headers:      headers,
	}
}

func (tf *ActionTemplateTF) ToAPI() any {
	var headers map[string]string
	if tf.Headers != nil {
		headers = make(map[string]string, len(tf.Headers))
		for k, v := range tf.Headers {
			headers[k] = v.ValueString()
		}
	}
	return &ActionTemplate{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       tf.Name.ValueString(),
		Type:       tf.Type.ValueString(),
		URL:        tf.URL.ValueStringPointer(),
		Payload:    tf.Payload.ValueStringPointer(),
		Content:    tf.Content.ValueStringPointer(),
		Headers:    headers,
	}
}
