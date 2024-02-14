// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	ip2location "github.com/leighmacdonald/gbans/pkg/ip2location"

	mock "github.com/stretchr/testify/mock"

	net "net"

	steamid "github.com/leighmacdonald/steamid/v3/steamid"

	zap "go.uber.org/zap"
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

// GetASNRecordByIP provides a mock function with given fields: ctx, ipAddr, asnRecord
func (_m *MockNetworkUsecase) GetASNRecordByIP(ctx context.Context, ipAddr net.IP, asnRecord *ip2location.ASNRecord) error {
	ret := _m.Called(ctx, ipAddr, asnRecord)

	if len(ret) == 0 {
		panic("no return value specified for GetASNRecordByIP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, net.IP, *ip2location.ASNRecord) error); ok {
		r0 = rf(ctx, ipAddr, asnRecord)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNetworkUsecase_GetASNRecordByIP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetASNRecordByIP'
type MockNetworkUsecase_GetASNRecordByIP_Call struct {
	*mock.Call
}

// GetASNRecordByIP is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAddr net.IP
//   - asnRecord *ip2location.ASNRecord
func (_e *MockNetworkUsecase_Expecter) GetASNRecordByIP(ctx interface{}, ipAddr interface{}, asnRecord interface{}) *MockNetworkUsecase_GetASNRecordByIP_Call {
	return &MockNetworkUsecase_GetASNRecordByIP_Call{Call: _e.mock.On("GetASNRecordByIP", ctx, ipAddr, asnRecord)}
}

func (_c *MockNetworkUsecase_GetASNRecordByIP_Call) Run(run func(ctx context.Context, ipAddr net.IP, asnRecord *ip2location.ASNRecord)) *MockNetworkUsecase_GetASNRecordByIP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(net.IP), args[2].(*ip2location.ASNRecord))
	})
	return _c
}

func (_c *MockNetworkUsecase_GetASNRecordByIP_Call) Return(_a0 error) *MockNetworkUsecase_GetASNRecordByIP_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_GetASNRecordByIP_Call) RunAndReturn(run func(context.Context, net.IP, *ip2location.ASNRecord) error) *MockNetworkUsecase_GetASNRecordByIP_Call {
	_c.Call.Return(run)
	return _c
}

// GetASNRecordsByNum provides a mock function with given fields: ctx, asNum
func (_m *MockNetworkUsecase) GetASNRecordsByNum(ctx context.Context, asNum int64) (ip2location.ASNRecords, error) {
	ret := _m.Called(ctx, asNum)

	if len(ret) == 0 {
		panic("no return value specified for GetASNRecordsByNum")
	}

	var r0 ip2location.ASNRecords
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (ip2location.ASNRecords, error)); ok {
		return rf(ctx, asNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) ip2location.ASNRecords); ok {
		r0 = rf(ctx, asNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ip2location.ASNRecords)
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

func (_c *MockNetworkUsecase_GetASNRecordsByNum_Call) Return(_a0 ip2location.ASNRecords, _a1 error) *MockNetworkUsecase_GetASNRecordsByNum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkUsecase_GetASNRecordsByNum_Call) RunAndReturn(run func(context.Context, int64) (ip2location.ASNRecords, error)) *MockNetworkUsecase_GetASNRecordsByNum_Call {
	_c.Call.Return(run)
	return _c
}

// GetLocationRecord provides a mock function with given fields: ctx, ipAddr, record
func (_m *MockNetworkUsecase) GetLocationRecord(ctx context.Context, ipAddr net.IP, record *ip2location.LocationRecord) error {
	ret := _m.Called(ctx, ipAddr, record)

	if len(ret) == 0 {
		panic("no return value specified for GetLocationRecord")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, net.IP, *ip2location.LocationRecord) error); ok {
		r0 = rf(ctx, ipAddr, record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNetworkUsecase_GetLocationRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLocationRecord'
type MockNetworkUsecase_GetLocationRecord_Call struct {
	*mock.Call
}

// GetLocationRecord is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAddr net.IP
//   - record *ip2location.LocationRecord
func (_e *MockNetworkUsecase_Expecter) GetLocationRecord(ctx interface{}, ipAddr interface{}, record interface{}) *MockNetworkUsecase_GetLocationRecord_Call {
	return &MockNetworkUsecase_GetLocationRecord_Call{Call: _e.mock.On("GetLocationRecord", ctx, ipAddr, record)}
}

func (_c *MockNetworkUsecase_GetLocationRecord_Call) Run(run func(ctx context.Context, ipAddr net.IP, record *ip2location.LocationRecord)) *MockNetworkUsecase_GetLocationRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(net.IP), args[2].(*ip2location.LocationRecord))
	})
	return _c
}

func (_c *MockNetworkUsecase_GetLocationRecord_Call) Return(_a0 error) *MockNetworkUsecase_GetLocationRecord_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_GetLocationRecord_Call) RunAndReturn(run func(context.Context, net.IP, *ip2location.LocationRecord) error) *MockNetworkUsecase_GetLocationRecord_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonIPHistory provides a mock function with given fields: ctx, sid64, limit
func (_m *MockNetworkUsecase) GetPersonIPHistory(ctx context.Context, sid64 steamid.SID64, limit uint64) (domain.PersonConnections, error) {
	ret := _m.Called(ctx, sid64, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonIPHistory")
	}

	var r0 domain.PersonConnections
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SID64, uint64) (domain.PersonConnections, error)); ok {
		return rf(ctx, sid64, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SID64, uint64) domain.PersonConnections); ok {
		r0 = rf(ctx, sid64, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.PersonConnections)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.SID64, uint64) error); ok {
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
//   - sid64 steamid.SID64
//   - limit uint64
func (_e *MockNetworkUsecase_Expecter) GetPersonIPHistory(ctx interface{}, sid64 interface{}, limit interface{}) *MockNetworkUsecase_GetPersonIPHistory_Call {
	return &MockNetworkUsecase_GetPersonIPHistory_Call{Call: _e.mock.On("GetPersonIPHistory", ctx, sid64, limit)}
}

func (_c *MockNetworkUsecase_GetPersonIPHistory_Call) Run(run func(ctx context.Context, sid64 steamid.SID64, limit uint64)) *MockNetworkUsecase_GetPersonIPHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SID64), args[2].(uint64))
	})
	return _c
}

func (_c *MockNetworkUsecase_GetPersonIPHistory_Call) Return(_a0 domain.PersonConnections, _a1 error) *MockNetworkUsecase_GetPersonIPHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkUsecase_GetPersonIPHistory_Call) RunAndReturn(run func(context.Context, steamid.SID64, uint64) (domain.PersonConnections, error)) *MockNetworkUsecase_GetPersonIPHistory_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlayerMostRecentIP provides a mock function with given fields: ctx, steamID
func (_m *MockNetworkUsecase) GetPlayerMostRecentIP(ctx context.Context, steamID steamid.SID64) net.IP {
	ret := _m.Called(ctx, steamID)

	if len(ret) == 0 {
		panic("no return value specified for GetPlayerMostRecentIP")
	}

	var r0 net.IP
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SID64) net.IP); ok {
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
//   - steamID steamid.SID64
func (_e *MockNetworkUsecase_Expecter) GetPlayerMostRecentIP(ctx interface{}, steamID interface{}) *MockNetworkUsecase_GetPlayerMostRecentIP_Call {
	return &MockNetworkUsecase_GetPlayerMostRecentIP_Call{Call: _e.mock.On("GetPlayerMostRecentIP", ctx, steamID)}
}

func (_c *MockNetworkUsecase_GetPlayerMostRecentIP_Call) Run(run func(ctx context.Context, steamID steamid.SID64)) *MockNetworkUsecase_GetPlayerMostRecentIP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SID64))
	})
	return _c
}

func (_c *MockNetworkUsecase_GetPlayerMostRecentIP_Call) Return(_a0 net.IP) *MockNetworkUsecase_GetPlayerMostRecentIP_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_GetPlayerMostRecentIP_Call) RunAndReturn(run func(context.Context, steamid.SID64) net.IP) *MockNetworkUsecase_GetPlayerMostRecentIP_Call {
	_c.Call.Return(run)
	return _c
}

// GetProxyRecord provides a mock function with given fields: ctx, ipAddr, proxyRecord
func (_m *MockNetworkUsecase) GetProxyRecord(ctx context.Context, ipAddr net.IP, proxyRecord *ip2location.ProxyRecord) error {
	ret := _m.Called(ctx, ipAddr, proxyRecord)

	if len(ret) == 0 {
		panic("no return value specified for GetProxyRecord")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, net.IP, *ip2location.ProxyRecord) error); ok {
		r0 = rf(ctx, ipAddr, proxyRecord)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNetworkUsecase_GetProxyRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProxyRecord'
type MockNetworkUsecase_GetProxyRecord_Call struct {
	*mock.Call
}

// GetProxyRecord is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAddr net.IP
//   - proxyRecord *ip2location.ProxyRecord
func (_e *MockNetworkUsecase_Expecter) GetProxyRecord(ctx interface{}, ipAddr interface{}, proxyRecord interface{}) *MockNetworkUsecase_GetProxyRecord_Call {
	return &MockNetworkUsecase_GetProxyRecord_Call{Call: _e.mock.On("GetProxyRecord", ctx, ipAddr, proxyRecord)}
}

func (_c *MockNetworkUsecase_GetProxyRecord_Call) Run(run func(ctx context.Context, ipAddr net.IP, proxyRecord *ip2location.ProxyRecord)) *MockNetworkUsecase_GetProxyRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(net.IP), args[2].(*ip2location.ProxyRecord))
	})
	return _c
}

func (_c *MockNetworkUsecase_GetProxyRecord_Call) Return(_a0 error) *MockNetworkUsecase_GetProxyRecord_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_GetProxyRecord_Call) RunAndReturn(run func(context.Context, net.IP, *ip2location.ProxyRecord) error) *MockNetworkUsecase_GetProxyRecord_Call {
	_c.Call.Return(run)
	return _c
}

// InsertBlockListData provides a mock function with given fields: ctx, log, blockListData
func (_m *MockNetworkUsecase) InsertBlockListData(ctx context.Context, log *zap.Logger, blockListData *ip2location.BlockListData) error {
	ret := _m.Called(ctx, log, blockListData)

	if len(ret) == 0 {
		panic("no return value specified for InsertBlockListData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *zap.Logger, *ip2location.BlockListData) error); ok {
		r0 = rf(ctx, log, blockListData)
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
//   - log *zap.Logger
//   - blockListData *ip2location.BlockListData
func (_e *MockNetworkUsecase_Expecter) InsertBlockListData(ctx interface{}, log interface{}, blockListData interface{}) *MockNetworkUsecase_InsertBlockListData_Call {
	return &MockNetworkUsecase_InsertBlockListData_Call{Call: _e.mock.On("InsertBlockListData", ctx, log, blockListData)}
}

func (_c *MockNetworkUsecase_InsertBlockListData_Call) Run(run func(ctx context.Context, log *zap.Logger, blockListData *ip2location.BlockListData)) *MockNetworkUsecase_InsertBlockListData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*zap.Logger), args[2].(*ip2location.BlockListData))
	})
	return _c
}

func (_c *MockNetworkUsecase_InsertBlockListData_Call) Return(_a0 error) *MockNetworkUsecase_InsertBlockListData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworkUsecase_InsertBlockListData_Call) RunAndReturn(run func(context.Context, *zap.Logger, *ip2location.BlockListData) error) *MockNetworkUsecase_InsertBlockListData_Call {
	_c.Call.Return(run)
	return _c
}

// IsMatch provides a mock function with given fields: addr
func (_m *MockNetworkUsecase) IsMatch(addr net.IP) (string, bool) {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for IsMatch")
	}

	var r0 string
	var r1 bool
	if rf, ok := ret.Get(0).(func(net.IP) (string, bool)); ok {
		return rf(addr)
	}
	if rf, ok := ret.Get(0).(func(net.IP) string); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(net.IP) bool); ok {
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
//   - addr net.IP
func (_e *MockNetworkUsecase_Expecter) IsMatch(addr interface{}) *MockNetworkUsecase_IsMatch_Call {
	return &MockNetworkUsecase_IsMatch_Call{Call: _e.mock.On("IsMatch", addr)}
}

func (_c *MockNetworkUsecase_IsMatch_Call) Run(run func(addr net.IP)) *MockNetworkUsecase_IsMatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.IP))
	})
	return _c
}

func (_c *MockNetworkUsecase_IsMatch_Call) Return(_a0 string, _a1 bool) *MockNetworkUsecase_IsMatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworkUsecase_IsMatch_Call) RunAndReturn(run func(net.IP) (string, bool)) *MockNetworkUsecase_IsMatch_Call {
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
func (_m *MockNetworkUsecase) QueryConnectionHistory(ctx context.Context, opts domain.ConnectionHistoryQueryFilter) ([]domain.PersonConnection, int64, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for QueryConnectionHistory")
	}

	var r0 []domain.PersonConnection
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ConnectionHistoryQueryFilter) ([]domain.PersonConnection, int64, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ConnectionHistoryQueryFilter) []domain.PersonConnection); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PersonConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ConnectionHistoryQueryFilter) int64); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.ConnectionHistoryQueryFilter) error); ok {
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
//   - opts domain.ConnectionHistoryQueryFilter
func (_e *MockNetworkUsecase_Expecter) QueryConnectionHistory(ctx interface{}, opts interface{}) *MockNetworkUsecase_QueryConnectionHistory_Call {
	return &MockNetworkUsecase_QueryConnectionHistory_Call{Call: _e.mock.On("QueryConnectionHistory", ctx, opts)}
}

func (_c *MockNetworkUsecase_QueryConnectionHistory_Call) Run(run func(ctx context.Context, opts domain.ConnectionHistoryQueryFilter)) *MockNetworkUsecase_QueryConnectionHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ConnectionHistoryQueryFilter))
	})
	return _c
}

func (_c *MockNetworkUsecase_QueryConnectionHistory_Call) Return(_a0 []domain.PersonConnection, _a1 int64, _a2 error) *MockNetworkUsecase_QueryConnectionHistory_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockNetworkUsecase_QueryConnectionHistory_Call) RunAndReturn(run func(context.Context, domain.ConnectionHistoryQueryFilter) ([]domain.PersonConnection, int64, error)) *MockNetworkUsecase_QueryConnectionHistory_Call {
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
