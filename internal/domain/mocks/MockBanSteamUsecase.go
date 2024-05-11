// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"

	netip "net/netip"

	steamid "github.com/leighmacdonald/steamid/v4/steamid"

	time "time"
)

// MockBanSteamUsecase is an autogenerated mock type for the BanSteamUsecase type
type MockBanSteamUsecase struct {
	mock.Mock
}

type MockBanSteamUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBanSteamUsecase) EXPECT() *MockBanSteamUsecase_Expecter {
	return &MockBanSteamUsecase_Expecter{mock: &_m.Mock}
}

// Ban provides a mock function with given fields: ctx, curUser, banSteam
func (_m *MockBanSteamUsecase) Ban(ctx context.Context, curUser domain.PersonInfo, banSteam *domain.BanSteam) error {
	ret := _m.Called(ctx, curUser, banSteam)

	if len(ret) == 0 {
		panic("no return value specified for Ban")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, *domain.BanSteam) error); ok {
		r0 = rf(ctx, curUser, banSteam)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanSteamUsecase_Ban_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ban'
type MockBanSteamUsecase_Ban_Call struct {
	*mock.Call
}

// Ban is a helper method to define mock.On call
//   - ctx context.Context
//   - curUser domain.PersonInfo
//   - banSteam *domain.BanSteam
func (_e *MockBanSteamUsecase_Expecter) Ban(ctx interface{}, curUser interface{}, banSteam interface{}) *MockBanSteamUsecase_Ban_Call {
	return &MockBanSteamUsecase_Ban_Call{Call: _e.mock.On("Ban", ctx, curUser, banSteam)}
}

func (_c *MockBanSteamUsecase_Ban_Call) Run(run func(ctx context.Context, curUser domain.PersonInfo, banSteam *domain.BanSteam)) *MockBanSteamUsecase_Ban_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PersonInfo), args[2].(*domain.BanSteam))
	})
	return _c
}

func (_c *MockBanSteamUsecase_Ban_Call) Return(_a0 error) *MockBanSteamUsecase_Ban_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanSteamUsecase_Ban_Call) RunAndReturn(run func(context.Context, domain.PersonInfo, *domain.BanSteam) error) *MockBanSteamUsecase_Ban_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, ban, hardDelete
func (_m *MockBanSteamUsecase) Delete(ctx context.Context, ban *domain.BanSteam, hardDelete bool) error {
	ret := _m.Called(ctx, ban, hardDelete)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BanSteam, bool) error); ok {
		r0 = rf(ctx, ban, hardDelete)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanSteamUsecase_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockBanSteamUsecase_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - ban *domain.BanSteam
//   - hardDelete bool
func (_e *MockBanSteamUsecase_Expecter) Delete(ctx interface{}, ban interface{}, hardDelete interface{}) *MockBanSteamUsecase_Delete_Call {
	return &MockBanSteamUsecase_Delete_Call{Call: _e.mock.On("Delete", ctx, ban, hardDelete)}
}

func (_c *MockBanSteamUsecase_Delete_Call) Run(run func(ctx context.Context, ban *domain.BanSteam, hardDelete bool)) *MockBanSteamUsecase_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.BanSteam), args[2].(bool))
	})
	return _c
}

func (_c *MockBanSteamUsecase_Delete_Call) Return(_a0 error) *MockBanSteamUsecase_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanSteamUsecase_Delete_Call) RunAndReturn(run func(context.Context, *domain.BanSteam, bool) error) *MockBanSteamUsecase_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Expired provides a mock function with given fields: ctx
func (_m *MockBanSteamUsecase) Expired(ctx context.Context) ([]domain.BanSteam, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Expired")
	}

	var r0 []domain.BanSteam
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.BanSteam, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.BanSteam); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BanSteam)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanSteamUsecase_Expired_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Expired'
type MockBanSteamUsecase_Expired_Call struct {
	*mock.Call
}

// Expired is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBanSteamUsecase_Expecter) Expired(ctx interface{}) *MockBanSteamUsecase_Expired_Call {
	return &MockBanSteamUsecase_Expired_Call{Call: _e.mock.On("Expired", ctx)}
}

func (_c *MockBanSteamUsecase_Expired_Call) Run(run func(ctx context.Context)) *MockBanSteamUsecase_Expired_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBanSteamUsecase_Expired_Call) Return(_a0 []domain.BanSteam, _a1 error) *MockBanSteamUsecase_Expired_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanSteamUsecase_Expired_Call) RunAndReturn(run func(context.Context) ([]domain.BanSteam, error)) *MockBanSteamUsecase_Expired_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, filter
func (_m *MockBanSteamUsecase) Get(ctx context.Context, filter domain.SteamBansQueryFilter) ([]domain.BannedSteamPerson, int64, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []domain.BannedSteamPerson
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.SteamBansQueryFilter) ([]domain.BannedSteamPerson, int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.SteamBansQueryFilter) []domain.BannedSteamPerson); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BannedSteamPerson)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.SteamBansQueryFilter) int64); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.SteamBansQueryFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBanSteamUsecase_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockBanSteamUsecase_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - filter domain.SteamBansQueryFilter
func (_e *MockBanSteamUsecase_Expecter) Get(ctx interface{}, filter interface{}) *MockBanSteamUsecase_Get_Call {
	return &MockBanSteamUsecase_Get_Call{Call: _e.mock.On("Get", ctx, filter)}
}

func (_c *MockBanSteamUsecase_Get_Call) Run(run func(ctx context.Context, filter domain.SteamBansQueryFilter)) *MockBanSteamUsecase_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.SteamBansQueryFilter))
	})
	return _c
}

func (_c *MockBanSteamUsecase_Get_Call) Return(_a0 []domain.BannedSteamPerson, _a1 int64, _a2 error) *MockBanSteamUsecase_Get_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockBanSteamUsecase_Get_Call) RunAndReturn(run func(context.Context, domain.SteamBansQueryFilter) ([]domain.BannedSteamPerson, int64, error)) *MockBanSteamUsecase_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByBanID provides a mock function with given fields: ctx, banID, deletedOk
func (_m *MockBanSteamUsecase) GetByBanID(ctx context.Context, banID int64, deletedOk bool) (domain.BannedSteamPerson, error) {
	ret := _m.Called(ctx, banID, deletedOk)

	if len(ret) == 0 {
		panic("no return value specified for GetByBanID")
	}

	var r0 domain.BannedSteamPerson
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, bool) (domain.BannedSteamPerson, error)); ok {
		return rf(ctx, banID, deletedOk)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, bool) domain.BannedSteamPerson); ok {
		r0 = rf(ctx, banID, deletedOk)
	} else {
		r0 = ret.Get(0).(domain.BannedSteamPerson)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, bool) error); ok {
		r1 = rf(ctx, banID, deletedOk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanSteamUsecase_GetByBanID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByBanID'
type MockBanSteamUsecase_GetByBanID_Call struct {
	*mock.Call
}

// GetByBanID is a helper method to define mock.On call
//   - ctx context.Context
//   - banID int64
//   - deletedOk bool
func (_e *MockBanSteamUsecase_Expecter) GetByBanID(ctx interface{}, banID interface{}, deletedOk interface{}) *MockBanSteamUsecase_GetByBanID_Call {
	return &MockBanSteamUsecase_GetByBanID_Call{Call: _e.mock.On("GetByBanID", ctx, banID, deletedOk)}
}

func (_c *MockBanSteamUsecase_GetByBanID_Call) Run(run func(ctx context.Context, banID int64, deletedOk bool)) *MockBanSteamUsecase_GetByBanID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(bool))
	})
	return _c
}

func (_c *MockBanSteamUsecase_GetByBanID_Call) Return(_a0 domain.BannedSteamPerson, _a1 error) *MockBanSteamUsecase_GetByBanID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanSteamUsecase_GetByBanID_Call) RunAndReturn(run func(context.Context, int64, bool) (domain.BannedSteamPerson, error)) *MockBanSteamUsecase_GetByBanID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByLastIP provides a mock function with given fields: ctx, lastIP, deletedOk
func (_m *MockBanSteamUsecase) GetByLastIP(ctx context.Context, lastIP netip.Addr, deletedOk bool) (domain.BannedSteamPerson, error) {
	ret := _m.Called(ctx, lastIP, deletedOk)

	if len(ret) == 0 {
		panic("no return value specified for GetByLastIP")
	}

	var r0 domain.BannedSteamPerson
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr, bool) (domain.BannedSteamPerson, error)); ok {
		return rf(ctx, lastIP, deletedOk)
	}
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr, bool) domain.BannedSteamPerson); ok {
		r0 = rf(ctx, lastIP, deletedOk)
	} else {
		r0 = ret.Get(0).(domain.BannedSteamPerson)
	}

	if rf, ok := ret.Get(1).(func(context.Context, netip.Addr, bool) error); ok {
		r1 = rf(ctx, lastIP, deletedOk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanSteamUsecase_GetByLastIP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByLastIP'
type MockBanSteamUsecase_GetByLastIP_Call struct {
	*mock.Call
}

// GetByLastIP is a helper method to define mock.On call
//   - ctx context.Context
//   - lastIP netip.Addr
//   - deletedOk bool
func (_e *MockBanSteamUsecase_Expecter) GetByLastIP(ctx interface{}, lastIP interface{}, deletedOk interface{}) *MockBanSteamUsecase_GetByLastIP_Call {
	return &MockBanSteamUsecase_GetByLastIP_Call{Call: _e.mock.On("GetByLastIP", ctx, lastIP, deletedOk)}
}

func (_c *MockBanSteamUsecase_GetByLastIP_Call) Run(run func(ctx context.Context, lastIP netip.Addr, deletedOk bool)) *MockBanSteamUsecase_GetByLastIP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(netip.Addr), args[2].(bool))
	})
	return _c
}

func (_c *MockBanSteamUsecase_GetByLastIP_Call) Return(_a0 domain.BannedSteamPerson, _a1 error) *MockBanSteamUsecase_GetByLastIP_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanSteamUsecase_GetByLastIP_Call) RunAndReturn(run func(context.Context, netip.Addr, bool) (domain.BannedSteamPerson, error)) *MockBanSteamUsecase_GetByLastIP_Call {
	_c.Call.Return(run)
	return _c
}

// GetBySteamID provides a mock function with given fields: ctx, sid64, deletedOk
func (_m *MockBanSteamUsecase) GetBySteamID(ctx context.Context, sid64 steamid.SteamID, deletedOk bool) (domain.BannedSteamPerson, error) {
	ret := _m.Called(ctx, sid64, deletedOk)

	if len(ret) == 0 {
		panic("no return value specified for GetBySteamID")
	}

	var r0 domain.BannedSteamPerson
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, bool) (domain.BannedSteamPerson, error)); ok {
		return rf(ctx, sid64, deletedOk)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, bool) domain.BannedSteamPerson); ok {
		r0 = rf(ctx, sid64, deletedOk)
	} else {
		r0 = ret.Get(0).(domain.BannedSteamPerson)
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.SteamID, bool) error); ok {
		r1 = rf(ctx, sid64, deletedOk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanSteamUsecase_GetBySteamID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBySteamID'
type MockBanSteamUsecase_GetBySteamID_Call struct {
	*mock.Call
}

// GetBySteamID is a helper method to define mock.On call
//   - ctx context.Context
//   - sid64 steamid.SteamID
//   - deletedOk bool
func (_e *MockBanSteamUsecase_Expecter) GetBySteamID(ctx interface{}, sid64 interface{}, deletedOk interface{}) *MockBanSteamUsecase_GetBySteamID_Call {
	return &MockBanSteamUsecase_GetBySteamID_Call{Call: _e.mock.On("GetBySteamID", ctx, sid64, deletedOk)}
}

func (_c *MockBanSteamUsecase_GetBySteamID_Call) Run(run func(ctx context.Context, sid64 steamid.SteamID, deletedOk bool)) *MockBanSteamUsecase_GetBySteamID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID), args[2].(bool))
	})
	return _c
}

func (_c *MockBanSteamUsecase_GetBySteamID_Call) Return(_a0 domain.BannedSteamPerson, _a1 error) *MockBanSteamUsecase_GetBySteamID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanSteamUsecase_GetBySteamID_Call) RunAndReturn(run func(context.Context, steamid.SteamID, bool) (domain.BannedSteamPerson, error)) *MockBanSteamUsecase_GetBySteamID_Call {
	_c.Call.Return(run)
	return _c
}

// GetOlderThan provides a mock function with given fields: ctx, filter, since
func (_m *MockBanSteamUsecase) GetOlderThan(ctx context.Context, filter domain.QueryFilter, since time.Time) ([]domain.BanSteam, error) {
	ret := _m.Called(ctx, filter, since)

	if len(ret) == 0 {
		panic("no return value specified for GetOlderThan")
	}

	var r0 []domain.BanSteam
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.QueryFilter, time.Time) ([]domain.BanSteam, error)); ok {
		return rf(ctx, filter, since)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.QueryFilter, time.Time) []domain.BanSteam); ok {
		r0 = rf(ctx, filter, since)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BanSteam)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.QueryFilter, time.Time) error); ok {
		r1 = rf(ctx, filter, since)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanSteamUsecase_GetOlderThan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOlderThan'
type MockBanSteamUsecase_GetOlderThan_Call struct {
	*mock.Call
}

// GetOlderThan is a helper method to define mock.On call
//   - ctx context.Context
//   - filter domain.QueryFilter
//   - since time.Time
func (_e *MockBanSteamUsecase_Expecter) GetOlderThan(ctx interface{}, filter interface{}, since interface{}) *MockBanSteamUsecase_GetOlderThan_Call {
	return &MockBanSteamUsecase_GetOlderThan_Call{Call: _e.mock.On("GetOlderThan", ctx, filter, since)}
}

func (_c *MockBanSteamUsecase_GetOlderThan_Call) Run(run func(ctx context.Context, filter domain.QueryFilter, since time.Time)) *MockBanSteamUsecase_GetOlderThan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.QueryFilter), args[2].(time.Time))
	})
	return _c
}

func (_c *MockBanSteamUsecase_GetOlderThan_Call) Return(_a0 []domain.BanSteam, _a1 error) *MockBanSteamUsecase_GetOlderThan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanSteamUsecase_GetOlderThan_Call) RunAndReturn(run func(context.Context, domain.QueryFilter, time.Time) ([]domain.BanSteam, error)) *MockBanSteamUsecase_GetOlderThan_Call {
	_c.Call.Return(run)
	return _c
}

// IsFriendBanned provides a mock function with given fields: steamID
func (_m *MockBanSteamUsecase) IsFriendBanned(steamID steamid.SteamID) (steamid.SteamID, bool) {
	ret := _m.Called(steamID)

	if len(ret) == 0 {
		panic("no return value specified for IsFriendBanned")
	}

	var r0 steamid.SteamID
	var r1 bool
	if rf, ok := ret.Get(0).(func(steamid.SteamID) (steamid.SteamID, bool)); ok {
		return rf(steamID)
	}
	if rf, ok := ret.Get(0).(func(steamid.SteamID) steamid.SteamID); ok {
		r0 = rf(steamID)
	} else {
		r0 = ret.Get(0).(steamid.SteamID)
	}

	if rf, ok := ret.Get(1).(func(steamid.SteamID) bool); ok {
		r1 = rf(steamID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockBanSteamUsecase_IsFriendBanned_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsFriendBanned'
type MockBanSteamUsecase_IsFriendBanned_Call struct {
	*mock.Call
}

// IsFriendBanned is a helper method to define mock.On call
//   - steamID steamid.SteamID
func (_e *MockBanSteamUsecase_Expecter) IsFriendBanned(steamID interface{}) *MockBanSteamUsecase_IsFriendBanned_Call {
	return &MockBanSteamUsecase_IsFriendBanned_Call{Call: _e.mock.On("IsFriendBanned", steamID)}
}

func (_c *MockBanSteamUsecase_IsFriendBanned_Call) Run(run func(steamID steamid.SteamID)) *MockBanSteamUsecase_IsFriendBanned_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(steamid.SteamID))
	})
	return _c
}

func (_c *MockBanSteamUsecase_IsFriendBanned_Call) Return(_a0 steamid.SteamID, _a1 bool) *MockBanSteamUsecase_IsFriendBanned_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanSteamUsecase_IsFriendBanned_Call) RunAndReturn(run func(steamid.SteamID) (steamid.SteamID, bool)) *MockBanSteamUsecase_IsFriendBanned_Call {
	_c.Call.Return(run)
	return _c
}

// IsOnIPWithBan provides a mock function with given fields: ctx, curUser, steamID, address
func (_m *MockBanSteamUsecase) IsOnIPWithBan(ctx context.Context, curUser domain.PersonInfo, steamID steamid.SteamID, address netip.Addr) (bool, error) {
	ret := _m.Called(ctx, curUser, steamID, address)

	if len(ret) == 0 {
		panic("no return value specified for IsOnIPWithBan")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, steamid.SteamID, netip.Addr) (bool, error)); ok {
		return rf(ctx, curUser, steamID, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, steamid.SteamID, netip.Addr) bool); ok {
		r0 = rf(ctx, curUser, steamID, address)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PersonInfo, steamid.SteamID, netip.Addr) error); ok {
		r1 = rf(ctx, curUser, steamID, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanSteamUsecase_IsOnIPWithBan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsOnIPWithBan'
type MockBanSteamUsecase_IsOnIPWithBan_Call struct {
	*mock.Call
}

// IsOnIPWithBan is a helper method to define mock.On call
//   - ctx context.Context
//   - curUser domain.PersonInfo
//   - steamID steamid.SteamID
//   - address netip.Addr
func (_e *MockBanSteamUsecase_Expecter) IsOnIPWithBan(ctx interface{}, curUser interface{}, steamID interface{}, address interface{}) *MockBanSteamUsecase_IsOnIPWithBan_Call {
	return &MockBanSteamUsecase_IsOnIPWithBan_Call{Call: _e.mock.On("IsOnIPWithBan", ctx, curUser, steamID, address)}
}

func (_c *MockBanSteamUsecase_IsOnIPWithBan_Call) Run(run func(ctx context.Context, curUser domain.PersonInfo, steamID steamid.SteamID, address netip.Addr)) *MockBanSteamUsecase_IsOnIPWithBan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PersonInfo), args[2].(steamid.SteamID), args[3].(netip.Addr))
	})
	return _c
}

func (_c *MockBanSteamUsecase_IsOnIPWithBan_Call) Return(_a0 bool, _a1 error) *MockBanSteamUsecase_IsOnIPWithBan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanSteamUsecase_IsOnIPWithBan_Call) RunAndReturn(run func(context.Context, domain.PersonInfo, steamid.SteamID, netip.Addr) (bool, error)) *MockBanSteamUsecase_IsOnIPWithBan_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, ban
func (_m *MockBanSteamUsecase) Save(ctx context.Context, ban *domain.BanSteam) error {
	ret := _m.Called(ctx, ban)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BanSteam) error); ok {
		r0 = rf(ctx, ban)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanSteamUsecase_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockBanSteamUsecase_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - ban *domain.BanSteam
func (_e *MockBanSteamUsecase_Expecter) Save(ctx interface{}, ban interface{}) *MockBanSteamUsecase_Save_Call {
	return &MockBanSteamUsecase_Save_Call{Call: _e.mock.On("Save", ctx, ban)}
}

func (_c *MockBanSteamUsecase_Save_Call) Run(run func(ctx context.Context, ban *domain.BanSteam)) *MockBanSteamUsecase_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.BanSteam))
	})
	return _c
}

func (_c *MockBanSteamUsecase_Save_Call) Return(_a0 error) *MockBanSteamUsecase_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanSteamUsecase_Save_Call) RunAndReturn(run func(context.Context, *domain.BanSteam) error) *MockBanSteamUsecase_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Stats provides a mock function with given fields: ctx, stats
func (_m *MockBanSteamUsecase) Stats(ctx context.Context, stats *domain.Stats) error {
	ret := _m.Called(ctx, stats)

	if len(ret) == 0 {
		panic("no return value specified for Stats")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Stats) error); ok {
		r0 = rf(ctx, stats)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanSteamUsecase_Stats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stats'
type MockBanSteamUsecase_Stats_Call struct {
	*mock.Call
}

// Stats is a helper method to define mock.On call
//   - ctx context.Context
//   - stats *domain.Stats
func (_e *MockBanSteamUsecase_Expecter) Stats(ctx interface{}, stats interface{}) *MockBanSteamUsecase_Stats_Call {
	return &MockBanSteamUsecase_Stats_Call{Call: _e.mock.On("Stats", ctx, stats)}
}

func (_c *MockBanSteamUsecase_Stats_Call) Run(run func(ctx context.Context, stats *domain.Stats)) *MockBanSteamUsecase_Stats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Stats))
	})
	return _c
}

func (_c *MockBanSteamUsecase_Stats_Call) Return(_a0 error) *MockBanSteamUsecase_Stats_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanSteamUsecase_Stats_Call) RunAndReturn(run func(context.Context, *domain.Stats) error) *MockBanSteamUsecase_Stats_Call {
	_c.Call.Return(run)
	return _c
}

// Unban provides a mock function with given fields: ctx, targetSID, reason
func (_m *MockBanSteamUsecase) Unban(ctx context.Context, targetSID steamid.SteamID, reason string) (bool, error) {
	ret := _m.Called(ctx, targetSID, reason)

	if len(ret) == 0 {
		panic("no return value specified for Unban")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, string) (bool, error)); ok {
		return rf(ctx, targetSID, reason)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, string) bool); ok {
		r0 = rf(ctx, targetSID, reason)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.SteamID, string) error); ok {
		r1 = rf(ctx, targetSID, reason)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanSteamUsecase_Unban_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unban'
type MockBanSteamUsecase_Unban_Call struct {
	*mock.Call
}

// Unban is a helper method to define mock.On call
//   - ctx context.Context
//   - targetSID steamid.SteamID
//   - reason string
func (_e *MockBanSteamUsecase_Expecter) Unban(ctx interface{}, targetSID interface{}, reason interface{}) *MockBanSteamUsecase_Unban_Call {
	return &MockBanSteamUsecase_Unban_Call{Call: _e.mock.On("Unban", ctx, targetSID, reason)}
}

func (_c *MockBanSteamUsecase_Unban_Call) Run(run func(ctx context.Context, targetSID steamid.SteamID, reason string)) *MockBanSteamUsecase_Unban_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID), args[2].(string))
	})
	return _c
}

func (_c *MockBanSteamUsecase_Unban_Call) Return(_a0 bool, _a1 error) *MockBanSteamUsecase_Unban_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanSteamUsecase_Unban_Call) RunAndReturn(run func(context.Context, steamid.SteamID, string) (bool, error)) *MockBanSteamUsecase_Unban_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBanSteamUsecase creates a new instance of MockBanSteamUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBanSteamUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBanSteamUsecase {
	mock := &MockBanSteamUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
