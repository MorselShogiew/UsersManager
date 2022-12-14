// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

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

type Status int32

const (
	Status_Success       Status = 0
	Status_InternalError Status = 1
	Status_NotFound      Status = 2
	Status_NothinToDo    Status = 3
	Status_RequestError  Status = 4
	Status_Unknown       Status = 5
)

var Status_name = map[int32]string{
	0: "Success",
	1: "InternalError",
	2: "NotFound",
	3: "NothinToDo",
	4: "RequestError",
	5: "Unknown",
}

var Status_value = map[string]int32{
	"Success":       0,
	"InternalError": 1,
	"NotFound":      2,
	"NothinToDo":    3,
	"RequestError":  4,
	"Unknown":       5,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SaveUserRequest struct {
	Data                 *User    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveUserRequest) Reset()         { *m = SaveUserRequest{} }
func (m *SaveUserRequest) String() string { return proto.CompactTextString(m) }
func (*SaveUserRequest) ProtoMessage()    {}
func (*SaveUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *SaveUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveUserRequest.Unmarshal(m, b)
}
func (m *SaveUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveUserRequest.Marshal(b, m, deterministic)
}
func (m *SaveUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveUserRequest.Merge(m, src)
}
func (m *SaveUserRequest) XXX_Size() int {
	return xxx_messageInfo_SaveUserRequest.Size(m)
}
func (m *SaveUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveUserRequest proto.InternalMessageInfo

func (m *SaveUserRequest) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

type SaveUserResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=grpc_psql_example.proto.user.Status" json:"status,omitempty"`
	Data                 *User    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveUserResponse) Reset()         { *m = SaveUserResponse{} }
func (m *SaveUserResponse) String() string { return proto.CompactTextString(m) }
func (*SaveUserResponse) ProtoMessage()    {}
func (*SaveUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *SaveUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveUserResponse.Unmarshal(m, b)
}
func (m *SaveUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveUserResponse.Marshal(b, m, deterministic)
}
func (m *SaveUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveUserResponse.Merge(m, src)
}
func (m *SaveUserResponse) XXX_Size() int {
	return xxx_messageInfo_SaveUserResponse.Size(m)
}
func (m *SaveUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveUserResponse proto.InternalMessageInfo

func (m *SaveUserResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Success
}

func (m *SaveUserResponse) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteUserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteUserResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=grpc_psql_example.proto.user.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Success
}

type GetUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersRequest) Reset()         { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()    {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *GetUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersRequest.Unmarshal(m, b)
}
func (m *GetUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersRequest.Merge(m, src)
}
func (m *GetUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetUsersRequest.Size(m)
}
func (m *GetUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersRequest proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("grpc_psql_example.proto.user.Status", Status_name, Status_value)
	proto.RegisterType((*User)(nil), "grpc_psql_example.proto.user.User")
	proto.RegisterType((*SaveUserRequest)(nil), "grpc_psql_example.proto.user.SaveUserRequest")
	proto.RegisterType((*SaveUserResponse)(nil), "grpc_psql_example.proto.user.SaveUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "grpc_psql_example.proto.user.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "grpc_psql_example.proto.user.DeleteUserResponse")
	proto.RegisterType((*GetUsersRequest)(nil), "grpc_psql_example.proto.user.GetUsersRequest")
}

func init() {
	proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7)
}

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0xd7, 0x09, 0x66, 0x5a, 0x52, 0x67, 0x04, 0x52, 0x14, 0x71, 0xa8, 0x16, 0x0e,
	0x15, 0x52, 0xec, 0x12, 0x24, 0x4e, 0x1c, 0x10, 0x2a, 0xa0, 0x5e, 0x7a, 0x70, 0xe8, 0x85, 0x4b,
	0xb5, 0xb5, 0x47, 0xad, 0x15, 0x7b, 0xd7, 0xd9, 0x5d, 0x43, 0x1e, 0x81, 0x97, 0xe3, 0x9d, 0x90,
	0xd7, 0x4e, 0x42, 0x82, 0x94, 0x04, 0xa9, 0x17, 0x6b, 0xd7, 0xfe, 0xff, 0x6f, 0x67, 0xe6, 0xf7,
	0xc2, 0x8b, 0x52, 0x49, 0x23, 0xa3, 0x4a, 0x93, 0xb2, 0x8f, 0xd0, 0xee, 0xf1, 0xe5, 0xbd, 0x2a,
	0x93, 0xdb, 0x52, 0xcf, 0xf3, 0x5b, 0x5a, 0xf0, 0xa2, 0xcc, 0xa9, 0xf9, 0x10, 0xd6, 0x1a, 0xf6,
	0x11, 0xbc, 0x1b, 0x4d, 0x0a, 0xfb, 0xe0, 0x66, 0xe9, 0xd0, 0x39, 0x73, 0xce, 0xbb, 0xb1, 0x9b,
	0xa5, 0x88, 0xe0, 0x09, 0x5e, 0xd0, 0xd0, 0x3d, 0x73, 0xce, 0x9f, 0xc6, 0x76, 0x8d, 0xcf, 0xa1,
	0x4b, 0x05, 0xcf, 0xf2, 0xe1, 0x91, 0x7d, 0xd9, 0x6c, 0xd8, 0x15, 0x9c, 0x4e, 0xf9, 0x0f, 0xaa,
	0x29, 0x31, 0xcd, 0x2b, 0xd2, 0x06, 0xdf, 0x83, 0x97, 0x72, 0xc3, 0x2d, 0xee, 0x78, 0xc2, 0xc2,
	0x5d, 0x15, 0x84, 0xd6, 0x68, 0xf5, 0xec, 0x97, 0x03, 0xc1, 0x9a, 0xa5, 0x4b, 0x29, 0x34, 0xe1,
	0x07, 0xe8, 0x69, 0xc3, 0x4d, 0xa5, 0x2d, 0xae, 0x3f, 0x79, 0xbd, 0x1b, 0x37, 0xb5, 0xda, 0xb8,
	0xf5, 0xac, 0x4a, 0x71, 0xff, 0xb3, 0x94, 0x57, 0x30, 0xb8, 0xa4, 0x9c, 0xcc, 0x46, 0x5f, 0x5b,
	0x43, 0x62, 0x31, 0xe0, 0xdf, 0xa2, 0xc7, 0x28, 0x98, 0x0d, 0xe0, 0xf4, 0x2b, 0x99, 0x1a, 0xa8,
	0xdb, 0x63, 0xdf, 0x10, 0xf4, 0x1a, 0x11, 0x1e, 0xc3, 0x93, 0x69, 0x95, 0x24, 0xa4, 0x75, 0xd0,
	0xc1, 0x01, 0x3c, 0xbb, 0x12, 0x86, 0x94, 0xe0, 0xf9, 0x67, 0xa5, 0xa4, 0x0a, 0x1c, 0x3c, 0x01,
	0xff, 0x5a, 0x9a, 0x2f, 0xb2, 0x12, 0x69, 0xe0, 0x62, 0x1f, 0xe0, 0x5a, 0x9a, 0x87, 0x4c, 0x7c,
	0x93, 0x97, 0x32, 0x38, 0xc2, 0x00, 0x4e, 0x5a, 0x64, 0xa3, 0xf7, 0x6a, 0xde, 0x8d, 0x98, 0x09,
	0xf9, 0x53, 0x04, 0xdd, 0xc9, 0x6f, 0x17, 0xfc, 0xa6, 0x91, 0x52, 0xe2, 0x0c, 0xfc, 0x65, 0x12,
	0x38, 0xde, 0xd3, 0xc0, 0x66, 0xfa, 0xa3, 0xf0, 0x50, 0x79, 0x33, 0x2f, 0xd6, 0xc1, 0x39, 0xc0,
	0x7a, 0x8e, 0x18, 0xed, 0xf6, 0xff, 0x13, 0xcb, 0xe8, 0xe2, 0x70, 0xc3, 0xea, 0x48, 0x0e, 0xfe,
	0x72, 0xcc, 0xfb, 0xfa, 0xdb, 0x8a, 0x63, 0x74, 0xc0, 0x4f, 0xc4, 0x3a, 0x17, 0xce, 0xa7, 0xb7,
	0xdf, 0xa3, 0xfb, 0xcc, 0x3c, 0x54, 0x77, 0x61, 0x22, 0x8b, 0x88, 0x16, 0xba, 0xa8, 0x44, 0x1a,
	0xd5, 0xde, 0x71, 0xed, 0x1d, 0xb7, 0xde, 0x68, 0x7d, 0x6d, 0xef, 0x7a, 0x76, 0xfd, 0xee, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xb0, 0x12, 0x34, 0xcb, 0x03, 0x00, 0x00,
}
