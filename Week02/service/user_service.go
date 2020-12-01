package service

import "Week02/dao"

type UserService struct {
	dao *dao.User
}

func NewUserService() *UserService {
	return &UserService{dao: dao.NewUser()}
}

func (this *UserService) GetUser(id int) (*dao.User, error) {
	return this.dao.GetUserById(id)
}
