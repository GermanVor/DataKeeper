package storage

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/GermanVor/data-keeper/internal/common"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Impl struct {
	dbPool *pgxpool.Pool
}

func (s *Impl) SignIn(ctx context.Context, userData *UserData) (*UserOutput, error) {
	var loggin pgtype.Text
	err := s.dbPool.QueryRow(ctx, getLogginSQL, userData.Secret).Scan(&loggin)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("uknown secret")
		} else {
			return nil, err
		}
	}

	if loggin.Status != pgtype.Null {
		return nil, errors.New("user with this secret already registered")
	}

	salt, _ := getSalt()
	saltStr := hex.EncodeToString(salt)

	passStr, err := getPass(userData.Password, salt)
	if err != nil {
		return nil, err
	}

	userID := 0
	err = s.dbPool.QueryRow(
		ctx,
		setUserSQL,
		userData.Secret,
		getLogin(userData.Login),
		passStr,
		saltStr,
		userData.Email,
	).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return &UserOutput{
		UserID: strconv.Itoa(userID),
		Secret: userData.Secret,
	}, nil
}

func (s *Impl) LogIn(ctx context.Context, login, password string) (*UserOutput, error) {
	loginStr := getLogin(login)

	saltStr := ""
	err := s.dbPool.QueryRow(context.TODO(), getSaltSQL, loginStr).Scan(&saltStr)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("unknown user")
		} else {
			return nil, err
		}
	}

	salt, err := hex.DecodeString(saltStr)
	if err != nil {
		return nil, err
	}

	passStr, err := getPass(password, salt)
	if err != nil {
		return nil, err
	}

	userID := 0
	var secret string
	err = s.dbPool.QueryRow(ctx, loginSQL, loginStr, passStr).Scan(&secret, &userID)
	if err != nil {
		return nil, err
	}

	return &UserOutput{
		UserID: strconv.Itoa(userID),
		Secret: secret,
	}, nil
}

func (s *Impl) GetSecret(ctx context.Context, userID string) (string, error) {
	secret := ""
	err := s.dbPool.QueryRow(ctx, getSecretSQL, userID).Scan(&secret)
	if err != nil {
		return "", err
	}

	return secret, nil
}

var _ Interface = (*Impl)(nil)

func createDefaultUser(conn *pgxpool.Pool) {
	_, err := conn.Exec(context.TODO(), "INSERT INTO users (secret) VALUES ('"+common.DEFAULT_USER_SECRET+"')")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Init(databaseURI string) Interface {
	conn, err := pgxpool.Connect(context.TODO(), databaseURI)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = conn.Exec(context.TODO(), createTableSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	createDefaultUser(conn)

	return &Impl{
		dbPool: conn,
	}
}
