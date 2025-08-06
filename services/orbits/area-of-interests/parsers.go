package areas_of_interest

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (aoi *AreaOfInterest) ToMap() map[string]any {
	aoiMap := make(map[string]any)
	aoiMap["id"] = aoi.ID
	aoiMap["name"] = aoi.Name
	if aoi.Shape != nil {
		aoiMap["shape"] = []map[string]any{aoi.Shape.ToMap()}
	}
	aoiMap["tags"] = helper.ParseToMaps(aoi.Tags)
	aoiMap["created_at"] = aoi.CreatedAt
	aoiMap["created_by"] = aoi.CreatedBy
	aoiMap["last_modified_at"] = aoi.LastModifiedAt
	aoiMap["last_modified_by"] = aoi.LastModifiedBy
	return aoiMap
}

func (shape *AreaOfInterestShape) ToMap() map[string]any {
	shapeMap := make(map[string]any)

	if shape.Type == "POINT" {
		if shape.Geolocation != nil {
			shapeMap["geolocation"] = []map[string]any{shape.Geolocation.ToMap()}
		}
	}

	if shape.Type == "CIRCLE" {
		if shape.CenterGeolocation != nil {
			shapeMap["center_geolocation"] = []map[string]any{shape.CenterGeolocation.ToMap()}
		}
		shapeMap["radius_in_meters"] = shape.RadiusInMeters
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
	geopointMap["latitude"] = geopoint.Latitude
	geopointMap["longitude"] = geopoint.Longitude
	geopointMap["altitude"] = geopoint.Altitude
	return geopointMap
}

func (aoi *AreaOfInterest) FromMap(aoiMap map[string]any) error {
	aoi.ID = aoiMap["id"].(string)
	aoi.Name = aoiMap["name"].(string)

	if aoi.Shape != nil {
		aoiMap["shape"] = []map[string]any{aoi.Shape.ToMap()}
	}

	if len(aoiMap["shape"].([]any)) > 0 && aoiMap["shape"].([]any)[0] != nil {
		aoi.Shape = new(AreaOfInterestShape)
		if err := aoi.Shape.FromMap(aoiMap["shape"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}

	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](aoiMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		aoi.Tags = tags
	}

	aoi.CreatedAt = aoiMap["created_at"].(string)
	aoi.CreatedBy = aoiMap["created_by"].(string)
	aoi.LastModifiedAt = aoiMap["last_modified_at"].(string)
	aoi.LastModifiedBy = aoiMap["last_modified_by"].(string)

	return nil
}

func (shape *AreaOfInterestShape) FromMap(shapeMap map[string]any) error {

	shape.Type = shapeMap["type"].(string)

	if shape.Type == "POINT" {
		if len(shapeMap["geolocation"].([]any)) > 0 && shapeMap["geolocation"].([]any)[0] != nil {
			shape.Geolocation = new(GeoPoint)
			if err := shape.Geolocation.FromMap(shapeMap["geolocation"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}
	}

	if shape.Type == "CIRCLE" {
		if len(shapeMap["center_geolocation"].([]any)) > 0 && shapeMap["center_geolocation"].([]any)[0] != nil {
			shape.CenterGeolocation = new(GeoPoint)
			if err := shape.CenterGeolocation.FromMap(shapeMap["center_geolocation"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
			shape.RadiusInMeters = shapeMap["radius_in_meters"].(float64)
		}
	}

	if shape.Type == "POLYGON" {
		if shapeMap["vertices_geolocation"] != nil {
			if vertices_geolocation, err := helper.ParseFromMaps[GeoPoint](
				shapeMap["vertices_geolocation"].(*schema.Set).List(),
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
	geopoint.Latitude = geopointMap["latitude"].(float64)
	geopoint.Longitude = geopointMap["longitude"].(float64)
	geopoint.Altitude = geopointMap["altitude"].(float64)
	return nil
}
