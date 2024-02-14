// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	domain "github.com/leighmacdonald/gbans/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockDemoUsecase is an autogenerated mock type for the DemoUsecase type
type MockDemoUsecase struct {
	mock.Mock
}

type MockDemoUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDemoUsecase) EXPECT() *MockDemoUsecase_Expecter {
	return &MockDemoUsecase_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, name, content, demoName, serverID
func (_m *MockDemoUsecase) Create(ctx context.Context, name string, content io.Reader, demoName string, serverID int) (*domain.DemoFile, error) {
	ret := _m.Called(ctx, name, content, demoName, serverID)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.DemoFile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, string, int) (*domain.DemoFile, error)); ok {
		return rf(ctx, name, content, demoName, serverID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, string, int) *domain.DemoFile); ok {
		r0 = rf(ctx, name, content, demoName, serverID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.DemoFile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, io.Reader, string, int) error); ok {
		r1 = rf(ctx, name, content, demoName, serverID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDemoUsecase_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockDemoUsecase_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - content io.Reader
//   - demoName string
//   - serverID int
func (_e *MockDemoUsecase_Expecter) Create(ctx interface{}, name interface{}, content interface{}, demoName interface{}, serverID interface{}) *MockDemoUsecase_Create_Call {
	return &MockDemoUsecase_Create_Call{Call: _e.mock.On("Create", ctx, name, content, demoName, serverID)}
}

func (_c *MockDemoUsecase_Create_Call) Run(run func(ctx context.Context, name string, content io.Reader, demoName string, serverID int)) *MockDemoUsecase_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(io.Reader), args[3].(string), args[4].(int))
	})
	return _c
}

func (_c *MockDemoUsecase_Create_Call) Return(_a0 *domain.DemoFile, _a1 error) *MockDemoUsecase_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDemoUsecase_Create_Call) RunAndReturn(run func(context.Context, string, io.Reader, string, int) (*domain.DemoFile, error)) *MockDemoUsecase_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DropDemo provides a mock function with given fields: ctx, demoFile
func (_m *MockDemoUsecase) DropDemo(ctx context.Context, demoFile *domain.DemoFile) error {
	ret := _m.Called(ctx, demoFile)

	if len(ret) == 0 {
		panic("no return value specified for DropDemo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.DemoFile) error); ok {
		r0 = rf(ctx, demoFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDemoUsecase_DropDemo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropDemo'
type MockDemoUsecase_DropDemo_Call struct {
	*mock.Call
}

// DropDemo is a helper method to define mock.On call
//   - ctx context.Context
//   - demoFile *domain.DemoFile
func (_e *MockDemoUsecase_Expecter) DropDemo(ctx interface{}, demoFile interface{}) *MockDemoUsecase_DropDemo_Call {
	return &MockDemoUsecase_DropDemo_Call{Call: _e.mock.On("DropDemo", ctx, demoFile)}
}

func (_c *MockDemoUsecase_DropDemo_Call) Run(run func(ctx context.Context, demoFile *domain.DemoFile)) *MockDemoUsecase_DropDemo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.DemoFile))
	})
	return _c
}

func (_c *MockDemoUsecase_DropDemo_Call) Return(_a0 error) *MockDemoUsecase_DropDemo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDemoUsecase_DropDemo_Call) RunAndReturn(run func(context.Context, *domain.DemoFile) error) *MockDemoUsecase_DropDemo_Call {
	_c.Call.Return(run)
	return _c
}

// ExpiredDemos provides a mock function with given fields: ctx, limit
func (_m *MockDemoUsecase) ExpiredDemos(ctx context.Context, limit uint64) ([]domain.DemoInfo, error) {
	ret := _m.Called(ctx, limit)

	if len(ret) == 0 {
		panic("no return value specified for ExpiredDemos")
	}

	var r0 []domain.DemoInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]domain.DemoInfo, error)); ok {
		return rf(ctx, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []domain.DemoInfo); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.DemoInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDemoUsecase_ExpiredDemos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpiredDemos'
type MockDemoUsecase_ExpiredDemos_Call struct {
	*mock.Call
}

// ExpiredDemos is a helper method to define mock.On call
//   - ctx context.Context
//   - limit uint64
func (_e *MockDemoUsecase_Expecter) ExpiredDemos(ctx interface{}, limit interface{}) *MockDemoUsecase_ExpiredDemos_Call {
	return &MockDemoUsecase_ExpiredDemos_Call{Call: _e.mock.On("ExpiredDemos", ctx, limit)}
}

func (_c *MockDemoUsecase_ExpiredDemos_Call) Run(run func(ctx context.Context, limit uint64)) *MockDemoUsecase_ExpiredDemos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockDemoUsecase_ExpiredDemos_Call) Return(_a0 []domain.DemoInfo, _a1 error) *MockDemoUsecase_ExpiredDemos_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDemoUsecase_ExpiredDemos_Call) RunAndReturn(run func(context.Context, uint64) ([]domain.DemoInfo, error)) *MockDemoUsecase_ExpiredDemos_Call {
	_c.Call.Return(run)
	return _c
}

// GetDemoByID provides a mock function with given fields: ctx, demoID, demoFile
func (_m *MockDemoUsecase) GetDemoByID(ctx context.Context, demoID int64, demoFile *domain.DemoFile) error {
	ret := _m.Called(ctx, demoID, demoFile)

	if len(ret) == 0 {
		panic("no return value specified for GetDemoByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.DemoFile) error); ok {
		r0 = rf(ctx, demoID, demoFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDemoUsecase_GetDemoByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDemoByID'
type MockDemoUsecase_GetDemoByID_Call struct {
	*mock.Call
}

// GetDemoByID is a helper method to define mock.On call
//   - ctx context.Context
//   - demoID int64
//   - demoFile *domain.DemoFile
func (_e *MockDemoUsecase_Expecter) GetDemoByID(ctx interface{}, demoID interface{}, demoFile interface{}) *MockDemoUsecase_GetDemoByID_Call {
	return &MockDemoUsecase_GetDemoByID_Call{Call: _e.mock.On("GetDemoByID", ctx, demoID, demoFile)}
}

func (_c *MockDemoUsecase_GetDemoByID_Call) Run(run func(ctx context.Context, demoID int64, demoFile *domain.DemoFile)) *MockDemoUsecase_GetDemoByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*domain.DemoFile))
	})
	return _c
}

func (_c *MockDemoUsecase_GetDemoByID_Call) Return(_a0 error) *MockDemoUsecase_GetDemoByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDemoUsecase_GetDemoByID_Call) RunAndReturn(run func(context.Context, int64, *domain.DemoFile) error) *MockDemoUsecase_GetDemoByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetDemoByName provides a mock function with given fields: ctx, demoName, demoFile
func (_m *MockDemoUsecase) GetDemoByName(ctx context.Context, demoName string, demoFile *domain.DemoFile) error {
	ret := _m.Called(ctx, demoName, demoFile)

	if len(ret) == 0 {
		panic("no return value specified for GetDemoByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.DemoFile) error); ok {
		r0 = rf(ctx, demoName, demoFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDemoUsecase_GetDemoByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDemoByName'
type MockDemoUsecase_GetDemoByName_Call struct {
	*mock.Call
}

// GetDemoByName is a helper method to define mock.On call
//   - ctx context.Context
//   - demoName string
//   - demoFile *domain.DemoFile
func (_e *MockDemoUsecase_Expecter) GetDemoByName(ctx interface{}, demoName interface{}, demoFile interface{}) *MockDemoUsecase_GetDemoByName_Call {
	return &MockDemoUsecase_GetDemoByName_Call{Call: _e.mock.On("GetDemoByName", ctx, demoName, demoFile)}
}

func (_c *MockDemoUsecase_GetDemoByName_Call) Run(run func(ctx context.Context, demoName string, demoFile *domain.DemoFile)) *MockDemoUsecase_GetDemoByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*domain.DemoFile))
	})
	return _c
}

func (_c *MockDemoUsecase_GetDemoByName_Call) Return(_a0 error) *MockDemoUsecase_GetDemoByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDemoUsecase_GetDemoByName_Call) RunAndReturn(run func(context.Context, string, *domain.DemoFile) error) *MockDemoUsecase_GetDemoByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetDemos provides a mock function with given fields: ctx, opts
func (_m *MockDemoUsecase) GetDemos(ctx context.Context, opts domain.DemoFilter) ([]domain.DemoFile, int64, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetDemos")
	}

	var r0 []domain.DemoFile
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.DemoFilter) ([]domain.DemoFile, int64, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.DemoFilter) []domain.DemoFile); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.DemoFile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.DemoFilter) int64); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.DemoFilter) error); ok {
		r2 = rf(ctx, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockDemoUsecase_GetDemos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDemos'
type MockDemoUsecase_GetDemos_Call struct {
	*mock.Call
}

// GetDemos is a helper method to define mock.On call
//   - ctx context.Context
//   - opts domain.DemoFilter
func (_e *MockDemoUsecase_Expecter) GetDemos(ctx interface{}, opts interface{}) *MockDemoUsecase_GetDemos_Call {
	return &MockDemoUsecase_GetDemos_Call{Call: _e.mock.On("GetDemos", ctx, opts)}
}

func (_c *MockDemoUsecase_GetDemos_Call) Run(run func(ctx context.Context, opts domain.DemoFilter)) *MockDemoUsecase_GetDemos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.DemoFilter))
	})
	return _c
}

func (_c *MockDemoUsecase_GetDemos_Call) Return(_a0 []domain.DemoFile, _a1 int64, _a2 error) *MockDemoUsecase_GetDemos_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockDemoUsecase_GetDemos_Call) RunAndReturn(run func(context.Context, domain.DemoFilter) ([]domain.DemoFile, int64, error)) *MockDemoUsecase_GetDemos_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *MockDemoUsecase) Start(ctx context.Context) {
	_m.Called(ctx)
}

// MockDemoUsecase_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockDemoUsecase_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDemoUsecase_Expecter) Start(ctx interface{}) *MockDemoUsecase_Start_Call {
	return &MockDemoUsecase_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *MockDemoUsecase_Start_Call) Run(run func(ctx context.Context)) *MockDemoUsecase_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDemoUsecase_Start_Call) Return() *MockDemoUsecase_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDemoUsecase_Start_Call) RunAndReturn(run func(context.Context)) *MockDemoUsecase_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDemoUsecase creates a new instance of MockDemoUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDemoUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDemoUsecase {
	mock := &MockDemoUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
