// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control.proto

package tio_control_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BuildStatus struct {
	User                 string    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image                string    `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Api                  string    `protobuf:"bytes,4,opt,name=api,proto3" json:"api,omitempty"`
	Rate                 int32     `protobuf:"varint,5,opt,name=rate,proto3" json:"rate,omitempty"`
	Raw                  string    `protobuf:"bytes,6,opt,name=raw,proto3" json:"raw,omitempty"`
	Status               JobStatus `protobuf:"varint,7,opt,name=status,proto3,enum=JobStatus" json:"status,omitempty"`
	Sid                  int32     `protobuf:"varint,8,opt,name=sid,proto3" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BuildStatus) Reset()         { *m = BuildStatus{} }
func (m *BuildStatus) String() string { return proto.CompactTextString(m) }
func (*BuildStatus) ProtoMessage()    {}
func (*BuildStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0}
}

func (m *BuildStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildStatus.Unmarshal(m, b)
}
func (m *BuildStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildStatus.Marshal(b, m, deterministic)
}
func (m *BuildStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildStatus.Merge(m, src)
}
func (m *BuildStatus) XXX_Size() int {
	return xxx_messageInfo_BuildStatus.Size(m)
}
func (m *BuildStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildStatus.DiscardUnknown(m)
}

var xxx_messageInfo_BuildStatus proto.InternalMessageInfo

func (m *BuildStatus) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *BuildStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuildStatus) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *BuildStatus) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *BuildStatus) GetRate() int32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *BuildStatus) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

func (m *BuildStatus) GetStatus() JobStatus {
	if m != nil {
		return m.Status
	}
	return JobStatus_BuildSucc
}

func (m *BuildStatus) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

type BuildReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildReply) Reset()         { *m = BuildReply{} }
func (m *BuildReply) String() string { return proto.CompactTextString(m) }
func (*BuildReply) ProtoMessage()    {}
func (*BuildReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{1}
}

func (m *BuildReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildReply.Unmarshal(m, b)
}
func (m *BuildReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildReply.Marshal(b, m, deterministic)
}
func (m *BuildReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildReply.Merge(m, src)
}
func (m *BuildReply) XXX_Size() int {
	return xxx_messageInfo_BuildReply.Size(m)
}
func (m *BuildReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildReply.DiscardUnknown(m)
}

var xxx_messageInfo_BuildReply proto.InternalMessageInfo

func (m *BuildReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BuildReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildStatus)(nil), "BuildStatus")
	proto.RegisterType((*BuildReply)(nil), "BuildReply")
}

func init() { proto.RegisterFile("control.proto", fileDescriptor_0c5120591600887d) }

var fileDescriptor_0c5120591600887d = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4f, 0xbb, 0x40,
	0x10, 0xc5, 0xff, 0xfc, 0x5b, 0x50, 0x07, 0x25, 0x75, 0xe3, 0x61, 0xd3, 0x13, 0xe1, 0xc4, 0x45,
	0xa2, 0xf8, 0x0d, 0xea, 0xc1, 0xc4, 0x23, 0x8d, 0x17, 0x6f, 0x5b, 0x98, 0x34, 0x9b, 0x00, 0x8b,
	0xbb, 0x4b, 0x1b, 0x3f, 0x9a, 0xdf, 0xce, 0xcc, 0x6c, 0x8d, 0x3d, 0x7a, 0xfb, 0xcd, 0x63, 0x86,
	0xf7, 0xf2, 0x16, 0x6e, 0x5a, 0x33, 0x7a, 0x6b, 0xfa, 0x6a, 0xb2, 0xc6, 0x9b, 0x75, 0xda, 0x9a,
	0xd1, 0xf9, 0x30, 0x14, 0x5f, 0x11, 0xa4, 0x9b, 0x59, 0xf7, 0xdd, 0xd6, 0x2b, 0x3f, 0x3b, 0x21,
	0x60, 0x39, 0x3b, 0xb4, 0x32, 0xca, 0xa3, 0xf2, 0xaa, 0x61, 0x26, 0x6d, 0x54, 0x03, 0xca, 0xff,
	0x41, 0x23, 0x16, 0x77, 0x10, 0xeb, 0x41, 0xed, 0x51, 0x2e, 0x58, 0x0c, 0x83, 0x58, 0xc1, 0x42,
	0x4d, 0x5a, 0x2e, 0x59, 0x23, 0xa4, 0x5b, 0xab, 0x3c, 0xca, 0x38, 0x8f, 0xca, 0xb8, 0x61, 0xa6,
	0x2d, 0xab, 0x8e, 0x32, 0x09, 0x5b, 0x56, 0x1d, 0x45, 0x01, 0x89, 0x63, 0x7f, 0x79, 0x91, 0x47,
	0x65, 0x56, 0x43, 0xf5, 0x6a, 0x76, 0x21, 0x51, 0x73, 0xfa, 0x42, 0x57, 0x4e, 0x77, 0xf2, 0x92,
	0x7f, 0x44, 0x58, 0xd4, 0x00, 0x1c, 0xbd, 0xc1, 0xa9, 0xff, 0x24, 0xa7, 0xd6, 0x74, 0xc8, 0xc9,
	0xe3, 0x86, 0x99, 0x6e, 0x06, 0xb7, 0x3f, 0x05, 0x27, 0xac, 0x3f, 0x20, 0x7b, 0x0e, 0x6d, 0x6c,
	0xd1, 0x1e, 0x74, 0x8b, 0xe2, 0x01, 0x6e, 0xdf, 0xa6, 0x4e, 0x79, 0x3c, 0xaf, 0xe1, 0xba, 0x3a,
	0x9b, 0xd6, 0x69, 0xf5, 0xeb, 0x53, 0xfc, 0x13, 0xf7, 0x90, 0xbd, 0xa0, 0xff, 0xeb, 0xfa, 0x66,
	0xf5, 0x9e, 0x79, 0x6d, 0xaa, 0x9f, 0x47, 0x38, 0x3c, 0xee, 0x12, 0xee, 0xfe, 0xe9, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x55, 0xc0, 0xbc, 0x81, 0x99, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControlServiceClient is the client API for ControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlServiceClient interface {
	UpdateBuildStatus(ctx context.Context, in *BuildStatus, opts ...grpc.CallOption) (*BuildReply, error)
	GetBuildStatus(ctx context.Context, in *BuildStatus, opts ...grpc.CallOption) (*BuildReply, error)
}

type controlServiceClient struct {
	cc *grpc.ClientConn
}

func NewControlServiceClient(cc *grpc.ClientConn) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) UpdateBuildStatus(ctx context.Context, in *BuildStatus, opts ...grpc.CallOption) (*BuildReply, error) {
	out := new(BuildReply)
	err := c.cc.Invoke(ctx, "/ControlService/UpdateBuildStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) GetBuildStatus(ctx context.Context, in *BuildStatus, opts ...grpc.CallOption) (*BuildReply, error) {
	out := new(BuildReply)
	err := c.cc.Invoke(ctx, "/ControlService/GetBuildStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServiceServer is the server API for ControlService service.
type ControlServiceServer interface {
	UpdateBuildStatus(context.Context, *BuildStatus) (*BuildReply, error)
	GetBuildStatus(context.Context, *BuildStatus) (*BuildReply, error)
}

// UnimplementedControlServiceServer can be embedded to have forward compatible implementations.
type UnimplementedControlServiceServer struct {
}

func (*UnimplementedControlServiceServer) UpdateBuildStatus(ctx context.Context, req *BuildStatus) (*BuildReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBuildStatus not implemented")
}
func (*UnimplementedControlServiceServer) GetBuildStatus(ctx context.Context, req *BuildStatus) (*BuildReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildStatus not implemented")
}

func RegisterControlServiceServer(s *grpc.Server, srv ControlServiceServer) {
	s.RegisterService(&_ControlService_serviceDesc, srv)
}

func _ControlService_UpdateBuildStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).UpdateBuildStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ControlService/UpdateBuildStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).UpdateBuildStatus(ctx, req.(*BuildStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_GetBuildStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).GetBuildStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ControlService/GetBuildStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).GetBuildStatus(ctx, req.(*BuildStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateBuildStatus",
			Handler:    _ControlService_UpdateBuildStatus_Handler,
		},
		{
			MethodName: "GetBuildStatus",
			Handler:    _ControlService_GetBuildStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}
