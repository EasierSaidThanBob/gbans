// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	SendLog(ctx context.Context, opts ...grpc.CallOption) (Agent_SendLogClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) SendLog(ctx context.Context, opts ...grpc.CallOption) (Agent_SendLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/relay.Agent/SendLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentSendLogClient{stream}
	return x, nil
}

type Agent_SendLogClient interface {
	Send(*LogEntry) error
	CloseAndRecv() (*SendLogSummary, error)
	grpc.ClientStream
}

type agentSendLogClient struct {
	grpc.ClientStream
}

func (x *agentSendLogClient) Send(m *LogEntry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentSendLogClient) CloseAndRecv() (*SendLogSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendLogSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	SendLog(Agent_SendLogServer) error
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) SendLog(Agent_SendLogServer) error {
	return status.Errorf(codes.Unimplemented, "method SendLog not implemented")
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

func _Agent_SendLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).SendLog(&agentSendLogServer{stream})
}

type Agent_SendLogServer interface {
	SendAndClose(*SendLogSummary) error
	Recv() (*LogEntry, error)
	grpc.ServerStream
}

type agentSendLogServer struct {
	grpc.ServerStream
}

func (x *agentSendLogServer) SendAndClose(m *SendLogSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentSendLogServer) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relay.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendLog",
			Handler:       _Agent_SendLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/server_event.proto",
}
