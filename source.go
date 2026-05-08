package katarive

import (
	"context"

	"github.com/hashicorp/go-plugin"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/plugin/v1"
	"google.golang.org/grpc"
)

type sourceGRPCClient struct {
	client pb.SourceServiceClient
}

func (c *sourceGRPCClient) GetSourceServiceMetadata(
	ctx context.Context,
	req *pb.GetSourceServiceMetadataRequest,
	opt ...grpc.CallOption,
) (*pb.GetSourceServiceMetadataResponse, error) {
	return c.client.GetSourceServiceMetadata(ctx, req)
}
func (c *sourceGRPCClient) GetSource(
	ctx context.Context,
	req *pb.GetSourceRequest,
	opt ...grpc.CallOption,
) (*pb.GetSourceResponse, error) {
	return c.client.GetSource(ctx, req)
}
func (c *sourceGRPCClient) ListSources(
	ctx context.Context,
	req *pb.ListSourcesRequest,
	opt ...grpc.CallOption,
) (*pb.ListSourcesResponse, error) {
	return c.client.ListSources(ctx, req)
}

// Check Source implementation
var _ pb.SourceServiceClient = new(sourceGRPCClient)

type sourceGRPCServer struct {
	pb.UnimplementedSourceServiceServer
	Impl pb.SourceServiceServer
}

func (s *sourceGRPCServer) GetSourceServiceMetadata(
	ctx context.Context,
	req *pb.GetSourceServiceMetadataRequest,
) (*pb.GetSourceServiceMetadataResponse, error) {
	return s.Impl.GetSourceServiceMetadata(ctx, req)
}
func (s *sourceGRPCServer) GetSource(
	ctx context.Context,
	req *pb.GetSourceRequest,
) (*pb.GetSourceResponse, error) {
	return s.Impl.GetSource(ctx, req)
}
func (s *sourceGRPCServer) ListSources(
	ctx context.Context,
	req *pb.ListSourcesRequest,
) (*pb.ListSourcesResponse, error) {
	return s.Impl.ListSources(ctx, req)
}

// Check SourceServiceServer implementation
var _ pb.SourceServiceServer = new(sourceGRPCServer)

type SourcePlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl pb.SourceServiceServer
}

func (p *SourcePlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterSourceServiceServer(server, &sourceGRPCServer{Impl: p.Impl})
	return nil
}
func (p *SourcePlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	conn *grpc.ClientConn,
) (interface{}, error) {
	return &sourceGRPCClient{client: pb.NewSourceServiceClient(conn)}, nil
}

// Check plugin.GRPCPlugin implementation
var _ plugin.GRPCPlugin = new(SourcePlugin)
