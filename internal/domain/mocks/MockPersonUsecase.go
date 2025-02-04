// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"

	net "net"

	steamid "github.com/leighmacdonald/steamid/v4/steamid"
)

// MockPersonUsecase is an autogenerated mock type for the PersonUsecase type
type MockPersonUsecase struct {
	mock.Mock
}

type MockPersonUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPersonUsecase) EXPECT() *MockPersonUsecase_Expecter {
	return &MockPersonUsecase_Expecter{mock: &_m.Mock}
}

// DropPerson provides a mock function with given fields: ctx, steamID
func (_m *MockPersonUsecase) DropPerson(ctx context.Context, steamID steamid.SteamID) error {
	ret := _m.Called(ctx, steamID)

	if len(ret) == 0 {
		panic("no return value specified for DropPerson")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID) error); ok {
		r0 = rf(ctx, steamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPersonUsecase_DropPerson_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropPerson'
type MockPersonUsecase_DropPerson_Call struct {
	*mock.Call
}

// DropPerson is a helper method to define mock.On call
//   - ctx context.Context
//   - steamID steamid.SteamID
func (_e *MockPersonUsecase_Expecter) DropPerson(ctx interface{}, steamID interface{}) *MockPersonUsecase_DropPerson_Call {
	return &MockPersonUsecase_DropPerson_Call{Call: _e.mock.On("DropPerson", ctx, steamID)}
}

func (_c *MockPersonUsecase_DropPerson_Call) Run(run func(ctx context.Context, steamID steamid.SteamID)) *MockPersonUsecase_DropPerson_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID))
	})
	return _c
}

func (_c *MockPersonUsecase_DropPerson_Call) Return(_a0 error) *MockPersonUsecase_DropPerson_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPersonUsecase_DropPerson_Call) RunAndReturn(run func(context.Context, steamid.SteamID) error) *MockPersonUsecase_DropPerson_Call {
	_c.Call.Return(run)
	return _c
}

// GetExpiredProfiles provides a mock function with given fields: ctx, limit
func (_m *MockPersonUsecase) GetExpiredProfiles(ctx context.Context, limit uint64) ([]domain.Person, error) {
	ret := _m.Called(ctx, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetExpiredProfiles")
	}

	var r0 []domain.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]domain.Person, error)); ok {
		return rf(ctx, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []domain.Person); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetExpiredProfiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpiredProfiles'
type MockPersonUsecase_GetExpiredProfiles_Call struct {
	*mock.Call
}

// GetExpiredProfiles is a helper method to define mock.On call
//   - ctx context.Context
//   - limit uint64
func (_e *MockPersonUsecase_Expecter) GetExpiredProfiles(ctx interface{}, limit interface{}) *MockPersonUsecase_GetExpiredProfiles_Call {
	return &MockPersonUsecase_GetExpiredProfiles_Call{Call: _e.mock.On("GetExpiredProfiles", ctx, limit)}
}

func (_c *MockPersonUsecase_GetExpiredProfiles_Call) Run(run func(ctx context.Context, limit uint64)) *MockPersonUsecase_GetExpiredProfiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockPersonUsecase_GetExpiredProfiles_Call) Return(_a0 []domain.Person, _a1 error) *MockPersonUsecase_GetExpiredProfiles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetExpiredProfiles_Call) RunAndReturn(run func(context.Context, uint64) ([]domain.Person, error)) *MockPersonUsecase_GetExpiredProfiles_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrCreatePersonBySteamID provides a mock function with given fields: ctx, sid64
func (_m *MockPersonUsecase) GetOrCreatePersonBySteamID(ctx context.Context, sid64 steamid.SteamID) (domain.Person, error) {
	ret := _m.Called(ctx, sid64)

	if len(ret) == 0 {
		panic("no return value specified for GetOrCreatePersonBySteamID")
	}

	var r0 domain.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID) (domain.Person, error)); ok {
		return rf(ctx, sid64)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID) domain.Person); ok {
		r0 = rf(ctx, sid64)
	} else {
		r0 = ret.Get(0).(domain.Person)
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.SteamID) error); ok {
		r1 = rf(ctx, sid64)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetOrCreatePersonBySteamID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrCreatePersonBySteamID'
type MockPersonUsecase_GetOrCreatePersonBySteamID_Call struct {
	*mock.Call
}

// GetOrCreatePersonBySteamID is a helper method to define mock.On call
//   - ctx context.Context
//   - sid64 steamid.SteamID
func (_e *MockPersonUsecase_Expecter) GetOrCreatePersonBySteamID(ctx interface{}, sid64 interface{}) *MockPersonUsecase_GetOrCreatePersonBySteamID_Call {
	return &MockPersonUsecase_GetOrCreatePersonBySteamID_Call{Call: _e.mock.On("GetOrCreatePersonBySteamID", ctx, sid64)}
}

func (_c *MockPersonUsecase_GetOrCreatePersonBySteamID_Call) Run(run func(ctx context.Context, sid64 steamid.SteamID)) *MockPersonUsecase_GetOrCreatePersonBySteamID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID))
	})
	return _c
}

func (_c *MockPersonUsecase_GetOrCreatePersonBySteamID_Call) Return(_a0 domain.Person, _a1 error) *MockPersonUsecase_GetOrCreatePersonBySteamID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetOrCreatePersonBySteamID_Call) RunAndReturn(run func(context.Context, steamid.SteamID) (domain.Person, error)) *MockPersonUsecase_GetOrCreatePersonBySteamID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPeople provides a mock function with given fields: ctx, filter
func (_m *MockPersonUsecase) GetPeople(ctx context.Context, filter domain.PlayerQuery) (domain.People, int64, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetPeople")
	}

	var r0 domain.People
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PlayerQuery) (domain.People, int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PlayerQuery) domain.People); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.People)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PlayerQuery) int64); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.PlayerQuery) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockPersonUsecase_GetPeople_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPeople'
type MockPersonUsecase_GetPeople_Call struct {
	*mock.Call
}

// GetPeople is a helper method to define mock.On call
//   - ctx context.Context
//   - filter domain.PlayerQuery
func (_e *MockPersonUsecase_Expecter) GetPeople(ctx interface{}, filter interface{}) *MockPersonUsecase_GetPeople_Call {
	return &MockPersonUsecase_GetPeople_Call{Call: _e.mock.On("GetPeople", ctx, filter)}
}

func (_c *MockPersonUsecase_GetPeople_Call) Run(run func(ctx context.Context, filter domain.PlayerQuery)) *MockPersonUsecase_GetPeople_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PlayerQuery))
	})
	return _c
}

func (_c *MockPersonUsecase_GetPeople_Call) Return(_a0 domain.People, _a1 int64, _a2 error) *MockPersonUsecase_GetPeople_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockPersonUsecase_GetPeople_Call) RunAndReturn(run func(context.Context, domain.PlayerQuery) (domain.People, int64, error)) *MockPersonUsecase_GetPeople_Call {
	_c.Call.Return(run)
	return _c
}

// GetPeopleBySteamID provides a mock function with given fields: ctx, steamIds
func (_m *MockPersonUsecase) GetPeopleBySteamID(ctx context.Context, steamIds steamid.Collection) (domain.People, error) {
	ret := _m.Called(ctx, steamIds)

	if len(ret) == 0 {
		panic("no return value specified for GetPeopleBySteamID")
	}

	var r0 domain.People
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.Collection) (domain.People, error)); ok {
		return rf(ctx, steamIds)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.Collection) domain.People); ok {
		r0 = rf(ctx, steamIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.People)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.Collection) error); ok {
		r1 = rf(ctx, steamIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetPeopleBySteamID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPeopleBySteamID'
type MockPersonUsecase_GetPeopleBySteamID_Call struct {
	*mock.Call
}

// GetPeopleBySteamID is a helper method to define mock.On call
//   - ctx context.Context
//   - steamIds steamid.Collection
func (_e *MockPersonUsecase_Expecter) GetPeopleBySteamID(ctx interface{}, steamIds interface{}) *MockPersonUsecase_GetPeopleBySteamID_Call {
	return &MockPersonUsecase_GetPeopleBySteamID_Call{Call: _e.mock.On("GetPeopleBySteamID", ctx, steamIds)}
}

func (_c *MockPersonUsecase_GetPeopleBySteamID_Call) Run(run func(ctx context.Context, steamIds steamid.Collection)) *MockPersonUsecase_GetPeopleBySteamID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.Collection))
	})
	return _c
}

func (_c *MockPersonUsecase_GetPeopleBySteamID_Call) Return(_a0 domain.People, _a1 error) *MockPersonUsecase_GetPeopleBySteamID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetPeopleBySteamID_Call) RunAndReturn(run func(context.Context, steamid.Collection) (domain.People, error)) *MockPersonUsecase_GetPeopleBySteamID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonByDiscordID provides a mock function with given fields: ctx, discordID
func (_m *MockPersonUsecase) GetPersonByDiscordID(ctx context.Context, discordID string) (domain.Person, error) {
	ret := _m.Called(ctx, discordID)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonByDiscordID")
	}

	var r0 domain.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Person, error)); ok {
		return rf(ctx, discordID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Person); ok {
		r0 = rf(ctx, discordID)
	} else {
		r0 = ret.Get(0).(domain.Person)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, discordID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetPersonByDiscordID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonByDiscordID'
type MockPersonUsecase_GetPersonByDiscordID_Call struct {
	*mock.Call
}

// GetPersonByDiscordID is a helper method to define mock.On call
//   - ctx context.Context
//   - discordID string
func (_e *MockPersonUsecase_Expecter) GetPersonByDiscordID(ctx interface{}, discordID interface{}) *MockPersonUsecase_GetPersonByDiscordID_Call {
	return &MockPersonUsecase_GetPersonByDiscordID_Call{Call: _e.mock.On("GetPersonByDiscordID", ctx, discordID)}
}

func (_c *MockPersonUsecase_GetPersonByDiscordID_Call) Run(run func(ctx context.Context, discordID string)) *MockPersonUsecase_GetPersonByDiscordID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockPersonUsecase_GetPersonByDiscordID_Call) Return(_a0 domain.Person, _a1 error) *MockPersonUsecase_GetPersonByDiscordID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetPersonByDiscordID_Call) RunAndReturn(run func(context.Context, string) (domain.Person, error)) *MockPersonUsecase_GetPersonByDiscordID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonBySteamID provides a mock function with given fields: ctx, sid64
func (_m *MockPersonUsecase) GetPersonBySteamID(ctx context.Context, sid64 steamid.SteamID) (domain.Person, error) {
	ret := _m.Called(ctx, sid64)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonBySteamID")
	}

	var r0 domain.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID) (domain.Person, error)); ok {
		return rf(ctx, sid64)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID) domain.Person); ok {
		r0 = rf(ctx, sid64)
	} else {
		r0 = ret.Get(0).(domain.Person)
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.SteamID) error); ok {
		r1 = rf(ctx, sid64)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetPersonBySteamID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonBySteamID'
type MockPersonUsecase_GetPersonBySteamID_Call struct {
	*mock.Call
}

// GetPersonBySteamID is a helper method to define mock.On call
//   - ctx context.Context
//   - sid64 steamid.SteamID
func (_e *MockPersonUsecase_Expecter) GetPersonBySteamID(ctx interface{}, sid64 interface{}) *MockPersonUsecase_GetPersonBySteamID_Call {
	return &MockPersonUsecase_GetPersonBySteamID_Call{Call: _e.mock.On("GetPersonBySteamID", ctx, sid64)}
}

func (_c *MockPersonUsecase_GetPersonBySteamID_Call) Run(run func(ctx context.Context, sid64 steamid.SteamID)) *MockPersonUsecase_GetPersonBySteamID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID))
	})
	return _c
}

func (_c *MockPersonUsecase_GetPersonBySteamID_Call) Return(_a0 domain.Person, _a1 error) *MockPersonUsecase_GetPersonBySteamID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetPersonBySteamID_Call) RunAndReturn(run func(context.Context, steamid.SteamID) (domain.Person, error)) *MockPersonUsecase_GetPersonBySteamID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonMessageByID provides a mock function with given fields: ctx, personMessageID
func (_m *MockPersonUsecase) GetPersonMessageByID(ctx context.Context, personMessageID int64) (domain.PersonMessage, error) {
	ret := _m.Called(ctx, personMessageID)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonMessageByID")
	}

	var r0 domain.PersonMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.PersonMessage, error)); ok {
		return rf(ctx, personMessageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.PersonMessage); ok {
		r0 = rf(ctx, personMessageID)
	} else {
		r0 = ret.Get(0).(domain.PersonMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, personMessageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetPersonMessageByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonMessageByID'
type MockPersonUsecase_GetPersonMessageByID_Call struct {
	*mock.Call
}

// GetPersonMessageByID is a helper method to define mock.On call
//   - ctx context.Context
//   - personMessageID int64
func (_e *MockPersonUsecase_Expecter) GetPersonMessageByID(ctx interface{}, personMessageID interface{}) *MockPersonUsecase_GetPersonMessageByID_Call {
	return &MockPersonUsecase_GetPersonMessageByID_Call{Call: _e.mock.On("GetPersonMessageByID", ctx, personMessageID)}
}

func (_c *MockPersonUsecase_GetPersonMessageByID_Call) Run(run func(ctx context.Context, personMessageID int64)) *MockPersonUsecase_GetPersonMessageByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockPersonUsecase_GetPersonMessageByID_Call) Return(_a0 domain.PersonMessage, _a1 error) *MockPersonUsecase_GetPersonMessageByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetPersonMessageByID_Call) RunAndReturn(run func(context.Context, int64) (domain.PersonMessage, error)) *MockPersonUsecase_GetPersonMessageByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonSettings provides a mock function with given fields: ctx, steamID
func (_m *MockPersonUsecase) GetPersonSettings(ctx context.Context, steamID steamid.SteamID) (domain.PersonSettings, error) {
	ret := _m.Called(ctx, steamID)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonSettings")
	}

	var r0 domain.PersonSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID) (domain.PersonSettings, error)); ok {
		return rf(ctx, steamID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID) domain.PersonSettings); ok {
		r0 = rf(ctx, steamID)
	} else {
		r0 = ret.Get(0).(domain.PersonSettings)
	}

	if rf, ok := ret.Get(1).(func(context.Context, steamid.SteamID) error); ok {
		r1 = rf(ctx, steamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetPersonSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonSettings'
type MockPersonUsecase_GetPersonSettings_Call struct {
	*mock.Call
}

// GetPersonSettings is a helper method to define mock.On call
//   - ctx context.Context
//   - steamID steamid.SteamID
func (_e *MockPersonUsecase_Expecter) GetPersonSettings(ctx interface{}, steamID interface{}) *MockPersonUsecase_GetPersonSettings_Call {
	return &MockPersonUsecase_GetPersonSettings_Call{Call: _e.mock.On("GetPersonSettings", ctx, steamID)}
}

func (_c *MockPersonUsecase_GetPersonSettings_Call) Run(run func(ctx context.Context, steamID steamid.SteamID)) *MockPersonUsecase_GetPersonSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID))
	})
	return _c
}

func (_c *MockPersonUsecase_GetPersonSettings_Call) Return(_a0 domain.PersonSettings, _a1 error) *MockPersonUsecase_GetPersonSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetPersonSettings_Call) RunAndReturn(run func(context.Context, steamid.SteamID) (domain.PersonSettings, error)) *MockPersonUsecase_GetPersonSettings_Call {
	_c.Call.Return(run)
	return _c
}

// GetSteamIdsAbove provides a mock function with given fields: ctx, privilege
func (_m *MockPersonUsecase) GetSteamIdsAbove(ctx context.Context, privilege domain.Privilege) (steamid.Collection, error) {
	ret := _m.Called(ctx, privilege)

	if len(ret) == 0 {
		panic("no return value specified for GetSteamIdsAbove")
	}

	var r0 steamid.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Privilege) (steamid.Collection, error)); ok {
		return rf(ctx, privilege)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Privilege) steamid.Collection); ok {
		r0 = rf(ctx, privilege)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(steamid.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Privilege) error); ok {
		r1 = rf(ctx, privilege)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetSteamIdsAbove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSteamIdsAbove'
type MockPersonUsecase_GetSteamIdsAbove_Call struct {
	*mock.Call
}

// GetSteamIdsAbove is a helper method to define mock.On call
//   - ctx context.Context
//   - privilege domain.Privilege
func (_e *MockPersonUsecase_Expecter) GetSteamIdsAbove(ctx interface{}, privilege interface{}) *MockPersonUsecase_GetSteamIdsAbove_Call {
	return &MockPersonUsecase_GetSteamIdsAbove_Call{Call: _e.mock.On("GetSteamIdsAbove", ctx, privilege)}
}

func (_c *MockPersonUsecase_GetSteamIdsAbove_Call) Run(run func(ctx context.Context, privilege domain.Privilege)) *MockPersonUsecase_GetSteamIdsAbove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Privilege))
	})
	return _c
}

func (_c *MockPersonUsecase_GetSteamIdsAbove_Call) Return(_a0 steamid.Collection, _a1 error) *MockPersonUsecase_GetSteamIdsAbove_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetSteamIdsAbove_Call) RunAndReturn(run func(context.Context, domain.Privilege) (steamid.Collection, error)) *MockPersonUsecase_GetSteamIdsAbove_Call {
	_c.Call.Return(run)
	return _c
}

// GetSteamsAtAddress provides a mock function with given fields: ctx, addr
func (_m *MockPersonUsecase) GetSteamsAtAddress(ctx context.Context, addr net.IP) (steamid.Collection, error) {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for GetSteamsAtAddress")
	}

	var r0 steamid.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, net.IP) (steamid.Collection, error)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, net.IP) steamid.Collection); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(steamid.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, net.IP) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_GetSteamsAtAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSteamsAtAddress'
type MockPersonUsecase_GetSteamsAtAddress_Call struct {
	*mock.Call
}

// GetSteamsAtAddress is a helper method to define mock.On call
//   - ctx context.Context
//   - addr net.IP
func (_e *MockPersonUsecase_Expecter) GetSteamsAtAddress(ctx interface{}, addr interface{}) *MockPersonUsecase_GetSteamsAtAddress_Call {
	return &MockPersonUsecase_GetSteamsAtAddress_Call{Call: _e.mock.On("GetSteamsAtAddress", ctx, addr)}
}

func (_c *MockPersonUsecase_GetSteamsAtAddress_Call) Run(run func(ctx context.Context, addr net.IP)) *MockPersonUsecase_GetSteamsAtAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(net.IP))
	})
	return _c
}

func (_c *MockPersonUsecase_GetSteamsAtAddress_Call) Return(_a0 steamid.Collection, _a1 error) *MockPersonUsecase_GetSteamsAtAddress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_GetSteamsAtAddress_Call) RunAndReturn(run func(context.Context, net.IP) (steamid.Collection, error)) *MockPersonUsecase_GetSteamsAtAddress_Call {
	_c.Call.Return(run)
	return _c
}

// QueryProfile provides a mock function with given fields: ctx, query
func (_m *MockPersonUsecase) QueryProfile(ctx context.Context, query string) (domain.ProfileResponse, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for QueryProfile")
	}

	var r0 domain.ProfileResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.ProfileResponse, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ProfileResponse); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(domain.ProfileResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_QueryProfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryProfile'
type MockPersonUsecase_QueryProfile_Call struct {
	*mock.Call
}

// QueryProfile is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
func (_e *MockPersonUsecase_Expecter) QueryProfile(ctx interface{}, query interface{}) *MockPersonUsecase_QueryProfile_Call {
	return &MockPersonUsecase_QueryProfile_Call{Call: _e.mock.On("QueryProfile", ctx, query)}
}

func (_c *MockPersonUsecase_QueryProfile_Call) Run(run func(ctx context.Context, query string)) *MockPersonUsecase_QueryProfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockPersonUsecase_QueryProfile_Call) Return(_a0 domain.ProfileResponse, _a1 error) *MockPersonUsecase_QueryProfile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_QueryProfile_Call) RunAndReturn(run func(context.Context, string) (domain.ProfileResponse, error)) *MockPersonUsecase_QueryProfile_Call {
	_c.Call.Return(run)
	return _c
}

// SavePerson provides a mock function with given fields: ctx, person
func (_m *MockPersonUsecase) SavePerson(ctx context.Context, person *domain.Person) error {
	ret := _m.Called(ctx, person)

	if len(ret) == 0 {
		panic("no return value specified for SavePerson")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Person) error); ok {
		r0 = rf(ctx, person)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPersonUsecase_SavePerson_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SavePerson'
type MockPersonUsecase_SavePerson_Call struct {
	*mock.Call
}

// SavePerson is a helper method to define mock.On call
//   - ctx context.Context
//   - person *domain.Person
func (_e *MockPersonUsecase_Expecter) SavePerson(ctx interface{}, person interface{}) *MockPersonUsecase_SavePerson_Call {
	return &MockPersonUsecase_SavePerson_Call{Call: _e.mock.On("SavePerson", ctx, person)}
}

func (_c *MockPersonUsecase_SavePerson_Call) Run(run func(ctx context.Context, person *domain.Person)) *MockPersonUsecase_SavePerson_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Person))
	})
	return _c
}

func (_c *MockPersonUsecase_SavePerson_Call) Return(_a0 error) *MockPersonUsecase_SavePerson_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPersonUsecase_SavePerson_Call) RunAndReturn(run func(context.Context, *domain.Person) error) *MockPersonUsecase_SavePerson_Call {
	_c.Call.Return(run)
	return _c
}

// SavePersonSettings provides a mock function with given fields: ctx, user, req
func (_m *MockPersonUsecase) SavePersonSettings(ctx context.Context, user domain.PersonInfo, req domain.PersonSettingsUpdate) (domain.PersonSettings, error) {
	ret := _m.Called(ctx, user, req)

	if len(ret) == 0 {
		panic("no return value specified for SavePersonSettings")
	}

	var r0 domain.PersonSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, domain.PersonSettingsUpdate) (domain.PersonSettings, error)); ok {
		return rf(ctx, user, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, domain.PersonSettingsUpdate) domain.PersonSettings); ok {
		r0 = rf(ctx, user, req)
	} else {
		r0 = ret.Get(0).(domain.PersonSettings)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PersonInfo, domain.PersonSettingsUpdate) error); ok {
		r1 = rf(ctx, user, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonUsecase_SavePersonSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SavePersonSettings'
type MockPersonUsecase_SavePersonSettings_Call struct {
	*mock.Call
}

// SavePersonSettings is a helper method to define mock.On call
//   - ctx context.Context
//   - user domain.PersonInfo
//   - req domain.PersonSettingsUpdate
func (_e *MockPersonUsecase_Expecter) SavePersonSettings(ctx interface{}, user interface{}, req interface{}) *MockPersonUsecase_SavePersonSettings_Call {
	return &MockPersonUsecase_SavePersonSettings_Call{Call: _e.mock.On("SavePersonSettings", ctx, user, req)}
}

func (_c *MockPersonUsecase_SavePersonSettings_Call) Run(run func(ctx context.Context, user domain.PersonInfo, req domain.PersonSettingsUpdate)) *MockPersonUsecase_SavePersonSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PersonInfo), args[2].(domain.PersonSettingsUpdate))
	})
	return _c
}

func (_c *MockPersonUsecase_SavePersonSettings_Call) Return(_a0 domain.PersonSettings, _a1 error) *MockPersonUsecase_SavePersonSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPersonUsecase_SavePersonSettings_Call) RunAndReturn(run func(context.Context, domain.PersonInfo, domain.PersonSettingsUpdate) (domain.PersonSettings, error)) *MockPersonUsecase_SavePersonSettings_Call {
	_c.Call.Return(run)
	return _c
}

// SetPermissionLevel provides a mock function with given fields: ctx, steamID, level
func (_m *MockPersonUsecase) SetPermissionLevel(ctx context.Context, steamID steamid.SteamID, level domain.Privilege) error {
	ret := _m.Called(ctx, steamID, level)

	if len(ret) == 0 {
		panic("no return value specified for SetPermissionLevel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, domain.Privilege) error); ok {
		r0 = rf(ctx, steamID, level)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPersonUsecase_SetPermissionLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPermissionLevel'
type MockPersonUsecase_SetPermissionLevel_Call struct {
	*mock.Call
}

// SetPermissionLevel is a helper method to define mock.On call
//   - ctx context.Context
//   - steamID steamid.SteamID
//   - level domain.Privilege
func (_e *MockPersonUsecase_Expecter) SetPermissionLevel(ctx interface{}, steamID interface{}, level interface{}) *MockPersonUsecase_SetPermissionLevel_Call {
	return &MockPersonUsecase_SetPermissionLevel_Call{Call: _e.mock.On("SetPermissionLevel", ctx, steamID, level)}
}

func (_c *MockPersonUsecase_SetPermissionLevel_Call) Run(run func(ctx context.Context, steamID steamid.SteamID, level domain.Privilege)) *MockPersonUsecase_SetPermissionLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID), args[2].(domain.Privilege))
	})
	return _c
}

func (_c *MockPersonUsecase_SetPermissionLevel_Call) Return(_a0 error) *MockPersonUsecase_SetPermissionLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPersonUsecase_SetPermissionLevel_Call) RunAndReturn(run func(context.Context, steamid.SteamID, domain.Privilege) error) *MockPersonUsecase_SetPermissionLevel_Call {
	_c.Call.Return(run)
	return _c
}

// SetSteam provides a mock function with given fields: ctx, sid64, discordID
func (_m *MockPersonUsecase) SetSteam(ctx context.Context, sid64 steamid.SteamID, discordID string) error {
	ret := _m.Called(ctx, sid64, discordID)

	if len(ret) == 0 {
		panic("no return value specified for SetSteam")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.SteamID, string) error); ok {
		r0 = rf(ctx, sid64, discordID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPersonUsecase_SetSteam_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSteam'
type MockPersonUsecase_SetSteam_Call struct {
	*mock.Call
}

// SetSteam is a helper method to define mock.On call
//   - ctx context.Context
//   - sid64 steamid.SteamID
//   - discordID string
func (_e *MockPersonUsecase_Expecter) SetSteam(ctx interface{}, sid64 interface{}, discordID interface{}) *MockPersonUsecase_SetSteam_Call {
	return &MockPersonUsecase_SetSteam_Call{Call: _e.mock.On("SetSteam", ctx, sid64, discordID)}
}

func (_c *MockPersonUsecase_SetSteam_Call) Run(run func(ctx context.Context, sid64 steamid.SteamID, discordID string)) *MockPersonUsecase_SetSteam_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.SteamID), args[2].(string))
	})
	return _c
}

func (_c *MockPersonUsecase_SetSteam_Call) Return(_a0 error) *MockPersonUsecase_SetSteam_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPersonUsecase_SetSteam_Call) RunAndReturn(run func(context.Context, steamid.SteamID, string) error) *MockPersonUsecase_SetSteam_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPersonUsecase creates a new instance of MockPersonUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPersonUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPersonUsecase {
	mock := &MockPersonUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
