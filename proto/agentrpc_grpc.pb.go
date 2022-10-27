// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/agentrpc.proto

package proto

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

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	// Clone invokes MySQL CLONE command initializes the cloned database for MOCO.
	// It does _not_ start the replication (START REPLICA).  Actually, it works as follows.
	//
	// 1. Configure `clone_donor_valid_list` global variable to allow the donor instance.
	//
	// 2. Invoke `CLONE INSTANCE` with `user` and `password` in the CloneRequest.
	//
	// 3. Initialize the database for MOCO using `init_user` and `init_password`.
	//
	// For 2, the user must have BACKUP_ADMIN and REPLICATION SLAVE privilege.
	// For 3, the init_user must have ALL privilege with GRANT OPTION.
	// The init_user is used only via UNIX domain socket, so its host can be `localhost`.
	//
	// The donor database should have prepared these two users beforehand.
	Clone(ctx context.Context, in *CloneRequest, opts ...grpc.CallOption) (*CloneResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Clone(ctx context.Context, in *CloneRequest, opts ...grpc.CallOption) (*CloneResponse, error) {
	out := new(CloneResponse)
	err := c.cc.Invoke(ctx, "/moco.Agent/Clone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	// Clone invokes MySQL CLONE command initializes the cloned database for MOCO.
	// It does _not_ start the replication (START REPLICA).  Actually, it works as follows.
	//
	// 1. Configure `clone_donor_valid_list` global variable to allow the donor instance.
	//
	// 2. Invoke `CLONE INSTANCE` with `user` and `password` in the CloneRequest.
	//
	// 3. Initialize the database for MOCO using `init_user` and `init_password`.
	//
	// For 2, the user must have BACKUP_ADMIN and REPLICATION SLAVE privilege.
	// For 3, the init_user must have ALL privilege with GRANT OPTION.
	// The init_user is used only via UNIX domain socket, so its host can be `localhost`.
	//
	// The donor database should have prepared these two users beforehand.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clone not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_Clone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Clone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moco.Agent/Clone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Clone(ctx, req.(*CloneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moco.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Clone",
			Handler:    _Agent_Clone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/agentrpc.proto",
}
