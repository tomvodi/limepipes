// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// ApiHandler is an autogenerated mock type for the ApiHandler type
type ApiHandler struct {
	mock.Mock
}

type ApiHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *ApiHandler) EXPECT() *ApiHandler_Expecter {
	return &ApiHandler_Expecter{mock: &_m.Mock}
}

// AssignTunesToSet provides a mock function with given fields: c
func (_m *ApiHandler) AssignTunesToSet(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_AssignTunesToSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AssignTunesToSet'
type ApiHandler_AssignTunesToSet_Call struct {
	*mock.Call
}

// AssignTunesToSet is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) AssignTunesToSet(c interface{}) *ApiHandler_AssignTunesToSet_Call {
	return &ApiHandler_AssignTunesToSet_Call{Call: _e.mock.On("AssignTunesToSet", c)}
}

func (_c *ApiHandler_AssignTunesToSet_Call) Run(run func(c *gin.Context)) *ApiHandler_AssignTunesToSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_AssignTunesToSet_Call) Return() *ApiHandler_AssignTunesToSet_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_AssignTunesToSet_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_AssignTunesToSet_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSet provides a mock function with given fields: c
func (_m *ApiHandler) CreateSet(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_CreateSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSet'
type ApiHandler_CreateSet_Call struct {
	*mock.Call
}

// CreateSet is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) CreateSet(c interface{}) *ApiHandler_CreateSet_Call {
	return &ApiHandler_CreateSet_Call{Call: _e.mock.On("CreateSet", c)}
}

func (_c *ApiHandler_CreateSet_Call) Run(run func(c *gin.Context)) *ApiHandler_CreateSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_CreateSet_Call) Return() *ApiHandler_CreateSet_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_CreateSet_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_CreateSet_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTune provides a mock function with given fields: c
func (_m *ApiHandler) CreateTune(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_CreateTune_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTune'
type ApiHandler_CreateTune_Call struct {
	*mock.Call
}

// CreateTune is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) CreateTune(c interface{}) *ApiHandler_CreateTune_Call {
	return &ApiHandler_CreateTune_Call{Call: _e.mock.On("CreateTune", c)}
}

func (_c *ApiHandler_CreateTune_Call) Run(run func(c *gin.Context)) *ApiHandler_CreateTune_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_CreateTune_Call) Return() *ApiHandler_CreateTune_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_CreateTune_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_CreateTune_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSet provides a mock function with given fields: c
func (_m *ApiHandler) DeleteSet(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_DeleteSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSet'
type ApiHandler_DeleteSet_Call struct {
	*mock.Call
}

// DeleteSet is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) DeleteSet(c interface{}) *ApiHandler_DeleteSet_Call {
	return &ApiHandler_DeleteSet_Call{Call: _e.mock.On("DeleteSet", c)}
}

func (_c *ApiHandler_DeleteSet_Call) Run(run func(c *gin.Context)) *ApiHandler_DeleteSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_DeleteSet_Call) Return() *ApiHandler_DeleteSet_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_DeleteSet_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_DeleteSet_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTune provides a mock function with given fields: c
func (_m *ApiHandler) DeleteTune(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_DeleteTune_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTune'
type ApiHandler_DeleteTune_Call struct {
	*mock.Call
}

// DeleteTune is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) DeleteTune(c interface{}) *ApiHandler_DeleteTune_Call {
	return &ApiHandler_DeleteTune_Call{Call: _e.mock.On("DeleteTune", c)}
}

func (_c *ApiHandler_DeleteTune_Call) Run(run func(c *gin.Context)) *ApiHandler_DeleteTune_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_DeleteTune_Call) Return() *ApiHandler_DeleteTune_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_DeleteTune_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_DeleteTune_Call {
	_c.Call.Return(run)
	return _c
}

// GetSet provides a mock function with given fields: c
func (_m *ApiHandler) GetSet(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_GetSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSet'
type ApiHandler_GetSet_Call struct {
	*mock.Call
}

// GetSet is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) GetSet(c interface{}) *ApiHandler_GetSet_Call {
	return &ApiHandler_GetSet_Call{Call: _e.mock.On("GetSet", c)}
}

func (_c *ApiHandler_GetSet_Call) Run(run func(c *gin.Context)) *ApiHandler_GetSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_GetSet_Call) Return() *ApiHandler_GetSet_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_GetSet_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_GetSet_Call {
	_c.Call.Return(run)
	return _c
}

// GetTune provides a mock function with given fields: c
func (_m *ApiHandler) GetTune(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_GetTune_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTune'
type ApiHandler_GetTune_Call struct {
	*mock.Call
}

// GetTune is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) GetTune(c interface{}) *ApiHandler_GetTune_Call {
	return &ApiHandler_GetTune_Call{Call: _e.mock.On("GetTune", c)}
}

func (_c *ApiHandler_GetTune_Call) Run(run func(c *gin.Context)) *ApiHandler_GetTune_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_GetTune_Call) Return() *ApiHandler_GetTune_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_GetTune_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_GetTune_Call {
	_c.Call.Return(run)
	return _c
}

// Health provides a mock function with given fields: c
func (_m *ApiHandler) Health(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_Health_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Health'
type ApiHandler_Health_Call struct {
	*mock.Call
}

// Health is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) Health(c interface{}) *ApiHandler_Health_Call {
	return &ApiHandler_Health_Call{Call: _e.mock.On("Health", c)}
}

func (_c *ApiHandler_Health_Call) Run(run func(c *gin.Context)) *ApiHandler_Health_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_Health_Call) Return() *ApiHandler_Health_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_Health_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_Health_Call {
	_c.Call.Return(run)
	return _c
}

// Home provides a mock function with given fields: c
func (_m *ApiHandler) Home(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_Home_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Home'
type ApiHandler_Home_Call struct {
	*mock.Call
}

// Home is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) Home(c interface{}) *ApiHandler_Home_Call {
	return &ApiHandler_Home_Call{Call: _e.mock.On("Home", c)}
}

func (_c *ApiHandler_Home_Call) Run(run func(c *gin.Context)) *ApiHandler_Home_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_Home_Call) Return() *ApiHandler_Home_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_Home_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_Home_Call {
	_c.Call.Return(run)
	return _c
}

// ImportFile provides a mock function with given fields: c
func (_m *ApiHandler) ImportFile(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_ImportFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ImportFile'
type ApiHandler_ImportFile_Call struct {
	*mock.Call
}

// ImportFile is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) ImportFile(c interface{}) *ApiHandler_ImportFile_Call {
	return &ApiHandler_ImportFile_Call{Call: _e.mock.On("ImportFile", c)}
}

func (_c *ApiHandler_ImportFile_Call) Run(run func(c *gin.Context)) *ApiHandler_ImportFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_ImportFile_Call) Return() *ApiHandler_ImportFile_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_ImportFile_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_ImportFile_Call {
	_c.Call.Return(run)
	return _c
}

// ListSets provides a mock function with given fields: c
func (_m *ApiHandler) ListSets(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_ListSets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSets'
type ApiHandler_ListSets_Call struct {
	*mock.Call
}

// ListSets is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) ListSets(c interface{}) *ApiHandler_ListSets_Call {
	return &ApiHandler_ListSets_Call{Call: _e.mock.On("ListSets", c)}
}

func (_c *ApiHandler_ListSets_Call) Run(run func(c *gin.Context)) *ApiHandler_ListSets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_ListSets_Call) Return() *ApiHandler_ListSets_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_ListSets_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_ListSets_Call {
	_c.Call.Return(run)
	return _c
}

// ListTunes provides a mock function with given fields: c
func (_m *ApiHandler) ListTunes(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_ListTunes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTunes'
type ApiHandler_ListTunes_Call struct {
	*mock.Call
}

// ListTunes is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) ListTunes(c interface{}) *ApiHandler_ListTunes_Call {
	return &ApiHandler_ListTunes_Call{Call: _e.mock.On("ListTunes", c)}
}

func (_c *ApiHandler_ListTunes_Call) Run(run func(c *gin.Context)) *ApiHandler_ListTunes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_ListTunes_Call) Return() *ApiHandler_ListTunes_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_ListTunes_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_ListTunes_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSet provides a mock function with given fields: c
func (_m *ApiHandler) UpdateSet(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_UpdateSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSet'
type ApiHandler_UpdateSet_Call struct {
	*mock.Call
}

// UpdateSet is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) UpdateSet(c interface{}) *ApiHandler_UpdateSet_Call {
	return &ApiHandler_UpdateSet_Call{Call: _e.mock.On("UpdateSet", c)}
}

func (_c *ApiHandler_UpdateSet_Call) Run(run func(c *gin.Context)) *ApiHandler_UpdateSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_UpdateSet_Call) Return() *ApiHandler_UpdateSet_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_UpdateSet_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_UpdateSet_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTune provides a mock function with given fields: c
func (_m *ApiHandler) UpdateTune(c *gin.Context) {
	_m.Called(c)
}

// ApiHandler_UpdateTune_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTune'
type ApiHandler_UpdateTune_Call struct {
	*mock.Call
}

// UpdateTune is a helper method to define mock.On call
//   - c *gin.Context
func (_e *ApiHandler_Expecter) UpdateTune(c interface{}) *ApiHandler_UpdateTune_Call {
	return &ApiHandler_UpdateTune_Call{Call: _e.mock.On("UpdateTune", c)}
}

func (_c *ApiHandler_UpdateTune_Call) Run(run func(c *gin.Context)) *ApiHandler_UpdateTune_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ApiHandler_UpdateTune_Call) Return() *ApiHandler_UpdateTune_Call {
	_c.Call.Return()
	return _c
}

func (_c *ApiHandler_UpdateTune_Call) RunAndReturn(run func(*gin.Context)) *ApiHandler_UpdateTune_Call {
	_c.Call.Return(run)
	return _c
}

// NewApiHandler creates a new instance of ApiHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApiHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *ApiHandler {
	mock := &ApiHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
