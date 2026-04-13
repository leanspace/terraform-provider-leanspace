package sensors

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type SensorTF struct {
	general_objects.AuditModelTF
	SatelliteID   types.String                 `tfsdk:"satellite_id"`
	Name          types.String                 `tfsdk:"name"`
	ApertureShape *ApertureShapeTF             `tfsdk:"aperture_shape"`
	Tags          []general_objects.KeyValueTF `tfsdk:"tags"`
}

type ApertureShapeTF struct {
	Type                        types.String         `tfsdk:"type"`
	ApertureCenter              *Vector3DTF          `tfsdk:"aperture_center"`
	HalfApertureAngle           *HalfApertureAngleTF `tfsdk:"half_aperture_angle"`
	FirstAxisVector             *Vector3DTF          `tfsdk:"first_axis_vector"`
	FirstAxisHalfApertureAngle  *HalfApertureAngleTF `tfsdk:"first_axis_half_aperture_angle"`
	SecondAxisVector            *Vector3DTF          `tfsdk:"second_axis_vector"`
	SecondAxisHalfApertureAngle *HalfApertureAngleTF `tfsdk:"second_axis_half_aperture_angle"`
}

type Vector3DTF struct {
	X types.Float64 `tfsdk:"x"`
	Y types.Float64 `tfsdk:"y"`
	Z types.Float64 `tfsdk:"z"`
}

type HalfApertureAngleTF struct {
	Degrees types.Float64 `tfsdk:"degrees"`
}

func vector3DToTF(v *Vector3D) *Vector3DTF {
	if v == nil {
		return nil
	}
	return &Vector3DTF{
		X: helper.TFFloat64Value(v.X),
		Y: helper.TFFloat64Value(v.Y),
		Z: helper.TFFloat64Value(v.Z),
	}
}

func vector3DFromTF(tf *Vector3DTF) *Vector3D {
	if tf == nil {
		return nil
	}
	return &Vector3D{
		X: helper.FromTFFloat64(tf.X),
		Y: helper.FromTFFloat64(tf.Y),
		Z: helper.FromTFFloat64(tf.Z),
	}
}

func halfApertureAngleToTF(h *HalfApertureAngle) *HalfApertureAngleTF {
	if h == nil {
		return nil
	}
	return &HalfApertureAngleTF{
		Degrees: helper.TFFloat64Value(h.Degrees),
	}
}

func halfApertureAngleFromTF(tf *HalfApertureAngleTF) *HalfApertureAngle {
	if tf == nil {
		return nil
	}
	return &HalfApertureAngle{
		Degrees: helper.FromTFFloat64(tf.Degrees),
	}
}

func (x *Sensor) ToTF() any {
	var apertureShape *ApertureShapeTF
	if x.ApertureShape != nil {
		apertureShape = &ApertureShapeTF{
			Type:                        helper.TFStringValue(x.ApertureShape.Type),
			ApertureCenter:              vector3DToTF(x.ApertureShape.ApertureCenter),
			HalfApertureAngle:           halfApertureAngleToTF(x.ApertureShape.HalfApertureAngle),
			FirstAxisVector:             vector3DToTF(x.ApertureShape.FirstAxisVector),
			FirstAxisHalfApertureAngle:  halfApertureAngleToTF(x.ApertureShape.FirstAxisHalfApertureAngle),
			SecondAxisVector:            vector3DToTF(x.ApertureShape.SecondAxisVector),
			SecondAxisHalfApertureAngle: halfApertureAngleToTF(x.ApertureShape.SecondAxisHalfApertureAngle),
		}
	}
	return &SensorTF{
		AuditModelTF:  general_objects.AuditModelToTF(&x.AuditModel),
		SatelliteID:   helper.TFStringValue(x.SatelliteID),
		Name:          helper.TFStringValue(x.Name),
		ApertureShape: apertureShape,
		Tags:          general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *SensorTF) ToAPI() any {
	var apertureShape *ApertureShape
	if tf.ApertureShape != nil {
		apertureShape = &ApertureShape{
			Type:                        helper.FromTFString(tf.ApertureShape.Type),
			ApertureCenter:              vector3DFromTF(tf.ApertureShape.ApertureCenter),
			HalfApertureAngle:           halfApertureAngleFromTF(tf.ApertureShape.HalfApertureAngle),
			FirstAxisVector:             vector3DFromTF(tf.ApertureShape.FirstAxisVector),
			FirstAxisHalfApertureAngle:  halfApertureAngleFromTF(tf.ApertureShape.FirstAxisHalfApertureAngle),
			SecondAxisVector:            vector3DFromTF(tf.ApertureShape.SecondAxisVector),
			SecondAxisHalfApertureAngle: halfApertureAngleFromTF(tf.ApertureShape.SecondAxisHalfApertureAngle),
		}
	}
	return &Sensor{
		AuditModel:    general_objects.AuditModelFromTF(tf.AuditModelTF),
		SatelliteID:   helper.FromTFString(tf.SatelliteID),
		Name:          helper.FromTFString(tf.Name),
		ApertureShape: apertureShape,
		Tags:          general_objects.KeyValuesFromTF(tf.Tags),
	}
}
