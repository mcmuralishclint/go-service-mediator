// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/mediators.proto

package mediators

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type MediationInput struct {
	Service              string          `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Version              string          `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Endpoint             string          `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Verb                 string          `protobuf:"bytes,4,opt,name=verb,proto3" json:"verb,omitempty"`
	RequestData          *_struct.Struct `protobuf:"bytes,5,opt,name=request_data,json=requestData,proto3" json:"request_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MediationInput) Reset()         { *m = MediationInput{} }
func (m *MediationInput) String() string { return proto.CompactTextString(m) }
func (*MediationInput) ProtoMessage()    {}
func (*MediationInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_788e2471c8e29459, []int{0}
}

func (m *MediationInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediationInput.Unmarshal(m, b)
}
func (m *MediationInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediationInput.Marshal(b, m, deterministic)
}
func (m *MediationInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediationInput.Merge(m, src)
}
func (m *MediationInput) XXX_Size() int {
	return xxx_messageInfo_MediationInput.Size(m)
}
func (m *MediationInput) XXX_DiscardUnknown() {
	xxx_messageInfo_MediationInput.DiscardUnknown(m)
}

var xxx_messageInfo_MediationInput proto.InternalMessageInfo

func (m *MediationInput) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *MediationInput) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *MediationInput) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *MediationInput) GetVerb() string {
	if m != nil {
		return m.Verb
	}
	return ""
}

func (m *MediationInput) GetRequestData() *_struct.Struct {
	if m != nil {
		return m.RequestData
	}
	return nil
}

type MediationOutput struct {
	Output               *_struct.Struct `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MediationOutput) Reset()         { *m = MediationOutput{} }
func (m *MediationOutput) String() string { return proto.CompactTextString(m) }
func (*MediationOutput) ProtoMessage()    {}
func (*MediationOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_788e2471c8e29459, []int{1}
}

func (m *MediationOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediationOutput.Unmarshal(m, b)
}
func (m *MediationOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediationOutput.Marshal(b, m, deterministic)
}
func (m *MediationOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediationOutput.Merge(m, src)
}
func (m *MediationOutput) XXX_Size() int {
	return xxx_messageInfo_MediationOutput.Size(m)
}
func (m *MediationOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_MediationOutput.DiscardUnknown(m)
}

var xxx_messageInfo_MediationOutput proto.InternalMessageInfo

func (m *MediationOutput) GetOutput() *_struct.Struct {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*MediationInput)(nil), "MediationInput")
	proto.RegisterType((*MediationOutput)(nil), "MediationOutput")
}

func init() { proto.RegisterFile("pb/mediators.proto", fileDescriptor_788e2471c8e29459) }

var fileDescriptor_788e2471c8e29459 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0x39, 0x8d, 0x89, 0xce, 0xa9, 0x91, 0x6d, 0x5c, 0x0e, 0x8b, 0x90, 0x2a, 0x85, 0xec,
	0x41, 0x2c, 0x04, 0xcb, 0x60, 0x63, 0x21, 0x42, 0xec, 0x6c, 0x64, 0x2f, 0x37, 0x86, 0x05, 0xdd,
	0x39, 0x77, 0x67, 0xef, 0x77, 0xf9, 0x13, 0x25, 0x73, 0x97, 0x04, 0x2d, 0xd2, 0xbd, 0xb7, 0x6f,
	0x76, 0xe6, 0xe3, 0x81, 0x6a, 0xaa, 0xf2, 0x0b, 0x6b, 0x67, 0x99, 0x42, 0x34, 0x4d, 0x20, 0xa6,
	0xe2, 0x66, 0x4d, 0xb4, 0xfe, 0xc4, 0x52, 0x5c, 0x95, 0x3e, 0xca, 0xc8, 0x21, 0xad, 0xb8, 0x4b,
	0xa7, 0x3f, 0x19, 0x5c, 0x3e, 0xcb, 0x0f, 0x47, 0xfe, 0xc9, 0x37, 0x89, 0x95, 0x86, 0x51, 0xc4,
	0xd0, 0xba, 0x15, 0xea, 0x6c, 0x92, 0xcd, 0xce, 0x96, 0x5b, 0xbb, 0x49, 0x5a, 0x0c, 0xd1, 0x91,
	0xd7, 0x47, 0x5d, 0xd2, 0x5b, 0x55, 0xc0, 0x29, 0xfa, 0xba, 0x21, 0xe7, 0x59, 0x1f, 0x4b, 0xb4,
	0xf3, 0x4a, 0xc1, 0xa0, 0xc5, 0x50, 0xe9, 0x81, 0xbc, 0x8b, 0x56, 0x0f, 0x70, 0x1e, 0xf0, 0x3b,
	0x61, 0xe4, 0xf7, 0xda, 0xb2, 0xd5, 0x27, 0x93, 0x6c, 0x96, 0xcf, 0xaf, 0x4d, 0xc7, 0x6a, 0xb6,
	0xac, 0xe6, 0x55, 0x58, 0x97, 0x79, 0x3f, 0xfc, 0x68, 0xd9, 0x4e, 0x17, 0x30, 0xde, 0x11, 0xbf,
	0x24, 0xde, 0x20, 0x97, 0x30, 0x24, 0x51, 0x42, 0x7c, 0x60, 0x51, 0x3f, 0x36, 0xbf, 0x87, 0x51,
	0xb7, 0x03, 0xd5, 0xed, 0x5e, 0x8e, 0xcd, 0xdf, 0x2a, 0x8a, 0x2b, 0xf3, 0xef, 0xd2, 0xe2, 0xe2,
	0x2d, 0x37, 0xfb, 0x8a, 0xab, 0xa1, 0x1c, 0xb8, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xf5,
	0x2d, 0xd0, 0x79, 0x01, 0x00, 0x00,
}
