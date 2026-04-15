package groundstation_links

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type LeafSpaceGroundStationLinkTF struct {
	ID                         types.String `tfsdk:"id"`
	LeafspaceGroundStationId   types.String `tfsdk:"leafspace_ground_station_id"`
	LeafspaceGroundStationName types.String `tfsdk:"leafspace_ground_station_name"`
	LeanspaceGroundStationId   types.String `tfsdk:"leanspace_ground_station_id"`
	LeanspaceGroundStationName types.String `tfsdk:"leanspace_ground_station_name"`
}

func (s *LeafSpaceGroundStationLink) ToTF() any {
	return general_objects.ReflectToTF[LeafSpaceGroundStationLinkTF](s)
}

func (tf *LeafSpaceGroundStationLinkTF) ToAPI() any {
	return general_objects.ReflectFromTF[LeafSpaceGroundStationLink](tf)
}
