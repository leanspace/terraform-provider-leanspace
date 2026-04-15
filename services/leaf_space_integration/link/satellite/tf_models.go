package satellite_links

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type LeafSpaceSatelliteLinkTF struct {
	ID                     types.String `tfsdk:"id"`
	LeafspaceSatelliteId   types.String `tfsdk:"leafspace_satellite_id"`
	LeafspaceSatelliteName types.String `tfsdk:"leafspace_satellite_name"`
	LeanspaceSatelliteId   types.String `tfsdk:"leanspace_satellite_id"`
	LeanspaceSatelliteName types.String `tfsdk:"leanspace_satellite_name"`
}

func (s *LeafSpaceSatelliteLink) ToTF() any {
	return general_objects.ReflectToTF[LeafSpaceSatelliteLinkTF](s)
}

func (tf *LeafSpaceSatelliteLinkTF) ToAPI() any {
	return general_objects.ReflectFromTF[LeafSpaceSatelliteLink](tf)
}
