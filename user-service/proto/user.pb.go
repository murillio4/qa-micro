// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	proto/user.proto

It has these top-level messages:
	User
	Role
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id        string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email     string  `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	FirstName string  `protobuf:"bytes,4,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string  `protobuf:"bytes,5,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Name      string  `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	Picture   string  `protobuf:"bytes,7,opt,name=picture" json:"picture,omitempty"`
	Role      []*Role `protobuf:"bytes,8,rep,name=role" json:"role,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *User) GetRole() []*Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type Role struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*Role)(nil), "Role")
}

func init() { proto.RegisterFile("proto/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x0a, 0x82, 0x40,
	0x14, 0x85, 0x71, 0x1a, 0xff, 0xae, 0x14, 0x31, 0xb4, 0x98, 0x88, 0x40, 0x5c, 0x49, 0x0b, 0x03,
	0x7b, 0x8f, 0x16, 0x46, 0xeb, 0x98, 0xf4, 0x06, 0x03, 0x63, 0x23, 0xa3, 0xf6, 0x5e, 0xbd, 0x61,
	0x78, 0x43, 0x5a, 0xb4, 0xbb, 0xe7, 0xfb, 0x16, 0xe7, 0x1e, 0x58, 0x77, 0xce, 0x0e, 0xf6, 0x38,
	0xf6, 0xe8, 0x0a, 0x3a, 0xb3, 0xb7, 0x07, 0xfc, 0xda, 0xa3, 0x13, 0x2b, 0x60, 0xba, 0x91, 0x5e,
	0xea, 0xe5, 0x71, 0xc5, 0x74, 0x23, 0x36, 0xe0, 0x63, 0xab, 0xb4, 0x91, 0x8c, 0xd0, 0x37, 0x88,
	0x3d, 0xc0, 0x43, 0xbb, 0x7e, 0xb8, 0x3d, 0x55, 0x8b, 0x92, 0x93, 0x8a, 0x89, 0x9c, 0x55, 0x8b,
	0x62, 0x07, 0xb1, 0x51, 0xb3, 0xf5, 0xc9, 0x46, 0x13, 0x20, 0x29, 0x80, 0x13, 0x0f, 0x88, 0xd3,
	0x2d, 0x24, 0x84, 0x9d, 0xae, 0x87, 0xd1, 0xa1, 0x0c, 0x09, 0xcf, 0x51, 0x6c, 0x81, 0x3b, 0x6b,
	0x50, 0x46, 0xe9, 0x22, 0x4f, 0x4a, 0xbf, 0xa8, 0xac, 0xc1, 0x8a, 0x50, 0x76, 0x00, 0x3e, 0xa5,
	0xbf, 0x97, 0xe7, 0x02, 0xf6, 0x2b, 0x28, 0x97, 0x90, 0x4c, 0xf3, 0x2e, 0xe8, 0x5e, 0xba, 0xc6,
	0x7b, 0x40, 0xab, 0x4f, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xb3, 0x3d, 0xa6, 0x09, 0x01,
	0x00, 0x00,
}
