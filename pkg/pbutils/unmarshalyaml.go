package pbutils

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

// ProtoUnmarshalYAML converts YAML to JSON and Unmarshals it to a Message
func ProtoUnmarshalYAML(input []byte, m proto.Message) error {
	var iface interface{}
	if err := yaml.Unmarshal(input, &iface); err != nil {
		return err
	}
	jsonData, err := json.Marshal(iface)
	if err != nil {
		return err
	}
	return protojson.UnmarshalOptions{DiscardUnknown: false}.Unmarshal(jsonData, m)
}
