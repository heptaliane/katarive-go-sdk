package katarive

import (
	"context"

	"github.com/hashicorp/go-plugin"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/plugin/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type narratorGRPCClient struct {
	pb.UnimplementedNarratorServiceServer
	client pb.NarratorServiceClient
}

func (c *narratorGRPCClient) GetNarratorServiceMetadata(
	ctx context.Context,
	req *pb.GetNarratorServiceMetadataRequest,
) (*pb.GetNarratorServiceMetadataResponse, error) {
	return c.client.GetNarratorServiceMetadata(ctx, req)
}
func (c *narratorGRPCClient) Narrate(
	ctx context.Context,
	req *pb.NarrateRequest,
) (*pb.NarrateResponse, error) {
	return c.client.Narrate(ctx, req)
}

// Check Narrator implementation
var _ pb.NarratorServiceServer = new(narratorGRPCClient)

type narratorGRPCServer struct {
	pb.UnimplementedNarratorServiceServer
	Impl pb.NarratorServiceServer
}

func (s *narratorGRPCServer) GetNarratorServiceMetadata(
	ctx context.Context,
	req *pb.GetNarratorServiceMetadataRequest,
) (*pb.GetNarratorServiceMetadataResponse, error) {
	return s.Impl.GetNarratorServiceMetadata(ctx, req)
}
func (s *narratorGRPCServer) Narrate(
	ctx context.Context,
	req *pb.NarrateRequest,
) (*pb.NarrateResponse, error) {
	return s.Impl.Narrate(ctx, req)
}

// Check NarratorServiceServer implementation
var _ pb.NarratorServiceServer = new(narratorGRPCServer)

type NarratorPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl pb.NarratorServiceServer
}

func (p *NarratorPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterNarratorServiceServer(server, &narratorGRPCServer{Impl: p.Impl})
	reflection.Register(server)
	return nil
}
func (p *NarratorPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	conn *grpc.ClientConn,
) (interface{}, error) {
	return &narratorGRPCClient{client: pb.NewNarratorServiceClient(conn)}, nil
}

// Check plugin.GRPCPlugin implementation
var _ plugin.GRPCPlugin = new(NarratorPlugin)
