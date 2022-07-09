package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nlm/protoc-gen-mock/demopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	flagListen   = flag.String("listen", ":8080", "listen address")
	flagEndpoint = flag.String("endpoint", "localhost:9090", "grpc server endpoint")
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := demopb.RegisterApiHandlerFromEndpoint(ctx, mux, *flagEndpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(*flagListen, mux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
