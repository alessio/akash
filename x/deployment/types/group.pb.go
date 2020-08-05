// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/group.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/ovrclk/akash/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Group_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	GroupStateInvalid      Group_State = 0
	GroupOpen              Group_State = 1
	GroupOrdered           Group_State = 2
	GroupMatched           Group_State = 3
	GroupInsufficientFunds Group_State = 4
	GroupClosed            Group_State = 5
)

var Group_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "ordered",
	3: "matched",
	4: "insufficient_funds",
	5: "closed",
}

var Group_State_value = map[string]int32{
	"invalid":            0,
	"open":               1,
	"ordered":            2,
	"matched":            3,
	"insufficient_funds": 4,
	"closed":             5,
}

func (x Group_State) String() string {
	return proto.EnumName(Group_State_name, int32(x))
}

func (Group_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{3, 0}
}

// MsgCloseGroup defines SDK message to close a single Group within a Deployment.
type MsgCloseGroup struct {
	ID GroupID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *MsgCloseGroup) Reset()         { *m = MsgCloseGroup{} }
func (m *MsgCloseGroup) String() string { return proto.CompactTextString(m) }
func (*MsgCloseGroup) ProtoMessage()    {}
func (*MsgCloseGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{0}
}
func (m *MsgCloseGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseGroup.Merge(m, src)
}
func (m *MsgCloseGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseGroup proto.InternalMessageInfo

func (m *MsgCloseGroup) GetID() GroupID {
	if m != nil {
		return m.ID
	}
	return GroupID{}
}

// GroupID stores owner, deployment sequence number and group sequence number
type GroupID struct {
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	DSeq  uint64                                        `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq  uint32                                        `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
}

func (m *GroupID) Reset()      { *m = GroupID{} }
func (*GroupID) ProtoMessage() {}
func (*GroupID) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{1}
}
func (m *GroupID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupID.Merge(m, src)
}
func (m *GroupID) XXX_Size() int {
	return m.Size()
}
func (m *GroupID) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupID.DiscardUnknown(m)
}

var xxx_messageInfo_GroupID proto.InternalMessageInfo

func (m *GroupID) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *GroupID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *GroupID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

type GroupSpec struct {
	Name         string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Requirements []types.Attribute `protobuf:"bytes,2,rep,name=requirements,proto3" json:"requirements" yaml:"requirements"`
	Resources    []Resource        `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources" yaml:"resources"`
}

func (m *GroupSpec) Reset()         { *m = GroupSpec{} }
func (m *GroupSpec) String() string { return proto.CompactTextString(m) }
func (*GroupSpec) ProtoMessage()    {}
func (*GroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{2}
}
func (m *GroupSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupSpec.Merge(m, src)
}
func (m *GroupSpec) XXX_Size() int {
	return m.Size()
}
func (m *GroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GroupSpec proto.InternalMessageInfo

type Group struct {
	GroupID   GroupID     `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"id" yaml:"id"`
	State     Group_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.Group_State" json:"state" yaml:"state"`
	GroupSpec GroupSpec   `protobuf:"bytes,3,opt,name=group_spec,json=groupSpec,proto3" json:"spec" yaml:"spec"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{3}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Group.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetGroupID() GroupID {
	if m != nil {
		return m.GroupID
	}
	return GroupID{}
}

func (m *Group) GetState() Group_State {
	if m != nil {
		return m.State
	}
	return GroupStateInvalid
}

func (m *Group) GetGroupSpec() GroupSpec {
	if m != nil {
		return m.GroupSpec
	}
	return GroupSpec{}
}

type Resource struct {
	Unit  types1.Unit `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit" yaml:"unit"`
	Count uint32      `protobuf:"varint,2,opt,name=count,proto3" json:"count" yaml:"count"`
	Price types.Coin  `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_940c80f4f1e11e54, []int{4}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetUnit() types1.Unit {
	if m != nil {
		return m.Unit
	}
	return types1.Unit{}
}

func (m *Resource) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Resource) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("akash.deployment.Group_State", Group_State_name, Group_State_value)
	proto.RegisterType((*MsgCloseGroup)(nil), "akash.deployment.MsgCloseGroup")
	proto.RegisterType((*GroupID)(nil), "akash.deployment.GroupID")
	proto.RegisterType((*GroupSpec)(nil), "akash.deployment.GroupSpec")
	proto.RegisterType((*Group)(nil), "akash.deployment.Group")
	proto.RegisterType((*Resource)(nil), "akash.deployment.Resource")
}

func init() { proto.RegisterFile("akash/deployment/group.proto", fileDescriptor_940c80f4f1e11e54) }

var fileDescriptor_940c80f4f1e11e54 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0xb6, 0x1d, 0x67, 0xdb, 0x4c, 0x52, 0xf0, 0xce, 0x2e, 0x6c, 0xd7, 0xa5, 0x19, 0x6b, 0x24,
	0xa4, 0x8a, 0xd5, 0xda, 0x52, 0xf6, 0x56, 0x71, 0xa9, 0xb7, 0x62, 0x29, 0xd2, 0x82, 0xe4, 0x0a,
	0x81, 0xb8, 0x94, 0xc4, 0x33, 0x75, 0x47, 0x4d, 0x3c, 0xae, 0xc7, 0x5e, 0xe8, 0x2f, 0x60, 0x95,
	0x13, 0x47, 0x2e, 0x91, 0x2a, 0xf1, 0x4f, 0x38, 0xed, 0x71, 0x8f, 0x70, 0x31, 0x28, 0xbd, 0xac,
	0x22, 0x4e, 0x39, 0x72, 0x42, 0x33, 0xe3, 0xc8, 0x09, 0xb0, 0x12, 0xa7, 0xe4, 0x7d, 0xef, 0x7b,
	0xdf, 0xb3, 0xdf, 0xfb, 0xfc, 0xc0, 0x07, 0xc3, 0xcb, 0xa1, 0xb8, 0x08, 0x08, 0xcd, 0xc6, 0xfc,
	0x7a, 0x42, 0xd3, 0x22, 0x48, 0x72, 0x5e, 0x66, 0x7e, 0x96, 0xf3, 0x82, 0x43, 0x47, 0x65, 0xfd,
	0x26, 0xeb, 0xde, 0x4f, 0x78, 0xc2, 0x55, 0x32, 0x90, 0xff, 0x34, 0xcf, 0xbd, 0xaf, 0x55, 0x72,
	0x2a, 0x78, 0x99, 0xc7, 0xb4, 0x46, 0xef, 0xc5, 0x5c, 0x4c, 0xb8, 0x08, 0xf4, 0x8f, 0x06, 0xf1,
	0x19, 0xd8, 0x79, 0x2e, 0x92, 0xa7, 0x63, 0x2e, 0xe8, 0x33, 0xd9, 0x09, 0x7e, 0x0a, 0x2c, 0x46,
	0x76, 0x4d, 0xcf, 0x3c, 0xe8, 0x0e, 0x1e, 0xfa, 0xff, 0x6c, 0xe8, 0x2b, 0xd2, 0xc9, 0x71, 0xb8,
	0xff, 0xaa, 0x42, 0xc6, 0xbc, 0x42, 0xd6, 0xc9, 0xf1, 0xa2, 0x42, 0x16, 0x23, 0xcb, 0x0a, 0x75,
	0xae, 0x87, 0x93, 0xf1, 0x21, 0x66, 0x04, 0x47, 0x16, 0x23, 0x87, 0xf6, 0x9b, 0x1b, 0x64, 0xe0,
	0xdf, 0x4d, 0xb0, 0x55, 0x17, 0xc1, 0x6f, 0x41, 0x9b, 0x7f, 0x97, 0xd2, 0x5c, 0xc9, 0xf7, 0xc2,
	0xcf, 0x16, 0x15, 0xd2, 0xc0, 0xb2, 0x42, 0x3d, 0x5d, 0xac, 0x42, 0xfc, 0x57, 0x85, 0x1e, 0x27,
	0xac, 0xb8, 0x28, 0x47, 0x7e, 0xcc, 0x27, 0xc1, 0xc6, 0x83, 0x3f, 0x16, 0xe4, 0x32, 0x28, 0xae,
	0x33, 0x2a, 0xfc, 0xa3, 0x38, 0x3e, 0x22, 0x24, 0xa7, 0x42, 0x44, 0x5a, 0x07, 0x3e, 0x01, 0x36,
	0x11, 0xf4, 0x6a, 0xd7, 0xf2, 0xcc, 0x03, 0x3b, 0x44, 0xf3, 0x0a, 0xd9, 0xc7, 0xa7, 0xf4, 0x6a,
	0x51, 0x21, 0x85, 0x2f, 0x2b, 0xd4, 0xd5, 0x7d, 0x64, 0x84, 0x23, 0x05, 0xca, 0xa2, 0x44, 0x16,
	0xb5, 0x3c, 0xf3, 0x60, 0x47, 0x17, 0x3d, 0xab, 0x8b, 0x92, 0x8d, 0xa2, 0x44, 0x17, 0xc9, 0x9f,
	0xc3, 0xed, 0x9f, 0x6e, 0x90, 0xf1, 0xe6, 0x06, 0x99, 0xf8, 0x07, 0x0b, 0x74, 0xd4, 0x1b, 0x9e,
	0x66, 0x34, 0x86, 0x8f, 0x80, 0x9d, 0x0e, 0x27, 0x54, 0xbd, 0x62, 0x27, 0x7c, 0x20, 0x45, 0x64,
	0xdc, 0x88, 0xc8, 0x08, 0x47, 0x0a, 0x84, 0x23, 0xd0, 0xcb, 0xe9, 0x55, 0xc9, 0x72, 0x2a, 0x87,
	0x2b, 0x76, 0x2d, 0xaf, 0x75, 0xd0, 0x1d, 0xdc, 0xf5, 0xeb, 0x15, 0x1d, 0x15, 0x45, 0xce, 0x46,
	0x65, 0x41, 0xc3, 0x47, 0x72, 0xdc, 0x8b, 0x0a, 0x6d, 0xd0, 0x97, 0x15, 0xba, 0xa7, 0x35, 0xd7,
	0x51, 0x1c, 0x6d, 0x90, 0xe0, 0x19, 0xe8, 0xac, 0x8c, 0x20, 0x76, 0x5b, 0xaa, 0x81, 0xfb, 0xef,
	0xbd, 0x46, 0x35, 0x25, 0xfc, 0xb0, 0xee, 0xd4, 0x14, 0x2d, 0x2b, 0xe4, 0xac, 0xda, 0xd4, 0x10,
	0x8e, 0x9a, 0xf4, 0xe1, 0xf6, 0x4b, 0x3d, 0x09, 0x03, 0xff, 0xd9, 0x02, 0x6d, 0xed, 0xa2, 0xaf,
	0xc1, 0xb6, 0x32, 0xee, 0xd9, 0xff, 0xf1, 0x12, 0xae, 0xbd, 0xb4, 0xf2, 0xc9, 0x7f, 0x19, 0x6a,
	0x4b, 0xc9, 0x9d, 0x10, 0xf8, 0x39, 0x68, 0x8b, 0x62, 0x58, 0x50, 0xb5, 0xe2, 0x77, 0x06, 0xfb,
	0x6f, 0x91, 0xf5, 0x4f, 0x25, 0x29, 0x7c, 0x28, 0x2d, 0xa6, 0xf8, 0x8d, 0xc5, 0x54, 0x88, 0x23,
	0x0d, 0xc3, 0xaf, 0x00, 0xd0, 0x4f, 0x2a, 0x32, 0x1a, 0x2b, 0x0b, 0x74, 0x07, 0x7b, 0x6f, 0x11,
	0x95, 0x0b, 0x0e, 0xf7, 0xea, 0x01, 0xd9, 0xb2, 0xa0, 0x59, 0xab, 0x8c, 0x70, 0xd4, 0x49, 0x56,
	0x3c, 0xfc, 0x9b, 0x09, 0xda, 0xea, 0x21, 0x20, 0x06, 0x5b, 0x2c, 0x7d, 0x31, 0x1c, 0x33, 0xe2,
	0x18, 0xee, 0x7b, 0xd3, 0x99, 0x77, 0x57, 0xab, 0xc9, 0xe4, 0x89, 0x4e, 0xc0, 0x07, 0xc0, 0xe6,
	0x19, 0x4d, 0x1d, 0xd3, 0xdd, 0x99, 0xce, 0x3c, 0xed, 0xa7, 0x2f, 0x32, 0x9a, 0xc2, 0x7d, 0xb0,
	0xc5, 0x73, 0x42, 0x73, 0x4a, 0x1c, 0xcb, 0x75, 0xa6, 0x33, 0xaf, 0xa7, 0x73, 0x1a, 0x93, 0xe9,
	0xc9, 0xb0, 0x88, 0x2f, 0x28, 0x71, 0x5a, 0x6b, 0xe9, 0xe7, 0x1a, 0x83, 0x03, 0x00, 0x59, 0x2a,
	0xca, 0xf3, 0x73, 0x16, 0x33, 0x9a, 0x16, 0x67, 0xe7, 0x65, 0x4a, 0x84, 0x63, 0xbb, 0xee, 0x74,
	0xe6, 0xbd, 0xaf, 0xc7, 0xbd, 0x96, 0xfe, 0x44, 0x66, 0xe1, 0x1e, 0xb8, 0x13, 0xcb, 0x7b, 0x40,
	0x9c, 0xb6, 0xfb, 0xee, 0x74, 0xe6, 0x75, 0x15, 0x4f, 0x9d, 0x08, 0xe2, 0xda, 0x2f, 0x7f, 0xee,
	0x1b, 0xf5, 0xa7, 0xfd, 0x8b, 0x09, 0xb6, 0x57, 0xbe, 0x81, 0x1f, 0x03, 0xbb, 0x4c, 0x59, 0x51,
	0x6f, 0xbb, 0x5b, 0x4f, 0xf0, 0xcb, 0x94, 0x15, 0xcd, 0xc4, 0x24, 0xa1, 0x99, 0x98, 0x8c, 0x70,
	0xa4, 0x40, 0x18, 0x80, 0x76, 0xcc, 0xcb, 0xb4, 0x50, 0x5b, 0xdd, 0xd1, 0x6b, 0x53, 0x40, 0xb3,
	0x36, 0x15, 0xe2, 0x48, 0xc3, 0xf0, 0x08, 0xb4, 0xb3, 0x9c, 0xc5, 0xb4, 0xde, 0x58, 0x6f, 0xf5,
	0xc9, 0x3c, 0xe5, 0x2c, 0xd5, 0xc7, 0x49, 0x4a, 0x28, 0x4a, 0x23, 0xa1, 0x42, 0x1c, 0x69, 0x58,
	0xbd, 0x84, 0x19, 0x1e, 0xbf, 0x9a, 0xf7, 0xcd, 0xd7, 0xf3, 0xbe, 0xf9, 0xc7, 0xbc, 0x6f, 0xfe,
	0x78, 0xdb, 0x37, 0x5e, 0xdf, 0xf6, 0x8d, 0x5f, 0x6f, 0xfb, 0xc6, 0x37, 0x1f, 0xad, 0x5d, 0x20,
	0xfe, 0x22, 0x8f, 0xc7, 0x97, 0x81, 0xbe, 0xab, 0xdf, 0xaf, 0xdf, 0x67, 0x75, 0x89, 0x46, 0x77,
	0xd4, 0x35, 0x7d, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x0b, 0xc8, 0x3c, 0xc0, 0x05,
	0x00, 0x00,
}

func (this *GroupID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GroupID)
	if !ok {
		that2, ok := that.(GroupID)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Owner, that1.Owner) {
		return false
	}
	if this.DSeq != that1.DSeq {
		return false
	}
	if this.GSeq != that1.GSeq {
		return false
	}
	return true
}
func (this *Resource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Resource)
	if !ok {
		that2, ok := that.(Resource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Unit.Equal(&that1.Unit) {
		return false
	}
	if this.Count != that1.Count {
		return false
	}
	if !this.Price.Equal(&that1.Price) {
		return false
	}
	return true
}
func (m *MsgCloseGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GroupID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GSeq != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GroupSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Resources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGroup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Requirements) > 0 {
		for iNdEx := len(m.Requirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGroup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Group) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GroupSpec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.GroupID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Count != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Unit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCloseGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func (m *GroupID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovGroup(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovGroup(uint64(m.GSeq))
	}
	return n
}

func (m *GroupSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if len(m.Requirements) > 0 {
		for _, e := range m.Requirements {
			l = e.Size()
			n += 1 + l + sovGroup(uint64(l))
		}
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = e.Size()
			n += 1 + l + sovGroup(uint64(l))
		}
	}
	return n
}

func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GroupID.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.State != 0 {
		n += 1 + sovGroup(uint64(m.State))
	}
	l = m.GroupSpec.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Unit.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.Count != 0 {
		n += 1 + sovGroup(uint64(m.Count))
	}
	l = m.Price.Size()
	n += 1 + l + sovGroup(uint64(l))
	return n
}

func sovGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroup(x uint64) (n int) {
	return sovGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCloseGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCloseGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requirements = append(m.Requirements, types.Attribute{})
			if err := m.Requirements[len(m.Requirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, Resource{})
			if err := m.Resources[len(m.Resources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Group_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Unit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroup
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroup = fmt.Errorf("proto: unexpected end of group")
)
