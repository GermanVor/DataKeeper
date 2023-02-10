package common

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	JWT_CTX_NAME     = "jwt"
	USER_ID_CTX_NAME = "userId"

	DEFAULT_USER_SERVICE_ADDR    = ":1234"
	DEFAULT_STORAGE_SERVICE_ADDR = ":5678"
	DEFAULT_DATA_BASE_DSN        = "postgres://zzman:@localhost:5432/postgres"
	DEFAULT_USER_SECRET          = "7f9c2ba4e88f827d616045507605853ed73b8093f6efbc88eb1a6eacfa66ef26"

	DEFAULT_ENV_PATH = ".env"
)

func LoadEnvFile(envFilePath *string, defaultFilePath string) {
	flag.StringVar(envFilePath, "p", defaultFilePath, "path to the file to download variables")
	flag.Parse()

	err := godotenv.Load(*envFilePath)
	if err != nil && *envFilePath != defaultFilePath {
		fmt.Println(err.Error(), *envFilePath)
		os.Exit(1)
	}
}
