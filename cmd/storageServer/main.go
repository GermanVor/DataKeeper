package main

import (
	"flag"

	"github.com/GermanVor/data-keeper/cmd/storageServer/service"
	"github.com/GermanVor/data-keeper/cmd/storageServer/storage"
)

var (
	addr        = ":5678"
	userAddr    = ":1234"
	dataBaseDSN = "postgres://zzman:@localhost:5432/postgres"
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
