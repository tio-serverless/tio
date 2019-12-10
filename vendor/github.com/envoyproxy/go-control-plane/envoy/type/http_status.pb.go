// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/http_status.proto

package envoy_type

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type StatusCode int32

const (
	StatusCode_Empty                         StatusCode = 0
	StatusCode_Continue                      StatusCode = 100
	StatusCode_OK                            StatusCode = 200
	StatusCode_Created                       StatusCode = 201
	StatusCode_Accepted                      StatusCode = 202
	StatusCode_NonAuthoritativeInformation   StatusCode = 203
	StatusCode_NoContent                     StatusCode = 204
	StatusCode_ResetContent                  StatusCode = 205
	StatusCode_PartialContent                StatusCode = 206
	StatusCode_MultiStatus                   StatusCode = 207
	StatusCode_AlreadyReported               StatusCode = 208
	StatusCode_IMUsed                        StatusCode = 226
	StatusCode_MultipleChoices               StatusCode = 300
	StatusCode_MovedPermanently              StatusCode = 301
	StatusCode_Found                         StatusCode = 302
	StatusCode_SeeOther                      StatusCode = 303
	StatusCode_NotModified                   StatusCode = 304
	StatusCode_UseProxy                      StatusCode = 305
	StatusCode_TemporaryRedirect             StatusCode = 307
	StatusCode_PermanentRedirect             StatusCode = 308
	StatusCode_BadRequest                    StatusCode = 400
	StatusCode_Unauthorized                  StatusCode = 401
	StatusCode_PaymentRequired               StatusCode = 402
	StatusCode_Forbidden                     StatusCode = 403
	StatusCode_NotFound                      StatusCode = 404
	StatusCode_MethodNotAllowed              StatusCode = 405
	StatusCode_NotAcceptable                 StatusCode = 406
	StatusCode_ProxyAuthenticationRequired   StatusCode = 407
	StatusCode_RequestTimeout                StatusCode = 408
	StatusCode_Conflict                      StatusCode = 409
	StatusCode_Gone                          StatusCode = 410
	StatusCode_LengthRequired                StatusCode = 411
	StatusCode_PreconditionFailed            StatusCode = 412
	StatusCode_PayloadTooLarge               StatusCode = 413
	StatusCode_URITooLong                    StatusCode = 414
	StatusCode_UnsupportedMediaType          StatusCode = 415
	StatusCode_RangeNotSatisfiable           StatusCode = 416
	StatusCode_ExpectationFailed             StatusCode = 417
	StatusCode_MisdirectedRequest            StatusCode = 421
	StatusCode_UnprocessableEntity           StatusCode = 422
	StatusCode_Locked                        StatusCode = 423
	StatusCode_FailedDependency              StatusCode = 424
	StatusCode_UpgradeRequired               StatusCode = 426
	StatusCode_PreconditionRequired          StatusCode = 428
	StatusCode_TooManyRequests               StatusCode = 429
	StatusCode_RequestHeaderFieldsTooLarge   StatusCode = 431
	StatusCode_InternalServerError           StatusCode = 500
	StatusCode_NotImplemented                StatusCode = 501
	StatusCode_BadGateway                    StatusCode = 502
	StatusCode_ServiceUnavailable            StatusCode = 503
	StatusCode_GatewayTimeout                StatusCode = 504
	StatusCode_HTTPVersionNotSupported       StatusCode = 505
	StatusCode_VariantAlsoNegotiates         StatusCode = 506
	StatusCode_InsufficientStorage           StatusCode = 507
	StatusCode_LoopDetected                  StatusCode = 508
	StatusCode_NotExtended                   StatusCode = 510
	StatusCode_NetworkAuthenticationRequired StatusCode = 511
)

var StatusCode_name = map[int32]string{
	0:   "Empty",
	100: "Continue",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "NonAuthoritativeInformation",
	204: "NoContent",
	205: "ResetContent",
	206: "PartialContent",
	207: "MultiStatus",
	208: "AlreadyReported",
	226: "IMUsed",
	300: "MultipleChoices",
	301: "MovedPermanently",
	302: "Found",
	303: "SeeOther",
	304: "NotModified",
	305: "UseProxy",
	307: "TemporaryRedirect",
	308: "PermanentRedirect",
	400: "BadRequest",
	401: "Unauthorized",
	402: "PaymentRequired",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	406: "NotAcceptable",
	407: "ProxyAuthenticationRequired",
	408: "RequestTimeout",
	409: "Conflict",
	410: "Gone",
	411: "LengthRequired",
	412: "PreconditionFailed",
	413: "PayloadTooLarge",
	414: "URITooLong",
	415: "UnsupportedMediaType",
	416: "RangeNotSatisfiable",
	417: "ExpectationFailed",
	421: "MisdirectedRequest",
	422: "UnprocessableEntity",
	423: "Locked",
	424: "FailedDependency",
	426: "UpgradeRequired",
	428: "PreconditionRequired",
	429: "TooManyRequests",
	431: "RequestHeaderFieldsTooLarge",
	500: "InternalServerError",
	501: "NotImplemented",
	502: "BadGateway",
	503: "ServiceUnavailable",
	504: "GatewayTimeout",
	505: "HTTPVersionNotSupported",
	506: "VariantAlsoNegotiates",
	507: "InsufficientStorage",
	508: "LoopDetected",
	510: "NotExtended",
	511: "NetworkAuthenticationRequired",
}

var StatusCode_value = map[string]int32{
	"Empty":                         0,
	"Continue":                      100,
	"OK":                            200,
	"Created":                       201,
	"Accepted":                      202,
	"NonAuthoritativeInformation":   203,
	"NoContent":                     204,
	"ResetContent":                  205,
	"PartialContent":                206,
	"MultiStatus":                   207,
	"AlreadyReported":               208,
	"IMUsed":                        226,
	"MultipleChoices":               300,
	"MovedPermanently":              301,
	"Found":                         302,
	"SeeOther":                      303,
	"NotModified":                   304,
	"UseProxy":                      305,
	"TemporaryRedirect":             307,
	"PermanentRedirect":             308,
	"BadRequest":                    400,
	"Unauthorized":                  401,
	"PaymentRequired":               402,
	"Forbidden":                     403,
	"NotFound":                      404,
	"MethodNotAllowed":              405,
	"NotAcceptable":                 406,
	"ProxyAuthenticationRequired":   407,
	"RequestTimeout":                408,
	"Conflict":                      409,
	"Gone":                          410,
	"LengthRequired":                411,
	"PreconditionFailed":            412,
	"PayloadTooLarge":               413,
	"URITooLong":                    414,
	"UnsupportedMediaType":          415,
	"RangeNotSatisfiable":           416,
	"ExpectationFailed":             417,
	"MisdirectedRequest":            421,
	"UnprocessableEntity":           422,
	"Locked":                        423,
	"FailedDependency":              424,
	"UpgradeRequired":               426,
	"PreconditionRequired":          428,
	"TooManyRequests":               429,
	"RequestHeaderFieldsTooLarge":   431,
	"InternalServerError":           500,
	"NotImplemented":                501,
	"BadGateway":                    502,
	"ServiceUnavailable":            503,
	"GatewayTimeout":                504,
	"HTTPVersionNotSupported":       505,
	"VariantAlsoNegotiates":         506,
	"InsufficientStorage":           507,
	"LoopDetected":                  508,
	"NotExtended":                   510,
	"NetworkAuthenticationRequired": 511,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7544d7adacd3389b, []int{0}
}

type HttpStatus struct {
	Code                 StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.type.StatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HttpStatus) Reset()         { *m = HttpStatus{} }
func (m *HttpStatus) String() string { return proto.CompactTextString(m) }
func (*HttpStatus) ProtoMessage()    {}
func (*HttpStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7544d7adacd3389b, []int{0}
}

func (m *HttpStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpStatus.Unmarshal(m, b)
}
func (m *HttpStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpStatus.Marshal(b, m, deterministic)
}
func (m *HttpStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStatus.Merge(m, src)
}
func (m *HttpStatus) XXX_Size() int {
	return xxx_messageInfo_HttpStatus.Size(m)
}
func (m *HttpStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStatus proto.InternalMessageInfo

func (m *HttpStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Empty
}

func init() {
	proto.RegisterEnum("envoy.type.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*HttpStatus)(nil), "envoy.type.HttpStatus")
}

func init() { proto.RegisterFile("envoy/type/http_status.proto", fileDescriptor_7544d7adacd3389b) }

var fileDescriptor_7544d7adacd3389b = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x49, 0x6f, 0x5c, 0x45,
	0x10, 0xce, 0x9b, 0x8e, 0x93, 0xb8, 0xe3, 0xd8, 0x95, 0xce, 0x62, 0x13, 0x82, 0x64, 0xe5, 0x84,
	0x90, 0xb0, 0x25, 0xb8, 0x72, 0xb1, 0x1d, 0x3b, 0x36, 0x78, 0x26, 0xa3, 0x59, 0x72, 0x45, 0xed,
	0xd7, 0x35, 0x33, 0xad, 0xbc, 0xe9, 0x7a, 0xe9, 0x57, 0x33, 0xf6, 0xe3, 0xc8, 0x2f, 0x60, 0xdf,
	0xd7, 0x03, 0x8b, 0x50, 0x42, 0x40, 0xc0, 0x7f, 0x60, 0x87, 0xdf, 0xc0, 0x6f, 0x60, 0x0d, 0x08,
	0x50, 0xf7, 0x2c, 0xf6, 0x25, 0x27, 0xfb, 0x55, 0xd7, 0xf2, 0x2d, 0x35, 0x25, 0x2f, 0xa3, 0x1b,
	0x52, 0xb9, 0xca, 0x65, 0x8e, 0xab, 0x3d, 0xe6, 0xfc, 0xe9, 0x82, 0x35, 0x0f, 0x8a, 0x95, 0xdc,
	0x13, 0x93, 0x92, 0xf1, 0x75, 0x25, 0xbc, 0x5e, 0x5a, 0x1c, 0xea, 0xcc, 0x1a, 0xcd, 0xb8, 0x3a,
	0xf9, 0x67, 0x94, 0x74, 0xe5, 0x49, 0x29, 0xb7, 0x99, 0xf3, 0x66, 0x2c, 0x54, 0x4f, 0xc8, 0xe3,
	0x29, 0x19, 0x5c, 0x4a, 0x96, 0x93, 0x87, 0xe7, 0x1f, 0xbb, 0xb8, 0x72, 0xd8, 0x61, 0x65, 0x94,
	0xb1, 0x41, 0x06, 0xd7, 0xe1, 0xde, 0xfa, 0xcc, 0xb3, 0x49, 0x65, 0xf9, 0xd8, 0xe8, 0x2f, 0x24,
	0x8d, 0x58, 0xf5, 0xc8, 0x57, 0xb3, 0x52, 0x1e, 0xa6, 0xa9, 0x59, 0x39, 0xb3, 0xd9, 0xcf, 0xb9,
	0x84, 0x63, 0x6a, 0x4e, 0x9e, 0xda, 0x20, 0xc7, 0xd6, 0x0d, 0x10, 0x8c, 0x3a, 0x29, 0x2b, 0xd7,
	0x9f, 0x82, 0xaf, 0x13, 0x35, 0x27, 0x4f, 0x6e, 0x78, 0xd4, 0x8c, 0x06, 0xbe, 0x49, 0xd4, 0x19,
	0x79, 0x6a, 0x2d, 0x4d, 0x31, 0x0f, 0x9f, 0xdf, 0x26, 0x6a, 0x59, 0x3e, 0x58, 0x23, 0xb7, 0x36,
	0xe0, 0x1e, 0x79, 0xcb, 0x9a, 0xed, 0x10, 0x77, 0x5c, 0x87, 0x7c, 0x5f, 0xb3, 0x25, 0x07, 0xdf,
	0x25, 0x6a, 0x5e, 0xce, 0xd6, 0x28, 0xf4, 0x45, 0xc7, 0xf0, 0x7d, 0xa2, 0xce, 0xca, 0xb9, 0x06,
	0x16, 0xc8, 0x93, 0xd0, 0x0f, 0x89, 0x3a, 0x27, 0xe7, 0xeb, 0xda, 0xb3, 0xd5, 0xd9, 0x24, 0xf8,
	0x63, 0xa2, 0x40, 0x9e, 0xae, 0x0e, 0x32, 0xb6, 0x23, 0xac, 0xf0, 0x53, 0xa2, 0xce, 0xcb, 0x85,
	0xb5, 0xcc, 0xa3, 0x36, 0x65, 0x03, 0x73, 0xf2, 0x01, 0xc1, 0xcf, 0x89, 0x3a, 0x2d, 0x4f, 0xec,
	0x54, 0xdb, 0x05, 0x1a, 0xf8, 0x25, 0xa6, 0xc4, 0xa2, 0x3c, 0xc3, 0x8d, 0x1e, 0xd9, 0x14, 0x0b,
	0xb8, 0x5d, 0x51, 0x17, 0x24, 0x54, 0x69, 0x88, 0xa6, 0x8e, 0xbe, 0xaf, 0x1d, 0x3a, 0xce, 0x4a,
	0xb8, 0x53, 0x51, 0x52, 0xce, 0x6c, 0xd1, 0xc0, 0x19, 0xf8, 0xb4, 0x12, 0x68, 0x35, 0x11, 0xaf,
	0x73, 0x0f, 0x3d, 0xdc, 0xad, 0x84, 0xe1, 0x35, 0xe2, 0x2a, 0x19, 0xdb, 0xb1, 0x68, 0xe0, 0xb3,
	0x98, 0xd0, 0x2e, 0xb0, 0xee, 0xe9, 0xa0, 0x84, 0xcf, 0x2b, 0xea, 0xa2, 0x3c, 0xdb, 0xc2, 0x7e,
	0x4e, 0x5e, 0xfb, 0xb2, 0x81, 0xc6, 0x7a, 0x4c, 0x19, 0xbe, 0x88, 0xf1, 0xe9, 0x94, 0x69, 0xfc,
	0xcb, 0x8a, 0x5a, 0x90, 0x72, 0x5d, 0x9b, 0x06, 0xde, 0x1a, 0x60, 0xc1, 0xf0, 0x9c, 0x08, 0x32,
	0xb4, 0x9d, 0x1e, 0xe9, 0xf6, 0x0c, 0x1a, 0x78, 0x5e, 0x04, 0xf0, 0x75, 0x5d, 0xf6, 0x63, 0xe5,
	0xad, 0x81, 0xf5, 0x68, 0xe0, 0x05, 0x11, 0xf4, 0xdb, 0x22, 0xbf, 0x67, 0x8d, 0x41, 0x07, 0x2f,
	0x8a, 0x00, 0xa4, 0x46, 0x3c, 0x02, 0xfe, 0x92, 0x88, 0xdc, 0x90, 0x7b, 0x64, 0x6a, 0xc4, 0x6b,
	0x59, 0x46, 0xfb, 0x68, 0xe0, 0x65, 0xa1, 0x94, 0x3c, 0x13, 0x02, 0xd1, 0x29, 0xbd, 0x97, 0x21,
	0xbc, 0x22, 0x82, 0x57, 0x11, 0x7f, 0x70, 0x0b, 0x1d, 0xdb, 0x34, 0x7a, 0x34, 0x9d, 0xf5, 0xaa,
	0x08, 0x46, 0x8c, 0x21, 0xb6, 0x6c, 0x1f, 0x69, 0xc0, 0xf0, 0x5a, 0x1c, 0xb8, 0x41, 0xae, 0x93,
	0xd9, 0x94, 0xe1, 0x75, 0xa1, 0x66, 0xe5, 0xf1, 0x6b, 0xe4, 0x10, 0xde, 0x88, 0xe9, 0xbb, 0xe8,
	0xba, 0xdc, 0x9b, 0xf6, 0x78, 0x53, 0xa8, 0x45, 0xa9, 0xea, 0x1e, 0x53, 0x72, 0xc6, 0x86, 0xf6,
	0x5b, 0xda, 0x66, 0x68, 0xe0, 0xad, 0x09, 0xbd, 0x8c, 0xb4, 0x69, 0x11, 0xed, 0x6a, 0xdf, 0x45,
	0x78, 0x5b, 0x04, 0x61, 0xda, 0x8d, 0x9d, 0x10, 0x21, 0xd7, 0x85, 0x77, 0x84, 0x7a, 0x40, 0x9e,
	0x6f, 0xbb, 0x62, 0x90, 0x8f, 0x1c, 0xae, 0xa2, 0xb1, 0xba, 0x55, 0xe6, 0x08, 0xef, 0x0a, 0xb5,
	0x24, 0xcf, 0x35, 0xb4, 0xeb, 0x62, 0x8d, 0xb8, 0xa9, 0xd9, 0x16, 0x1d, 0x1b, 0xa9, 0xbd, 0x27,
	0x82, 0xec, 0x9b, 0x07, 0x39, 0xa6, 0xac, 0x8f, 0xcc, 0x7c, 0x3f, 0x82, 0xa9, 0xda, 0x62, 0x64,
	0x03, 0x4e, 0xe5, 0xff, 0x20, 0xb6, 0x6a, 0xbb, 0xdc, 0x53, 0x8a, 0x45, 0x11, 0x9a, 0x6c, 0x3a,
	0xb6, 0x5c, 0xc2, 0x87, 0x22, 0xec, 0xd3, 0x2e, 0xa5, 0x37, 0xd1, 0xc0, 0x47, 0x51, 0xdd, 0x51,
	0xb3, 0xab, 0x98, 0xa3, 0x33, 0xe8, 0xd2, 0x12, 0x3e, 0x8e, 0x54, 0xda, 0x79, 0xd7, 0x6b, 0x83,
	0x53, 0xe6, 0x9f, 0x44, 0xe4, 0x47, 0x99, 0x4f, 0x9f, 0x6e, 0xc7, 0x82, 0x16, 0x51, 0x55, 0xbb,
	0x72, 0x8c, 0xa1, 0x80, 0x3b, 0xd1, 0x90, 0xf1, 0xe7, 0x36, 0x6a, 0x83, 0x7e, 0xcb, 0x62, 0x66,
	0x8a, 0xa9, 0x3a, 0x77, 0x23, 0xcc, 0x1d, 0xc7, 0xe8, 0x9d, 0xce, 0x9a, 0xe8, 0x87, 0xe8, 0x37,
	0xbd, 0x27, 0x0f, 0xbf, 0x46, 0xed, 0x6b, 0xc4, 0x3b, 0xfd, 0x3c, 0xc3, 0xb0, 0x31, 0x68, 0xe0,
	0x37, 0x31, 0xde, 0xb2, 0x6b, 0x9a, 0x71, 0x5f, 0x97, 0xf0, 0x7b, 0xe4, 0x1f, 0xea, 0x6c, 0x8a,
	0x6d, 0xa7, 0x87, 0xda, 0x66, 0x51, 0xb0, 0x3f, 0x62, 0xf9, 0x38, 0x6d, 0xe2, 0xf4, 0x9f, 0x42,
	0x5d, 0x96, 0x8b, 0xdb, 0xad, 0x56, 0xfd, 0x06, 0xfa, 0xc2, 0x92, 0x0b, 0x2a, 0x4f, 0x6c, 0x80,
	0xbf, 0x84, 0xba, 0x24, 0x2f, 0xdc, 0xd0, 0xde, 0x6a, 0xc7, 0x6b, 0x59, 0x41, 0x35, 0xec, 0x12,
	0x5b, 0xcd, 0x58, 0xc0, 0xbd, 0x31, 0xce, 0x62, 0xd0, 0xe9, 0xd8, 0xd4, 0xa2, 0xe3, 0x26, 0x93,
	0xd7, 0x5d, 0x84, 0xbf, 0xe3, 0x9e, 0xef, 0x12, 0xe5, 0x57, 0x91, 0xa3, 0x05, 0xf0, 0x8f, 0x18,
	0xff, 0xb8, 0x36, 0x0f, 0x38, 0x28, 0x6a, 0xe0, 0x5f, 0xa1, 0xae, 0xc8, 0x87, 0x6a, 0xc8, 0xfb,
	0xe4, 0x6f, 0xde, 0x67, 0x37, 0xff, 0x13, 0xeb, 0x8f, 0xca, 0x25, 0x4b, 0xa3, 0x5b, 0x97, 0x87,
	0x2d, 0x3e, 0x72, 0xf6, 0xd6, 0x17, 0x0e, 0xaf, 0x63, 0x3d, 0x1c, 0xcc, 0x7a, 0xb2, 0x77, 0x22,
	0x5e, 0xce, 0xc7, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x85, 0x1b, 0x12, 0xf8, 0x7e, 0x05, 0x00,
	0x00,
}