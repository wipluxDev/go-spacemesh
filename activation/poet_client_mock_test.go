// Code generated by MockGen. DO NOT EDIT.
// Source: ./nipst.go

// Package activation is a generated GoMock package.
package activation

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/spacemeshos/go-spacemesh/common/types"
	config "github.com/spacemeshos/post/config"
	shared "github.com/spacemeshos/post/shared"
)

// MockPostProverClient is a mock of PostProverClient interface.
type MockPostProverClient struct {
	ctrl     *gomock.Controller
	recorder *MockPostProverClientMockRecorder
}

// MockPostProverClientMockRecorder is the mock recorder for MockPostProverClient.
type MockPostProverClientMockRecorder struct {
	mock *MockPostProverClient
}

// NewMockPostProverClient creates a new mock instance.
func NewMockPostProverClient(ctrl *gomock.Controller) *MockPostProverClient {
	mock := &MockPostProverClient{ctrl: ctrl}
	mock.recorder = &MockPostProverClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostProverClient) EXPECT() *MockPostProverClientMockRecorder {
	return m.recorder
}

// Cfg mocks base method.
func (m *MockPostProverClient) Cfg() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cfg")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Cfg indicates an expected call of Cfg.
func (mr *MockPostProverClientMockRecorder) Cfg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cfg", reflect.TypeOf((*MockPostProverClient)(nil).Cfg))
}

// Execute mocks base method.
func (m *MockPostProverClient) Execute(challenge []byte) (*types.PostProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", challenge)
	ret0, _ := ret[0].(*types.PostProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockPostProverClientMockRecorder) Execute(challenge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockPostProverClient)(nil).Execute), challenge)
}

// Initialize mocks base method.
func (m *MockPostProverClient) Initialize() (*types.PostProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(*types.PostProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Initialize indicates an expected call of Initialize.
func (mr *MockPostProverClientMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockPostProverClient)(nil).Initialize))
}

// IsInitialized mocks base method.
func (m *MockPostProverClient) IsInitialized() (bool, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInitialized")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsInitialized indicates an expected call of IsInitialized.
func (mr *MockPostProverClientMockRecorder) IsInitialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInitialized", reflect.TypeOf((*MockPostProverClient)(nil).IsInitialized))
}

// Reset mocks base method.
func (m *MockPostProverClient) Reset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockPostProverClientMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockPostProverClient)(nil).Reset))
}

// SetLogger mocks base method.
func (m *MockPostProverClient) SetLogger(logger shared.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", logger)
}

// SetLogger indicates an expected call of SetLogger.
func (mr *MockPostProverClientMockRecorder) SetLogger(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockPostProverClient)(nil).SetLogger), logger)
}

// SetParams mocks base method.
func (m *MockPostProverClient) SetParams(datadir string, space uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetParams", datadir, space)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetParams indicates an expected call of SetParams.
func (mr *MockPostProverClientMockRecorder) SetParams(datadir, space interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParams", reflect.TypeOf((*MockPostProverClient)(nil).SetParams), datadir, space)
}

// VerifyInitAllowed mocks base method.
func (m *MockPostProverClient) VerifyInitAllowed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyInitAllowed")
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyInitAllowed indicates an expected call of VerifyInitAllowed.
func (mr *MockPostProverClientMockRecorder) VerifyInitAllowed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyInitAllowed", reflect.TypeOf((*MockPostProverClient)(nil).VerifyInitAllowed))
}

// MockPoetProvingServiceClient is a mock of PoetProvingServiceClient interface.
type MockPoetProvingServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockPoetProvingServiceClientMockRecorder
}

// MockPoetProvingServiceClientMockRecorder is the mock recorder for MockPoetProvingServiceClient.
type MockPoetProvingServiceClientMockRecorder struct {
	mock *MockPoetProvingServiceClient
}

// NewMockPoetProvingServiceClient creates a new mock instance.
func NewMockPoetProvingServiceClient(ctrl *gomock.Controller) *MockPoetProvingServiceClient {
	mock := &MockPoetProvingServiceClient{ctrl: ctrl}
	mock.recorder = &MockPoetProvingServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoetProvingServiceClient) EXPECT() *MockPoetProvingServiceClientMockRecorder {
	return m.recorder
}

// PoetServiceID mocks base method.
func (m *MockPoetProvingServiceClient) PoetServiceID(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PoetServiceID", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PoetServiceID indicates an expected call of PoetServiceID.
func (mr *MockPoetProvingServiceClientMockRecorder) PoetServiceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PoetServiceID", reflect.TypeOf((*MockPoetProvingServiceClient)(nil).PoetServiceID), arg0)
}

// Submit mocks base method.
func (m *MockPoetProvingServiceClient) Submit(ctx context.Context, challenge types.Hash32) (*types.PoetRound, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Submit", ctx, challenge)
	ret0, _ := ret[0].(*types.PoetRound)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Submit indicates an expected call of Submit.
func (mr *MockPoetProvingServiceClientMockRecorder) Submit(ctx, challenge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Submit", reflect.TypeOf((*MockPoetProvingServiceClient)(nil).Submit), ctx, challenge)
}

// MockpoetDbAPI is a mock of poetDbAPI interface.
type MockpoetDbAPI struct {
	ctrl     *gomock.Controller
	recorder *MockpoetDbAPIMockRecorder
}

// MockpoetDbAPIMockRecorder is the mock recorder for MockpoetDbAPI.
type MockpoetDbAPIMockRecorder struct {
	mock *MockpoetDbAPI
}

// NewMockpoetDbAPI creates a new mock instance.
func NewMockpoetDbAPI(ctrl *gomock.Controller) *MockpoetDbAPI {
	mock := &MockpoetDbAPI{ctrl: ctrl}
	mock.recorder = &MockpoetDbAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpoetDbAPI) EXPECT() *MockpoetDbAPIMockRecorder {
	return m.recorder
}

// GetMembershipMap mocks base method.
func (m *MockpoetDbAPI) GetMembershipMap(proofRef []byte) (map[types.Hash32]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembershipMap", proofRef)
	ret0, _ := ret[0].(map[types.Hash32]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembershipMap indicates an expected call of GetMembershipMap.
func (mr *MockpoetDbAPIMockRecorder) GetMembershipMap(proofRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembershipMap", reflect.TypeOf((*MockpoetDbAPI)(nil).GetMembershipMap), proofRef)
}

// SubscribeToProofRef mocks base method.
func (m *MockpoetDbAPI) SubscribeToProofRef(poetID []byte, roundID string) chan []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToProofRef", poetID, roundID)
	ret0, _ := ret[0].(chan []byte)
	return ret0
}

// SubscribeToProofRef indicates an expected call of SubscribeToProofRef.
func (mr *MockpoetDbAPIMockRecorder) SubscribeToProofRef(poetID, roundID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToProofRef", reflect.TypeOf((*MockpoetDbAPI)(nil).SubscribeToProofRef), poetID, roundID)
}

// UnsubscribeFromProofRef mocks base method.
func (m *MockpoetDbAPI) UnsubscribeFromProofRef(poetID []byte, roundID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsubscribeFromProofRef", poetID, roundID)
}

// UnsubscribeFromProofRef indicates an expected call of UnsubscribeFromProofRef.
func (mr *MockpoetDbAPIMockRecorder) UnsubscribeFromProofRef(poetID, roundID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromProofRef", reflect.TypeOf((*MockpoetDbAPI)(nil).UnsubscribeFromProofRef), poetID, roundID)
}
