package sensors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (sensor *Sensor) ToMap() map[string]any {
	sensorMap := sensor.ToAuditMap()
	sensorMap["name"] = helper.NilIfEmpty(sensor.Name)
	sensorMap["satellite_id"] = helper.NilIfEmpty(sensor.SatelliteID)
	if sensor.ApertureShape != nil {
		sensorMap["aperture_shape"] = sensor.ApertureShape.ToMap()
	}
	sensorMap["tags"] = helper.ParseToMaps(sensor.Tags)
	return sensorMap
}

func (shape *ApertureShape) ToMap() map[string]any {
	shapeMap := make(map[string]any)

	shapeMap["type"] = helper.NilIfEmpty(shape.Type)

	if shape.ApertureCenter != nil {
		shapeMap["aperture_center"] = shape.ApertureCenter.ToMap()
	}

	if shape.Type == "CIRCULAR" {
		if shape.HalfApertureAngle != nil {
			shapeMap["half_aperture_angle"] = shape.HalfApertureAngle.ToMap()
		}
	}

	if shape.Type == "RECTANGULAR" {
		if shape.FirstAxisVector != nil {
			shapeMap["first_axis_vector"] = shape.FirstAxisVector.ToMap()
		}
		if shape.FirstAxisHalfApertureAngle != nil {
			shapeMap["first_axis_half_aperture_angle"] = shape.FirstAxisHalfApertureAngle.ToMap()
		}
		if shape.SecondAxisVector != nil {
			shapeMap["second_axis_vector"] = shape.SecondAxisVector.ToMap()
		}
		if shape.SecondAxisHalfApertureAngle != nil {
			shapeMap["second_axis_half_aperture_angle"] = shape.SecondAxisHalfApertureAngle.ToMap()
		}
	}

	return shapeMap
}

func (vector *Vector3D) ToMap() map[string]any {
	vector3DMap := make(map[string]any)
	vector3DMap["x"] = helper.NilIfEmpty(vector.X)
	vector3DMap["y"] = helper.NilIfEmpty(vector.Y)
	vector3DMap["z"] = helper.NilIfEmpty(vector.Z)
	return vector3DMap
}

func (halfApertureAngle *HalfApertureAngle) ToMap() map[string]any {
	hapfApertureAngleMap := make(map[string]any)
	hapfApertureAngleMap["degrees"] = helper.NilIfEmpty(halfApertureAngle.Degrees)
	return hapfApertureAngleMap
}

func (sensor *Sensor) FromMap(sensorMap map[string]any) error {
	sensor.FromAuditMap(sensorMap)
	sensor.SatelliteID = helper.CastString(sensorMap, "satellite_id")
	sensor.Name = helper.CastString(sensorMap, "name")

	sensor.ApertureShape = new(ApertureShape)
	if err := sensor.ApertureShape.FromMap(helper.CastMapAny(sensorMap, "aperture_shape")); err != nil {
		return err
	}

	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(sensorMap, "tags")); err != nil {
		return err
	} else {
		sensor.Tags = tags
	}
	return nil
}

func (shape *ApertureShape) FromMap(shapeMap map[string]any) error {

	shape.Type = helper.CastString(shapeMap, "type")

	if shapeMap["aperture_center"] != nil {
		shape.ApertureCenter = new(Vector3D)
		if err := shape.ApertureCenter.FromMap(helper.CastMapAny(shapeMap, "aperture_center")); err != nil {
			return err
		}
	}

	if shape.Type == "CIRCULAR" {
		shape.HalfApertureAngle = new(HalfApertureAngle)
		if err := shape.HalfApertureAngle.FromMap(helper.CastMapAny(shapeMap, "half_aperture_angle")); err != nil {
			return err
		}
	}

	if shape.Type == "RECTANGULAR" {

		shape.FirstAxisVector = new(Vector3D)
		if err := shape.FirstAxisVector.FromMap(helper.CastMapAny(shapeMap, "first_axis_vector")); err != nil {
			return err
		}

		shape.FirstAxisHalfApertureAngle = new(HalfApertureAngle)
		if err := shape.FirstAxisHalfApertureAngle.FromMap(helper.CastMapAny(shapeMap, "first_axis_half_aperture_angle")); err != nil {
			return err
		}

		shape.SecondAxisVector = new(Vector3D)
		if err := shape.SecondAxisVector.FromMap(helper.CastMapAny(shapeMap, "second_axis_vector")); err != nil {
			return err
		}

		shape.SecondAxisHalfApertureAngle = new(HalfApertureAngle)
		if err := shape.SecondAxisHalfApertureAngle.FromMap(helper.CastMapAny(shapeMap, "second_axis_half_aperture_angle")); err != nil {
			return err
		}
	}

	return nil
}

func (vector *Vector3D) FromMap(vectorMap map[string]any) error {
	vector.X = helper.CastFloat64(vectorMap, "x")
	vector.Y = helper.CastFloat64(vectorMap, "y")
	vector.Z = helper.CastFloat64(vectorMap, "z")
	return nil
}

func (halfApertureAngle *HalfApertureAngle) FromMap(halfApertureAngleMap map[string]any) error {
	halfApertureAngle.Degrees = helper.CastFloat64(halfApertureAngleMap, "degrees")
	return nil
}
