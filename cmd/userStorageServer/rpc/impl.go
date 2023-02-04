package userrpc

import (
	"context"
	"strings"

	"github.com/GermanVor/data-keeper/cmd/userStorageServer/storage"
	"github.com/GermanVor/data-keeper/internal/common"
	pb "github.com/GermanVor/data-keeper/proto/user"
	"github.com/golang-jwt/jwt/v4"
)

type UserRPCImpl struct {
	pb.UnimplementedUserServer
	stor storage.Interface
}

// TODO возвращать на ружу свои ошибки

func buildUserToke(userId, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		common.USER_ID_CTX_NAME: userId,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserRPCImpl) LogIn(ctx context.Context, in *pb.LogInRequest) (*pb.LogInResponse, error) {
	userOutput, err := s.stor.LogIn(ctx, in.Login, in.Password)
	if err != nil {
		return nil, err
	}

	tokenStr, err := buildUserToke(userOutput.UserID, userOutput.Secret)
	if err != nil {
		return nil, err
	}

	resp := &pb.LogInResponse{
		Token: tokenStr,
	}

	return resp, nil
}

func (s *UserRPCImpl) CheckAccess(ctx context.Context, in *pb.CheckAccessRequest) (*pb.CheckAccessResponse, error) {
	claim := jwt.MapClaims{}
	token, parts, err := jwt.NewParser().ParseUnverified(in.Token, claim)

	if err != nil {
		return nil, err
	}

	var userId string
	var ok bool
	if userId, ok = claim[common.USER_ID_CTX_NAME].(string); !ok {
		return nil, nil
	}

	secret, err := s.stor.GetSecret(ctx, userId)
	if err != nil {
		return nil, err
	}

	if err = token.Method.Verify(strings.Join(parts[0:2], "."), parts[2], []byte(secret)); err != nil {
		return nil, err
	}

	return &pb.CheckAccessResponse{
		UserId: userId,
	}, nil
}

func (s *UserRPCImpl) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	userOutput, err := s.stor.SignIn(ctx, &storage.UserData{
		Login:    in.Login,
		Password: in.Password,
		Email:    in.Email,
		Secret:   in.Secret,
	})
	if err != nil {
		return nil, err
	}

	tokenStr, err := buildUserToke(userOutput.UserID, userOutput.Secret)
	if err != nil {
		return nil, err
	}

	resp := &pb.SignInResponse{
		Token: tokenStr,
	}

	return resp, nil
}

func Init(stor storage.Interface) *UserRPCImpl {
	return &UserRPCImpl{
		stor: stor,
	}
}
