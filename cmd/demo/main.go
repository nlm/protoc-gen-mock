package main

import (
	"errors"
	"flag"
	"log"
	"net"

	"github.com/nlm/protoc-gen-mock/demopb"
	"google.golang.org/grpc"
)

// type DemoApiServer struct {
// 	demopb.UnimplementedApiServer
// }

var (
	flagListen = flag.String("listen", ":9090", "listen address")
)

func main() {
	flag.Parse()
	server := grpc.NewServer()

	// Classic Register
	// demoApiServer := DemoApiServer{}
	// demopb.RegisterApiServer(server, &demoApiServer)

	// Mock Register
	ms := demopb.RegisterMockApiServer(server)
	ms.RegisterMockResponse("GetPerson", &demopb.Person{
		Id:    "123",
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Type:  1,
	})
	ms.RegisterMockResponse("ListPersons", errors.New("example error"))

	// Listen
	lis, err := net.Listen("tcp", *flagListen)
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)
}
