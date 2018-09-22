package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"
)

const passwordSalt = "test"

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	row := db.QueryRow(
		`SELECT id, email, firstname, lastname
		FROM public.User
		WHERE email = $1
		AND password = $2`, email, pwd)
	fmt.Println("pwd: ", pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	return result, nil
}
