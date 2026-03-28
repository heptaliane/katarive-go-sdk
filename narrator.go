package katarive

import (
	"context"

	"github.com/hashicorp/go-plugin"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/plugin/v1"
	"google.golang.org/grpc"
)

type Narrator interface {
	GetMetadata(ctx context.Context) (*pb.GetMetadataResponse, error)
	Synthesize(ctx context.Context, lines []string, options map[string]string) (*pb.SynthesizeResponse, error)
}

type speakerGRPCClient struct {
	client pb.NarratorServiceClient
}

func (c *speakerGRPCClient) GetMetadata(ctx context.Context) (*pb.GetMetadataResponse, error) {
	return c.client.GetMetadata(ctx, &pb.GetMetadataRequest{})
}
func (c *speakerGRPCClient) Synthesize(ctx context.Context, lines []string, options map[string]string) (*pb.SynthesizeResponse, error) {
	return c.client.Synthesize(ctx, &pb.SynthesizeRequest{
		Lines:   lines,
		Options: options,
	})
}

// Check Narrator implementation
var _ Narrator = new(speakerGRPCClient)

type speakerGRPCServer struct {
	pb.UnimplementedNarratorServiceServer
	Impl Narrator
}

func (s *speakerGRPCServer) GetMetadata(ctx context.Context, _req *pb.GetMetadataRequest) (*pb.GetMetadataResponse, error) {
	return s.Impl.GetMetadata(ctx)
}
func (s *speakerGRPCServer) Synthesize(ctx context.Context, req *pb.SynthesizeRequest) (*pb.SynthesizeResponse, error) {
	return s.Impl.Synthesize(ctx, req.Lines, req.Options)
}

// Check NarratorServiceServer implementation
var _ pb.NarratorServiceServer = new(speakerGRPCServer)

type NarratorPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Narrator
}

func (p *NarratorPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterNarratorServiceServer(server, &speakerGRPCServer{Impl: p.Impl})
	return nil
}
func (p *NarratorPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &speakerGRPCClient{client: pb.NewNarratorServiceClient(conn)}, nil
}

// Check plugin.GRPCPlugin implementation
var _ plugin.GRPCPlugin = new(NarratorPlugin)
