// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/meroedu/meroedu/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// TagUseCase is an autogenerated mock type for the TagUseCase type
type TagUseCase struct {
	mock.Mock
}

// CreateTag provides a mock function with given fields: ctx, Tag
func (_m *TagUseCase) CreateTag(ctx context.Context, Tag *domain.Tag) error {
	ret := _m.Called(ctx, Tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Tag) error); ok {
		r0 = rf(ctx, Tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTag provides a mock function with given fields: ctx, id
func (_m *TagUseCase) DeleteTag(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, start, limit
func (_m *TagUseCase) GetAll(ctx context.Context, start int, limit int) ([]domain.Tag, error) {
	ret := _m.Called(ctx, start, limit)

	var r0 []domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []domain.Tag); ok {
		r0 = rf(ctx, start, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, start, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TagUseCase) GetByID(ctx context.Context, id int64) (*domain.Tag, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Tag); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *TagUseCase) GetByName(ctx context.Context, name string) (*domain.Tag, error) {
	ret := _m.Called(ctx, name)

	var r0 *domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Tag); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTag provides a mock function with given fields: ctx, Tag, id
func (_m *TagUseCase) UpdateTag(ctx context.Context, Tag *domain.Tag, id int64) error {
	ret := _m.Called(ctx, Tag, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Tag, int64) error); ok {
		r0 = rf(ctx, Tag, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
