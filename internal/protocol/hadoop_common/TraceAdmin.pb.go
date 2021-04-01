// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TraceAdmin.proto

package hadoop_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListSpanReceiversRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListSpanReceiversRequestProto) Reset()                    { *m = ListSpanReceiversRequestProto{} }
func (m *ListSpanReceiversRequestProto) String() string            { return proto.CompactTextString(m) }
func (*ListSpanReceiversRequestProto) ProtoMessage()               {}
func (*ListSpanReceiversRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type SpanReceiverListInfo struct {
	Id               *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	ClassName        *string `protobuf:"bytes,2,req,name=className" json:"className,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SpanReceiverListInfo) Reset()                    { *m = SpanReceiverListInfo{} }
func (m *SpanReceiverListInfo) String() string            { return proto.CompactTextString(m) }
func (*SpanReceiverListInfo) ProtoMessage()               {}
func (*SpanReceiverListInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *SpanReceiverListInfo) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SpanReceiverListInfo) GetClassName() string {
	if m != nil && m.ClassName != nil {
		return *m.ClassName
	}
	return ""
}

type ListSpanReceiversResponseProto struct {
	Descriptions     []*SpanReceiverListInfo `protobuf:"bytes,1,rep,name=descriptions" json:"descriptions,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListSpanReceiversResponseProto) Reset()                    { *m = ListSpanReceiversResponseProto{} }
func (m *ListSpanReceiversResponseProto) String() string            { return proto.CompactTextString(m) }
func (*ListSpanReceiversResponseProto) ProtoMessage()               {}
func (*ListSpanReceiversResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ListSpanReceiversResponseProto) GetDescriptions() []*SpanReceiverListInfo {
	if m != nil {
		return m.Descriptions
	}
	return nil
}

type ConfigPair struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ConfigPair) Reset()                    { *m = ConfigPair{} }
func (m *ConfigPair) String() string            { return proto.CompactTextString(m) }
func (*ConfigPair) ProtoMessage()               {}
func (*ConfigPair) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *ConfigPair) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ConfigPair) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type AddSpanReceiverRequestProto struct {
	ClassName        *string       `protobuf:"bytes,1,req,name=className" json:"className,omitempty"`
	Config           []*ConfigPair `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *AddSpanReceiverRequestProto) Reset()                    { *m = AddSpanReceiverRequestProto{} }
func (m *AddSpanReceiverRequestProto) String() string            { return proto.CompactTextString(m) }
func (*AddSpanReceiverRequestProto) ProtoMessage()               {}
func (*AddSpanReceiverRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *AddSpanReceiverRequestProto) GetClassName() string {
	if m != nil && m.ClassName != nil {
		return *m.ClassName
	}
	return ""
}

func (m *AddSpanReceiverRequestProto) GetConfig() []*ConfigPair {
	if m != nil {
		return m.Config
	}
	return nil
}

type AddSpanReceiverResponseProto struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AddSpanReceiverResponseProto) Reset()                    { *m = AddSpanReceiverResponseProto{} }
func (m *AddSpanReceiverResponseProto) String() string            { return proto.CompactTextString(m) }
func (*AddSpanReceiverResponseProto) ProtoMessage()               {}
func (*AddSpanReceiverResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *AddSpanReceiverResponseProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type RemoveSpanReceiverRequestProto struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveSpanReceiverRequestProto) Reset()                    { *m = RemoveSpanReceiverRequestProto{} }
func (m *RemoveSpanReceiverRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveSpanReceiverRequestProto) ProtoMessage()               {}
func (*RemoveSpanReceiverRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *RemoveSpanReceiverRequestProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type RemoveSpanReceiverResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveSpanReceiverResponseProto) Reset()         { *m = RemoveSpanReceiverResponseProto{} }
func (m *RemoveSpanReceiverResponseProto) String() string { return proto.CompactTextString(m) }
func (*RemoveSpanReceiverResponseProto) ProtoMessage()    {}
func (*RemoveSpanReceiverResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{7}
}

func init() {
	proto.RegisterType((*ListSpanReceiversRequestProto)(nil), "hadoop.common.ListSpanReceiversRequestProto")
	proto.RegisterType((*SpanReceiverListInfo)(nil), "hadoop.common.SpanReceiverListInfo")
	proto.RegisterType((*ListSpanReceiversResponseProto)(nil), "hadoop.common.ListSpanReceiversResponseProto")
	proto.RegisterType((*ConfigPair)(nil), "hadoop.common.ConfigPair")
	proto.RegisterType((*AddSpanReceiverRequestProto)(nil), "hadoop.common.AddSpanReceiverRequestProto")
	proto.RegisterType((*AddSpanReceiverResponseProto)(nil), "hadoop.common.AddSpanReceiverResponseProto")
	proto.RegisterType((*RemoveSpanReceiverRequestProto)(nil), "hadoop.common.RemoveSpanReceiverRequestProto")
	proto.RegisterType((*RemoveSpanReceiverResponseProto)(nil), "hadoop.common.RemoveSpanReceiverResponseProto")
}

func init() { proto.RegisterFile("TraceAdmin.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xef, 0x4e, 0xea, 0x30,
	0x18, 0xc6, 0xb3, 0x91, 0x73, 0x12, 0xde, 0xc3, 0x39, 0x07, 0x1a, 0x3e, 0x0c, 0x44, 0xc0, 0xf9,
	0x85, 0xf8, 0x67, 0x2a, 0xf1, 0x06, 0x40, 0x13, 0x63, 0x62, 0x0c, 0x29, 0xde, 0x40, 0xd3, 0x15,
	0xa8, 0xb2, 0x76, 0xb6, 0x63, 0x89, 0x77, 0xe0, 0x65, 0x78, 0x75, 0x5e, 0x87, 0x61, 0x4a, 0xd6,
	0x15, 0x04, 0xbf, 0xed, 0xcf, 0xf3, 0xf6, 0xf9, 0x3d, 0xcf, 0x5b, 0xa8, 0x3e, 0x28, 0x42, 0xd9,
	0x20, 0x8c, 0xb8, 0x08, 0x62, 0x25, 0x13, 0x89, 0xfe, 0xce, 0x48, 0x28, 0x65, 0x1c, 0x50, 0x19,
	0x45, 0x52, 0xf8, 0x1d, 0xd8, 0xbf, 0xe3, 0x3a, 0x19, 0xc7, 0x44, 0x60, 0x46, 0x19, 0x4f, 0x99,
	0xd2, 0x98, 0x3d, 0x2f, 0x98, 0x4e, 0x46, 0x4b, 0xbd, 0x7f, 0x0d, 0x75, 0xf3, 0xe7, 0x52, 0x7c,
	0x2b, 0x26, 0x12, 0xfd, 0x03, 0x97, 0x87, 0x9e, 0xd3, 0x75, 0x7b, 0x25, 0xec, 0xf2, 0x10, 0xb5,
	0xa0, 0x4c, 0xe7, 0x44, 0xeb, 0x7b, 0x12, 0x31, 0xcf, 0xed, 0xba, 0xbd, 0x32, 0xce, 0x3f, 0xf8,
	0x1c, 0xda, 0x1b, 0x6c, 0x74, 0x2c, 0x85, 0x66, 0x99, 0x0f, 0xba, 0x81, 0x4a, 0xc8, 0x34, 0x55,
	0x3c, 0x4e, 0xb8, 0x14, 0xda, 0x73, 0xba, 0xa5, 0xde, 0x9f, 0xfe, 0x61, 0x50, 0xc0, 0x0d, 0x36,
	0xa1, 0xe0, 0xc2, 0xa0, 0x7f, 0x09, 0x70, 0x25, 0xc5, 0x84, 0x4f, 0x47, 0x84, 0x2b, 0x54, 0x85,
	0xd2, 0x13, 0x7b, 0xc9, 0x38, 0xcb, 0x78, 0xf9, 0x88, 0xea, 0xf0, 0x2b, 0x25, 0xf3, 0xc5, 0x0a,
	0xf2, 0xf3, 0xc5, 0x17, 0xb0, 0x37, 0x08, 0x43, 0xf3, 0x78, 0xb3, 0x85, 0x62, 0x3a, 0xc7, 0x4a,
	0x87, 0x2e, 0xe0, 0x37, 0xcd, 0x2c, 0x3d, 0x37, 0xa3, 0x6e, 0x58, 0xd4, 0x39, 0x0f, 0xfe, 0x12,
	0xfa, 0x01, 0xb4, 0xd6, 0xfc, 0xcc, 0x3a, 0xac, 0x7a, 0xfd, 0x73, 0x68, 0x63, 0x16, 0xc9, 0x94,
	0x7d, 0x8b, 0x68, 0x4f, 0x1c, 0x40, 0x67, 0xd3, 0x84, 0x61, 0xd2, 0x7f, 0x77, 0xa1, 0x96, 0x5f,
	0x90, 0x31, 0x53, 0x29, 0xa7, 0x0c, 0xc5, 0x50, 0x9b, 0xdb, 0xbb, 0x42, 0x27, 0x56, 0xa4, 0xad,
	0x97, 0xa6, 0x79, 0xba, 0x5b, 0x6d, 0x86, 0x7d, 0x84, 0xff, 0xa4, 0x58, 0x06, 0x3a, 0xb2, 0x4e,
	0xd8, 0xb2, 0x9c, 0xe6, 0xf1, 0x2e, 0xad, 0xe9, 0xa5, 0x01, 0xa9, 0xb5, 0x5a, 0x90, 0x0d, 0xbc,
	0xbd, 0xeb, 0x66, 0xf0, 0x03, 0xb9, 0x61, 0x3a, 0x3c, 0x83, 0x86, 0x54, 0xd3, 0x80, 0xc4, 0x84,
	0xce, 0xd8, 0x6a, 0x36, 0x51, 0x84, 0x72, 0x31, 0x1d, 0x56, 0xf2, 0x15, 0x8c, 0x86, 0xaf, 0x8e,
	0xf3, 0xe6, 0x38, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x09, 0x34, 0x4e, 0xb8, 0x03, 0x00,
	0x00,
}
