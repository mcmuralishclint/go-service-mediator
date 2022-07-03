run:
	go run main.go
mediator-grpc:
	protoc --go-grpc_out=pb --go_out=pb pb/mediators.proto
test:
	go test ./... -cover
