// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockSRCDSUsecase is an autogenerated mock type for the SRCDSUsecase type
type MockSRCDSUsecase struct {
	mock.Mock
}

type MockSRCDSUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSRCDSUsecase) EXPECT() *MockSRCDSUsecase_Expecter {
	return &MockSRCDSUsecase_Expecter{mock: &_m.Mock}
}

// Report provides a mock function with given fields: ctx, currentUser, req
func (_m *MockSRCDSUsecase) Report(ctx context.Context, currentUser domain.UserProfile, req domain.CreateReportReq) (*domain.Report, error) {
	ret := _m.Called(ctx, currentUser, req)

	if len(ret) == 0 {
		panic("no return value specified for Report")
	}

	var r0 *domain.Report
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, domain.CreateReportReq) (*domain.Report, error)); ok {
		return rf(ctx, currentUser, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, domain.CreateReportReq) *domain.Report); ok {
		r0 = rf(ctx, currentUser, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Report)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.UserProfile, domain.CreateReportReq) error); ok {
		r1 = rf(ctx, currentUser, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSRCDSUsecase_Report_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Report'
type MockSRCDSUsecase_Report_Call struct {
	*mock.Call
}

// Report is a helper method to define mock.On call
//   - ctx context.Context
//   - currentUser domain.UserProfile
//   - req domain.CreateReportReq
func (_e *MockSRCDSUsecase_Expecter) Report(ctx interface{}, currentUser interface{}, req interface{}) *MockSRCDSUsecase_Report_Call {
	return &MockSRCDSUsecase_Report_Call{Call: _e.mock.On("Report", ctx, currentUser, req)}
}

func (_c *MockSRCDSUsecase_Report_Call) Run(run func(ctx context.Context, currentUser domain.UserProfile, req domain.CreateReportReq)) *MockSRCDSUsecase_Report_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.UserProfile), args[2].(domain.CreateReportReq))
	})
	return _c
}

func (_c *MockSRCDSUsecase_Report_Call) Return(_a0 *domain.Report, _a1 error) *MockSRCDSUsecase_Report_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSRCDSUsecase_Report_Call) RunAndReturn(run func(context.Context, domain.UserProfile, domain.CreateReportReq) (*domain.Report, error)) *MockSRCDSUsecase_Report_Call {
	_c.Call.Return(run)
	return _c
}

// ServerAuth provides a mock function with given fields: ctx, req
func (_m *MockSRCDSUsecase) ServerAuth(ctx context.Context, req domain.ServerAuthReq) (string, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ServerAuth")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ServerAuthReq) (string, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ServerAuthReq) string); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ServerAuthReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSRCDSUsecase_ServerAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServerAuth'
type MockSRCDSUsecase_ServerAuth_Call struct {
	*mock.Call
}

// ServerAuth is a helper method to define mock.On call
//   - ctx context.Context
//   - req domain.ServerAuthReq
func (_e *MockSRCDSUsecase_Expecter) ServerAuth(ctx interface{}, req interface{}) *MockSRCDSUsecase_ServerAuth_Call {
	return &MockSRCDSUsecase_ServerAuth_Call{Call: _e.mock.On("ServerAuth", ctx, req)}
}

func (_c *MockSRCDSUsecase_ServerAuth_Call) Run(run func(ctx context.Context, req domain.ServerAuthReq)) *MockSRCDSUsecase_ServerAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ServerAuthReq))
	})
	return _c
}

func (_c *MockSRCDSUsecase_ServerAuth_Call) Return(_a0 string, _a1 error) *MockSRCDSUsecase_ServerAuth_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSRCDSUsecase_ServerAuth_Call) RunAndReturn(run func(context.Context, domain.ServerAuthReq) (string, error)) *MockSRCDSUsecase_ServerAuth_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSRCDSUsecase creates a new instance of MockSRCDSUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSRCDSUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSRCDSUsecase {
	mock := &MockSRCDSUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
