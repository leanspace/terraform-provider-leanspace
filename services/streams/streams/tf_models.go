package streams

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type StreamTF struct {
	general_objects.AuditModelTF
	Version       types.Int64                  `tfsdk:"version"`
	Name          types.String                 `tfsdk:"name"`
	Description   types.String                 `tfsdk:"description"`
	Tags          []general_objects.KeyValueTF `tfsdk:"tags"`
	AssetId       types.String                 `tfsdk:"asset_id"`
	Configuration *ConfigurationTF             `tfsdk:"configuration"`
	Mappings      []MappingTF                  `tfsdk:"mappings"`
}

type ConfigurationTF struct {
	Endianness   types.String            `tfsdk:"endianness"`
	Structure    *ElementListTF          `tfsdk:"structure"`
	Metadata     *MetadataTF             `tfsdk:"metadata"`
	Computations *ElementListWithValidTF `tfsdk:"computations"`
}

type ElementListTF struct {
	Elements []StreamComponentTF `tfsdk:"elements"`
}

type ElementListWithValidTF struct {
	Elements []ComputationTF `tfsdk:"elements"`
	Valid    types.Bool      `tfsdk:"valid"`
}

type StreamComponentTF struct {
	Name       types.String        `tfsdk:"name"`
	Order      types.Int64         `tfsdk:"order"`
	Path       types.String        `tfsdk:"path"`
	Type       types.String        `tfsdk:"type"`
	Repetitive *RepetitiveTF       `tfsdk:"repetitive"`
	Length     *LengthTF           `tfsdk:"length"`
	Processor  types.String        `tfsdk:"processor"`
	DataType   types.String        `tfsdk:"data_type"`
	Endianness types.String        `tfsdk:"endianness"`
	Expression *SwitchExpressionTF `tfsdk:"expression"`
	Elements   []StreamComponentTF `tfsdk:"elements"`
}

type RepetitiveTF struct {
	Value types.Int64  `tfsdk:"value"`
	Path  types.String `tfsdk:"path"`
}

type LengthTF struct {
	Type  types.String `tfsdk:"type"`
	Unit  types.String `tfsdk:"unit"`
	Value types.Int64  `tfsdk:"value"`
	Path  types.String `tfsdk:"path"`
}

type SwitchExpressionTF struct {
	SwitchOn types.String     `tfsdk:"switch_on"`
	Options  []SwitchOptionTF `tfsdk:"options"`
}

type SwitchOptionTF struct {
	Component types.String   `tfsdk:"component"`
	Value     *SwitchValueTF `tfsdk:"value"`
}

type SwitchValueTF struct {
	DataType types.String `tfsdk:"data_type"`
	Data     types.String `tfsdk:"data"`
}

type ComputationTF struct {
	Name       types.String `tfsdk:"name"`
	Order      types.Int64  `tfsdk:"order"`
	Type       types.String `tfsdk:"type"`
	DataType   types.String `tfsdk:"data_type"`
	Expression types.String `tfsdk:"expression"`
}

type MetadataTF struct {
	Timestamp *TimestampDefinitionTF `tfsdk:"timestamp"`
}

type TimestampDefinitionTF struct {
	Expression types.String `tfsdk:"expression"`
}

type MappingTF struct {
	MetricId   types.String `tfsdk:"metric_id"`
	Expression types.String `tfsdk:"expression"`
}

// --- ToTF conversions ---

func streamComponentToTF(sc StreamComponent) StreamComponentTF {
	tf := StreamComponentTF{
		Name:       helper.TFStringValue(sc.Name),
		Order:      helper.TFInt64Value(sc.Order),
		Path:       helper.TFStringValue(sc.Path),
		Type:       helper.TFStringValue(sc.Type),
		Processor:  helper.TFStringValue(sc.Processor),
		DataType:   helper.TFStringValue(sc.DataType),
		Endianness: helper.TFStringValue(sc.Endianness),
	}

	if sc.Repetitive != nil {
		tf.Repetitive = &RepetitiveTF{
			Value: helper.TFInt64Value(sc.Repetitive.Value),
			Path:  helper.TFStringValue(sc.Repetitive.Path),
		}
	}

	if sc.Length != nil {
		tf.Length = &LengthTF{
			Type:  helper.TFStringValue(sc.Length.Type),
			Unit:  helper.TFStringValue(sc.Length.Unit),
			Value: helper.TFInt64Value(sc.Length.Value),
			Path:  helper.TFStringValue(sc.Length.Path),
		}
	}

	if sc.Expression != nil {
		opts := make([]SwitchOptionTF, len(sc.Expression.Options))
		for i, o := range sc.Expression.Options {
			opts[i] = SwitchOptionTF{
				Component: helper.TFStringValue(o.Component),
				Value: &SwitchValueTF{
					DataType: helper.TFStringValue(o.Value.DataType),
					Data:     helper.TFStringValue(fmt.Sprint(o.Value.Data)),
				},
			}
		}
		tf.Expression = &SwitchExpressionTF{
			SwitchOn: helper.TFStringValue(sc.Expression.SwitchOn),
			Options:  opts,
		}
	}

	if sc.Elements != nil {
		tf.Elements = make([]StreamComponentTF, len(sc.Elements))
		for i, e := range sc.Elements {
			tf.Elements[i] = streamComponentToTF(e)
		}
	}

	return tf
}

func streamComponentFromTF(tf StreamComponentTF) StreamComponent {
	sc := StreamComponent{
		Name:       helper.FromTFString(tf.Name),
		Order:      helper.FromTFInt64(tf.Order),
		Path:       helper.FromTFString(tf.Path),
		Type:       helper.FromTFString(tf.Type),
		Processor:  helper.FromTFString(tf.Processor),
		DataType:   helper.FromTFString(tf.DataType),
		Endianness: helper.FromTFString(tf.Endianness),
	}

	if tf.Repetitive != nil {
		sc.Repetitive = &Repetitive{
			Value: helper.FromTFInt64(tf.Repetitive.Value),
			Path:  helper.FromTFString(tf.Repetitive.Path),
		}
	}

	if tf.Length != nil {
		sc.Length = &Length{
			Type:  helper.FromTFString(tf.Length.Type),
			Unit:  helper.FromTFString(tf.Length.Unit),
			Value: helper.FromTFInt64(tf.Length.Value),
			Path:  helper.FromTFString(tf.Length.Path),
		}
	}

	if tf.Expression != nil {
		opts := make([]SwitchOption, len(tf.Expression.Options))
		for i, o := range tf.Expression.Options {
			opts[i] = SwitchOption{
				Component: helper.FromTFString(o.Component),
			}
			if o.Value != nil {
				opts[i].Value = SwitchValue[any]{
					DataType: helper.FromTFString(o.Value.DataType),
					Data:     helper.FromTFString(o.Value.Data),
				}
			}
		}
		sc.Expression = &SwitchExpression{
			SwitchOn: helper.FromTFString(tf.Expression.SwitchOn),
			Options:  opts,
		}
	}

	if tf.Elements != nil {
		sc.Elements = make([]StreamComponent, len(tf.Elements))
		for i, e := range tf.Elements {
			sc.Elements[i] = streamComponentFromTF(e)
		}
	}

	return sc
}

func (x *Stream) ToTF() interface{} {
	// Structure elements
	structElements := make([]StreamComponentTF, len(x.Configuration.Structure.Elements))
	for i, e := range x.Configuration.Structure.Elements {
		structElements[i] = streamComponentToTF(e)
	}

	// Computations
	compElements := make([]ComputationTF, len(x.Configuration.Computations.Elements))
	for i, c := range x.Configuration.Computations.Elements {
		compElements[i] = ComputationTF{
			Name:       helper.TFStringValue(c.Name),
			Order:      helper.TFInt64Value(c.Order),
			Type:       helper.TFStringValue(c.Type),
			DataType:   helper.TFStringValue(c.DataType),
			Expression: helper.TFStringValue(c.Expression),
		}
	}

	// Mappings
	mappings := make([]MappingTF, len(x.Mappings))
	for i, m := range x.Mappings {
		mappings[i] = MappingTF{
			MetricId:   helper.TFStringValue(m.MetricId),
			Expression: helper.TFStringValue(m.Expression),
		}
	}

	return &StreamTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Version:      helper.TFInt64Value(x.Version),
		Name:         helper.TFStringValue(x.Name),
		Description:  helper.TFStringValue(x.Description),
		Tags:         general_objects.KeyValuesToTF(x.Tags),
		AssetId:      helper.TFStringValue(x.AssetId),
		Configuration: &ConfigurationTF{
			Endianness: helper.TFStringValue(x.Configuration.Endianness),
			Structure: &ElementListTF{
				Elements: structElements,
			},
			Metadata: &MetadataTF{
				Timestamp: &TimestampDefinitionTF{
					Expression: helper.TFStringValue(x.Configuration.Metadata.Timestamp.Expression),
				},
			},
			Computations: &ElementListWithValidTF{
				Elements: compElements,
				Valid:    helper.TFBoolValue(x.Configuration.Computations.Valid),
			},
		},
		Mappings: mappings,
	}
}

func (tf *StreamTF) ToAPI() interface{} {
	// Structure elements
	var structElements []StreamComponent
	if tf.Configuration != nil && tf.Configuration.Structure != nil {
		structElements = make([]StreamComponent, len(tf.Configuration.Structure.Elements))
		for i, e := range tf.Configuration.Structure.Elements {
			structElements[i] = streamComponentFromTF(e)
		}
	}

	// Computations
	var compElements []Computation
	var compValid bool
	if tf.Configuration != nil && tf.Configuration.Computations != nil {
		compElements = make([]Computation, len(tf.Configuration.Computations.Elements))
		for i, c := range tf.Configuration.Computations.Elements {
			compElements[i] = Computation{
				Name:       helper.FromTFString(c.Name),
				Order:      helper.FromTFInt64(c.Order),
				Type:       helper.FromTFString(c.Type),
				DataType:   helper.FromTFString(c.DataType),
				Expression: helper.FromTFString(c.Expression),
			}
		}
		compValid = helper.FromTFBool(tf.Configuration.Computations.Valid)
	}

	// Metadata
	var metadata Metadata
	if tf.Configuration != nil && tf.Configuration.Metadata != nil && tf.Configuration.Metadata.Timestamp != nil {
		metadata.Timestamp.Expression = helper.FromTFString(tf.Configuration.Metadata.Timestamp.Expression)
	}

	// Mappings
	mappings := make([]Mapping, len(tf.Mappings))
	for i, m := range tf.Mappings {
		mappings[i] = Mapping{
			MetricId:   helper.FromTFString(m.MetricId),
			Expression: helper.FromTFString(m.Expression),
		}
	}

	var config Configuration
	if tf.Configuration != nil {
		config = Configuration{
			Endianness: helper.FromTFString(tf.Configuration.Endianness),
			Structure: ElementList[StreamComponent, *StreamComponent]{
				Elements: structElements,
			},
			Metadata: metadata,
			Computations: ElementListWithValid[Computation, *Computation]{
				Elements: compElements,
				Valid:    compValid,
			},
		}
	}

	return &Stream{
		AuditModel:    general_objects.AuditModelFromTF(tf.AuditModelTF),
		Version:       helper.FromTFInt64(tf.Version),
		Name:          helper.FromTFString(tf.Name),
		Description:   helper.FromTFString(tf.Description),
		Tags:          general_objects.KeyValuesFromTF(tf.Tags),
		AssetId:       helper.FromTFString(tf.AssetId),
		Configuration: config,
		Mappings:      mappings,
	}
}
