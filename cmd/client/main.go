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
	addr            = ":5678"
	userAddr        = ":1234"
	secretValue     = common.DEFAULT_USER_SECRET
	secretValuePath = ""
)

type UserClient struct {
}

func init() {
	flag.StringVar(&addr, "a", addr, "address of the service")
	flag.StringVar(&userAddr, "ua", userAddr, "address of the user service")
	flag.StringVar(&secretValuePath, "s", secretValuePath, "path to the file with secret")
}

var (
	loginComm  = "l"
	signInComm = "s"
)

func main() {
	flag.Parse()

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

	ctx := context.Background()
	if ans == signInComm {
		fmt.Print("\nEnter email: ")
		email, _ := reader.ReadString('\n')

		if secretValuePath != "" {
			if secretFromFile, err := ioutil.ReadFile(secretValuePath); err == nil {
				secretValue = string(secretFromFile[:])
			} else {
				log.Println(err.Error())
				return
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
		return
	}

	client.Start(reader, ctx)
}
