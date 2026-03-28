package katarive

import (
	"context"

	"github.com/hashicorp/go-plugin"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/plugin/v1"
	"google.golang.org/grpc"
)

type Speaker interface {
	GetMetadata(ctx context.Context) (*pb.GetMetadataResponse, error)
	Synthesize(ctx context.Context, lines []string, options map[string]string) (*pb.SynthesizeResponse, error)
}

type speakerGRPCClient struct {
	client pb.SpeakerServiceClient
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

// Check Speaker implementation
var _ Speaker = new(speakerGRPCClient)

type speakerGRPCServer struct {
	pb.UnimplementedSpeakerServiceServer
	Impl Speaker
}

func (s *speakerGRPCServer) GetMetadata(ctx context.Context, _req *pb.GetMetadataRequest) (*pb.GetMetadataResponse, error) {
	return s.Impl.GetMetadata(ctx)
}
func (s *speakerGRPCServer) Synthesize(ctx context.Context, req *pb.SynthesizeRequest) (*pb.SynthesizeResponse, error) {
	return s.Impl.Synthesize(ctx, req.Lines, req.Options)
}

// Check SpeakerServiceServer implementation
var _ pb.SpeakerServiceServer = new(speakerGRPCServer)

type SpeakerPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Speaker
}

func (p *SpeakerPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterSpeakerServiceServer(server, &speakerGRPCServer{Impl: p.Impl})
	return nil
}
func (p *SpeakerPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &speakerGRPCClient{client: pb.NewSpeakerServiceClient(conn)}, nil
}

// Check plugin.GRPCPlugin implementation
var _ plugin.GRPCPlugin = new(SpeakerPlugin)
