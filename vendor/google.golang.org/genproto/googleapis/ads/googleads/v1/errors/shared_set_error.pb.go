// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/shared_set_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible shared set errors.
type SharedSetErrorEnum_SharedSetError int32

const (
	// Enum unspecified.
	SharedSetErrorEnum_UNSPECIFIED SharedSetErrorEnum_SharedSetError = 0
	// The received error code is not known in this version.
	SharedSetErrorEnum_UNKNOWN SharedSetErrorEnum_SharedSetError = 1
	// The customer cannot create this type of shared set.
	SharedSetErrorEnum_CUSTOMER_CANNOT_CREATE_SHARED_SET_OF_THIS_TYPE SharedSetErrorEnum_SharedSetError = 2
	// A shared set with this name already exists.
	SharedSetErrorEnum_DUPLICATE_NAME SharedSetErrorEnum_SharedSetError = 3
	// Removed shared sets cannot be mutated.
	SharedSetErrorEnum_SHARED_SET_REMOVED SharedSetErrorEnum_SharedSetError = 4
	// The shared set cannot be removed because it is in use.
	SharedSetErrorEnum_SHARED_SET_IN_USE SharedSetErrorEnum_SharedSetError = 5
)

var SharedSetErrorEnum_SharedSetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CUSTOMER_CANNOT_CREATE_SHARED_SET_OF_THIS_TYPE",
	3: "DUPLICATE_NAME",
	4: "SHARED_SET_REMOVED",
	5: "SHARED_SET_IN_USE",
}
var SharedSetErrorEnum_SharedSetError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CUSTOMER_CANNOT_CREATE_SHARED_SET_OF_THIS_TYPE": 2,
	"DUPLICATE_NAME":     3,
	"SHARED_SET_REMOVED": 4,
	"SHARED_SET_IN_USE":  5,
}

func (x SharedSetErrorEnum_SharedSetError) String() string {
	return proto.EnumName(SharedSetErrorEnum_SharedSetError_name, int32(x))
}
func (SharedSetErrorEnum_SharedSetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_shared_set_error_26e7a47f141f39e8, []int{0, 0}
}

// Container for enum describing possible shared set errors.
type SharedSetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharedSetErrorEnum) Reset()         { *m = SharedSetErrorEnum{} }
func (m *SharedSetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*SharedSetErrorEnum) ProtoMessage()    {}
func (*SharedSetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_set_error_26e7a47f141f39e8, []int{0}
}
func (m *SharedSetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSetErrorEnum.Unmarshal(m, b)
}
func (m *SharedSetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSetErrorEnum.Marshal(b, m, deterministic)
}
func (dst *SharedSetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSetErrorEnum.Merge(dst, src)
}
func (m *SharedSetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_SharedSetErrorEnum.Size(m)
}
func (m *SharedSetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SharedSetErrorEnum)(nil), "google.ads.googleads.v1.errors.SharedSetErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.SharedSetErrorEnum_SharedSetError", SharedSetErrorEnum_SharedSetError_name, SharedSetErrorEnum_SharedSetError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/shared_set_error.proto", fileDescriptor_shared_set_error_26e7a47f141f39e8)
}

var fileDescriptor_shared_set_error_26e7a47f141f39e8 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x8e, 0x9b, 0x30,
	0x18, 0xc5, 0x0b, 0xe9, 0x1f, 0xc9, 0x91, 0x52, 0xea, 0xaa, 0x5d, 0x54, 0x55, 0x16, 0x1c, 0xc0,
	0x88, 0x56, 0xdd, 0xb8, 0x2b, 0x07, 0x9c, 0x04, 0xb5, 0x31, 0x08, 0x03, 0x55, 0x2b, 0x24, 0x8b,
	0x16, 0x44, 0x23, 0x25, 0x38, 0xc2, 0x4c, 0x0e, 0x34, 0xd2, 0x6c, 0xe6, 0x0c, 0x73, 0x82, 0x39,
	0xca, 0x2c, 0xe6, 0x0c, 0x23, 0xf0, 0x24, 0x4a, 0x16, 0x33, 0x2b, 0x9e, 0x9e, 0x7e, 0xef, 0xe1,
	0xef, 0xfb, 0xc0, 0xb7, 0x5a, 0xca, 0x7a, 0x53, 0x39, 0x45, 0xa9, 0x1c, 0x2d, 0x7b, 0xb5, 0x77,
	0x9d, 0xaa, 0x6d, 0x65, 0xab, 0x1c, 0xf5, 0xbf, 0x68, 0xab, 0x52, 0xa8, 0xaa, 0x13, 0x83, 0x83,
	0x76, 0xad, 0xec, 0x24, 0x9c, 0x6a, 0x16, 0x15, 0xa5, 0x42, 0xc7, 0x18, 0xda, 0xbb, 0x48, 0xc7,
	0x3e, 0x7d, 0x3e, 0xd4, 0xee, 0xd6, 0x4e, 0xd1, 0x34, 0xb2, 0x2b, 0xba, 0xb5, 0x6c, 0x94, 0x4e,
	0xdb, 0x37, 0x06, 0x80, 0x7c, 0x28, 0xe6, 0x55, 0x47, 0xfb, 0x04, 0x6d, 0x2e, 0xb6, 0xf6, 0x95,
	0x01, 0x26, 0xe7, 0x36, 0x7c, 0x0b, 0xc6, 0x29, 0xe3, 0x11, 0xf5, 0x82, 0x79, 0x40, 0x7d, 0xeb,
	0x05, 0x1c, 0x83, 0x37, 0x29, 0xfb, 0xc1, 0xc2, 0x5f, 0xcc, 0x32, 0xe0, 0x17, 0x80, 0xbc, 0x94,
	0x27, 0xe1, 0x8a, 0xc6, 0xc2, 0x23, 0x8c, 0x85, 0x89, 0xf0, 0x62, 0x4a, 0x12, 0x2a, 0xf8, 0x92,
	0xc4, 0xd4, 0x17, 0x9c, 0x26, 0x22, 0x9c, 0x8b, 0x64, 0x19, 0x70, 0x91, 0xfc, 0x8e, 0xa8, 0x65,
	0x42, 0x08, 0x26, 0x7e, 0x1a, 0xfd, 0x0c, 0xbc, 0x1e, 0x63, 0x64, 0x45, 0xad, 0x11, 0xfc, 0x08,
	0xe0, 0x49, 0x20, 0xa6, 0xab, 0x30, 0xa3, 0xbe, 0xf5, 0x12, 0x7e, 0x00, 0xef, 0x4e, 0xfc, 0x80,
	0x89, 0x94, 0x53, 0xeb, 0xd5, 0xec, 0xde, 0x00, 0xf6, 0x3f, 0xb9, 0x45, 0xcf, 0xef, 0x60, 0xf6,
	0xfe, 0x7c, 0x96, 0xa8, 0x1f, 0x3d, 0x32, 0xfe, 0xf8, 0x8f, 0xb1, 0x5a, 0x6e, 0x8a, 0xa6, 0x46,
	0xb2, 0xad, 0x9d, 0xba, 0x6a, 0x86, 0xc5, 0x1c, 0x2e, 0xb0, 0x5b, 0xab, 0xa7, 0x0e, 0xf2, 0x5d,
	0x7f, 0x2e, 0xcd, 0xd1, 0x82, 0x90, 0x6b, 0x73, 0xba, 0xd0, 0x65, 0xa4, 0x54, 0x48, 0xcb, 0x5e,
	0x65, 0x2e, 0x1a, 0x7e, 0xa9, 0x6e, 0x0f, 0x40, 0x4e, 0x4a, 0x95, 0x1f, 0x81, 0x3c, 0x73, 0x73,
	0x0d, 0xdc, 0x99, 0xb6, 0x76, 0x31, 0x26, 0xa5, 0xc2, 0xf8, 0x88, 0x60, 0x9c, 0xb9, 0x18, 0x6b,
	0xe8, 0xef, 0xeb, 0xe1, 0x75, 0x5f, 0x1f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x19, 0xb4, 0xb0, 0x71,
	0x2d, 0x02, 0x00, 0x00,
}
