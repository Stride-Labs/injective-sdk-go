// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: injective_accounts_rpc.proto

package injective_accounts_rpcpb

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

// InjectiveAccountsRPCClient is the client API for InjectiveAccountsRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InjectiveAccountsRPCClient interface {
	// Provide the account's portfolio value in USD.
	Portfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error)
	// List order states by order hashes
	OrderStates(ctx context.Context, in *OrderStatesRequest, opts ...grpc.CallOption) (*OrderStatesResponse, error)
	// List all subaccounts IDs of an account address
	SubaccountsList(ctx context.Context, in *SubaccountsListRequest, opts ...grpc.CallOption) (*SubaccountsListResponse, error)
	// List subaccount balances for the provided denoms.
	SubaccountBalancesList(ctx context.Context, in *SubaccountBalancesListRequest, opts ...grpc.CallOption) (*SubaccountBalancesListResponse, error)
	// Gets a balance for specific coin denom
	SubaccountBalanceEndpoint(ctx context.Context, in *SubaccountBalanceRequest, opts ...grpc.CallOption) (*SubaccountBalanceResponse, error)
	// StreamSubaccountBalance streams new balance changes for a specified
	// subaccount and denoms. If no denoms are provided, all denom changes are
	// streamed.
	StreamSubaccountBalance(ctx context.Context, in *StreamSubaccountBalanceRequest, opts ...grpc.CallOption) (InjectiveAccountsRPC_StreamSubaccountBalanceClient, error)
	// Get subaccount's deposits and withdrawals history
	SubaccountHistory(ctx context.Context, in *SubaccountHistoryRequest, opts ...grpc.CallOption) (*SubaccountHistoryResponse, error)
	// Get subaccount's orders summary
	SubaccountOrderSummary(ctx context.Context, in *SubaccountOrderSummaryRequest, opts ...grpc.CallOption) (*SubaccountOrderSummaryResponse, error)
	// Provide historical trading rewards
	Rewards(ctx context.Context, in *RewardsRequest, opts ...grpc.CallOption) (*RewardsResponse, error)
}

type injectiveAccountsRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewInjectiveAccountsRPCClient(cc grpc.ClientConnInterface) InjectiveAccountsRPCClient {
	return &injectiveAccountsRPCClient{cc}
}

func (c *injectiveAccountsRPCClient) Portfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error) {
	out := new(PortfolioResponse)
	err := c.cc.Invoke(ctx, "/injective_accounts_rpc.InjectiveAccountsRPC/Portfolio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAccountsRPCClient) OrderStates(ctx context.Context, in *OrderStatesRequest, opts ...grpc.CallOption) (*OrderStatesResponse, error) {
	out := new(OrderStatesResponse)
	err := c.cc.Invoke(ctx, "/injective_accounts_rpc.InjectiveAccountsRPC/OrderStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAccountsRPCClient) SubaccountsList(ctx context.Context, in *SubaccountsListRequest, opts ...grpc.CallOption) (*SubaccountsListResponse, error) {
	out := new(SubaccountsListResponse)
	err := c.cc.Invoke(ctx, "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAccountsRPCClient) SubaccountBalancesList(ctx context.Context, in *SubaccountBalancesListRequest, opts ...grpc.CallOption) (*SubaccountBalancesListResponse, error) {
	out := new(SubaccountBalancesListResponse)
	err := c.cc.Invoke(ctx, "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountBalancesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAccountsRPCClient) SubaccountBalanceEndpoint(ctx context.Context, in *SubaccountBalanceRequest, opts ...grpc.CallOption) (*SubaccountBalanceResponse, error) {
	out := new(SubaccountBalanceResponse)
	err := c.cc.Invoke(ctx, "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountBalanceEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAccountsRPCClient) StreamSubaccountBalance(ctx context.Context, in *StreamSubaccountBalanceRequest, opts ...grpc.CallOption) (InjectiveAccountsRPC_StreamSubaccountBalanceClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveAccountsRPC_ServiceDesc.Streams[0], "/injective_accounts_rpc.InjectiveAccountsRPC/StreamSubaccountBalance", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveAccountsRPCStreamSubaccountBalanceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveAccountsRPC_StreamSubaccountBalanceClient interface {
	Recv() (*StreamSubaccountBalanceResponse, error)
	grpc.ClientStream
}

type injectiveAccountsRPCStreamSubaccountBalanceClient struct {
	grpc.ClientStream
}

func (x *injectiveAccountsRPCStreamSubaccountBalanceClient) Recv() (*StreamSubaccountBalanceResponse, error) {
	m := new(StreamSubaccountBalanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveAccountsRPCClient) SubaccountHistory(ctx context.Context, in *SubaccountHistoryRequest, opts ...grpc.CallOption) (*SubaccountHistoryResponse, error) {
	out := new(SubaccountHistoryResponse)
	err := c.cc.Invoke(ctx, "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAccountsRPCClient) SubaccountOrderSummary(ctx context.Context, in *SubaccountOrderSummaryRequest, opts ...grpc.CallOption) (*SubaccountOrderSummaryResponse, error) {
	out := new(SubaccountOrderSummaryResponse)
	err := c.cc.Invoke(ctx, "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountOrderSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAccountsRPCClient) Rewards(ctx context.Context, in *RewardsRequest, opts ...grpc.CallOption) (*RewardsResponse, error) {
	out := new(RewardsResponse)
	err := c.cc.Invoke(ctx, "/injective_accounts_rpc.InjectiveAccountsRPC/Rewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjectiveAccountsRPCServer is the server API for InjectiveAccountsRPC service.
// All implementations must embed UnimplementedInjectiveAccountsRPCServer
// for forward compatibility
type InjectiveAccountsRPCServer interface {
	// Provide the account's portfolio value in USD.
	Portfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error)
	// List order states by order hashes
	OrderStates(context.Context, *OrderStatesRequest) (*OrderStatesResponse, error)
	// List all subaccounts IDs of an account address
	SubaccountsList(context.Context, *SubaccountsListRequest) (*SubaccountsListResponse, error)
	// List subaccount balances for the provided denoms.
	SubaccountBalancesList(context.Context, *SubaccountBalancesListRequest) (*SubaccountBalancesListResponse, error)
	// Gets a balance for specific coin denom
	SubaccountBalanceEndpoint(context.Context, *SubaccountBalanceRequest) (*SubaccountBalanceResponse, error)
	// StreamSubaccountBalance streams new balance changes for a specified
	// subaccount and denoms. If no denoms are provided, all denom changes are
	// streamed.
	StreamSubaccountBalance(*StreamSubaccountBalanceRequest, InjectiveAccountsRPC_StreamSubaccountBalanceServer) error
	// Get subaccount's deposits and withdrawals history
	SubaccountHistory(context.Context, *SubaccountHistoryRequest) (*SubaccountHistoryResponse, error)
	// Get subaccount's orders summary
	SubaccountOrderSummary(context.Context, *SubaccountOrderSummaryRequest) (*SubaccountOrderSummaryResponse, error)
	// Provide historical trading rewards
	Rewards(context.Context, *RewardsRequest) (*RewardsResponse, error)
	mustEmbedUnimplementedInjectiveAccountsRPCServer()
}

// UnimplementedInjectiveAccountsRPCServer must be embedded to have forward compatible implementations.
type UnimplementedInjectiveAccountsRPCServer struct {
}

func (UnimplementedInjectiveAccountsRPCServer) Portfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Portfolio not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) OrderStates(context.Context, *OrderStatesRequest) (*OrderStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderStates not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) SubaccountsList(context.Context, *SubaccountsListRequest) (*SubaccountsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountsList not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) SubaccountBalancesList(context.Context, *SubaccountBalancesListRequest) (*SubaccountBalancesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountBalancesList not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) SubaccountBalanceEndpoint(context.Context, *SubaccountBalanceRequest) (*SubaccountBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountBalanceEndpoint not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) StreamSubaccountBalance(*StreamSubaccountBalanceRequest, InjectiveAccountsRPC_StreamSubaccountBalanceServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSubaccountBalance not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) SubaccountHistory(context.Context, *SubaccountHistoryRequest) (*SubaccountHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountHistory not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) SubaccountOrderSummary(context.Context, *SubaccountOrderSummaryRequest) (*SubaccountOrderSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountOrderSummary not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) Rewards(context.Context, *RewardsRequest) (*RewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rewards not implemented")
}
func (UnimplementedInjectiveAccountsRPCServer) mustEmbedUnimplementedInjectiveAccountsRPCServer() {}

// UnsafeInjectiveAccountsRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InjectiveAccountsRPCServer will
// result in compilation errors.
type UnsafeInjectiveAccountsRPCServer interface {
	mustEmbedUnimplementedInjectiveAccountsRPCServer()
}

func RegisterInjectiveAccountsRPCServer(s grpc.ServiceRegistrar, srv InjectiveAccountsRPCServer) {
	s.RegisterService(&InjectiveAccountsRPC_ServiceDesc, srv)
}

func _InjectiveAccountsRPC_Portfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAccountsRPCServer).Portfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_accounts_rpc.InjectiveAccountsRPC/Portfolio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAccountsRPCServer).Portfolio(ctx, req.(*PortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAccountsRPC_OrderStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAccountsRPCServer).OrderStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_accounts_rpc.InjectiveAccountsRPC/OrderStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAccountsRPCServer).OrderStates(ctx, req.(*OrderStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAccountsRPC_SubaccountsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAccountsRPCServer).SubaccountsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAccountsRPCServer).SubaccountsList(ctx, req.(*SubaccountsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAccountsRPC_SubaccountBalancesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountBalancesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAccountsRPCServer).SubaccountBalancesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountBalancesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAccountsRPCServer).SubaccountBalancesList(ctx, req.(*SubaccountBalancesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAccountsRPC_SubaccountBalanceEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAccountsRPCServer).SubaccountBalanceEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountBalanceEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAccountsRPCServer).SubaccountBalanceEndpoint(ctx, req.(*SubaccountBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAccountsRPC_StreamSubaccountBalance_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamSubaccountBalanceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveAccountsRPCServer).StreamSubaccountBalance(m, &injectiveAccountsRPCStreamSubaccountBalanceServer{stream})
}

type InjectiveAccountsRPC_StreamSubaccountBalanceServer interface {
	Send(*StreamSubaccountBalanceResponse) error
	grpc.ServerStream
}

type injectiveAccountsRPCStreamSubaccountBalanceServer struct {
	grpc.ServerStream
}

func (x *injectiveAccountsRPCStreamSubaccountBalanceServer) Send(m *StreamSubaccountBalanceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveAccountsRPC_SubaccountHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAccountsRPCServer).SubaccountHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAccountsRPCServer).SubaccountHistory(ctx, req.(*SubaccountHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAccountsRPC_SubaccountOrderSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubaccountOrderSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAccountsRPCServer).SubaccountOrderSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_accounts_rpc.InjectiveAccountsRPC/SubaccountOrderSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAccountsRPCServer).SubaccountOrderSummary(ctx, req.(*SubaccountOrderSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAccountsRPC_Rewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAccountsRPCServer).Rewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_accounts_rpc.InjectiveAccountsRPC/Rewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAccountsRPCServer).Rewards(ctx, req.(*RewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InjectiveAccountsRPC_ServiceDesc is the grpc.ServiceDesc for InjectiveAccountsRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InjectiveAccountsRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "injective_accounts_rpc.InjectiveAccountsRPC",
	HandlerType: (*InjectiveAccountsRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Portfolio",
			Handler:    _InjectiveAccountsRPC_Portfolio_Handler,
		},
		{
			MethodName: "OrderStates",
			Handler:    _InjectiveAccountsRPC_OrderStates_Handler,
		},
		{
			MethodName: "SubaccountsList",
			Handler:    _InjectiveAccountsRPC_SubaccountsList_Handler,
		},
		{
			MethodName: "SubaccountBalancesList",
			Handler:    _InjectiveAccountsRPC_SubaccountBalancesList_Handler,
		},
		{
			MethodName: "SubaccountBalanceEndpoint",
			Handler:    _InjectiveAccountsRPC_SubaccountBalanceEndpoint_Handler,
		},
		{
			MethodName: "SubaccountHistory",
			Handler:    _InjectiveAccountsRPC_SubaccountHistory_Handler,
		},
		{
			MethodName: "SubaccountOrderSummary",
			Handler:    _InjectiveAccountsRPC_SubaccountOrderSummary_Handler,
		},
		{
			MethodName: "Rewards",
			Handler:    _InjectiveAccountsRPC_Rewards_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSubaccountBalance",
			Handler:       _InjectiveAccountsRPC_StreamSubaccountBalance_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "injective_accounts_rpc.proto",
}
