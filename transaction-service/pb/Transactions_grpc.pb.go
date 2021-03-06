// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransactionServiceRPCClient is the client API for TransactionServiceRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceRPCClient interface {
	CreditAccount(ctx context.Context, in *CreditAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DebitAccount(ctx context.Context, in *DebitAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type transactionServiceRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceRPCClient(cc grpc.ClientConnInterface) TransactionServiceRPCClient {
	return &transactionServiceRPCClient{cc}
}

func (c *transactionServiceRPCClient) CreditAccount(ctx context.Context, in *CreditAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/transaction.TransactionServiceRPC/CreditAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceRPCClient) DebitAccount(ctx context.Context, in *DebitAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/transaction.TransactionServiceRPC/DebitAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceRPCServer is the server API for TransactionServiceRPC service.
// All implementations should embed UnimplementedTransactionServiceRPCServer
// for forward compatibility
type TransactionServiceRPCServer interface {
	CreditAccount(context.Context, *CreditAccountRequest) (*emptypb.Empty, error)
	DebitAccount(context.Context, *DebitAccountRequest) (*emptypb.Empty, error)
}

// UnimplementedTransactionServiceRPCServer should be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceRPCServer struct {
}

func (UnimplementedTransactionServiceRPCServer) CreditAccount(context.Context, *CreditAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditAccount not implemented")
}
func (UnimplementedTransactionServiceRPCServer) DebitAccount(context.Context, *DebitAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebitAccount not implemented")
}

// UnsafeTransactionServiceRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceRPCServer will
// result in compilation errors.
type UnsafeTransactionServiceRPCServer interface {
	mustEmbedUnimplementedTransactionServiceRPCServer()
}

func RegisterTransactionServiceRPCServer(s grpc.ServiceRegistrar, srv TransactionServiceRPCServer) {
	s.RegisterService(&TransactionServiceRPC_ServiceDesc, srv)
}

func _TransactionServiceRPC_CreditAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceRPCServer).CreditAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionServiceRPC/CreditAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceRPCServer).CreditAccount(ctx, req.(*CreditAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionServiceRPC_DebitAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceRPCServer).DebitAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionServiceRPC/DebitAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceRPCServer).DebitAccount(ctx, req.(*DebitAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionServiceRPC_ServiceDesc is the grpc.ServiceDesc for TransactionServiceRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionServiceRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TransactionServiceRPC",
	HandlerType: (*TransactionServiceRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreditAccount",
			Handler:    _TransactionServiceRPC_CreditAccount_Handler,
		},
		{
			MethodName: "DebitAccount",
			Handler:    _TransactionServiceRPC_DebitAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction-service/pb/Transactions.proto",
}
