package streams

import (
	"encoding/base64"
	"strconv"

	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (stream *Stream) ToMap() map[string]any {
	streamMap := make(map[string]any)
	streamMap["id"] = stream.ID
	streamMap["version"] = stream.Version
	streamMap["name"] = stream.Name
	streamMap["description"] = stream.Description
	streamMap["asset_id"] = stream.AssetId
	streamMap["configuration"] = []any{stream.Configuration.ToMap()}
	streamMap["mappings"] = helper.ParseToMaps(stream.Mappings)
	streamMap["created_at"] = stream.CreatedAt
	streamMap["created_by"] = stream.CreatedBy
	streamMap["last_modified_at"] = stream.LastModifiedAt
	streamMap["last_modified_by"] = stream.LastModifiedBy
	return streamMap
}

func (configuration *Configuration) ToMap() map[string]any {
	configMap := make(map[string]any)
	configMap["endianness"] = configuration.Endianness
	configMap["structure"] = []any{configuration.Structure.ToMap()}
	configMap["metadata"] = []any{configuration.Metadata.ToMap()}
	configMap["computations"] = []any{configuration.Computations.ToMap()}
	configMap["valid"] = configuration.Valid
	configMap["errors"] = helper.ParseToMaps(configuration.Errors)
	return configMap
}

func (streamComp *StreamComponent) ToMap() map[string]any {
	streamCompMap := make(map[string]any)
	streamCompMap["name"] = streamComp.Name
	streamCompMap["order"] = streamComp.Order
	streamCompMap["path"] = streamComp.Path
	streamCompMap["type"] = streamComp.Type
	streamCompMap["valid"] = streamComp.Valid
	streamCompMap["errors"] = helper.ParseToMaps(streamComp.Errors)

	if streamComp.Repetitive != nil {
		streamCompMap["repetitive"] = []map[string]any{streamComp.Repetitive.ToMap()}
	}

	if streamComp.Type == "FIELD" {
		streamCompMap["length"] = []map[string]any{streamComp.Length.ToMap()}
		streamCompMap["processor"] = streamComp.Processor
		streamCompMap["data_type"] = streamComp.DataType
		streamCompMap["endianness"] = streamComp.Endianness
	}
	if streamComp.Type == "SWITCH" {
		streamCompMap["expression"] = []any{streamComp.Expression.ToMap()}
	}
	if streamComp.Type == "SWITCH" || streamComp.Type == "CONTAINER" {
		streamCompMap["elements"] = helper.ParseToMaps(streamComp.Elements)
	}

	return streamCompMap
}

func (repetitive *Repetitive) ToMap() map[string]any {
	repetitiveMap := make(map[string]any)
	if repetitive != nil && repetitive.Value != 0 {
		repetitiveMap["value"] = repetitive.Value
	}
	if repetitive != nil && repetitive.Path != "" {
		repetitiveMap["path"] = repetitive.Path
	}
	return repetitiveMap
}

func (length *Length) ToMap() map[string]any {
	lengthMap := make(map[string]any)
	lengthMap["type"] = length.Type
	lengthMap["unit"] = length.Unit
	lengthMap["value"] = length.Value
	lengthMap["path"] = length.Path
	return lengthMap
}

func (switchExp *SwitchExpression) ToMap() map[string]any {
	switchExpMap := make(map[string]any)
	switchExpMap["switch_on"] = switchExp.SwitchOn
	switchExpMap["options"] = helper.ParseToMaps(switchExp.Options)
	return switchExpMap
}

func (switchOption *SwitchOption) ToMap() map[string]any {
	switchOptionMap := make(map[string]any)
	switchOptionMap["value"] = []any{switchOption.Value.ToMap()}
	switchOptionMap["component"] = switchOption.Component
	return switchOptionMap
}

func (switchValue *SwitchValue[T]) ToMap() map[string]any {
	switchValueMap := make(map[string]any)
	switchValueMap["data_type"] = switchValue.DataType
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
	metadataMap["packet_id"] = []any{metadata.PacketID.ToMap()}
	metadataMap["timestamp"] = []any{metadata.Timestamp.ToMap()}
	metadataMap["valid"] = metadata.Valid
	metadataMap["errors"] = helper.ParseToMaps(metadata.Errors)
	return metadataMap
}

func (timestampDef *TimestampDefinition) ToMap() map[string]any {
	timestampDefMap := make(map[string]any)
	timestampDefMap["expression"] = timestampDef.Expression
	timestampDefMap["valid"] = timestampDef.Valid
	timestampDefMap["errors"] = helper.ParseToMaps(timestampDef.Errors)
	return timestampDefMap
}

func (elementList *ElementList[T, PT]) ToMap() map[string]any {
	elementListMap := make(map[string]any)
	elementListMap["elements"] = helper.ParseToMaps[T, PT](elementList.Elements)
	elementListMap["valid"] = elementList.Valid
	elementListMap["errors"] = helper.ParseToMaps(elementList.Errors)
	return elementListMap
}

func (computation *Computation) ToMap() map[string]any {
	computationMap := make(map[string]any)
	computationMap["name"] = computation.Name
	computationMap["order"] = computation.Order
	computationMap["type"] = computation.Type
	computationMap["data_type"] = computation.DataType
	computationMap["expression"] = computation.Expression
	computationMap["valid"] = computation.Valid
	computationMap["errors"] = helper.ParseToMaps(computation.Errors)
	return computationMap
}

func (mapping *Mapping) ToMap() map[string]any {
	mappingMap := make(map[string]any)
	mappingMap["metric_id"] = mapping.MetricId
	mappingMap["component"] = mapping.Component
	mappingMap["expression"] = mapping.Expression
	return mappingMap
}

func (elemStatus *ElementStatus) ToMap() map[string]any {
	elemStatusMap := make(map[string]any)
	elemStatusMap["valid"] = elemStatus.Valid
	elemStatusMap["errors"] = helper.ParseToMaps(elemStatus.Errors)
	return elemStatusMap
}

func (err *Error) ToMap() map[string]any {
	errMap := make(map[string]any)
	errMap["code"] = err.Code
	errMap["message"] = err.Message
	return errMap
}

func (stream *Stream) FromMap(streamMap map[string]any) error {
	stream.ID = streamMap["id"].(string)
	stream.Version = streamMap["version"].(int)
	stream.Name = streamMap["name"].(string)
	stream.Description = streamMap["description"].(string)
	stream.AssetId = streamMap["asset_id"].(string)
	if len(streamMap["configuration"].([]any)) > 0 {
		if err := stream.Configuration.FromMap(streamMap["configuration"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if mappings, err := helper.ParseFromMaps[Mapping](streamMap["mappings"].(*schema.Set).List()); err != nil {
		return err
	} else {
		stream.Mappings = mappings
	}
	stream.CreatedAt = streamMap["created_at"].(string)
	stream.CreatedBy = streamMap["created_by"].(string)
	stream.LastModifiedAt = streamMap["last_modified_at"].(string)
	stream.LastModifiedBy = streamMap["last_modified_by"].(string)

	return nil
}

func (configuration *Configuration) FromMap(configMap map[string]any) error {
	configuration.Endianness = configMap["endianness"].(string)
	if len(configMap["structure"].([]any)) > 0 {
		if err := configuration.Structure.FromMap(configMap["structure"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if len(configMap["metadata"].([]any)) > 0 {
		if err := configuration.Metadata.FromMap(configMap["metadata"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if len(configMap["computations"].([]any)) > 0 {
		if err := configuration.Computations.FromMap(configMap["computations"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	configuration.Valid = configMap["valid"].(bool)
	if errors, err := helper.ParseFromMaps[Error](configMap["errors"].(*schema.Set).List()); err != nil {
		return err
	} else {
		configuration.Errors = errors
	}
	return nil
}

func (streamComp *StreamComponent) FromMap(streamCompMap map[string]any) error {
	streamComp.Name = streamCompMap["name"].(string)
	streamComp.Order = streamCompMap["order"].(int)
	streamComp.Path = streamCompMap["path"].(string)
	streamComp.Type = streamCompMap["type"].(string)
	streamComp.Valid = streamCompMap["valid"].(bool)
	if len(streamCompMap["repetitive"].([]any)) > 0 && streamCompMap["repetitive"].([]any)[0] != nil {
		streamComp.Repetitive = new(Repetitive)
		if err := streamComp.Repetitive.FromMap(streamCompMap["repetitive"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	errors, err := helper.ParseFromMaps[Error](streamCompMap["errors"].(*schema.Set).List())
	streamComp.Errors = errors
	if err != nil {
		return err
	}

	if streamComp.Type == "FIELD" {
		if len(streamCompMap["length"].([]any)) > 0 && streamCompMap["length"].([]any)[0] != nil {
			streamComp.Length = new(Length)
			if err := streamComp.Length.FromMap(streamCompMap["length"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}
		streamComp.Processor = streamCompMap["processor"].(string)
		streamComp.DataType = streamCompMap["data_type"].(string)
		streamComp.Endianness = streamCompMap["endianness"].(string)
	}
	if streamComp.Type == "SWITCH" {
		if len(streamCompMap["expression"].([]any)) > 0 {
			if err := streamComp.Expression.FromMap(streamCompMap["expression"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}
	}
	if streamComp.Type == "SWITCH" || streamComp.Type == "CONTAINER" {
		if len(streamCompMap["elements"].([]any)) > 0 {
			elements, err := helper.ParseFromMaps[StreamComponent](streamCompMap["elements"].([]any))
			streamComp.Elements = elements
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (repetitive *Repetitive) FromMap(repetitiveMap map[string]any) error {
	repetitive.Value = repetitiveMap["value"].(int)
	repetitive.Path = repetitiveMap["path"].(string)
	return nil
}

func (length *Length) FromMap(lengthMap map[string]any) error {
	length.Type = lengthMap["type"].(string)
	length.Unit = lengthMap["unit"].(string)
	length.Value = lengthMap["value"].(int)
	length.Path = lengthMap["path"].(string)
	return nil
}

func (switchExp *SwitchExpression) FromMap(switchExpMap map[string]any) error {
	switchExp.SwitchOn = switchExpMap["switch_on"].(string)
	if options, err := helper.ParseFromMaps[SwitchOption](switchExpMap["options"].([]any)); err != nil {
		return err
	} else {
		switchExp.Options = options
	}
	return nil
}

func (switchOption *SwitchOption) FromMap(switchOptionMap map[string]any) error {
	if err := switchOption.Value.FromMap(switchOptionMap["value"].([]any)[0].(map[string]any)); err != nil {
		return err
	}
	switchOption.Component = switchOptionMap["component"].(string)
	return nil
}

func (switchValue *SwitchValue[T]) FromMap(switchValueMap map[string]any) error {
	switchValue.DataType = switchValueMap["data_type"].(string)
	switchValue.Data = switchValueMap["data"].(T)
	return nil
}

func (metadata *Metadata) FromMap(metadataMap map[string]any) error {
	if len(metadataMap["packet_id"].([]any)) > 0 {
		if err := metadata.PacketID.FromMap(metadataMap["packet_id"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if len(metadataMap["timestamp"].([]any)) > 0 {
		if err := metadata.Timestamp.FromMap(metadataMap["timestamp"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	metadata.Valid = metadataMap["valid"].(bool)
	if errors, err := helper.ParseFromMaps[Error](metadataMap["errors"].(*schema.Set).List()); err != nil {
		return err
	} else {
		metadata.Errors = errors
	}
	return nil
}

func (timestampDef *TimestampDefinition) FromMap(timestampDefMap map[string]any) error {
	timestampDef.Expression = timestampDefMap["expression"].(string)
	timestampDef.Valid = timestampDefMap["valid"].(bool)
	if errors, err := helper.ParseFromMaps[Error](timestampDefMap["errors"].(*schema.Set).List()); err != nil {
		return err
	} else {
		timestampDef.Errors = errors
	}
	return nil
}

func (elementList *ElementList[T, PT]) FromMap(elementListMap map[string]any) error {
	if elems, err := helper.ParseFromMaps[T, PT](elementListMap["elements"].([]any)); err != nil {
		return err
	} else {
		elementList.Elements = elems
	}
	elementList.Valid = elementListMap["valid"].(bool)
	if errors, err := helper.ParseFromMaps[Error](elementListMap["errors"].(*schema.Set).List()); err != nil {
		return err
	} else {
		elementList.Errors = errors
	}
	return nil
}

func (computation *Computation) FromMap(computationMap map[string]any) error {
	computation.Name = computationMap["name"].(string)
	computation.Order = computationMap["order"].(int)
	computation.Type = computationMap["type"].(string)
	computation.DataType = computationMap["data_type"].(string)
	computation.Expression = computationMap["expression"].(string)
	computation.Valid = computationMap["valid"].(bool)
	if errors, err := helper.ParseFromMaps[Error](computationMap["errors"].(*schema.Set).List()); err != nil {
		return err
	} else {
		computation.Errors = errors
	}
	return nil
}

func (mapping *Mapping) FromMap(mappingMap map[string]any) error {
	mapping.MetricId = mappingMap["metric_id"].(string)
	mapping.Component = mappingMap["component"].(string)
	mapping.Expression = mappingMap["expression"].(string)
	return nil
}

func (elemStatus *ElementStatus) FromMap(elemStatusMap map[string]any) error {
	elemStatus.Valid = elemStatusMap["valid"].(bool)
	if errors, err := helper.ParseFromMaps[Error](elemStatusMap["errors"].(*schema.Set).List()); err != nil {
		return err
	} else {
		elemStatus.Errors = errors
	}
	return nil
}

func (err *Error) FromMap(errorMap map[string]any) error {
	err.Code = errorMap["code"].(string)
	err.Message = errorMap["message"].(string)
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
