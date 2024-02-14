// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"

	steamid "github.com/leighmacdonald/steamid/v3/steamid"
)

// MockBanGroupRepository is an autogenerated mock type for the BanGroupRepository type
type MockBanGroupRepository struct {
	mock.Mock
}

type MockBanGroupRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBanGroupRepository) EXPECT() *MockBanGroupRepository_Expecter {
	return &MockBanGroupRepository_Expecter{mock: &_m.Mock}
}

// Ban provides a mock function with given fields: ctx, banGroup
func (_m *MockBanGroupRepository) Ban(ctx context.Context, banGroup *domain.BanGroup) error {
	ret := _m.Called(ctx, banGroup)

	if len(ret) == 0 {
		panic("no return value specified for Ban")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BanGroup) error); ok {
		r0 = rf(ctx, banGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanGroupRepository_Ban_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ban'
type MockBanGroupRepository_Ban_Call struct {
	*mock.Call
}

// Ban is a helper method to define mock.On call
//   - ctx context.Context
//   - banGroup *domain.BanGroup
func (_e *MockBanGroupRepository_Expecter) Ban(ctx interface{}, banGroup interface{}) *MockBanGroupRepository_Ban_Call {
	return &MockBanGroupRepository_Ban_Call{Call: _e.mock.On("Ban", ctx, banGroup)}
}

func (_c *MockBanGroupRepository_Ban_Call) Run(run func(ctx context.Context, banGroup *domain.BanGroup)) *MockBanGroupRepository_Ban_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.BanGroup))
	})
	return _c
}

func (_c *MockBanGroupRepository_Ban_Call) Return(_a0 error) *MockBanGroupRepository_Ban_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanGroupRepository_Ban_Call) RunAndReturn(run func(context.Context, *domain.BanGroup) error) *MockBanGroupRepository_Ban_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, banGroup
func (_m *MockBanGroupRepository) Delete(ctx context.Context, banGroup *domain.BanGroup) error {
	ret := _m.Called(ctx, banGroup)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BanGroup) error); ok {
		r0 = rf(ctx, banGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanGroupRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockBanGroupRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - banGroup *domain.BanGroup
func (_e *MockBanGroupRepository_Expecter) Delete(ctx interface{}, banGroup interface{}) *MockBanGroupRepository_Delete_Call {
	return &MockBanGroupRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, banGroup)}
}

func (_c *MockBanGroupRepository_Delete_Call) Run(run func(ctx context.Context, banGroup *domain.BanGroup)) *MockBanGroupRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.BanGroup))
	})
	return _c
}

func (_c *MockBanGroupRepository_Delete_Call) Return(_a0 error) *MockBanGroupRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanGroupRepository_Delete_Call) RunAndReturn(run func(context.Context, *domain.BanGroup) error) *MockBanGroupRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, filter
func (_m *MockBanGroupRepository) Get(ctx context.Context, filter domain.GroupBansQueryFilter) ([]domain.BannedGroupPerson, int64, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []domain.BannedGroupPerson
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GroupBansQueryFilter) ([]domain.BannedGroupPerson, int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GroupBansQueryFilter) []domain.BannedGroupPerson); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BannedGroupPerson)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GroupBansQueryFilter) int64); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.GroupBansQueryFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBanGroupRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockBanGroupRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - filter domain.GroupBansQueryFilter
func (_e *MockBanGroupRepository_Expecter) Get(ctx interface{}, filter interface{}) *MockBanGroupRepository_Get_Call {
	return &MockBanGroupRepository_Get_Call{Call: _e.mock.On("Get", ctx, filter)}
}

func (_c *MockBanGroupRepository_Get_Call) Run(run func(ctx context.Context, filter domain.GroupBansQueryFilter)) *MockBanGroupRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.GroupBansQueryFilter))
	})
	return _c
}

func (_c *MockBanGroupRepository_Get_Call) Return(_a0 []domain.BannedGroupPerson, _a1 int64, _a2 error) *MockBanGroupRepository_Get_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockBanGroupRepository_Get_Call) RunAndReturn(run func(context.Context, domain.GroupBansQueryFilter) ([]domain.BannedGroupPerson, int64, error)) *MockBanGroupRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByGID provides a mock function with given fields: ctx, groupID, banGroup
func (_m *MockBanGroupRepository) GetByGID(ctx context.Context, groupID steamid.GID, banGroup *domain.BanGroup) error {
	ret := _m.Called(ctx, groupID, banGroup)

	if len(ret) == 0 {
		panic("no return value specified for GetByGID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, steamid.GID, *domain.BanGroup) error); ok {
		r0 = rf(ctx, groupID, banGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanGroupRepository_GetByGID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByGID'
type MockBanGroupRepository_GetByGID_Call struct {
	*mock.Call
}

// GetByGID is a helper method to define mock.On call
//   - ctx context.Context
//   - groupID steamid.GID
//   - banGroup *domain.BanGroup
func (_e *MockBanGroupRepository_Expecter) GetByGID(ctx interface{}, groupID interface{}, banGroup interface{}) *MockBanGroupRepository_GetByGID_Call {
	return &MockBanGroupRepository_GetByGID_Call{Call: _e.mock.On("GetByGID", ctx, groupID, banGroup)}
}

func (_c *MockBanGroupRepository_GetByGID_Call) Run(run func(ctx context.Context, groupID steamid.GID, banGroup *domain.BanGroup)) *MockBanGroupRepository_GetByGID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(steamid.GID), args[2].(*domain.BanGroup))
	})
	return _c
}

func (_c *MockBanGroupRepository_GetByGID_Call) Return(_a0 error) *MockBanGroupRepository_GetByGID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanGroupRepository_GetByGID_Call) RunAndReturn(run func(context.Context, steamid.GID, *domain.BanGroup) error) *MockBanGroupRepository_GetByGID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, banGroupID, banGroup
func (_m *MockBanGroupRepository) GetByID(ctx context.Context, banGroupID int64, banGroup *domain.BanGroup) error {
	ret := _m.Called(ctx, banGroupID, banGroup)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.BanGroup) error); ok {
		r0 = rf(ctx, banGroupID, banGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanGroupRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockBanGroupRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - banGroupID int64
//   - banGroup *domain.BanGroup
func (_e *MockBanGroupRepository_Expecter) GetByID(ctx interface{}, banGroupID interface{}, banGroup interface{}) *MockBanGroupRepository_GetByID_Call {
	return &MockBanGroupRepository_GetByID_Call{Call: _e.mock.On("GetByID", ctx, banGroupID, banGroup)}
}

func (_c *MockBanGroupRepository_GetByID_Call) Run(run func(ctx context.Context, banGroupID int64, banGroup *domain.BanGroup)) *MockBanGroupRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*domain.BanGroup))
	})
	return _c
}

func (_c *MockBanGroupRepository_GetByID_Call) Return(_a0 error) *MockBanGroupRepository_GetByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanGroupRepository_GetByID_Call) RunAndReturn(run func(context.Context, int64, *domain.BanGroup) error) *MockBanGroupRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetMembersList provides a mock function with given fields: ctx, parentID, list
func (_m *MockBanGroupRepository) GetMembersList(ctx context.Context, parentID int64, list *domain.MembersList) error {
	ret := _m.Called(ctx, parentID, list)

	if len(ret) == 0 {
		panic("no return value specified for GetMembersList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.MembersList) error); ok {
		r0 = rf(ctx, parentID, list)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanGroupRepository_GetMembersList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMembersList'
type MockBanGroupRepository_GetMembersList_Call struct {
	*mock.Call
}

// GetMembersList is a helper method to define mock.On call
//   - ctx context.Context
//   - parentID int64
//   - list *domain.MembersList
func (_e *MockBanGroupRepository_Expecter) GetMembersList(ctx interface{}, parentID interface{}, list interface{}) *MockBanGroupRepository_GetMembersList_Call {
	return &MockBanGroupRepository_GetMembersList_Call{Call: _e.mock.On("GetMembersList", ctx, parentID, list)}
}

func (_c *MockBanGroupRepository_GetMembersList_Call) Run(run func(ctx context.Context, parentID int64, list *domain.MembersList)) *MockBanGroupRepository_GetMembersList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*domain.MembersList))
	})
	return _c
}

func (_c *MockBanGroupRepository_GetMembersList_Call) Return(_a0 error) *MockBanGroupRepository_GetMembersList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanGroupRepository_GetMembersList_Call) RunAndReturn(run func(context.Context, int64, *domain.MembersList) error) *MockBanGroupRepository_GetMembersList_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, banGroup
func (_m *MockBanGroupRepository) Save(ctx context.Context, banGroup *domain.BanGroup) error {
	ret := _m.Called(ctx, banGroup)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BanGroup) error); ok {
		r0 = rf(ctx, banGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanGroupRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockBanGroupRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - banGroup *domain.BanGroup
func (_e *MockBanGroupRepository_Expecter) Save(ctx interface{}, banGroup interface{}) *MockBanGroupRepository_Save_Call {
	return &MockBanGroupRepository_Save_Call{Call: _e.mock.On("Save", ctx, banGroup)}
}

func (_c *MockBanGroupRepository_Save_Call) Run(run func(ctx context.Context, banGroup *domain.BanGroup)) *MockBanGroupRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.BanGroup))
	})
	return _c
}

func (_c *MockBanGroupRepository_Save_Call) Return(_a0 error) *MockBanGroupRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanGroupRepository_Save_Call) RunAndReturn(run func(context.Context, *domain.BanGroup) error) *MockBanGroupRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// SaveMembersList provides a mock function with given fields: ctx, list
func (_m *MockBanGroupRepository) SaveMembersList(ctx context.Context, list *domain.MembersList) error {
	ret := _m.Called(ctx, list)

	if len(ret) == 0 {
		panic("no return value specified for SaveMembersList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.MembersList) error); ok {
		r0 = rf(ctx, list)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanGroupRepository_SaveMembersList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveMembersList'
type MockBanGroupRepository_SaveMembersList_Call struct {
	*mock.Call
}

// SaveMembersList is a helper method to define mock.On call
//   - ctx context.Context
//   - list *domain.MembersList
func (_e *MockBanGroupRepository_Expecter) SaveMembersList(ctx interface{}, list interface{}) *MockBanGroupRepository_SaveMembersList_Call {
	return &MockBanGroupRepository_SaveMembersList_Call{Call: _e.mock.On("SaveMembersList", ctx, list)}
}

func (_c *MockBanGroupRepository_SaveMembersList_Call) Run(run func(ctx context.Context, list *domain.MembersList)) *MockBanGroupRepository_SaveMembersList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.MembersList))
	})
	return _c
}

func (_c *MockBanGroupRepository_SaveMembersList_Call) Return(_a0 error) *MockBanGroupRepository_SaveMembersList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanGroupRepository_SaveMembersList_Call) RunAndReturn(run func(context.Context, *domain.MembersList) error) *MockBanGroupRepository_SaveMembersList_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBanGroupRepository creates a new instance of MockBanGroupRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBanGroupRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBanGroupRepository {
	mock := &MockBanGroupRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
