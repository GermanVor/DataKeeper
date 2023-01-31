package rpc

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	datakeeperPB "github.com/GermanVor/data-keeper/proto/datakeeper"
	userPB "github.com/GermanVor/data-keeper/proto/user"
)

func readData(data *datakeeperPB.Data, reader *bufio.Reader) error {
	fmt.Print("Value : ")

	switch data.DataType {
	case datakeeperPB.DataType_LOG_PASS:
		fmt.Println(string(data.Data[:]))
	case datakeeperPB.DataType_TEXT:
		fmt.Println(string(data.Data[:]))
	case datakeeperPB.DataType_BANK_CARD:
		fmt.Println(string(data.Data[:]))
	default:
		for {
			fmt.Print("Enter path to save Data : ")
			filePath, _ := reader.ReadString('\n')
			filePath = strings.TrimSpace(filePath)

			err := ioutil.WriteFile(filePath, data.Data, 0644)
			if err != nil {
				return err
			}
		}
	}

	fmt.Print("Meta : ")
	fmt.Println(data.Meta)

	return nil
}

func getData(reader *bufio.Reader) ([]byte, error) {
	fmt.Println("Enter path to new Data (empty string to skip) : ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	var data []byte
	var err error

	if filePath != "" {
		data, err = ioutil.ReadFile(filePath)
	}

	return data, err
}

func getMeta(reader *bufio.Reader) (map[string]string, error) {
	fmt.Println("Enter path to new Meta (empty string to skip) : ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	var meta map[string]string

	if filePath != "" {
		newMeta, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(newMeta, &meta)
		if err != nil {
			return nil, err
		}
	}

	return meta, nil
}

type Impl struct {
	userClient       userPB.UserClient
	dataKeeperClient datakeeperPB.DataKeeperClient
	token            string
	reader           *bufio.Reader

	ctx context.Context
}

type LogIn struct {
	Login    string
	Password string
}

type SignIn struct {
	Login    string
	Password string
	Email    string
	Secret   string
}

type Interface interface {
	LogIn(context.Context, *LogIn) error
	SignIn(context.Context, *SignIn) error
	Start(reader *bufio.Reader, ctx context.Context)
}

func (s *Impl) LogIn(ctx context.Context, req *LogIn) error {
	resp, err := s.userClient.LogIn(ctx, &userPB.LogInRequest{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	s.token = resp.Token
	return nil
}
func (s *Impl) SignIn(ctx context.Context, req *SignIn) error {
	resp, err := s.userClient.SignIn(ctx, &userPB.SignInRequest{
		Login:    req.Login,
		Password: req.Password,
		Email:    req.Email,
		Secret:   req.Secret,
	})
	if err != nil {
		return err
	}

	s.token = resp.Token
	return nil
}

func (s *Impl) new() {
	for {
		fmt.Println("Exit - q")
		fmt.Print("Enter DataType (LOG_PASS - lp, TEXT - t, BANK_CARD - bc, OTHER - o) : ")
		dataType, _ := s.reader.ReadString('\n')
		dataType = strings.TrimSpace(dataType)
		if dataType == "q" {
			return
		}

		req := &datakeeperPB.NewRequest{}

		switch dataType {
		case "lp":
			req.DataType = datakeeperPB.DataType_LOG_PASS
		case "t":
			req.DataType = datakeeperPB.DataType_TEXT
		case "bc":
			req.DataType = datakeeperPB.DataType_BANK_CARD
		case "o":
			req.DataType = datakeeperPB.DataType_OTHER
		default:
			fmt.Println("Unknown DataType. Try again.")
			continue
		}

		var err error

		req.Data, err = getData(s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		req.Meta, err = getMeta(s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		resp, err := s.dataKeeperClient.New(s.ctx, req)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		fmt.Println("Success, New Id", resp.Id)
	}
}

func (s *Impl) get() {
	for {
		fmt.Println("Exit - q")
		fmt.Print("Enter item Id : ")
		id, _ := s.reader.ReadString('\n')
		id = strings.TrimSpace(id)
		if id == "q" {
			return
		}

		resp, err := s.dataKeeperClient.Get(s.ctx, &datakeeperPB.GetRequest{
			Id: id,
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		err = readData(resp.Data, s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}
	}
}

func (s *Impl) set() {
	for {
		fmt.Println("Exit - q")
		fmt.Print("Enter item Id : ")
		id, _ := s.reader.ReadString('\n')
		id = strings.TrimSpace(id)
		if id == "q" {
			return
		}

		prevData, err := s.dataKeeperClient.Get(s.ctx, &datakeeperPB.GetRequest{
			Id: id,
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		newData := &datakeeperPB.SetRequest{
			Id:   id,
			Data: prevData.Data.Data,
			Meta: prevData.Data.Meta,
		}

		newData.Data, err = getData(s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		newData.Meta, err = getMeta(s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		_, err = s.dataKeeperClient.Set(s.ctx, newData)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		fmt.Println("success")
	}
}

func (s *Impl) list() {
	offset := int32(0)
	limit := int32(5)

	for {
		resp, err := s.dataKeeperClient.GetBatch(s.ctx, &datakeeperPB.GetBatchRequest{
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		for _, item := range resp.DataArray {
			err := readData(item, s.reader)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if len(resp.DataArray) == 0 {
			fmt.Println("End")
			return
		}

		fmt.Println("Exit - q")
		fmt.Println("Next - any other symbol")
		command, _ := s.reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "q" {
			return
		}

		offset += limit
	}
}

func (s *Impl) delete() {
	for {
		fmt.Println("Exit - q")
		fmt.Print("Enter item Id : ")
		id, _ := s.reader.ReadString('\n')
		id = strings.TrimSpace(id)
		if id == "q" {
			return
		}

		_, err := s.dataKeeperClient.Delete(s.ctx, &datakeeperPB.DeleteRequest{
			Id: id,
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Try again")
			continue
		}

		fmt.Println("Item deleted")
		return
	}
}

func (s *Impl) Start(reader *bufio.Reader, ctx context.Context) {
	s.reader = reader

	s.ctx = metadata.AppendToOutgoingContext(
		ctx,
		"jwt",
		s.token,
	)

	fmt.Println()

	for {
		fmt.Print("Next command.\n\tExit - q\n\tNew - n\n\tGet - g\n\tSet - s\n\tList - l\n\tDelete - d\n: ")
		command, _ := s.reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "q":
			return
		case "n":
			s.new()
		case "g":
			s.get()
		case "s":
			s.set()
		case "d":
			s.delete()
		case "l":
			s.list()
		default:
			fmt.Println("Unknown command. Try again.")
		}
	}
}

var _ Interface = (*Impl)(nil)

// "localhost:5678"
func Init(userServiceAddr, serviceAddr string) (Interface, error) {
	userConn, err := grpc.Dial(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(serviceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Impl{
		userClient:       userPB.NewUserClient(userConn),
		dataKeeperClient: datakeeperPB.NewDataKeeperClient(conn),
	}, nil
}
