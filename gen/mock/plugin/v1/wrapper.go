package v1

import (
	"context"
	gomock "go.uber.org/mock/gomock"

	mock "github.com/heptaliane/katarive-go-sdk/gen/mock/plugin/v1/gen"
	pb "github.com/heptaliane/katarive-go-sdk/gen/pb/plugin/v1"
)

// Wrapper for gen/mock_narrator.go

type MockNarratorServiceClient = mock.MockNarratorServiceClient

var NewMockNarratorServiceClient = mock.NewMockNarratorServiceClient

type MockNarratorServiceServer struct {
	*mock.MockNarratorServiceServer
	pb.UnimplementedNarratorServiceServer
}

func (s *MockNarratorServiceServer) Narrate(
	ctx context.Context,
	req *pb.NarrateRequest,
) (*pb.NarrateResponse, error) {
	return s.MockNarratorServiceServer.Narrate(ctx, req)
}
func (s *MockNarratorServiceServer) GetNarratorServiceMetadata(
	ctx context.Context,
	req *pb.GetNarratorServiceMetadataRequest,
) (*pb.GetNarratorServiceMetadataResponse, error) {
	return s.MockNarratorServiceServer.GetNarratorServiceMetadata(ctx, req)
}
func NewMockNarratorServiceServer(ctrl *gomock.Controller) *MockNarratorServiceServer {
	return &MockNarratorServiceServer{
		MockNarratorServiceServer: mock.NewMockNarratorServiceServer(ctrl),
	}
}

// Ensure NarratorServiceServer implementation
var _ pb.NarratorServiceServer = new(MockNarratorServiceServer)

// Wrapper for gen/mock_source.go

type MockSourceServiceClient = mock.MockSourceServiceClient

var NewMockSourceServiceClient = mock.NewMockSourceServiceClient

type MockSourceServiceServer struct {
	*mock.MockSourceServiceServer
	pb.UnimplementedSourceServiceServer
}

func (s *MockSourceServiceServer) GetSourceItem(
	ctx context.Context,
	req *pb.GetSourceItemRequest,
) (*pb.GetSourceItemResponse, error) {
	return s.MockSourceServiceServer.GetSourceItem(ctx, req)
}
func (s *MockSourceServiceServer) GetSourceServiceMetadata(
	ctx context.Context,
	req *pb.GetSourceServiceMetadataRequest,
) (*pb.GetSourceServiceMetadataResponse, error) {
	return s.MockSourceServiceServer.GetSourceServiceMetadata(ctx, req)
}
func (s *MockSourceServiceServer) GetSourceCollection(
	ctx context.Context,
	req *pb.GetSourceCollectionRequest,
) (*pb.GetSourceCollectionResponse, error) {
	return s.MockSourceServiceServer.GetSourceCollection(ctx, req)
}
func NewMockSourceServiceServer(ctrl *gomock.Controller) *MockSourceServiceServer {
	return &MockSourceServiceServer{
		MockSourceServiceServer: mock.NewMockSourceServiceServer(ctrl),
	}
}

// Ensure SourceServiceServer implementation
var _ pb.SourceServiceServer = new(MockSourceServiceServer)
