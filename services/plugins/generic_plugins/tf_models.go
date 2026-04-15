package generic_plugins

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

// sourceCodeLinkAttrTypes matches sourceCodeLinkSchema attribute names and types.
var sourceCodeLinkAttrTypes = map[string]attr.Type{
	"expiration_time": types.StringType,
	"source_code_id":  types.StringType,
	"url":             types.StringType,
}

type GenericPluginTF struct {
	general_objects.AuditModelTF
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
	Type           types.String `tfsdk:"type"`
	Language       types.String `tfsdk:"language"`
	SourceCodeLink types.List   `tfsdk:"source_code_link"`
	Status         types.String `tfsdk:"status"`
	FilePath       types.String `tfsdk:"source_code_path"`
	FileSha        types.String `tfsdk:"source_code_sha"`
}

func (x *GenericPlugin) ToTF() any {
	elemType := types.ObjectType{AttrTypes: sourceCodeLinkAttrTypes}
	linkObj, _ := types.ObjectValue(sourceCodeLinkAttrTypes, map[string]attr.Value{
		"expiration_time": helper.TFStringValue(x.SourceCodeLink.ExpirationTime),
		"source_code_id":  helper.TFStringValue(x.SourceCodeLink.SourceCodeId),
		"url":             helper.TFStringValue(x.SourceCodeLink.Url),
	})
	sourceCodeLink, _ := types.ListValue(elemType, []attr.Value{linkObj})

	return &GenericPluginTF{
		AuditModelTF:   general_objects.AuditModelToTF(&x.AuditModel),
		Name:           helper.TFStringValue(x.Name),
		Description:    helper.TFStringPtrValue(x.Description),
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
		Description: helper.FromTFStringPtr(tf.Description),
		Type:        helper.FromTFString(tf.Type),
		Language:    helper.FromTFString(tf.Language),
		Status:      helper.FromTFString(tf.Status),
		FilePath:    helper.FromTFString(tf.FilePath),
		FileSha:     helper.FromTFString(tf.FileSha),
	}
	elems := tf.SourceCodeLink.Elements()
	if len(elems) > 0 {
		if obj, ok := elems[0].(types.Object); ok {
			attrs := obj.Attributes()
			gp.SourceCodeLink = SourceCodeLink{
				ExpirationTime: helper.FromTFString(attrs["expiration_time"].(types.String)),
				SourceCodeId:   helper.FromTFString(attrs["source_code_id"].(types.String)),
				Url:            helper.FromTFString(attrs["url"].(types.String)),
			}
		}
	}
	return gp
}
