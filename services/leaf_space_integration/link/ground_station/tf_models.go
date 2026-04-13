package groundstation_links

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

type LeafSpaceGroundStationLinkTF struct {
	ID                           types.String `tfsdk:"id"`
	LeafspaceGroundStationId     types.String `tfsdk:"leafspace_ground_station_id"`
	LeafspaceGroundStationName   types.String `tfsdk:"leafspace_ground_station_name"`
	LeanspaceGroundStationId     types.String `tfsdk:"leanspace_ground_station_id"`
	LeanspaceGroundStationName   types.String `tfsdk:"leanspace_ground_station_name"`
}

func (s *LeafSpaceGroundStationLink) ToTF() any {
	return &LeafSpaceGroundStationLinkTF{
		ID:                         helper.TFStringValue(s.ID),
		LeafspaceGroundStationId:   helper.TFStringValue(s.LeafspaceGroundStationId),
		LeafspaceGroundStationName: helper.TFStringValue(s.LeafspaceGroundStationName),
		LeanspaceGroundStationId:   helper.TFStringValue(s.LeanspaceGroundStationId),
		LeanspaceGroundStationName: helper.TFStringValue(s.LeanspaceGroundStationName),
	}
}

func (tf *LeafSpaceGroundStationLinkTF) ToAPI() any {
	return &LeafSpaceGroundStationLink{
		ID:                         helper.FromTFString(tf.ID),
		LeafspaceGroundStationId:   helper.FromTFString(tf.LeafspaceGroundStationId),
		LeafspaceGroundStationName: helper.FromTFString(tf.LeafspaceGroundStationName),
		LeanspaceGroundStationId:   helper.FromTFString(tf.LeanspaceGroundStationId),
		LeanspaceGroundStationName: helper.FromTFString(tf.LeanspaceGroundStationName),
	}
}
