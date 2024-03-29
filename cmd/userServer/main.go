package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GermanVor/data-keeper/cmd/userServer/service"
	"github.com/GermanVor/data-keeper/cmd/userServer/storage"
	"github.com/GermanVor/data-keeper/internal/common"
)

var (
	addr        = common.DEFAULT_USER_SERVICE_ADDR
	dataBaseDSN = common.DEFAULT_DATA_BASE_DSN
	envFilePath = common.DEFAULT_ENV_PATH
)

func initConfig() {
	common.LoadEnvFile(&envFilePath, envFilePath)

	if envAddr, ok := os.LookupEnv("ADDR"); ok {
		addr = envAddr
	}

	if envDataBaseDSN, ok := os.LookupEnv("DATA_BASE_DSN"); ok {
		dataBaseDSN = envDataBaseDSN
	}

	flag.StringVar(&addr, "a", addr, "address of the service")
	flag.StringVar(&dataBaseDSN, "d", dataBaseDSN, "")
	flag.Parse()
}

func main() {
	initConfig()

	stor := storage.Init(dataBaseDSN)
	s := service.Init(addr, stor)

	err := s.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
