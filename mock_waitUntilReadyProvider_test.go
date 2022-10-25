// Code generated by mockery v2.14.0. DO NOT EDIT.

package gocb

import (
	context "context"

	gocbcore "github.com/couchbase/gocbcore/v10"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// mockWaitUntilReadyProvider is an autogenerated mock type for the waitUntilReadyProvider type
type mockWaitUntilReadyProvider struct {
	mock.Mock
}

// WaitUntilReady provides a mock function with given fields: ctx, deadline, opts
func (_m *mockWaitUntilReadyProvider) WaitUntilReady(ctx context.Context, deadline time.Time, opts gocbcore.WaitUntilReadyOptions) error {
	ret := _m.Called(ctx, deadline, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, gocbcore.WaitUntilReadyOptions) error); ok {
		r0 = rf(ctx, deadline, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewMockWaitUntilReadyProvider interface {
	mock.TestingT
	Cleanup(func())
}

// newMockWaitUntilReadyProvider creates a new instance of mockWaitUntilReadyProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockWaitUntilReadyProvider(t mockConstructorTestingTnewMockWaitUntilReadyProvider) *mockWaitUntilReadyProvider {
	mock := &mockWaitUntilReadyProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
