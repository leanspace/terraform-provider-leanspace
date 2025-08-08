package sensors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (sensor *Sensor) ToMap() map[string]any {
	sensorMap := make(map[string]any)
	sensorMap["id"] = sensor.ID
	sensorMap["name"] = sensor.Name
	sensorMap["satellite_id"] = sensor.SatelliteID
	if sensor.ApertureShape != nil {
		sensorMap["aperture_shape"] = []map[string]any{sensor.ApertureShape.ToMap()}
	}
	sensorMap["tags"] = helper.ParseToMaps(sensor.Tags)
	sensorMap["created_at"] = sensor.CreatedAt
	sensorMap["created_by"] = sensor.CreatedBy
	sensorMap["last_modified_at"] = sensor.LastModifiedAt
	sensorMap["last_modified_by"] = sensor.LastModifiedBy
	return sensorMap
}

func (shape *ApertureShape) ToMap() map[string]any {
	shapeMap := make(map[string]any)

	shapeMap["type"] = shape.Type

	if shape.ApertureCenter != nil {
		shapeMap["aperture_center"] = []map[string]any{shape.ApertureCenter.ToMap()}
	}

	if shape.Type == "CIRCULAR" {
		if shape.HalfApertureAngle != nil {
			shapeMap["half_aperture_angle"] = []map[string]any{shape.HalfApertureAngle.ToMap()}
		}
	}

	if shape.Type == "RECTANGULAR" {
		if shape.FirstAxisVector != nil {
			shapeMap["first_axis_vector"] = []map[string]any{shape.FirstAxisVector.ToMap()}
		}
		if shape.FirstAxisHalfApertureAngle != nil {
			shapeMap["first_axis_half_aperture_angle"] = []map[string]any{shape.FirstAxisHalfApertureAngle.ToMap()}
		}
		if shape.SecondAxisVector != nil {
			shapeMap["second_axis_vector"] = []map[string]any{shape.SecondAxisVector.ToMap()}
		}
		if shape.SecondAxisHalfApertureAngle != nil {
			shapeMap["second_axis_half_aperture_angle"] = []map[string]any{shape.SecondAxisHalfApertureAngle.ToMap()}
		}
	}

	return shapeMap
}

func (vector *Vector3D) ToMap() map[string]any {
	vector3DMap := make(map[string]any)
	vector3DMap["x"] = vector.X
	vector3DMap["y"] = vector.Y
	vector3DMap["z"] = vector.Z
	return vector3DMap
}

func (halfApertureAngle *HalfApertureAngle) ToMap() map[string]any {
	hapfApertureAngleMap := make(map[string]any)
	hapfApertureAngleMap["degrees"] = halfApertureAngle.Degrees
	return hapfApertureAngleMap
}

func (sensor *Sensor) FromMap(sensorMap map[string]any) error {
	sensor.ID = sensorMap["id"].(string)
	sensor.SatelliteID = sensorMap["satellite_id"].(string)
	sensor.Name = sensorMap["name"].(string)

	if len(sensorMap["aperture_shape"].([]any)) > 0 && sensorMap["aperture_shape"].([]any)[0] != nil {
		sensor.ApertureShape = new(ApertureShape)
		if err := sensor.ApertureShape.FromMap(sensorMap["aperture_shape"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}

	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](sensorMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		sensor.Tags = tags
	}

	sensor.CreatedAt = sensorMap["created_at"].(string)
	sensor.CreatedBy = sensorMap["created_by"].(string)
	sensor.LastModifiedAt = sensorMap["last_modified_at"].(string)
	sensor.LastModifiedBy = sensorMap["last_modified_by"].(string)

	return nil
}

func (shape *ApertureShape) FromMap(shapeMap map[string]any) error {

	shape.Type = shapeMap["type"].(string)

	if shape.Type == "CIRCULAR" {
		if len(shapeMap["half_aperture_angle"].([]any)) > 0 && shapeMap["half_aperture_angle"].([]any)[0] != nil {
			shape.HalfApertureAngle = new(HalfApertureAngle)
			if err := shape.HalfApertureAngle.FromMap(shapeMap["half_aperture_angle"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}
	}

	if shape.Type == "RECTANGULAR" {

		if len(shapeMap["first_axis_vector"].([]any)) > 0 && shapeMap["first_axis_vector"].([]any)[0] != nil {
			shape.FirstAxisVector = new(Vector3D)
			if err := shape.FirstAxisVector.FromMap(shapeMap["first_axis_vector"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}

		if len(shapeMap["first_axis_half_aperture_angle"].([]any)) > 0 && shapeMap["first_axis_half_aperture_angle"].([]any)[0] != nil {
			shape.FirstAxisHalfApertureAngle = new(HalfApertureAngle)
			if err := shape.FirstAxisHalfApertureAngle.FromMap(shapeMap["first_axis_half_aperture_angle"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}

		if len(shapeMap["second_axis_vector"].([]any)) > 0 && shapeMap["second_axis_vector"].([]any)[0] != nil {
			shape.SecondAxisVector = new(Vector3D)
			if err := shape.SecondAxisVector.FromMap(shapeMap["second_axis_vector"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}

		if len(shapeMap["second_axis_half_aperture_angle"].([]any)) > 0 && shapeMap["second_axis_half_aperture_angle"].([]any)[0] != nil {
			shape.SecondAxisHalfApertureAngle = new(HalfApertureAngle)
			if err := shape.SecondAxisHalfApertureAngle.FromMap(shapeMap["second_axis_half_aperture_angle"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (vector *Vector3D) FromMap(vectorMap map[string]any) error {
	vector.X = vectorMap["x"].(float64)
	vector.Y = vectorMap["y"].(float64)
	vector.Z = vectorMap["z"].(float64)
	return nil
}

func (halfApertureAngle *HalfApertureAngle) FromMap(halfApertureAngleMap map[string]any) error {
	halfApertureAngle.Degrees = halfApertureAngleMap["degrees"].(float64)
	return nil
}
