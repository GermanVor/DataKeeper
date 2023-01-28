package storage

import (
	"context"
)

type UserData struct {
	Login    string
	Password string
	Email    string
	Secret   string
}

type SignOutput struct {
	UserId string
	Secret string
}

type Interface interface {
	LogIn(ctx context.Context, userData *UserData) (string, error)
	SignIn(ctx context.Context, login, password string) (*SignOutput, error)
	GetSecret(ctx context.Context, userId string) (string, error)
}
