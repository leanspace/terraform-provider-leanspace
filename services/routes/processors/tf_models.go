package processors

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ProcessorTF struct {
	general_objects.AuditModelTF
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Version     types.String `tfsdk:"version"`
	Type        types.String `tfsdk:"type"`
	FilePath    types.String `tfsdk:"file_path"`
	FileSha     types.String `tfsdk:"file_sha"`
}

func (x *Processor) ToTF() any {
	return &ProcessorTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Description:  helper.TFStringPtrValue(x.Description),
		Version:      helper.TFStringValue(x.Version),
		Type:         helper.TFStringValue(x.Type),
		FilePath:     helper.TFStringValue(x.FilePath),
		FileSha:      helper.TFStringValue(x.FileSha),
	}
}

func (tf *ProcessorTF) ToAPI() any {
	return &Processor{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:        helper.FromTFString(tf.Name),
		Description: helper.FromTFStringPtr(tf.Description),
		Version:     helper.FromTFString(tf.Version),
		Type:        helper.FromTFString(tf.Type),
		FilePath:    helper.FromTFString(tf.FilePath),
		FileSha:     helper.FromTFString(tf.FileSha),
	}
}
