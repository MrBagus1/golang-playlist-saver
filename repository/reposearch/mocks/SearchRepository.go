// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	record "playlist-saver/model/record"

	mock "github.com/stretchr/testify/mock"
)

// SearchRepository is an autogenerated mock type for the SearchRepository type
type SearchRepository struct {
	mock.Mock
}

// GetByUrlId provides a mock function with given fields: ctx, Id
func (_m *SearchRepository) GetByUrlId(ctx context.Context, Id string) record.YoutubeData {
	ret := _m.Called(ctx, Id)

	var r0 record.YoutubeData
	if rf, ok := ret.Get(0).(func(context.Context, string) record.YoutubeData); ok {
		r0 = rf(ctx, Id)
	} else {
		r0 = ret.Get(0).(record.YoutubeData)
	}

	return r0
}

// SearchYtById provides a mock function with given fields: ctx, Id
func (_m *SearchRepository) SearchYtById(ctx context.Context, Id string) record.YoutubeData {
	ret := _m.Called(ctx, Id)

	var r0 record.YoutubeData
	if rf, ok := ret.Get(0).(func(context.Context, string) record.YoutubeData); ok {
		r0 = rf(ctx, Id)
	} else {
		r0 = ret.Get(0).(record.YoutubeData)
	}

	return r0
}

// SearchYtByParam provides a mock function with given fields: ctx, Search
func (_m *SearchRepository) SearchYtByParam(ctx context.Context, Search string) ([]record.YoutubeData, error) {
	ret := _m.Called(ctx, Search)

	var r0 []record.YoutubeData
	if rf, ok := ret.Get(0).(func(context.Context, string) []record.YoutubeData); ok {
		r0 = rf(ctx, Search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]record.YoutubeData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, Search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
