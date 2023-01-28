package main

import (
	"context"
	"errors"

	"github.com/GermanVor/data-keeper/cmd/storageServer/service"
	"github.com/GermanVor/data-keeper/cmd/storageServer/storage"
)

type StorageMock struct{}

func (s *StorageMock) New(ctx context.Context, newData *storage.NewData) (string, error) {
	return "", nil
}

func (s *StorageMock) Get(ctx context.Context, getData *storage.GetData) (*storage.Data, error) {
	return nil, errors.New("")
}

func (s *StorageMock) Set(ctx context.Context, setData *storage.SetData) (*storage.Data, error) {
	return nil, errors.New("")
}

func (s *StorageMock) Delete(ctx context.Context, deleteData *storage.DeleteData) (bool, error) {
	return false, nil
}

var _ storage.Interface = (*StorageMock)(nil)

func main() {
	s := service.Init(":1234", &StorageMock{}, ":5678")

	s.Start()
}
