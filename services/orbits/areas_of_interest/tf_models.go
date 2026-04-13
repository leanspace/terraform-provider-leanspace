package areas_of_interest

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type AreaOfInterestTF struct {
	general_objects.AuditModelTF
	Name  types.String                 `tfsdk:"name"`
	Shape *AreaOfInterestShapeTF       `tfsdk:"shape"`
	Tags  []general_objects.KeyValueTF `tfsdk:"tags"`
}

type AreaOfInterestShapeTF struct {
	Type                types.String  `tfsdk:"type"`
	Geolocation         *GeoPointTF   `tfsdk:"geolocation"`
	CenterGeolocation   *GeoPointTF   `tfsdk:"center_geolocation"`
	RadiusInMeters      types.Float64 `tfsdk:"radius_in_meters"`
	VerticesGeolocation []GeoPointTF  `tfsdk:"vertices_geolocation"`
}

type GeoPointTF struct {
	Latitude  types.Float64 `tfsdk:"latitude"`
	Longitude types.Float64 `tfsdk:"longitude"`
	Altitude  types.Float64 `tfsdk:"altitude"`
}

func geoPointToTF(gp *GeoPoint) *GeoPointTF {
	if gp == nil {
		return nil
	}
	return &GeoPointTF{
		Latitude:  helper.TFFloat64Value(gp.Latitude),
		Longitude: helper.TFFloat64Value(gp.Longitude),
		Altitude:  helper.TFFloat64Value(gp.Altitude),
	}
}

func geoPointFromTF(tf *GeoPointTF) *GeoPoint {
	if tf == nil {
		return nil
	}
	return &GeoPoint{
		Latitude:  helper.FromTFFloat64(tf.Latitude),
		Longitude: helper.FromTFFloat64(tf.Longitude),
		Altitude:  helper.FromTFFloat64(tf.Altitude),
	}
}

func (x *AreaOfInterest) ToTF() any {
	var shape *AreaOfInterestShapeTF
	if x.Shape != nil {
		verticesGeolocation := make([]GeoPointTF, len(x.Shape.VerticesGeolocation))
		for i, v := range x.Shape.VerticesGeolocation {
			verticesGeolocation[i] = GeoPointTF{
				Latitude:  helper.TFFloat64Value(v.Latitude),
				Longitude: helper.TFFloat64Value(v.Longitude),
				Altitude:  helper.TFFloat64Value(v.Altitude),
			}
		}
		shape = &AreaOfInterestShapeTF{
			Type:                helper.TFStringValue(x.Shape.Type),
			Geolocation:         geoPointToTF(x.Shape.Geolocation),
			CenterGeolocation:   geoPointToTF(x.Shape.CenterGeolocation),
			RadiusInMeters:      helper.TFFloat64PtrValue(x.Shape.RadiusInMeters),
			VerticesGeolocation: verticesGeolocation,
		}
	}
	return &AreaOfInterestTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Shape:        shape,
		Tags:         general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *AreaOfInterestTF) ToAPI() any {
	var shape *AreaOfInterestShape
	if tf.Shape != nil {
		verticesGeolocation := make([]GeoPoint, len(tf.Shape.VerticesGeolocation))
		for i, v := range tf.Shape.VerticesGeolocation {
			verticesGeolocation[i] = GeoPoint{
				Latitude:  helper.FromTFFloat64(v.Latitude),
				Longitude: helper.FromTFFloat64(v.Longitude),
				Altitude:  helper.FromTFFloat64(v.Altitude),
			}
		}
		shape = &AreaOfInterestShape{
			Type:                helper.FromTFString(tf.Shape.Type),
			Geolocation:         geoPointFromTF(tf.Shape.Geolocation),
			CenterGeolocation:   geoPointFromTF(tf.Shape.CenterGeolocation),
			RadiusInMeters:      helper.FromTFFloat64Ptr(tf.Shape.RadiusInMeters),
			VerticesGeolocation: verticesGeolocation,
		}
	}
	return &AreaOfInterest{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		Shape:      shape,
		Tags:       general_objects.KeyValuesFromTF(tf.Tags),
	}
}
