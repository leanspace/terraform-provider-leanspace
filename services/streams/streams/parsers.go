package streams

import (
	"encoding/base64"
	"strconv"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (stream *Stream) ToMap() map[string]any {
	streamMap := stream.ToAuditMap()
	streamMap["version"] = helper.NilIfEmpty(stream.Version)
	streamMap["name"] = helper.NilIfEmpty(stream.Name)
	streamMap["description"] = helper.NilIfEmpty(stream.Description)
	streamMap["tags"] = helper.ParseToMaps(stream.Tags)
	streamMap["asset_id"] = helper.NilIfEmpty(stream.AssetId)
	streamMap["configuration"] = stream.Configuration.ToMap()
	streamMap["mappings"] = helper.ParseToMaps(stream.Mappings)
	return streamMap
}

func (configuration *Configuration) ToMap() map[string]any {
	configMap := make(map[string]any)
	configMap["endianness"] = helper.NilIfEmpty(configuration.Endianness)
	configMap["structure"] = configuration.Structure.ToMap()
	configMap["metadata"] = configuration.Metadata.ToMap()
	configMap["computations"] = configuration.Computations.ToMap()
	return configMap
}

func (streamComp *StreamComponent) ToMap() map[string]any {
	streamCompMap := make(map[string]any)
	streamCompMap["name"] = helper.NilIfEmpty(streamComp.Name)
	streamCompMap["order"] = helper.NilIfEmpty(streamComp.Order)
	streamCompMap["path"] = helper.NilIfEmpty(streamComp.Path)
	streamCompMap["type"] = helper.NilIfEmpty(streamComp.Type)

	if streamComp.Repetitive != nil {
		streamCompMap["repetitive"] = streamComp.Repetitive.ToMap()
	}

	if streamComp.Type == "FIELD" {
		streamCompMap["length"] = streamComp.Length.ToMap()
		streamCompMap["processor"] = helper.NilIfEmpty(streamComp.Processor)
		streamCompMap["data_type"] = helper.NilIfEmpty(streamComp.DataType)
		streamCompMap["endianness"] = helper.NilIfEmpty(streamComp.Endianness)
	}
	if streamComp.Type == "SWITCH" {
		streamCompMap["expression"] = streamComp.Expression.ToMap()
	}
	if streamComp.Type == "SWITCH" || streamComp.Type == "CONTAINER" {
		streamCompMap["elements"] = helper.ParseToMaps(streamComp.Elements)
	}

	return streamCompMap
}

func (repetitive *Repetitive) ToMap() map[string]any {
	repetitiveMap := make(map[string]any)
	if repetitive != nil && repetitive.Value != 0 {
		repetitiveMap["value"] = helper.NilIfEmpty(repetitive.Value)
	}
	if repetitive != nil && repetitive.Path != "" {
		repetitiveMap["path"] = helper.NilIfEmpty(repetitive.Path)
	}
	return repetitiveMap
}

func (length *Length) ToMap() map[string]any {
	lengthMap := make(map[string]any)
	lengthMap["type"] = helper.NilIfEmpty(length.Type)
	lengthMap["unit"] = helper.NilIfEmpty(length.Unit)
	lengthMap["value"] = helper.NilIfEmpty(length.Value)
	lengthMap["path"] = helper.NilIfEmpty(length.Path)
	return lengthMap
}

func (switchExp *SwitchExpression) ToMap() map[string]any {
	switchExpMap := make(map[string]any)
	switchExpMap["switch_on"] = helper.NilIfEmpty(switchExp.SwitchOn)
	switchExpMap["options"] = helper.ParseToMaps(switchExp.Options)
	return switchExpMap
}

func (switchOption *SwitchOption) ToMap() map[string]any {
	switchOptionMap := make(map[string]any)
	switchOptionMap["value"] = switchOption.Value.ToMap()
	switchOptionMap["component"] = helper.NilIfEmpty(switchOption.Component)
	return switchOptionMap
}

func (switchValue *SwitchValue[T]) ToMap() map[string]any {
	switchValueMap := make(map[string]any)
	switchValueMap["data_type"] = helper.NilIfEmpty(switchValue.DataType)
	switch switchValue.DataType {
	case "INTEGER", "UINTEGER", "DECIMAL":
		switchValueMap["data"] = helper.ParseFloat(any(switchValue.Data).(float64))
	case "TEXT", "BINARY", "TIMESTAMP", "DATE":
		switchValueMap["data"] = any(switchValue.Data).(string)
	case "BOOLEAN":
		switchValueMap["data"] = strconv.FormatBool(any(switchValue.Data).(bool))
	}
	return switchValueMap
}

func (metadata *Metadata) ToMap() map[string]any {
	metadataMap := make(map[string]any)
	metadataMap["timestamp"] = metadata.Timestamp.ToMap()
	return metadataMap
}

func (timestampDef *TimestampDefinition) ToMap() map[string]any {
	timestampDefMap := make(map[string]any)
	timestampDefMap["expression"] = helper.NilIfEmpty(timestampDef.Expression)
	return timestampDefMap
}

func (elementList *ElementList[T, PT]) ToMap() map[string]any {
	elementListMap := make(map[string]any)
	elementListMap["elements"] = helper.ParseToMaps[T, PT](elementList.Elements)
	return elementListMap
}

func (elementList *ElementListWithValid[T, PT]) ToMap() map[string]any {
	elementListMap := make(map[string]any)
	elementListMap["elements"] = helper.ParseToMaps[T, PT](elementList.Elements)
	elementListMap["valid"] = helper.NilIfEmpty(elementList.Valid)
	return elementListMap
}

func (computation *Computation) ToMap() map[string]any {
	computationMap := make(map[string]any)
	computationMap["name"] = helper.NilIfEmpty(computation.Name)
	computationMap["order"] = helper.NilIfEmpty(computation.Order)
	computationMap["type"] = helper.NilIfEmpty(computation.Type)
	computationMap["data_type"] = helper.NilIfEmpty(computation.DataType)
	computationMap["expression"] = helper.NilIfEmpty(computation.Expression)
	return computationMap
}

func (mapping *Mapping) ToMap() map[string]any {
	mappingMap := make(map[string]any)
	mappingMap["metric_id"] = helper.NilIfEmpty(mapping.MetricId)
	mappingMap["expression"] = helper.NilIfEmpty(mapping.Expression)
	return mappingMap
}

func (stream *Stream) FromMap(streamMap map[string]any) error {
	stream.FromAuditMap(streamMap)
	stream.Version = helper.CastInt(streamMap, "version")
	stream.Name = helper.CastString(streamMap, "name")
	stream.Description = helper.CastString(streamMap, "description")
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(streamMap, "tags")); err != nil {
		return err
	} else {
		stream.Tags = tags
	}
	stream.AssetId = helper.CastString(streamMap, "asset_id")
	if err := stream.Configuration.FromMap(helper.CastMapAny(streamMap, "configuration")); err != nil {
		return err
	}
	if mappings, err := helper.ParseFromMaps[Mapping](helper.CastSlice(streamMap, "mappings")); err != nil {
		return err
	} else {
		stream.Mappings = mappings
	}
	return nil
}

func (configuration *Configuration) FromMap(configMap map[string]any) error {
	configuration.Endianness = helper.CastString(configMap, "endianness")
	if err := configuration.Structure.FromMap(helper.CastMapAny(configMap, "structure")); err != nil {
		return err
	}
	if err := configuration.Metadata.FromMap(helper.CastMapAny(configMap, "metadata")); err != nil {
		return err
	}
	if err := configuration.Computations.FromMap(helper.CastMapAny(configMap, "computations")); err != nil {
		return err
	}
	return nil
}

func (streamComp *StreamComponent) FromMap(streamCompMap map[string]any) error {
	streamComp.Name = helper.CastString(streamCompMap, "name")
	streamComp.Order = helper.CastInt(streamCompMap, "order")
	streamComp.Path = helper.CastString(streamCompMap, "path")
	streamComp.Type = helper.CastString(streamCompMap, "type")
	if streamCompMap["repetitive"] != nil {
		streamComp.Repetitive = &Repetitive{}
		if err := streamComp.Repetitive.FromMap(helper.CastMapAny(streamCompMap, "repetitive")); err != nil {
			return err
		}
	}

	if streamComp.Type == "FIELD" {
		streamComp.Length = new(Length)
		if err := streamComp.Length.FromMap(helper.CastMapAny(streamCompMap, "length")); err != nil {
			return err
		}
		streamComp.Processor = helper.CastString(streamCompMap, "processor")
		streamComp.DataType = helper.CastString(streamCompMap, "data_type")
		streamComp.Endianness = helper.CastString(streamCompMap, "endianness")
	}
	if streamComp.Type == "SWITCH" {
		if streamComp.Expression == nil {
			streamComp.Expression = &SwitchExpression{}
		}
		if err := streamComp.Expression.FromMap(helper.CastMapAny(streamCompMap, "expression")); err != nil {
			return err
		}
	}
	if streamComp.Type == "SWITCH" || streamComp.Type == "CONTAINER" {
		if len(helper.CastSlice(streamCompMap, "elements")) > 0 {
			elements, err := helper.ParseFromMaps[StreamComponent](helper.CastSlice(streamCompMap, "elements"))
			streamComp.Elements = elements
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (repetitive *Repetitive) FromMap(repetitiveMap map[string]any) error {
	repetitive.Value = helper.CastInt(repetitiveMap, "value")
	repetitive.Path = helper.CastString(repetitiveMap, "path")
	return nil
}

func (length *Length) FromMap(lengthMap map[string]any) error {
	length.Type = helper.CastString(lengthMap, "type")
	length.Unit = helper.CastString(lengthMap, "unit")
	length.Value = helper.CastInt(lengthMap, "value")
	length.Path = helper.CastString(lengthMap, "path")
	return nil
}

func (switchExp *SwitchExpression) FromMap(switchExpMap map[string]any) error {
	switchExp.SwitchOn = helper.CastString(switchExpMap, "switch_on")
	if options, err := helper.ParseFromMaps[SwitchOption](helper.CastSlice(switchExpMap, "options")); err != nil {
		return err
	} else {
		switchExp.Options = options
	}
	return nil
}

func (switchOption *SwitchOption) FromMap(switchOptionMap map[string]any) error {
	if err := switchOption.Value.FromMap(helper.CastMapAny(switchOptionMap, "value")); err != nil {
		return err
	}
	switchOption.Component = helper.CastString(switchOptionMap, "component")
	return nil
}

func (switchValue *SwitchValue[T]) FromMap(switchValueMap map[string]any) error {
	switchValue.DataType = helper.CastString(switchValueMap, "data_type")
	switchValue.Data = switchValueMap["data"].(T)
	return nil
}

func (metadata *Metadata) FromMap(metadataMap map[string]any) error {
	if err := metadata.Timestamp.FromMap(helper.CastMapAny(metadataMap, "timestamp")); err != nil {
		return err
	}
	return nil
}

func (timestampDef *TimestampDefinition) FromMap(timestampDefMap map[string]any) error {
	timestampDef.Expression = helper.CastString(timestampDefMap, "expression")
	return nil
}

func (elementList *ElementList[T, PT]) FromMap(elementListMap map[string]any) error {
	if elems, err := helper.ParseFromMaps[T, PT](helper.CastSlice(elementListMap, "elements")); err != nil {
		return err
	} else {
		elementList.Elements = elems
	}
	return nil
}

func (elementList *ElementListWithValid[T, PT]) FromMap(elementListMap map[string]any) error {
	if elems, err := helper.ParseFromMaps[T, PT](helper.CastSlice(elementListMap, "elements")); err != nil {
		return err
	} else {
		elementList.Elements = elems
	}
	elementList.Valid = helper.CastBool(elementListMap, "valid")
	return nil
}

func (computation *Computation) FromMap(computationMap map[string]any) error {
	computation.Name = helper.CastString(computationMap, "name")
	computation.Order = helper.CastInt(computationMap, "order")
	computation.Type = helper.CastString(computationMap, "type")
	computation.DataType = helper.CastString(computationMap, "data_type")
	computation.Expression = helper.CastString(computationMap, "expression")
	return nil
}

func (mapping *Mapping) FromMap(mappingMap map[string]any) error {
	mapping.MetricId = helper.CastString(mappingMap, "metric_id")
	mapping.Expression = helper.CastString(mappingMap, "expression")
	return nil
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64Decode(str string) (string, error) {
	if val, err := base64.StdEncoding.DecodeString(str); err != nil {
		return "", err
	} else {
		return string(val), nil
	}
}

func recursiveUpdateStreamComponent(streamComps []StreamComponent, path string) {
	for index := range streamComps {
		component := &streamComps[index]
		component.Path = path + "." + component.Name
		component.Order = index
		if component.Type == "CONTAINER" || component.Type == "SWITCH" {
			recursiveUpdateStreamComponent(component.Elements, component.Path)
		}
	}
}

func (stream *Stream) PreMarshallProcess() error {
	// Encode expressions to Base64
	computations := stream.Configuration.Computations.Elements
	for i := range computations {
		computations[i].Order = i
		computations[i].Type = "COMPUTATION"
		computations[i].Expression = base64Encode(computations[i].Expression)
	}
	recursiveUpdateStreamComponent(stream.Configuration.Structure.Elements, "structure")
	stream.Configuration.Metadata.Timestamp.Expression = base64Encode(stream.Configuration.Metadata.Timestamp.Expression)
	return nil
}

func (stream *Stream) PostUnmarshallProcess() error {
	computations := stream.Configuration.Computations.Elements
	for i := range computations {
		if value, err := base64Decode(computations[i].Expression); err != nil {
			return err
		} else {
			computations[i].Expression = value
		}
	}
	if value, err := base64Decode(stream.Configuration.Metadata.Timestamp.Expression); err != nil {
		return err
	} else {
		stream.Configuration.Metadata.Timestamp.Expression = value
	}
	return nil
}
