package scenarios

import (
	"errors"

	"github.com/nlm/protoc-gen-mock/pkg/pb/scenariopb"
	"google.golang.org/protobuf/encoding/protojson"
)

type JSONMockServer interface {
	RegisterJSONMockContent(method string, data []byte) error
	RegisterJSONMockStatus(method string, data []byte) error
}

var (
	ErrEmptyEndpoint = errors.New("empty endpoint")
)

func RegisterScenario(ms JSONMockServer, s *scenariopb.Scenario) error {
	for k, v := range s.GetEndpoints() {
		if content := v.GetContent(); content != nil {
			bytes, err := protojson.Marshal(content)
			if err != nil {
				return err
			}
			// log.Print(k, " -content-> ", string(bytes))
			if err := ms.RegisterJSONMockContent(k, bytes); err != nil {
				return err
			}
		} else if status := v.GetStatus(); status != nil {
			bytes, err := protojson.Marshal(status)
			if err != nil {
				return err
			}
			// log.Print(k, " -status-> ", string(bytes))
			if err := ms.RegisterJSONMockStatus(k, bytes); err != nil {
				return err
			}
		} else {
			// log.Print("skipping")
			return ErrEmptyEndpoint
		}
	}
	return nil
}
