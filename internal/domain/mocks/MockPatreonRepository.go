// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockPatreonRepository is an autogenerated mock type for the PatreonRepository type
type MockPatreonRepository struct {
	mock.Mock
}

type MockPatreonRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPatreonRepository) EXPECT() *MockPatreonRepository_Expecter {
	return &MockPatreonRepository_Expecter{mock: &_m.Mock}
}

// GetPatreonAuth provides a mock function with given fields: ctx
func (_m *MockPatreonRepository) GetPatreonAuth(ctx context.Context) (string, string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPatreonAuth")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) string); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockPatreonRepository_GetPatreonAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPatreonAuth'
type MockPatreonRepository_GetPatreonAuth_Call struct {
	*mock.Call
}

// GetPatreonAuth is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPatreonRepository_Expecter) GetPatreonAuth(ctx interface{}) *MockPatreonRepository_GetPatreonAuth_Call {
	return &MockPatreonRepository_GetPatreonAuth_Call{Call: _e.mock.On("GetPatreonAuth", ctx)}
}

func (_c *MockPatreonRepository_GetPatreonAuth_Call) Run(run func(ctx context.Context)) *MockPatreonRepository_GetPatreonAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPatreonRepository_GetPatreonAuth_Call) Return(_a0 string, _a1 string, _a2 error) *MockPatreonRepository_GetPatreonAuth_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockPatreonRepository_GetPatreonAuth_Call) RunAndReturn(run func(context.Context) (string, string, error)) *MockPatreonRepository_GetPatreonAuth_Call {
	_c.Call.Return(run)
	return _c
}

// SetPatreonAuth provides a mock function with given fields: ctx, accessToken, refreshToken
func (_m *MockPatreonRepository) SetPatreonAuth(ctx context.Context, accessToken string, refreshToken string) error {
	ret := _m.Called(ctx, accessToken, refreshToken)

	if len(ret) == 0 {
		panic("no return value specified for SetPatreonAuth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, accessToken, refreshToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPatreonRepository_SetPatreonAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPatreonAuth'
type MockPatreonRepository_SetPatreonAuth_Call struct {
	*mock.Call
}

// SetPatreonAuth is a helper method to define mock.On call
//   - ctx context.Context
//   - accessToken string
//   - refreshToken string
func (_e *MockPatreonRepository_Expecter) SetPatreonAuth(ctx interface{}, accessToken interface{}, refreshToken interface{}) *MockPatreonRepository_SetPatreonAuth_Call {
	return &MockPatreonRepository_SetPatreonAuth_Call{Call: _e.mock.On("SetPatreonAuth", ctx, accessToken, refreshToken)}
}

func (_c *MockPatreonRepository_SetPatreonAuth_Call) Run(run func(ctx context.Context, accessToken string, refreshToken string)) *MockPatreonRepository_SetPatreonAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockPatreonRepository_SetPatreonAuth_Call) Return(_a0 error) *MockPatreonRepository_SetPatreonAuth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPatreonRepository_SetPatreonAuth_Call) RunAndReturn(run func(context.Context, string, string) error) *MockPatreonRepository_SetPatreonAuth_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPatreonRepository creates a new instance of MockPatreonRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPatreonRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPatreonRepository {
	mock := &MockPatreonRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
