// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/authpb/auth.proto

package authpb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LoginRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5393af7afe3ec9a6, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type LoginResponse struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5393af7afe3ec9a6, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type CallBackRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallBackRequest) Reset()         { *m = CallBackRequest{} }
func (m *CallBackRequest) String() string { return proto.CompactTextString(m) }
func (*CallBackRequest) ProtoMessage()    {}
func (*CallBackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5393af7afe3ec9a6, []int{2}
}

func (m *CallBackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallBackRequest.Unmarshal(m, b)
}
func (m *CallBackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallBackRequest.Marshal(b, m, deterministic)
}
func (m *CallBackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallBackRequest.Merge(m, src)
}
func (m *CallBackRequest) XXX_Size() int {
	return xxx_messageInfo_CallBackRequest.Size(m)
}
func (m *CallBackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallBackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallBackRequest proto.InternalMessageInfo

func (m *CallBackRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type CallBackResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallBackResponse) Reset()         { *m = CallBackResponse{} }
func (m *CallBackResponse) String() string { return proto.CompactTextString(m) }
func (*CallBackResponse) ProtoMessage()    {}
func (*CallBackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5393af7afe3ec9a6, []int{3}
}

func (m *CallBackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallBackResponse.Unmarshal(m, b)
}
func (m *CallBackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallBackResponse.Marshal(b, m, deterministic)
}
func (m *CallBackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallBackResponse.Merge(m, src)
}
func (m *CallBackResponse) XXX_Size() int {
	return xxx_messageInfo_CallBackResponse.Size(m)
}
func (m *CallBackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallBackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallBackResponse proto.InternalMessageInfo

func (m *CallBackResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*CallBackRequest)(nil), "auth.CallBackRequest")
	proto.RegisterType((*CallBackResponse)(nil), "auth.CallBackResponse")
}

func init() {
	proto.RegisterFile("src/authpb/auth.proto", fileDescriptor_5393af7afe3ec9a6)
}

var fileDescriptor_5393af7afe3ec9a6 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0x4a, 0xd6,
	0x4f, 0x2c, 0x2d, 0xc9, 0x28, 0x48, 0x02, 0x53, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c,
	0x20, 0xb6, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62,
	0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x8d, 0x92, 0x02, 0x17, 0x8f,
	0x4f, 0x7e, 0x7a, 0x66, 0x5e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x00, 0x17, 0x73,
	0x5e, 0x69, 0xae, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x88, 0xa9, 0xa4, 0xc8, 0xc5, 0x0b,
	0x55, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0x45, 0x89, 0x2a, 0x17, 0xbf, 0x73, 0x62, 0x4e,
	0x8e, 0x53, 0x62, 0x72, 0x36, 0xcc, 0x1c, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0xb0, 0x2a,
	0xce, 0x20, 0x30, 0x5b, 0x49, 0x8b, 0x4b, 0x00, 0xa1, 0x0c, 0x6a, 0x98, 0x18, 0x17, 0x5b, 0x51,
	0x6a, 0x71, 0x69, 0x4e, 0x09, 0x54, 0x25, 0x94, 0x67, 0xb4, 0x8e, 0x91, 0x8b, 0xdb, 0xb1, 0xb4,
	0x24, 0x23, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x95, 0x8b, 0x15, 0xec, 0x0a, 0x21,
	0x21, 0x3d, 0xb0, 0x0f, 0x91, 0x1d, 0x2d, 0x25, 0x8c, 0x22, 0x06, 0x31, 0x59, 0x49, 0xb8, 0xe9,
	0xf2, 0x93, 0xc9, 0x4c, 0xbc, 0x42, 0xdc, 0xe0, 0x20, 0xd1, 0xcf, 0x01, 0xeb, 0x8e, 0xe0, 0xe2,
	0x80, 0x39, 0x41, 0x48, 0x14, 0xa2, 0x0b, 0xcd, 0xe5, 0x52, 0x62, 0xe8, 0xc2, 0x50, 0xf3, 0x64,
	0xc1, 0xe6, 0x89, 0x0b, 0x89, 0x42, 0xcc, 0x4b, 0x4e, 0xcc, 0xc9, 0x49, 0x4a, 0x4c, 0xce, 0xd6,
	0x87, 0x04, 0xaf, 0x13, 0x47, 0x14, 0x1b, 0x24, 0x06, 0x92, 0xd8, 0xc0, 0x21, 0x6b, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x04, 0x4e, 0x55, 0x7a, 0x96, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CallBack(ctx context.Context, in *CallBackRequest, opts ...grpc.CallOption) (*CallBackResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CallBack(ctx context.Context, in *CallBackRequest, opts ...grpc.CallOption) (*CallBackResponse, error) {
	out := new(CallBackResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CallBack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CallBack(context.Context, *CallBackRequest) (*CallBackResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) CallBack(ctx context.Context, req *CallBackRequest) (*CallBackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallBack not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CallBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallBackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CallBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CallBack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CallBack(ctx, req.(*CallBackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "CallBack",
			Handler:    _AuthService_CallBack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/authpb/auth.proto",
}