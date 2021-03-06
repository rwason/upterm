// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Identifier_Type int32

const (
	Identifier_HOST   Identifier_Type = 0
	Identifier_CLIENT Identifier_Type = 1
)

var Identifier_Type_name = map[int32]string{
	0: "HOST",
	1: "CLIENT",
}

var Identifier_Type_value = map[string]int32{
	"HOST":   0,
	"CLIENT": 1,
}

func (x Identifier_Type) String() string {
	return proto.EnumName(Identifier_Type_name, int32(x))
}

func (Identifier_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3, 0}
}

type GetSessionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionRequest) Reset()         { *m = GetSessionRequest{} }
func (m *GetSessionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionRequest) ProtoMessage()    {}
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *GetSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionRequest.Unmarshal(m, b)
}
func (m *GetSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionRequest.Marshal(b, m, deterministic)
}
func (m *GetSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionRequest.Merge(m, src)
}
func (m *GetSessionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionRequest.Size(m)
}
func (m *GetSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionRequest proto.InternalMessageInfo

type GetSessionResponse struct {
	SessionId            string    `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Command              []string  `protobuf:"bytes,2,rep,name=command,proto3" json:"command,omitempty"`
	ForceCommand         []string  `protobuf:"bytes,3,rep,name=force_command,json=forceCommand,proto3" json:"force_command,omitempty"`
	Host                 string    `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	NodeAddr             string    `protobuf:"bytes,5,opt,name=node_addr,json=nodeAddr,proto3" json:"node_addr,omitempty"`
	ConnectedClients     []*Client `protobuf:"bytes,6,rep,name=connected_clients,json=connectedClients,proto3" json:"connected_clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetSessionResponse) Reset()         { *m = GetSessionResponse{} }
func (m *GetSessionResponse) String() string { return proto.CompactTextString(m) }
func (*GetSessionResponse) ProtoMessage()    {}
func (*GetSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *GetSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionResponse.Unmarshal(m, b)
}
func (m *GetSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionResponse.Marshal(b, m, deterministic)
}
func (m *GetSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionResponse.Merge(m, src)
}
func (m *GetSessionResponse) XXX_Size() int {
	return xxx_messageInfo_GetSessionResponse.Size(m)
}
func (m *GetSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionResponse proto.InternalMessageInfo

func (m *GetSessionResponse) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *GetSessionResponse) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *GetSessionResponse) GetForceCommand() []string {
	if m != nil {
		return m.ForceCommand
	}
	return nil
}

func (m *GetSessionResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GetSessionResponse) GetNodeAddr() string {
	if m != nil {
		return m.NodeAddr
	}
	return ""
}

func (m *GetSessionResponse) GetConnectedClients() []*Client {
	if m != nil {
		return m.ConnectedClients
	}
	return nil
}

type Client struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	PublicKeyFingerprint string   `protobuf:"bytes,4,opt,name=public_key_fingerprint,json=publicKeyFingerprint,proto3" json:"public_key_fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Client) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Client) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Client) GetPublicKeyFingerprint() string {
	if m != nil {
		return m.PublicKeyFingerprint
	}
	return ""
}

type Identifier struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Identifier_Type `protobuf:"varint,2,opt,name=type,proto3,enum=api.Identifier_Type" json:"type,omitempty"`
	NodeAddr             string          `protobuf:"bytes,3,opt,name=node_addr,json=nodeAddr,proto3" json:"node_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Identifier) GetType() Identifier_Type {
	if m != nil {
		return m.Type
	}
	return Identifier_HOST
}

func (m *Identifier) GetNodeAddr() string {
	if m != nil {
		return m.NodeAddr
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.Identifier_Type", Identifier_Type_name, Identifier_Type_value)
	proto.RegisterType((*GetSessionRequest)(nil), "api.GetSessionRequest")
	proto.RegisterType((*GetSessionResponse)(nil), "api.GetSessionResponse")
	proto.RegisterType((*Client)(nil), "api.Client")
	proto.RegisterType((*Identifier)(nil), "api.Identifier")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0xed, 0x1f, 0x2b, 0x3d, 0xf7, 0x7a, 0xd3, 0x3b, 0xde, 0x5c, 0x1b, 0xc4, 0x84, 0xd4,
	0x4d, 0x57, 0x90, 0xa0, 0x0b, 0xb7, 0x84, 0xf8, 0x87, 0x68, 0x24, 0x29, 0xec, 0x9b, 0xd2, 0x39,
	0xe0, 0x44, 0x98, 0x19, 0x67, 0x06, 0x92, 0x26, 0x6e, 0xf4, 0x15, 0x7c, 0x34, 0xd7, 0xee, 0x7c,
	0x10, 0xd3, 0x29, 0x85, 0x4b, 0xd8, 0x9d, 0xf9, 0x7d, 0x5f, 0x7a, 0xce, 0x77, 0x7a, 0x20, 0x2c,
	0x24, 0x1b, 0x48, 0x25, 0x8c, 0x20, 0x5e, 0x21, 0x59, 0xb7, 0xb7, 0x16, 0x62, 0xbd, 0xc1, 0x61,
	0x21, 0xd9, 0xb0, 0xe0, 0x5c, 0x98, 0xc2, 0x30, 0xc1, 0x75, 0x63, 0x49, 0x9e, 0xc1, 0xed, 0x07,
	0x34, 0x73, 0xd4, 0x9a, 0x09, 0x9e, 0xe1, 0xf7, 0x1d, 0x6a, 0x93, 0xfc, 0x75, 0x80, 0x3c, 0xa4,
	0x5a, 0x0a, 0xae, 0x91, 0xbc, 0x04, 0xd0, 0x0d, 0xca, 0x19, 0x8d, 0x9d, 0xbe, 0x93, 0x86, 0x59,
	0x78, 0x20, 0x53, 0x4a, 0x62, 0x78, 0x52, 0x8a, 0xed, 0xb6, 0xe0, 0x34, 0x76, 0xfb, 0x5e, 0x1a,
	0x66, 0xed, 0x93, 0xbc, 0x82, 0xa7, 0x2b, 0xa1, 0x4a, 0xcc, 0x5b, 0xdd, 0xb3, 0xfa, 0xb5, 0x85,
	0x93, 0x83, 0x89, 0x80, 0xff, 0x55, 0x68, 0x13, 0xfb, 0xf6, 0xbb, 0xb6, 0x26, 0x2f, 0x20, 0xe4,
	0x82, 0x62, 0x5e, 0x50, 0xaa, 0xe2, 0xc7, 0x56, 0xe8, 0xd4, 0x60, 0x4c, 0xa9, 0x22, 0x6f, 0xe1,
	0xb6, 0x14, 0x9c, 0x63, 0x69, 0x90, 0xe6, 0xe5, 0x86, 0x21, 0x37, 0x3a, 0x0e, 0xfa, 0x5e, 0x7a,
	0x35, 0xba, 0x1a, 0xd4, 0x4b, 0x98, 0x58, 0x96, 0x45, 0x47, 0x57, 0x03, 0x74, 0xf2, 0x03, 0x82,
	0xa6, 0x24, 0x37, 0xe0, 0x1e, 0xa3, 0xb8, 0xcc, 0x66, 0xd8, 0xa3, 0xaa, 0x03, 0xc5, 0xae, 0x85,
	0xed, 0xb3, 0x1e, 0xcf, 0x4e, 0xe1, 0x35, 0xe3, 0xd5, 0x35, 0x79, 0x03, 0xf7, 0x72, 0xb7, 0xdc,
	0xb0, 0x32, 0xff, 0x86, 0x55, 0xbe, 0x62, 0x7c, 0x8d, 0x4a, 0x2a, 0xc6, 0xdb, 0x10, 0x77, 0x8d,
	0xfa, 0x09, 0xab, 0xf7, 0x27, 0x2d, 0xf9, 0xe9, 0x00, 0x4c, 0x29, 0x72, 0xc3, 0x56, 0x0c, 0xd5,
	0xc5, 0x08, 0x29, 0xf8, 0xa6, 0x92, 0x68, 0xfb, 0xdf, 0x8c, 0xee, 0x6c, 0x92, 0x93, 0x7d, 0xb0,
	0xa8, 0x24, 0x66, 0xd6, 0x71, 0xbe, 0x1d, 0xef, 0x7c, 0x3b, 0x49, 0x0f, 0xfc, 0xda, 0x4a, 0x3a,
	0xe0, 0x7f, 0x9c, 0xcd, 0x17, 0xd1, 0x23, 0x02, 0x10, 0x4c, 0x3e, 0x4f, 0xdf, 0x7d, 0x59, 0x44,
	0xce, 0x28, 0x87, 0xeb, 0x31, 0xdd, 0x32, 0x3e, 0x47, 0xb5, 0x67, 0x25, 0x92, 0x19, 0xc0, 0xe9,
	0x87, 0x93, 0x7b, 0xdb, 0xf4, 0xe2, 0x2e, 0xba, 0xcf, 0x2f, 0x78, 0x73, 0x19, 0x49, 0xf4, 0xeb,
	0xcf, 0xbf, 0xdf, 0x2e, 0x90, 0xce, 0xf0, 0x70, 0x0e, 0xcb, 0xc0, 0x9e, 0xd7, 0xeb, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x15, 0x76, 0x9b, 0x1e, 0x8e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error)
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error) {
	out := new(GetSessionResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error)
}

// UnimplementedAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (*UnimplementedAdminServiceServer) GetSession(ctx context.Context, req *GetSessionRequest) (*GetSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _AdminService_GetSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
