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

// MotorServiceClient is the client API for MotorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MotorServiceClient interface {
	// SetPower sets the percentage of the motor's total power that should be employed
	// expressed a value between -1 and 1 where negative values indiciate a backwards
	// direction and positive values a forward direction
	SetPower(ctx context.Context, in *MotorServiceSetPowerRequest, opts ...grpc.CallOption) (*MotorServiceSetPowerResponse, error)
	// Go instructs the motor to turn using a specified percentage of its total power,
	// expressed as a value between -1 and 1 where negative values indiciate a backwards
	// direction and positive values a forward direction
	Go(ctx context.Context, in *MotorServiceGoRequest, opts ...grpc.CallOption) (*MotorServiceGoResponse, error)
	// GoFor instructs the motor to turn at a specified speed, which is expressed in RPM,
	// for a specified number of rotations relative to its starting position
	// This method will return an error if MotorPositionSupported is false
	GoFor(ctx context.Context, in *MotorServiceGoForRequest, opts ...grpc.CallOption) (*MotorServiceGoForResponse, error)
	// GoTo requests the robot's motor to move to a specific position that
	// is relative to its home position at a specified speed which is expressed in RPM
	// This method will return an error if MotorPositionSupported is false
	GoTo(ctx context.Context, in *MotorServiceGoToRequest, opts ...grpc.CallOption) (*MotorServiceGoToResponse, error)
	// To Do (FA): This will be deprecated in favor of a  MotorStop method
	// GoTillStop moves a motor until it is stopped
	// The logic to trigger the "stop" mechanism is up to the underlying motor implementation
	GoTillStop(ctx context.Context, in *MotorServiceGoTillStopRequest, opts ...grpc.CallOption) (*MotorServiceGoTillStopResponse, error)
	// ResetZeroPosition sets the current position of the motor as the new zero position
	// This method will return an error if MotorPositionSupported is false
	ResetZeroPosition(ctx context.Context, in *MotorServiceResetZeroPositionRequest, opts ...grpc.CallOption) (*MotorServiceResetZeroPositionResponse, error)
	// Position reports the position of the robot's motor relative to its zero position
	// This method will return an error if MotorPositionSupported is false
	Position(ctx context.Context, in *MotorServicePositionRequest, opts ...grpc.CallOption) (*MotorServicePositionResponse, error)
	// PositionSupported returns whether or not the robot's motor supports reporting of its position
	PositionSupported(ctx context.Context, in *MotorServicePositionSupportedRequest, opts ...grpc.CallOption) (*MotorServicePositionSupportedResponse, error)
	// Stop turns the robot's motor off
	// To Do (FA): This will be deprecated
	Stop(ctx context.Context, in *MotorServiceStopRequest, opts ...grpc.CallOption) (*MotorServiceStopResponse, error)
	// IsOn returns true if the robot's motor off
	// To Do (FA): This will be deprecated
	IsOn(ctx context.Context, in *MotorServiceIsOnRequest, opts ...grpc.CallOption) (*MotorServiceIsOnResponse, error)
}

type motorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMotorServiceClient(cc grpc.ClientConnInterface) MotorServiceClient {
	return &motorServiceClient{cc}
}

func (c *motorServiceClient) SetPower(ctx context.Context, in *MotorServiceSetPowerRequest, opts ...grpc.CallOption) (*MotorServiceSetPowerResponse, error) {
	out := new(MotorServiceSetPowerResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/SetPower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) Go(ctx context.Context, in *MotorServiceGoRequest, opts ...grpc.CallOption) (*MotorServiceGoResponse, error) {
	out := new(MotorServiceGoResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/Go", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) GoFor(ctx context.Context, in *MotorServiceGoForRequest, opts ...grpc.CallOption) (*MotorServiceGoForResponse, error) {
	out := new(MotorServiceGoForResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/GoFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) GoTo(ctx context.Context, in *MotorServiceGoToRequest, opts ...grpc.CallOption) (*MotorServiceGoToResponse, error) {
	out := new(MotorServiceGoToResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/GoTo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) GoTillStop(ctx context.Context, in *MotorServiceGoTillStopRequest, opts ...grpc.CallOption) (*MotorServiceGoTillStopResponse, error) {
	out := new(MotorServiceGoTillStopResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/GoTillStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) ResetZeroPosition(ctx context.Context, in *MotorServiceResetZeroPositionRequest, opts ...grpc.CallOption) (*MotorServiceResetZeroPositionResponse, error) {
	out := new(MotorServiceResetZeroPositionResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/ResetZeroPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) Position(ctx context.Context, in *MotorServicePositionRequest, opts ...grpc.CallOption) (*MotorServicePositionResponse, error) {
	out := new(MotorServicePositionResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/Position", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) PositionSupported(ctx context.Context, in *MotorServicePositionSupportedRequest, opts ...grpc.CallOption) (*MotorServicePositionSupportedResponse, error) {
	out := new(MotorServicePositionSupportedResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/PositionSupported", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) Stop(ctx context.Context, in *MotorServiceStopRequest, opts ...grpc.CallOption) (*MotorServiceStopResponse, error) {
	out := new(MotorServiceStopResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorServiceClient) IsOn(ctx context.Context, in *MotorServiceIsOnRequest, opts ...grpc.CallOption) (*MotorServiceIsOnResponse, error) {
	out := new(MotorServiceIsOnResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.v1.MotorService/IsOn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MotorServiceServer is the server API for MotorService service.
// All implementations must embed UnimplementedMotorServiceServer
// for forward compatibility
type MotorServiceServer interface {
	// SetPower sets the percentage of the motor's total power that should be employed
	// expressed a value between -1 and 1 where negative values indiciate a backwards
	// direction and positive values a forward direction
	SetPower(context.Context, *MotorServiceSetPowerRequest) (*MotorServiceSetPowerResponse, error)
	// Go instructs the motor to turn using a specified percentage of its total power,
	// expressed as a value between -1 and 1 where negative values indiciate a backwards
	// direction and positive values a forward direction
	Go(context.Context, *MotorServiceGoRequest) (*MotorServiceGoResponse, error)
	// GoFor instructs the motor to turn at a specified speed, which is expressed in RPM,
	// for a specified number of rotations relative to its starting position
	// This method will return an error if MotorPositionSupported is false
	GoFor(context.Context, *MotorServiceGoForRequest) (*MotorServiceGoForResponse, error)
	// GoTo requests the robot's motor to move to a specific position that
	// is relative to its home position at a specified speed which is expressed in RPM
	// This method will return an error if MotorPositionSupported is false
	GoTo(context.Context, *MotorServiceGoToRequest) (*MotorServiceGoToResponse, error)
	// To Do (FA): This will be deprecated in favor of a  MotorStop method
	// GoTillStop moves a motor until it is stopped
	// The logic to trigger the "stop" mechanism is up to the underlying motor implementation
	GoTillStop(context.Context, *MotorServiceGoTillStopRequest) (*MotorServiceGoTillStopResponse, error)
	// ResetZeroPosition sets the current position of the motor as the new zero position
	// This method will return an error if MotorPositionSupported is false
	ResetZeroPosition(context.Context, *MotorServiceResetZeroPositionRequest) (*MotorServiceResetZeroPositionResponse, error)
	// Position reports the position of the robot's motor relative to its zero position
	// This method will return an error if MotorPositionSupported is false
	Position(context.Context, *MotorServicePositionRequest) (*MotorServicePositionResponse, error)
	// PositionSupported returns whether or not the robot's motor supports reporting of its position
	PositionSupported(context.Context, *MotorServicePositionSupportedRequest) (*MotorServicePositionSupportedResponse, error)
	// Stop turns the robot's motor off
	// To Do (FA): This will be deprecated
	Stop(context.Context, *MotorServiceStopRequest) (*MotorServiceStopResponse, error)
	// IsOn returns true if the robot's motor off
	// To Do (FA): This will be deprecated
	IsOn(context.Context, *MotorServiceIsOnRequest) (*MotorServiceIsOnResponse, error)
	mustEmbedUnimplementedMotorServiceServer()
}

// UnimplementedMotorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMotorServiceServer struct {
}

func (UnimplementedMotorServiceServer) SetPower(context.Context, *MotorServiceSetPowerRequest) (*MotorServiceSetPowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPower not implemented")
}
func (UnimplementedMotorServiceServer) Go(context.Context, *MotorServiceGoRequest) (*MotorServiceGoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Go not implemented")
}
func (UnimplementedMotorServiceServer) GoFor(context.Context, *MotorServiceGoForRequest) (*MotorServiceGoForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoFor not implemented")
}
func (UnimplementedMotorServiceServer) GoTo(context.Context, *MotorServiceGoToRequest) (*MotorServiceGoToResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoTo not implemented")
}
func (UnimplementedMotorServiceServer) GoTillStop(context.Context, *MotorServiceGoTillStopRequest) (*MotorServiceGoTillStopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoTillStop not implemented")
}
func (UnimplementedMotorServiceServer) ResetZeroPosition(context.Context, *MotorServiceResetZeroPositionRequest) (*MotorServiceResetZeroPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetZeroPosition not implemented")
}
func (UnimplementedMotorServiceServer) Position(context.Context, *MotorServicePositionRequest) (*MotorServicePositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Position not implemented")
}
func (UnimplementedMotorServiceServer) PositionSupported(context.Context, *MotorServicePositionSupportedRequest) (*MotorServicePositionSupportedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PositionSupported not implemented")
}
func (UnimplementedMotorServiceServer) Stop(context.Context, *MotorServiceStopRequest) (*MotorServiceStopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedMotorServiceServer) IsOn(context.Context, *MotorServiceIsOnRequest) (*MotorServiceIsOnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsOn not implemented")
}
func (UnimplementedMotorServiceServer) mustEmbedUnimplementedMotorServiceServer() {}

// UnsafeMotorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MotorServiceServer will
// result in compilation errors.
type UnsafeMotorServiceServer interface {
	mustEmbedUnimplementedMotorServiceServer()
}

func RegisterMotorServiceServer(s grpc.ServiceRegistrar, srv MotorServiceServer) {
	s.RegisterService(&MotorService_ServiceDesc, srv)
}

func _MotorService_SetPower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServiceSetPowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).SetPower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/SetPower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).SetPower(ctx, req.(*MotorServiceSetPowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_Go_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServiceGoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).Go(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/Go",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).Go(ctx, req.(*MotorServiceGoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_GoFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServiceGoForRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).GoFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/GoFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).GoFor(ctx, req.(*MotorServiceGoForRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_GoTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServiceGoToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).GoTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/GoTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).GoTo(ctx, req.(*MotorServiceGoToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_GoTillStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServiceGoTillStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).GoTillStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/GoTillStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).GoTillStop(ctx, req.(*MotorServiceGoTillStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_ResetZeroPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServiceResetZeroPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).ResetZeroPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/ResetZeroPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).ResetZeroPosition(ctx, req.(*MotorServiceResetZeroPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_Position_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServicePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).Position(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/Position",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).Position(ctx, req.(*MotorServicePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_PositionSupported_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServicePositionSupportedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).PositionSupported(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/PositionSupported",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).PositionSupported(ctx, req.(*MotorServicePositionSupportedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServiceStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).Stop(ctx, req.(*MotorServiceStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorService_IsOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorServiceIsOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorServiceServer).IsOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.v1.MotorService/IsOn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorServiceServer).IsOn(ctx, req.(*MotorServiceIsOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MotorService_ServiceDesc is the grpc.ServiceDesc for MotorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MotorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api.component.v1.MotorService",
	HandlerType: (*MotorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPower",
			Handler:    _MotorService_SetPower_Handler,
		},
		{
			MethodName: "Go",
			Handler:    _MotorService_Go_Handler,
		},
		{
			MethodName: "GoFor",
			Handler:    _MotorService_GoFor_Handler,
		},
		{
			MethodName: "GoTo",
			Handler:    _MotorService_GoTo_Handler,
		},
		{
			MethodName: "GoTillStop",
			Handler:    _MotorService_GoTillStop_Handler,
		},
		{
			MethodName: "ResetZeroPosition",
			Handler:    _MotorService_ResetZeroPosition_Handler,
		},
		{
			MethodName: "Position",
			Handler:    _MotorService_Position_Handler,
		},
		{
			MethodName: "PositionSupported",
			Handler:    _MotorService_PositionSupported_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _MotorService_Stop_Handler,
		},
		{
			MethodName: "IsOn",
			Handler:    _MotorService_IsOn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/component/v1/motor.proto",
}
