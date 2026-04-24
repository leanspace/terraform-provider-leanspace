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
		Name:       types.StringValue(sc.Name),
		Order:      helper.TFInt64Value(sc.Order),
		Path:       types.StringValue(sc.Path),
		Type:       types.StringValue(sc.Type),
		Processor:  types.StringPointerValue(sc.Processor),
		DataType:   types.StringPointerValue(sc.DataType),
		Endianness: types.StringPointerValue(sc.Endianness),
	}

	if sc.Repetitive != nil {
		tf.Repetitive = &RepetitiveTF{
			Value: helper.TFIntPtrValue(sc.Repetitive.Value),
			Path:  types.StringPointerValue(sc.Repetitive.Path),
		}
	}

	if sc.Length != nil {
		tf.Length = &LengthTF{
			Type:  types.StringValue(sc.Length.Type),
			Unit:  types.StringValue(sc.Length.Unit),
			Value: helper.TFIntPtrValue(sc.Length.Value),
			Path:  types.StringPointerValue(sc.Length.Path),
		}
	}

	if sc.Expression != nil {
		opts := make([]SwitchOptionTF, len(sc.Expression.Options))
		for i, o := range sc.Expression.Options {
			opts[i] = SwitchOptionTF{
				Component: types.StringValue(o.Component),
				Value: &SwitchValueTF{
					DataType: types.StringValue(o.Value.DataType),
					Data:     types.StringValue(fmt.Sprint(o.Value.Data)),
				},
			}
		}
		tf.Expression = &SwitchExpressionTF{
			SwitchOn: types.StringValue(sc.Expression.SwitchOn),
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
		Name:       tf.Name.ValueString(),
		Order:      helper.FromTFInt64(tf.Order),
		Path:       tf.Path.ValueString(),
		Type:       tf.Type.ValueString(),
		Processor:  tf.Processor.ValueStringPointer(),
		DataType:   tf.DataType.ValueStringPointer(),
		Endianness: tf.Endianness.ValueStringPointer(),
	}

	if tf.Repetitive != nil {
		sc.Repetitive = &Repetitive{
			Value: helper.FromTFIntPtr(tf.Repetitive.Value),
			Path:  tf.Repetitive.Path.ValueStringPointer(),
		}
	}

	if tf.Length != nil {
		sc.Length = &Length{
			Type:  tf.Length.Type.ValueString(),
			Unit:  tf.Length.Unit.ValueString(),
			Value: helper.FromTFIntPtr(tf.Length.Value),
			Path:  tf.Length.Path.ValueStringPointer(),
		}
	}

	if tf.Expression != nil {
		opts := make([]SwitchOption, len(tf.Expression.Options))
		for i, o := range tf.Expression.Options {
			opts[i] = SwitchOption{
				Component: o.Component.ValueString(),
			}
			if o.Value != nil {
				opts[i].Value = SwitchValue[any]{
					DataType: o.Value.DataType.ValueString(),
					Data:     o.Value.Data.ValueString(),
				}
			}
		}
		sc.Expression = &SwitchExpression{
			SwitchOn: tf.Expression.SwitchOn.ValueString(),
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
			Name:       types.StringValue(c.Name),
			Order:      helper.TFInt64Value(c.Order),
			Type:       types.StringValue(c.Type),
			DataType:   types.StringValue(c.DataType),
			Expression: types.StringValue(c.Expression),
		}
	}

	// Mappings
	mappings := make([]MappingTF, len(x.Mappings))
	for i, m := range x.Mappings {
		mappings[i] = MappingTF{
			MetricId:   types.StringValue(m.MetricId),
			Expression: types.StringPointerValue(m.Expression),
		}
	}

	return &StreamTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Version:      helper.TFInt64Value(x.Version),
		Name:         types.StringValue(x.Name),
		Description:  types.StringPointerValue(x.Description),
		Tags:         general_objects.KeyValuesToTF(x.Tags),
		AssetId:      types.StringValue(x.AssetId),
		Configuration: &ConfigurationTF{
			Endianness: types.StringValue(x.Configuration.Endianness),
			Structure: &ElementListTF{
				Elements: structElements,
			},
			Metadata: &MetadataTF{
				Timestamp: &TimestampDefinitionTF{
					Expression: types.StringValue(x.Configuration.Metadata.Timestamp.Expression),
				},
			},
			Computations: &ElementListWithValidTF{
				Elements: compElements,
				Valid:    types.BoolValue(x.Configuration.Computations.Valid),
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
				Name:       c.Name.ValueString(),
				Order:      helper.FromTFInt64(c.Order),
				Type:       c.Type.ValueString(),
				DataType:   c.DataType.ValueString(),
				Expression: c.Expression.ValueString(),
			}
		}
		compValid = tf.Configuration.Computations.Valid.ValueBool()
	}

	// Metadata
	var metadata Metadata
	if tf.Configuration != nil && tf.Configuration.Metadata != nil && tf.Configuration.Metadata.Timestamp != nil {
		metadata.Timestamp.Expression = tf.Configuration.Metadata.Timestamp.Expression.ValueString()
	}

	// Mappings
	mappings := make([]Mapping, len(tf.Mappings))
	for i, m := range tf.Mappings {
		mappings[i] = Mapping{
			MetricId:   m.MetricId.ValueString(),
			Expression: m.Expression.ValueStringPointer(),
		}
	}

	var config Configuration
	if tf.Configuration != nil {
		config = Configuration{
			Endianness: tf.Configuration.Endianness.ValueString(),
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
		Name:          tf.Name.ValueString(),
		Description:   tf.Description.ValueStringPointer(),
		Tags:          general_objects.KeyValuesFromTF(tf.Tags),
		AssetId:       tf.AssetId.ValueString(),
		Configuration: config,
		Mappings:      mappings,
	}
}
