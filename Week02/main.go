package main

import (
	"database/sql"

	"github.com/pkg/errors"
)

func main() {
	UserServiceGetUser()
}

type User struct {
	ID       uint64
	Username string
}

// Service 层
func UserServiceGetUser() (*User, error) {
	id := uint64(192222211111)
	user, err := DaoGetUser(id)

	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

// Dao 层
func DaoGetUser(id uint64) (*User, error) {
	var user User
	if user, err := getFromDB(id); errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrapf(err, "User not found")
	}

	return &user, nil
}
