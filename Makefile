pb-generate:
	buf generate https://github.com/heptaliane/katarive-proto.git --path plugin --clean

generate: pb-generate
	go generate ./...
