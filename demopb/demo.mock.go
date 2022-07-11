// Code generated by protoc-gen-mock. DO NOT EDIT.
// source: demo.proto

package demopb

import (
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

import (
	context "context"
	errors "errors"
	grpc "google.golang.org/grpc"
	protojson "google.golang.org/protobuf/encoding/protojson"
	status "google.golang.org/genproto/googleapis/rpc/status"
	spb "google.golang.org/grpc/status"
)

var (
	ErrWrongArgType  = errors.New("wrong argument type for this method")
	ErrUnknownMethod = errors.New("unknown method name")
	ErrEmptyResponse = errors.New("empty response to register")
)

type MockApiServer struct {
	UnimplementedApiServer
	contents struct {
		ListPersons  *ListPersonsResponse
		GetPerson    *Person
		CreatePerson *Person
		DeletePerson *emptypb.Empty
	}
	errors struct {
		ListPersons  error
		GetPerson    error
		CreatePerson error
		DeletePerson error
	}
	callbacks struct {
		ListPersons  func(*MockApiServer)
		GetPerson    func(*MockApiServer)
		CreatePerson func(*MockApiServer)
		DeletePerson func(*MockApiServer)
	}
}

// RegisterMockResponse registers a response that is return at method invocation.
func (ms *MockApiServer) RegisterMockResponse(method string, response any) error {
	switch method {
	case "ListPersons":
		switch r := response.(type) {
		case error:
			ms.errors.ListPersons = r
		case *ListPersonsResponse:
			ms.contents.ListPersons = r
		default:
			return ErrWrongArgType
		}
	case "GetPerson":
		switch r := response.(type) {
		case error:
			ms.errors.GetPerson = r
		case *Person:
			ms.contents.GetPerson = r
		default:
			return ErrWrongArgType
		}
	case "CreatePerson":
		switch r := response.(type) {
		case error:
			ms.errors.CreatePerson = r
		case *Person:
			ms.contents.CreatePerson = r
		default:
			return ErrWrongArgType
		}
	case "DeletePerson":
		switch r := response.(type) {
		case error:
			ms.errors.DeletePerson = r
		case *emptypb.Empty:
			ms.contents.DeletePerson = r
		default:
			return ErrWrongArgType
		}
	default:
		return ErrUnknownMethod
	}
	return nil
}

// RegisterMockCallback registers a callback that is called after method invocation.
func (ms *MockApiServer) RegisterMockCallback(method string, callback func(*MockApiServer)) error {
	switch method {
	case "ListPersons":
		ms.callbacks.ListPersons = callback
	case "GetPerson":
		ms.callbacks.GetPerson = callback
	case "CreatePerson":
		ms.callbacks.CreatePerson = callback
	case "DeletePerson":
		ms.callbacks.DeletePerson = callback
	default:
		return ErrUnknownMethod
	}
	return nil
}

// RegisterJSONMockContent registers a JSON string as a Mock content,
// making sure that the format is respected
func (ms *MockApiServer) RegisterJSONMockContent(method string, payload []byte) error {
	switch method {
	case "ListPersons":
		var content = new(ListPersonsResponse)
		if err := protojson.Unmarshal(payload, content); err != nil {
			return err
		}
		ms.contents.ListPersons = content
	case "GetPerson":
		var content = new(Person)
		if err := protojson.Unmarshal(payload, content); err != nil {
			return err
		}
		ms.contents.GetPerson = content
	case "CreatePerson":
		var content = new(Person)
		if err := protojson.Unmarshal(payload, content); err != nil {
			return err
		}
		ms.contents.CreatePerson = content
	case "DeletePerson":
		var content = new(emptypb.Empty)
		if err := protojson.Unmarshal(payload, content); err != nil {
			return err
		}
		ms.contents.DeletePerson = content
	default:
		return ErrUnknownMethod
	}
	return nil
}

// RegisterJSONMockStatus registers a JSON string as a Mock status,
// making sure that the format is respected
func (ms *MockApiServer) RegisterJSONMockStatus(method string, payload []byte) error {
	switch method {
	case "ListPersons":
		var sta = new(status.Status)
		if err := protojson.Unmarshal(payload, sta); err != nil {
			return err
		}
		ms.errors.ListPersons = spb.ErrorProto(sta)
	case "GetPerson":
		var sta = new(status.Status)
		if err := protojson.Unmarshal(payload, sta); err != nil {
			return err
		}
		ms.errors.GetPerson = spb.ErrorProto(sta)
	case "CreatePerson":
		var sta = new(status.Status)
		if err := protojson.Unmarshal(payload, sta); err != nil {
			return err
		}
		ms.errors.CreatePerson = spb.ErrorProto(sta)
	case "DeletePerson":
		var sta = new(status.Status)
		if err := protojson.Unmarshal(payload, sta); err != nil {
			return err
		}
		ms.errors.DeletePerson = spb.ErrorProto(sta)
	default:
		return ErrUnknownMethod
	}
	return nil
}
func (ms *MockApiServer) ListPersons(ctx context.Context, req *ListPersonsRequest) (*ListPersonsResponse, error) {
	if ms.callbacks.ListPersons != nil {
		defer ms.callbacks.ListPersons(ms)
	}
	if ms.errors.ListPersons != nil {
		return nil, ms.errors.ListPersons
	}
	if ms.contents.ListPersons != nil {
		return ms.contents.ListPersons, nil
	}
	return &ListPersonsResponse{
		Persons:    nil,
		TotalCount: 0,
	}, nil
}

func (ms *MockApiServer) GetPerson(ctx context.Context, req *GetPersonRequest) (*Person, error) {
	if ms.callbacks.GetPerson != nil {
		defer ms.callbacks.GetPerson(ms)
	}
	if ms.errors.GetPerson != nil {
		return nil, ms.errors.GetPerson
	}
	if ms.contents.GetPerson != nil {
		return ms.contents.GetPerson, nil
	}
	return &Person{
		Id:    "485c917c-3a44-40be-aecc-89d2684cc9b4",
		Name:  "Gayle Ullrich",
		Email: "kariannedurgan@hettinger.name",
		Type:  0,
	}, nil
}

func (ms *MockApiServer) CreatePerson(ctx context.Context, req *CreatePersonRequest) (*Person, error) {
	if ms.callbacks.CreatePerson != nil {
		defer ms.callbacks.CreatePerson(ms)
	}
	if ms.errors.CreatePerson != nil {
		return nil, ms.errors.CreatePerson
	}
	if ms.contents.CreatePerson != nil {
		return ms.contents.CreatePerson, nil
	}
	return &Person{
		Id:    "f0f99e4a-5399-4d6a-9bbd-278c24aa74ef",
		Name:  "Ross Legros",
		Email: "kevinrau@schiller.info",
		Type:  0,
	}, nil
}

func (ms *MockApiServer) DeletePerson(ctx context.Context, req *DeletePersonRequest) (*emptypb.Empty, error) {
	if ms.callbacks.DeletePerson != nil {
		defer ms.callbacks.DeletePerson(ms)
	}
	if ms.errors.DeletePerson != nil {
		return nil, ms.errors.DeletePerson
	}
	if ms.contents.DeletePerson != nil {
		return ms.contents.DeletePerson, nil
	}
	return &emptypb.Empty{}, nil
}

func RegisterMockApiServer(s grpc.ServiceRegistrar) *MockApiServer {
	ms := &MockApiServer{}
	RegisterApiServer(s, ms)
	return ms
}
