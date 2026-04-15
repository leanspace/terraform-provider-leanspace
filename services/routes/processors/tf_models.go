package processors

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	return general_objects.ReflectToTF[ProcessorTF](x)
}

func (tf *ProcessorTF) ToAPI() any {
	return general_objects.ReflectFromTF[Processor](tf)
}
