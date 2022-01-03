// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	categories "backend/business/categories"
	context "context"

	course "backend/business/course"

	mock "github.com/stretchr/testify/mock"

	teacher "backend/business/teacher"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CheckCategories provides a mock function with given fields: ctx, id
func (_m *Repository) CheckCategories(ctx context.Context, id uint) (categories.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 categories.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) categories.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTeacher provides a mock function with given fields: ctx, id
func (_m *Repository) CheckTeacher(ctx context.Context, id uint) (teacher.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 teacher.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) teacher.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(teacher.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, domain
func (_m *Repository) Create(ctx context.Context, domain course.Domain) (course.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 course.Domain
	if rf, ok := ret.Get(0).(func(context.Context, course.Domain) course.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(course.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, course.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]course.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []course.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []course.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]course.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCourseById provides a mock function with given fields: ctx, id
func (_m *Repository) GetCourseById(ctx context.Context, id uint) (course.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 course.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) course.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(course.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCourseByStudentId provides a mock function with given fields: ctx, courseIds
func (_m *Repository) GetCourseByStudentId(ctx context.Context, courseIds []uint) ([]course.Domain, error) {
	ret := _m.Called(ctx, courseIds)

	var r0 []course.Domain
	if rf, ok := ret.Get(0).(func(context.Context, []uint) []course.Domain); ok {
		r0 = rf(ctx, courseIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]course.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []uint) error); ok {
		r1 = rf(ctx, courseIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCourseByTeacherId provides a mock function with given fields: ctx, teacherId
func (_m *Repository) GetCourseByTeacherId(ctx context.Context, teacherId uint) ([]course.Domain, error) {
	ret := _m.Called(ctx, teacherId)

	var r0 []course.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) []course.Domain); ok {
		r0 = rf(ctx, teacherId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]course.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, teacherId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnrollmentsByStudentId provides a mock function with given fields: ctx, studentId
func (_m *Repository) GetEnrollmentsByStudentId(ctx context.Context, studentId uint) ([]course.CourseEnrollmentDomain, error) {
	ret := _m.Called(ctx, studentId)

	var r0 []course.CourseEnrollmentDomain
	if rf, ok := ret.Get(0).(func(context.Context, uint) []course.CourseEnrollmentDomain); ok {
		r0 = rf(ctx, studentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]course.CourseEnrollmentDomain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, studentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, domain
func (_m *Repository) Update(ctx context.Context, domain course.Domain) (course.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 course.Domain
	if rf, ok := ret.Get(0).(func(context.Context, course.Domain) course.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(course.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, course.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
