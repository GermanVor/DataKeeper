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

	"github.com/GermanVor/data-keeper/internal/common"
	datakeeperPB "github.com/GermanVor/data-keeper/proto/datakeeper"
	userPB "github.com/GermanVor/data-keeper/proto/user"
)

var (
	tryAgainStr         = "Try again."
	enterPathToSaveData = "Enter path to save Data : "

	quitallComm = "q"
	exitStr     = "Exit - " + quitallComm

	newComm = "n"
	newStr  = "New - " + newComm

	getComm = "g"
	getStr  = "Get - " + getComm

	setComm = "s"
	setStr  = "Set - " + setComm

	listComm = "l"
	listStr  = "List - " + listComm

	deleteComm = "d"
	deleteStr  = "Delete - " + deleteComm

	nextCommandStr = "Next command.\n\t" +
		exitStr + "\n\t" +
		newStr + "\n\t" +
		getStr + "\n\t" +
		setStr + "\n\t" +
		listStr + "\n\t" +
		deleteStr + "\n: "

	valueStr = "Value : "
	metaStr  = "Meta : "

	enterItemIdStr = "Enter item Id : "

	unknownCommandStr = "Unknown command."

	itemDeletedStr = "Item deleted"

	endStr          = "End"
	nextAnyOtherStr = "Next - any other symbol"

	successStr = "Success"

	logPassComm      = "lp"
	textComm         = "t"
	bankCardComm     = "bc"
	otherComm        = "o"
	enterDataTypeStr = "Enter DataType (" +
		"LOG_PASS - " + logPassComm + ", " +
		"TEXT - " + textComm + ", " +
		"BANK_CARD - " + bankCardComm + ", " +
		"OTHER - " + otherComm +
		") : "
	unknownDataTypeStr = "Unknown DataType."
	newIdStr           = "New Id."
)

func readData(data *datakeeperPB.Data, reader *bufio.Reader) error {
	fmt.Print(valueStr)

	switch data.DataType {
	case datakeeperPB.DataType_LOG_PASS,
		datakeeperPB.DataType_TEXT,
		datakeeperPB.DataType_BANK_CARD:
		fmt.Println(string(data.Data[:]))
	default:
		for {
			fmt.Print(enterPathToSaveData)
			filePath, _ := reader.ReadString('\n')
			filePath = strings.TrimSpace(filePath)

			err := ioutil.WriteFile(filePath, data.Data, 0644)
			if err != nil {
				return err
			}
		}
	}

	fmt.Print(metaStr)
	fmt.Println(data.Meta)

	return nil
}

const enterPathToNewData = "Enter path to new Data (empty string to skip) : "

func getData(reader *bufio.Reader) ([]byte, error) {
	fmt.Println(enterPathToNewData)
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	var data []byte
	var err error

	if filePath != "" {
		data, err = ioutil.ReadFile(filePath)
	}

	return data, err
}

const enterPathToNewMeta = "Enter path to new Meta (empty string to skip) : "

func getMeta(reader *bufio.Reader) (map[string]string, error) {
	fmt.Println(enterPathToNewMeta)
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

func (s *Impl) new(ctx context.Context) {
	for {
		fmt.Println(exitStr)
		fmt.Print(enterDataTypeStr)
		dataType, _ := s.reader.ReadString('\n')
		dataType = strings.TrimSpace(dataType)
		if dataType == quitallComm {
			return
		}

		req := &datakeeperPB.NewRequest{}

		switch dataType {
		case logPassComm:
			req.DataType = datakeeperPB.DataType_LOG_PASS
		case textComm:
			req.DataType = datakeeperPB.DataType_TEXT
		case bankCardComm:
			req.DataType = datakeeperPB.DataType_BANK_CARD
		case otherComm:
			req.DataType = datakeeperPB.DataType_OTHER
		default:
			fmt.Println(unknownDataTypeStr, tryAgainStr)
			continue
		}

		var err error

		req.Data, err = getData(s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}

		req.Meta, err = getMeta(s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}

		resp, err := s.dataKeeperClient.New(ctx, req)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}

		fmt.Println(successStr, newIdStr, resp.Id)
	}
}

func (s *Impl) get(ctx context.Context) {
	for {
		fmt.Println(exitStr)
		fmt.Print(enterItemIdStr)
		id, _ := s.reader.ReadString('\n')
		id = strings.TrimSpace(id)
		if id == quitallComm {
			return
		}

		resp, err := s.dataKeeperClient.Get(ctx, &datakeeperPB.GetRequest{
			Id: id,
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}

		err = readData(resp.Data, s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}
	}
}

func (s *Impl) set(ctx context.Context) {
	for {
		fmt.Println(exitStr)
		fmt.Print(enterItemIdStr)
		id, _ := s.reader.ReadString('\n')
		id = strings.TrimSpace(id)
		if id == quitallComm {
			return
		}

		prevData, err := s.dataKeeperClient.Get(ctx, &datakeeperPB.GetRequest{
			Id: id,
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
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
			fmt.Println(tryAgainStr)
			continue
		}

		newData.Meta, err = getMeta(s.reader)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}

		_, err = s.dataKeeperClient.Set(ctx, newData)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}

		fmt.Println(successStr)
	}
}

func (s *Impl) list(ctx context.Context) {
	offset := int32(0)
	limit := int32(5)

	for {
		resp, err := s.dataKeeperClient.GetBatch(ctx, &datakeeperPB.GetBatchRequest{
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}

		for _, item := range resp.DataArray {
			err := readData(item, s.reader)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if len(resp.DataArray) == 0 {
			fmt.Println(endStr)
			return
		}

		fmt.Println(exitStr)
		fmt.Println(nextAnyOtherStr)
		command, _ := s.reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == quitallComm {
			return
		}

		offset += limit
	}
}

func (s *Impl) delete(ctx context.Context) {
	for {
		fmt.Println(exitStr)
		fmt.Print(enterItemIdStr)
		id, _ := s.reader.ReadString('\n')
		id = strings.TrimSpace(id)
		if id == quitallComm {
			return
		}

		_, err := s.dataKeeperClient.Delete(ctx, &datakeeperPB.DeleteRequest{
			Id: id,
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(tryAgainStr)
			continue
		}

		fmt.Println(itemDeletedStr)
		return
	}
}

func (s *Impl) Start(reader *bufio.Reader, ctx context.Context) {
	s.reader = reader

	currentCtx := metadata.AppendToOutgoingContext(
		ctx,
		common.JWT_CTX_NAME,
		s.token,
	)

	fmt.Println()

	for {
		fmt.Print(nextCommandStr)
		command, _ := s.reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case quitallComm:
			return
		case newComm:
			s.new(currentCtx)
		case getComm:
			s.get(currentCtx)
		case setComm:
			s.set(currentCtx)
		case deleteComm:
			s.delete(currentCtx)
		case listComm:
			s.list(currentCtx)
		default:
			fmt.Println(unknownCommandStr, tryAgainStr)
		}
	}
}

var _ Interface = (*Impl)(nil)

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
