package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
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

func (repository Users) GetUsers() ([]model.User, error) {
	rows, err := repository.db.Query("SELECT id, name, nickname, email, createdAt FROM users")

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var users []model.User

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetUsersByNameOrNickname(nameOrNickname string) ([]model.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

	rows, err := repository.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE name LIKE ? or nickname LIKE ?",
		nameOrNickname, nameOrNickname,
	)
	if err != nil {
		return nil, err
	}

	var users []model.User

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetUserByID(id uint64) (model.User, error) {
	rows, err := repository.db.Query("SELECT id, name, nickname, email, createdAt FROM users WHERE id = ?", id)

	if err != nil {
		return model.User{}, err
	}

	defer rows.Close()

	var user model.User

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}
