package plugins

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type PluginTF struct {
	general_objects.AuditModelTF
	Type                             types.String `tfsdk:"type"`
	ImplementationClassName          types.String `tfsdk:"implementation_class_name"`
	Name                             types.String `tfsdk:"name"`
	Description                      types.String `tfsdk:"description"`
	SourceCodeFileDownloadAuthorized types.Bool   `tfsdk:"source_code_file_download_authorized"`
	FilePath                         types.String `tfsdk:"file_path"`
	SdkVersion                       types.String `tfsdk:"sdk_version"`
	SdkVersionFamily                 types.String `tfsdk:"sdk_version_family"`
	Status                           types.String `tfsdk:"status"`
	FileSha                          types.String `tfsdk:"file_sha"`
}

func (x *Plugin) ToTF() any {
	return &PluginTF{
		AuditModelTF:                    general_objects.AuditModelToTF(&x.AuditModel),
		Type:                            helper.TFStringValue(x.Type),
		ImplementationClassName:         helper.TFStringValue(x.ImplementationClassName),
		Name:                            helper.TFStringValue(x.Name),
		Description:                     helper.TFStringValue(x.Description),
		SourceCodeFileDownloadAuthorized: helper.TFBoolValue(x.SourceCodeFileDownloadAuthorized),
		FilePath:                        helper.TFStringValue(x.FilePath),
		SdkVersion:                      helper.TFStringValue(x.SdkVersion),
		SdkVersionFamily:                helper.TFStringValue(x.SdkVersionFamily),
		Status:                          helper.TFStringValue(x.Status),
		FileSha:                         helper.TFStringValue(x.FileSha),
	}
}

func (tf *PluginTF) ToAPI() any {
	return &Plugin{
		AuditModel:                       general_objects.AuditModelFromTF(tf.AuditModelTF),
		Type:                             helper.FromTFString(tf.Type),
		ImplementationClassName:          helper.FromTFString(tf.ImplementationClassName),
		Name:                             helper.FromTFString(tf.Name),
		Description:                      helper.FromTFString(tf.Description),
		SourceCodeFileDownloadAuthorized: helper.FromTFBool(tf.SourceCodeFileDownloadAuthorized),
		FilePath:                         helper.FromTFString(tf.FilePath),
		SdkVersion:                       helper.FromTFString(tf.SdkVersion),
		SdkVersionFamily:                 helper.FromTFString(tf.SdkVersionFamily),
		Status:                           helper.FromTFString(tf.Status),
		FileSha:                          helper.FromTFString(tf.FileSha),
	}
}
