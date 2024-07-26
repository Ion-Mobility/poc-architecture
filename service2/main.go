package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/ad3n/microservices/grpcs"
	"github.com/ad3n/microservices/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	godotenv.Load()
}

func main() {

	r := gin.Default()

	r.GET("/hello", middlewares.ValidateToken(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Service 2",
		})
	})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.EmptyServerOption{})

	grpcs.RegisterHelloServer(grpcServer, &grpcs.Hello{})

	go grpcServer.Serve(listen)

	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))

	grpcServer.Stop()
}
