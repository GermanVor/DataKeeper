package tests

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	"github.com/GermanVor/data-keeper/internal/user"
	storage "github.com/GermanVor/data-keeper/internal/userStorage"
	pb "github.com/GermanVor/data-keeper/proto/user"
	"github.com/bmizerany/assert"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

const USER_LOGIN = "USER_LOGIN_QWERTY"
const USER_ID = "USER_ID_QWERTY"
const SECRET = "SECRET_YTREWQ"

var TOKEN, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"userId": USER_ID,
}).SignedString([]byte(SECRET))

type StorageMock struct{}

func (s *StorageMock) LogIn(ctx context.Context, userData *storage.UserData) (string, error) {
	if userData.Login == USER_LOGIN {
		return USER_ID, nil
	}

	return "", errors.New("")
}

func (s *StorageMock) GetSecret(ctx context.Context, userId string) (string, error) {
	if userId == USER_ID {
		return SECRET, nil
	}

	return "", errors.New("")
}

func (s *StorageMock) SignIn(ctx context.Context, login, password string) (*storage.SignOutput, error) {
	if login == USER_LOGIN {
		return &storage.SignOutput{
			UserId: USER_ID,
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

	pb.RegisterUserServer(s, user.Init(stor))

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
			Login:  USER_LOGIN,
			Secret: SECRET,
		})

		require.NoError(t, err)
		assert.Equal(t, TOKEN, resp.Token)
	})

	t.Run("SignIn", func(t *testing.T) {
		resp, err := client.SignIn(ctx, &pb.SignInRequest{
			Login:    USER_LOGIN,
			Password: "",
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
			Login:  USER_LOGIN + "qwe",
			Secret: "",
		})

		assert.Equal(t, (*pb.LogInResponse)(nil), resp)
	})

	t.Run("Negative SignIn", func(t *testing.T) {
		resp, _ := client.SignIn(ctx, &pb.SignInRequest{
			Login:    USER_LOGIN + "qwe",
			Password: "",
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
