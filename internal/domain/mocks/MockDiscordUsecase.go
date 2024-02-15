// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	discordgo "github.com/bwmarrin/discordgo"
	domain "github.com/leighmacdonald/gbans/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockDiscordUsecase is an autogenerated mock type for the DiscordUsecase type
type MockDiscordUsecase struct {
	mock.Mock
}

type MockDiscordUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDiscordUsecase) EXPECT() *MockDiscordUsecase_Expecter {
	return &MockDiscordUsecase_Expecter{mock: &_m.Mock}
}

// FilterAdd provides a mock function with given fields: ctx, user, pattern, isRegex
func (_m *MockDiscordUsecase) FilterAdd(ctx context.Context, user domain.PersonInfo, pattern string, isRegex bool) (*discordgo.MessageEmbed, error) {
	ret := _m.Called(ctx, user, pattern, isRegex)

	if len(ret) == 0 {
		panic("no return value specified for FilterAdd")
	}

	var r0 *discordgo.MessageEmbed
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, string, bool) (*discordgo.MessageEmbed, error)); ok {
		return rf(ctx, user, pattern, isRegex)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PersonInfo, string, bool) *discordgo.MessageEmbed); ok {
		r0 = rf(ctx, user, pattern, isRegex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.MessageEmbed)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PersonInfo, string, bool) error); ok {
		r1 = rf(ctx, user, pattern, isRegex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiscordUsecase_FilterAdd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterAdd'
type MockDiscordUsecase_FilterAdd_Call struct {
	*mock.Call
}

// FilterAdd is a helper method to define mock.On call
//   - ctx context.Context
//   - user domain.PersonInfo
//   - pattern string
//   - isRegex bool
func (_e *MockDiscordUsecase_Expecter) FilterAdd(ctx interface{}, user interface{}, pattern interface{}, isRegex interface{}) *MockDiscordUsecase_FilterAdd_Call {
	return &MockDiscordUsecase_FilterAdd_Call{Call: _e.mock.On("FilterAdd", ctx, user, pattern, isRegex)}
}

func (_c *MockDiscordUsecase_FilterAdd_Call) Run(run func(ctx context.Context, user domain.PersonInfo, pattern string, isRegex bool)) *MockDiscordUsecase_FilterAdd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PersonInfo), args[2].(string), args[3].(bool))
	})
	return _c
}

func (_c *MockDiscordUsecase_FilterAdd_Call) Return(_a0 *discordgo.MessageEmbed, _a1 error) *MockDiscordUsecase_FilterAdd_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiscordUsecase_FilterAdd_Call) RunAndReturn(run func(context.Context, domain.PersonInfo, string, bool) (*discordgo.MessageEmbed, error)) *MockDiscordUsecase_FilterAdd_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterHandler provides a mock function with given fields: cmd, handler
func (_m *MockDiscordUsecase) RegisterHandler(cmd domain.Cmd, handler domain.SlashCommandHandler) error {
	ret := _m.Called(cmd, handler)

	if len(ret) == 0 {
		panic("no return value specified for RegisterHandler")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Cmd, domain.SlashCommandHandler) error); ok {
		r0 = rf(cmd, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDiscordUsecase_RegisterHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterHandler'
type MockDiscordUsecase_RegisterHandler_Call struct {
	*mock.Call
}

// RegisterHandler is a helper method to define mock.On call
//   - cmd domain.Cmd
//   - handler domain.SlashCommandHandler
func (_e *MockDiscordUsecase_Expecter) RegisterHandler(cmd interface{}, handler interface{}) *MockDiscordUsecase_RegisterHandler_Call {
	return &MockDiscordUsecase_RegisterHandler_Call{Call: _e.mock.On("RegisterHandler", cmd, handler)}
}

func (_c *MockDiscordUsecase_RegisterHandler_Call) Run(run func(cmd domain.Cmd, handler domain.SlashCommandHandler)) *MockDiscordUsecase_RegisterHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.Cmd), args[1].(domain.SlashCommandHandler))
	})
	return _c
}

func (_c *MockDiscordUsecase_RegisterHandler_Call) Return(_a0 error) *MockDiscordUsecase_RegisterHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscordUsecase_RegisterHandler_Call) RunAndReturn(run func(domain.Cmd, domain.SlashCommandHandler) error) *MockDiscordUsecase_RegisterHandler_Call {
	_c.Call.Return(run)
	return _c
}

// SendPayload provides a mock function with given fields: channelID, embed
func (_m *MockDiscordUsecase) SendPayload(channelID domain.DiscordChannel, embed *discordgo.MessageEmbed) {
	_m.Called(channelID, embed)
}

// MockDiscordUsecase_SendPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendPayload'
type MockDiscordUsecase_SendPayload_Call struct {
	*mock.Call
}

// SendPayload is a helper method to define mock.On call
//   - channelID domain.DiscordChannel
//   - embed *discordgo.MessageEmbed
func (_e *MockDiscordUsecase_Expecter) SendPayload(channelID interface{}, embed interface{}) *MockDiscordUsecase_SendPayload_Call {
	return &MockDiscordUsecase_SendPayload_Call{Call: _e.mock.On("SendPayload", channelID, embed)}
}

func (_c *MockDiscordUsecase_SendPayload_Call) Run(run func(channelID domain.DiscordChannel, embed *discordgo.MessageEmbed)) *MockDiscordUsecase_SendPayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.DiscordChannel), args[1].(*discordgo.MessageEmbed))
	})
	return _c
}

func (_c *MockDiscordUsecase_SendPayload_Call) Return() *MockDiscordUsecase_SendPayload_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDiscordUsecase_SendPayload_Call) RunAndReturn(run func(domain.DiscordChannel, *discordgo.MessageEmbed)) *MockDiscordUsecase_SendPayload_Call {
	_c.Call.Return(run)
	return _c
}

// Shutdown provides a mock function with given fields: guildID
func (_m *MockDiscordUsecase) Shutdown(guildID string) {
	_m.Called(guildID)
}

// MockDiscordUsecase_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type MockDiscordUsecase_Shutdown_Call struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
//   - guildID string
func (_e *MockDiscordUsecase_Expecter) Shutdown(guildID interface{}) *MockDiscordUsecase_Shutdown_Call {
	return &MockDiscordUsecase_Shutdown_Call{Call: _e.mock.On("Shutdown", guildID)}
}

func (_c *MockDiscordUsecase_Shutdown_Call) Run(run func(guildID string)) *MockDiscordUsecase_Shutdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDiscordUsecase_Shutdown_Call) Return() *MockDiscordUsecase_Shutdown_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDiscordUsecase_Shutdown_Call) RunAndReturn(run func(string)) *MockDiscordUsecase_Shutdown_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockDiscordUsecase) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDiscordUsecase_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockDiscordUsecase_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockDiscordUsecase_Expecter) Start() *MockDiscordUsecase_Start_Call {
	return &MockDiscordUsecase_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockDiscordUsecase_Start_Call) Run(run func()) *MockDiscordUsecase_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDiscordUsecase_Start_Call) Return(_a0 error) *MockDiscordUsecase_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscordUsecase_Start_Call) RunAndReturn(run func() error) *MockDiscordUsecase_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDiscordUsecase creates a new instance of MockDiscordUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDiscordUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDiscordUsecase {
	mock := &MockDiscordUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}