package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/mcmuralishclint/go-service-mediator/handlers"
	"github.com/mcmuralishclint/go-service-mediator/parser"
	"github.com/mcmuralishclint/go-service-mediator/pb/mediators"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/structpb"
)

type server struct {
	mediators.UnimplementedMediateServer
}

func main() {
	handlers.Config = parser.Parser("config/config.yaml")
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
	endpoint, baseUrl, err := handlers.RetreiveConfig(request)
	if err != nil {
		return &mediators.MediationOutput{}, errors.New("Invalid Request")
	}
	body := handlers.MakeHttpCall(endpoint, baseUrl, request.GetRequestData())
	m, _ := structpb.NewStruct(map[string]interface{}{
		"output": string(body),
	})
	return &mediators.MediationOutput{Output: m}, nil
}
