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
	db, err := sql.Open("mysql", "ca:mission@/ca_mission")
	if err != nil {
		println("sql open error")
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	insert, err := db.Prepare("INSERT INTO users (id, name) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(user.Id, user.Name)

	return err
}

func (repository *UserRepository) GetByUserId(userId string) (*model.User, error) {
	db, err := sql.Open("mysql", "ca:mission@/ca_mission")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	user := model.User{}

	err = db.QueryRow("SELECT * FROM users WHERE id = ?", userId).Scan(&user.Id, &user.Name)
	return &user, err
}

func (repository *UserRepository) Update(updatedUser model.User) error {
	db, err := sql.Open("mysql", "ca:mission@/ca_mission")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	update, err := db.Prepare("UPDATE users set name = ? where id = ? ")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(updatedUser.Name, updatedUser.Id)

	return err
}
