// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/internal/proto/pathenum/path_enum.proto

package pathenum

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

type PathEnum int32

const (
	PathEnum_ABC PathEnum = 0
	PathEnum_DEF PathEnum = 1
)

var PathEnum_name = map[int32]string{
	0: "ABC",
	1: "DEF",
}

var PathEnum_value = map[string]int32{
	"ABC": 0,
	"DEF": 1,
}

func (x PathEnum) String() string {
	return proto.EnumName(PathEnum_name, int32(x))
}

func (PathEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_83f42520181de088, []int{0}
}

type MessagePathEnum_NestedPathEnum int32

const (
	MessagePathEnum_GHI MessagePathEnum_NestedPathEnum = 0
	MessagePathEnum_JKL MessagePathEnum_NestedPathEnum = 1
)

var MessagePathEnum_NestedPathEnum_name = map[int32]string{
	0: "GHI",
	1: "JKL",
}

var MessagePathEnum_NestedPathEnum_value = map[string]int32{
	"GHI": 0,
	"JKL": 1,
}

func (x MessagePathEnum_NestedPathEnum) String() string {
	return proto.EnumName(MessagePathEnum_NestedPathEnum_name, int32(x))
}

func (MessagePathEnum_NestedPathEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_83f42520181de088, []int{0, 0}
}

type MessagePathEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagePathEnum) Reset()         { *m = MessagePathEnum{} }
func (m *MessagePathEnum) String() string { return proto.CompactTextString(m) }
func (*MessagePathEnum) ProtoMessage()    {}
func (*MessagePathEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f42520181de088, []int{0}
}

func (m *MessagePathEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagePathEnum.Unmarshal(m, b)
}
func (m *MessagePathEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagePathEnum.Marshal(b, m, deterministic)
}
func (m *MessagePathEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagePathEnum.Merge(m, src)
}
func (m *MessagePathEnum) XXX_Size() int {
	return xxx_messageInfo_MessagePathEnum.Size(m)
}
func (m *MessagePathEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagePathEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MessagePathEnum proto.InternalMessageInfo

type MessageWithPathEnum struct {
	Value                PathEnum `protobuf:"varint,1,opt,name=value,proto3,enum=grpc.gateway.examples.internal.pathenum.PathEnum" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageWithPathEnum) Reset()         { *m = MessageWithPathEnum{} }
func (m *MessageWithPathEnum) String() string { return proto.CompactTextString(m) }
func (*MessageWithPathEnum) ProtoMessage()    {}
func (*MessageWithPathEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f42520181de088, []int{1}
}

func (m *MessageWithPathEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithPathEnum.Unmarshal(m, b)
}
func (m *MessageWithPathEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithPathEnum.Marshal(b, m, deterministic)
}
func (m *MessageWithPathEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithPathEnum.Merge(m, src)
}
func (m *MessageWithPathEnum) XXX_Size() int {
	return xxx_messageInfo_MessageWithPathEnum.Size(m)
}
func (m *MessageWithPathEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithPathEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithPathEnum proto.InternalMessageInfo

func (m *MessageWithPathEnum) GetValue() PathEnum {
	if m != nil {
		return m.Value
	}
	return PathEnum_ABC
}

type MessageWithNestedPathEnum struct {
	Value                MessagePathEnum_NestedPathEnum `protobuf:"varint,1,opt,name=value,proto3,enum=grpc.gateway.examples.internal.pathenum.MessagePathEnum_NestedPathEnum" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MessageWithNestedPathEnum) Reset()         { *m = MessageWithNestedPathEnum{} }
func (m *MessageWithNestedPathEnum) String() string { return proto.CompactTextString(m) }
func (*MessageWithNestedPathEnum) ProtoMessage()    {}
func (*MessageWithNestedPathEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f42520181de088, []int{2}
}

func (m *MessageWithNestedPathEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithNestedPathEnum.Unmarshal(m, b)
}
func (m *MessageWithNestedPathEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithNestedPathEnum.Marshal(b, m, deterministic)
}
func (m *MessageWithNestedPathEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithNestedPathEnum.Merge(m, src)
}
func (m *MessageWithNestedPathEnum) XXX_Size() int {
	return xxx_messageInfo_MessageWithNestedPathEnum.Size(m)
}
func (m *MessageWithNestedPathEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithNestedPathEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithNestedPathEnum proto.InternalMessageInfo

func (m *MessageWithNestedPathEnum) GetValue() MessagePathEnum_NestedPathEnum {
	if m != nil {
		return m.Value
	}
	return MessagePathEnum_GHI
}

func init() {
	proto.RegisterEnum("grpc.gateway.examples.internal.pathenum.PathEnum", PathEnum_name, PathEnum_value)
	proto.RegisterEnum("grpc.gateway.examples.internal.pathenum.MessagePathEnum_NestedPathEnum", MessagePathEnum_NestedPathEnum_name, MessagePathEnum_NestedPathEnum_value)
	proto.RegisterType((*MessagePathEnum)(nil), "grpc.gateway.examples.internal.pathenum.MessagePathEnum")
	proto.RegisterType((*MessageWithPathEnum)(nil), "grpc.gateway.examples.internal.pathenum.MessageWithPathEnum")
	proto.RegisterType((*MessageWithNestedPathEnum)(nil), "grpc.gateway.examples.internal.pathenum.MessageWithNestedPathEnum")
}

func init() {
	proto.RegisterFile("examples/internal/proto/pathenum/path_enum.proto", fileDescriptor_83f42520181de088)
}

var fileDescriptor_83f42520181de088 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x48, 0x2c, 0xc9, 0x48, 0xcd, 0x2b, 0xcd, 0x05, 0x33, 0xe2, 0x41,
	0x2c, 0x3d, 0xb0, 0x84, 0x90, 0x7a, 0x7a, 0x51, 0x41, 0xb2, 0x5e, 0x7a, 0x62, 0x49, 0x6a, 0x79,
	0x62, 0xa5, 0x1e, 0x4c, 0xbb, 0x1e, 0x4c, 0xbb, 0x1e, 0x4c, 0xa3, 0x92, 0x29, 0x17, 0xbf, 0x6f,
	0x6a, 0x71, 0x71, 0x62, 0x7a, 0x6a, 0x40, 0x62, 0x49, 0x86, 0x2b, 0x48, 0x48, 0x89, 0x8b, 0xcf,
	0x2f, 0xb5, 0xb8, 0x24, 0x35, 0x05, 0x26, 0x22, 0xc4, 0xce, 0xc5, 0xec, 0xee, 0xe1, 0x29, 0xc0,
	0x00, 0x62, 0x78, 0x79, 0xfb, 0x08, 0x30, 0x2a, 0xc5, 0x71, 0x09, 0x43, 0xb5, 0x85, 0x67, 0x96,
	0x64, 0xc0, 0x15, 0xba, 0x73, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x19, 0x19, 0xea, 0x11, 0xe9, 0x0c, 0x3d, 0x98, 0x09, 0x41, 0x10, 0xfd, 0x4a, 0x55, 0x5c,
	0x92, 0x48, 0xe6, 0xa3, 0x39, 0x27, 0x16, 0xd5, 0x16, 0x77, 0xa2, 0x6d, 0x41, 0xf3, 0xa9, 0x1e,
	0xaa, 0xb9, 0x50, 0xbb, 0xb5, 0x64, 0xb8, 0x38, 0x90, 0x7d, 0xee, 0xe8, 0xe4, 0x0c, 0xf1, 0xb9,
	0x8b, 0xab, 0x9b, 0x00, 0xa3, 0x93, 0x67, 0x94, 0x7b, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x3e, 0xc8, 0x66, 0xdd, 0xd4, 0xe4, 0xfc, 0xe2, 0xca, 0xe2, 0x92, 0x54, 0x28,
	0x17, 0xea, 0x10, 0x7d, 0x42, 0x91, 0x96, 0xc4, 0x06, 0xe6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xce, 0x83, 0xf1, 0x07, 0xdf, 0x01, 0x00, 0x00,
}