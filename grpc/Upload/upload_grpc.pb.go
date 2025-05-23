// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: upload.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DFS_UploadPortsRequest_FullMethodName           = "/Upload.DFS/UploadPortsRequest"
	DFS_UploadFileRequest_FullMethodName            = "/Upload.DFS/UploadFileRequest"
	DFS_NodeMasterAckRequestUpload_FullMethodName   = "/Upload.DFS/NodeMasterAckRequestUpload"
	DFS_MasterClientAckRequestUpload_FullMethodName = "/Upload.DFS/MasterClientAckRequestUpload"
	DFS_NodeMasterAckRequest_FullMethodName         = "/Upload.DFS/NodeMasterAckRequest"
	DFS_MasterClientAckRequest_FullMethodName       = "/Upload.DFS/MasterClientAckRequest"
	DFS_DownloadPortsRequest_FullMethodName         = "/Upload.DFS/DownloadPortsRequest"
	DFS_DownloadFileRequest_FullMethodName          = "/Upload.DFS/DownloadFileRequest"
)

// DFSClient is the client API for DFS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DFSClient interface {
	// upload
	UploadPortsRequest(ctx context.Context, in *UploadRequestBody, opts ...grpc.CallOption) (*UploadResponseBody, error)
	UploadFileRequest(ctx context.Context, in *UploadFileRequestBody, opts ...grpc.CallOption) (*Empty, error)
	NodeMasterAckRequestUpload(ctx context.Context, in *NodeMasterAckRequestBodyUpload, opts ...grpc.CallOption) (*Empty, error)
	MasterClientAckRequestUpload(ctx context.Context, in *MasterClientAckRequestBodyUpload, opts ...grpc.CallOption) (*Empty, error)
	// download
	NodeMasterAckRequest(ctx context.Context, in *NodeMasterAckRequestBody, opts ...grpc.CallOption) (*Empty, error)
	MasterClientAckRequest(ctx context.Context, in *MasterClientAckRequestBody, opts ...grpc.CallOption) (*Empty, error)
	DownloadPortsRequest(ctx context.Context, in *DownloadPortsRequestBody, opts ...grpc.CallOption) (*DownloadPortsResponseBody, error)
	DownloadFileRequest(ctx context.Context, in *DownloadFileRequestBody, opts ...grpc.CallOption) (*DownloadFileResponseBody, error)
}

type dFSClient struct {
	cc grpc.ClientConnInterface
}

func NewDFSClient(cc grpc.ClientConnInterface) DFSClient {
	return &dFSClient{cc}
}

func (c *dFSClient) UploadPortsRequest(ctx context.Context, in *UploadRequestBody, opts ...grpc.CallOption) (*UploadResponseBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponseBody)
	err := c.cc.Invoke(ctx, DFS_UploadPortsRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) UploadFileRequest(ctx context.Context, in *UploadFileRequestBody, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, DFS_UploadFileRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) NodeMasterAckRequestUpload(ctx context.Context, in *NodeMasterAckRequestBodyUpload, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, DFS_NodeMasterAckRequestUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) MasterClientAckRequestUpload(ctx context.Context, in *MasterClientAckRequestBodyUpload, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, DFS_MasterClientAckRequestUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) NodeMasterAckRequest(ctx context.Context, in *NodeMasterAckRequestBody, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, DFS_NodeMasterAckRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) MasterClientAckRequest(ctx context.Context, in *MasterClientAckRequestBody, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, DFS_MasterClientAckRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) DownloadPortsRequest(ctx context.Context, in *DownloadPortsRequestBody, opts ...grpc.CallOption) (*DownloadPortsResponseBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadPortsResponseBody)
	err := c.cc.Invoke(ctx, DFS_DownloadPortsRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) DownloadFileRequest(ctx context.Context, in *DownloadFileRequestBody, opts ...grpc.CallOption) (*DownloadFileResponseBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadFileResponseBody)
	err := c.cc.Invoke(ctx, DFS_DownloadFileRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DFSServer is the server API for DFS service.
// All implementations must embed UnimplementedDFSServer
// for forward compatibility.
type DFSServer interface {
	// upload
	UploadPortsRequest(context.Context, *UploadRequestBody) (*UploadResponseBody, error)
	UploadFileRequest(context.Context, *UploadFileRequestBody) (*Empty, error)
	NodeMasterAckRequestUpload(context.Context, *NodeMasterAckRequestBodyUpload) (*Empty, error)
	MasterClientAckRequestUpload(context.Context, *MasterClientAckRequestBodyUpload) (*Empty, error)
	// download
	NodeMasterAckRequest(context.Context, *NodeMasterAckRequestBody) (*Empty, error)
	MasterClientAckRequest(context.Context, *MasterClientAckRequestBody) (*Empty, error)
	DownloadPortsRequest(context.Context, *DownloadPortsRequestBody) (*DownloadPortsResponseBody, error)
	DownloadFileRequest(context.Context, *DownloadFileRequestBody) (*DownloadFileResponseBody, error)
	mustEmbedUnimplementedDFSServer()
}

// UnimplementedDFSServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDFSServer struct{}

func (UnimplementedDFSServer) UploadPortsRequest(context.Context, *UploadRequestBody) (*UploadResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPortsRequest not implemented")
}
func (UnimplementedDFSServer) UploadFileRequest(context.Context, *UploadFileRequestBody) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFileRequest not implemented")
}
func (UnimplementedDFSServer) NodeMasterAckRequestUpload(context.Context, *NodeMasterAckRequestBodyUpload) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeMasterAckRequestUpload not implemented")
}
func (UnimplementedDFSServer) MasterClientAckRequestUpload(context.Context, *MasterClientAckRequestBodyUpload) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MasterClientAckRequestUpload not implemented")
}
func (UnimplementedDFSServer) NodeMasterAckRequest(context.Context, *NodeMasterAckRequestBody) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeMasterAckRequest not implemented")
}
func (UnimplementedDFSServer) MasterClientAckRequest(context.Context, *MasterClientAckRequestBody) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MasterClientAckRequest not implemented")
}
func (UnimplementedDFSServer) DownloadPortsRequest(context.Context, *DownloadPortsRequestBody) (*DownloadPortsResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadPortsRequest not implemented")
}
func (UnimplementedDFSServer) DownloadFileRequest(context.Context, *DownloadFileRequestBody) (*DownloadFileResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFileRequest not implemented")
}
func (UnimplementedDFSServer) mustEmbedUnimplementedDFSServer() {}
func (UnimplementedDFSServer) testEmbeddedByValue()             {}

// UnsafeDFSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DFSServer will
// result in compilation errors.
type UnsafeDFSServer interface {
	mustEmbedUnimplementedDFSServer()
}

func RegisterDFSServer(s grpc.ServiceRegistrar, srv DFSServer) {
	// If the following call pancis, it indicates UnimplementedDFSServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DFS_ServiceDesc, srv)
}

func _DFS_UploadPortsRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).UploadPortsRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DFS_UploadPortsRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).UploadPortsRequest(ctx, req.(*UploadRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_UploadFileRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).UploadFileRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DFS_UploadFileRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).UploadFileRequest(ctx, req.(*UploadFileRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_NodeMasterAckRequestUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeMasterAckRequestBodyUpload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).NodeMasterAckRequestUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DFS_NodeMasterAckRequestUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).NodeMasterAckRequestUpload(ctx, req.(*NodeMasterAckRequestBodyUpload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_MasterClientAckRequestUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterClientAckRequestBodyUpload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).MasterClientAckRequestUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DFS_MasterClientAckRequestUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).MasterClientAckRequestUpload(ctx, req.(*MasterClientAckRequestBodyUpload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_NodeMasterAckRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeMasterAckRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).NodeMasterAckRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DFS_NodeMasterAckRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).NodeMasterAckRequest(ctx, req.(*NodeMasterAckRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_MasterClientAckRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterClientAckRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).MasterClientAckRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DFS_MasterClientAckRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).MasterClientAckRequest(ctx, req.(*MasterClientAckRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_DownloadPortsRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadPortsRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).DownloadPortsRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DFS_DownloadPortsRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).DownloadPortsRequest(ctx, req.(*DownloadPortsRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_DownloadFileRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).DownloadFileRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DFS_DownloadFileRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).DownloadFileRequest(ctx, req.(*DownloadFileRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

// DFS_ServiceDesc is the grpc.ServiceDesc for DFS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DFS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Upload.DFS",
	HandlerType: (*DFSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadPortsRequest",
			Handler:    _DFS_UploadPortsRequest_Handler,
		},
		{
			MethodName: "UploadFileRequest",
			Handler:    _DFS_UploadFileRequest_Handler,
		},
		{
			MethodName: "NodeMasterAckRequestUpload",
			Handler:    _DFS_NodeMasterAckRequestUpload_Handler,
		},
		{
			MethodName: "MasterClientAckRequestUpload",
			Handler:    _DFS_MasterClientAckRequestUpload_Handler,
		},
		{
			MethodName: "NodeMasterAckRequest",
			Handler:    _DFS_NodeMasterAckRequest_Handler,
		},
		{
			MethodName: "MasterClientAckRequest",
			Handler:    _DFS_MasterClientAckRequest_Handler,
		},
		{
			MethodName: "DownloadPortsRequest",
			Handler:    _DFS_DownloadPortsRequest_Handler,
		},
		{
			MethodName: "DownloadFileRequest",
			Handler:    _DFS_DownloadFileRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload.proto",
}
