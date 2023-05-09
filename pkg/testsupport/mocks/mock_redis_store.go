// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/redis_store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockRedisStore is a mock of RedisStore interface.
type MockRedisStore struct {
	ctrl     *gomock.Controller
	recorder *MockRedisStoreMockRecorder
}

// MockRedisStoreMockRecorder is the mock recorder for MockRedisStore.
type MockRedisStoreMockRecorder struct {
	mock *MockRedisStore
}

// NewMockRedisStore creates a new mock instance.
func NewMockRedisStore(ctrl *gomock.Controller) *MockRedisStore {
	mock := &MockRedisStore{ctrl: ctrl}
	mock.recorder = &MockRedisStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisStore) EXPECT() *MockRedisStoreMockRecorder {
	return m.recorder
}

// BLMove mocks base method.
func (m *MockRedisStore) BLMove(ctx context.Context, timeout time.Duration, sourceQueueKey, destQueueKey string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BLMove", ctx, timeout, sourceQueueKey, destQueueKey)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BLMove indicates an expected call of BLMove.
func (mr *MockRedisStoreMockRecorder) BLMove(ctx, timeout, sourceQueueKey, destQueueKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BLMove", reflect.TypeOf((*MockRedisStore)(nil).BLMove), ctx, timeout, sourceQueueKey, destQueueKey)
}

// Dequeue mocks base method.
func (m *MockRedisStore) Dequeue(ctx context.Context, timeout time.Duration, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dequeue", ctx, timeout, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dequeue indicates an expected call of Dequeue.
func (mr *MockRedisStoreMockRecorder) Dequeue(ctx, timeout, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dequeue", reflect.TypeOf((*MockRedisStore)(nil).Dequeue), ctx, timeout, key)
}

// Enqueue mocks base method.
func (m *MockRedisStore) Enqueue(ctx context.Context, key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enqueue", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enqueue indicates an expected call of Enqueue.
func (mr *MockRedisStoreMockRecorder) Enqueue(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enqueue", reflect.TypeOf((*MockRedisStore)(nil).Enqueue), ctx, key, value)
}

// Get mocks base method.
func (m *MockRedisStore) Get(ctx context.Context, messageKey string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, messageKey)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRedisStoreMockRecorder) Get(ctx, messageKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisStore)(nil).Get), ctx, messageKey)
}

// SetAndEnqueue mocks base method.
func (m *MockRedisStore) SetAndEnqueue(ctx context.Context, messageKey string, value []byte, queueKey, messageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAndEnqueue", ctx, messageKey, value, queueKey, messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAndEnqueue indicates an expected call of SetAndEnqueue.
func (mr *MockRedisStoreMockRecorder) SetAndEnqueue(ctx, messageKey, value, queueKey, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAndEnqueue", reflect.TypeOf((*MockRedisStore)(nil).SetAndEnqueue), ctx, messageKey, value, queueKey, messageID)
}

// SetAndMove mocks base method.
func (m *MockRedisStore) SetAndMove(ctx context.Context, messageKey string, value []byte, sourceQueueKey, destQueueKey, messageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAndMove", ctx, messageKey, value, sourceQueueKey, destQueueKey, messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAndMove indicates an expected call of SetAndMove.
func (mr *MockRedisStoreMockRecorder) SetAndMove(ctx, messageKey, value, sourceQueueKey, destQueueKey, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAndMove", reflect.TypeOf((*MockRedisStore)(nil).SetAndMove), ctx, messageKey, value, sourceQueueKey, destQueueKey, messageID)
}

// SetAndZAdd mocks base method.
func (m *MockRedisStore) SetAndZAdd(ctx context.Context, messageKey string, value []byte, queueKey, messageID string, score float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAndZAdd", ctx, messageKey, value, queueKey, messageID, score)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAndZAdd indicates an expected call of SetAndZAdd.
func (mr *MockRedisStoreMockRecorder) SetAndZAdd(ctx, messageKey, value, queueKey, messageID, score interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAndZAdd", reflect.TypeOf((*MockRedisStore)(nil).SetAndZAdd), ctx, messageKey, value, queueKey, messageID, score)
}
