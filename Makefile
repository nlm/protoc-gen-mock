PROG=protoc-gen-mock
DEMOPB=demopb
PROTOINCLUDE=-I $(DEMOPB) -I ./protobuf/google -I ./protobuf/grpc/src/proto

.PHONY: build

build: $(PROG)

$(PROG): ./cmd/$(PROG)/*.go
	go build ./cmd/$(PROG)/

.PHONY: proto demoproto demo clean

proto: $(DEMOPB)/demo.proto
	protoc $(PROTOINCLUDE) --go_out=$(DEMOPB) --go-grpc_out=$(DEMOPB) --grpc-gateway_out=$(DEMOPB) demo.proto

mockproto: $(DEMOPB)/demo.proto $(PROG)
	protoc $(PROTOINCLUDE) --mock_out=$(DEMOPB) --plugin $(PROG)=$(PROG) demo.proto
	@echo "-----"
	@cat $(DEMOPB)/demo.mock.go
	@echo "-----"

rundemo: proto mockproto
	go run ./cmd/demo

rundemogw: proto mockproto
	go run ./cmd/demogw

clean:
	rm -f $(PROG) demo/demo.demo.go
