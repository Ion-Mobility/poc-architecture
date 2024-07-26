package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Env env

type env struct {
	AppPort    int
	AppSignKey string
}

func LoadEnv() {
	godotenv.Load()

	Env.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	Env.AppSignKey = os.Getenv("APP_SIGN_KEY")
}
