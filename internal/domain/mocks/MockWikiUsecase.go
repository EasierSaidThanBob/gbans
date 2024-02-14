// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockWikiUsecase is an autogenerated mock type for the WikiUsecase type
type MockWikiUsecase struct {
	mock.Mock
}

type MockWikiUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWikiUsecase) EXPECT() *MockWikiUsecase_Expecter {
	return &MockWikiUsecase_Expecter{mock: &_m.Mock}
}

// DeleteWikiPageBySlug provides a mock function with given fields: ctx, slug
func (_m *MockWikiUsecase) DeleteWikiPageBySlug(ctx context.Context, slug string) error {
	ret := _m.Called(ctx, slug)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWikiPageBySlug")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, slug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWikiUsecase_DeleteWikiPageBySlug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWikiPageBySlug'
type MockWikiUsecase_DeleteWikiPageBySlug_Call struct {
	*mock.Call
}

// DeleteWikiPageBySlug is a helper method to define mock.On call
//   - ctx context.Context
//   - slug string
func (_e *MockWikiUsecase_Expecter) DeleteWikiPageBySlug(ctx interface{}, slug interface{}) *MockWikiUsecase_DeleteWikiPageBySlug_Call {
	return &MockWikiUsecase_DeleteWikiPageBySlug_Call{Call: _e.mock.On("DeleteWikiPageBySlug", ctx, slug)}
}

func (_c *MockWikiUsecase_DeleteWikiPageBySlug_Call) Run(run func(ctx context.Context, slug string)) *MockWikiUsecase_DeleteWikiPageBySlug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockWikiUsecase_DeleteWikiPageBySlug_Call) Return(_a0 error) *MockWikiUsecase_DeleteWikiPageBySlug_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWikiUsecase_DeleteWikiPageBySlug_Call) RunAndReturn(run func(context.Context, string) error) *MockWikiUsecase_DeleteWikiPageBySlug_Call {
	_c.Call.Return(run)
	return _c
}

// GetWikiPageBySlug provides a mock function with given fields: ctx, user, slug
func (_m *MockWikiUsecase) GetWikiPageBySlug(ctx context.Context, user domain.PersonInfo, slug string) (domain.WikiPage, error) {
	ret := _m.Called(ctx, user, slug)

	if len(ret) == 0 {
		panic("no return value specified for GetWikiPageBySlug")
	}

	var r0 domain.WikiPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, string) (domain.WikiPage, error)); ok {
		return rf(ctx, user, slug)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, string) domain.WikiPage); ok {
		r0 = rf(ctx, user, slug)
	} else {
		r0 = ret.Get(0).(domain.WikiPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PersonInfo, string) error); ok {
		r1 = rf(ctx, user, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWikiUsecase_GetWikiPageBySlug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWikiPageBySlug'
type MockWikiUsecase_GetWikiPageBySlug_Call struct {
	*mock.Call
}

// GetWikiPageBySlug is a helper method to define mock.On call
//   - ctx context.Context
//   - user domain.PersonInfo
//   - slug string
func (_e *MockWikiUsecase_Expecter) GetWikiPageBySlug(ctx interface{}, user interface{}, slug interface{}) *MockWikiUsecase_GetWikiPageBySlug_Call {
	return &MockWikiUsecase_GetWikiPageBySlug_Call{Call: _e.mock.On("GetWikiPageBySlug", ctx, user, slug)}
}

func (_c *MockWikiUsecase_GetWikiPageBySlug_Call) Run(run func(ctx context.Context, user domain.PersonInfo, slug string)) *MockWikiUsecase_GetWikiPageBySlug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PersonInfo), args[2].(string))
	})
	return _c
}

func (_c *MockWikiUsecase_GetWikiPageBySlug_Call) Return(_a0 domain.WikiPage, _a1 error) *MockWikiUsecase_GetWikiPageBySlug_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWikiUsecase_GetWikiPageBySlug_Call) RunAndReturn(run func(context.Context, domain.PersonInfo, string) (domain.WikiPage, error)) *MockWikiUsecase_GetWikiPageBySlug_Call {
	_c.Call.Return(run)
	return _c
}

// SaveWikiPage provides a mock function with given fields: ctx, user, slug, body, level
func (_m *MockWikiUsecase) SaveWikiPage(ctx context.Context, user domain.PersonInfo, slug string, body string, level domain.Privilege) (domain.WikiPage, error) {
	ret := _m.Called(ctx, user, slug, body, level)

	if len(ret) == 0 {
		panic("no return value specified for SaveWikiPage")
	}

	var r0 domain.WikiPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, string, string, domain.Privilege) (domain.WikiPage, error)); ok {
		return rf(ctx, user, slug, body, level)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, string, string, domain.Privilege) domain.WikiPage); ok {
		r0 = rf(ctx, user, slug, body, level)
	} else {
		r0 = ret.Get(0).(domain.WikiPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PersonInfo, string, string, domain.Privilege) error); ok {
		r1 = rf(ctx, user, slug, body, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWikiUsecase_SaveWikiPage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveWikiPage'
type MockWikiUsecase_SaveWikiPage_Call struct {
	*mock.Call
}

// SaveWikiPage is a helper method to define mock.On call
//   - ctx context.Context
//   - user domain.PersonInfo
//   - slug string
//   - body string
//   - level domain.Privilege
func (_e *MockWikiUsecase_Expecter) SaveWikiPage(ctx interface{}, user interface{}, slug interface{}, body interface{}, level interface{}) *MockWikiUsecase_SaveWikiPage_Call {
	return &MockWikiUsecase_SaveWikiPage_Call{Call: _e.mock.On("SaveWikiPage", ctx, user, slug, body, level)}
}

func (_c *MockWikiUsecase_SaveWikiPage_Call) Run(run func(ctx context.Context, user domain.PersonInfo, slug string, body string, level domain.Privilege)) *MockWikiUsecase_SaveWikiPage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PersonInfo), args[2].(string), args[3].(string), args[4].(domain.Privilege))
	})
	return _c
}

func (_c *MockWikiUsecase_SaveWikiPage_Call) Return(_a0 domain.WikiPage, _a1 error) *MockWikiUsecase_SaveWikiPage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWikiUsecase_SaveWikiPage_Call) RunAndReturn(run func(context.Context, domain.PersonInfo, string, string, domain.Privilege) (domain.WikiPage, error)) *MockWikiUsecase_SaveWikiPage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWikiUsecase creates a new instance of MockWikiUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWikiUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWikiUsecase {
	mock := &MockWikiUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
