// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMetricsUsecase is an autogenerated mock type for the MetricsUsecase type
type MockMetricsUsecase struct {
	mock.Mock
}

type MockMetricsUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetricsUsecase) EXPECT() *MockMetricsUsecase_Expecter {
	return &MockMetricsUsecase_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with given fields: ctx
func (_m *MockMetricsUsecase) Start(ctx context.Context) {
	_m.Called(ctx)
}

// MockMetricsUsecase_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockMetricsUsecase_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMetricsUsecase_Expecter) Start(ctx interface{}) *MockMetricsUsecase_Start_Call {
	return &MockMetricsUsecase_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *MockMetricsUsecase_Start_Call) Run(run func(ctx context.Context)) *MockMetricsUsecase_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMetricsUsecase_Start_Call) Return() *MockMetricsUsecase_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMetricsUsecase_Start_Call) RunAndReturn(run func(context.Context)) *MockMetricsUsecase_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetricsUsecase creates a new instance of MockMetricsUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetricsUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetricsUsecase {
	mock := &MockMetricsUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}