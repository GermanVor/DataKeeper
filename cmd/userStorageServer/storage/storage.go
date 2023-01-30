package storage

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"strconv"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/scrypt"
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

func getLogin(login string) string {
	loginHash := sha1.Sum([]byte(login))
	return hex.EncodeToString(loginHash[:])
}

func getPass(pass string, salt []byte) (string, error) {
	passHash, err := scrypt.Key([]byte(pass), salt, 1<<14, 8, 1, 64)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(passHash), nil
}

func getSalt() ([]byte, error) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

var (
	// UPDATE users SET (login, pass, salt, email) = ('a', 'b', 'c', 'd') WHERE secret='test';
	setUserSQL = "UPDATE users SET (login, pass, salt, email) = " +
		"($2, $3, $4, $5) WHERE secret=$1 RETURNING users.userID"

	// SELECT secret, userID FROM users WHERE login=$1 and pass=$2
	loginSQL = "SELECT secret, userID FROM users WHERE login=$1 and pass=$2"

	// SELECT salt FROM users WHERE login=$1;
	getSaltSQL = "SELECT salt FROM users WHERE login=$1"

	// SELECT login FROM users WHERE secret=$1;
	getLogginSQL = "SELECT login FROM users WHERE secret=$1"

	// SELECT secret FROM users WHERE userID=$1;
	getSecretSQL = "SELECT secret FROM users WHERE userID=$1"

	// Create table
	createTableSQL = "CREATE TABLE IF NOT EXISTS users (" +
		"login text UNIQUE, " +
		"pass text, " +
		"salt text, " +
		"secret text UNIQUE," +
		"email text, " +
		"userID SERIAL " +
		");"
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

func Init(databaseURI string) Interface {
	conn, err := pgxpool.Connect(context.TODO(), databaseURI)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = conn.Exec(context.TODO(), createTableSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Impl{
		dbPool: conn,
	}
}
