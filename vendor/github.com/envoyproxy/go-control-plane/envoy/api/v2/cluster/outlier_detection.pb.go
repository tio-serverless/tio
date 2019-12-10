// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package envoy_api_v2_cluster

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type OutlierDetection struct {
	Consecutive_5Xx                        *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx,proto3" json:"consecutive_5xx,omitempty"`
	Interval                               *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	BaseEjectionTime                       *duration.Duration    `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	MaxEjectionPercent                     *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	EnforcingConsecutive_5Xx               *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
	EnforcingSuccessRate                   *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate,proto3" json:"enforcing_success_rate,omitempty"`
	SuccessRateMinimumHosts                *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts,proto3" json:"success_rate_minimum_hosts,omitempty"`
	SuccessRateRequestVolume               *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume,proto3" json:"success_rate_request_volume,omitempty"`
	SuccessRateStdevFactor                 *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor,proto3" json:"success_rate_stdev_factor,omitempty"`
	ConsecutiveGatewayFailure              *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure,proto3" json:"consecutive_gateway_failure,omitempty"`
	EnforcingConsecutiveGatewayFailure     *wrappers.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure,proto3" json:"enforcing_consecutive_gateway_failure,omitempty"`
	SplitExternalLocalOriginErrors         bool                  `protobuf:"varint,12,opt,name=split_external_local_origin_errors,json=splitExternalLocalOriginErrors,proto3" json:"split_external_local_origin_errors,omitempty"`
	ConsecutiveLocalOriginFailure          *wrappers.UInt32Value `protobuf:"bytes,13,opt,name=consecutive_local_origin_failure,json=consecutiveLocalOriginFailure,proto3" json:"consecutive_local_origin_failure,omitempty"`
	EnforcingConsecutiveLocalOriginFailure *wrappers.UInt32Value `protobuf:"bytes,14,opt,name=enforcing_consecutive_local_origin_failure,json=enforcingConsecutiveLocalOriginFailure,proto3" json:"enforcing_consecutive_local_origin_failure,omitempty"`
	EnforcingLocalOriginSuccessRate        *wrappers.UInt32Value `protobuf:"bytes,15,opt,name=enforcing_local_origin_success_rate,json=enforcingLocalOriginSuccessRate,proto3" json:"enforcing_local_origin_success_rate,omitempty"`
	XXX_NoUnkeyedLiteral                   struct{}              `json:"-"`
	XXX_unrecognized                       []byte                `json:"-"`
	XXX_sizecache                          int32                 `json:"-"`
}

func (m *OutlierDetection) Reset()         { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()    {}
func (*OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_56cd87362a3f00c9, []int{0}
}

func (m *OutlierDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierDetection.Unmarshal(m, b)
}
func (m *OutlierDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierDetection.Marshal(b, m, deterministic)
}
func (m *OutlierDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetection.Merge(m, src)
}
func (m *OutlierDetection) XXX_Size() int {
	return xxx_messageInfo_OutlierDetection.Size(m)
}
func (m *OutlierDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetection.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetection proto.InternalMessageInfo

func (m *OutlierDetection) GetConsecutive_5Xx() *wrappers.UInt32Value {
	if m != nil {
		return m.Consecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection) GetBaseEjectionTime() *duration.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection) GetMaxEjectionPercent() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutive_5Xx() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingSuccessRate() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateMinimumHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateRequestVolume() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateRequestVolume
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateStdevFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateStdevFactor
	}
	return nil
}

func (m *OutlierDetection) GetConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.ConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetSplitExternalLocalOriginErrors() bool {
	if m != nil {
		return m.SplitExternalLocalOriginErrors
	}
	return false
}

func (m *OutlierDetection) GetConsecutiveLocalOriginFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.ConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveLocalOriginFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingLocalOriginSuccessRate() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingLocalOriginSuccessRate
	}
	return nil
}

func init() {
	proto.RegisterType((*OutlierDetection)(nil), "envoy.api.v2.cluster.OutlierDetection")
}

func init() {
	proto.RegisterFile("envoy/api/v2/cluster/outlier_detection.proto", fileDescriptor_56cd87362a3f00c9)
}

var fileDescriptor_56cd87362a3f00c9 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0x69, 0xd9, 0x9f, 0xe2, 0xc1, 0x36, 0x59, 0x63, 0x73, 0x37, 0x18, 0x53, 0x11, 0x68,
	0x9a, 0x50, 0x22, 0x75, 0xda, 0x35, 0xa2, 0x5b, 0xc7, 0x1f, 0x01, 0x1b, 0x1d, 0x6c, 0x42, 0x43,
	0xb2, 0xbc, 0xf4, 0x6b, 0x31, 0x72, 0xe2, 0x60, 0x3b, 0x59, 0x76, 0xc9, 0x73, 0xf0, 0x06, 0x7b,
	0x34, 0x1e, 0x61, 0x57, 0xa8, 0x49, 0xd3, 0x26, 0x5d, 0x10, 0xe9, 0x5d, 0xa5, 0xef, 0x9c, 0xdf,
	0x39, 0xfe, 0x9c, 0x34, 0xe8, 0x05, 0x78, 0xa1, 0xbc, 0xb2, 0x99, 0xcf, 0xed, 0xb0, 0x69, 0x3b,
	0x22, 0xd0, 0x06, 0x94, 0x2d, 0x03, 0x23, 0x38, 0x28, 0xda, 0x05, 0x03, 0x8e, 0xe1, 0xd2, 0xb3,
	0x7c, 0x25, 0x8d, 0xc4, 0x2b, 0xb1, 0xda, 0x62, 0x3e, 0xb7, 0xc2, 0xa6, 0x35, 0x54, 0xaf, 0x6f,
	0xf6, 0xa5, 0xec, 0x0b, 0xb0, 0x63, 0xcd, 0x45, 0xd0, 0xb3, 0xbb, 0x81, 0x62, 0x63, 0xd7, 0xed,
	0xf9, 0xa5, 0x62, 0xbe, 0x0f, 0x4a, 0x0f, 0xe7, 0x6b, 0x21, 0x13, 0xbc, 0xcb, 0x0c, 0xd8, 0xe9,
	0x8f, 0x64, 0xd0, 0xf8, 0xbd, 0x80, 0x96, 0x8f, 0x92, 0x2a, 0x07, 0x69, 0x13, 0xdc, 0x46, 0x4b,
	0x8e, 0xf4, 0x34, 0x38, 0x81, 0xe1, 0x21, 0xd0, 0xbd, 0x28, 0x22, 0x95, 0xad, 0xca, 0xf6, 0x42,
	0xf3, 0x91, 0x95, 0xe4, 0x58, 0x69, 0x8e, 0xf5, 0xe5, 0xad, 0x67, 0x76, 0x9b, 0xa7, 0x4c, 0x04,
	0xd0, 0x59, 0xcc, 0x98, 0xf6, 0xa2, 0x08, 0xbf, 0x44, 0x35, 0xee, 0x19, 0x50, 0x21, 0x13, 0xa4,
	0x1a, 0xfb, 0xeb, 0xb7, 0xfc, 0x07, 0xc3, 0x73, 0xb4, 0x6a, 0x37, 0xad, 0xd9, 0xeb, 0x4a, 0x75,
	0xe7, 0x4e, 0x67, 0x64, 0xc2, 0x9f, 0x10, 0xbe, 0x60, 0x1a, 0x28, 0xfc, 0x48, 0x8a, 0x51, 0xc3,
	0x5d, 0x20, 0x77, 0xcb, 0xa3, 0x96, 0x07, 0xf6, 0xf6, 0xd0, 0xfd, 0x99, 0xbb, 0x80, 0xcf, 0xd0,
	0x8a, 0xcb, 0xa2, 0x31, 0xd1, 0x07, 0xe5, 0x80, 0x67, 0xc8, 0xcc, 0xff, 0xcf, 0xd7, 0x9a, 0xbf,
	0x69, 0xcd, 0xec, 0x54, 0x49, 0xb7, 0x83, 0x5d, 0x16, 0xa5, 0xd4, 0xe3, 0x04, 0x80, 0x19, 0xaa,
	0x83, 0xd7, 0x93, 0xca, 0xe1, 0x5e, 0x9f, 0x4e, 0x6e, 0x6f, 0x76, 0x1a, 0xfa, 0xda, 0x88, 0xb3,
	0x9f, 0xdf, 0xe7, 0x39, 0x5a, 0x1d, 0x47, 0xe8, 0xc0, 0x71, 0x40, 0x6b, 0xaa, 0x98, 0x01, 0x32,
	0x37, 0x0d, 0x7f, 0x65, 0x04, 0x39, 0x49, 0x18, 0x1d, 0x66, 0x00, 0x7f, 0x45, 0xeb, 0x59, 0x24,
	0x75, 0xb9, 0xc7, 0xdd, 0xc0, 0xa5, 0xdf, 0xa5, 0x36, 0x9a, 0xcc, 0x97, 0xb8, 0xfe, 0x35, 0x3d,
	0xc6, 0x7d, 0x48, 0xdc, 0x6f, 0x06, 0x66, 0x7c, 0x8e, 0x36, 0x72, 0x68, 0x05, 0x3f, 0x03, 0xd0,
	0x86, 0x86, 0x52, 0x04, 0x2e, 0x90, 0x5a, 0x09, 0x36, 0xc9, 0xb0, 0x3b, 0x89, 0xfd, 0x34, 0x76,
	0xe3, 0x33, 0x54, 0xcf, 0xc1, 0xb5, 0xe9, 0x42, 0x48, 0x7b, 0xcc, 0x31, 0x52, 0x91, 0x7b, 0x25,
	0xd0, 0xab, 0x19, 0xf4, 0xc9, 0xc0, 0x7c, 0x18, 0x7b, 0xf1, 0x37, 0xb4, 0x91, 0xbd, 0xc6, 0x3e,
	0x33, 0x70, 0xc9, 0xae, 0x68, 0x8f, 0x71, 0x11, 0x28, 0x20, 0xa8, 0x04, 0xba, 0x9e, 0x01, 0xbc,
	0x4e, 0xfc, 0x87, 0x89, 0x1d, 0x47, 0xe8, 0x59, 0xf1, 0xe3, 0x32, 0x99, 0xb3, 0x30, 0xcd, 0xd5,
	0x36, 0x8a, 0x1e, 0x9d, 0x89, 0xe4, 0x77, 0xa8, 0xa1, 0x7d, 0xc1, 0x0d, 0x85, 0xc8, 0x80, 0xf2,
	0x98, 0xa0, 0x42, 0x3a, 0x4c, 0x50, 0xa9, 0x78, 0x9f, 0x7b, 0x14, 0x94, 0x92, 0x4a, 0x93, 0xfb,
	0x5b, 0x95, 0xed, 0x5a, 0x67, 0x33, 0x56, 0xb6, 0x87, 0xc2, 0xf7, 0x03, 0xdd, 0x51, 0x2c, 0x6b,
	0xc7, 0x2a, 0x0c, 0x68, 0x2b, 0xdb, 0x3d, 0x07, 0x4a, 0x0f, 0xf0, 0xa0, 0xc4, 0xa2, 0x1e, 0x67,
	0x28, 0x99, 0x94, 0xb4, 0xf2, 0xaf, 0x0a, 0xda, 0x29, 0xde, 0x56, 0x61, 0xe2, 0xe2, 0x34, 0x2b,
	0x7b, 0x5e, 0xb4, 0xb2, 0x82, 0x0e, 0x1a, 0x3d, 0x1d, 0x57, 0xc8, 0xc5, 0xe6, 0xde, 0xc4, 0xa5,
	0x69, 0xb2, 0x9f, 0x8c, 0x88, 0x99, 0xc0, 0xcc, 0x4b, 0xd9, 0x92, 0xa8, 0xc1, 0xa5, 0x15, 0x7f,
	0x11, 0x7c, 0x25, 0xa3, 0x2b, 0xab, 0xe8, 0xe3, 0xd0, 0x7a, 0x38, 0xf9, 0x07, 0x7e, 0x3c, 0x48,
	0x3d, 0xae, 0x5c, 0x57, 0x57, 0xdb, 0xb1, 0xfe, 0x95, 0xcf, 0xad, 0xd3, 0xa6, 0xb5, 0x9f, 0xe8,
	0x3f, 0x9e, 0xfc, 0xf9, 0xd7, 0xe0, 0x62, 0x2e, 0x2e, 0xbc, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x91, 0x91, 0x3b, 0x5a, 0xb4, 0x06, 0x00, 0x00,
}