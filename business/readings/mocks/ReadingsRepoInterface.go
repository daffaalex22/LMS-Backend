// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	modules "backend/business/modules"
	context "context"

	mock "github.com/stretchr/testify/mock"

	readings "backend/business/readings"
)

// ReadingsRepoInterface is an autogenerated mock type for the ReadingsRepoInterface type
type ReadingsRepoInterface struct {
	mock.Mock
}

// CheckModule provides a mock function with given fields: ctx, id
func (_m *ReadingsRepoInterface) CheckModule(ctx context.Context, id uint) (modules.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 modules.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) modules.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(modules.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsAdd provides a mock function with given fields: ctx, domain
func (_m *ReadingsRepoInterface) ReadingsAdd(ctx context.Context, domain readings.Domain) (readings.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 readings.Domain
	if rf, ok := ret.Get(0).(func(context.Context, readings.Domain) readings.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(readings.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, readings.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsDelete provides a mock function with given fields: ctx, id
func (_m *ReadingsRepoInterface) ReadingsDelete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadingsGetById provides a mock function with given fields: ctx, id
func (_m *ReadingsRepoInterface) ReadingsGetById(ctx context.Context, id uint) (readings.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 readings.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) readings.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(readings.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsGetByModuleId provides a mock function with given fields: ctx, moduleId
func (_m *ReadingsRepoInterface) ReadingsGetByModuleId(ctx context.Context, moduleId uint) ([]readings.Domain, error) {
	ret := _m.Called(ctx, moduleId)

	var r0 []readings.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) []readings.Domain); ok {
		r0 = rf(ctx, moduleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]readings.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, moduleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsUpdate provides a mock function with given fields: ctx, domain, id
func (_m *ReadingsRepoInterface) ReadingsUpdate(ctx context.Context, domain readings.Domain, id uint) (readings.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 readings.Domain
	if rf, ok := ret.Get(0).(func(context.Context, readings.Domain, uint) readings.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(readings.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, readings.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}