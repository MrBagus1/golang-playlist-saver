// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	record "playlist-saver/model/record"

	mock "github.com/stretchr/testify/mock"
)

// PlaylistRepository is an autogenerated mock type for the PlaylistRepository type
type PlaylistRepository struct {
	mock.Mock
}

// CreatePlaylist provides a mock function with given fields: ctx, name, id
func (_m *PlaylistRepository) CreatePlaylist(ctx context.Context, name string, id int) (record.Playlist, error) {
	ret := _m.Called(ctx, name, id)

	var r0 record.Playlist
	if rf, ok := ret.Get(0).(func(context.Context, string, int) record.Playlist); ok {
		r0 = rf(ctx, name, id)
	} else {
		r0 = ret.Get(0).(record.Playlist)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, name, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePlaylist provides a mock function with given fields: ctx, id
func (_m *PlaylistRepository) DeletePlaylist(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditPlaylist provides a mock function with given fields: ctx, playlist, id
func (_m *PlaylistRepository) EditPlaylist(ctx context.Context, playlist record.Playlist, id int) (record.Playlist, error) {
	ret := _m.Called(ctx, playlist, id)

	var r0 record.Playlist
	if rf, ok := ret.Get(0).(func(context.Context, record.Playlist, int) record.Playlist); ok {
		r0 = rf(ctx, playlist, id)
	} else {
		r0 = ret.Get(0).(record.Playlist)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, record.Playlist, int) error); ok {
		r1 = rf(ctx, playlist, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPlaylist provides a mock function with given fields: ctx
func (_m *PlaylistRepository) GetAllPlaylist(ctx context.Context) ([]record.Playlist, error) {
	ret := _m.Called(ctx)

	var r0 []record.Playlist
	if rf, ok := ret.Get(0).(func(context.Context) []record.Playlist); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]record.Playlist)
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

// GetPlaylistByUserId provides a mock function with given fields: ctx, id
func (_m *PlaylistRepository) GetPlaylistByUserId(ctx context.Context, id int) ([]record.Playlist, error) {
	ret := _m.Called(ctx, id)

	var r0 []record.Playlist
	if rf, ok := ret.Get(0).(func(context.Context, int) []record.Playlist); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]record.Playlist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
