package dao

import (
	"database/sql"
	"errors"
	xerrors "github.com/pkg/errors"
)

var ErrNotData error = errors.New("not exist")

type User struct {
	Id   int
	Name string
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserById(id int) (*User, error) {
	if id < 100 {
		//不存在数据
		err := sql.ErrNoRows
		if err == sql.ErrNoRows {
			return nil, xerrors.Wrap(ErrNotData, "user not exist")
		}
	}
	return &User{1, "青青河边草"}, nil
}
