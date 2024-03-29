// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: auth.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*BoolResponse, error)
}

type registerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterServiceClient(cc grpc.ClientConnInterface) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/auth.RegisterService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServiceServer is the server API for RegisterService service.
// All implementations must embed UnimplementedRegisterServiceServer
// for forward compatibility
type RegisterServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*BoolResponse, error)
	mustEmbedUnimplementedRegisterServiceServer()
}

// UnimplementedRegisterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServiceServer struct {
}

func (UnimplementedRegisterServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedRegisterServiceServer) mustEmbedUnimplementedRegisterServiceServer() {}

// UnsafeRegisterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServiceServer will
// result in compilation errors.
type UnsafeRegisterServiceServer interface {
	mustEmbedUnimplementedRegisterServiceServer()
}

func RegisterRegisterServiceServer(s grpc.ServiceRegistrar, srv RegisterServiceServer) {
	s.RegisterService(&RegisterService_ServiceDesc, srv)
}

func _RegisterService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RegisterService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterService_ServiceDesc is the grpc.ServiceDesc for RegisterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegisterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _RegisterService_RegisterUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AccessToken, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AccessToken, error) {
	out := new(AccessToken)
	err := c.cc.Invoke(ctx, "/auth.LoginService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	LoginUser(context.Context, *LoginRequest) (*AccessToken, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) LoginUser(context.Context, *LoginRequest) (*AccessToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.LoginService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _LoginService_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	AuthorizeUser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthUserData, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) AuthorizeUser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthUserData, error) {
	out := new(AuthUserData)
	err := c.cc.Invoke(ctx, "/auth.AuthorizationService/AuthorizeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
// All implementations must embed UnimplementedAuthorizationServiceServer
// for forward compatibility
type AuthorizationServiceServer interface {
	AuthorizeUser(context.Context, *Empty) (*AuthUserData, error)
	mustEmbedUnimplementedAuthorizationServiceServer()
}

// UnimplementedAuthorizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (UnimplementedAuthorizationServiceServer) AuthorizeUser(context.Context, *Empty) (*AuthUserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeUser not implemented")
}
func (UnimplementedAuthorizationServiceServer) mustEmbedUnimplementedAuthorizationServiceServer() {}

// UnsafeAuthorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServiceServer will
// result in compilation errors.
type UnsafeAuthorizationServiceServer interface {
	mustEmbedUnimplementedAuthorizationServiceServer()
}

func RegisterAuthorizationServiceServer(s grpc.ServiceRegistrar, srv AuthorizationServiceServer) {
	s.RegisterService(&AuthorizationService_ServiceDesc, srv)
}

func _AuthorizationService_AuthorizeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).AuthorizeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthorizationService/AuthorizeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).AuthorizeUser(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorizationService_ServiceDesc is the grpc.ServiceDesc for AuthorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeUser",
			Handler:    _AuthorizationService_AuthorizeUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

// UserDataServiceClient is the client API for UserDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDataServiceClient interface {
	GetCurrentUserData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error)
	GetOtherUserData(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*OtherUser, error)
}

type userDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataServiceClient(cc grpc.ClientConnInterface) UserDataServiceClient {
	return &userDataServiceClient{cc}
}

func (c *userDataServiceClient) GetCurrentUserData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.UserDataService/GetCurrentUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataServiceClient) GetOtherUserData(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*OtherUser, error) {
	out := new(OtherUser)
	err := c.cc.Invoke(ctx, "/auth.UserDataService/GetOtherUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataServiceServer is the server API for UserDataService service.
// All implementations must embed UnimplementedUserDataServiceServer
// for forward compatibility
type UserDataServiceServer interface {
	GetCurrentUserData(context.Context, *Empty) (*User, error)
	GetOtherUserData(context.Context, *UserId) (*OtherUser, error)
	mustEmbedUnimplementedUserDataServiceServer()
}

// UnimplementedUserDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserDataServiceServer struct {
}

func (UnimplementedUserDataServiceServer) GetCurrentUserData(context.Context, *Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUserData not implemented")
}
func (UnimplementedUserDataServiceServer) GetOtherUserData(context.Context, *UserId) (*OtherUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOtherUserData not implemented")
}
func (UnimplementedUserDataServiceServer) mustEmbedUnimplementedUserDataServiceServer() {}

// UnsafeUserDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataServiceServer will
// result in compilation errors.
type UnsafeUserDataServiceServer interface {
	mustEmbedUnimplementedUserDataServiceServer()
}

func RegisterUserDataServiceServer(s grpc.ServiceRegistrar, srv UserDataServiceServer) {
	s.RegisterService(&UserDataService_ServiceDesc, srv)
}

func _UserDataService_GetCurrentUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).GetCurrentUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserDataService/GetCurrentUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).GetCurrentUserData(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataService_GetOtherUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).GetOtherUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserDataService/GetOtherUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).GetOtherUserData(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDataService_ServiceDesc is the grpc.ServiceDesc for UserDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.UserDataService",
	HandlerType: (*UserDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentUserData",
			Handler:    _UserDataService_GetCurrentUserData_Handler,
		},
		{
			MethodName: "GetOtherUserData",
			Handler:    _UserDataService_GetOtherUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
