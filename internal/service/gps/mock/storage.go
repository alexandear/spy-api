// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"
import model "github.com/devchallenge/spy-api/internal/model"
import time "time"

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// Save provides a mock function with given fields: phone, coordinate, timestamp
func (_m *Storage) Save(phone model.Phone, coordinate model.Coordinate, timestamp time.Time) error {
	ret := _m.Called(phone, coordinate, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Phone, model.Coordinate, time.Time) error); ok {
		r0 = rf(phone, coordinate, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
