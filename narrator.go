package katarive

import (
	"context"

	"github.com/hashicorp/go-plugin"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/plugin/v1"
	"google.golang.org/grpc"
)

type Narrator interface {
	GetMetadata(ctx context.Context) (*pb.GetMetadataResponse, error)
	Narrate(ctx context.Context, lines []string, options map[string]string) (*pb.NarrateResponse, error)
	JobStatus(ctx context.Context, jobId string) (*pb.JobStatusResponse, error)
}

type speakerGRPCClient struct {
	client pb.NarratorServiceClient
}

func (c *speakerGRPCClient) GetMetadata(ctx context.Context) (*pb.GetMetadataResponse, error) {
	return c.client.GetMetadata(ctx, &pb.GetMetadataRequest{})
}
func (c *speakerGRPCClient) Narrate(ctx context.Context, lines []string, options map[string]string) (*pb.NarrateResponse, error) {
	return c.client.Narrate(ctx, &pb.NarrateRequest{
		Lines:   lines,
		Options: options,
	})
}
func (c *speakerGRPCClient) JobStatus(ctx context.Context, jobId string) (*pb.JobStatusResponse, error) {
	return c.client.JobStatus(ctx, &pb.JobStatusRequest{
		JobId: jobId,
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
func (s *speakerGRPCServer) Narrate(ctx context.Context, req *pb.NarrateRequest) (*pb.NarrateResponse, error) {
	return s.Impl.Narrate(ctx, req.Lines, req.Options)
}
func (s *speakerGRPCServer) JobStatus(ctx context.Context, req *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
	return s.Impl.JobStatus(ctx, req.JobId)
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
