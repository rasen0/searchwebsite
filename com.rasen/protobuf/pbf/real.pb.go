// Code generated by protoc-gen-go. DO NOT EDIT.
// source: real.proto

package robpbf

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 用于服务器主动向客户端下发指令更新通知
type MissionInform struct {
	UpdateStamp          int64    `protobuf:"varint,1,opt,name=update_stamp,json=updateStamp,proto3" json:"update_stamp,omitempty"`
	MissionId            int32    `protobuf:"varint,2,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MissionInform) Reset()         { *m = MissionInform{} }
func (m *MissionInform) String() string { return proto.CompactTextString(m) }
func (*MissionInform) ProtoMessage()    {}
func (*MissionInform) Descriptor() ([]byte, []int) {
	return fileDescriptor_a785078f41c39600, []int{0}
}

func (m *MissionInform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissionInform.Unmarshal(m, b)
}
func (m *MissionInform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissionInform.Marshal(b, m, deterministic)
}
func (m *MissionInform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissionInform.Merge(m, src)
}
func (m *MissionInform) XXX_Size() int {
	return xxx_messageInfo_MissionInform.Size(m)
}
func (m *MissionInform) XXX_DiscardUnknown() {
	xxx_messageInfo_MissionInform.DiscardUnknown(m)
}

var xxx_messageInfo_MissionInform proto.InternalMessageInfo

func (m *MissionInform) GetUpdateStamp() int64 {
	if m != nil {
		return m.UpdateStamp
	}
	return 0
}

func (m *MissionInform) GetMissionId() int32 {
	if m != nil {
		return m.MissionId
	}
	return 0
}

func init() {
	proto.RegisterType((*MissionInform)(nil), "robpbf.MissionInform")
}

func init() { proto.RegisterFile("real.proto", fileDescriptor_a785078f41c39600) }

var fileDescriptor_a785078f41c39600 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4a, 0x4d, 0xcc,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0xca, 0x4f, 0x2a, 0x48, 0x4a, 0x53, 0x0a,
	0xe4, 0xe2, 0xf5, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xf3, 0xcc, 0x4b, 0xcb, 0x2f, 0xca, 0x15,
	0x52, 0xe4, 0xe2, 0x29, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x8d, 0x2f, 0x2e, 0x49, 0xcc, 0x2d, 0x90,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0xe2, 0x86, 0x88, 0x05, 0x83, 0x84, 0x84, 0x64, 0xb9, 0xb8,
	0x72, 0x21, 0x7a, 0xe2, 0x33, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x38, 0xa1, 0x22,
	0x9e, 0x29, 0x4e, 0x52, 0x5c, 0x7c, 0x79, 0x59, 0x45, 0xf9, 0x49, 0xf9, 0x25, 0x7a, 0x10, 0x4b,
	0x3c, 0x98, 0xa3, 0xa0, 0xd6, 0x25, 0xb1, 0x81, 0x6d, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x46, 0x3a, 0x9d, 0x45, 0x8b, 0x00, 0x00, 0x00,
}
