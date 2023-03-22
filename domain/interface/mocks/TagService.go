// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entitie "microservice/domain/entitie"

	mock "github.com/stretchr/testify/mock"
)

// TagService is an autogenerated mock type for the TagService type
type TagService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TagService) Delete(ctx context.Context, id uint) error {
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
func (_m *TagService) GetAll(ctx context.Context) ([]entitie.Tag, error) {
	ret := _m.Called(ctx)

	var r0 []entitie.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entitie.Tag, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entitie.Tag); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entitie.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TagService) GetByID(ctx context.Context, id uint) (entitie.Tag, error) {
	ret := _m.Called(ctx, id)

	var r0 entitie.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (entitie.Tag, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) entitie.Tag); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entitie.Tag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, title
func (_m *TagService) GetByName(ctx context.Context, title string) (entitie.Tag, error) {
	ret := _m.Called(ctx, title)

	var r0 entitie.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entitie.Tag, error)); ok {
		return rf(ctx, title)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entitie.Tag); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(entitie.Tag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *TagService) Store(_a0 context.Context, _a1 *entitie.Tag) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entitie.Tag) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, ar
func (_m *TagService) Update(ctx context.Context, ar *entitie.Tag) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entitie.Tag) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTagService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTagService creates a new instance of TagService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTagService(t mockConstructorTestingTNewTagService) *TagService {
	mock := &TagService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
