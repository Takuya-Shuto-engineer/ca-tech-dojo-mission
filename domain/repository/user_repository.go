package repository

import (
	"ca-mission/domain/model"
)

type UserRepository interface {
	// データベースに格納
	InsertDB(user model.User) error
	GetByUserIdDB(userId string) (*model.User, error)
	UpdateDB(updatedUser model.User) error

	// メモリに格納
	Insert(user model.User) error
	GetByUserId(userId string) (*model.User, error)
	Update(updatedUser model.User) error
}
