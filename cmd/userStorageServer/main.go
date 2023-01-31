package main

import (
	"flag"
	"fmt"

	"github.com/GermanVor/data-keeper/cmd/userStorageServer/service"
	"github.com/GermanVor/data-keeper/cmd/userStorageServer/storage"
)

var (
	addr        = ":1234"
	dataBaseDSN = "postgres://zzman:@localhost:5432/postgres"
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
