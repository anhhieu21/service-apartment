proto:
	protoc api/proto/apartment.proto --go-grpc_out=.
	protoc api/proto/apartment.proto --go_out=.
run-server:
	go run main.go