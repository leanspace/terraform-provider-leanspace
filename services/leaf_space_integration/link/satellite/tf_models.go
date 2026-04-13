package satellite_links

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

type LeafSpaceSatelliteLinkTF struct {
	ID                      types.String `tfsdk:"id"`
	LeafspaceSatelliteId    types.String `tfsdk:"leafspace_satellite_id"`
	LeafspaceSatelliteName  types.String `tfsdk:"leafspace_satellite_name"`
	LeanspaceSatelliteId    types.String `tfsdk:"leanspace_satellite_id"`
	LeanspaceSatelliteName  types.String `tfsdk:"leanspace_satellite_name"`
}

func (s *LeafSpaceSatelliteLink) ToTF() any {
	return &LeafSpaceSatelliteLinkTF{
		ID:                     helper.TFStringValue(s.ID),
		LeafspaceSatelliteId:   helper.TFStringValue(s.LeafspaceSatelliteId),
		LeafspaceSatelliteName: helper.TFStringValue(s.LeafspaceSatelliteName),
		LeanspaceSatelliteId:   helper.TFStringValue(s.LeanspaceSatelliteId),
		LeanspaceSatelliteName: helper.TFStringValue(s.LeanspaceSatelliteName),
	}
}

func (tf *LeafSpaceSatelliteLinkTF) ToAPI() any {
	return &LeafSpaceSatelliteLink{
		ID:                     helper.FromTFString(tf.ID),
		LeafspaceSatelliteId:   helper.FromTFString(tf.LeafspaceSatelliteId),
		LeafspaceSatelliteName: helper.FromTFString(tf.LeafspaceSatelliteName),
		LeanspaceSatelliteId:   helper.FromTFString(tf.LeanspaceSatelliteId),
		LeanspaceSatelliteName: helper.FromTFString(tf.LeanspaceSatelliteName),
	}
}
