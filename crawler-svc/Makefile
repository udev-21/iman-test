proto:
	protoc pkg/**/pb/*.proto --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=.

server:
	go run cmd/main.go