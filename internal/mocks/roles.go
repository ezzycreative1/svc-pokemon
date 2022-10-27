package mocks

import (
	"context"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type Roles struct {
	mock.Mock
}

type Roles_Expecter struct {
	mock *mock.Mock
}

func (_m *Roles) EXPECT() *Roles_Expecter {
	return &Roles_Expecter{mock: &_m.Mock}
}

// CreateRoles provides a mock function with given fields: ctx, role
func (_m *Roles) CreateRoles(ctx context.Context, role domain.Roles) (uint64, error) {
	ret := _m.Called(ctx, role)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, domain.Roles) uint64); ok {
		r0 = rf(ctx, role)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Roles) error); ok {
		r1 = rf(ctx, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Roles_CreateRoles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRoles'
type Roles_CreateRoles_Call struct {
	*mock.Call
}

// CreateRoles is a helper method to define mock.On call
//  - ctx context.Context
//  - roles domain.Roles
func (_e *Roles_CreateRoles_Call) CreateRoles(ctx interface{}, role interface{}) *Roles_CreateRoles_Call {
	return &Roles_CreateRoles_Call{Call: _e.mock.On("CreateRoles", ctx, role)}
}

func (_c *Roles_CreateRoles_Call) Run(run func(ctx context.Context, input domain.Roles)) *Roles_CreateRoles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Roles))
	})
	return _c
}

func (_c *Roles_CreateRoles_Call) Return(_a0 uint64, _a1 error) *Roles_CreateRoles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
