package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/nlm/protoc-gen-mock/demopb"
	"github.com/nlm/protoc-gen-mock/pkg/scenariopb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Mocks map[string]any
}

var (
	// flagConfigFile = flag.String("config", "config.yaml", "config file")
	flagListen   = flag.String("listen", ":9090", "listen address")
	flagScenario = flag.String("scenario", "", "scenario file")
)

// ProtoUnmarshalYaml converts YAML to JSON and Unmarshals it to a Message
func ProtoUnmarshalYaml(input []byte, m proto.Message) error {
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

func RegisterScenario(ms *demopb.MockApiServer, s *scenariopb.Scenario) error {
	for k, v := range s.GetEndpoints() {
		if content := v.GetContent(); content != nil {
			bytes, err := protojson.Marshal(v)
			if err != nil {
				return err
			}
			log.Print(k, " -content-> ", string(bytes))
			if err := ms.RegisterJSONMockContent(k, bytes); err != nil {
				return err
			}
		} else if status := v.GetStatus(); status != nil {
			bytes, err := protojson.Marshal(v)
			if err != nil {
				return err
			}
			log.Print(k, " -status-> ", string(bytes))
			if err := ms.RegisterJSONMockStatus(k, bytes); err != nil {
				return err
			}
		} else {
			log.Print("skipping")
		}
	}
	return nil
}

func main() {
	flag.Parse()
	server := grpc.NewServer()

	// Classic Register
	// demoApiServer := DemoApiServer{}
	// demopb.RegisterApiServer(server, &demoApiServer)

	// Mock Register
	ms := demopb.RegisterMockApiServer(server)

	// var config Config
	// configData, err := ioutil.ReadFile(*flagConfigFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := yaml.Unmarshal(configData, &config); err != nil {
	// 	log.Fatal(err)
	// }
	// for name, content := range config.Mocks {
	// 	ms.RegisterJSONMockResponse(name, content)
	// }
	// var config Config
	if *flagScenario != "" {
		scenarioData, err := ioutil.ReadFile(*flagScenario)
		if err != nil {
			log.Fatal(err)
		}
		scenario := scenariopb.Scenario{}
		if err := ProtoUnmarshalYaml(scenarioData, &scenario); err != nil {
			log.Fatal(err)
		}
		bytes, err := protojson.Marshal(&scenario)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(bytes))

		if err := RegisterScenario(ms, &scenario); err != nil {
			log.Fatal(err)
		}
	}
	// var scenario interface{}
	// if err := yaml.Unmarshal(scenarioData, &scenario); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(scenario)
	// if err := protojson.Unmarshal(scenarioData, &scenario); err != nil {
	// 	log.Fatal(err)
	// }
	// bytes, err := prototext.Marshal(&scenario)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(string(bytes))

	ms.RegisterMockResponse("GetPerson", &demopb.Person{
		Id:    "123",
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Type:  1,
	})
	//ms.RegisterMockResponse("ListPersons", errors.New("example error"))
	// ms.RegisterJSONMockResponse("GetPerson", []byte("{\"name\": \"bob\"}"))

	// Listen
	lis, err := net.Listen("tcp", *flagListen)
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)
}
