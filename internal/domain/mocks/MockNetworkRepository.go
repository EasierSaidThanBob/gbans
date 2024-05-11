// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	ip2location "github.com/leighmacdonald/gbans/pkg/ip2location"

	mock "github.com/stretchr/testify/mock"

	net "net"

	netip "net/netip"

	steamid "github.com/leighmacdonald/steamid/v4/steamid"
)

// MockNetworkRepository is an autogenerated mock type for the NetworkRepository type
type MockNetworkRepository struct {
	mock.Mock
}

type MockNetworkRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNetworkRepository) EXPECT() *MockNetworkRepository_Expecter {
	return &MockNetworkRepository_Expecter{mock: &_m.Mock}
}

// AddConnectionHistory provides a mock function with given fields: ctx, conn
func (_m *MockNetworkRepository) AddConnectionHistory(ctx context.Context, conn *domain.PersonConnection) error {
	ret := _m.Called(ctx, conn)

	if len(ret) == 0 {
		panic("no return value specified for AddConnectionHistory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.PersonConnection) error); ok {
		r0 = rf(ctx, conn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNetworkRepository_AddConnectionHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddConnectionHistory'
type MockNetworkRepository_AddConnectionHistory_Call struct {
	*mock.Call
}

// AddConnectionHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - conn *domain.PersonConnection
func (_e *MockNetworkRepository_Expecter) AddConnectionHistory(ctx interface{}, conn interface{}) *MockNetworkRepository_AddConnectionHistory_Call {
	return &MockNetworkRepository_AddConnectionHistory_Call{Call: _e.mock.On("AddConnectionHistory", ctx, conn)}
}

func (_c *MockNetworkRepository_AddConnectionHistory_Call) Run(run func(ctx context.Context, conn *domain.PersonConnection)) *MockNetworkRepository_AddConnectionHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.PersonConnection))
	})
	return _c
}

func (_c *MockNetworkRepository_AddConnectionHistory_Call) Return(_a0 error) *MockNetworkRepository_AddConnectionHistory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkRepository_AddConnectionHistory_Call) RunAndReturn(run func(context.Context, *domain.PersonConnection) error) *MockNetworkRepository_AddConnectionHistory_Call {
	_c.Call.Return(run)
	return _c
}

// GetASNRecordByIP provides a mock function with given fields: ctx, ipAddr
func (_m *MockNetworkRepository) GetASNRecordByIP(ctx context.Context, ipAddr netip.Addr) (domain.NetworkASN, error) {
	ret := _m.Called(ctx, ipAddr)

	if len(ret) == 0 {
		panic("no return value specified for GetASNRecordByIP")
	}

	var r0 domain.NetworkASN
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) (domain.NetworkASN, error)); ok {
		return rf(ctx, ipAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) domain.NetworkASN); ok {
		r0 = rf(ctx, ipAddr)
	} else {
		r0 = ret.Get(0).(domain.NetworkASN)
	}

	if rf, ok := ret.Get(1).(func(context.Context, netip.Addr) error); ok {
		r1 = rf(ctx, ipAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworkRepository_GetASNRecordByIP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetASNRecordByIP'
type MockNetworkRepository_GetASNRecordByIP_Call struct {
	*mock.Call
}

// GetASNRecordByIP is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAddr netip.Addr
func (_e *MockNetworkRepository_Expecter) GetASNRecordByIP(ctx interface{}, ipAddr interface{}) *MockNetworkRepository_GetASNRecordByIP_Call {
	return &MockNetworkRepository_GetASNRecordByIP_Call{Call: _e.mock.On("GetASNRecordByIP", ctx, ipAddr)}
}

func (_c *MockNetworkRepository_GetASNRecordByIP_Call) Run(run func(ctx context.Context, ipAddr netip.Addr)) *MockNetworkRepository_GetASNRecordByIP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(netip.Addr))
	})
	return _c
}

func (_c *MockNetworkRepository_GetASNRecordByIP_Call) Return(_a0 domain.NetworkASN, _a1 error) *MockNetworkRepository_GetASNRecordByIP_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkRepository_GetASNRecordByIP_Call) RunAndReturn(run func(context.Context, netip.Addr) (domain.NetworkASN, error)) *MockNetworkRepository_GetASNRecordByIP_Call {
	_c.Call.Return(run)
	return _c
}

// GetASNRecordsByNum provides a mock function with given fields: ctx, asNum
func (_m *MockNetworkRepository) GetASNRecordsByNum(ctx context.Context, asNum int64) ([]domain.NetworkASN, error) {
	ret := _m.Called(ctx, asNum)

	if len(ret) == 0 {
		panic("no return value specified for GetASNRecordsByNum")
	}

	var r0 []domain.NetworkASN
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]domain.NetworkASN, error)); ok {
		return rf(ctx, asNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []domain.NetworkASN); ok {
		r0 = rf(ctx, asNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.NetworkASN)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, asNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworkRepository_GetASNRecordsByNum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetASNRecordsByNum'
type MockNetworkRepository_GetASNRecordsByNum_Call struct {
	*mock.Call
}

// GetASNRecordsByNum is a helper method to define mock.On call
//   - ctx context.Context
//   - asNum int64
func (_e *MockNetworkRepository_Expecter) GetASNRecordsByNum(ctx interface{}, asNum interface{}) *MockNetworkRepository_GetASNRecordsByNum_Call {
	return &MockNetworkRepository_GetASNRecordsByNum_Call{Call: _e.mock.On("GetASNRecordsByNum", ctx, asNum)}
}

func (_c *MockNetworkRepository_GetASNRecordsByNum_Call) Run(run func(ctx context.Context, asNum int64)) *MockNetworkRepository_GetASNRecordsByNum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockNetworkRepository_GetASNRecordsByNum_Call) Return(_a0 []domain.NetworkASN, _a1 error) *MockNetworkRepository_GetASNRecordsByNum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkRepository_GetASNRecordsByNum_Call) RunAndReturn(run func(context.Context, int64) ([]domain.NetworkASN, error)) *MockNetworkRepository_GetASNRecordsByNum_Call {
	_c.Call.Return(run)
	return _c
}

// GetLocationRecord provides a mock function with given fields: ctx, ipAddr
func (_m *MockNetworkRepository) GetLocationRecord(ctx context.Context, ipAddr netip.Addr) (domain.NetworkLocation, error) {
	ret := _m.Called(ctx, ipAddr)

	if len(ret) == 0 {
		panic("no return value specified for GetLocationRecord")
	}

	var r0 domain.NetworkLocation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) (domain.NetworkLocation, error)); ok {
		return rf(ctx, ipAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) domain.NetworkLocation); ok {
		r0 = rf(ctx, ipAddr)
	} else {
		r0 = ret.Get(0).(domain.NetworkLocation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, netip.Addr) error); ok {
		r1 = rf(ctx, ipAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworkRepository_GetLocationRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLocationRecord'
type MockNetworkRepository_GetLocationRecord_Call struct {
	*mock.Call
}

// GetLocationRecord is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAddr netip.Addr
func (_e *MockNetworkRepository_Expecter) GetLocationRecord(ctx interface{}, ipAddr interface{}) *MockNetworkRepository_GetLocationRecord_Call {
	return &MockNetworkRepository_GetLocationRecord_Call{Call: _e.mock.On("GetLocationRecord", ctx, ipAddr)}
}

func (_c *MockNetworkRepository_GetLocationRecord_Call) Run(run func(ctx context.Context, ipAddr netip.Addr)) *MockNetworkRepository_GetLocationRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(netip.Addr))
	})
	return _c
}

func (_c *MockNetworkRepository_GetLocationRecord_Call) Return(_a0 domain.NetworkLocation, _a1 error) *MockNetworkRepository_GetLocationRecord_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkRepository_GetLocationRecord_Call) RunAndReturn(run func(context.Context, netip.Addr) (domain.NetworkLocation, error)) *MockNetworkRepository_GetLocationRecord_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonIPHistory provides a mock function with given fields: ctx, sid64, limit
func (_m *MockNetworkRepository) GetPersonIPHistory(ctx context.Context, sid64 steamid.SteamID, limit uint64) (domain.PersonConnections, error) {
	ret := _m.Called(ctx, sid64, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonIPHistory")
	}

	var r0 domain.PersonConnections
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, uint64) (domain.PersonConnections, error)); ok {
		return rf(ctx, sid64, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, uint64) domain.PersonConnections); ok {
		r0 = rf(ctx, sid64, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.PersonConnections)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.SteamID, uint64) error); ok {
		r1 = rf(ctx, sid64, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworkRepository_GetPersonIPHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonIPHistory'
type MockNetworkRepository_GetPersonIPHistory_Call struct {
	*mock.Call
}

// GetPersonIPHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - sid64 steamid.SteamID
//   - limit uint64
func (_e *MockNetworkRepository_Expecter) GetPersonIPHistory(ctx interface{}, sid64 interface{}, limit interface{}) *MockNetworkRepository_GetPersonIPHistory_Call {
	return &MockNetworkRepository_GetPersonIPHistory_Call{Call: _e.mock.On("GetPersonIPHistory", ctx, sid64, limit)}
}

func (_c *MockNetworkRepository_GetPersonIPHistory_Call) Run(run func(ctx context.Context, sid64 steamid.SteamID, limit uint64)) *MockNetworkRepository_GetPersonIPHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID), args[2].(uint64))
	})
	return _c
}

func (_c *MockNetworkRepository_GetPersonIPHistory_Call) Return(_a0 domain.PersonConnections, _a1 error) *MockNetworkRepository_GetPersonIPHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkRepository_GetPersonIPHistory_Call) RunAndReturn(run func(context.Context, steamid.SteamID, uint64) (domain.PersonConnections, error)) *MockNetworkRepository_GetPersonIPHistory_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlayerMostRecentIP provides a mock function with given fields: ctx, steamID
func (_m *MockNetworkRepository) GetPlayerMostRecentIP(ctx context.Context, steamID steamid.SteamID) net.IP {
	ret := _m.Called(ctx, steamID)

	if len(ret) == 0 {
		panic("no return value specified for GetPlayerMostRecentIP")
	}

	var r0 net.IP
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID) net.IP); ok {
		r0 = rf(ctx, steamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.IP)
		}
	}

	return r0
}

// MockNetworkRepository_GetPlayerMostRecentIP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlayerMostRecentIP'
type MockNetworkRepository_GetPlayerMostRecentIP_Call struct {
	*mock.Call
}

// GetPlayerMostRecentIP is a helper method to define mock.On call
//   - ctx context.Context
//   - steamID steamid.SteamID
func (_e *MockNetworkRepository_Expecter) GetPlayerMostRecentIP(ctx interface{}, steamID interface{}) *MockNetworkRepository_GetPlayerMostRecentIP_Call {
	return &MockNetworkRepository_GetPlayerMostRecentIP_Call{Call: _e.mock.On("GetPlayerMostRecentIP", ctx, steamID)}
}

func (_c *MockNetworkRepository_GetPlayerMostRecentIP_Call) Run(run func(ctx context.Context, steamID steamid.SteamID)) *MockNetworkRepository_GetPlayerMostRecentIP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID))
	})
	return _c
}

func (_c *MockNetworkRepository_GetPlayerMostRecentIP_Call) Return(_a0 net.IP) *MockNetworkRepository_GetPlayerMostRecentIP_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkRepository_GetPlayerMostRecentIP_Call) RunAndReturn(run func(context.Context, steamid.SteamID) net.IP) *MockNetworkRepository_GetPlayerMostRecentIP_Call {
	_c.Call.Return(run)
	return _c
}

// GetProxyRecord provides a mock function with given fields: ctx, ipAddr
func (_m *MockNetworkRepository) GetProxyRecord(ctx context.Context, ipAddr netip.Addr) (domain.NetworkProxy, error) {
	ret := _m.Called(ctx, ipAddr)

	if len(ret) == 0 {
		panic("no return value specified for GetProxyRecord")
	}

	var r0 domain.NetworkProxy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) (domain.NetworkProxy, error)); ok {
		return rf(ctx, ipAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) domain.NetworkProxy); ok {
		r0 = rf(ctx, ipAddr)
	} else {
		r0 = ret.Get(0).(domain.NetworkProxy)
	}

	if rf, ok := ret.Get(1).(func(context.Context, netip.Addr) error); ok {
		r1 = rf(ctx, ipAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworkRepository_GetProxyRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProxyRecord'
type MockNetworkRepository_GetProxyRecord_Call struct {
	*mock.Call
}

// GetProxyRecord is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAddr netip.Addr
func (_e *MockNetworkRepository_Expecter) GetProxyRecord(ctx interface{}, ipAddr interface{}) *MockNetworkRepository_GetProxyRecord_Call {
	return &MockNetworkRepository_GetProxyRecord_Call{Call: _e.mock.On("GetProxyRecord", ctx, ipAddr)}
}

func (_c *MockNetworkRepository_GetProxyRecord_Call) Run(run func(ctx context.Context, ipAddr netip.Addr)) *MockNetworkRepository_GetProxyRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(netip.Addr))
	})
	return _c
}

func (_c *MockNetworkRepository_GetProxyRecord_Call) Return(_a0 domain.NetworkProxy, _a1 error) *MockNetworkRepository_GetProxyRecord_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkRepository_GetProxyRecord_Call) RunAndReturn(run func(context.Context, netip.Addr) (domain.NetworkProxy, error)) *MockNetworkRepository_GetProxyRecord_Call {
	_c.Call.Return(run)
	return _c
}

// InsertBlockListData provides a mock function with given fields: ctx, blockListData
func (_m *MockNetworkRepository) InsertBlockListData(ctx context.Context, blockListData *ip2location.BlockListData) error {
	ret := _m.Called(ctx, blockListData)

	if len(ret) == 0 {
		panic("no return value specified for InsertBlockListData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ip2location.BlockListData) error); ok {
		r0 = rf(ctx, blockListData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNetworkRepository_InsertBlockListData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertBlockListData'
type MockNetworkRepository_InsertBlockListData_Call struct {
	*mock.Call
}

// InsertBlockListData is a helper method to define mock.On call
//   - ctx context.Context
//   - blockListData *ip2location.BlockListData
func (_e *MockNetworkRepository_Expecter) InsertBlockListData(ctx interface{}, blockListData interface{}) *MockNetworkRepository_InsertBlockListData_Call {
	return &MockNetworkRepository_InsertBlockListData_Call{Call: _e.mock.On("InsertBlockListData", ctx, blockListData)}
}

func (_c *MockNetworkRepository_InsertBlockListData_Call) Run(run func(ctx context.Context, blockListData *ip2location.BlockListData)) *MockNetworkRepository_InsertBlockListData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*ip2location.BlockListData))
	})
	return _c
}

func (_c *MockNetworkRepository_InsertBlockListData_Call) Return(_a0 error) *MockNetworkRepository_InsertBlockListData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkRepository_InsertBlockListData_Call) RunAndReturn(run func(context.Context, *ip2location.BlockListData) error) *MockNetworkRepository_InsertBlockListData_Call {
	_c.Call.Return(run)
	return _c
}

// QueryConnections provides a mock function with given fields: ctx, opts
func (_m *MockNetworkRepository) QueryConnections(ctx context.Context, opts domain.ConnectionHistoryQuery) ([]domain.PersonConnection, int64, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for QueryConnections")
	}

	var r0 []domain.PersonConnection
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ConnectionHistoryQuery) ([]domain.PersonConnection, int64, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ConnectionHistoryQuery) []domain.PersonConnection); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PersonConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ConnectionHistoryQuery) int64); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.ConnectionHistoryQuery) error); ok {
		r2 = rf(ctx, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockNetworkRepository_QueryConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryConnections'
type MockNetworkRepository_QueryConnections_Call struct {
	*mock.Call
}

// QueryConnections is a helper method to define mock.On call
//   - ctx context.Context
//   - opts domain.ConnectionHistoryQuery
func (_e *MockNetworkRepository_Expecter) QueryConnections(ctx interface{}, opts interface{}) *MockNetworkRepository_QueryConnections_Call {
	return &MockNetworkRepository_QueryConnections_Call{Call: _e.mock.On("QueryConnections", ctx, opts)}
}

func (_c *MockNetworkRepository_QueryConnections_Call) Run(run func(ctx context.Context, opts domain.ConnectionHistoryQuery)) *MockNetworkRepository_QueryConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ConnectionHistoryQuery))
	})
	return _c
}

func (_c *MockNetworkRepository_QueryConnections_Call) Return(_a0 []domain.PersonConnection, _a1 int64, _a2 error) *MockNetworkRepository_QueryConnections_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockNetworkRepository_QueryConnections_Call) RunAndReturn(run func(context.Context, domain.ConnectionHistoryQuery) ([]domain.PersonConnection, int64, error)) *MockNetworkRepository_QueryConnections_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNetworkRepository creates a new instance of MockNetworkRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNetworkRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNetworkRepository {
	mock := &MockNetworkRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
