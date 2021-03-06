// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/authpb/auth.proto

package authpb

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
	RedirectUrl          string   `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
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

func (m *LoginResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
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
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
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

func (m *CallBackResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTokenRequest) Reset()         { *m = VerifyTokenRequest{} }
func (m *VerifyTokenRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenRequest) ProtoMessage()    {}
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5393af7afe3ec9a6, []int{4}
}

func (m *VerifyTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenRequest.Unmarshal(m, b)
}
func (m *VerifyTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenRequest.Marshal(b, m, deterministic)
}
func (m *VerifyTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenRequest.Merge(m, src)
}
func (m *VerifyTokenRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenRequest.Size(m)
}
func (m *VerifyTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenRequest proto.InternalMessageInfo

func (m *VerifyTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyTokenResponse struct {
	IsLoggedId           bool     `protobuf:"varint,1,opt,name=is_logged_id,json=isLoggedId,proto3" json:"is_logged_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTokenResponse) Reset()         { *m = VerifyTokenResponse{} }
func (m *VerifyTokenResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenResponse) ProtoMessage()    {}
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5393af7afe3ec9a6, []int{5}
}

func (m *VerifyTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenResponse.Unmarshal(m, b)
}
func (m *VerifyTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenResponse.Marshal(b, m, deterministic)
}
func (m *VerifyTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenResponse.Merge(m, src)
}
func (m *VerifyTokenResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenResponse.Size(m)
}
func (m *VerifyTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenResponse proto.InternalMessageInfo

func (m *VerifyTokenResponse) GetIsLoggedId() bool {
	if m != nil {
		return m.IsLoggedId
	}
	return false
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*CallBackRequest)(nil), "auth.CallBackRequest")
	proto.RegisterType((*CallBackResponse)(nil), "auth.CallBackResponse")
	proto.RegisterType((*VerifyTokenRequest)(nil), "auth.VerifyTokenRequest")
	proto.RegisterType((*VerifyTokenResponse)(nil), "auth.VerifyTokenResponse")
}

func init() {
	proto.RegisterFile("src/authpb/auth.proto", fileDescriptor_5393af7afe3ec9a6)
}

var fileDescriptor_5393af7afe3ec9a6 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4e, 0xfa, 0x40,
	0x14, 0xc5, 0x03, 0x7f, 0x20, 0x70, 0x5b, 0xfe, 0x92, 0x01, 0x14, 0x1b, 0x4d, 0x6a, 0x13, 0x13,
	0xe2, 0x82, 0x26, 0xba, 0x70, 0x2d, 0xea, 0xc2, 0x84, 0x55, 0xfd, 0x08, 0x71, 0x43, 0x86, 0x76,
	0x2c, 0x13, 0xc6, 0x4e, 0x9d, 0x99, 0x92, 0xb8, 0x75, 0xe5, 0xde, 0x47, 0xf3, 0x15, 0x7c, 0x10,
	0xc3, 0x4c, 0x1b, 0xf9, 0x5a, 0xf5, 0xce, 0xed, 0xef, 0x9e, 0x93, 0x39, 0x73, 0xa1, 0x2b, 0x45,
	0xe8, 0xe3, 0x4c, 0xcd, 0xd2, 0xa9, 0xfe, 0x0c, 0x52, 0xc1, 0x15, 0x47, 0x95, 0x65, 0xed, 0x1c,
	0xc5, 0x9c, 0xc7, 0x8c, 0xf8, 0x38, 0xa5, 0x3e, 0x4e, 0x12, 0xae, 0xb0, 0xa2, 0x3c, 0x91, 0x86,
	0xf1, 0x5c, 0xb0, 0x47, 0x3c, 0xa6, 0x49, 0x40, 0xde, 0x32, 0x22, 0x15, 0x6a, 0xc1, 0xbf, 0x24,
	0x7b, 0xed, 0x95, 0xdc, 0x52, 0xbf, 0x1a, 0x2c, 0x4b, 0xef, 0x06, 0x9a, 0x39, 0x21, 0x53, 0x9e,
	0x48, 0xb2, 0x8d, 0xa0, 0x13, 0xb0, 0x05, 0x89, 0xa8, 0x20, 0xa1, 0x9a, 0x64, 0x82, 0xf5, 0xca,
	0x6e, 0xa9, 0xdf, 0x08, 0xac, 0xa2, 0xf7, 0x28, 0x98, 0x77, 0x0a, 0x7b, 0xd7, 0x98, 0xb1, 0x21,
	0x0e, 0xe7, 0x85, 0x15, 0x82, 0x4a, 0xc8, 0x23, 0xa2, 0x85, 0x1a, 0x81, 0xae, 0xbd, 0x3e, 0xb4,
	0xfe, 0xb0, 0xdc, 0xaf, 0x03, 0x55, 0xc5, 0xe7, 0x24, 0xc9, 0x41, 0x73, 0xf0, 0xce, 0x00, 0x3d,
	0x11, 0x41, 0x5f, 0xde, 0x1f, 0x96, 0xc7, 0x42, 0x73, 0x37, 0x7b, 0x09, 0xed, 0x35, 0x36, 0x17,
	0x76, 0xc1, 0xa6, 0x72, 0xc2, 0x78, 0x1c, 0x93, 0x68, 0x42, 0x23, 0x3d, 0x53, 0x0f, 0x80, 0xca,
	0x91, 0x6e, 0xdd, 0x45, 0xe7, 0x9f, 0x65, 0xb0, 0xae, 0x32, 0x35, 0xbb, 0x27, 0x62, 0x41, 0x43,
	0x82, 0x6e, 0xa1, 0xaa, 0xb3, 0x40, 0x68, 0xa0, 0x73, 0x5e, 0x8d, 0xce, 0x69, 0xaf, 0xf5, 0x8c,
	0x87, 0xd7, 0xfe, 0xf8, 0xfe, 0xf9, 0x2a, 0x37, 0x91, 0xa5, 0x1f, 0xc6, 0x67, 0x7a, 0x7a, 0x0c,
	0xf5, 0xe2, 0x96, 0xa8, 0x6b, 0xa6, 0x36, 0xc2, 0x71, 0xf6, 0x37, 0xdb, 0xb9, 0xde, 0xb1, 0xd6,
	0x3b, 0x40, 0x5d, 0xa3, 0x17, 0x62, 0xc6, 0xa6, 0x38, 0x9c, 0xfb, 0xe6, 0x91, 0xd1, 0x18, 0xac,
	0x95, 0x9b, 0xa2, 0x9e, 0x51, 0xd9, 0x0e, 0xca, 0x39, 0xdc, 0xf1, 0x27, 0xb7, 0xe8, 0x68, 0x8b,
	0xff, 0xc8, 0x36, 0x16, 0x0b, 0x8d, 0x0c, 0xeb, 0xcf, 0x35, 0xb3, 0x61, 0xd3, 0x9a, 0xde, 0x9c,
	0x8b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x20, 0x43, 0x42, 0x76, 0x02, 0x00, 0x00,
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
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
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

func (c *authServiceClient) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CallBack(context.Context, *CallBackRequest) (*CallBackResponse, error)
	VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error)
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
func (*UnimplementedAuthServiceServer) VerifyToken(ctx context.Context, req *VerifyTokenRequest) (*VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
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

func _AuthService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyToken(ctx, req.(*VerifyTokenRequest))
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
		{
			MethodName: "VerifyToken",
			Handler:    _AuthService_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/authpb/auth.proto",
}
