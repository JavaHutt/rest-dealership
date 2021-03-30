// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	model "github.com/JavaHutt/rest-dealership/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockVehicle is a mock of Vehicle interface.
type MockVehicle struct {
	ctrl     *gomock.Controller
	recorder *MockVehicleMockRecorder
}

// MockVehicleMockRecorder is the mock recorder for MockVehicle.
type MockVehicleMockRecorder struct {
	mock *MockVehicle
}

// NewMockVehicle creates a new mock instance.
func NewMockVehicle(ctrl *gomock.Controller) *MockVehicle {
	mock := &MockVehicle{ctrl: ctrl}
	mock.recorder = &MockVehicleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVehicle) EXPECT() *MockVehicleMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVehicle) Create(arg0 model.Vehicle) (*model.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*model.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVehicleMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVehicle)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockVehicle) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVehicleMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVehicle)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockVehicle) Get(id int) (*model.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVehicleMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVehicle)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockVehicle) GetAll() []model.Vehicle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Vehicle)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockVehicleMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockVehicle)(nil).GetAll))
}

// SeedData mocks base method.
func (m *MockVehicle) SeedData(vehicles []model.Vehicle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeedData", vehicles)
	ret0, _ := ret[0].(error)
	return ret0
}

// SeedData indicates an expected call of SeedData.
func (mr *MockVehicleMockRecorder) SeedData(vehicles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeedData", reflect.TypeOf((*MockVehicle)(nil).SeedData), vehicles)
}

// Update mocks base method.
func (m *MockVehicle) Update(id int, vehicle model.Vehicle) (*model.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, vehicle)
	ret0, _ := ret[0].(*model.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVehicleMockRecorder) Update(id, vehicle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVehicle)(nil).Update), id, vehicle)
}
