package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/GermanVor/data-keeper/internal/common"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/GermanVor/data-keeper/cmd/client/rpc"
)

var (
	addr            = common.DEFAULT_STORAGE_SERVICE_ADDR
	userAddr        = common.DEFAULT_USER_SERVICE_ADDR
	secretValue     = common.DEFAULT_USER_SECRET
	secretValuePath = ""
	envFilePath     = common.DEFAULT_ENV_PATH
)

func initConfig() {
	common.LoadEnvFile(&envFilePath, envFilePath)

	if envAddr, ok := os.LookupEnv("ADDR"); ok {
		addr = envAddr
	}

	if envUserAddr, ok := os.LookupEnv("USER_SERVICE_ADDR"); ok {
		userAddr = envUserAddr
	}

	if envSecretValuePath, ok := os.LookupEnv("SECRET_PATH"); ok {
		secretValuePath = envSecretValuePath
	}

	flag.StringVar(&addr, "a", addr, "address of the service")
	flag.StringVar(&userAddr, "ua", userAddr, "address of the user service")
	flag.StringVar(&secretValuePath, "s", secretValuePath, "path to the file with secret")
	flag.Parse()
}

var (
	loginComm  = "l"
	signInComm = "s"
)

func main() {
	initConfig()

	client, err := rpc.Init(userAddr, addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bufio.NewReader(os.Stdin)

	ans := ""
	for {
		fmt.Print("Shell you login (" + loginComm + ") or sign (" + signInComm + ") in ?: ")
		ans, _ = reader.ReadString('\n')
		ans = strings.TrimSpace(ans)

		if ans == loginComm || ans == signInComm {
			break
		}

		fmt.Println("Unknown command", ans, "Try again")
	}

	ctx := context.Background()

	for {
		fmt.Print("Enter login: ")
		login, _ := reader.ReadString('\n')
		login = strings.TrimSpace(login)

		fmt.Print("Enter password: ")
		bytePassword, err := terminal.ReadPassword(0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		password := strings.TrimSpace(string(bytePassword))

		if ans == signInComm {
			fmt.Print("\nEnter email: ")
			email, _ := reader.ReadString('\n')

			if secretValuePath != "" {
				if secretFromFile, err := ioutil.ReadFile(secretValuePath); err == nil {
					secretValue = string(secretFromFile[:])
				} else {
					log.Println(err.Error())
					continue
				}
			}

			fmt.Println(login, password, email)

			err = client.SignIn(ctx, &rpc.SignIn{
				Login:    login,
				Password: password,
				Email:    email,
				Secret:   secretValue,
			})
		} else {
			err = client.LogIn(ctx, &rpc.LogIn{
				Login:    login,
				Password: password,
			})
		}

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		break
	}

	client.Start(reader, ctx)
}
