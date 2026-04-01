package katarive

import (
	"context"

	"github.com/hashicorp/go-plugin"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/plugin/v1"
	"google.golang.org/grpc"
)

type Narrator interface {
	GetMetadata(ctx context.Context) (*pb.GetMetadataResponse, error)
	Narrate(ctx context.Context, path string, lines []string, options map[string]string) (*pb.NarrateResponse, error)
}

type narratorGRPCClient struct {
	client pb.NarratorServiceClient
}

func (c *narratorGRPCClient) GetMetadata(ctx context.Context) (*pb.GetMetadataResponse, error) {
	return c.client.GetMetadata(ctx, &pb.GetMetadataRequest{})
}
func (c *narratorGRPCClient) Narrate(ctx context.Context, path string, lines []string, options map[string]string) (*pb.NarrateResponse, error) {
	return c.client.Narrate(ctx, &pb.NarrateRequest{
		Path:    path,
		Lines:   lines,
		Options: options,
	})
}

// Check Narrator implementation
var _ Narrator = new(narratorGRPCClient)

type narratorGRPCServer struct {
	pb.UnimplementedNarratorServiceServer
	Impl Narrator
}

func (s *narratorGRPCServer) GetMetadata(ctx context.Context, _req *pb.GetMetadataRequest) (*pb.GetMetadataResponse, error) {
	return s.Impl.GetMetadata(ctx)
}
func (s *narratorGRPCServer) Narrate(ctx context.Context, req *pb.NarrateRequest) (*pb.NarrateResponse, error) {
	return s.Impl.Narrate(ctx, req.Path, req.Lines, req.Options)
}

// Check NarratorServiceServer implementation
var _ pb.NarratorServiceServer = new(narratorGRPCServer)

type NarratorPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Narrator
}

func (p *NarratorPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterNarratorServiceServer(server, &narratorGRPCServer{Impl: p.Impl})
	return nil
}
func (p *NarratorPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &narratorGRPCClient{client: pb.NewNarratorServiceClient(conn)}, nil
}

// Check plugin.GRPCPlugin implementation
var _ plugin.GRPCPlugin = new(NarratorPlugin)
