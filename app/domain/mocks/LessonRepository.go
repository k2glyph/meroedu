// Code generated by mockery v2.2.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/meroedu/meroedu/app/domain"
	mock "github.com/stretchr/testify/mock"
)

// LessonRepository is an autogenerated mock type for the LessonRepository type
type LessonRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *LessonRepository) GetByID(ctx context.Context, id int64) (domain.Lesson, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Lesson
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Lesson); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Lesson)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
