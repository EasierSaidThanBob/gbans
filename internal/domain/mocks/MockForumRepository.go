// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/leighmacdonald/gbans/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockForumRepository is an autogenerated mock type for the ForumRepository type
type MockForumRepository struct {
	mock.Mock
}

type MockForumRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockForumRepository) EXPECT() *MockForumRepository_Expecter {
	return &MockForumRepository_Expecter{mock: &_m.Mock}
}

// Forum provides a mock function with given fields: ctx, forumID, forum
func (_m *MockForumRepository) Forum(ctx context.Context, forumID int, forum *domain.Forum) error {
	ret := _m.Called(ctx, forumID, forum)

	if len(ret) == 0 {
		panic("no return value specified for Forum")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *domain.Forum) error); ok {
		r0 = rf(ctx, forumID, forum)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_Forum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Forum'
type MockForumRepository_Forum_Call struct {
	*mock.Call
}

// Forum is a helper method to define mock.On call
//   - ctx context.Context
//   - forumID int
//   - forum *domain.Forum
func (_e *MockForumRepository_Expecter) Forum(ctx interface{}, forumID interface{}, forum interface{}) *MockForumRepository_Forum_Call {
	return &MockForumRepository_Forum_Call{Call: _e.mock.On("Forum", ctx, forumID, forum)}
}

func (_c *MockForumRepository_Forum_Call) Run(run func(ctx context.Context, forumID int, forum *domain.Forum)) *MockForumRepository_Forum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(*domain.Forum))
	})
	return _c
}

func (_c *MockForumRepository_Forum_Call) Return(_a0 error) *MockForumRepository_Forum_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_Forum_Call) RunAndReturn(run func(context.Context, int, *domain.Forum) error) *MockForumRepository_Forum_Call {
	_c.Call.Return(run)
	return _c
}

// ForumCategories provides a mock function with given fields: ctx
func (_m *MockForumRepository) ForumCategories(ctx context.Context) ([]domain.ForumCategory, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ForumCategories")
	}

	var r0 []domain.ForumCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.ForumCategory, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.ForumCategory); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ForumCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockForumRepository_ForumCategories_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumCategories'
type MockForumRepository_ForumCategories_Call struct {
	*mock.Call
}

// ForumCategories is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockForumRepository_Expecter) ForumCategories(ctx interface{}) *MockForumRepository_ForumCategories_Call {
	return &MockForumRepository_ForumCategories_Call{Call: _e.mock.On("ForumCategories", ctx)}
}

func (_c *MockForumRepository_ForumCategories_Call) Run(run func(ctx context.Context)) *MockForumRepository_ForumCategories_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockForumRepository_ForumCategories_Call) Return(_a0 []domain.ForumCategory, _a1 error) *MockForumRepository_ForumCategories_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockForumRepository_ForumCategories_Call) RunAndReturn(run func(context.Context) ([]domain.ForumCategory, error)) *MockForumRepository_ForumCategories_Call {
	_c.Call.Return(run)
	return _c
}

// ForumCategory provides a mock function with given fields: ctx, categoryID, category
func (_m *MockForumRepository) ForumCategory(ctx context.Context, categoryID int, category *domain.ForumCategory) error {
	ret := _m.Called(ctx, categoryID, category)

	if len(ret) == 0 {
		panic("no return value specified for ForumCategory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *domain.ForumCategory) error); ok {
		r0 = rf(ctx, categoryID, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumCategory'
type MockForumRepository_ForumCategory_Call struct {
	*mock.Call
}

// ForumCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - categoryID int
//   - category *domain.ForumCategory
func (_e *MockForumRepository_Expecter) ForumCategory(ctx interface{}, categoryID interface{}, category interface{}) *MockForumRepository_ForumCategory_Call {
	return &MockForumRepository_ForumCategory_Call{Call: _e.mock.On("ForumCategory", ctx, categoryID, category)}
}

func (_c *MockForumRepository_ForumCategory_Call) Run(run func(ctx context.Context, categoryID int, category *domain.ForumCategory)) *MockForumRepository_ForumCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(*domain.ForumCategory))
	})
	return _c
}

func (_c *MockForumRepository_ForumCategory_Call) Return(_a0 error) *MockForumRepository_ForumCategory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumCategory_Call) RunAndReturn(run func(context.Context, int, *domain.ForumCategory) error) *MockForumRepository_ForumCategory_Call {
	_c.Call.Return(run)
	return _c
}

// ForumCategoryDelete provides a mock function with given fields: ctx, categoryID
func (_m *MockForumRepository) ForumCategoryDelete(ctx context.Context, categoryID int) error {
	ret := _m.Called(ctx, categoryID)

	if len(ret) == 0 {
		panic("no return value specified for ForumCategoryDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, categoryID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumCategoryDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumCategoryDelete'
type MockForumRepository_ForumCategoryDelete_Call struct {
	*mock.Call
}

// ForumCategoryDelete is a helper method to define mock.On call
//   - ctx context.Context
//   - categoryID int
func (_e *MockForumRepository_Expecter) ForumCategoryDelete(ctx interface{}, categoryID interface{}) *MockForumRepository_ForumCategoryDelete_Call {
	return &MockForumRepository_ForumCategoryDelete_Call{Call: _e.mock.On("ForumCategoryDelete", ctx, categoryID)}
}

func (_c *MockForumRepository_ForumCategoryDelete_Call) Run(run func(ctx context.Context, categoryID int)) *MockForumRepository_ForumCategoryDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockForumRepository_ForumCategoryDelete_Call) Return(_a0 error) *MockForumRepository_ForumCategoryDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumCategoryDelete_Call) RunAndReturn(run func(context.Context, int) error) *MockForumRepository_ForumCategoryDelete_Call {
	_c.Call.Return(run)
	return _c
}

// ForumCategorySave provides a mock function with given fields: ctx, category
func (_m *MockForumRepository) ForumCategorySave(ctx context.Context, category *domain.ForumCategory) error {
	ret := _m.Called(ctx, category)

	if len(ret) == 0 {
		panic("no return value specified for ForumCategorySave")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ForumCategory) error); ok {
		r0 = rf(ctx, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumCategorySave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumCategorySave'
type MockForumRepository_ForumCategorySave_Call struct {
	*mock.Call
}

// ForumCategorySave is a helper method to define mock.On call
//   - ctx context.Context
//   - category *domain.ForumCategory
func (_e *MockForumRepository_Expecter) ForumCategorySave(ctx interface{}, category interface{}) *MockForumRepository_ForumCategorySave_Call {
	return &MockForumRepository_ForumCategorySave_Call{Call: _e.mock.On("ForumCategorySave", ctx, category)}
}

func (_c *MockForumRepository_ForumCategorySave_Call) Run(run func(ctx context.Context, category *domain.ForumCategory)) *MockForumRepository_ForumCategorySave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.ForumCategory))
	})
	return _c
}

func (_c *MockForumRepository_ForumCategorySave_Call) Return(_a0 error) *MockForumRepository_ForumCategorySave_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumCategorySave_Call) RunAndReturn(run func(context.Context, *domain.ForumCategory) error) *MockForumRepository_ForumCategorySave_Call {
	_c.Call.Return(run)
	return _c
}

// ForumDelete provides a mock function with given fields: ctx, forumID
func (_m *MockForumRepository) ForumDelete(ctx context.Context, forumID int) error {
	ret := _m.Called(ctx, forumID)

	if len(ret) == 0 {
		panic("no return value specified for ForumDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, forumID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumDelete'
type MockForumRepository_ForumDelete_Call struct {
	*mock.Call
}

// ForumDelete is a helper method to define mock.On call
//   - ctx context.Context
//   - forumID int
func (_e *MockForumRepository_Expecter) ForumDelete(ctx interface{}, forumID interface{}) *MockForumRepository_ForumDelete_Call {
	return &MockForumRepository_ForumDelete_Call{Call: _e.mock.On("ForumDelete", ctx, forumID)}
}

func (_c *MockForumRepository_ForumDelete_Call) Run(run func(ctx context.Context, forumID int)) *MockForumRepository_ForumDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockForumRepository_ForumDelete_Call) Return(_a0 error) *MockForumRepository_ForumDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumDelete_Call) RunAndReturn(run func(context.Context, int) error) *MockForumRepository_ForumDelete_Call {
	_c.Call.Return(run)
	return _c
}

// ForumIncrMessageCount provides a mock function with given fields: ctx, forumID, incr
func (_m *MockForumRepository) ForumIncrMessageCount(ctx context.Context, forumID int, incr bool) error {
	ret := _m.Called(ctx, forumID, incr)

	if len(ret) == 0 {
		panic("no return value specified for ForumIncrMessageCount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, bool) error); ok {
		r0 = rf(ctx, forumID, incr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumIncrMessageCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumIncrMessageCount'
type MockForumRepository_ForumIncrMessageCount_Call struct {
	*mock.Call
}

// ForumIncrMessageCount is a helper method to define mock.On call
//   - ctx context.Context
//   - forumID int
//   - incr bool
func (_e *MockForumRepository_Expecter) ForumIncrMessageCount(ctx interface{}, forumID interface{}, incr interface{}) *MockForumRepository_ForumIncrMessageCount_Call {
	return &MockForumRepository_ForumIncrMessageCount_Call{Call: _e.mock.On("ForumIncrMessageCount", ctx, forumID, incr)}
}

func (_c *MockForumRepository_ForumIncrMessageCount_Call) Run(run func(ctx context.Context, forumID int, incr bool)) *MockForumRepository_ForumIncrMessageCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(bool))
	})
	return _c
}

func (_c *MockForumRepository_ForumIncrMessageCount_Call) Return(_a0 error) *MockForumRepository_ForumIncrMessageCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumIncrMessageCount_Call) RunAndReturn(run func(context.Context, int, bool) error) *MockForumRepository_ForumIncrMessageCount_Call {
	_c.Call.Return(run)
	return _c
}

// ForumMessage provides a mock function with given fields: ctx, messageID, forumMessage
func (_m *MockForumRepository) ForumMessage(ctx context.Context, messageID int64, forumMessage *domain.ForumMessage) error {
	ret := _m.Called(ctx, messageID, forumMessage)

	if len(ret) == 0 {
		panic("no return value specified for ForumMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.ForumMessage) error); ok {
		r0 = rf(ctx, messageID, forumMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumMessage'
type MockForumRepository_ForumMessage_Call struct {
	*mock.Call
}

// ForumMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - messageID int64
//   - forumMessage *domain.ForumMessage
func (_e *MockForumRepository_Expecter) ForumMessage(ctx interface{}, messageID interface{}, forumMessage interface{}) *MockForumRepository_ForumMessage_Call {
	return &MockForumRepository_ForumMessage_Call{Call: _e.mock.On("ForumMessage", ctx, messageID, forumMessage)}
}

func (_c *MockForumRepository_ForumMessage_Call) Run(run func(ctx context.Context, messageID int64, forumMessage *domain.ForumMessage)) *MockForumRepository_ForumMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*domain.ForumMessage))
	})
	return _c
}

func (_c *MockForumRepository_ForumMessage_Call) Return(_a0 error) *MockForumRepository_ForumMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumMessage_Call) RunAndReturn(run func(context.Context, int64, *domain.ForumMessage) error) *MockForumRepository_ForumMessage_Call {
	_c.Call.Return(run)
	return _c
}

// ForumMessageDelete provides a mock function with given fields: ctx, messageID
func (_m *MockForumRepository) ForumMessageDelete(ctx context.Context, messageID int64) error {
	ret := _m.Called(ctx, messageID)

	if len(ret) == 0 {
		panic("no return value specified for ForumMessageDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, messageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumMessageDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumMessageDelete'
type MockForumRepository_ForumMessageDelete_Call struct {
	*mock.Call
}

// ForumMessageDelete is a helper method to define mock.On call
//   - ctx context.Context
//   - messageID int64
func (_e *MockForumRepository_Expecter) ForumMessageDelete(ctx interface{}, messageID interface{}) *MockForumRepository_ForumMessageDelete_Call {
	return &MockForumRepository_ForumMessageDelete_Call{Call: _e.mock.On("ForumMessageDelete", ctx, messageID)}
}

func (_c *MockForumRepository_ForumMessageDelete_Call) Run(run func(ctx context.Context, messageID int64)) *MockForumRepository_ForumMessageDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockForumRepository_ForumMessageDelete_Call) Return(_a0 error) *MockForumRepository_ForumMessageDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumMessageDelete_Call) RunAndReturn(run func(context.Context, int64) error) *MockForumRepository_ForumMessageDelete_Call {
	_c.Call.Return(run)
	return _c
}

// ForumMessageSave provides a mock function with given fields: ctx, message
func (_m *MockForumRepository) ForumMessageSave(ctx context.Context, message *domain.ForumMessage) error {
	ret := _m.Called(ctx, message)

	if len(ret) == 0 {
		panic("no return value specified for ForumMessageSave")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ForumMessage) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumMessageSave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumMessageSave'
type MockForumRepository_ForumMessageSave_Call struct {
	*mock.Call
}

// ForumMessageSave is a helper method to define mock.On call
//   - ctx context.Context
//   - message *domain.ForumMessage
func (_e *MockForumRepository_Expecter) ForumMessageSave(ctx interface{}, message interface{}) *MockForumRepository_ForumMessageSave_Call {
	return &MockForumRepository_ForumMessageSave_Call{Call: _e.mock.On("ForumMessageSave", ctx, message)}
}

func (_c *MockForumRepository_ForumMessageSave_Call) Run(run func(ctx context.Context, message *domain.ForumMessage)) *MockForumRepository_ForumMessageSave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.ForumMessage))
	})
	return _c
}

func (_c *MockForumRepository_ForumMessageSave_Call) Return(_a0 error) *MockForumRepository_ForumMessageSave_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumMessageSave_Call) RunAndReturn(run func(context.Context, *domain.ForumMessage) error) *MockForumRepository_ForumMessageSave_Call {
	_c.Call.Return(run)
	return _c
}

// ForumMessageVoteApply provides a mock function with given fields: ctx, messageVote
func (_m *MockForumRepository) ForumMessageVoteApply(ctx context.Context, messageVote *domain.ForumMessageVote) error {
	ret := _m.Called(ctx, messageVote)

	if len(ret) == 0 {
		panic("no return value specified for ForumMessageVoteApply")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ForumMessageVote) error); ok {
		r0 = rf(ctx, messageVote)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumMessageVoteApply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumMessageVoteApply'
type MockForumRepository_ForumMessageVoteApply_Call struct {
	*mock.Call
}

// ForumMessageVoteApply is a helper method to define mock.On call
//   - ctx context.Context
//   - messageVote *domain.ForumMessageVote
func (_e *MockForumRepository_Expecter) ForumMessageVoteApply(ctx interface{}, messageVote interface{}) *MockForumRepository_ForumMessageVoteApply_Call {
	return &MockForumRepository_ForumMessageVoteApply_Call{Call: _e.mock.On("ForumMessageVoteApply", ctx, messageVote)}
}

func (_c *MockForumRepository_ForumMessageVoteApply_Call) Run(run func(ctx context.Context, messageVote *domain.ForumMessageVote)) *MockForumRepository_ForumMessageVoteApply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.ForumMessageVote))
	})
	return _c
}

func (_c *MockForumRepository_ForumMessageVoteApply_Call) Return(_a0 error) *MockForumRepository_ForumMessageVoteApply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumMessageVoteApply_Call) RunAndReturn(run func(context.Context, *domain.ForumMessageVote) error) *MockForumRepository_ForumMessageVoteApply_Call {
	_c.Call.Return(run)
	return _c
}

// ForumMessageVoteByID provides a mock function with given fields: ctx, messageVoteID, messageVote
func (_m *MockForumRepository) ForumMessageVoteByID(ctx context.Context, messageVoteID int64, messageVote *domain.ForumMessageVote) error {
	ret := _m.Called(ctx, messageVoteID, messageVote)

	if len(ret) == 0 {
		panic("no return value specified for ForumMessageVoteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.ForumMessageVote) error); ok {
		r0 = rf(ctx, messageVoteID, messageVote)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumMessageVoteByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumMessageVoteByID'
type MockForumRepository_ForumMessageVoteByID_Call struct {
	*mock.Call
}

// ForumMessageVoteByID is a helper method to define mock.On call
//   - ctx context.Context
//   - messageVoteID int64
//   - messageVote *domain.ForumMessageVote
func (_e *MockForumRepository_Expecter) ForumMessageVoteByID(ctx interface{}, messageVoteID interface{}, messageVote interface{}) *MockForumRepository_ForumMessageVoteByID_Call {
	return &MockForumRepository_ForumMessageVoteByID_Call{Call: _e.mock.On("ForumMessageVoteByID", ctx, messageVoteID, messageVote)}
}

func (_c *MockForumRepository_ForumMessageVoteByID_Call) Run(run func(ctx context.Context, messageVoteID int64, messageVote *domain.ForumMessageVote)) *MockForumRepository_ForumMessageVoteByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*domain.ForumMessageVote))
	})
	return _c
}

func (_c *MockForumRepository_ForumMessageVoteByID_Call) Return(_a0 error) *MockForumRepository_ForumMessageVoteByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumMessageVoteByID_Call) RunAndReturn(run func(context.Context, int64, *domain.ForumMessageVote) error) *MockForumRepository_ForumMessageVoteByID_Call {
	_c.Call.Return(run)
	return _c
}

// ForumMessages provides a mock function with given fields: ctx, filters
func (_m *MockForumRepository) ForumMessages(ctx context.Context, filters domain.ThreadMessagesQuery) ([]domain.ForumMessage, error) {
	ret := _m.Called(ctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for ForumMessages")
	}

	var r0 []domain.ForumMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ThreadMessagesQuery) ([]domain.ForumMessage, error)); ok {
		return rf(ctx, filters)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ThreadMessagesQuery) []domain.ForumMessage); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ForumMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ThreadMessagesQuery) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockForumRepository_ForumMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumMessages'
type MockForumRepository_ForumMessages_Call struct {
	*mock.Call
}

// ForumMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - filters domain.ThreadMessagesQuery
func (_e *MockForumRepository_Expecter) ForumMessages(ctx interface{}, filters interface{}) *MockForumRepository_ForumMessages_Call {
	return &MockForumRepository_ForumMessages_Call{Call: _e.mock.On("ForumMessages", ctx, filters)}
}

func (_c *MockForumRepository_ForumMessages_Call) Run(run func(ctx context.Context, filters domain.ThreadMessagesQuery)) *MockForumRepository_ForumMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ThreadMessagesQuery))
	})
	return _c
}

func (_c *MockForumRepository_ForumMessages_Call) Return(_a0 []domain.ForumMessage, _a1 error) *MockForumRepository_ForumMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockForumRepository_ForumMessages_Call) RunAndReturn(run func(context.Context, domain.ThreadMessagesQuery) ([]domain.ForumMessage, error)) *MockForumRepository_ForumMessages_Call {
	_c.Call.Return(run)
	return _c
}

// ForumRecentActivity provides a mock function with given fields: ctx, limit, permissionLevel
func (_m *MockForumRepository) ForumRecentActivity(ctx context.Context, limit uint64, permissionLevel domain.Privilege) ([]domain.ForumMessage, error) {
	ret := _m.Called(ctx, limit, permissionLevel)

	if len(ret) == 0 {
		panic("no return value specified for ForumRecentActivity")
	}

	var r0 []domain.ForumMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, domain.Privilege) ([]domain.ForumMessage, error)); ok {
		return rf(ctx, limit, permissionLevel)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, domain.Privilege) []domain.ForumMessage); ok {
		r0 = rf(ctx, limit, permissionLevel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ForumMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, domain.Privilege) error); ok {
		r1 = rf(ctx, limit, permissionLevel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockForumRepository_ForumRecentActivity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumRecentActivity'
type MockForumRepository_ForumRecentActivity_Call struct {
	*mock.Call
}

// ForumRecentActivity is a helper method to define mock.On call
//   - ctx context.Context
//   - limit uint64
//   - permissionLevel domain.Privilege
func (_e *MockForumRepository_Expecter) ForumRecentActivity(ctx interface{}, limit interface{}, permissionLevel interface{}) *MockForumRepository_ForumRecentActivity_Call {
	return &MockForumRepository_ForumRecentActivity_Call{Call: _e.mock.On("ForumRecentActivity", ctx, limit, permissionLevel)}
}

func (_c *MockForumRepository_ForumRecentActivity_Call) Run(run func(ctx context.Context, limit uint64, permissionLevel domain.Privilege)) *MockForumRepository_ForumRecentActivity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(domain.Privilege))
	})
	return _c
}

func (_c *MockForumRepository_ForumRecentActivity_Call) Return(_a0 []domain.ForumMessage, _a1 error) *MockForumRepository_ForumRecentActivity_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockForumRepository_ForumRecentActivity_Call) RunAndReturn(run func(context.Context, uint64, domain.Privilege) ([]domain.ForumMessage, error)) *MockForumRepository_ForumRecentActivity_Call {
	_c.Call.Return(run)
	return _c
}

// ForumSave provides a mock function with given fields: ctx, forum
func (_m *MockForumRepository) ForumSave(ctx context.Context, forum *domain.Forum) error {
	ret := _m.Called(ctx, forum)

	if len(ret) == 0 {
		panic("no return value specified for ForumSave")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Forum) error); ok {
		r0 = rf(ctx, forum)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumSave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumSave'
type MockForumRepository_ForumSave_Call struct {
	*mock.Call
}

// ForumSave is a helper method to define mock.On call
//   - ctx context.Context
//   - forum *domain.Forum
func (_e *MockForumRepository_Expecter) ForumSave(ctx interface{}, forum interface{}) *MockForumRepository_ForumSave_Call {
	return &MockForumRepository_ForumSave_Call{Call: _e.mock.On("ForumSave", ctx, forum)}
}

func (_c *MockForumRepository_ForumSave_Call) Run(run func(ctx context.Context, forum *domain.Forum)) *MockForumRepository_ForumSave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Forum))
	})
	return _c
}

func (_c *MockForumRepository_ForumSave_Call) Return(_a0 error) *MockForumRepository_ForumSave_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumSave_Call) RunAndReturn(run func(context.Context, *domain.Forum) error) *MockForumRepository_ForumSave_Call {
	_c.Call.Return(run)
	return _c
}

// ForumThread provides a mock function with given fields: ctx, forumThreadID, thread
func (_m *MockForumRepository) ForumThread(ctx context.Context, forumThreadID int64, thread *domain.ForumThread) error {
	ret := _m.Called(ctx, forumThreadID, thread)

	if len(ret) == 0 {
		panic("no return value specified for ForumThread")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.ForumThread) error); ok {
		r0 = rf(ctx, forumThreadID, thread)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumThread_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumThread'
type MockForumRepository_ForumThread_Call struct {
	*mock.Call
}

// ForumThread is a helper method to define mock.On call
//   - ctx context.Context
//   - forumThreadID int64
//   - thread *domain.ForumThread
func (_e *MockForumRepository_Expecter) ForumThread(ctx interface{}, forumThreadID interface{}, thread interface{}) *MockForumRepository_ForumThread_Call {
	return &MockForumRepository_ForumThread_Call{Call: _e.mock.On("ForumThread", ctx, forumThreadID, thread)}
}

func (_c *MockForumRepository_ForumThread_Call) Run(run func(ctx context.Context, forumThreadID int64, thread *domain.ForumThread)) *MockForumRepository_ForumThread_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*domain.ForumThread))
	})
	return _c
}

func (_c *MockForumRepository_ForumThread_Call) Return(_a0 error) *MockForumRepository_ForumThread_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumThread_Call) RunAndReturn(run func(context.Context, int64, *domain.ForumThread) error) *MockForumRepository_ForumThread_Call {
	_c.Call.Return(run)
	return _c
}

// ForumThreadDelete provides a mock function with given fields: ctx, forumThreadID
func (_m *MockForumRepository) ForumThreadDelete(ctx context.Context, forumThreadID int64) error {
	ret := _m.Called(ctx, forumThreadID)

	if len(ret) == 0 {
		panic("no return value specified for ForumThreadDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, forumThreadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumThreadDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumThreadDelete'
type MockForumRepository_ForumThreadDelete_Call struct {
	*mock.Call
}

// ForumThreadDelete is a helper method to define mock.On call
//   - ctx context.Context
//   - forumThreadID int64
func (_e *MockForumRepository_Expecter) ForumThreadDelete(ctx interface{}, forumThreadID interface{}) *MockForumRepository_ForumThreadDelete_Call {
	return &MockForumRepository_ForumThreadDelete_Call{Call: _e.mock.On("ForumThreadDelete", ctx, forumThreadID)}
}

func (_c *MockForumRepository_ForumThreadDelete_Call) Run(run func(ctx context.Context, forumThreadID int64)) *MockForumRepository_ForumThreadDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockForumRepository_ForumThreadDelete_Call) Return(_a0 error) *MockForumRepository_ForumThreadDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumThreadDelete_Call) RunAndReturn(run func(context.Context, int64) error) *MockForumRepository_ForumThreadDelete_Call {
	_c.Call.Return(run)
	return _c
}

// ForumThreadIncrView provides a mock function with given fields: ctx, forumThreadID
func (_m *MockForumRepository) ForumThreadIncrView(ctx context.Context, forumThreadID int64) error {
	ret := _m.Called(ctx, forumThreadID)

	if len(ret) == 0 {
		panic("no return value specified for ForumThreadIncrView")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, forumThreadID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumThreadIncrView_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumThreadIncrView'
type MockForumRepository_ForumThreadIncrView_Call struct {
	*mock.Call
}

// ForumThreadIncrView is a helper method to define mock.On call
//   - ctx context.Context
//   - forumThreadID int64
func (_e *MockForumRepository_Expecter) ForumThreadIncrView(ctx interface{}, forumThreadID interface{}) *MockForumRepository_ForumThreadIncrView_Call {
	return &MockForumRepository_ForumThreadIncrView_Call{Call: _e.mock.On("ForumThreadIncrView", ctx, forumThreadID)}
}

func (_c *MockForumRepository_ForumThreadIncrView_Call) Run(run func(ctx context.Context, forumThreadID int64)) *MockForumRepository_ForumThreadIncrView_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockForumRepository_ForumThreadIncrView_Call) Return(_a0 error) *MockForumRepository_ForumThreadIncrView_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumThreadIncrView_Call) RunAndReturn(run func(context.Context, int64) error) *MockForumRepository_ForumThreadIncrView_Call {
	_c.Call.Return(run)
	return _c
}

// ForumThreadSave provides a mock function with given fields: ctx, thread
func (_m *MockForumRepository) ForumThreadSave(ctx context.Context, thread *domain.ForumThread) error {
	ret := _m.Called(ctx, thread)

	if len(ret) == 0 {
		panic("no return value specified for ForumThreadSave")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ForumThread) error); ok {
		r0 = rf(ctx, thread)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockForumRepository_ForumThreadSave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumThreadSave'
type MockForumRepository_ForumThreadSave_Call struct {
	*mock.Call
}

// ForumThreadSave is a helper method to define mock.On call
//   - ctx context.Context
//   - thread *domain.ForumThread
func (_e *MockForumRepository_Expecter) ForumThreadSave(ctx interface{}, thread interface{}) *MockForumRepository_ForumThreadSave_Call {
	return &MockForumRepository_ForumThreadSave_Call{Call: _e.mock.On("ForumThreadSave", ctx, thread)}
}

func (_c *MockForumRepository_ForumThreadSave_Call) Run(run func(ctx context.Context, thread *domain.ForumThread)) *MockForumRepository_ForumThreadSave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.ForumThread))
	})
	return _c
}

func (_c *MockForumRepository_ForumThreadSave_Call) Return(_a0 error) *MockForumRepository_ForumThreadSave_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockForumRepository_ForumThreadSave_Call) RunAndReturn(run func(context.Context, *domain.ForumThread) error) *MockForumRepository_ForumThreadSave_Call {
	_c.Call.Return(run)
	return _c
}

// ForumThreads provides a mock function with given fields: ctx, filter
func (_m *MockForumRepository) ForumThreads(ctx context.Context, filter domain.ThreadQueryFilter) ([]domain.ThreadWithSource, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for ForumThreads")
	}

	var r0 []domain.ThreadWithSource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ThreadQueryFilter) ([]domain.ThreadWithSource, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ThreadQueryFilter) []domain.ThreadWithSource); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ThreadWithSource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ThreadQueryFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockForumRepository_ForumThreads_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForumThreads'
type MockForumRepository_ForumThreads_Call struct {
	*mock.Call
}

// ForumThreads is a helper method to define mock.On call
//   - ctx context.Context
//   - filter domain.ThreadQueryFilter
func (_e *MockForumRepository_Expecter) ForumThreads(ctx interface{}, filter interface{}) *MockForumRepository_ForumThreads_Call {
	return &MockForumRepository_ForumThreads_Call{Call: _e.mock.On("ForumThreads", ctx, filter)}
}

func (_c *MockForumRepository_ForumThreads_Call) Run(run func(ctx context.Context, filter domain.ThreadQueryFilter)) *MockForumRepository_ForumThreads_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ThreadQueryFilter))
	})
	return _c
}

func (_c *MockForumRepository_ForumThreads_Call) Return(_a0 []domain.ThreadWithSource, _a1 error) *MockForumRepository_ForumThreads_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockForumRepository_ForumThreads_Call) RunAndReturn(run func(context.Context, domain.ThreadQueryFilter) ([]domain.ThreadWithSource, error)) *MockForumRepository_ForumThreads_Call {
	_c.Call.Return(run)
	return _c
}

// Forums provides a mock function with given fields: ctx
func (_m *MockForumRepository) Forums(ctx context.Context) ([]domain.Forum, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Forums")
	}

	var r0 []domain.Forum
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Forum, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Forum); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Forum)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockForumRepository_Forums_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Forums'
type MockForumRepository_Forums_Call struct {
	*mock.Call
}

// Forums is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockForumRepository_Expecter) Forums(ctx interface{}) *MockForumRepository_Forums_Call {
	return &MockForumRepository_Forums_Call{Call: _e.mock.On("Forums", ctx)}
}

func (_c *MockForumRepository_Forums_Call) Run(run func(ctx context.Context)) *MockForumRepository_Forums_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockForumRepository_Forums_Call) Return(_a0 []domain.Forum, _a1 error) *MockForumRepository_Forums_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockForumRepository_Forums_Call) RunAndReturn(run func(context.Context) ([]domain.Forum, error)) *MockForumRepository_Forums_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockForumRepository creates a new instance of MockForumRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockForumRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockForumRepository {
	mock := &MockForumRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
