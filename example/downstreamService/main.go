package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/mcmuralishclint/go-service-mediator/pb/mediators"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := mediators.NewMediateClient(conn)
	g := gin.Default()

	g.GET("/api/:version/orchestrator/:service/:endpoint", func(ctx *gin.Context) {
		version := ctx.Param("version")
		service := ctx.Param("service")
		endpoint := ctx.Param("endpoint")
		verb := ctx.Request.Method
		requestData := structpb.Struct{}
		err := ctx.ShouldBindJSON(&requestData)
		if err != nil {
			log.Fatal("Unable to process endpoint")
		}
		req := &mediators.MediationInput{Service: service, Version: version, Endpoint: endpoint, Verb: verb, RequestData: &requestData}
		if response, err := client.Mediate(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": response})
		}
	})

	g.POST("/api/:version/orchestrator/:service/:endpoint", func(ctx *gin.Context) {
		version := ctx.Param("version")
		service := ctx.Param("service")
		endpoint := ctx.Param("endpoint")
		verb := ctx.Request.Method
		requestData := structpb.Struct{}
		err := ctx.ShouldBindJSON(&requestData)
		if err != nil {
			log.Fatal("Unable to process endpoint")
		}
		req := &mediators.MediationInput{Service: service, Version: version, Endpoint: endpoint, Verb: verb, RequestData: &requestData}
		response, err := client.Mediate(ctx, req)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": response})
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
