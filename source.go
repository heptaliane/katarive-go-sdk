package katarive

import (
	"context"

	"github.com/hashicorp/go-plugin"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/proto/v1"
	"google.golang.org/grpc"
)

type Source interface {
	GetSupportedPatterns(ctx context.Context) (*pb.GetSupportedPatternsResponse, error)
	GetSource(ctx context.Context, url string) (*pb.GetSourceResponse, error)
}

type sourceGRPCClient struct {
	client pb.SourceServiceClient
}

func (c *sourceGRPCClient) GetSupportedPatterns(ctx context.Context) (*pb.GetSupportedPatternsResponse, error) {
	return c.client.GetSupportedPatterns(ctx, &pb.GetSupportedPatternsRequest{})
}
func (c *sourceGRPCClient) GetSource(ctx context.Context, url string) (*pb.GetSourceResponse, error) {
	return c.client.GetSource(ctx, &pb.GetSourceRequest{Url: url})
}

// Check Source implementation
var _ Source = new(sourceGRPCClient)

type sourceGRPCServer struct {
	pb.UnimplementedSourceServiceServer
	Impl Source
}

func (s *sourceGRPCServer) GetSupportedPatterns(ctx context.Context, _req *pb.GetSupportedPatternsRequest) (*pb.GetSupportedPatternsResponse, error) {
	return s.Impl.GetSupportedPatterns(ctx)
}
func (s *sourceGRPCServer) GetSource(ctx context.Context, req *pb.GetSourceRequest) (*pb.GetSourceResponse, error) {
	return s.Impl.GetSource(ctx, req.Url)
}

type SourcePlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Source
}

func (p *SourcePlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterSourceServiceServer(server, &sourceGRPCServer{Impl: p.Impl})
	return nil
}
func (p *SourcePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &sourceGRPCClient{client: pb.NewSourceServiceClient(conn)}, nil
}

// Check plugin.GRPCPlugin implementation
var _ plugin.GRPCPlugin = new(SourcePlugin)
