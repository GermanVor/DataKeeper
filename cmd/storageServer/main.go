package main

import (
	"flag"

	"github.com/GermanVor/data-keeper/cmd/storageServer/service"
	"github.com/GermanVor/data-keeper/cmd/storageServer/storage"
	"github.com/GermanVor/data-keeper/internal/common"
)

var (
	addr        = common.DEFAULT_STORAGE_SERVICE_ADDR
	userAddr    = common.DEFAULT_USER_SERVICE_ADDR
	dataBaseDSN = common.DEFAULT_DATA_BASE_DSN
)

func init() {
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
