package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/nlm/protoc-gen-mock/demopb"
	"github.com/nlm/protoc-gen-mock/pkg/pb/scenariopb"
	"github.com/nlm/protoc-gen-mock/pkg/pbutils"
	"github.com/nlm/protoc-gen-mock/pkg/scenarios"
	"google.golang.org/grpc"
)

var (
	// flagConfigFile = flag.String("config", "config.yaml", "config file")
	flagListen   = flag.String("listen", ":9090", "listen address")
	flagScenario = flag.String("scenario", "", "scenario file")
)

func main() {
	flag.Parse()
	server := grpc.NewServer()

	// Register Mock.
	// This will automatically populate naive content mocks
	ms := demopb.RegisterMockApiServer(server)

	// Register using Scenario.
	// This makes possible to load different mock combinations from a data file
	if *flagScenario != "" {
		scenarioData, err := ioutil.ReadFile(*flagScenario)
		if err != nil {
			log.Fatal(err)
		}
		scenario := scenariopb.Scenario{}
		if err := pbutils.ProtoUnmarshalYAML(scenarioData, &scenario); err != nil {
			log.Fatal(err)
		}
		if err := scenarios.RegisterScenario(ms, &scenario); err != nil {
			log.Fatal(err)
		}
	}

	// Direct Registration (JSON).
	// Used by the Scenario facility to register content and status.
	ms.RegisterJSONMockContent("GetPerson", []byte(`{"id": "42", "name": "Bob"}`))

	// Direct Registration (Native).
	// Convenience wrapper giving the ability go register objects, status and errors.
	ms.RegisterMockResponse("GetPerson", &demopb.Person{
		Id:    "123",
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Type:  1,
	})

	// Listen and Serve
	lis, err := net.Listen("tcp", *flagListen)
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)
}
