  proto:
	protoc -I . protobuf/proxy_client.proto --go_out=plugins=grpc:.
proto2:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/*.proto --experimental_allow_proto3_optional


