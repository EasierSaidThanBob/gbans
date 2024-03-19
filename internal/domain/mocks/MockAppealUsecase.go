// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockAppealUsecase is an autogenerated mock type for the AppealUsecase type
type MockAppealUsecase struct {
	mock.Mock
}

type MockAppealUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAppealUsecase) EXPECT() *MockAppealUsecase_Expecter {
	return &MockAppealUsecase_Expecter{mock: &_m.Mock}
}

// CreateBanMessage provides a mock function with given fields: ctx, curUser, banID, newMsg
func (_m *MockAppealUsecase) CreateBanMessage(ctx context.Context, curUser domain.UserProfile, banID int64, newMsg string) (domain.BanAppealMessage, error) {
	ret := _m.Called(ctx, curUser, banID, newMsg)

	if len(ret) == 0 {
		panic("no return value specified for CreateBanMessage")
	}

	var r0 domain.BanAppealMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, int64, string) (domain.BanAppealMessage, error)); ok {
		return rf(ctx, curUser, banID, newMsg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, int64, string) domain.BanAppealMessage); ok {
		r0 = rf(ctx, curUser, banID, newMsg)
	} else {
		r0 = ret.Get(0).(domain.BanAppealMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.UserProfile, int64, string) error); ok {
		r1 = rf(ctx, curUser, banID, newMsg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAppealUsecase_CreateBanMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBanMessage'
type MockAppealUsecase_CreateBanMessage_Call struct {
	*mock.Call
}

// CreateBanMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - curUser domain.UserProfile
//   - banID int64
//   - newMsg string
func (_e *MockAppealUsecase_Expecter) CreateBanMessage(ctx interface{}, curUser interface{}, banID interface{}, newMsg interface{}) *MockAppealUsecase_CreateBanMessage_Call {
	return &MockAppealUsecase_CreateBanMessage_Call{Call: _e.mock.On("CreateBanMessage", ctx, curUser, banID, newMsg)}
}

func (_c *MockAppealUsecase_CreateBanMessage_Call) Run(run func(ctx context.Context, curUser domain.UserProfile, banID int64, newMsg string)) *MockAppealUsecase_CreateBanMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.UserProfile), args[2].(int64), args[3].(string))
	})
	return _c
}

func (_c *MockAppealUsecase_CreateBanMessage_Call) Return(_a0 domain.BanAppealMessage, _a1 error) *MockAppealUsecase_CreateBanMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAppealUsecase_CreateBanMessage_Call) RunAndReturn(run func(context.Context, domain.UserProfile, int64, string) (domain.BanAppealMessage, error)) *MockAppealUsecase_CreateBanMessage_Call {
	_c.Call.Return(run)
	return _c
}

// DropBanMessage provides a mock function with given fields: ctx, curUser, banMessageID
func (_m *MockAppealUsecase) DropBanMessage(ctx context.Context, curUser domain.UserProfile, banMessageID int64) error {
	ret := _m.Called(ctx, curUser, banMessageID)

	if len(ret) == 0 {
		panic("no return value specified for DropBanMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, int64) error); ok {
		r0 = rf(ctx, curUser, banMessageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAppealUsecase_DropBanMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropBanMessage'
type MockAppealUsecase_DropBanMessage_Call struct {
	*mock.Call
}

// DropBanMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - curUser domain.UserProfile
//   - banMessageID int64
func (_e *MockAppealUsecase_Expecter) DropBanMessage(ctx interface{}, curUser interface{}, banMessageID interface{}) *MockAppealUsecase_DropBanMessage_Call {
	return &MockAppealUsecase_DropBanMessage_Call{Call: _e.mock.On("DropBanMessage", ctx, curUser, banMessageID)}
}

func (_c *MockAppealUsecase_DropBanMessage_Call) Run(run func(ctx context.Context, curUser domain.UserProfile, banMessageID int64)) *MockAppealUsecase_DropBanMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.UserProfile), args[2].(int64))
	})
	return _c
}

func (_c *MockAppealUsecase_DropBanMessage_Call) Return(_a0 error) *MockAppealUsecase_DropBanMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAppealUsecase_DropBanMessage_Call) RunAndReturn(run func(context.Context, domain.UserProfile, int64) error) *MockAppealUsecase_DropBanMessage_Call {
	_c.Call.Return(run)
	return _c
}

// EditBanMessage provides a mock function with given fields: ctx, curUser, reportMessageID, newMsg
func (_m *MockAppealUsecase) EditBanMessage(ctx context.Context, curUser domain.UserProfile, reportMessageID int64, newMsg string) (domain.BanAppealMessage, error) {
	ret := _m.Called(ctx, curUser, reportMessageID, newMsg)

	if len(ret) == 0 {
		panic("no return value specified for EditBanMessage")
	}

	var r0 domain.BanAppealMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, int64, string) (domain.BanAppealMessage, error)); ok {
		return rf(ctx, curUser, reportMessageID, newMsg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, int64, string) domain.BanAppealMessage); ok {
		r0 = rf(ctx, curUser, reportMessageID, newMsg)
	} else {
		r0 = ret.Get(0).(domain.BanAppealMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.UserProfile, int64, string) error); ok {
		r1 = rf(ctx, curUser, reportMessageID, newMsg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAppealUsecase_EditBanMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditBanMessage'
type MockAppealUsecase_EditBanMessage_Call struct {
	*mock.Call
}

// EditBanMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - curUser domain.UserProfile
//   - reportMessageID int64
//   - newMsg string
func (_e *MockAppealUsecase_Expecter) EditBanMessage(ctx interface{}, curUser interface{}, reportMessageID interface{}, newMsg interface{}) *MockAppealUsecase_EditBanMessage_Call {
	return &MockAppealUsecase_EditBanMessage_Call{Call: _e.mock.On("EditBanMessage", ctx, curUser, reportMessageID, newMsg)}
}

func (_c *MockAppealUsecase_EditBanMessage_Call) Run(run func(ctx context.Context, curUser domain.UserProfile, reportMessageID int64, newMsg string)) *MockAppealUsecase_EditBanMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.UserProfile), args[2].(int64), args[3].(string))
	})
	return _c
}

func (_c *MockAppealUsecase_EditBanMessage_Call) Return(_a0 domain.BanAppealMessage, _a1 error) *MockAppealUsecase_EditBanMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAppealUsecase_EditBanMessage_Call) RunAndReturn(run func(context.Context, domain.UserProfile, int64, string) (domain.BanAppealMessage, error)) *MockAppealUsecase_EditBanMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetAppealsByActivity provides a mock function with given fields: ctx, opts
func (_m *MockAppealUsecase) GetAppealsByActivity(ctx context.Context, opts domain.AppealQueryFilter) ([]domain.AppealOverview, int64, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetAppealsByActivity")
	}

	var r0 []domain.AppealOverview
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.AppealQueryFilter) ([]domain.AppealOverview, int64, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.AppealQueryFilter) []domain.AppealOverview); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.AppealOverview)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.AppealQueryFilter) int64); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.AppealQueryFilter) error); ok {
		r2 = rf(ctx, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAppealUsecase_GetAppealsByActivity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAppealsByActivity'
type MockAppealUsecase_GetAppealsByActivity_Call struct {
	*mock.Call
}

// GetAppealsByActivity is a helper method to define mock.On call
//   - ctx context.Context
//   - opts domain.AppealQueryFilter
func (_e *MockAppealUsecase_Expecter) GetAppealsByActivity(ctx interface{}, opts interface{}) *MockAppealUsecase_GetAppealsByActivity_Call {
	return &MockAppealUsecase_GetAppealsByActivity_Call{Call: _e.mock.On("GetAppealsByActivity", ctx, opts)}
}

func (_c *MockAppealUsecase_GetAppealsByActivity_Call) Run(run func(ctx context.Context, opts domain.AppealQueryFilter)) *MockAppealUsecase_GetAppealsByActivity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.AppealQueryFilter))
	})
	return _c
}

func (_c *MockAppealUsecase_GetAppealsByActivity_Call) Return(_a0 []domain.AppealOverview, _a1 int64, _a2 error) *MockAppealUsecase_GetAppealsByActivity_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAppealUsecase_GetAppealsByActivity_Call) RunAndReturn(run func(context.Context, domain.AppealQueryFilter) ([]domain.AppealOverview, int64, error)) *MockAppealUsecase_GetAppealsByActivity_Call {
	_c.Call.Return(run)
	return _c
}

// GetBanMessageByID provides a mock function with given fields: ctx, banMessageID
func (_m *MockAppealUsecase) GetBanMessageByID(ctx context.Context, banMessageID int64) (domain.BanAppealMessage, error) {
	ret := _m.Called(ctx, banMessageID)

	if len(ret) == 0 {
		panic("no return value specified for GetBanMessageByID")
	}

	var r0 domain.BanAppealMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.BanAppealMessage, error)); ok {
		return rf(ctx, banMessageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.BanAppealMessage); ok {
		r0 = rf(ctx, banMessageID)
	} else {
		r0 = ret.Get(0).(domain.BanAppealMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, banMessageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAppealUsecase_GetBanMessageByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBanMessageByID'
type MockAppealUsecase_GetBanMessageByID_Call struct {
	*mock.Call
}

// GetBanMessageByID is a helper method to define mock.On call
//   - ctx context.Context
//   - banMessageID int64
func (_e *MockAppealUsecase_Expecter) GetBanMessageByID(ctx interface{}, banMessageID interface{}) *MockAppealUsecase_GetBanMessageByID_Call {
	return &MockAppealUsecase_GetBanMessageByID_Call{Call: _e.mock.On("GetBanMessageByID", ctx, banMessageID)}
}

func (_c *MockAppealUsecase_GetBanMessageByID_Call) Run(run func(ctx context.Context, banMessageID int64)) *MockAppealUsecase_GetBanMessageByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockAppealUsecase_GetBanMessageByID_Call) Return(_a0 domain.BanAppealMessage, _a1 error) *MockAppealUsecase_GetBanMessageByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAppealUsecase_GetBanMessageByID_Call) RunAndReturn(run func(context.Context, int64) (domain.BanAppealMessage, error)) *MockAppealUsecase_GetBanMessageByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetBanMessages provides a mock function with given fields: ctx, userProfile, banID
func (_m *MockAppealUsecase) GetBanMessages(ctx context.Context, userProfile domain.UserProfile, banID int64) ([]domain.BanAppealMessage, error) {
	ret := _m.Called(ctx, userProfile, banID)

	if len(ret) == 0 {
		panic("no return value specified for GetBanMessages")
	}

	var r0 []domain.BanAppealMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, int64) ([]domain.BanAppealMessage, error)); ok {
		return rf(ctx, userProfile, banID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserProfile, int64) []domain.BanAppealMessage); ok {
		r0 = rf(ctx, userProfile, banID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BanAppealMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.UserProfile, int64) error); ok {
		r1 = rf(ctx, userProfile, banID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAppealUsecase_GetBanMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBanMessages'
type MockAppealUsecase_GetBanMessages_Call struct {
	*mock.Call
}

// GetBanMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - userProfile domain.UserProfile
//   - banID int64
func (_e *MockAppealUsecase_Expecter) GetBanMessages(ctx interface{}, userProfile interface{}, banID interface{}) *MockAppealUsecase_GetBanMessages_Call {
	return &MockAppealUsecase_GetBanMessages_Call{Call: _e.mock.On("GetBanMessages", ctx, userProfile, banID)}
}

func (_c *MockAppealUsecase_GetBanMessages_Call) Run(run func(ctx context.Context, userProfile domain.UserProfile, banID int64)) *MockAppealUsecase_GetBanMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.UserProfile), args[2].(int64))
	})
	return _c
}

func (_c *MockAppealUsecase_GetBanMessages_Call) Return(_a0 []domain.BanAppealMessage, _a1 error) *MockAppealUsecase_GetBanMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAppealUsecase_GetBanMessages_Call) RunAndReturn(run func(context.Context, domain.UserProfile, int64) ([]domain.BanAppealMessage, error)) *MockAppealUsecase_GetBanMessages_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAppealUsecase creates a new instance of MockAppealUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAppealUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAppealUsecase {
	mock := &MockAppealUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
