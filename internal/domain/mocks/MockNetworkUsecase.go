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

// MockNetworkUsecase is an autogenerated mock type for the NetworkUsecase type
type MockNetworkUsecase struct {
	mock.Mock
}

type MockNetworkUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNetworkUsecase) EXPECT() *MockNetworkUsecase_Expecter {
	return &MockNetworkUsecase_Expecter{mock: &_m.Mock}
}

// AddConnectionHistory provides a mock function with given fields: ctx, conn
func (_m *MockNetworkUsecase) AddConnectionHistory(ctx context.Context, conn *domain.PersonConnection) error {
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

// MockNetworkUsecase_AddConnectionHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddConnectionHistory'
type MockNetworkUsecase_AddConnectionHistory_Call struct {
	*mock.Call
}

// AddConnectionHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - conn *domain.PersonConnection
func (_e *MockNetworkUsecase_Expecter) AddConnectionHistory(ctx interface{}, conn interface{}) *MockNetworkUsecase_AddConnectionHistory_Call {
	return &MockNetworkUsecase_AddConnectionHistory_Call{Call: _e.mock.On("AddConnectionHistory", ctx, conn)}
}

func (_c *MockNetworkUsecase_AddConnectionHistory_Call) Run(run func(ctx context.Context, conn *domain.PersonConnection)) *MockNetworkUsecase_AddConnectionHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.PersonConnection))
	})
	return _c
}

func (_c *MockNetworkUsecase_AddConnectionHistory_Call) Return(_a0 error) *MockNetworkUsecase_AddConnectionHistory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_AddConnectionHistory_Call) RunAndReturn(run func(context.Context, *domain.PersonConnection) error) *MockNetworkUsecase_AddConnectionHistory_Call {
	_c.Call.Return(run)
	return _c
}

// AddRemoteSource provides a mock function with given fields: ctx, name, url
func (_m *MockNetworkUsecase) AddRemoteSource(ctx context.Context, name string, url string) (int64, error) {
	ret := _m.Called(ctx, name, url)

	if len(ret) == 0 {
		panic("no return value specified for AddRemoteSource")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (int64, error)); ok {
		return rf(ctx, name, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, name, url)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworkUsecase_AddRemoteSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRemoteSource'
type MockNetworkUsecase_AddRemoteSource_Call struct {
	*mock.Call
}

// AddRemoteSource is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - url string
func (_e *MockNetworkUsecase_Expecter) AddRemoteSource(ctx interface{}, name interface{}, url interface{}) *MockNetworkUsecase_AddRemoteSource_Call {
	return &MockNetworkUsecase_AddRemoteSource_Call{Call: _e.mock.On("AddRemoteSource", ctx, name, url)}
}

func (_c *MockNetworkUsecase_AddRemoteSource_Call) Run(run func(ctx context.Context, name string, url string)) *MockNetworkUsecase_AddRemoteSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockNetworkUsecase_AddRemoteSource_Call) Return(_a0 int64, _a1 error) *MockNetworkUsecase_AddRemoteSource_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkUsecase_AddRemoteSource_Call) RunAndReturn(run func(context.Context, string, string) (int64, error)) *MockNetworkUsecase_AddRemoteSource_Call {
	_c.Call.Return(run)
	return _c
}

// AddWhitelist provides a mock function with given fields: id, network
func (_m *MockNetworkUsecase) AddWhitelist(id int, network *net.IPNet) {
	_m.Called(id, network)
}

// MockNetworkUsecase_AddWhitelist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddWhitelist'
type MockNetworkUsecase_AddWhitelist_Call struct {
	*mock.Call
}

// AddWhitelist is a helper method to define mock.On call
//   - id int
//   - network *net.IPNet
func (_e *MockNetworkUsecase_Expecter) AddWhitelist(id interface{}, network interface{}) *MockNetworkUsecase_AddWhitelist_Call {
	return &MockNetworkUsecase_AddWhitelist_Call{Call: _e.mock.On("AddWhitelist", id, network)}
}

func (_c *MockNetworkUsecase_AddWhitelist_Call) Run(run func(id int, network *net.IPNet)) *MockNetworkUsecase_AddWhitelist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(*net.IPNet))
	})
	return _c
}

func (_c *MockNetworkUsecase_AddWhitelist_Call) Return() *MockNetworkUsecase_AddWhitelist_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNetworkUsecase_AddWhitelist_Call) RunAndReturn(run func(int, *net.IPNet)) *MockNetworkUsecase_AddWhitelist_Call {
	_c.Call.Return(run)
	return _c
}

// GetASNRecordsByNum provides a mock function with given fields: ctx, asNum
func (_m *MockNetworkUsecase) GetASNRecordsByNum(ctx context.Context, asNum int64) ([]domain.NetworkASN, error) {
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

// MockNetworkUsecase_GetASNRecordsByNum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetASNRecordsByNum'
type MockNetworkUsecase_GetASNRecordsByNum_Call struct {
	*mock.Call
}

// GetASNRecordsByNum is a helper method to define mock.On call
//   - ctx context.Context
//   - asNum int64
func (_e *MockNetworkUsecase_Expecter) GetASNRecordsByNum(ctx interface{}, asNum interface{}) *MockNetworkUsecase_GetASNRecordsByNum_Call {
	return &MockNetworkUsecase_GetASNRecordsByNum_Call{Call: _e.mock.On("GetASNRecordsByNum", ctx, asNum)}
}

func (_c *MockNetworkUsecase_GetASNRecordsByNum_Call) Run(run func(ctx context.Context, asNum int64)) *MockNetworkUsecase_GetASNRecordsByNum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockNetworkUsecase_GetASNRecordsByNum_Call) Return(_a0 []domain.NetworkASN, _a1 error) *MockNetworkUsecase_GetASNRecordsByNum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkUsecase_GetASNRecordsByNum_Call) RunAndReturn(run func(context.Context, int64) ([]domain.NetworkASN, error)) *MockNetworkUsecase_GetASNRecordsByNum_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonIPHistory provides a mock function with given fields: ctx, sid64, limit
func (_m *MockNetworkUsecase) GetPersonIPHistory(ctx context.Context, sid64 steamid.SteamID, limit uint64) (domain.PersonConnections, error) {
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

// MockNetworkUsecase_GetPersonIPHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonIPHistory'
type MockNetworkUsecase_GetPersonIPHistory_Call struct {
	*mock.Call
}

// GetPersonIPHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - sid64 steamid.SteamID
//   - limit uint64
func (_e *MockNetworkUsecase_Expecter) GetPersonIPHistory(ctx interface{}, sid64 interface{}, limit interface{}) *MockNetworkUsecase_GetPersonIPHistory_Call {
	return &MockNetworkUsecase_GetPersonIPHistory_Call{Call: _e.mock.On("GetPersonIPHistory", ctx, sid64, limit)}
}

func (_c *MockNetworkUsecase_GetPersonIPHistory_Call) Run(run func(ctx context.Context, sid64 steamid.SteamID, limit uint64)) *MockNetworkUsecase_GetPersonIPHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID), args[2].(uint64))
	})
	return _c
}

func (_c *MockNetworkUsecase_GetPersonIPHistory_Call) Return(_a0 domain.PersonConnections, _a1 error) *MockNetworkUsecase_GetPersonIPHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkUsecase_GetPersonIPHistory_Call) RunAndReturn(run func(context.Context, steamid.SteamID, uint64) (domain.PersonConnections, error)) *MockNetworkUsecase_GetPersonIPHistory_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlayerMostRecentIP provides a mock function with given fields: ctx, steamID
func (_m *MockNetworkUsecase) GetPlayerMostRecentIP(ctx context.Context, steamID steamid.SteamID) net.IP {
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

// MockNetworkUsecase_GetPlayerMostRecentIP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlayerMostRecentIP'
type MockNetworkUsecase_GetPlayerMostRecentIP_Call struct {
	*mock.Call
}

// GetPlayerMostRecentIP is a helper method to define mock.On call
//   - ctx context.Context
//   - steamID steamid.SteamID
func (_e *MockNetworkUsecase_Expecter) GetPlayerMostRecentIP(ctx interface{}, steamID interface{}) *MockNetworkUsecase_GetPlayerMostRecentIP_Call {
	return &MockNetworkUsecase_GetPlayerMostRecentIP_Call{Call: _e.mock.On("GetPlayerMostRecentIP", ctx, steamID)}
}

func (_c *MockNetworkUsecase_GetPlayerMostRecentIP_Call) Run(run func(ctx context.Context, steamID steamid.SteamID)) *MockNetworkUsecase_GetPlayerMostRecentIP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID))
	})
	return _c
}

func (_c *MockNetworkUsecase_GetPlayerMostRecentIP_Call) Return(_a0 net.IP) *MockNetworkUsecase_GetPlayerMostRecentIP_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_GetPlayerMostRecentIP_Call) RunAndReturn(run func(context.Context, steamid.SteamID) net.IP) *MockNetworkUsecase_GetPlayerMostRecentIP_Call {
	_c.Call.Return(run)
	return _c
}

// InsertBlockListData provides a mock function with given fields: ctx, blockListData
func (_m *MockNetworkUsecase) InsertBlockListData(ctx context.Context, blockListData *ip2location.BlockListData) error {
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

// MockNetworkUsecase_InsertBlockListData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertBlockListData'
type MockNetworkUsecase_InsertBlockListData_Call struct {
	*mock.Call
}

// InsertBlockListData is a helper method to define mock.On call
//   - ctx context.Context
//   - blockListData *ip2location.BlockListData
func (_e *MockNetworkUsecase_Expecter) InsertBlockListData(ctx interface{}, blockListData interface{}) *MockNetworkUsecase_InsertBlockListData_Call {
	return &MockNetworkUsecase_InsertBlockListData_Call{Call: _e.mock.On("InsertBlockListData", ctx, blockListData)}
}

func (_c *MockNetworkUsecase_InsertBlockListData_Call) Run(run func(ctx context.Context, blockListData *ip2location.BlockListData)) *MockNetworkUsecase_InsertBlockListData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*ip2location.BlockListData))
	})
	return _c
}

func (_c *MockNetworkUsecase_InsertBlockListData_Call) Return(_a0 error) *MockNetworkUsecase_InsertBlockListData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_InsertBlockListData_Call) RunAndReturn(run func(context.Context, *ip2location.BlockListData) error) *MockNetworkUsecase_InsertBlockListData_Call {
	_c.Call.Return(run)
	return _c
}

// IsMatch provides a mock function with given fields: addr
func (_m *MockNetworkUsecase) IsMatch(addr netip.Addr) (string, bool) {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for IsMatch")
	}

	var r0 string
	var r1 bool
	if rf, ok := ret.Get(0).(func(netip.Addr) (string, bool)); ok {
		return rf(addr)
	}
	if rf, ok := ret.Get(0).(func(netip.Addr) string); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(netip.Addr) bool); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockNetworkUsecase_IsMatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsMatch'
type MockNetworkUsecase_IsMatch_Call struct {
	*mock.Call
}

// IsMatch is a helper method to define mock.On call
//   - addr netip.Addr
func (_e *MockNetworkUsecase_Expecter) IsMatch(addr interface{}) *MockNetworkUsecase_IsMatch_Call {
	return &MockNetworkUsecase_IsMatch_Call{Call: _e.mock.On("IsMatch", addr)}
}

func (_c *MockNetworkUsecase_IsMatch_Call) Run(run func(addr netip.Addr)) *MockNetworkUsecase_IsMatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(netip.Addr))
	})
	return _c
}

func (_c *MockNetworkUsecase_IsMatch_Call) Return(_a0 string, _a1 bool) *MockNetworkUsecase_IsMatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkUsecase_IsMatch_Call) RunAndReturn(run func(netip.Addr) (string, bool)) *MockNetworkUsecase_IsMatch_Call {
	_c.Call.Return(run)
	return _c
}

// LoadNetBlocks provides a mock function with given fields: ctx
func (_m *MockNetworkUsecase) LoadNetBlocks(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LoadNetBlocks")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNetworkUsecase_LoadNetBlocks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadNetBlocks'
type MockNetworkUsecase_LoadNetBlocks_Call struct {
	*mock.Call
}

// LoadNetBlocks is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockNetworkUsecase_Expecter) LoadNetBlocks(ctx interface{}) *MockNetworkUsecase_LoadNetBlocks_Call {
	return &MockNetworkUsecase_LoadNetBlocks_Call{Call: _e.mock.On("LoadNetBlocks", ctx)}
}

func (_c *MockNetworkUsecase_LoadNetBlocks_Call) Run(run func(ctx context.Context)) *MockNetworkUsecase_LoadNetBlocks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockNetworkUsecase_LoadNetBlocks_Call) Return(_a0 error) *MockNetworkUsecase_LoadNetBlocks_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_LoadNetBlocks_Call) RunAndReturn(run func(context.Context) error) *MockNetworkUsecase_LoadNetBlocks_Call {
	_c.Call.Return(run)
	return _c
}

// QueryConnectionHistory provides a mock function with given fields: ctx, opts
func (_m *MockNetworkUsecase) QueryConnectionHistory(ctx context.Context, opts domain.ConnectionHistoryQuery) ([]domain.PersonConnection, int64, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for QueryConnectionHistory")
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

// MockNetworkUsecase_QueryConnectionHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryConnectionHistory'
type MockNetworkUsecase_QueryConnectionHistory_Call struct {
	*mock.Call
}

// QueryConnectionHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - opts domain.ConnectionHistoryQuery
func (_e *MockNetworkUsecase_Expecter) QueryConnectionHistory(ctx interface{}, opts interface{}) *MockNetworkUsecase_QueryConnectionHistory_Call {
	return &MockNetworkUsecase_QueryConnectionHistory_Call{Call: _e.mock.On("QueryConnectionHistory", ctx, opts)}
}

func (_c *MockNetworkUsecase_QueryConnectionHistory_Call) Run(run func(ctx context.Context, opts domain.ConnectionHistoryQuery)) *MockNetworkUsecase_QueryConnectionHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ConnectionHistoryQuery))
	})
	return _c
}

func (_c *MockNetworkUsecase_QueryConnectionHistory_Call) Return(_a0 []domain.PersonConnection, _a1 int64, _a2 error) *MockNetworkUsecase_QueryConnectionHistory_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockNetworkUsecase_QueryConnectionHistory_Call) RunAndReturn(run func(context.Context, domain.ConnectionHistoryQuery) ([]domain.PersonConnection, int64, error)) *MockNetworkUsecase_QueryConnectionHistory_Call {
	_c.Call.Return(run)
	return _c
}

// QueryNetwork provides a mock function with given fields: ctx, ip
func (_m *MockNetworkUsecase) QueryNetwork(ctx context.Context, ip netip.Addr) (domain.NetworkDetails, error) {
	ret := _m.Called(ctx, ip)

	if len(ret) == 0 {
		panic("no return value specified for QueryNetwork")
	}

	var r0 domain.NetworkDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) (domain.NetworkDetails, error)); ok {
		return rf(ctx, ip)
	}
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) domain.NetworkDetails); ok {
		r0 = rf(ctx, ip)
	} else {
		r0 = ret.Get(0).(domain.NetworkDetails)
	}

	if rf, ok := ret.Get(1).(func(context.Context, netip.Addr) error); ok {
		r1 = rf(ctx, ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworkUsecase_QueryNetwork_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryNetwork'
type MockNetworkUsecase_QueryNetwork_Call struct {
	*mock.Call
}

// QueryNetwork is a helper method to define mock.On call
//   - ctx context.Context
//   - ip netip.Addr
func (_e *MockNetworkUsecase_Expecter) QueryNetwork(ctx interface{}, ip interface{}) *MockNetworkUsecase_QueryNetwork_Call {
	return &MockNetworkUsecase_QueryNetwork_Call{Call: _e.mock.On("QueryNetwork", ctx, ip)}
}

func (_c *MockNetworkUsecase_QueryNetwork_Call) Run(run func(ctx context.Context, ip netip.Addr)) *MockNetworkUsecase_QueryNetwork_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(netip.Addr))
	})
	return _c
}

func (_c *MockNetworkUsecase_QueryNetwork_Call) Return(_a0 domain.NetworkDetails, _a1 error) *MockNetworkUsecase_QueryNetwork_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkUsecase_QueryNetwork_Call) RunAndReturn(run func(context.Context, netip.Addr) (domain.NetworkDetails, error)) *MockNetworkUsecase_QueryNetwork_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveWhitelist provides a mock function with given fields: id
func (_m *MockNetworkUsecase) RemoveWhitelist(id int) {
	_m.Called(id)
}

// MockNetworkUsecase_RemoveWhitelist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveWhitelist'
type MockNetworkUsecase_RemoveWhitelist_Call struct {
	*mock.Call
}

// RemoveWhitelist is a helper method to define mock.On call
//   - id int
func (_e *MockNetworkUsecase_Expecter) RemoveWhitelist(id interface{}) *MockNetworkUsecase_RemoveWhitelist_Call {
	return &MockNetworkUsecase_RemoveWhitelist_Call{Call: _e.mock.On("RemoveWhitelist", id)}
}

func (_c *MockNetworkUsecase_RemoveWhitelist_Call) Run(run func(id int)) *MockNetworkUsecase_RemoveWhitelist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockNetworkUsecase_RemoveWhitelist_Call) Return() *MockNetworkUsecase_RemoveWhitelist_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNetworkUsecase_RemoveWhitelist_Call) RunAndReturn(run func(int)) *MockNetworkUsecase_RemoveWhitelist_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *MockNetworkUsecase) Start(ctx context.Context) {
	_m.Called(ctx)
}

// MockNetworkUsecase_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockNetworkUsecase_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockNetworkUsecase_Expecter) Start(ctx interface{}) *MockNetworkUsecase_Start_Call {
	return &MockNetworkUsecase_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *MockNetworkUsecase_Start_Call) Run(run func(ctx context.Context)) *MockNetworkUsecase_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockNetworkUsecase_Start_Call) Return() *MockNetworkUsecase_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNetworkUsecase_Start_Call) RunAndReturn(run func(context.Context)) *MockNetworkUsecase_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNetworkUsecase creates a new instance of MockNetworkUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNetworkUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNetworkUsecase {
	mock := &MockNetworkUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
