// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"

	steamid "github.com/leighmacdonald/steamid/v4/steamid"
)

// MockReportRepository is an autogenerated mock type for the ReportRepository type
type MockReportRepository struct {
	mock.Mock
}

type MockReportRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReportRepository) EXPECT() *MockReportRepository_Expecter {
	return &MockReportRepository_Expecter{mock: &_m.Mock}
}

// DropReport provides a mock function with given fields: ctx, report
func (_m *MockReportRepository) DropReport(ctx context.Context, report *domain.Report) error {
	ret := _m.Called(ctx, report)

	if len(ret) == 0 {
		panic("no return value specified for DropReport")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Report) error); ok {
		r0 = rf(ctx, report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReportRepository_DropReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropReport'
type MockReportRepository_DropReport_Call struct {
	*mock.Call
}

// DropReport is a helper method to define mock.On call
//   - ctx context.Context
//   - report *domain.Report
func (_e *MockReportRepository_Expecter) DropReport(ctx interface{}, report interface{}) *MockReportRepository_DropReport_Call {
	return &MockReportRepository_DropReport_Call{Call: _e.mock.On("DropReport", ctx, report)}
}

func (_c *MockReportRepository_DropReport_Call) Run(run func(ctx context.Context, report *domain.Report)) *MockReportRepository_DropReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Report))
	})
	return _c
}

func (_c *MockReportRepository_DropReport_Call) Return(_a0 error) *MockReportRepository_DropReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReportRepository_DropReport_Call) RunAndReturn(run func(context.Context, *domain.Report) error) *MockReportRepository_DropReport_Call {
	_c.Call.Return(run)
	return _c
}

// DropReportMessage provides a mock function with given fields: ctx, message
func (_m *MockReportRepository) DropReportMessage(ctx context.Context, message *domain.ReportMessage) error {
	ret := _m.Called(ctx, message)

	if len(ret) == 0 {
		panic("no return value specified for DropReportMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ReportMessage) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReportRepository_DropReportMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropReportMessage'
type MockReportRepository_DropReportMessage_Call struct {
	*mock.Call
}

// DropReportMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - message *domain.ReportMessage
func (_e *MockReportRepository_Expecter) DropReportMessage(ctx interface{}, message interface{}) *MockReportRepository_DropReportMessage_Call {
	return &MockReportRepository_DropReportMessage_Call{Call: _e.mock.On("DropReportMessage", ctx, message)}
}

func (_c *MockReportRepository_DropReportMessage_Call) Run(run func(ctx context.Context, message *domain.ReportMessage)) *MockReportRepository_DropReportMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.ReportMessage))
	})
	return _c
}

func (_c *MockReportRepository_DropReportMessage_Call) Return(_a0 error) *MockReportRepository_DropReportMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReportRepository_DropReportMessage_Call) RunAndReturn(run func(context.Context, *domain.ReportMessage) error) *MockReportRepository_DropReportMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetReport provides a mock function with given fields: ctx, reportID
func (_m *MockReportRepository) GetReport(ctx context.Context, reportID int64) (domain.Report, error) {
	ret := _m.Called(ctx, reportID)

	if len(ret) == 0 {
		panic("no return value specified for GetReport")
	}

	var r0 domain.Report
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.Report, error)); ok {
		return rf(ctx, reportID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Report); ok {
		r0 = rf(ctx, reportID)
	} else {
		r0 = ret.Get(0).(domain.Report)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, reportID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReportRepository_GetReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReport'
type MockReportRepository_GetReport_Call struct {
	*mock.Call
}

// GetReport is a helper method to define mock.On call
//   - ctx context.Context
//   - reportID int64
func (_e *MockReportRepository_Expecter) GetReport(ctx interface{}, reportID interface{}) *MockReportRepository_GetReport_Call {
	return &MockReportRepository_GetReport_Call{Call: _e.mock.On("GetReport", ctx, reportID)}
}

func (_c *MockReportRepository_GetReport_Call) Run(run func(ctx context.Context, reportID int64)) *MockReportRepository_GetReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockReportRepository_GetReport_Call) Return(_a0 domain.Report, _a1 error) *MockReportRepository_GetReport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReportRepository_GetReport_Call) RunAndReturn(run func(context.Context, int64) (domain.Report, error)) *MockReportRepository_GetReport_Call {
	_c.Call.Return(run)
	return _c
}

// GetReportBySteamID provides a mock function with given fields: ctx, authorID, steamID
func (_m *MockReportRepository) GetReportBySteamID(ctx context.Context, authorID steamid.SteamID, steamID steamid.SteamID) (domain.Report, error) {
	ret := _m.Called(ctx, authorID, steamID)

	if len(ret) == 0 {
		panic("no return value specified for GetReportBySteamID")
	}

	var r0 domain.Report
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, steamid.SteamID) (domain.Report, error)); ok {
		return rf(ctx, authorID, steamID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, steamid.SteamID) domain.Report); ok {
		r0 = rf(ctx, authorID, steamID)
	} else {
		r0 = ret.Get(0).(domain.Report)
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.SteamID, steamid.SteamID) error); ok {
		r1 = rf(ctx, authorID, steamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReportRepository_GetReportBySteamID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReportBySteamID'
type MockReportRepository_GetReportBySteamID_Call struct {
	*mock.Call
}

// GetReportBySteamID is a helper method to define mock.On call
//   - ctx context.Context
//   - authorID steamid.SteamID
//   - steamID steamid.SteamID
func (_e *MockReportRepository_Expecter) GetReportBySteamID(ctx interface{}, authorID interface{}, steamID interface{}) *MockReportRepository_GetReportBySteamID_Call {
	return &MockReportRepository_GetReportBySteamID_Call{Call: _e.mock.On("GetReportBySteamID", ctx, authorID, steamID)}
}

func (_c *MockReportRepository_GetReportBySteamID_Call) Run(run func(ctx context.Context, authorID steamid.SteamID, steamID steamid.SteamID)) *MockReportRepository_GetReportBySteamID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID), args[2].(steamid.SteamID))
	})
	return _c
}

func (_c *MockReportRepository_GetReportBySteamID_Call) Return(_a0 domain.Report, _a1 error) *MockReportRepository_GetReportBySteamID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReportRepository_GetReportBySteamID_Call) RunAndReturn(run func(context.Context, steamid.SteamID, steamid.SteamID) (domain.Report, error)) *MockReportRepository_GetReportBySteamID_Call {
	_c.Call.Return(run)
	return _c
}

// GetReportMessageByID provides a mock function with given fields: ctx, reportMessageID
func (_m *MockReportRepository) GetReportMessageByID(ctx context.Context, reportMessageID int64) (domain.ReportMessage, error) {
	ret := _m.Called(ctx, reportMessageID)

	if len(ret) == 0 {
		panic("no return value specified for GetReportMessageByID")
	}

	var r0 domain.ReportMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.ReportMessage, error)); ok {
		return rf(ctx, reportMessageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.ReportMessage); ok {
		r0 = rf(ctx, reportMessageID)
	} else {
		r0 = ret.Get(0).(domain.ReportMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, reportMessageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReportRepository_GetReportMessageByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReportMessageByID'
type MockReportRepository_GetReportMessageByID_Call struct {
	*mock.Call
}

// GetReportMessageByID is a helper method to define mock.On call
//   - ctx context.Context
//   - reportMessageID int64
func (_e *MockReportRepository_Expecter) GetReportMessageByID(ctx interface{}, reportMessageID interface{}) *MockReportRepository_GetReportMessageByID_Call {
	return &MockReportRepository_GetReportMessageByID_Call{Call: _e.mock.On("GetReportMessageByID", ctx, reportMessageID)}
}

func (_c *MockReportRepository_GetReportMessageByID_Call) Run(run func(ctx context.Context, reportMessageID int64)) *MockReportRepository_GetReportMessageByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockReportRepository_GetReportMessageByID_Call) Return(_a0 domain.ReportMessage, _a1 error) *MockReportRepository_GetReportMessageByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReportRepository_GetReportMessageByID_Call) RunAndReturn(run func(context.Context, int64) (domain.ReportMessage, error)) *MockReportRepository_GetReportMessageByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetReportMessages provides a mock function with given fields: ctx, reportID
func (_m *MockReportRepository) GetReportMessages(ctx context.Context, reportID int64) ([]domain.ReportMessage, error) {
	ret := _m.Called(ctx, reportID)

	if len(ret) == 0 {
		panic("no return value specified for GetReportMessages")
	}

	var r0 []domain.ReportMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]domain.ReportMessage, error)); ok {
		return rf(ctx, reportID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []domain.ReportMessage); ok {
		r0 = rf(ctx, reportID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ReportMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, reportID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReportRepository_GetReportMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReportMessages'
type MockReportRepository_GetReportMessages_Call struct {
	*mock.Call
}

// GetReportMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - reportID int64
func (_e *MockReportRepository_Expecter) GetReportMessages(ctx interface{}, reportID interface{}) *MockReportRepository_GetReportMessages_Call {
	return &MockReportRepository_GetReportMessages_Call{Call: _e.mock.On("GetReportMessages", ctx, reportID)}
}

func (_c *MockReportRepository_GetReportMessages_Call) Run(run func(ctx context.Context, reportID int64)) *MockReportRepository_GetReportMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockReportRepository_GetReportMessages_Call) Return(_a0 []domain.ReportMessage, _a1 error) *MockReportRepository_GetReportMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReportRepository_GetReportMessages_Call) RunAndReturn(run func(context.Context, int64) ([]domain.ReportMessage, error)) *MockReportRepository_GetReportMessages_Call {
	_c.Call.Return(run)
	return _c
}

// GetReports provides a mock function with given fields: ctx, opts
func (_m *MockReportRepository) GetReports(ctx context.Context, opts domain.ReportQueryFilter) ([]domain.Report, int64, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetReports")
	}

	var r0 []domain.Report
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ReportQueryFilter) ([]domain.Report, int64, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ReportQueryFilter) []domain.Report); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Report)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ReportQueryFilter) int64); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.ReportQueryFilter) error); ok {
		r2 = rf(ctx, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockReportRepository_GetReports_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReports'
type MockReportRepository_GetReports_Call struct {
	*mock.Call
}

// GetReports is a helper method to define mock.On call
//   - ctx context.Context
//   - opts domain.ReportQueryFilter
func (_e *MockReportRepository_Expecter) GetReports(ctx interface{}, opts interface{}) *MockReportRepository_GetReports_Call {
	return &MockReportRepository_GetReports_Call{Call: _e.mock.On("GetReports", ctx, opts)}
}

func (_c *MockReportRepository_GetReports_Call) Run(run func(ctx context.Context, opts domain.ReportQueryFilter)) *MockReportRepository_GetReports_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ReportQueryFilter))
	})
	return _c
}

func (_c *MockReportRepository_GetReports_Call) Return(_a0 []domain.Report, _a1 int64, _a2 error) *MockReportRepository_GetReports_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockReportRepository_GetReports_Call) RunAndReturn(run func(context.Context, domain.ReportQueryFilter) ([]domain.Report, int64, error)) *MockReportRepository_GetReports_Call {
	_c.Call.Return(run)
	return _c
}

// SaveReport provides a mock function with given fields: ctx, report
func (_m *MockReportRepository) SaveReport(ctx context.Context, report *domain.Report) error {
	ret := _m.Called(ctx, report)

	if len(ret) == 0 {
		panic("no return value specified for SaveReport")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Report) error); ok {
		r0 = rf(ctx, report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReportRepository_SaveReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveReport'
type MockReportRepository_SaveReport_Call struct {
	*mock.Call
}

// SaveReport is a helper method to define mock.On call
//   - ctx context.Context
//   - report *domain.Report
func (_e *MockReportRepository_Expecter) SaveReport(ctx interface{}, report interface{}) *MockReportRepository_SaveReport_Call {
	return &MockReportRepository_SaveReport_Call{Call: _e.mock.On("SaveReport", ctx, report)}
}

func (_c *MockReportRepository_SaveReport_Call) Run(run func(ctx context.Context, report *domain.Report)) *MockReportRepository_SaveReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Report))
	})
	return _c
}

func (_c *MockReportRepository_SaveReport_Call) Return(_a0 error) *MockReportRepository_SaveReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReportRepository_SaveReport_Call) RunAndReturn(run func(context.Context, *domain.Report) error) *MockReportRepository_SaveReport_Call {
	_c.Call.Return(run)
	return _c
}

// SaveReportMessage provides a mock function with given fields: ctx, message
func (_m *MockReportRepository) SaveReportMessage(ctx context.Context, message *domain.ReportMessage) error {
	ret := _m.Called(ctx, message)

	if len(ret) == 0 {
		panic("no return value specified for SaveReportMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ReportMessage) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReportRepository_SaveReportMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveReportMessage'
type MockReportRepository_SaveReportMessage_Call struct {
	*mock.Call
}

// SaveReportMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - message *domain.ReportMessage
func (_e *MockReportRepository_Expecter) SaveReportMessage(ctx interface{}, message interface{}) *MockReportRepository_SaveReportMessage_Call {
	return &MockReportRepository_SaveReportMessage_Call{Call: _e.mock.On("SaveReportMessage", ctx, message)}
}

func (_c *MockReportRepository_SaveReportMessage_Call) Run(run func(ctx context.Context, message *domain.ReportMessage)) *MockReportRepository_SaveReportMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.ReportMessage))
	})
	return _c
}

func (_c *MockReportRepository_SaveReportMessage_Call) Return(_a0 error) *MockReportRepository_SaveReportMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReportRepository_SaveReportMessage_Call) RunAndReturn(run func(context.Context, *domain.ReportMessage) error) *MockReportRepository_SaveReportMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReportRepository creates a new instance of MockReportRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReportRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReportRepository {
	mock := &MockReportRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
