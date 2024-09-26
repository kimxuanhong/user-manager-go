package repository

import (
	"github.com/kimxuanhong/user-manager-go/pkg/core/dto"
	"github.com/kimxuanhong/user-manager-go/pkg/core/repository"
	"sync"
)

type userRepository struct {
	Repository
}

var instanceUserRepository *userRepository
var userRepositoryOnce sync.Once

func NewUserRepository() repository.UserRepository {
	userRepositoryOnce.Do(func() {
		instanceUserRepository = &userRepository{}
	})
	return instanceUserRepository
}

func (r *userRepository) FindById(id string) (*dto.UserDto, error) {
	return &dto.UserDto{}, nil
}

func (r *userRepository) UpdateStatus(user *dto.UserDto) (*dto.UserDto, error) {
	return &dto.UserDto{}, nil
}
