// Automatically generated by MockGen. DO NOT EDIT!
// Source: heka/plugins/amqp (interfaces: AMQPConnection)

package amqp

import (
    gomock "github.com/rafrombrc/gomock/gomock"
    amqp "github.com/streadway/amqp"
)

// Mock of AMQPConnection interface
type MockAMQPConnection struct {
	ctrl     *gomock.Controller
	recorder *_MockAMQPConnectionRecorder
}

// Recorder for MockAMQPConnection (not exported)
type _MockAMQPConnectionRecorder struct {
	mock *MockAMQPConnection
}

func NewMockAMQPConnection(ctrl *gomock.Controller) *MockAMQPConnection {
	mock := &MockAMQPConnection{ctrl: ctrl}
	mock.recorder = &_MockAMQPConnectionRecorder{mock}
	return mock
}

func (_m *MockAMQPConnection) EXPECT() *_MockAMQPConnectionRecorder {
	return _m.recorder
}

func (_m *MockAMQPConnection) Channel() (*amqp.Channel, error) {
	ret := _m.ctrl.Call(_m, "Channel")
	ret0, _ := ret[0].(*amqp.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAMQPConnectionRecorder) Channel() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Channel")
}

func (_m *MockAMQPConnection) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAMQPConnectionRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockAMQPConnection) NotifyClose(_param0 chan *amqp.Error) chan *amqp.Error {
	ret := _m.ctrl.Call(_m, "NotifyClose", _param0)
	ret0, _ := ret[0].(chan *amqp.Error)
	return ret0
}

func (_mr *_MockAMQPConnectionRecorder) NotifyClose(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NotifyClose", arg0)
}
