// Copyright 2019 Publr Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prksu/publr/service/sites (interfaces: SiteDatastore)

// Package mock_sites is a generated GoMock package.
package mock_sites

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/prksu/publr/pkg/api/sites/v1alpha3"
	reflect "reflect"
)

// MockSiteDatastore is a mock of SiteDatastore interface
type MockSiteDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockSiteDatastoreMockRecorder
}

// MockSiteDatastoreMockRecorder is the mock recorder for MockSiteDatastore
type MockSiteDatastoreMockRecorder struct {
	mock *MockSiteDatastore
}

// NewMockSiteDatastore creates a new mock instance
func NewMockSiteDatastore(ctrl *gomock.Controller) *MockSiteDatastore {
	mock := &MockSiteDatastore{ctrl: ctrl}
	mock.recorder = &MockSiteDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSiteDatastore) EXPECT() *MockSiteDatastoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockSiteDatastore) Create(arg0 context.Context, arg1 *v1alpha3.Site) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockSiteDatastoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSiteDatastore)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockSiteDatastore) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockSiteDatastoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSiteDatastore)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockSiteDatastore) Get(arg0 context.Context, arg1 string) (*v1alpha3.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha3.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSiteDatastoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSiteDatastore)(nil).Get), arg0, arg1)
}
