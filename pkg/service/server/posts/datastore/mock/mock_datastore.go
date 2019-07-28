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
// Source: pkg/service/server/posts/datastore/datastore.go

// Package mock_datastore is a generated GoMock package.
package mock_datastore

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/prksu/publr/pkg/api/posts/v1alpha2"
	reflect "reflect"
)

// MockPostDatastore is a mock of PostDatastore interface
type MockPostDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockPostDatastoreMockRecorder
}

// MockPostDatastoreMockRecorder is the mock recorder for MockPostDatastore
type MockPostDatastoreMockRecorder struct {
	mock *MockPostDatastore
}

// NewMockPostDatastore creates a new mock instance
func NewMockPostDatastore(ctrl *gomock.Controller) *MockPostDatastore {
	mock := &MockPostDatastore{ctrl: ctrl}
	mock.recorder = &MockPostDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostDatastore) EXPECT() *MockPostDatastoreMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockPostDatastore) List(ctx context.Context, sitedomain, author string, start, limit int) ([]*v1alpha2.Post, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, sitedomain, author, start, limit)
	ret0, _ := ret[0].([]*v1alpha2.Post)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockPostDatastoreMockRecorder) List(ctx, sitedomain, author, start, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPostDatastore)(nil).List), ctx, sitedomain, author, start, limit)
}

// Create mocks base method
func (m *MockPostDatastore) Create(ctx context.Context, sitedomain, author string, post *v1alpha2.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, sitedomain, author, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockPostDatastoreMockRecorder) Create(ctx, sitedomain, author, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostDatastore)(nil).Create), ctx, sitedomain, author, post)
}

// Get mocks base method
func (m *MockPostDatastore) Get(ctx context.Context, sitedomain, author, slug string) (*v1alpha2.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, sitedomain, author, slug)
	ret0, _ := ret[0].(*v1alpha2.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPostDatastoreMockRecorder) Get(ctx, sitedomain, author, slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostDatastore)(nil).Get), ctx, sitedomain, author, slug)
}

// Update mocks base method
func (m *MockPostDatastore) Update(ctx context.Context, sitedomain, author, slug string, post *v1alpha2.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, sitedomain, author, slug, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockPostDatastoreMockRecorder) Update(ctx, sitedomain, author, slug, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostDatastore)(nil).Update), ctx, sitedomain, author, slug, post)
}

// Delete mocks base method
func (m *MockPostDatastore) Delete(ctx context.Context, sitedomain, author, slug string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, sitedomain, author, slug)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPostDatastoreMockRecorder) Delete(ctx, sitedomain, author, slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostDatastore)(nil).Delete), ctx, sitedomain, author, slug)
}

// Search mocks base method
func (m *MockPostDatastore) Search(ctx context.Context, query string) ([]*v1alpha2.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, query)
	ret0, _ := ret[0].([]*v1alpha2.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockPostDatastoreMockRecorder) Search(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockPostDatastore)(nil).Search), ctx, query)
}