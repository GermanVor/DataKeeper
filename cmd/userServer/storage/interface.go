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

type Storager interface {
	// Registration of user. User should have secret
	SignIn(ctx context.Context, userData *UserData) (*UserOutput, error)

	// Autorisation
	LogIn(ctx context.Context, login, password string) (*UserOutput, error)

	// Exchange userId on user secret to verify user JWT token
	GetSecret(ctx context.Context, userID string) (string, error)
}
