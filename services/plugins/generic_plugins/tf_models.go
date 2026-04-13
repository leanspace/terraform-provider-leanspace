package generic_plugins

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type GenericPluginTF struct {
	general_objects.AuditModelTF
	Name           types.String       `tfsdk:"name"`
	Description    types.String       `tfsdk:"description"`
	Type           types.String       `tfsdk:"type"`
	Language       types.String       `tfsdk:"language"`
	SourceCodeLink []SourceCodeLinkTF `tfsdk:"source_code_link"`
	Status         types.String       `tfsdk:"status"`
	FilePath       types.String       `tfsdk:"source_code_path"`
	FileSha        types.String       `tfsdk:"source_code_sha"`
}

type SourceCodeLinkTF struct {
	ExpirationTime types.String `tfsdk:"expiration_time"`
	SourceCodeId   types.String `tfsdk:"source_code_id"`
	Url            types.String `tfsdk:"url"`
}

func (x *GenericPlugin) ToTF() any {
	var sourceCodeLink []SourceCodeLinkTF
	sourceCodeLink = []SourceCodeLinkTF{{
		ExpirationTime: helper.TFStringValue(x.SourceCodeLink.ExpirationTime),
		SourceCodeId:   helper.TFStringValue(x.SourceCodeLink.SourceCodeId),
		Url:            helper.TFStringValue(x.SourceCodeLink.Url),
	}}

	return &GenericPluginTF{
		AuditModelTF:   general_objects.AuditModelToTF(&x.AuditModel),
		Name:           helper.TFStringValue(x.Name),
		Description:    helper.TFStringValue(x.Description),
		Type:           helper.TFStringValue(x.Type),
		Language:       helper.TFStringValue(x.Language),
		SourceCodeLink: sourceCodeLink,
		Status:         helper.TFStringValue(x.Status),
		FilePath:       helper.TFStringValue(x.FilePath),
		FileSha:        helper.TFStringValue(x.FileSha),
	}
}

func (tf *GenericPluginTF) ToAPI() any {
	gp := &GenericPlugin{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:        helper.FromTFString(tf.Name),
		Description: helper.FromTFString(tf.Description),
		Type:        helper.FromTFString(tf.Type),
		Language:    helper.FromTFString(tf.Language),
		Status:      helper.FromTFString(tf.Status),
		FilePath:    helper.FromTFString(tf.FilePath),
		FileSha:     helper.FromTFString(tf.FileSha),
	}
	if len(tf.SourceCodeLink) > 0 {
		gp.SourceCodeLink = SourceCodeLink{
			ExpirationTime: helper.FromTFString(tf.SourceCodeLink[0].ExpirationTime),
			SourceCodeId:   helper.FromTFString(tf.SourceCodeLink[0].SourceCodeId),
			Url:            helper.FromTFString(tf.SourceCodeLink[0].Url),
		}
	}
	return gp
}
