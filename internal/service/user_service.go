package service

import "sync"

type UserService interface {
	GetUserById(id string)
}

type UserServiceImpl struct {
}

var instanceUserService *UserServiceImpl
var userServiceOnce sync.Once

func NewUserService() UserService {
	userServiceOnce.Do(func() {
		instanceUserService = &UserServiceImpl{}
	})
	return instanceUserService
}

func (r *UserServiceImpl) GetUserById(id string) {
	// Implement logic here
}
