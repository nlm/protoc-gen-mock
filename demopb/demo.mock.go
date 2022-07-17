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
	protomock "github.com/nlm/protoc-gen-mock/pkg/protomock"
)

var (
	ErrWrongArgType  = errors.New("wrong argument type for this method")
	ErrUnknownMethod = errors.New("unknown method name")
	ErrEmptyResponse = errors.New("empty response to register")
)

// Api is the main service

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
	defaults struct {
		ListPersons  *ListPersonsResponse
		GetPerson    *Person
		CreatePerson *Person
		DeletePerson *emptypb.Empty
	}
}

// RegisterMockResponse registers a response that is returned at method invocation.
func (ms *MockApiServer) RegisterMockResponse(method string, response any) error {
	switch method {
	// ListPersons lists the persons present in the database.

	case "ListPersons":
		switch r := response.(type) {
		case error:
			ms.errors.ListPersons = r
		case *ListPersonsResponse:
			ms.contents.ListPersons = r
		default:
			return ErrWrongArgType
		}
	// GetPerson retrives one person from the database.

	case "GetPerson":
		switch r := response.(type) {
		case error:
			ms.errors.GetPerson = r
		case *Person:
			ms.contents.GetPerson = r
		default:
			return ErrWrongArgType
		}
	// CreatePerson creates a new person and stores it in the database.

	case "CreatePerson":
		switch r := response.(type) {
		case error:
			ms.errors.CreatePerson = r
		case *Person:
			ms.contents.CreatePerson = r
		default:
			return ErrWrongArgType
		}
	// DeletePerson remove a person from the database

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

// ListPersons lists the persons present in the database.

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
	return ms.defaults.ListPersons, nil
}

// GetPerson retrives one person from the database.

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
	return ms.defaults.GetPerson, nil
}

// CreatePerson creates a new person and stores it in the database.

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
	return ms.defaults.CreatePerson, nil
}

// DeletePerson remove a person from the database

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
	return ms.defaults.DeletePerson, nil
}

func (ms *MockApiServer) initDefaults() {
	protomock.Seed(6904809466881647088)
	ms.defaults.ListPersons = new(ListPersonsResponse)
	protomock.Mock(ms.defaults.ListPersons)
	ms.defaults.GetPerson = new(Person)
	protomock.Mock(ms.defaults.GetPerson)
	ms.defaults.CreatePerson = new(Person)
	protomock.Mock(ms.defaults.CreatePerson)
	ms.defaults.DeletePerson = new(emptypb.Empty)
	protomock.Mock(ms.defaults.DeletePerson)
}

func RegisterMockApiServer(s grpc.ServiceRegistrar) *MockApiServer {
	ms := &MockApiServer{}
	ms.initDefaults()
	RegisterApiServer(s, ms)
	return ms
}
