package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mcmuralishclint/go-service-mediator/parser"
	"github.com/mcmuralishclint/go-service-mediator/pb/mediators"
	"google.golang.org/protobuf/types/known/structpb"
)

var Config *parser.Config

func RetreiveConfig(request *mediators.MediationInput) (parser.Endpoint, string, error) {
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
				if (endpoint.Name == request.Endpoint) && (endpoint.Verb == request.GetVerb()) {
					return endpoint, baseUrl, nil
				}
			}
		}
	}
	return parser.Endpoint{}, "", errors.New("No config found")
}

func MakeHttpCall(endpoint parser.Endpoint, baseUrl string, requestData *structpb.Struct) []byte {
	body, _ := json.Marshal(requestData.Fields["body"])

	request_url := baseUrl + endpoint.Url
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest(endpoint.Verb, request_url, bytes.NewBuffer(body))
	request.Header.Set("Content-type", endpoint.ContentType)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	return body
}
