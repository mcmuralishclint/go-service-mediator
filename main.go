package main

import (
	"context"
	"fmt"
	"net"

	"github.com/mcmuralishclint/go-service-mediator/parser"
	"github.com/mcmuralishclint/go-service-mediator/pb/mediators"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var Config *parser.Config

type server struct {
	mediators.UnimplementedMediateServer
}

func main() {
	Config = parser.Parser("config/config.yaml")
	fmt.Println("Config read successfully")
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	mediators.RegisterMediateServer(srv, &server{})
	reflection.Register(srv)
	fmt.Println("Grpc server up and running")
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Mediate(ctx context.Context, request *mediators.MediationInput) (*mediators.MediationOutput, error) {
	service := request.GetService()
	version := request.GetVersion()
	endpoint := request.GetEndpoint()
	requestData := request.GetRequestData()
	fmt.Println(service, version, endpoint, requestData)
	return &mediators.MediationOutput{}, nil
}
