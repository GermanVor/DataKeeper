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

type UserOutput struct {
	UserID string
	Secret string
}

type Interface interface {
	SignIn(ctx context.Context, userData *UserData) (*UserOutput, error)
	LogIn(ctx context.Context, login, password string) (*UserOutput, error)
	GetSecret(ctx context.Context, userID string) (string, error)
}
