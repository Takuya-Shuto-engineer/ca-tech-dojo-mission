package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"ca-mission/domain/model"
	"ca-mission/domain/repository"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() (repository.UserRepository, error) {
	db, err := sql.Open("mysql", "ca:mission@/ca_mission")
	return &UserRepository{DB: db}, err
}

func (repository *UserRepository) DBClose() error {
	err := repository.DB.Close()
	return err
}

func (repository *UserRepository) Insert(user model.User) error {
	// databaseにuserを新規登録する
	insert, err := repository.DB.Prepare("INSERT INTO users (id, name) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(user.Id, user.Name)

	return err
}

func (repository *UserRepository) GetByUserId(userId string) (*model.User, error) {
	user := model.User{}

	err := repository.DB.QueryRow("SELECT * FROM users WHERE id = ?", userId).Scan(&user.Id, &user.Name)
	return &user, err
}

func (repository *UserRepository) Update(updatedUser model.User) error {
	update, err := repository.DB.Prepare("UPDATE users set name = ? where id = ? ")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(updatedUser.Name, updatedUser.Id)

	return err
}
