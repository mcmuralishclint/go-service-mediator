run:
	go run main.go
run-server:
	go run example/upstreamService/main.go
run-client:
	go run example/downstreamService/main.go
mediator-grpc:
	protoc --go-grpc_out=pb --go_out=pb pb/mediators.proto
# export PATH="$PATH:$(go env GOPATH)/bin"

test:
	go test ./... -cover
