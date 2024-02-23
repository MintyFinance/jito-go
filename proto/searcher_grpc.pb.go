// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: searcher.proto

package proto

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SearcherService_SubscribeBundleResults_FullMethodName      = "/searcher.SearcherService/SubscribeBundleResults"
	SearcherService_SubscribeMempool_FullMethodName            = "/searcher.SearcherService/SubscribeMempool"
	SearcherService_SendBundle_FullMethodName                  = "/searcher.SearcherService/SendBundle"
	SearcherService_GetNextScheduledLeader_FullMethodName      = "/searcher.SearcherService/GetNextScheduledLeader"
	SearcherService_GetConnectedLeaders_FullMethodName         = "/searcher.SearcherService/GetConnectedLeaders"
	SearcherService_GetConnectedLeadersRegioned_FullMethodName = "/searcher.SearcherService/GetConnectedLeadersRegioned"
	SearcherService_GetTipAccounts_FullMethodName              = "/searcher.SearcherService/GetTipAccounts"
	SearcherService_GetRegions_FullMethodName                  = "/searcher.SearcherService/GetRegions"
)

// SearcherServiceClient is the client API for SearcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearcherServiceClient interface {
	// Searchers can invoke this endpoint to subscribe to their respective bundle results.
	// A success result would indicate the bundle won its state auction and was submitted to the validator.
	SubscribeBundleResults(ctx context.Context, in *SubscribeBundleResultsRequest, opts ...grpc.CallOption) (SearcherService_SubscribeBundleResultsClient, error)
	// Subscribe to mempool transactions based on a few filters
	SubscribeMempool(ctx context.Context, in *MempoolSubscription, opts ...grpc.CallOption) (SearcherService_SubscribeMempoolClient, error)
	SendBundle(ctx context.Context, in *SendBundleRequest, opts ...grpc.CallOption) (*SendBundleResponse, error)
	// Returns the next scheduled leader connected to the block engine.
	GetNextScheduledLeader(ctx context.Context, in *NextScheduledLeaderRequest, opts ...grpc.CallOption) (*NextScheduledLeaderResponse, error)
	// Returns leader slots for connected jito validator-watcher during the current epoch. Only returns data for this region.
	GetConnectedLeaders(ctx context.Context, in *ConnectedLeadersRequest, opts ...grpc.CallOption) (*ConnectedLeadersResponse, error)
	// Returns leader slots for connected jito validator-watcher during the current epoch.
	GetConnectedLeadersRegioned(ctx context.Context, in *ConnectedLeadersRegionedRequest, opts ...grpc.CallOption) (*ConnectedLeadersRegionedResponse, error)
	// Returns the tip accounts searchers shall transfer funds to for the leader to claim.
	GetTipAccounts(ctx context.Context, in *GetTipAccountsRequest, opts ...grpc.CallOption) (*GetTipAccountsResponse, error)
	// Returns region the client is directly connected to, along with all available regions
	GetRegions(ctx context.Context, in *GetRegionsRequest, opts ...grpc.CallOption) (*GetRegionsResponse, error)
}

type searcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearcherServiceClient(cc grpc.ClientConnInterface) SearcherServiceClient {
	return &searcherServiceClient{cc}
}

func (c *searcherServiceClient) SubscribeBundleResults(ctx context.Context, in *SubscribeBundleResultsRequest, opts ...grpc.CallOption) (SearcherService_SubscribeBundleResultsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearcherService_ServiceDesc.Streams[0], SearcherService_SubscribeBundleResults_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &searcherServiceSubscribeBundleResultsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearcherService_SubscribeBundleResultsClient interface {
	Recv() (*BundleResult, error)
	grpc.ClientStream
}

type searcherServiceSubscribeBundleResultsClient struct {
	grpc.ClientStream
}

func (x *searcherServiceSubscribeBundleResultsClient) Recv() (*BundleResult, error) {
	m := new(BundleResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searcherServiceClient) SubscribeMempool(ctx context.Context, in *MempoolSubscription, opts ...grpc.CallOption) (SearcherService_SubscribeMempoolClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearcherService_ServiceDesc.Streams[1], SearcherService_SubscribeMempool_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &searcherServiceSubscribeMempoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearcherService_SubscribeMempoolClient interface {
	Recv() (*PendingTxNotification, error)
	grpc.ClientStream
}

type searcherServiceSubscribeMempoolClient struct {
	grpc.ClientStream
}

func (x *searcherServiceSubscribeMempoolClient) Recv() (*PendingTxNotification, error) {
	m := new(PendingTxNotification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searcherServiceClient) SendBundle(ctx context.Context, in *SendBundleRequest, opts ...grpc.CallOption) (*SendBundleResponse, error) {
	out := new(SendBundleResponse)
	err := c.cc.Invoke(ctx, SearcherService_SendBundle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searcherServiceClient) GetNextScheduledLeader(ctx context.Context, in *NextScheduledLeaderRequest, opts ...grpc.CallOption) (*NextScheduledLeaderResponse, error) {
	out := new(NextScheduledLeaderResponse)
	err := c.cc.Invoke(ctx, SearcherService_GetNextScheduledLeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searcherServiceClient) GetConnectedLeaders(ctx context.Context, in *ConnectedLeadersRequest, opts ...grpc.CallOption) (*ConnectedLeadersResponse, error) {
	out := new(ConnectedLeadersResponse)
	err := c.cc.Invoke(ctx, SearcherService_GetConnectedLeaders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searcherServiceClient) GetConnectedLeadersRegioned(ctx context.Context, in *ConnectedLeadersRegionedRequest, opts ...grpc.CallOption) (*ConnectedLeadersRegionedResponse, error) {
	out := new(ConnectedLeadersRegionedResponse)
	err := c.cc.Invoke(ctx, SearcherService_GetConnectedLeadersRegioned_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searcherServiceClient) GetTipAccounts(ctx context.Context, in *GetTipAccountsRequest, opts ...grpc.CallOption) (*GetTipAccountsResponse, error) {
	out := new(GetTipAccountsResponse)
	err := c.cc.Invoke(ctx, SearcherService_GetTipAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searcherServiceClient) GetRegions(ctx context.Context, in *GetRegionsRequest, opts ...grpc.CallOption) (*GetRegionsResponse, error) {
	out := new(GetRegionsResponse)
	err := c.cc.Invoke(ctx, SearcherService_GetRegions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearcherServiceServer is the server API for SearcherService service.
// All implementations must embed UnimplementedSearcherServiceServer
// for forward compatibility
type SearcherServiceServer interface {
	// Searchers can invoke this endpoint to subscribe to their respective bundle results.
	// A success result would indicate the bundle won its state auction and was submitted to the validator.
	SubscribeBundleResults(*SubscribeBundleResultsRequest, SearcherService_SubscribeBundleResultsServer) error
	// Subscribe to mempool transactions based on a few filters
	SubscribeMempool(*MempoolSubscription, SearcherService_SubscribeMempoolServer) error
	SendBundle(context.Context, *SendBundleRequest) (*SendBundleResponse, error)
	// Returns the next scheduled leader connected to the block engine.
	GetNextScheduledLeader(context.Context, *NextScheduledLeaderRequest) (*NextScheduledLeaderResponse, error)
	// Returns leader slots for connected jito validator-watcher during the current epoch. Only returns data for this region.
	GetConnectedLeaders(context.Context, *ConnectedLeadersRequest) (*ConnectedLeadersResponse, error)
	// Returns leader slots for connected jito validator-watcher during the current epoch.
	GetConnectedLeadersRegioned(context.Context, *ConnectedLeadersRegionedRequest) (*ConnectedLeadersRegionedResponse, error)
	// Returns the tip accounts searchers shall transfer funds to for the leader to claim.
	GetTipAccounts(context.Context, *GetTipAccountsRequest) (*GetTipAccountsResponse, error)
	// Returns region the client is directly connected to, along with all available regions
	GetRegions(context.Context, *GetRegionsRequest) (*GetRegionsResponse, error)
	mustEmbedUnimplementedSearcherServiceServer()
}

// UnimplementedSearcherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearcherServiceServer struct {
}

func (UnimplementedSearcherServiceServer) SubscribeBundleResults(*SubscribeBundleResultsRequest, SearcherService_SubscribeBundleResultsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeBundleResults not implemented")
}
func (UnimplementedSearcherServiceServer) SubscribeMempool(*MempoolSubscription, SearcherService_SubscribeMempoolServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeMempool not implemented")
}
func (UnimplementedSearcherServiceServer) SendBundle(context.Context, *SendBundleRequest) (*SendBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBundle not implemented")
}
func (UnimplementedSearcherServiceServer) GetNextScheduledLeader(context.Context, *NextScheduledLeaderRequest) (*NextScheduledLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextScheduledLeader not implemented")
}
func (UnimplementedSearcherServiceServer) GetConnectedLeaders(context.Context, *ConnectedLeadersRequest) (*ConnectedLeadersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectedLeaders not implemented")
}
func (UnimplementedSearcherServiceServer) GetConnectedLeadersRegioned(context.Context, *ConnectedLeadersRegionedRequest) (*ConnectedLeadersRegionedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectedLeadersRegioned not implemented")
}
func (UnimplementedSearcherServiceServer) GetTipAccounts(context.Context, *GetTipAccountsRequest) (*GetTipAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTipAccounts not implemented")
}
func (UnimplementedSearcherServiceServer) GetRegions(context.Context, *GetRegionsRequest) (*GetRegionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegions not implemented")
}
func (UnimplementedSearcherServiceServer) mustEmbedUnimplementedSearcherServiceServer() {}

// UnsafeSearcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearcherServiceServer will
// result in compilation errors.
type UnsafeSearcherServiceServer interface {
	mustEmbedUnimplementedSearcherServiceServer()
}

func RegisterSearcherServiceServer(s grpc.ServiceRegistrar, srv SearcherServiceServer) {
	s.RegisterService(&SearcherService_ServiceDesc, srv)
}

func _SearcherService_SubscribeBundleResults_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeBundleResultsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearcherServiceServer).SubscribeBundleResults(m, &searcherServiceSubscribeBundleResultsServer{stream})
}

type SearcherService_SubscribeBundleResultsServer interface {
	Send(*BundleResult) error
	grpc.ServerStream
}

type searcherServiceSubscribeBundleResultsServer struct {
	grpc.ServerStream
}

func (x *searcherServiceSubscribeBundleResultsServer) Send(m *BundleResult) error {
	return x.ServerStream.SendMsg(m)
}

func _SearcherService_SubscribeMempool_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MempoolSubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearcherServiceServer).SubscribeMempool(m, &searcherServiceSubscribeMempoolServer{stream})
}

type SearcherService_SubscribeMempoolServer interface {
	Send(*PendingTxNotification) error
	grpc.ServerStream
}

type searcherServiceSubscribeMempoolServer struct {
	grpc.ServerStream
}

func (x *searcherServiceSubscribeMempoolServer) Send(m *PendingTxNotification) error {
	return x.ServerStream.SendMsg(m)
}

func _SearcherService_SendBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherServiceServer).SendBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearcherService_SendBundle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherServiceServer).SendBundle(ctx, req.(*SendBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearcherService_GetNextScheduledLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextScheduledLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherServiceServer).GetNextScheduledLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearcherService_GetNextScheduledLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherServiceServer).GetNextScheduledLeader(ctx, req.(*NextScheduledLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearcherService_GetConnectedLeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectedLeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherServiceServer).GetConnectedLeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearcherService_GetConnectedLeaders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherServiceServer).GetConnectedLeaders(ctx, req.(*ConnectedLeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearcherService_GetConnectedLeadersRegioned_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectedLeadersRegionedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherServiceServer).GetConnectedLeadersRegioned(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearcherService_GetConnectedLeadersRegioned_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherServiceServer).GetConnectedLeadersRegioned(ctx, req.(*ConnectedLeadersRegionedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearcherService_GetTipAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTipAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherServiceServer).GetTipAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearcherService_GetTipAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherServiceServer).GetTipAccounts(ctx, req.(*GetTipAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearcherService_GetRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherServiceServer).GetRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearcherService_GetRegions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherServiceServer).GetRegions(ctx, req.(*GetRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearcherService_ServiceDesc is the grpc.ServiceDesc for SearcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "searcher.SearcherService",
	HandlerType: (*SearcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBundle",
			Handler:    _SearcherService_SendBundle_Handler,
		},
		{
			MethodName: "GetNextScheduledLeader",
			Handler:    _SearcherService_GetNextScheduledLeader_Handler,
		},
		{
			MethodName: "GetConnectedLeaders",
			Handler:    _SearcherService_GetConnectedLeaders_Handler,
		},
		{
			MethodName: "GetConnectedLeadersRegioned",
			Handler:    _SearcherService_GetConnectedLeadersRegioned_Handler,
		},
		{
			MethodName: "GetTipAccounts",
			Handler:    _SearcherService_GetTipAccounts_Handler,
		},
		{
			MethodName: "GetRegions",
			Handler:    _SearcherService_GetRegions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeBundleResults",
			Handler:       _SearcherService_SubscribeBundleResults_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeMempool",
			Handler:       _SearcherService_SubscribeMempool_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "searcher.proto",
}
