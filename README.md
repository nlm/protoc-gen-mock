# protoc-gen-mock

mock generator from protobuf and grpc

## build

build the `protoc-gen-mock` plugin

```sh
make build
```

## run the demo

This demo includes a gRPC server and a gRPC gateway for convenience.
What you will want to do is play with `cmd/demo/main.go` to manipulate mocks.

get submodules (if you want/need to regen proto files)

```sh
git submodule init
git submodule update
```

run the demo gRPC server, it will listen on port 9090

```bash
make rundemo
```

in another terminal, run the demo gRPC gateway, it will listen on port 8080

```sh
make rundemogw
```

in another terminal, query the gRPC gateway

```sh
$ curl -s http://localhost:8080/api/v1/persons
{"code":2,"message":"example error","details":[]}

$ curl -s http://localhost:8080/api/v1/persons/jdoe
{"id":"123","name":"John Doe","email":"jdoe@example.com"}
```