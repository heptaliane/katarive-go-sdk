package katarive

import (
	"context"

	"github.com/hashicorp/go-plugin"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/plugin/v1"
	"google.golang.org/grpc"
)

type Narrator interface {
	Narrate(
		ctx context.Context,
		path string,
		text string,
		options map[string]string,
	) (*pb.NarrateResponse, error)

	GetNarratorServiceMetadata(
		ctx context.Context,
	) (*pb.GetNarratorServiceMetadataResponse, error)
}

type narratorGRPCClient struct {
	client pb.NarratorServiceClient
}

func (c *narratorGRPCClient) GetNarratorServiceMetadata(
	ctx context.Context,
) (*pb.GetNarratorServiceMetadataResponse, error) {
	return c.client.GetNarratorServiceMetadata(ctx, &pb.GetNarratorServiceMetadataRequest{})
}
func (c *narratorGRPCClient) Narrate(
	ctx context.Context,
	path string,
	text string,
	options map[string]string,
) (*pb.NarrateResponse, error) {
	return c.client.Narrate(ctx, &pb.NarrateRequest{
		Path:    path,
		Text:    text,
		Options: options,
	})
}

// Check Narrator implementation
var _ Narrator = new(narratorGRPCClient)

type narratorGRPCServer struct {
	pb.UnimplementedNarratorServiceServer
	Impl Narrator
}

func (s *narratorGRPCServer) GetNarratorServiceMetadata(
	ctx context.Context,
	_req *pb.GetNarratorServiceMetadataRequest,
) (*pb.GetNarratorServiceMetadataResponse, error) {
	return s.Impl.GetNarratorServiceMetadata(ctx)
}
func (s *narratorGRPCServer) Narrate(
	ctx context.Context,
	req *pb.NarrateRequest,
) (*pb.NarrateResponse, error) {
	return s.Impl.Narrate(ctx, req.Path, req.Text, req.Options)
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
func (p *NarratorPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	conn *grpc.ClientConn,
) (interface{}, error) {
	return &narratorGRPCClient{client: pb.NewNarratorServiceClient(conn)}, nil
}

// Check plugin.GRPCPlugin implementation
var _ plugin.GRPCPlugin = new(NarratorPlugin)
