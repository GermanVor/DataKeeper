package main

import (
	"flag"
	"os"

	"github.com/GermanVor/data-keeper/cmd/storageServer/service"
	"github.com/GermanVor/data-keeper/cmd/storageServer/storage"
	"github.com/GermanVor/data-keeper/internal/common"
	"github.com/joho/godotenv"
)

var (
	addr        = common.DEFAULT_STORAGE_SERVICE_ADDR
	userAddr    = common.DEFAULT_USER_SERVICE_ADDR
	dataBaseDSN = common.DEFAULT_DATA_BASE_DSN
)

func init() {
	godotenv.Load(".env")

	if envAddr, ok := os.LookupEnv("ADDR"); ok {
		addr = envAddr
	}

	if envUserAddr, ok := os.LookupEnv("USER_SERVICE_ADDR"); ok {
		userAddr = envUserAddr
	}

	if envDataBaseDSN, ok := os.LookupEnv("DATA_BASE_DSN"); ok {
		dataBaseDSN = envDataBaseDSN
	}

	flag.StringVar(&addr, "a", addr, "address of the service")
	flag.StringVar(&userAddr, "ua", userAddr, "address of the user service")
	flag.StringVar(&dataBaseDSN, "d", dataBaseDSN, "")
}

func main() {
	flag.Parse()

	stor := storage.Init(dataBaseDSN)
	s := service.Init(addr, stor, userAddr)

	s.Start()
}
