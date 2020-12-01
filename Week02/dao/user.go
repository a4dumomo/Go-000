package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	Id   int
	Name string
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserById(id int) (*User, error) {
	if id < 100 {
		return nil, errors.Wrap(sql.ErrNoRows, "user not exitst")
	} else {
		return &User{1, "青青河边草"}, nil
	}
}
