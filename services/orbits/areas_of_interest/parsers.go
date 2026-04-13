package areas_of_interest

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (aoi *AreaOfInterest) ToMap() map[string]any {
	aoiMap := aoi.ToAuditMap()
	aoiMap["name"] = helper.NilIfEmpty(aoi.Name)
	if aoi.Shape != nil {
		aoiMap["shape"] = aoi.Shape.ToMap()
	}
	aoiMap["tags"] = helper.ParseToMaps(aoi.Tags)
	return aoiMap
}

func (shape *AreaOfInterestShape) ToMap() map[string]any {
	shapeMap := make(map[string]any)

	shapeMap["type"] = helper.NilIfEmpty(shape.Type)

	if shape.Type == "POINT" {
		if shape.Geolocation != nil {
			shapeMap["geolocation"] = shape.Geolocation.ToMap()
		}
	}

	if shape.Type == "CIRCLE" {
		if shape.CenterGeolocation != nil {
			shapeMap["center_geolocation"] = shape.CenterGeolocation.ToMap()
		}
		shapeMap["radius_in_meters"] = helper.NilIfEmpty(shape.RadiusInMeters)
	}

	if shape.Type == "POLYGON" {
		if shape.VerticesGeolocation != nil {
			shapeMap["vertices_geolocation"] = helper.ParseToMaps(shape.VerticesGeolocation)
		}
	}
	return shapeMap
}

func (geopoint *GeoPoint) ToMap() map[string]any {
	geopointMap := make(map[string]any)
	geopointMap["latitude"] = helper.NilIfEmpty(geopoint.Latitude)
	geopointMap["longitude"] = helper.NilIfEmpty(geopoint.Longitude)
	geopointMap["altitude"] = helper.NilIfEmpty(geopoint.Altitude)
	return geopointMap
}

func (aoi *AreaOfInterest) FromMap(aoiMap map[string]any) error {
	aoi.FromAuditMap(aoiMap)
	aoi.Name = helper.CastString(aoiMap, "name")

	aoi.Shape = new(AreaOfInterestShape)
	if err := aoi.Shape.FromMap(helper.CastMapAny(aoiMap, "shape")); err != nil {
		return err
	}

	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(aoiMap, "tags")); err != nil {
		return err
	} else {
		aoi.Tags = tags
	}
	return nil
}

func (shape *AreaOfInterestShape) FromMap(shapeMap map[string]any) error {

	shape.Type = helper.CastString(shapeMap, "type")

	if shape.Type == "POINT" {
		shape.Geolocation = new(GeoPoint)
		if err := shape.Geolocation.FromMap(helper.CastMapAny(shapeMap, "geolocation")); err != nil {
			return err
		}
	}

	if shape.Type == "CIRCLE" {
		shape.CenterGeolocation = new(GeoPoint)
		if err := shape.CenterGeolocation.FromMap(helper.CastMapAny(shapeMap, "center_geolocation")); err != nil {
			return err
		}
		shape.RadiusInMeters = helper.CastFloat64(shapeMap, "radius_in_meters")
	}

	if shape.Type == "POLYGON" {
		if shapeMap["vertices_geolocation"] != nil {
			if vertices_geolocation, err := helper.ParseFromMaps[GeoPoint](
				helper.CastSlice(shapeMap, "vertices_geolocation"),
			); err != nil {
				return err
			} else {
				shape.VerticesGeolocation = vertices_geolocation
			}
		}
	}

	return nil
}

func (geopoint *GeoPoint) FromMap(geopointMap map[string]any) error {
	geopoint.Latitude = helper.CastFloat64(geopointMap, "latitude")
	geopoint.Longitude = helper.CastFloat64(geopointMap, "longitude")
	geopoint.Altitude = helper.CastFloat64(geopointMap, "altitude")
	return nil
}
