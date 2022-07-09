package demopb

import "google.golang.org/grpc"
import "context"

type MockApiServerResponses struct {
	ListPersons *ListPersonsResponse
	GetPerson   *Person
}
type MockApiServerErrors struct {
	ListPersons error
	GetPerson   error
}
type MockApiServer struct {
	UnimplementedApiServer
	responses MockApiServerResponses
	errors    MockApiServerErrors
}

// func (ms *MockApiServer) RegisterListPersonsResponse(response *ListPersonsResponse) {
// ms.responses.ListPersons = response
// }
// func (ms *MockApiServer) RegisterGetPersonResponse(response *Person) {
// ms.responses.GetPerson = response
// }
func (ms *MockApiServer) RegisterMockResponse(method string, response any) {
	switch method {
	case "ListPersons":
		switch r := response.(type) {
		case error:
			ms.errors.ListPersons = r
		case *ListPersonsResponse:
			ms.responses.ListPersons = r
		default:
			panic("wrong argument type for this method")
		}
	case "GetPerson":
		switch r := response.(type) {
		case error:
			ms.errors.GetPerson = r
		case *Person:
			ms.responses.GetPerson = r
		default:
			panic("wrong argument type for this method")
		}
	default:
		panic("unknown method name: " + method)
	}
}
func (ms *MockApiServer) ListPersons(ctx context.Context, req *ListPersonsRequest) (*ListPersonsResponse, error) {
	if ms.errors.ListPersons != nil {
		return nil, ms.errors.ListPersons
	}
	if ms.responses.ListPersons != nil {
		return ms.responses.ListPersons, nil
	}
	return &ListPersonsResponse{
		// Options: [] true
		Persons: nil,
		// Options: [] true
		TotalCount: 0,
	}, nil
}
func (ms *MockApiServer) GetPerson(ctx context.Context, req *GetPersonRequest) (*Person, error) {
	if ms.errors.GetPerson != nil {
		return nil, ms.errors.GetPerson
	}
	if ms.responses.GetPerson != nil {
		return ms.responses.GetPerson, nil
	}
	return &Person{
		// Options: [] true
		Id: "string",
		// Options: [] true
		Name: "string",
		// Options: [] true
		Email: "string",
		// Options: [] true
		Type: 0,
	}, nil
}
func RegisterMockApiServer(s grpc.ServiceRegistrar) *MockApiServer {
	ms := &MockApiServer{}
	RegisterApiServer(s, ms)
	return ms
}
