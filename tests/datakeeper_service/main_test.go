package tests

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	"github.com/GermanVor/data-keeper/internal/datakeeper"
	storage "github.com/GermanVor/data-keeper/internal/storage"
	pb "github.com/GermanVor/data-keeper/proto/datakeeper"
	"github.com/bmizerany/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var (
	USER_ID   = "USER_ID_QWE"
	ID        = "ID"
	DATA_TYPE = pb.DataType_TEXT
	DATA      = []byte("qwerty")
)

const bufSize = 1024 * 1024

type StorageMock struct{}

func (s *StorageMock) New(ctx context.Context, newData *storage.NewData) (string, error) {
	if newData.UserId == USER_ID {
		return ID, nil
	}

	return "", errors.New("")
}

func (s *StorageMock) Get(ctx context.Context, getData *storage.GetData) (*storage.Data, error) {
	if getData.UserId == USER_ID && getData.Id == ID {
		return &storage.Data{
			Id:       ID,
			DataType: DATA_TYPE.Format(),
			Data:     DATA,
			Meta:     make(map[string]string),
		}, nil
	}

	return nil, errors.New("")
}

func (s *StorageMock) Set(ctx context.Context, setData *storage.SetData) (*storage.Data, error) {
	if setData.UserId == USER_ID && setData.Id == ID {
		return &storage.Data{
			Id:       ID,
			DataType: DATA_TYPE.Format(),
			Data:     setData.Data,
			Meta:     setData.Meta,
		}, nil
	}

	return nil, errors.New("")
}

func (s *StorageMock) Delete(ctx context.Context, deleteData *storage.DeleteData) (bool, error) {
	if deleteData.UserId == USER_ID && deleteData.Id == ID {
		return true, nil
	}

	return false, nil
}

var _ storage.Interface = (*StorageMock)(nil)

var stor = &StorageMock{}

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	pb.RegisterDataKeeperServer(s, datakeeper.Init(stor))

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
	client := pb.NewDataKeeperClient(conn)

	t.Run("New", func(t *testing.T) {
		resp, err := client.New(ctx, &pb.NewRequest{
			UserId:   USER_ID,
			DataType: DATA_TYPE,
			Data:     DATA,
			Meta:     make(map[string]string),
		})

		require.NoError(t, err)
		assert.Equal(t, ID, resp.Id)
	})

	t.Run("Get", func(t *testing.T) {
		resp, err := client.Get(ctx, &pb.GetRequest{
			UserId: USER_ID,
			Id:     ID,
		})

		require.NoError(t, err)
		assert.Equal(t, ID, resp.Id)
		assert.Equal(t, DATA, resp.Data)
		assert.Equal(t, DATA_TYPE, resp.DataType)
	})

	t.Run("Set", func(t *testing.T) {
		newData := []byte("ytrewq")
		newMeta := map[string]string{"qwe": "rty"}

		resp, err := client.Set(ctx, &pb.SetRequest{
			UserId: USER_ID,
			Id:     ID,
			Data:   newData,
			Meta:   newMeta,
		})

		require.NoError(t, err)
		assert.Equal(t, newData, resp.Data)
		assert.Equal(t, ID, resp.Id)
		assert.Equal(t, DATA_TYPE, resp.DataType)
		assert.Equal(t, newMeta, resp.Meta)
	})
}
