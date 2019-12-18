// Code generated by protoc-gen-go. DO NOT EDIT.
// source: const.proto

package tio_control_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type JobStatus int32

const (
	JobStatus_Ready        JobStatus = 0
	JobStatus_BuildSucc    JobStatus = 1
	JobStatus_BuildFailed  JobStatus = 2
	JobStatus_BuildIng     JobStatus = 3
	JobStatus_DeployIng    JobStatus = 4
	JobStatus_DeploySuc    JobStatus = 5
	JobStatus_DeployFailed JobStatus = 6
)

var JobStatus_name = map[int32]string{
	0: "Ready",
	1: "BuildSucc",
	2: "BuildFailed",
	3: "BuildIng",
	4: "DeployIng",
	5: "DeploySuc",
	6: "DeployFailed",
}

var JobStatus_value = map[string]int32{
	"Ready":        0,
	"BuildSucc":    1,
	"BuildFailed":  2,
	"BuildIng":     3,
	"DeployIng":    4,
	"DeploySuc":    5,
	"DeployFailed": 6,
}

func (x JobStatus) String() string {
	return proto.EnumName(JobStatus_name, int32(x))
}

func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

type CommonRespCode int32

const (
	CommonRespCode_RespSucc  CommonRespCode = 0
	CommonRespCode_RespFaild CommonRespCode = -1
)

var CommonRespCode_name = map[int32]string{
	0:  "RespSucc",
	-1: "RespFaild",
}

var CommonRespCode_value = map[string]int32{
	"RespSucc":  0,
	"RespFaild": -1,
}

func (x CommonRespCode) String() string {
	return proto.EnumName(CommonRespCode_name, int32(x))
}

func (CommonRespCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

type TioReply struct {
	Code                 CommonRespCode `protobuf:"varint,1,opt,name=code,proto3,enum=CommonRespCode" json:"code,omitempty"`
	Msg                  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TioReply) Reset()         { *m = TioReply{} }
func (m *TioReply) String() string { return proto.CompactTextString(m) }
func (*TioReply) ProtoMessage()    {}
func (*TioReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

func (m *TioReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioReply.Unmarshal(m, b)
}
func (m *TioReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioReply.Marshal(b, m, deterministic)
}
func (m *TioReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioReply.Merge(m, src)
}
func (m *TioReply) XXX_Size() int {
	return xxx_messageInfo_TioReply.Size(m)
}
func (m *TioReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TioReply.DiscardUnknown(m)
}

var xxx_messageInfo_TioReply proto.InternalMessageInfo

func (m *TioReply) GetCode() CommonRespCode {
	if m != nil {
		return m.Code
	}
	return CommonRespCode_RespSucc
}

func (m *TioReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TioUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioUserRequest) Reset()         { *m = TioUserRequest{} }
func (m *TioUserRequest) String() string { return proto.CompactTextString(m) }
func (*TioUserRequest) ProtoMessage()    {}
func (*TioUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

func (m *TioUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioUserRequest.Unmarshal(m, b)
}
func (m *TioUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioUserRequest.Marshal(b, m, deterministic)
}
func (m *TioUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioUserRequest.Merge(m, src)
}
func (m *TioUserRequest) XXX_Size() int {
	return xxx_messageInfo_TioUserRequest.Size(m)
}
func (m *TioUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TioUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TioUserRequest proto.InternalMessageInfo

func (m *TioUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TioUserRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type TioUserReply struct {
	Code                 CommonRespCode `protobuf:"varint,1,opt,name=Code,proto3,enum=CommonRespCode" json:"Code,omitempty"`
	Token                *TioToken      `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	User                 *TioUserInfo   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TioUserReply) Reset()         { *m = TioUserReply{} }
func (m *TioUserReply) String() string { return proto.CompactTextString(m) }
func (*TioUserReply) ProtoMessage()    {}
func (*TioUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{2}
}

func (m *TioUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioUserReply.Unmarshal(m, b)
}
func (m *TioUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioUserReply.Marshal(b, m, deterministic)
}
func (m *TioUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioUserReply.Merge(m, src)
}
func (m *TioUserReply) XXX_Size() int {
	return xxx_messageInfo_TioUserReply.Size(m)
}
func (m *TioUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TioUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_TioUserReply proto.InternalMessageInfo

func (m *TioUserReply) GetCode() CommonRespCode {
	if m != nil {
		return m.Code
	}
	return CommonRespCode_RespSucc
}

func (m *TioUserReply) GetToken() *TioToken {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *TioUserReply) GetUser() *TioUserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type TioUserInfo struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioUserInfo) Reset()         { *m = TioUserInfo{} }
func (m *TioUserInfo) String() string { return proto.CompactTextString(m) }
func (*TioUserInfo) ProtoMessage()    {}
func (*TioUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{3}
}

func (m *TioUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioUserInfo.Unmarshal(m, b)
}
func (m *TioUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioUserInfo.Marshal(b, m, deterministic)
}
func (m *TioUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioUserInfo.Merge(m, src)
}
func (m *TioUserInfo) XXX_Size() int {
	return xxx_messageInfo_TioUserInfo.Size(m)
}
func (m *TioUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TioUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TioUserInfo proto.InternalMessageInfo

func (m *TioUserInfo) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type TioToken struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
	Bucket               string   `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	CallBackUrl          string   `protobuf:"bytes,4,opt,name=callBackUrl,proto3" json:"callBackUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioToken) Reset()         { *m = TioToken{} }
func (m *TioToken) String() string { return proto.CompactTextString(m) }
func (*TioToken) ProtoMessage()    {}
func (*TioToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{4}
}

func (m *TioToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioToken.Unmarshal(m, b)
}
func (m *TioToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioToken.Marshal(b, m, deterministic)
}
func (m *TioToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioToken.Merge(m, src)
}
func (m *TioToken) XXX_Size() int {
	return xxx_messageInfo_TioToken.Size(m)
}
func (m *TioToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TioToken.DiscardUnknown(m)
}

var xxx_messageInfo_TioToken proto.InternalMessageInfo

func (m *TioToken) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *TioToken) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *TioToken) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *TioToken) GetCallBackUrl() string {
	if m != nil {
		return m.CallBackUrl
	}
	return ""
}

type TioAgentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioAgentRequest) Reset()         { *m = TioAgentRequest{} }
func (m *TioAgentRequest) String() string { return proto.CompactTextString(m) }
func (*TioAgentRequest) ProtoMessage()    {}
func (*TioAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{5}
}

func (m *TioAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioAgentRequest.Unmarshal(m, b)
}
func (m *TioAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioAgentRequest.Marshal(b, m, deterministic)
}
func (m *TioAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioAgentRequest.Merge(m, src)
}
func (m *TioAgentRequest) XXX_Size() int {
	return xxx_messageInfo_TioAgentRequest.Size(m)
}
func (m *TioAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TioAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TioAgentRequest proto.InternalMessageInfo

func (m *TioAgentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TioAgentReply struct {
	Code                 CommonRespCode `protobuf:"varint,1,opt,name=Code,proto3,enum=CommonRespCode" json:"Code,omitempty"`
	Address              string         `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TioAgentReply) Reset()         { *m = TioAgentReply{} }
func (m *TioAgentReply) String() string { return proto.CompactTextString(m) }
func (*TioAgentReply) ProtoMessage()    {}
func (*TioAgentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{6}
}

func (m *TioAgentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioAgentReply.Unmarshal(m, b)
}
func (m *TioAgentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioAgentReply.Marshal(b, m, deterministic)
}
func (m *TioAgentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioAgentReply.Merge(m, src)
}
func (m *TioAgentReply) XXX_Size() int {
	return xxx_messageInfo_TioAgentReply.Size(m)
}
func (m *TioAgentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TioAgentReply.DiscardUnknown(m)
}

var xxx_messageInfo_TioAgentReply proto.InternalMessageInfo

func (m *TioAgentReply) GetCode() CommonRespCode {
	if m != nil {
		return m.Code
	}
	return CommonRespCode_RespSucc
}

func (m *TioAgentReply) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type TioLogRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Flowing              bool     `protobuf:"varint,2,opt,name=flowing,proto3" json:"flowing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioLogRequest) Reset()         { *m = TioLogRequest{} }
func (m *TioLogRequest) String() string { return proto.CompactTextString(m) }
func (*TioLogRequest) ProtoMessage()    {}
func (*TioLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{7}
}

func (m *TioLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioLogRequest.Unmarshal(m, b)
}
func (m *TioLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioLogRequest.Marshal(b, m, deterministic)
}
func (m *TioLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioLogRequest.Merge(m, src)
}
func (m *TioLogRequest) XXX_Size() int {
	return xxx_messageInfo_TioLogRequest.Size(m)
}
func (m *TioLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TioLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TioLogRequest proto.InternalMessageInfo

func (m *TioLogRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TioLogRequest) GetFlowing() bool {
	if m != nil {
		return m.Flowing
	}
	return false
}

type TioLogReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TioLogReply) Reset()         { *m = TioLogReply{} }
func (m *TioLogReply) String() string { return proto.CompactTextString(m) }
func (*TioLogReply) ProtoMessage()    {}
func (*TioLogReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{8}
}

func (m *TioLogReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TioLogReply.Unmarshal(m, b)
}
func (m *TioLogReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TioLogReply.Marshal(b, m, deterministic)
}
func (m *TioLogReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TioLogReply.Merge(m, src)
}
func (m *TioLogReply) XXX_Size() int {
	return xxx_messageInfo_TioLogReply.Size(m)
}
func (m *TioLogReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TioLogReply.DiscardUnknown(m)
}

var xxx_messageInfo_TioLogReply proto.InternalMessageInfo

func (m *TioLogReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("JobStatus", JobStatus_name, JobStatus_value)
	proto.RegisterEnum("CommonRespCode", CommonRespCode_name, CommonRespCode_value)
	proto.RegisterType((*TioReply)(nil), "TioReply")
	proto.RegisterType((*TioUserRequest)(nil), "TioUserRequest")
	proto.RegisterType((*TioUserReply)(nil), "TioUserReply")
	proto.RegisterType((*TioUserInfo)(nil), "TioUserInfo")
	proto.RegisterType((*TioToken)(nil), "TioToken")
	proto.RegisterType((*TioAgentRequest)(nil), "TioAgentRequest")
	proto.RegisterType((*TioAgentReply)(nil), "TioAgentReply")
	proto.RegisterType((*TioLogRequest)(nil), "TioLogRequest")
	proto.RegisterType((*TioLogReply)(nil), "TioLogReply")
}

func init() { proto.RegisterFile("const.proto", fileDescriptor_5adb9555099c2688) }

var fileDescriptor_5adb9555099c2688 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6b, 0xdb, 0x4e,
	0x10, 0x8d, 0xe2, 0x4f, 0x8d, 0x1d, 0x5b, 0xec, 0x21, 0xe8, 0xf0, 0x83, 0x18, 0xfd, 0x28, 0x0d,
	0x39, 0x08, 0x9a, 0x42, 0x4f, 0xed, 0x21, 0x76, 0x29, 0xa4, 0x2d, 0x3d, 0xac, 0x95, 0x4b, 0x6f,
	0xeb, 0xdd, 0x89, 0x10, 0x96, 0x34, 0xaa, 0x76, 0x95, 0xe0, 0x5b, 0xff, 0xf3, 0x96, 0x5d, 0x49,
	0x75, 0x7a, 0x31, 0xf5, 0xc1, 0xcc, 0x7b, 0x6f, 0x3e, 0x9f, 0x24, 0x98, 0x49, 0x2a, 0xb5, 0x89,
	0xab, 0x9a, 0x0c, 0x45, 0x77, 0x30, 0x4d, 0x32, 0xe2, 0x58, 0xe5, 0x07, 0xf6, 0x3f, 0x0c, 0x25,
	0x29, 0x0c, 0xbd, 0x95, 0x77, 0xbd, 0xb8, 0x5d, 0xc6, 0x1b, 0x2a, 0x0a, 0x2a, 0x39, 0xea, 0x6a,
	0x43, 0x0a, 0xb9, 0x13, 0x59, 0x00, 0x83, 0x42, 0xa7, 0xe1, 0xf9, 0xca, 0xbb, 0xf6, 0xb9, 0x0d,
	0xa3, 0xf7, 0xb0, 0x48, 0x32, 0x7a, 0xd0, 0x58, 0x73, 0xfc, 0xd1, 0xa0, 0x36, 0x8c, 0xc1, 0xb0,
	0x14, 0x45, 0xdb, 0xc8, 0xe7, 0x2e, 0x66, 0x97, 0x30, 0xae, 0x84, 0xd6, 0xcf, 0xaa, 0x2b, 0xed,
	0x50, 0xf4, 0x04, 0xf3, 0x3f, 0xd5, 0xdd, 0x12, 0x9b, 0x53, 0x4b, 0xd8, 0x7f, 0x76, 0x05, 0x23,
	0x43, 0x7b, 0x2c, 0x5d, 0xaf, 0xd9, 0xad, 0x1f, 0x27, 0x19, 0x25, 0x96, 0xe0, 0x2d, 0xcf, 0x56,
	0x30, 0x6c, 0x34, 0xd6, 0xe1, 0xc0, 0xe9, 0xf3, 0xb8, 0x1b, 0x71, 0x5f, 0x3e, 0x12, 0x77, 0x4a,
	0x74, 0x05, 0xb3, 0x17, 0xa4, 0x3d, 0xab, 0xc9, 0x94, 0x9b, 0x3a, 0xe2, 0x36, 0x8c, 0x7e, 0x7a,
	0xce, 0x1a, 0xd7, 0x96, 0xfd, 0x07, 0xbe, 0x90, 0x12, 0xb5, 0xfe, 0x82, 0x87, 0xee, 0xac, 0x23,
	0x61, 0x55, 0x8d, 0xb2, 0x46, 0x63, 0xd5, 0xf6, 0xbc, 0x23, 0x61, 0x2f, 0xdf, 0x35, 0x72, 0x8f,
	0xc6, 0x6d, 0xe3, 0xf3, 0x0e, 0xb1, 0x15, 0xcc, 0xa4, 0xc8, 0xf3, 0xb5, 0x90, 0xfb, 0x87, 0x3a,
	0x0f, 0x87, 0x4e, 0x7c, 0x49, 0x45, 0xaf, 0x60, 0x99, 0x64, 0x74, 0x97, 0x62, 0x69, 0x4e, 0x58,
	0x1b, 0x7d, 0x83, 0x8b, 0x63, 0xda, 0x3f, 0x7b, 0x18, 0xc2, 0x44, 0x28, 0x55, 0xa3, 0xd6, 0xdd,
	0xca, 0x3d, 0x8c, 0x3e, 0xb8, 0x7e, 0x5f, 0x29, 0x3d, 0xf5, 0x3c, 0x43, 0x98, 0x3c, 0xe6, 0xf4,
	0x9c, 0x95, 0xed, 0xbb, 0x30, 0xe5, 0x3d, 0x8c, 0x5e, 0x3b, 0x67, 0x5d, 0xb9, 0x5d, 0x26, 0x84,
	0x49, 0x81, 0x5a, 0x8b, 0xb4, 0xaf, 0xef, 0xe1, 0x8d, 0x01, 0xff, 0x33, 0xed, 0xb6, 0x46, 0x98,
	0x46, 0x33, 0x1f, 0x46, 0x1c, 0x85, 0x3a, 0x04, 0x67, 0xec, 0x02, 0xfc, 0x75, 0x93, 0xe5, 0x6a,
	0xdb, 0x48, 0x19, 0x78, 0x6c, 0x09, 0x33, 0x07, 0x3f, 0x89, 0x2c, 0x47, 0x15, 0x9c, 0xb3, 0x39,
	0x4c, 0x1d, 0x71, 0x5f, 0xa6, 0xc1, 0xc0, 0x66, 0x7f, 0xc4, 0x2a, 0xa7, 0x83, 0x85, 0xc3, 0x23,
	0xdc, 0x36, 0x32, 0x18, 0xb1, 0x00, 0xe6, 0x2d, 0xec, 0xaa, 0xc7, 0x37, 0xef, 0x60, 0xf1, 0xb7,
	0x1f, 0xb6, 0x9f, 0x8d, 0xdd, 0xb8, 0x33, 0x76, 0x09, 0xbe, 0x45, 0x36, 0x5f, 0x05, 0xbf, 0xfa,
	0x9f, 0xb7, 0x0e, 0xbe, 0x2f, 0x4c, 0x46, 0xb1, 0xa4, 0xd2, 0xd4, 0x94, 0xc7, 0x4f, 0x6f, 0x76,
	0x63, 0xf7, 0x09, 0xbd, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x2c, 0x81, 0xf8, 0x51, 0x03,
	0x00, 0x00,
}
