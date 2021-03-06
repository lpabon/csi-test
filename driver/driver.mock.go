// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../vendor/github.com/container-storage-interface/spec/lib/go/csi/csi.pb.go

package driver

import (
	. "github.com/container-storage-interface/spec/lib/go/csi"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Mock of isVolumeCapability_AccessType interface
type MockisVolumeCapability_AccessType struct {
	ctrl     *gomock.Controller
	recorder *_MockisVolumeCapability_AccessTypeRecorder
}

// Recorder for MockisVolumeCapability_AccessType (not exported)
type _MockisVolumeCapability_AccessTypeRecorder struct {
	mock *MockisVolumeCapability_AccessType
}

func NewMockisVolumeCapability_AccessType(ctrl *gomock.Controller) *MockisVolumeCapability_AccessType {
	mock := &MockisVolumeCapability_AccessType{ctrl: ctrl}
	mock.recorder = &_MockisVolumeCapability_AccessTypeRecorder{mock}
	return mock
}

func (_m *MockisVolumeCapability_AccessType) EXPECT() *_MockisVolumeCapability_AccessTypeRecorder {
	return _m.recorder
}

func (_m *MockisVolumeCapability_AccessType) isVolumeCapability_AccessType() {
	_m.ctrl.Call(_m, "isVolumeCapability_AccessType")
}

func (_mr *_MockisVolumeCapability_AccessTypeRecorder) isVolumeCapability_AccessType() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "isVolumeCapability_AccessType")
}

// Mock of isControllerServiceCapability_Type interface
type MockisControllerServiceCapability_Type struct {
	ctrl     *gomock.Controller
	recorder *_MockisControllerServiceCapability_TypeRecorder
}

// Recorder for MockisControllerServiceCapability_Type (not exported)
type _MockisControllerServiceCapability_TypeRecorder struct {
	mock *MockisControllerServiceCapability_Type
}

func NewMockisControllerServiceCapability_Type(ctrl *gomock.Controller) *MockisControllerServiceCapability_Type {
	mock := &MockisControllerServiceCapability_Type{ctrl: ctrl}
	mock.recorder = &_MockisControllerServiceCapability_TypeRecorder{mock}
	return mock
}

func (_m *MockisControllerServiceCapability_Type) EXPECT() *_MockisControllerServiceCapability_TypeRecorder {
	return _m.recorder
}

func (_m *MockisControllerServiceCapability_Type) isControllerServiceCapability_Type() {
	_m.ctrl.Call(_m, "isControllerServiceCapability_Type")
}

func (_mr *_MockisControllerServiceCapability_TypeRecorder) isControllerServiceCapability_Type() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "isControllerServiceCapability_Type")
}

// Mock of isNodeServiceCapability_Type interface
type MockisNodeServiceCapability_Type struct {
	ctrl     *gomock.Controller
	recorder *_MockisNodeServiceCapability_TypeRecorder
}

// Recorder for MockisNodeServiceCapability_Type (not exported)
type _MockisNodeServiceCapability_TypeRecorder struct {
	mock *MockisNodeServiceCapability_Type
}

func NewMockisNodeServiceCapability_Type(ctrl *gomock.Controller) *MockisNodeServiceCapability_Type {
	mock := &MockisNodeServiceCapability_Type{ctrl: ctrl}
	mock.recorder = &_MockisNodeServiceCapability_TypeRecorder{mock}
	return mock
}

func (_m *MockisNodeServiceCapability_Type) EXPECT() *_MockisNodeServiceCapability_TypeRecorder {
	return _m.recorder
}

func (_m *MockisNodeServiceCapability_Type) isNodeServiceCapability_Type() {
	_m.ctrl.Call(_m, "isNodeServiceCapability_Type")
}

func (_mr *_MockisNodeServiceCapability_TypeRecorder) isNodeServiceCapability_Type() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "isNodeServiceCapability_Type")
}

// Mock of IdentityClient interface
type MockIdentityClient struct {
	ctrl     *gomock.Controller
	recorder *_MockIdentityClientRecorder
}

// Recorder for MockIdentityClient (not exported)
type _MockIdentityClientRecorder struct {
	mock *MockIdentityClient
}

func NewMockIdentityClient(ctrl *gomock.Controller) *MockIdentityClient {
	mock := &MockIdentityClient{ctrl: ctrl}
	mock.recorder = &_MockIdentityClientRecorder{mock}
	return mock
}

func (_m *MockIdentityClient) EXPECT() *_MockIdentityClientRecorder {
	return _m.recorder
}

func (_m *MockIdentityClient) GetSupportedVersions(ctx context.Context, in *GetSupportedVersionsRequest, opts ...grpc.CallOption) (*GetSupportedVersionsResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetSupportedVersions", _s...)
	ret0, _ := ret[0].(*GetSupportedVersionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIdentityClientRecorder) GetSupportedVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSupportedVersions", _s...)
}

func (_m *MockIdentityClient) GetPluginInfo(ctx context.Context, in *GetPluginInfoRequest, opts ...grpc.CallOption) (*GetPluginInfoResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetPluginInfo", _s...)
	ret0, _ := ret[0].(*GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIdentityClientRecorder) GetPluginInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPluginInfo", _s...)
}

// Mock of IdentityServer interface
type MockIdentityServer struct {
	ctrl     *gomock.Controller
	recorder *_MockIdentityServerRecorder
}

// Recorder for MockIdentityServer (not exported)
type _MockIdentityServerRecorder struct {
	mock *MockIdentityServer
}

func NewMockIdentityServer(ctrl *gomock.Controller) *MockIdentityServer {
	mock := &MockIdentityServer{ctrl: ctrl}
	mock.recorder = &_MockIdentityServerRecorder{mock}
	return mock
}

func (_m *MockIdentityServer) EXPECT() *_MockIdentityServerRecorder {
	return _m.recorder
}

func (_m *MockIdentityServer) GetSupportedVersions(_param0 context.Context, _param1 *GetSupportedVersionsRequest) (*GetSupportedVersionsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetSupportedVersions", _param0, _param1)
	ret0, _ := ret[0].(*GetSupportedVersionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIdentityServerRecorder) GetSupportedVersions(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSupportedVersions", arg0, arg1)
}

func (_m *MockIdentityServer) GetPluginInfo(_param0 context.Context, _param1 *GetPluginInfoRequest) (*GetPluginInfoResponse, error) {
	ret := _m.ctrl.Call(_m, "GetPluginInfo", _param0, _param1)
	ret0, _ := ret[0].(*GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIdentityServerRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPluginInfo", arg0, arg1)
}

// Mock of ControllerClient interface
type MockControllerClient struct {
	ctrl     *gomock.Controller
	recorder *_MockControllerClientRecorder
}

// Recorder for MockControllerClient (not exported)
type _MockControllerClientRecorder struct {
	mock *MockControllerClient
}

func NewMockControllerClient(ctrl *gomock.Controller) *MockControllerClient {
	mock := &MockControllerClient{ctrl: ctrl}
	mock.recorder = &_MockControllerClientRecorder{mock}
	return mock
}

func (_m *MockControllerClient) EXPECT() *_MockControllerClientRecorder {
	return _m.recorder
}

func (_m *MockControllerClient) CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CreateVolume", _s...)
	ret0, _ := ret[0].(*CreateVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) CreateVolume(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateVolume", _s...)
}

func (_m *MockControllerClient) DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*DeleteVolumeResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DeleteVolume", _s...)
	ret0, _ := ret[0].(*DeleteVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) DeleteVolume(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteVolume", _s...)
}

func (_m *MockControllerClient) ControllerPublishVolume(ctx context.Context, in *ControllerPublishVolumeRequest, opts ...grpc.CallOption) (*ControllerPublishVolumeResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ControllerPublishVolume", _s...)
	ret0, _ := ret[0].(*ControllerPublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) ControllerPublishVolume(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerPublishVolume", _s...)
}

func (_m *MockControllerClient) ControllerUnpublishVolume(ctx context.Context, in *ControllerUnpublishVolumeRequest, opts ...grpc.CallOption) (*ControllerUnpublishVolumeResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ControllerUnpublishVolume", _s...)
	ret0, _ := ret[0].(*ControllerUnpublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) ControllerUnpublishVolume(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerUnpublishVolume", _s...)
}

func (_m *MockControllerClient) ValidateVolumeCapabilities(ctx context.Context, in *ValidateVolumeCapabilitiesRequest, opts ...grpc.CallOption) (*ValidateVolumeCapabilitiesResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ValidateVolumeCapabilities", _s...)
	ret0, _ := ret[0].(*ValidateVolumeCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) ValidateVolumeCapabilities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateVolumeCapabilities", _s...)
}

func (_m *MockControllerClient) ListVolumes(ctx context.Context, in *ListVolumesRequest, opts ...grpc.CallOption) (*ListVolumesResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListVolumes", _s...)
	ret0, _ := ret[0].(*ListVolumesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) ListVolumes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListVolumes", _s...)
}

func (_m *MockControllerClient) GetCapacity(ctx context.Context, in *GetCapacityRequest, opts ...grpc.CallOption) (*GetCapacityResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetCapacity", _s...)
	ret0, _ := ret[0].(*GetCapacityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) GetCapacity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCapacity", _s...)
}

func (_m *MockControllerClient) ControllerProbe(ctx context.Context, in *ControllerProbeRequest, opts ...grpc.CallOption) (*ControllerProbeResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ControllerProbe", _s...)
	ret0, _ := ret[0].(*ControllerProbeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) ControllerProbe(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerProbe", _s...)
}

func (_m *MockControllerClient) ControllerGetCapabilities(ctx context.Context, in *ControllerGetCapabilitiesRequest, opts ...grpc.CallOption) (*ControllerGetCapabilitiesResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ControllerGetCapabilities", _s...)
	ret0, _ := ret[0].(*ControllerGetCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerClientRecorder) ControllerGetCapabilities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerGetCapabilities", _s...)
}

// Mock of ControllerServer interface
type MockControllerServer struct {
	ctrl     *gomock.Controller
	recorder *_MockControllerServerRecorder
}

// Recorder for MockControllerServer (not exported)
type _MockControllerServerRecorder struct {
	mock *MockControllerServer
}

func NewMockControllerServer(ctrl *gomock.Controller) *MockControllerServer {
	mock := &MockControllerServer{ctrl: ctrl}
	mock.recorder = &_MockControllerServerRecorder{mock}
	return mock
}

func (_m *MockControllerServer) EXPECT() *_MockControllerServerRecorder {
	return _m.recorder
}

func (_m *MockControllerServer) CreateVolume(_param0 context.Context, _param1 *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateVolume", _param0, _param1)
	ret0, _ := ret[0].(*CreateVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) CreateVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateVolume", arg0, arg1)
}

func (_m *MockControllerServer) DeleteVolume(_param0 context.Context, _param1 *DeleteVolumeRequest) (*DeleteVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteVolume", _param0, _param1)
	ret0, _ := ret[0].(*DeleteVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) DeleteVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteVolume", arg0, arg1)
}

func (_m *MockControllerServer) ControllerPublishVolume(_param0 context.Context, _param1 *ControllerPublishVolumeRequest) (*ControllerPublishVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "ControllerPublishVolume", _param0, _param1)
	ret0, _ := ret[0].(*ControllerPublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ControllerPublishVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerPublishVolume", arg0, arg1)
}

func (_m *MockControllerServer) ControllerUnpublishVolume(_param0 context.Context, _param1 *ControllerUnpublishVolumeRequest) (*ControllerUnpublishVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "ControllerUnpublishVolume", _param0, _param1)
	ret0, _ := ret[0].(*ControllerUnpublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ControllerUnpublishVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerUnpublishVolume", arg0, arg1)
}

func (_m *MockControllerServer) ValidateVolumeCapabilities(_param0 context.Context, _param1 *ValidateVolumeCapabilitiesRequest) (*ValidateVolumeCapabilitiesResponse, error) {
	ret := _m.ctrl.Call(_m, "ValidateVolumeCapabilities", _param0, _param1)
	ret0, _ := ret[0].(*ValidateVolumeCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ValidateVolumeCapabilities(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateVolumeCapabilities", arg0, arg1)
}

func (_m *MockControllerServer) ListVolumes(_param0 context.Context, _param1 *ListVolumesRequest) (*ListVolumesResponse, error) {
	ret := _m.ctrl.Call(_m, "ListVolumes", _param0, _param1)
	ret0, _ := ret[0].(*ListVolumesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ListVolumes(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListVolumes", arg0, arg1)
}

func (_m *MockControllerServer) GetCapacity(_param0 context.Context, _param1 *GetCapacityRequest) (*GetCapacityResponse, error) {
	ret := _m.ctrl.Call(_m, "GetCapacity", _param0, _param1)
	ret0, _ := ret[0].(*GetCapacityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) GetCapacity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCapacity", arg0, arg1)
}

func (_m *MockControllerServer) ControllerProbe(_param0 context.Context, _param1 *ControllerProbeRequest) (*ControllerProbeResponse, error) {
	ret := _m.ctrl.Call(_m, "ControllerProbe", _param0, _param1)
	ret0, _ := ret[0].(*ControllerProbeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ControllerProbe(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerProbe", arg0, arg1)
}

func (_m *MockControllerServer) ControllerGetCapabilities(_param0 context.Context, _param1 *ControllerGetCapabilitiesRequest) (*ControllerGetCapabilitiesResponse, error) {
	ret := _m.ctrl.Call(_m, "ControllerGetCapabilities", _param0, _param1)
	ret0, _ := ret[0].(*ControllerGetCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerServerRecorder) ControllerGetCapabilities(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ControllerGetCapabilities", arg0, arg1)
}

// Mock of NodeClient interface
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *_MockNodeClientRecorder
}

// Recorder for MockNodeClient (not exported)
type _MockNodeClientRecorder struct {
	mock *MockNodeClient
}

func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &_MockNodeClientRecorder{mock}
	return mock
}

func (_m *MockNodeClient) EXPECT() *_MockNodeClientRecorder {
	return _m.recorder
}

func (_m *MockNodeClient) NodePublishVolume(ctx context.Context, in *NodePublishVolumeRequest, opts ...grpc.CallOption) (*NodePublishVolumeResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NodePublishVolume", _s...)
	ret0, _ := ret[0].(*NodePublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeClientRecorder) NodePublishVolume(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodePublishVolume", _s...)
}

func (_m *MockNodeClient) NodeUnpublishVolume(ctx context.Context, in *NodeUnpublishVolumeRequest, opts ...grpc.CallOption) (*NodeUnpublishVolumeResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NodeUnpublishVolume", _s...)
	ret0, _ := ret[0].(*NodeUnpublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeClientRecorder) NodeUnpublishVolume(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeUnpublishVolume", _s...)
}

func (_m *MockNodeClient) GetNodeID(ctx context.Context, in *GetNodeIDRequest, opts ...grpc.CallOption) (*GetNodeIDResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetNodeID", _s...)
	ret0, _ := ret[0].(*GetNodeIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeClientRecorder) GetNodeID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNodeID", _s...)
}

func (_m *MockNodeClient) NodeProbe(ctx context.Context, in *NodeProbeRequest, opts ...grpc.CallOption) (*NodeProbeResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NodeProbe", _s...)
	ret0, _ := ret[0].(*NodeProbeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeClientRecorder) NodeProbe(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeProbe", _s...)
}

func (_m *MockNodeClient) NodeGetCapabilities(ctx context.Context, in *NodeGetCapabilitiesRequest, opts ...grpc.CallOption) (*NodeGetCapabilitiesResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NodeGetCapabilities", _s...)
	ret0, _ := ret[0].(*NodeGetCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeClientRecorder) NodeGetCapabilities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeGetCapabilities", _s...)
}

// Mock of NodeServer interface
type MockNodeServer struct {
	ctrl     *gomock.Controller
	recorder *_MockNodeServerRecorder
}

// Recorder for MockNodeServer (not exported)
type _MockNodeServerRecorder struct {
	mock *MockNodeServer
}

func NewMockNodeServer(ctrl *gomock.Controller) *MockNodeServer {
	mock := &MockNodeServer{ctrl: ctrl}
	mock.recorder = &_MockNodeServerRecorder{mock}
	return mock
}

func (_m *MockNodeServer) EXPECT() *_MockNodeServerRecorder {
	return _m.recorder
}

func (_m *MockNodeServer) NodePublishVolume(_param0 context.Context, _param1 *NodePublishVolumeRequest) (*NodePublishVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "NodePublishVolume", _param0, _param1)
	ret0, _ := ret[0].(*NodePublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodePublishVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodePublishVolume", arg0, arg1)
}

func (_m *MockNodeServer) NodeUnpublishVolume(_param0 context.Context, _param1 *NodeUnpublishVolumeRequest) (*NodeUnpublishVolumeResponse, error) {
	ret := _m.ctrl.Call(_m, "NodeUnpublishVolume", _param0, _param1)
	ret0, _ := ret[0].(*NodeUnpublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodeUnpublishVolume(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeUnpublishVolume", arg0, arg1)
}

func (_m *MockNodeServer) GetNodeID(_param0 context.Context, _param1 *GetNodeIDRequest) (*GetNodeIDResponse, error) {
	ret := _m.ctrl.Call(_m, "GetNodeID", _param0, _param1)
	ret0, _ := ret[0].(*GetNodeIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) GetNodeID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNodeID", arg0, arg1)
}

func (_m *MockNodeServer) NodeProbe(_param0 context.Context, _param1 *NodeProbeRequest) (*NodeProbeResponse, error) {
	ret := _m.ctrl.Call(_m, "NodeProbe", _param0, _param1)
	ret0, _ := ret[0].(*NodeProbeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodeProbe(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeProbe", arg0, arg1)
}

func (_m *MockNodeServer) NodeGetCapabilities(_param0 context.Context, _param1 *NodeGetCapabilitiesRequest) (*NodeGetCapabilitiesResponse, error) {
	ret := _m.ctrl.Call(_m, "NodeGetCapabilities", _param0, _param1)
	ret0, _ := ret[0].(*NodeGetCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNodeServerRecorder) NodeGetCapabilities(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeGetCapabilities", arg0, arg1)
}
