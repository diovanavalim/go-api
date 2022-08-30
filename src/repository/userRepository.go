package repository

import (
	"api/src/model"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func CreateUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) CreateUser(user model.User) (uint64, error) {
	stmt, err := repository.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)",
	)

	defer stmt.Close()

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(user.Name, user.Nickname, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
