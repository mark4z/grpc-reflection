export PATH=PATH:/Users/mark4zlv/go/bin

gen-api:
	protoc --proto_path=./vendor --proto_path=./ --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./api/helloworld.proto

curl:
	grpcurl --plaintext -protoset=./dep.protoset localhost:50051 describe

protosets:
	protoc --proto_path=. --proto_path=./vendor \
		--descriptor_set_out=dep.protoset \
		--include_imports \
		api/helloworld.proto