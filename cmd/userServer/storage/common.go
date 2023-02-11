package storage

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/scrypt"
)

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

// Returns a hash of the login to write to the database
func getLogin(login string) string {
	loginHash := sha1.Sum([]byte(login))
	return hex.EncodeToString(loginHash[:])
}

// Returns a hash of the salted the password to write to the database
func getPass(pass string, salt []byte) (string, error) {
	passHash, err := scrypt.Key([]byte(pass), salt, 1<<14, 8, 1, 64)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(passHash), nil
}

// Returns random salt
func getSalt() ([]byte, error) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}
