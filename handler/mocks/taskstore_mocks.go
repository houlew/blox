// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the License). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the license file accompanying this file. This file is distributed
// on an AS IS BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: handler/store/taskstore.go

package mocks

import (
	context "context"
	types0 "github.com/aws/amazon-ecs-event-stream-handler/handler/store/types"
	types "github.com/aws/amazon-ecs-event-stream-handler/handler/types"
	gomock "github.com/golang/mock/gomock"
)

// Mock of TaskStore interface
type MockTaskStore struct {
	ctrl     *gomock.Controller
	recorder *_MockTaskStoreRecorder
}

// Recorder for MockTaskStore (not exported)
type _MockTaskStoreRecorder struct {
	mock *MockTaskStore
}

func NewMockTaskStore(ctrl *gomock.Controller) *MockTaskStore {
	mock := &MockTaskStore{ctrl: ctrl}
	mock.recorder = &_MockTaskStoreRecorder{mock}
	return mock
}

func (_m *MockTaskStore) EXPECT() *_MockTaskStoreRecorder {
	return _m.recorder
}

func (_m *MockTaskStore) AddTask(task string) error {
	ret := _m.ctrl.Call(_m, "AddTask", task)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTaskStoreRecorder) AddTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddTask", arg0)
}

func (_m *MockTaskStore) GetTask(cluster string, taskARN string) (*types.Task, error) {
	ret := _m.ctrl.Call(_m, "GetTask", cluster, taskARN)
	ret0, _ := ret[0].(*types.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTaskStoreRecorder) GetTask(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTask", arg0, arg1)
}

func (_m *MockTaskStore) ListTasks() ([]types.Task, error) {
	ret := _m.ctrl.Call(_m, "ListTasks")
	ret0, _ := ret[0].([]types.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTaskStoreRecorder) ListTasks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasks")
}

func (_m *MockTaskStore) FilterTasks(filterKey string, filterValue string) ([]types.Task, error) {
	ret := _m.ctrl.Call(_m, "FilterTasks", filterKey, filterValue)
	ret0, _ := ret[0].([]types.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTaskStoreRecorder) FilterTasks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilterTasks", arg0, arg1)
}

func (_m *MockTaskStore) StreamTasks(ctx context.Context) (chan types0.TaskErrorWrapper, error) {
	ret := _m.ctrl.Call(_m, "StreamTasks", ctx)
	ret0, _ := ret[0].(chan types0.TaskErrorWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTaskStoreRecorder) StreamTasks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StreamTasks", arg0)
}
