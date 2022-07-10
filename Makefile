PROG=protoc-gen-mock
DEMOPB=demopb
SCENARIOPB=pkg/pb/scenariopb
PROTOINCLUDE=-I ./protobuf/google -I ./protobuf/grpc/src/proto

.PHONY: build

build: $(PROG)

$(PROG): ./cmd/$(PROG)/*.go $(SCENARIOPB)/scenario.pb.go
	go build ./cmd/$(PROG)/

$(SCENARIOPB)/scenario.pb.go: $(SCENARIOPB)/scenario.proto
	protoc -I $(SCENARIOPB) $(PROTOINCLUDE) --go_out=$(SCENARIOPB) --go_opt=paths=source_relative scenario.proto

.PHONY: proto

proto: $(DEMOPB)/demo.pb.go $(DEMOPB)/demo_grpc.pb.go $(DEMOPB)/demo.pb.gw.go $(DEMOPB)/demo.mock.go

$(DEMOPB)/demo.pb.go: $(DEMOPB)/demo.proto
	protoc -I $(DEMOPB) $(PROTOINCLUDE) --go_out=$(DEMOPB) --go_opt=paths=source_relative demo.proto

$(DEMOPB)/demo_grpc.pb.go: $(DEMOPB)/demo.proto
	protoc -I $(DEMOPB) $(PROTOINCLUDE) --go-grpc_out=$(DEMOPB) --go-grpc_opt=paths=source_relative demo.proto

$(DEMOPB)/demo.pb.gw.go: $(DEMOPB)/demo.proto
	protoc -I $(DEMOPB) $(PROTOINCLUDE) --grpc-gateway_out=$(DEMOPB) --grpc-gateway_opt=paths=source_relative demo.proto

$(DEMOPB)/demo.mock.go: $(DEMOPB)/demo.proto $(PROG)
	protoc -I $(DEMOPB) $(PROTOINCLUDE) --mock_out=$(DEMOPB) --plugin $(PROG)=$(PROG) --mock_opt=paths=source_relative demo.proto
	@echo "-----"
	@cat $(DEMOPB)/demo.mock.go
	@echo "-----"

.PHONY: rundemo

rundemo: proto
	go run ./cmd/demo

.PHONY: rundemogw

rundemogw: proto
	go run ./cmd/demogw

.PHONY: clean

clean:
	rm -f $(PROG)
