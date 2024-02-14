// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/gofrs/uuid/v5"
)

// MockMediaRepository is an autogenerated mock type for the MediaRepository type
type MockMediaRepository struct {
	mock.Mock
}

type MockMediaRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMediaRepository) EXPECT() *MockMediaRepository_Expecter {
	return &MockMediaRepository_Expecter{mock: &_m.Mock}
}

// GetMediaByAssetID provides a mock function with given fields: ctx, _a1, media
func (_m *MockMediaRepository) GetMediaByAssetID(ctx context.Context, _a1 uuid.UUID, media *domain.Media) error {
	ret := _m.Called(ctx, _a1, media)

	if len(ret) == 0 {
		panic("no return value specified for GetMediaByAssetID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *domain.Media) error); ok {
		r0 = rf(ctx, _a1, media)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaRepository_GetMediaByAssetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMediaByAssetID'
type MockMediaRepository_GetMediaByAssetID_Call struct {
	*mock.Call
}

// GetMediaByAssetID is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 uuid.UUID
//   - media *domain.Media
func (_e *MockMediaRepository_Expecter) GetMediaByAssetID(ctx interface{}, _a1 interface{}, media interface{}) *MockMediaRepository_GetMediaByAssetID_Call {
	return &MockMediaRepository_GetMediaByAssetID_Call{Call: _e.mock.On("GetMediaByAssetID", ctx, _a1, media)}
}

func (_c *MockMediaRepository_GetMediaByAssetID_Call) Run(run func(ctx context.Context, _a1 uuid.UUID, media *domain.Media)) *MockMediaRepository_GetMediaByAssetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(*domain.Media))
	})
	return _c
}

func (_c *MockMediaRepository_GetMediaByAssetID_Call) Return(_a0 error) *MockMediaRepository_GetMediaByAssetID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaRepository_GetMediaByAssetID_Call) RunAndReturn(run func(context.Context, uuid.UUID, *domain.Media) error) *MockMediaRepository_GetMediaByAssetID_Call {
	_c.Call.Return(run)
	return _c
}

// GetMediaByID provides a mock function with given fields: ctx, mediaID, media
func (_m *MockMediaRepository) GetMediaByID(ctx context.Context, mediaID int, media *domain.Media) error {
	ret := _m.Called(ctx, mediaID, media)

	if len(ret) == 0 {
		panic("no return value specified for GetMediaByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *domain.Media) error); ok {
		r0 = rf(ctx, mediaID, media)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaRepository_GetMediaByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMediaByID'
type MockMediaRepository_GetMediaByID_Call struct {
	*mock.Call
}

// GetMediaByID is a helper method to define mock.On call
//   - ctx context.Context
//   - mediaID int
//   - media *domain.Media
func (_e *MockMediaRepository_Expecter) GetMediaByID(ctx interface{}, mediaID interface{}, media interface{}) *MockMediaRepository_GetMediaByID_Call {
	return &MockMediaRepository_GetMediaByID_Call{Call: _e.mock.On("GetMediaByID", ctx, mediaID, media)}
}

func (_c *MockMediaRepository_GetMediaByID_Call) Run(run func(ctx context.Context, mediaID int, media *domain.Media)) *MockMediaRepository_GetMediaByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(*domain.Media))
	})
	return _c
}

func (_c *MockMediaRepository_GetMediaByID_Call) Return(_a0 error) *MockMediaRepository_GetMediaByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaRepository_GetMediaByID_Call) RunAndReturn(run func(context.Context, int, *domain.Media) error) *MockMediaRepository_GetMediaByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetMediaByName provides a mock function with given fields: ctx, name, media
func (_m *MockMediaRepository) GetMediaByName(ctx context.Context, name string, media *domain.Media) error {
	ret := _m.Called(ctx, name, media)

	if len(ret) == 0 {
		panic("no return value specified for GetMediaByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.Media) error); ok {
		r0 = rf(ctx, name, media)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaRepository_GetMediaByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMediaByName'
type MockMediaRepository_GetMediaByName_Call struct {
	*mock.Call
}

// GetMediaByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - media *domain.Media
func (_e *MockMediaRepository_Expecter) GetMediaByName(ctx interface{}, name interface{}, media interface{}) *MockMediaRepository_GetMediaByName_Call {
	return &MockMediaRepository_GetMediaByName_Call{Call: _e.mock.On("GetMediaByName", ctx, name, media)}
}

func (_c *MockMediaRepository_GetMediaByName_Call) Run(run func(ctx context.Context, name string, media *domain.Media)) *MockMediaRepository_GetMediaByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*domain.Media))
	})
	return _c
}

func (_c *MockMediaRepository_GetMediaByName_Call) Return(_a0 error) *MockMediaRepository_GetMediaByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaRepository_GetMediaByName_Call) RunAndReturn(run func(context.Context, string, *domain.Media) error) *MockMediaRepository_GetMediaByName_Call {
	_c.Call.Return(run)
	return _c
}

// SaveMedia provides a mock function with given fields: ctx, media
func (_m *MockMediaRepository) SaveMedia(ctx context.Context, media *domain.Media) error {
	ret := _m.Called(ctx, media)

	if len(ret) == 0 {
		panic("no return value specified for SaveMedia")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Media) error); ok {
		r0 = rf(ctx, media)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMediaRepository_SaveMedia_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveMedia'
type MockMediaRepository_SaveMedia_Call struct {
	*mock.Call
}

// SaveMedia is a helper method to define mock.On call
//   - ctx context.Context
//   - media *domain.Media
func (_e *MockMediaRepository_Expecter) SaveMedia(ctx interface{}, media interface{}) *MockMediaRepository_SaveMedia_Call {
	return &MockMediaRepository_SaveMedia_Call{Call: _e.mock.On("SaveMedia", ctx, media)}
}

func (_c *MockMediaRepository_SaveMedia_Call) Run(run func(ctx context.Context, media *domain.Media)) *MockMediaRepository_SaveMedia_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Media))
	})
	return _c
}

func (_c *MockMediaRepository_SaveMedia_Call) Return(_a0 error) *MockMediaRepository_SaveMedia_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMediaRepository_SaveMedia_Call) RunAndReturn(run func(context.Context, *domain.Media) error) *MockMediaRepository_SaveMedia_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMediaRepository creates a new instance of MockMediaRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMediaRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMediaRepository {
	mock := &MockMediaRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
