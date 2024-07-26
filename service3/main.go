package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ad3n/microservices/grpcs"
	"github.com/ad3n/microservices/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	godotenv.Load()
}

func main() {

	r := gin.Default()

	r.GET("/insecure", func(c *gin.Context) {
		connection, err := grpc.NewClient(os.Getenv("GRPC_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Backend down",
			})

			return
		}

		response, err := grpcs.NewHelloClient(connection).SayHello(context.TODO(), &grpcs.Unsecure{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"message": response.Message,
		})
	})

	r.GET("/secure", middlewares.ValidateToken(), func(c *gin.Context) {
		connection, err := grpc.NewClient(os.Getenv("GRPC_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Backend down",
			})

			return
		}

		response, err := grpcs.NewHelloClient(connection).SayHelloSecure(context.TODO(), &grpcs.Secure{
			UserId: c.Request.Header.Get("X-User-ID"),
			Email:  c.Request.Header.Get("X-User-Email"),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"message": response.Message,
		})
	})

	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
