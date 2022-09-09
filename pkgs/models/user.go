package models

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int64  `json:"userId"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

var ErrUserEmailNotFound = errors.New("email not found")
var ErrUserPasswordInvalid = errors.New("invalid password")

func InsertUser(ctx context.Context, email string, password string) (int64, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	res, err := db.ExecContext(ctx, "INSERT INTO user (email, password) VALUES (?,?);", email, passwordHash)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func SelectUserByEmail(ctx context.Context, email string, password string) (User, error) {
	var user User
	err := db.QueryRowContext(ctx, "SELECT userId, email, password FROM user WHERE email = ?;", email).Scan(&user.Id, &user.Email, &user.PasswordHash)
	switch {
	case err == sql.ErrNoRows:
		return user, ErrUserEmailNotFound
	case err != nil:
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	switch {
	case err == bcrypt.ErrMismatchedHashAndPassword:
		return user, ErrUserPasswordInvalid
	case err != nil:
		return user, err
	}

	return user, nil
}
