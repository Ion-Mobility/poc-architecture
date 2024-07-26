package main

import (
	"fmt"

	"github.com/ad3n/microservices/configs"
	"github.com/ad3n/microservices/controllers"
	"github.com/ad3n/microservices/repositories"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnv()
}

func main() {

	var repository = repositories.UserRepository{}
	var login = controllers.Login{Repository: &repository}

	r := gin.Default()

	r.POST("/login", login.Auth)
	r.POST("/validate", login.Validate)

	r.Run(fmt.Sprintf(":%d", configs.Env.AppPort))
}
