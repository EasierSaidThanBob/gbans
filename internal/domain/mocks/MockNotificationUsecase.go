// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"

	steamid "github.com/leighmacdonald/steamid/v3/steamid"
)

// MockNotificationUsecase is an autogenerated mock type for the NotificationUsecase type
type MockNotificationUsecase struct {
	mock.Mock
}

type MockNotificationUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNotificationUsecase) EXPECT() *MockNotificationUsecase_Expecter {
	return &MockNotificationUsecase_Expecter{mock: &_m.Mock}
}

// GetPersonNotifications provides a mock function with given fields: ctx, filters
func (_m *MockNotificationUsecase) GetPersonNotifications(ctx context.Context, filters domain.NotificationQuery) ([]domain.UserNotification, int64, error) {
	ret := _m.Called(ctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonNotifications")
	}

	var r0 []domain.UserNotification
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.NotificationQuery) ([]domain.UserNotification, int64, error)); ok {
		return rf(ctx, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.NotificationQuery) []domain.UserNotification); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserNotification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.NotificationQuery) int64); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.NotificationQuery) error); ok {
		r2 = rf(ctx, filters)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockNotificationUsecase_GetPersonNotifications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonNotifications'
type MockNotificationUsecase_GetPersonNotifications_Call struct {
	*mock.Call
}

// GetPersonNotifications is a helper method to define mock.On call
//   - ctx context.Context
//   - filters domain.NotificationQuery
func (_e *MockNotificationUsecase_Expecter) GetPersonNotifications(ctx interface{}, filters interface{}) *MockNotificationUsecase_GetPersonNotifications_Call {
	return &MockNotificationUsecase_GetPersonNotifications_Call{Call: _e.mock.On("GetPersonNotifications", ctx, filters)}
}

func (_c *MockNotificationUsecase_GetPersonNotifications_Call) Run(run func(ctx context.Context, filters domain.NotificationQuery)) *MockNotificationUsecase_GetPersonNotifications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.NotificationQuery))
	})
	return _c
}

func (_c *MockNotificationUsecase_GetPersonNotifications_Call) Return(_a0 []domain.UserNotification, _a1 int64, _a2 error) *MockNotificationUsecase_GetPersonNotifications_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockNotificationUsecase_GetPersonNotifications_Call) RunAndReturn(run func(context.Context, domain.NotificationQuery) ([]domain.UserNotification, int64, error)) *MockNotificationUsecase_GetPersonNotifications_Call {
	_c.Call.Return(run)
	return _c
}

// SendNotification provides a mock function with given fields: ctx, targetID, severity, message, link
func (_m *MockNotificationUsecase) SendNotification(ctx context.Context, targetID steamid.SID64, severity domain.NotificationSeverity, message string, link string) error {
	ret := _m.Called(ctx, targetID, severity, message, link)

	if len(ret) == 0 {
		panic("no return value specified for SendNotification")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SID64, domain.NotificationSeverity, string, string) error); ok {
		r0 = rf(ctx, targetID, severity, message, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationUsecase_SendNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendNotification'
type MockNotificationUsecase_SendNotification_Call struct {
	*mock.Call
}

// SendNotification is a helper method to define mock.On call
//   - ctx context.Context
//   - targetID steamid.SID64
//   - severity domain.NotificationSeverity
//   - message string
//   - link string
func (_e *MockNotificationUsecase_Expecter) SendNotification(ctx interface{}, targetID interface{}, severity interface{}, message interface{}, link interface{}) *MockNotificationUsecase_SendNotification_Call {
	return &MockNotificationUsecase_SendNotification_Call{Call: _e.mock.On("SendNotification", ctx, targetID, severity, message, link)}
}

func (_c *MockNotificationUsecase_SendNotification_Call) Run(run func(ctx context.Context, targetID steamid.SID64, severity domain.NotificationSeverity, message string, link string)) *MockNotificationUsecase_SendNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SID64), args[2].(domain.NotificationSeverity), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockNotificationUsecase_SendNotification_Call) Return(_a0 error) *MockNotificationUsecase_SendNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationUsecase_SendNotification_Call) RunAndReturn(run func(context.Context, steamid.SID64, domain.NotificationSeverity, string, string) error) *MockNotificationUsecase_SendNotification_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNotificationUsecase creates a new instance of MockNotificationUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNotificationUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNotificationUsecase {
	mock := &MockNotificationUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
