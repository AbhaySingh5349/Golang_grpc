// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: protos/blog.proto

package protos

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

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*BlogID, error)
	ReadBlog(ctx context.Context, in *BlogID, opts ...grpc.CallOption) (*Blog, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*BlogID, error) {
	out := new(BlogID)
	err := c.cc.Invoke(ctx, "/protos.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *BlogID, opts ...grpc.CallOption) (*Blog, error) {
	out := new(Blog)
	err := c.cc.Invoke(ctx, "/protos.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	CreateBlog(context.Context, *Blog) (*BlogID, error)
	ReadBlog(context.Context, *BlogID) (*Blog, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) CreateBlog(context.Context, *Blog) (*BlogID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogServiceServer) ReadBlog(context.Context, *BlogID) (*Blog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*Blog))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*BlogID))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/blog.proto",
}
