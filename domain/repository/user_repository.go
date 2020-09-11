package repository

import (
	"ca-mission/domain/model"
)

type UserRepository interface {
	DBClose() error
	Insert(user model.User) error
	GetByUserId(userId string) (*model.User, error)
	Update(updatedUser model.User) error
}
