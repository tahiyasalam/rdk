// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// PoseTrackerServiceClient is the client API for PoseTrackerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoseTrackerServiceClient interface {
	// GetPoses returns the current pose of each body tracked by the pose tracker
	GetPoses(ctx context.Context, in *GetPosesRequest, opts ...grpc.CallOption) (*GetPosesResponse, error)
}

type poseTrackerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPoseTrackerServiceClient(cc grpc.ClientConnInterface) PoseTrackerServiceClient {
	return &poseTrackerServiceClient{cc}
}

func (c *poseTrackerServiceClient) GetPoses(ctx context.Context, in *GetPosesRequest, opts ...grpc.CallOption) (*GetPosesResponse, error) {
	out := new(GetPosesResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.posetracker.v1.PoseTrackerService/GetPoses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoseTrackerServiceServer is the server API for PoseTrackerService service.
// All implementations must embed UnimplementedPoseTrackerServiceServer
// for forward compatibility
type PoseTrackerServiceServer interface {
	// GetPoses returns the current pose of each body tracked by the pose tracker
	GetPoses(context.Context, *GetPosesRequest) (*GetPosesResponse, error)
	mustEmbedUnimplementedPoseTrackerServiceServer()
}

// UnimplementedPoseTrackerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPoseTrackerServiceServer struct {
}

func (UnimplementedPoseTrackerServiceServer) GetPoses(context.Context, *GetPosesRequest) (*GetPosesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoses not implemented")
}
func (UnimplementedPoseTrackerServiceServer) mustEmbedUnimplementedPoseTrackerServiceServer() {}

// UnsafePoseTrackerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoseTrackerServiceServer will
// result in compilation errors.
type UnsafePoseTrackerServiceServer interface {
	mustEmbedUnimplementedPoseTrackerServiceServer()
}

func RegisterPoseTrackerServiceServer(s grpc.ServiceRegistrar, srv PoseTrackerServiceServer) {
	s.RegisterService(&PoseTrackerService_ServiceDesc, srv)
}

func _PoseTrackerService_GetPoses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPosesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoseTrackerServiceServer).GetPoses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.posetracker.v1.PoseTrackerService/GetPoses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoseTrackerServiceServer).GetPoses(ctx, req.(*GetPosesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoseTrackerService_ServiceDesc is the grpc.ServiceDesc for PoseTrackerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoseTrackerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api.component.posetracker.v1.PoseTrackerService",
	HandlerType: (*PoseTrackerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPoses",
			Handler:    _PoseTrackerService_GetPoses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/component/posetracker/v1/pose_tracker.proto",
}
