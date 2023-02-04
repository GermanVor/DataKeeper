package main

import (
	"flag"
	"fmt"

	"github.com/GermanVor/data-keeper/cmd/userServer/service"
	"github.com/GermanVor/data-keeper/cmd/userServer/storage"
	"github.com/GermanVor/data-keeper/internal/common"
)

var (
	addr        = common.DEFAULT_USER_SERVICE_ADDR
	dataBaseDSN = common.DEFAULT_DATA_BASE_DSN
)

func init() {
	flag.StringVar(&addr, "a", addr, "address of the service")
	flag.StringVar(&dataBaseDSN, "d", dataBaseDSN, "")
}

func main() {
	flag.Parse()

	stor := storage.Init(dataBaseDSN)
	s := service.Init(addr, stor)

	err := s.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
