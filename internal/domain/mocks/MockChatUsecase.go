// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockChatUsecase is an autogenerated mock type for the ChatUsecase type
type MockChatUsecase struct {
	mock.Mock
}

type MockChatUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockChatUsecase) EXPECT() *MockChatUsecase_Expecter {
	return &MockChatUsecase_Expecter{mock: &_m.Mock}
}

// AddChatHistory provides a mock function with given fields: ctx, message
func (_m *MockChatUsecase) AddChatHistory(ctx context.Context, message *domain.PersonMessage) error {
	ret := _m.Called(ctx, message)

	if len(ret) == 0 {
		panic("no return value specified for AddChatHistory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.PersonMessage) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockChatUsecase_AddChatHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddChatHistory'
type MockChatUsecase_AddChatHistory_Call struct {
	*mock.Call
}

// AddChatHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - message *domain.PersonMessage
func (_e *MockChatUsecase_Expecter) AddChatHistory(ctx interface{}, message interface{}) *MockChatUsecase_AddChatHistory_Call {
	return &MockChatUsecase_AddChatHistory_Call{Call: _e.mock.On("AddChatHistory", ctx, message)}
}

func (_c *MockChatUsecase_AddChatHistory_Call) Run(run func(ctx context.Context, message *domain.PersonMessage)) *MockChatUsecase_AddChatHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.PersonMessage))
	})
	return _c
}

func (_c *MockChatUsecase_AddChatHistory_Call) Return(_a0 error) *MockChatUsecase_AddChatHistory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChatUsecase_AddChatHistory_Call) RunAndReturn(run func(context.Context, *domain.PersonMessage) error) *MockChatUsecase_AddChatHistory_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonMessage provides a mock function with given fields: ctx, messageID
func (_m *MockChatUsecase) GetPersonMessage(ctx context.Context, messageID int64) (domain.QueryChatHistoryResult, error) {
	ret := _m.Called(ctx, messageID)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonMessage")
	}

	var r0 domain.QueryChatHistoryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.QueryChatHistoryResult, error)); ok {
		return rf(ctx, messageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.QueryChatHistoryResult); ok {
		r0 = rf(ctx, messageID)
	} else {
		r0 = ret.Get(0).(domain.QueryChatHistoryResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChatUsecase_GetPersonMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonMessage'
type MockChatUsecase_GetPersonMessage_Call struct {
	*mock.Call
}

// GetPersonMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - messageID int64
func (_e *MockChatUsecase_Expecter) GetPersonMessage(ctx interface{}, messageID interface{}) *MockChatUsecase_GetPersonMessage_Call {
	return &MockChatUsecase_GetPersonMessage_Call{Call: _e.mock.On("GetPersonMessage", ctx, messageID)}
}

func (_c *MockChatUsecase_GetPersonMessage_Call) Run(run func(ctx context.Context, messageID int64)) *MockChatUsecase_GetPersonMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockChatUsecase_GetPersonMessage_Call) Return(_a0 domain.QueryChatHistoryResult, _a1 error) *MockChatUsecase_GetPersonMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChatUsecase_GetPersonMessage_Call) RunAndReturn(run func(context.Context, int64) (domain.QueryChatHistoryResult, error)) *MockChatUsecase_GetPersonMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonMessageContext provides a mock function with given fields: ctx, messageID, paddedMessageCount
func (_m *MockChatUsecase) GetPersonMessageContext(ctx context.Context, messageID int64, paddedMessageCount int) ([]domain.QueryChatHistoryResult, error) {
	ret := _m.Called(ctx, messageID, paddedMessageCount)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonMessageContext")
	}

	var r0 []domain.QueryChatHistoryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) ([]domain.QueryChatHistoryResult, error)); ok {
		return rf(ctx, messageID, paddedMessageCount)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) []domain.QueryChatHistoryResult); ok {
		r0 = rf(ctx, messageID, paddedMessageCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.QueryChatHistoryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int) error); ok {
		r1 = rf(ctx, messageID, paddedMessageCount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChatUsecase_GetPersonMessageContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonMessageContext'
type MockChatUsecase_GetPersonMessageContext_Call struct {
	*mock.Call
}

// GetPersonMessageContext is a helper method to define mock.On call
//   - ctx context.Context
//   - messageID int64
//   - paddedMessageCount int
func (_e *MockChatUsecase_Expecter) GetPersonMessageContext(ctx interface{}, messageID interface{}, paddedMessageCount interface{}) *MockChatUsecase_GetPersonMessageContext_Call {
	return &MockChatUsecase_GetPersonMessageContext_Call{Call: _e.mock.On("GetPersonMessageContext", ctx, messageID, paddedMessageCount)}
}

func (_c *MockChatUsecase_GetPersonMessageContext_Call) Run(run func(ctx context.Context, messageID int64, paddedMessageCount int)) *MockChatUsecase_GetPersonMessageContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int))
	})
	return _c
}

func (_c *MockChatUsecase_GetPersonMessageContext_Call) Return(_a0 []domain.QueryChatHistoryResult, _a1 error) *MockChatUsecase_GetPersonMessageContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChatUsecase_GetPersonMessageContext_Call) RunAndReturn(run func(context.Context, int64, int) ([]domain.QueryChatHistoryResult, error)) *MockChatUsecase_GetPersonMessageContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryChatHistory provides a mock function with given fields: ctx, user, filters
func (_m *MockChatUsecase) QueryChatHistory(ctx context.Context, user domain.PersonInfo, filters domain.ChatHistoryQueryFilter) ([]domain.QueryChatHistoryResult, error) {
	ret := _m.Called(ctx, user, filters)

	if len(ret) == 0 {
		panic("no return value specified for QueryChatHistory")
	}

	var r0 []domain.QueryChatHistoryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, domain.ChatHistoryQueryFilter) ([]domain.QueryChatHistoryResult, error)); ok {
		return rf(ctx, user, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, domain.ChatHistoryQueryFilter) []domain.QueryChatHistoryResult); ok {
		r0 = rf(ctx, user, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.QueryChatHistoryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PersonInfo, domain.ChatHistoryQueryFilter) error); ok {
		r1 = rf(ctx, user, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChatUsecase_QueryChatHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryChatHistory'
type MockChatUsecase_QueryChatHistory_Call struct {
	*mock.Call
}

// QueryChatHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - user domain.PersonInfo
//   - filters domain.ChatHistoryQueryFilter
func (_e *MockChatUsecase_Expecter) QueryChatHistory(ctx interface{}, user interface{}, filters interface{}) *MockChatUsecase_QueryChatHistory_Call {
	return &MockChatUsecase_QueryChatHistory_Call{Call: _e.mock.On("QueryChatHistory", ctx, user, filters)}
}

func (_c *MockChatUsecase_QueryChatHistory_Call) Run(run func(ctx context.Context, user domain.PersonInfo, filters domain.ChatHistoryQueryFilter)) *MockChatUsecase_QueryChatHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PersonInfo), args[2].(domain.ChatHistoryQueryFilter))
	})
	return _c
}

func (_c *MockChatUsecase_QueryChatHistory_Call) Return(_a0 []domain.QueryChatHistoryResult, _a1 error) *MockChatUsecase_QueryChatHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChatUsecase_QueryChatHistory_Call) RunAndReturn(run func(context.Context, domain.PersonInfo, domain.ChatHistoryQueryFilter) ([]domain.QueryChatHistoryResult, error)) *MockChatUsecase_QueryChatHistory_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *MockChatUsecase) Start(ctx context.Context) {
	_m.Called(ctx)
}

// MockChatUsecase_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockChatUsecase_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockChatUsecase_Expecter) Start(ctx interface{}) *MockChatUsecase_Start_Call {
	return &MockChatUsecase_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *MockChatUsecase_Start_Call) Run(run func(ctx context.Context)) *MockChatUsecase_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockChatUsecase_Start_Call) Return() *MockChatUsecase_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockChatUsecase_Start_Call) RunAndReturn(run func(context.Context)) *MockChatUsecase_Start_Call {
	_c.Call.Return(run)
	return _c
}

// TopChatters provides a mock function with given fields: ctx, count
func (_m *MockChatUsecase) TopChatters(ctx context.Context, count uint64) ([]domain.TopChatterResult, error) {
	ret := _m.Called(ctx, count)

	if len(ret) == 0 {
		panic("no return value specified for TopChatters")
	}

	var r0 []domain.TopChatterResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]domain.TopChatterResult, error)); ok {
		return rf(ctx, count)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []domain.TopChatterResult); ok {
		r0 = rf(ctx, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.TopChatterResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChatUsecase_TopChatters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TopChatters'
type MockChatUsecase_TopChatters_Call struct {
	*mock.Call
}

// TopChatters is a helper method to define mock.On call
//   - ctx context.Context
//   - count uint64
func (_e *MockChatUsecase_Expecter) TopChatters(ctx interface{}, count interface{}) *MockChatUsecase_TopChatters_Call {
	return &MockChatUsecase_TopChatters_Call{Call: _e.mock.On("TopChatters", ctx, count)}
}

func (_c *MockChatUsecase_TopChatters_Call) Run(run func(ctx context.Context, count uint64)) *MockChatUsecase_TopChatters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockChatUsecase_TopChatters_Call) Return(_a0 []domain.TopChatterResult, _a1 error) *MockChatUsecase_TopChatters_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChatUsecase_TopChatters_Call) RunAndReturn(run func(context.Context, uint64) ([]domain.TopChatterResult, error)) *MockChatUsecase_TopChatters_Call {
	_c.Call.Return(run)
	return _c
}

// WarningState provides a mock function with given fields:
func (_m *MockChatUsecase) WarningState() map[string][]domain.UserWarning {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for WarningState")
	}

	var r0 map[string][]domain.UserWarning
	if rf, ok := ret.Get(0).(func() map[string][]domain.UserWarning); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]domain.UserWarning)
		}
	}

	return r0
}

// MockChatUsecase_WarningState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WarningState'
type MockChatUsecase_WarningState_Call struct {
	*mock.Call
}

// WarningState is a helper method to define mock.On call
func (_e *MockChatUsecase_Expecter) WarningState() *MockChatUsecase_WarningState_Call {
	return &MockChatUsecase_WarningState_Call{Call: _e.mock.On("WarningState")}
}

func (_c *MockChatUsecase_WarningState_Call) Run(run func()) *MockChatUsecase_WarningState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockChatUsecase_WarningState_Call) Return(_a0 map[string][]domain.UserWarning) *MockChatUsecase_WarningState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChatUsecase_WarningState_Call) RunAndReturn(run func() map[string][]domain.UserWarning) *MockChatUsecase_WarningState_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockChatUsecase creates a new instance of MockChatUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockChatUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockChatUsecase {
	mock := &MockChatUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
