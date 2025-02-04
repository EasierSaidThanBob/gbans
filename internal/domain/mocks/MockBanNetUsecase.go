// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"

	netip "net/netip"
)

// MockBanNetUsecase is an autogenerated mock type for the BanNetUsecase type
type MockBanNetUsecase struct {
	mock.Mock
}

type MockBanNetUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBanNetUsecase) EXPECT() *MockBanNetUsecase_Expecter {
	return &MockBanNetUsecase_Expecter{mock: &_m.Mock}
}

// Ban provides a mock function with given fields: ctx, banNet
func (_m *MockBanNetUsecase) Ban(ctx context.Context, banNet *domain.BanCIDR) error {
	ret := _m.Called(ctx, banNet)

	if len(ret) == 0 {
		panic("no return value specified for Ban")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BanCIDR) error); ok {
		r0 = rf(ctx, banNet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanNetUsecase_Ban_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ban'
type MockBanNetUsecase_Ban_Call struct {
	*mock.Call
}

// Ban is a helper method to define mock.On call
//   - ctx context.Context
//   - banNet *domain.BanCIDR
func (_e *MockBanNetUsecase_Expecter) Ban(ctx interface{}, banNet interface{}) *MockBanNetUsecase_Ban_Call {
	return &MockBanNetUsecase_Ban_Call{Call: _e.mock.On("Ban", ctx, banNet)}
}

func (_c *MockBanNetUsecase_Ban_Call) Run(run func(ctx context.Context, banNet *domain.BanCIDR)) *MockBanNetUsecase_Ban_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.BanCIDR))
	})
	return _c
}

func (_c *MockBanNetUsecase_Ban_Call) Return(_a0 error) *MockBanNetUsecase_Ban_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanNetUsecase_Ban_Call) RunAndReturn(run func(context.Context, *domain.BanCIDR) error) *MockBanNetUsecase_Ban_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, banNet
func (_m *MockBanNetUsecase) Delete(ctx context.Context, banNet *domain.BanCIDR) error {
	ret := _m.Called(ctx, banNet)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BanCIDR) error); ok {
		r0 = rf(ctx, banNet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanNetUsecase_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockBanNetUsecase_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - banNet *domain.BanCIDR
func (_e *MockBanNetUsecase_Expecter) Delete(ctx interface{}, banNet interface{}) *MockBanNetUsecase_Delete_Call {
	return &MockBanNetUsecase_Delete_Call{Call: _e.mock.On("Delete", ctx, banNet)}
}

func (_c *MockBanNetUsecase_Delete_Call) Run(run func(ctx context.Context, banNet *domain.BanCIDR)) *MockBanNetUsecase_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.BanCIDR))
	})
	return _c
}

func (_c *MockBanNetUsecase_Delete_Call) Return(_a0 error) *MockBanNetUsecase_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanNetUsecase_Delete_Call) RunAndReturn(run func(context.Context, *domain.BanCIDR) error) *MockBanNetUsecase_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Expired provides a mock function with given fields: ctx
func (_m *MockBanNetUsecase) Expired(ctx context.Context) ([]domain.BanCIDR, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Expired")
	}

	var r0 []domain.BanCIDR
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.BanCIDR, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.BanCIDR); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BanCIDR)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanNetUsecase_Expired_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Expired'
type MockBanNetUsecase_Expired_Call struct {
	*mock.Call
}

// Expired is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBanNetUsecase_Expecter) Expired(ctx interface{}) *MockBanNetUsecase_Expired_Call {
	return &MockBanNetUsecase_Expired_Call{Call: _e.mock.On("Expired", ctx)}
}

func (_c *MockBanNetUsecase_Expired_Call) Run(run func(ctx context.Context)) *MockBanNetUsecase_Expired_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBanNetUsecase_Expired_Call) Return(_a0 []domain.BanCIDR, _a1 error) *MockBanNetUsecase_Expired_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanNetUsecase_Expired_Call) RunAndReturn(run func(context.Context) ([]domain.BanCIDR, error)) *MockBanNetUsecase_Expired_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, filter
func (_m *MockBanNetUsecase) Get(ctx context.Context, filter domain.CIDRBansQueryFilter) ([]domain.BannedCIDRPerson, int64, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []domain.BannedCIDRPerson
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CIDRBansQueryFilter) ([]domain.BannedCIDRPerson, int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CIDRBansQueryFilter) []domain.BannedCIDRPerson); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BannedCIDRPerson)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CIDRBansQueryFilter) int64); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.CIDRBansQueryFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBanNetUsecase_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockBanNetUsecase_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - filter domain.CIDRBansQueryFilter
func (_e *MockBanNetUsecase_Expecter) Get(ctx interface{}, filter interface{}) *MockBanNetUsecase_Get_Call {
	return &MockBanNetUsecase_Get_Call{Call: _e.mock.On("Get", ctx, filter)}
}

func (_c *MockBanNetUsecase_Get_Call) Run(run func(ctx context.Context, filter domain.CIDRBansQueryFilter)) *MockBanNetUsecase_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.CIDRBansQueryFilter))
	})
	return _c
}

func (_c *MockBanNetUsecase_Get_Call) Return(_a0 []domain.BannedCIDRPerson, _a1 int64, _a2 error) *MockBanNetUsecase_Get_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockBanNetUsecase_Get_Call) RunAndReturn(run func(context.Context, domain.CIDRBansQueryFilter) ([]domain.BannedCIDRPerson, int64, error)) *MockBanNetUsecase_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByAddress provides a mock function with given fields: ctx, ipAddr
func (_m *MockBanNetUsecase) GetByAddress(ctx context.Context, ipAddr netip.Addr) ([]domain.BanCIDR, error) {
	ret := _m.Called(ctx, ipAddr)

	if len(ret) == 0 {
		panic("no return value specified for GetByAddress")
	}

	var r0 []domain.BanCIDR
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) ([]domain.BanCIDR, error)); ok {
		return rf(ctx, ipAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, netip.Addr) []domain.BanCIDR); ok {
		r0 = rf(ctx, ipAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BanCIDR)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, netip.Addr) error); ok {
		r1 = rf(ctx, ipAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBanNetUsecase_GetByAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByAddress'
type MockBanNetUsecase_GetByAddress_Call struct {
	*mock.Call
}

// GetByAddress is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAddr netip.Addr
func (_e *MockBanNetUsecase_Expecter) GetByAddress(ctx interface{}, ipAddr interface{}) *MockBanNetUsecase_GetByAddress_Call {
	return &MockBanNetUsecase_GetByAddress_Call{Call: _e.mock.On("GetByAddress", ctx, ipAddr)}
}

func (_c *MockBanNetUsecase_GetByAddress_Call) Run(run func(ctx context.Context, ipAddr netip.Addr)) *MockBanNetUsecase_GetByAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(netip.Addr))
	})
	return _c
}

func (_c *MockBanNetUsecase_GetByAddress_Call) Return(_a0 []domain.BanCIDR, _a1 error) *MockBanNetUsecase_GetByAddress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBanNetUsecase_GetByAddress_Call) RunAndReturn(run func(context.Context, netip.Addr) ([]domain.BanCIDR, error)) *MockBanNetUsecase_GetByAddress_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, netID, banNet
func (_m *MockBanNetUsecase) GetByID(ctx context.Context, netID int64, banNet *domain.BanCIDR) error {
	ret := _m.Called(ctx, netID, banNet)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.BanCIDR) error); ok {
		r0 = rf(ctx, netID, banNet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanNetUsecase_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockBanNetUsecase_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - netID int64
//   - banNet *domain.BanCIDR
func (_e *MockBanNetUsecase_Expecter) GetByID(ctx interface{}, netID interface{}, banNet interface{}) *MockBanNetUsecase_GetByID_Call {
	return &MockBanNetUsecase_GetByID_Call{Call: _e.mock.On("GetByID", ctx, netID, banNet)}
}

func (_c *MockBanNetUsecase_GetByID_Call) Run(run func(ctx context.Context, netID int64, banNet *domain.BanCIDR)) *MockBanNetUsecase_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*domain.BanCIDR))
	})
	return _c
}

func (_c *MockBanNetUsecase_GetByID_Call) Return(_a0 error) *MockBanNetUsecase_GetByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanNetUsecase_GetByID_Call) RunAndReturn(run func(context.Context, int64, *domain.BanCIDR) error) *MockBanNetUsecase_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, banNet
func (_m *MockBanNetUsecase) Save(ctx context.Context, banNet *domain.BanCIDR) error {
	ret := _m.Called(ctx, banNet)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BanCIDR) error); ok {
		r0 = rf(ctx, banNet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBanNetUsecase_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockBanNetUsecase_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - banNet *domain.BanCIDR
func (_e *MockBanNetUsecase_Expecter) Save(ctx interface{}, banNet interface{}) *MockBanNetUsecase_Save_Call {
	return &MockBanNetUsecase_Save_Call{Call: _e.mock.On("Save", ctx, banNet)}
}

func (_c *MockBanNetUsecase_Save_Call) Run(run func(ctx context.Context, banNet *domain.BanCIDR)) *MockBanNetUsecase_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.BanCIDR))
	})
	return _c
}

func (_c *MockBanNetUsecase_Save_Call) Return(_a0 error) *MockBanNetUsecase_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBanNetUsecase_Save_Call) RunAndReturn(run func(context.Context, *domain.BanCIDR) error) *MockBanNetUsecase_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBanNetUsecase creates a new instance of MockBanNetUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBanNetUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBanNetUsecase {
	mock := &MockBanNetUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
