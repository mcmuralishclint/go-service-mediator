package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcmuralishclint/go-service-mediator/pb/mediators"
	"google.golang.org/grpc"
)

type RequestData struct {
	Header     map[string]interface{} `json:"header"`
	Body       map[string]interface{} `json:"body"`
	QueryParam map[string]interface{} `json:"queryparam"`
	PathParam  map[string]interface{} `json:"pathparam"`
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := mediators.NewMediateClient(conn)
	g := gin.Default()

	g.POST("/api/:version/orchestrator/:service/:endpoint", func(ctx *gin.Context) {
		version := ctx.Param("version")
		service := ctx.Param("service")
		endpoint := ctx.Param("endpoint")
		requestData := RequestData{}
		err := ctx.ShouldBindJSON(&requestData)
		if err != nil {
			fmt.Println("Error when processing endpoint")
		} else {
			log.Println(version)
			log.Println(service)
			log.Println(endpoint)
			log.Println(requestData)
		}
		req := &mediators.MediationInput{Service: service, Version: version, Endpoint: endpoint}
		if response, err := client.Mediate(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": response})
		}
	})
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
