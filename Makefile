generate_grpc_code:
	protoc --proto_path=. ./chat.proto --go_out=.