package et

// MockOption is a function that can be passed to MockEndpoint or MockService to configure the mocking behavior.
type MockOption func(*mockOptions)

type mockOptions struct {
	runMiddleware bool
}

// RunMiddleware is a MockOption that sets whether to run the middleware chain
// prior to invoking the mock.
func RunMiddleware(enabled bool) (_ MockOption) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/et/mocking.go#L20-L24
	doPanic("encore apps must be run using the encore command")
	return
}

// MockEndpoint allows you to mock out an endpoint in your tests; Any calls made to the endpoint
// during this test or any of its sub-tests will be routed to the mock you provide.
//
// Your mocked function must match the signature of the endpoint you are mocking.
//
// For example, if you have an endpoint defined as:
//
//	//encore:api public
//	func MyAPI(ctx context.Context, req *MyAPIRequest) (*MyAPIResponse, error) {
//		...
//	}
//
// You can mock it out in your test as:
//
//	et.MockEndpoint(mysvc.MyAPI, func(ctx context.Context, req *MyAPIRequest) (*MyAPIResponse, error) {
//		...
//	})
//
// If you want to mock out a single endpoint method on a service, you can use the generated helper
// package function as the `originalEndpoint` argument to this function, however if you want to
// mock out more than one API method on a service, consider using [MockService].
//
// Note: if you use [MockService] to mock a service and then use this function to mock
// an endpoint on that service, the endpoint mock will take precedence over the service mock.
//
// Setting the mock to nil will remove the endpoint mock.
func MockEndpoint[T any](originalEndpoint T, mock T, opts ...MockOption) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/et/mocking.go#L52-L72
	doPanic("encore apps must be run using the encore command")
	return
}

// MockService allows you to mock out a service in your tests; Any calls made to the service
// during this test or any of its sub-tests will be routed to the mock you provide.
//
// Your mock must implement the all the API methods of the service which are used during the
// test(s). If you do not implement a method, it will panic when that method is called.
//
// If you want to ensure compile time safety, you can use the Interface type generated for
// the service, which will ensure that you implement all the methods. For example:
//
//	package svca
//
//	import (
//		"testing"
//		"encore.dev/et"
//
//		"encore.app/svcb"
//	)
//
//	func TestServiceA(t *testing.T) {
//		et.MockService[svcb.Interface]("svcb", &myMockType{})
//		SomeFuncInThisPackageWhichUltimatelyCallsServiceB()
//	}
//
// Setting the mock to nil will remove the service mock.
func MockService[T any](serviceName string, mock T, opts ...MockOption) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/et/mocking.go#L98-L109
	doPanic("encore apps must be run using the encore command")
	return
}
