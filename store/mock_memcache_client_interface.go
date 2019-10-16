// Code generated by mockery v1.0.0. DO NOT EDIT.

package store

import memcache "github.com/bradfitz/gomemcache/memcache"
import mock "github.com/stretchr/testify/mock"

// MockMemcacheClientInterface is an autogenerated mock type for the MemcacheClientInterface type
type MockMemcacheClientInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: item
func (_m *MockMemcacheClientInterface) Delete(item string) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *MockMemcacheClientInterface) Get(key string) (*memcache.Item, error) {
	ret := _m.Called(key)

	var r0 *memcache.Item
	if rf, ok := ret.Get(0).(func(string) *memcache.Item); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*memcache.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: item
func (_m *MockMemcacheClientInterface) Set(item *memcache.Item) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*memcache.Item) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
