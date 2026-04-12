package gen

//go:generate mockgen -package mock -destination mock/plugin/v1/gen/mock_narrator.go ./pb/plugin/v1 NarratorServiceClient,NarratorServiceServer
//go:generate mockgen -package mock -destination mock/plugin/v1/gen/mock_source.go ./pb/plugin/v1 SourceServiceClient,SourceServiceServer
