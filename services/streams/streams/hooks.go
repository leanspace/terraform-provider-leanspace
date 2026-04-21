package streams

import (
	"encoding/base64"

	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

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

func (stream *Stream) PostReadProcess(_ *provider.Client, newValue any) error {
	newStream, ok := newValue.(*Stream)
	if !ok || newStream == nil {
		return nil
	}
	newStream.Tags = general_objects.ReorderKeyValues(stream.Tags, newStream.Tags)
	return nil
}
