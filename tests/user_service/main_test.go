package tests

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	userRPC "github.com/GermanVor/data-keeper/cmd/userStorageServer/rpc"
	"github.com/GermanVor/data-keeper/cmd/userStorageServer/storage"
	"github.com/GermanVor/data-keeper/internal/common"
	pb "github.com/GermanVor/data-keeper/proto/user"
	"github.com/bmizerany/assert"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize    = 1024 * 1024
	USER_LOGIN = "USER_LOGIN_QWERTY"
	USER_ID    = "USER_ID_QWERTY"
	SECRET     = "SECRET_YTREWQ"
	PASSWORD   = "PASSWORD"
)

var TOKEN, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	common.USER_ID_CTX_NAME: USER_ID,
}).SignedString([]byte(SECRET))

type StorageMock struct{}

func (s *StorageMock) LogIn(ctx context.Context, login, password string) (*storage.UserOutput, error) {
	if login == USER_LOGIN {
		return &storage.UserOutput{
			UserID: USER_ID,
			Secret: SECRET,
		}, nil
	}

	return nil, errors.New("")
}

func (s *StorageMock) GetSecret(ctx context.Context, userId string) (string, error) {
	if userId == USER_ID {
		return SECRET, nil
	}

	return "", errors.New("")
}

func (s *StorageMock) SignIn(ctx context.Context, userData *storage.UserData) (*storage.UserOutput, error) {
	if userData.Login == USER_LOGIN && userData.Secret == SECRET {
		return &storage.UserOutput{
			UserID: USER_ID,
			Secret: SECRET,
		}, nil
	}

	return nil, errors.New("")
}

var _ storage.Interface = (*StorageMock)(nil)

var stor = &StorageMock{}

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	pb.RegisterUserServer(s, userRPC.Init(stor))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestBase(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)

	t.Run("LogIn", func(t *testing.T) {
		resp, err := client.LogIn(ctx, &pb.LogInRequest{
			Login:    USER_LOGIN,
			Password: PASSWORD,
		})

		require.NoError(t, err)
		assert.Equal(t, TOKEN, resp.Token)
	})

	t.Run("SignIn", func(t *testing.T) {
		resp, err := client.SignIn(ctx, &pb.SignInRequest{
			Login:    USER_LOGIN,
			Password: PASSWORD,
			Email:    "",
			Secret:   SECRET,
		})

		require.NoError(t, err)
		assert.Equal(t, TOKEN, resp.Token)
	})

	t.Run("CheckAccess", func(t *testing.T) {
		resp, err := client.CheckAccess(ctx, &pb.CheckAccessRequest{
			Token: TOKEN,
		})

		require.NoError(t, err)
		assert.Equal(t, USER_ID, resp.UserId)
	})
}

func TestNegative(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)

	t.Run("Negative LogIn", func(t *testing.T) {
		resp, _ := client.LogIn(ctx, &pb.LogInRequest{
			Login:    USER_LOGIN + "qwe",
			Password: PASSWORD,
		})

		assert.Equal(t, (*pb.LogInResponse)(nil), resp)
	})

	t.Run("Negative SignIn", func(t *testing.T) {
		resp, _ := client.SignIn(ctx, &pb.SignInRequest{
			Login:    USER_LOGIN,
			Password: PASSWORD + "qwe",
		})

		assert.Equal(t, (*pb.SignInResponse)(nil), resp)
	})

	t.Run("Negative CheckAccess", func(t *testing.T) {
		resp, _ := client.CheckAccess(ctx, &pb.CheckAccessRequest{
			Token: "qwe",
		})

		assert.Equal(t, (*pb.CheckAccessResponse)(nil), resp)
	})
}
