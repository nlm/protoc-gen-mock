PROG=protoc-gen-mock
DEMOPB=demopb
TESTPB=pkg/pb/testpb
SCENARIOPB=pkg/pb/scenariopb
MOCKPB=pkg/pb/mockpb
PROTOINCLUDE=-I ./protobuf/google -I ./protobuf/grpc/src/proto -I ./pkg/pb
LDFLAGS=-ldflags="-s -w"

.PHONY: build

build: $(PROG)

$(PROG): ./cmd/$(PROG)/*.go ./pkg/*/*.go $(SCENARIOPB)/scenario.pb.go $(MOCKPB)/mockoptions.pb.go
	go build $(LDFLAGS) ./cmd/$(PROG)/

$(SCENARIOPB)/scenario.pb.go: $(SCENARIOPB)/scenario.proto
	protoc -I $(SCENARIOPB) $(PROTOINCLUDE) --go_out=$(SCENARIOPB) --go_opt=paths=source_relative scenario.proto

$(TESTPB)/test.pb.go: $(TESTPB)/test.proto
	protoc -I $(TESTPB) $(PROTOINCLUDE) --go_out=$(TESTPB) --go_opt=paths=source_relative test.proto

$(MOCKPB)/mockoptions.pb.go: $(MOCKPB)/mockoptions.proto
	protoc -I $(MOCKPB) $(PROTOINCLUDE) --go_out=$(MOCKPB) --go_opt=paths=source_relative mockoptions.proto

.PHONY: test

test: $(TESTPB)/test.pb.go
	go test ./cmd/... ./pkg/...

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
	go run ./cmd/demo -scenario ./scenarios/demo1.yaml

.PHONY: rundemogw

rundemogw: proto
	go run ./cmd/demogw

.PHONY: clean

clean:
	rm -f $(PROG)

.PHONY: protoclean

protoclean:
	rm -f $(DEMOPB)/demo.pb.go $(DEMOPB)/demo.mock.go $(DEMOPB)/demo.pb.gw.go $(DEMOPB)/demo_grpc.pb.go
	rm -f $(SCENARIOPB)/scenario.pb.go
	rm -f $(MOCKPB)/annotations.pb.go
