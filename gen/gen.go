package gen

//go:generate mockgen -package mock -destination pb/plugin/v1/mock/mock_narrator.go ./pb/plugin/v1 NarratorServiceClient,NarratorServiceServer,UnsafeNarratorServiceServer
//go:generate mockgen -package mock -destination pb/plugin/v1/mock/mock_source.go ./pb/plugin/v1 SourceServiceClient,SourceServiceServer,UnsafeSourceServiceServer
