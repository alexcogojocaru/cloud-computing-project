run:
	go run -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" .

gen:
	protoc --proto_path=schemas schemas/*.proto --go_out=proto-gen
	protoc --proto_path=schemas schemas/*.proto --go-grpc_out=proto-gen
