// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: corndogs/v1alpha1/corndogs.proto

package corndogsv1alpha1

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

// CorndogsServiceClient is the client API for CorndogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CorndogsServiceClient interface {
	SubmitTask(ctx context.Context, in *SubmitTaskRequest, opts ...grpc.CallOption) (*SubmitTaskResponse, error)
	GetTaskStateByID(ctx context.Context, in *GetTaskStateByIDRequest, opts ...grpc.CallOption) (*GetTaskStateByIDResponse, error)
	GetNextTask(ctx context.Context, in *GetNextTaskRequest, opts ...grpc.CallOption) (*GetNextTaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error)
	CancelTask(ctx context.Context, in *CancelTaskRequest, opts ...grpc.CallOption) (*CancelTaskResponse, error)
	CleanUpTimedOut(ctx context.Context, in *CleanUpTimedOutRequest, opts ...grpc.CallOption) (*CleanUpTimedOutResponse, error)
	// Metrics
	GetQueues(ctx context.Context, in *GetQueuesRequest, opts ...grpc.CallOption) (*GetQueuesResponse, error)
	GetQueueTaskCounts(ctx context.Context, in *GetQueueTaskCountsRequest, opts ...grpc.CallOption) (*GetQueueTaskCountsResponse, error)
	GetTaskStateCounts(ctx context.Context, in *GetTaskStateCountsRequest, opts ...grpc.CallOption) (*GetTaskStateCountsResponse, error)
	GetQueueAndStateCounts(ctx context.Context, in *GetQueueAndStateCountsRequest, opts ...grpc.CallOption) (*GetQueueAndStateCountsResponse, error)
}

type corndogsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCorndogsServiceClient(cc grpc.ClientConnInterface) CorndogsServiceClient {
	return &corndogsServiceClient{cc}
}

func (c *corndogsServiceClient) SubmitTask(ctx context.Context, in *SubmitTaskRequest, opts ...grpc.CallOption) (*SubmitTaskResponse, error) {
	out := new(SubmitTaskResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/SubmitTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) GetTaskStateByID(ctx context.Context, in *GetTaskStateByIDRequest, opts ...grpc.CallOption) (*GetTaskStateByIDResponse, error) {
	out := new(GetTaskStateByIDResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/GetTaskStateByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) GetNextTask(ctx context.Context, in *GetNextTaskRequest, opts ...grpc.CallOption) (*GetNextTaskResponse, error) {
	out := new(GetNextTaskResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/GetNextTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error) {
	out := new(CompleteTaskResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/CompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) CancelTask(ctx context.Context, in *CancelTaskRequest, opts ...grpc.CallOption) (*CancelTaskResponse, error) {
	out := new(CancelTaskResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/CancelTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) CleanUpTimedOut(ctx context.Context, in *CleanUpTimedOutRequest, opts ...grpc.CallOption) (*CleanUpTimedOutResponse, error) {
	out := new(CleanUpTimedOutResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/CleanUpTimedOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) GetQueues(ctx context.Context, in *GetQueuesRequest, opts ...grpc.CallOption) (*GetQueuesResponse, error) {
	out := new(GetQueuesResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/GetQueues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) GetQueueTaskCounts(ctx context.Context, in *GetQueueTaskCountsRequest, opts ...grpc.CallOption) (*GetQueueTaskCountsResponse, error) {
	out := new(GetQueueTaskCountsResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/GetQueueTaskCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) GetTaskStateCounts(ctx context.Context, in *GetTaskStateCountsRequest, opts ...grpc.CallOption) (*GetTaskStateCountsResponse, error) {
	out := new(GetTaskStateCountsResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/GetTaskStateCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corndogsServiceClient) GetQueueAndStateCounts(ctx context.Context, in *GetQueueAndStateCountsRequest, opts ...grpc.CallOption) (*GetQueueAndStateCountsResponse, error) {
	out := new(GetQueueAndStateCountsResponse)
	err := c.cc.Invoke(ctx, "/corndogs.v1alpha1.CorndogsService/GetQueueAndStateCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CorndogsServiceServer is the server API for CorndogsService service.
// All implementations should embed UnimplementedCorndogsServiceServer
// for forward compatibility
type CorndogsServiceServer interface {
	SubmitTask(context.Context, *SubmitTaskRequest) (*SubmitTaskResponse, error)
	GetTaskStateByID(context.Context, *GetTaskStateByIDRequest) (*GetTaskStateByIDResponse, error)
	GetNextTask(context.Context, *GetNextTaskRequest) (*GetNextTaskResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error)
	CancelTask(context.Context, *CancelTaskRequest) (*CancelTaskResponse, error)
	CleanUpTimedOut(context.Context, *CleanUpTimedOutRequest) (*CleanUpTimedOutResponse, error)
	// Metrics
	GetQueues(context.Context, *GetQueuesRequest) (*GetQueuesResponse, error)
	GetQueueTaskCounts(context.Context, *GetQueueTaskCountsRequest) (*GetQueueTaskCountsResponse, error)
	GetTaskStateCounts(context.Context, *GetTaskStateCountsRequest) (*GetTaskStateCountsResponse, error)
	GetQueueAndStateCounts(context.Context, *GetQueueAndStateCountsRequest) (*GetQueueAndStateCountsResponse, error)
}

// UnimplementedCorndogsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCorndogsServiceServer struct {
}

func (UnimplementedCorndogsServiceServer) SubmitTask(context.Context, *SubmitTaskRequest) (*SubmitTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTask not implemented")
}
func (UnimplementedCorndogsServiceServer) GetTaskStateByID(context.Context, *GetTaskStateByIDRequest) (*GetTaskStateByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskStateByID not implemented")
}
func (UnimplementedCorndogsServiceServer) GetNextTask(context.Context, *GetNextTaskRequest) (*GetNextTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextTask not implemented")
}
func (UnimplementedCorndogsServiceServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedCorndogsServiceServer) CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTask not implemented")
}
func (UnimplementedCorndogsServiceServer) CancelTask(context.Context, *CancelTaskRequest) (*CancelTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTask not implemented")
}
func (UnimplementedCorndogsServiceServer) CleanUpTimedOut(context.Context, *CleanUpTimedOutRequest) (*CleanUpTimedOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanUpTimedOut not implemented")
}
func (UnimplementedCorndogsServiceServer) GetQueues(context.Context, *GetQueuesRequest) (*GetQueuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueues not implemented")
}
func (UnimplementedCorndogsServiceServer) GetQueueTaskCounts(context.Context, *GetQueueTaskCountsRequest) (*GetQueueTaskCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueueTaskCounts not implemented")
}
func (UnimplementedCorndogsServiceServer) GetTaskStateCounts(context.Context, *GetTaskStateCountsRequest) (*GetTaskStateCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskStateCounts not implemented")
}
func (UnimplementedCorndogsServiceServer) GetQueueAndStateCounts(context.Context, *GetQueueAndStateCountsRequest) (*GetQueueAndStateCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueueAndStateCounts not implemented")
}

// UnsafeCorndogsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CorndogsServiceServer will
// result in compilation errors.
type UnsafeCorndogsServiceServer interface {
	mustEmbedUnimplementedCorndogsServiceServer()
}

func RegisterCorndogsServiceServer(s grpc.ServiceRegistrar, srv CorndogsServiceServer) {
	s.RegisterService(&CorndogsService_ServiceDesc, srv)
}

func _CorndogsService_SubmitTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).SubmitTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/SubmitTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).SubmitTask(ctx, req.(*SubmitTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_GetTaskStateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskStateByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).GetTaskStateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/GetTaskStateByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).GetTaskStateByID(ctx, req.(*GetTaskStateByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_GetNextTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNextTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).GetNextTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/GetNextTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).GetNextTask(ctx, req.(*GetNextTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/CompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).CompleteTask(ctx, req.(*CompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).CancelTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/CancelTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).CancelTask(ctx, req.(*CancelTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_CleanUpTimedOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanUpTimedOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).CleanUpTimedOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/CleanUpTimedOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).CleanUpTimedOut(ctx, req.(*CleanUpTimedOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_GetQueues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).GetQueues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/GetQueues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).GetQueues(ctx, req.(*GetQueuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_GetQueueTaskCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueueTaskCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).GetQueueTaskCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/GetQueueTaskCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).GetQueueTaskCounts(ctx, req.(*GetQueueTaskCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_GetTaskStateCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskStateCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).GetTaskStateCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/GetTaskStateCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).GetTaskStateCounts(ctx, req.(*GetTaskStateCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CorndogsService_GetQueueAndStateCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueueAndStateCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorndogsServiceServer).GetQueueAndStateCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/corndogs.v1alpha1.CorndogsService/GetQueueAndStateCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorndogsServiceServer).GetQueueAndStateCounts(ctx, req.(*GetQueueAndStateCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CorndogsService_ServiceDesc is the grpc.ServiceDesc for CorndogsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CorndogsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "corndogs.v1alpha1.CorndogsService",
	HandlerType: (*CorndogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTask",
			Handler:    _CorndogsService_SubmitTask_Handler,
		},
		{
			MethodName: "GetTaskStateByID",
			Handler:    _CorndogsService_GetTaskStateByID_Handler,
		},
		{
			MethodName: "GetNextTask",
			Handler:    _CorndogsService_GetNextTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _CorndogsService_UpdateTask_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _CorndogsService_CompleteTask_Handler,
		},
		{
			MethodName: "CancelTask",
			Handler:    _CorndogsService_CancelTask_Handler,
		},
		{
			MethodName: "CleanUpTimedOut",
			Handler:    _CorndogsService_CleanUpTimedOut_Handler,
		},
		{
			MethodName: "GetQueues",
			Handler:    _CorndogsService_GetQueues_Handler,
		},
		{
			MethodName: "GetQueueTaskCounts",
			Handler:    _CorndogsService_GetQueueTaskCounts_Handler,
		},
		{
			MethodName: "GetTaskStateCounts",
			Handler:    _CorndogsService_GetTaskStateCounts_Handler,
		},
		{
			MethodName: "GetQueueAndStateCounts",
			Handler:    _CorndogsService_GetQueueAndStateCounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "corndogs/v1alpha1/corndogs.proto",
}
