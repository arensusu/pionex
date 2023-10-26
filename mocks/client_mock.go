// Code generated by MockGen. DO NOT EDIT.
// Source: domain/client.go
//
// Generated by this command:
//
//	mockgen -source=domain/client.go -destination=mocks/client_mock.go
//
// Package mock_domain is a generated GoMock package.
package mocks

import (
	url "net/url"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// HttpGet mocks base method.
func (m *MockClient) HttpGet(api string, param url.Values) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HttpGet", api, param)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HttpGet indicates an expected call of HttpGet.
func (mr *MockClientMockRecorder) HttpGet(api, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HttpGet", reflect.TypeOf((*MockClient)(nil).HttpGet), api, param)
}

// HttpPost mocks base method.
func (m *MockClient) HttpPost(api string, body any) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HttpPost", api, body)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HttpPost indicates an expected call of HttpPost.
func (mr *MockClientMockRecorder) HttpPost(api, body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HttpPost", reflect.TypeOf((*MockClient)(nil).HttpPost), api, body)
}

// Sign mocks base method.
func (m *MockClient) Sign(method, api, timestamp string, param url.Values, body string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", method, api, timestamp, param, body)
	ret0, _ := ret[0].(string)
	return ret0
}

// Sign indicates an expected call of Sign.
func (mr *MockClientMockRecorder) Sign(method, api, timestamp, param, body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockClient)(nil).Sign), method, api, timestamp, param, body)
}
