package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"

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
	// service := request.GetService()
	// version := request.GetVersion()
	// endpoint := request.GetEndpoint()
	// requestData := request.GetRequestData()
	fmt.Println(request)
	endpoint, baseUrl, err := retreiveConfig(request)
	if err != nil {
		return &mediators.MediationOutput{}, errors.New("Invalid Request")
	}
	fmt.Println(endpoint)
	MakeHttpCall(endpoint, baseUrl)
	return &mediators.MediationOutput{}, nil
}

func retreiveConfig(request *mediators.MediationInput) (parser.Endpoint, string, error) {
	for _, service := range Config.Services {
		serviceName := service["service1"].Name
		if request.GetService() == serviceName {
			versions := service[serviceName].Version
			baseUrl := service[serviceName].BaseUrl
			versionToFind, err := strconv.Atoi(request.GetVersion()[1:])
			if err != nil {
				fmt.Println("Invalid Version")
				return parser.Endpoint{}, baseUrl, errors.New("Invalid version")
			}
			endpoints := versions[versionToFind-1]["endpoints"]
			for _, endpoint := range endpoints {
				if endpoint.Name == request.Endpoint {
					return endpoint, baseUrl, nil
				}
			}
		}
	}
	return parser.Endpoint{}, "", errors.New("No config found")
}

func MakeHttpCall(endpoint parser.Endpoint, baseUrl string) {
	request := baseUrl + endpoint.Url
	fmt.Println(request)
	resp, err := http.Get(request)
	if err != nil {
		log.Fatal("Request failed")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
