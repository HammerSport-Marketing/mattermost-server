// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	cache "github.com/mattermost/mattermost-server/v5/services/cache"
	mock "github.com/stretchr/testify/mock"
)

// CacheProvider is an autogenerated mock type for the CacheProvider type
type CacheProvider struct {
	mock.Mock
}

// NewCache provides a mock function with given fields: size
func (_m *CacheProvider) NewCache(size int) cache.Cache {
	ret := _m.Called(size)

	var r0 cache.Cache
	if rf, ok := ret.Get(0).(func(int) cache.Cache); ok {
		r0 = rf(size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Cache)
		}
	}

	return r0
}

// Connect provides a mock function with given fields:
func (_m *CacheProvider) Connect() {
	_m.Called()
}

// Close provides a mock function with given fields:
func (_m *CacheProvider) Close() {
	_m.Called()
}

// NewCacheWithParams provides a mock function with given fields: size, name, defaultExpiry, invalidateClusterEvent
func (_m *CacheProvider) NewCacheWithParams(size int, name string, defaultExpiry int64, invalidateClusterEvent string) cache.Cache {
	ret := _m.Called(size, name, defaultExpiry, invalidateClusterEvent)

	var r0 cache.Cache
	if rf, ok := ret.Get(0).(func(int, string, int64, string) cache.Cache); ok {
		r0 = rf(size, name, defaultExpiry, invalidateClusterEvent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Cache)
		}
	}

	return r0
}
