package repository

import (
	"github.com/kimxuanhong/user-manager-go/pkg/core/dto"
)

type UserRepository interface {
	FindById(id string) (*dto.UserDto, error)
	UpdateStatus(user *dto.UserDto) (*dto.UserDto, error)
}
