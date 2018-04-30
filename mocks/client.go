// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber-go/dosa (interfaces: Client,AdminClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	dosa "github.com/uber-go/dosa"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateIfNotExists mocks base method
func (m *MockClient) CreateIfNotExists(arg0 context.Context, arg1 dosa.DomainObject) error {
	ret := m.ctrl.Call(m, "CreateIfNotExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIfNotExists indicates an expected call of CreateIfNotExists
func (mr *MockClientMockRecorder) CreateIfNotExists(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIfNotExists", reflect.TypeOf((*MockClient)(nil).CreateIfNotExists), arg0, arg1)
}

// GetRegistrar mocks base method
func (m *MockClient) GetRegistrar() dosa.Registrar {
	ret := m.ctrl.Call(m, "GetRegistrar")
	ret0, _ := ret[0].(dosa.Registrar)
	return ret0
}

// GetRegistrar indicates an expected call of GetRegistrar
func (mr *MockClientMockRecorder) GetRegistrar() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistrar", reflect.TypeOf((*MockClient)(nil).GetRegistrar))
}

// Initialize mocks base method
func (m *MockClient) Initialize(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockClientMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockClient)(nil).Initialize), arg0)
}

// MultiRead mocks base method
func (m *MockClient) MultiRead(arg0 context.Context, arg1 []string, arg2 ...dosa.DomainObject) (dosa.MultiResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiRead", varargs...)
	ret0, _ := ret[0].(dosa.MultiResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiRead indicates an expected call of MultiRead
func (mr *MockClientMockRecorder) MultiRead(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiRead", reflect.TypeOf((*MockClient)(nil).MultiRead), varargs...)
}

// Range mocks base method
func (m *MockClient) Range(arg0 context.Context, arg1 *dosa.RangeOp) ([]dosa.DomainObject, string, error) {
	ret := m.ctrl.Call(m, "Range", arg0, arg1)
	ret0, _ := ret[0].([]dosa.DomainObject)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Range indicates an expected call of Range
func (mr *MockClientMockRecorder) Range(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Range", reflect.TypeOf((*MockClient)(nil).Range), arg0, arg1)
}

// Read mocks base method
func (m *MockClient) Read(arg0 context.Context, arg1 []string, arg2 dosa.DomainObject) error {
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read
func (mr *MockClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockClient)(nil).Read), arg0, arg1, arg2)
}

// Remove mocks base method
func (m *MockClient) Remove(arg0 context.Context, arg1 dosa.DomainObject) error {
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockClientMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockClient)(nil).Remove), arg0, arg1)
}

// RemoveRange mocks base method
func (m *MockClient) RemoveRange(arg0 context.Context, arg1 *dosa.RemoveRangeOp) error {
	ret := m.ctrl.Call(m, "RemoveRange", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRange indicates an expected call of RemoveRange
func (mr *MockClientMockRecorder) RemoveRange(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRange", reflect.TypeOf((*MockClient)(nil).RemoveRange), arg0, arg1)
}

// ScanEverything mocks base method
func (m *MockClient) ScanEverything(arg0 context.Context, arg1 *dosa.ScanOp) ([]dosa.DomainObject, string, error) {
	ret := m.ctrl.Call(m, "ScanEverything", arg0, arg1)
	ret0, _ := ret[0].([]dosa.DomainObject)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ScanEverything indicates an expected call of ScanEverything
func (mr *MockClientMockRecorder) ScanEverything(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanEverything", reflect.TypeOf((*MockClient)(nil).ScanEverything), arg0, arg1)
}

// Upsert mocks base method
func (m *MockClient) Upsert(arg0 context.Context, arg1 []string, arg2 dosa.DomainObject) error {
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockClientMockRecorder) Upsert(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockClient)(nil).Upsert), arg0, arg1, arg2)
}

// WalkRange mocks base method
func (m *MockClient) WalkRange(arg0 context.Context, arg1 *dosa.RangeOp, arg2 func(dosa.DomainObject) error) error {
	ret := m.ctrl.Call(m, "WalkRange", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkRange indicates an expected call of WalkRange
func (mr *MockClientMockRecorder) WalkRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkRange", reflect.TypeOf((*MockClient)(nil).WalkRange), arg0, arg1, arg2)
}

// MockAdminClient is a mock of AdminClient interface
type MockAdminClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdminClientMockRecorder
}

// MockAdminClientMockRecorder is the mock recorder for MockAdminClient
type MockAdminClientMockRecorder struct {
	mock *MockAdminClient
}

// NewMockAdminClient creates a new mock instance
func NewMockAdminClient(ctrl *gomock.Controller) *MockAdminClient {
	mock := &MockAdminClient{ctrl: ctrl}
	mock.recorder = &MockAdminClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdminClient) EXPECT() *MockAdminClientMockRecorder {
	return m.recorder
}

// CanUpsertSchema mocks base method
func (m *MockAdminClient) CanUpsertSchema(arg0 context.Context, arg1 string) (*dosa.SchemaStatus, error) {
	ret := m.ctrl.Call(m, "CanUpsertSchema", arg0, arg1)
	ret0, _ := ret[0].(*dosa.SchemaStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanUpsertSchema indicates an expected call of CanUpsertSchema
func (mr *MockAdminClientMockRecorder) CanUpsertSchema(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanUpsertSchema", reflect.TypeOf((*MockAdminClient)(nil).CanUpsertSchema), arg0, arg1)
}

// CheckSchemaStatus mocks base method
func (m *MockAdminClient) CheckSchemaStatus(arg0 context.Context, arg1 string, arg2 int32) (*dosa.SchemaStatus, error) {
	ret := m.ctrl.Call(m, "CheckSchemaStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dosa.SchemaStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSchemaStatus indicates an expected call of CheckSchemaStatus
func (mr *MockAdminClientMockRecorder) CheckSchemaStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSchemaStatus", reflect.TypeOf((*MockAdminClient)(nil).CheckSchemaStatus), arg0, arg1, arg2)
}

// CreateScope mocks base method
func (m *MockAdminClient) CreateScope(arg0 context.Context, arg1 *dosa.ScopeMetadata) error {
	ret := m.ctrl.Call(m, "CreateScope", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateScope indicates an expected call of CreateScope
func (mr *MockAdminClientMockRecorder) CreateScope(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScope", reflect.TypeOf((*MockAdminClient)(nil).CreateScope), arg0, arg1)
}

// Directories mocks base method
func (m *MockAdminClient) Directories(arg0 []string) dosa.AdminClient {
	ret := m.ctrl.Call(m, "Directories", arg0)
	ret0, _ := ret[0].(dosa.AdminClient)
	return ret0
}

// Directories indicates an expected call of Directories
func (mr *MockAdminClientMockRecorder) Directories(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Directories", reflect.TypeOf((*MockAdminClient)(nil).Directories), arg0)
}

// DropScope mocks base method
func (m *MockAdminClient) DropScope(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "DropScope", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropScope indicates an expected call of DropScope
func (mr *MockAdminClientMockRecorder) DropScope(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropScope", reflect.TypeOf((*MockAdminClient)(nil).DropScope), arg0, arg1)
}

// Excludes mocks base method
func (m *MockAdminClient) Excludes(arg0 []string) dosa.AdminClient {
	ret := m.ctrl.Call(m, "Excludes", arg0)
	ret0, _ := ret[0].(dosa.AdminClient)
	return ret0
}

// Excludes indicates an expected call of Excludes
func (mr *MockAdminClientMockRecorder) Excludes(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Excludes", reflect.TypeOf((*MockAdminClient)(nil).Excludes), arg0)
}

// GetSchema mocks base method
func (m *MockAdminClient) GetSchema() ([]*dosa.EntityDefinition, error) {
	ret := m.ctrl.Call(m, "GetSchema")
	ret0, _ := ret[0].([]*dosa.EntityDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchema indicates an expected call of GetSchema
func (mr *MockAdminClientMockRecorder) GetSchema() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchema", reflect.TypeOf((*MockAdminClient)(nil).GetSchema))
}

// Scope mocks base method
func (m *MockAdminClient) Scope(arg0 string) dosa.AdminClient {
	ret := m.ctrl.Call(m, "Scope", arg0)
	ret0, _ := ret[0].(dosa.AdminClient)
	return ret0
}

// Scope indicates an expected call of Scope
func (mr *MockAdminClientMockRecorder) Scope(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scope", reflect.TypeOf((*MockAdminClient)(nil).Scope), arg0)
}

// TruncateScope mocks base method
func (m *MockAdminClient) TruncateScope(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "TruncateScope", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TruncateScope indicates an expected call of TruncateScope
func (mr *MockAdminClientMockRecorder) TruncateScope(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateScope", reflect.TypeOf((*MockAdminClient)(nil).TruncateScope), arg0, arg1)
}

// UpsertSchema mocks base method
func (m *MockAdminClient) UpsertSchema(arg0 context.Context, arg1 string) (*dosa.SchemaStatus, error) {
	ret := m.ctrl.Call(m, "UpsertSchema", arg0, arg1)
	ret0, _ := ret[0].(*dosa.SchemaStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertSchema indicates an expected call of UpsertSchema
func (mr *MockAdminClientMockRecorder) UpsertSchema(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSchema", reflect.TypeOf((*MockAdminClient)(nil).UpsertSchema), arg0, arg1)
}
