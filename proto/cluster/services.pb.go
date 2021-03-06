// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: proto/cluster/services.proto

package prysm_internal_cluster

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PrivateKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName      string `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	NumberOfKeys uint64 `protobuf:"varint,2,opt,name=number_of_keys,json=numberOfKeys,proto3" json:"number_of_keys,omitempty"`
}

func (x *PrivateKeyRequest) Reset() {
	*x = PrivateKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKeyRequest) ProtoMessage() {}

func (x *PrivateKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKeyRequest.ProtoReflect.Descriptor instead.
func (*PrivateKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_cluster_services_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateKeyRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *PrivateKeyRequest) GetNumberOfKeys() uint64 {
	if x != nil {
		return x.NumberOfKeys
	}
	return 0
}

type PrivateKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	PrivateKey  []byte       `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PrivateKeys *PrivateKeys `protobuf:"bytes,2,opt,name=private_keys,json=privateKeys,proto3" json:"private_keys,omitempty"`
}

func (x *PrivateKeyResponse) Reset() {
	*x = PrivateKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKeyResponse) ProtoMessage() {}

func (x *PrivateKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKeyResponse.ProtoReflect.Descriptor instead.
func (*PrivateKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_cluster_services_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Do not use.
func (x *PrivateKeyResponse) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *PrivateKeyResponse) GetPrivateKeys() *PrivateKeys {
	if x != nil {
		return x.PrivateKeys
	}
	return nil
}

type PrivateKeys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKeys [][]byte `protobuf:"bytes,1,rep,name=private_keys,json=privateKeys,proto3" json:"private_keys,omitempty"`
}

func (x *PrivateKeys) Reset() {
	*x = PrivateKeys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKeys) ProtoMessage() {}

func (x *PrivateKeys) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKeys.ProtoReflect.Descriptor instead.
func (*PrivateKeys) Descriptor() ([]byte, []int) {
	return file_proto_cluster_services_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateKeys) GetPrivateKeys() [][]byte {
	if x != nil {
		return x.PrivateKeys
	}
	return nil
}

var File_proto_cluster_services_proto protoreflect.FileDescriptor

var file_proto_cluster_services_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x70, 0x72, 0x79, 0x73, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x6f, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x81, 0x01, 0x0a,
	0x12, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x46, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x73, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73,
	0x22, 0x30, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x73, 0x32, 0x75, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x2e, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x70, 0x72, 0x79, 0x73, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_cluster_services_proto_rawDescOnce sync.Once
	file_proto_cluster_services_proto_rawDescData = file_proto_cluster_services_proto_rawDesc
)

func file_proto_cluster_services_proto_rawDescGZIP() []byte {
	file_proto_cluster_services_proto_rawDescOnce.Do(func() {
		file_proto_cluster_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cluster_services_proto_rawDescData)
	})
	return file_proto_cluster_services_proto_rawDescData
}

var file_proto_cluster_services_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_cluster_services_proto_goTypes = []interface{}{
	(*PrivateKeyRequest)(nil),  // 0: prysm.internal.cluster.PrivateKeyRequest
	(*PrivateKeyResponse)(nil), // 1: prysm.internal.cluster.PrivateKeyResponse
	(*PrivateKeys)(nil),        // 2: prysm.internal.cluster.PrivateKeys
}
var file_proto_cluster_services_proto_depIdxs = []int32{
	2, // 0: prysm.internal.cluster.PrivateKeyResponse.private_keys:type_name -> prysm.internal.cluster.PrivateKeys
	0, // 1: prysm.internal.cluster.PrivateKeyService.Request:input_type -> prysm.internal.cluster.PrivateKeyRequest
	1, // 2: prysm.internal.cluster.PrivateKeyService.Request:output_type -> prysm.internal.cluster.PrivateKeyResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cluster_services_proto_init() }
func file_proto_cluster_services_proto_init() {
	if File_proto_cluster_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cluster_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateKeyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cluster_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateKeyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cluster_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateKeys); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cluster_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cluster_services_proto_goTypes,
		DependencyIndexes: file_proto_cluster_services_proto_depIdxs,
		MessageInfos:      file_proto_cluster_services_proto_msgTypes,
	}.Build()
	File_proto_cluster_services_proto = out.File
	file_proto_cluster_services_proto_rawDesc = nil
	file_proto_cluster_services_proto_goTypes = nil
	file_proto_cluster_services_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PrivateKeyServiceClient is the client API for PrivateKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrivateKeyServiceClient interface {
	Request(ctx context.Context, in *PrivateKeyRequest, opts ...grpc.CallOption) (*PrivateKeyResponse, error)
}

type privateKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateKeyServiceClient(cc grpc.ClientConnInterface) PrivateKeyServiceClient {
	return &privateKeyServiceClient{cc}
}

func (c *privateKeyServiceClient) Request(ctx context.Context, in *PrivateKeyRequest, opts ...grpc.CallOption) (*PrivateKeyResponse, error) {
	out := new(PrivateKeyResponse)
	err := c.cc.Invoke(ctx, "/prysm.internal.cluster.PrivateKeyService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateKeyServiceServer is the server API for PrivateKeyService service.
type PrivateKeyServiceServer interface {
	Request(context.Context, *PrivateKeyRequest) (*PrivateKeyResponse, error)
}

// UnimplementedPrivateKeyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPrivateKeyServiceServer struct {
}

func (*UnimplementedPrivateKeyServiceServer) Request(context.Context, *PrivateKeyRequest) (*PrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}

func RegisterPrivateKeyServiceServer(s *grpc.Server, srv PrivateKeyServiceServer) {
	s.RegisterService(&_PrivateKeyService_serviceDesc, srv)
}

func _PrivateKeyService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateKeyServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prysm.internal.cluster.PrivateKeyService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateKeyServiceServer).Request(ctx, req.(*PrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrivateKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prysm.internal.cluster.PrivateKeyService",
	HandlerType: (*PrivateKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _PrivateKeyService_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cluster/services.proto",
}
