package main

import (
	"context"
	"errors"

	"github.com/GermanVor/data-keeper/cmd/userStorageServer/service"
	"github.com/GermanVor/data-keeper/cmd/userStorageServer/storage"
)

type StorageMock struct{}

func (s *StorageMock) LogIn(ctx context.Context, userData *storage.UserData) (string, error) {
	return "", errors.New("")
}

func (s *StorageMock) GetSecret(ctx context.Context, userId string) (string, error) {
	return "", errors.New("")
}

func (s *StorageMock) SignIn(ctx context.Context, login, password string) (*storage.SignOutput, error) {
	return nil, errors.New("")
}

var _ storage.Interface = (*StorageMock)(nil)

func main() {
	s := service.Init(":1234", &StorageMock{})

	s.Start()
}
