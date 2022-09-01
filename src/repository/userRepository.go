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

func (repository Users) UpdateUser(id uint64, user model.User) error {
	stmt, err := repository.db.Prepare("UPDATE users SET name = ?, nickname = ?, email = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.Nickname, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (repository Users) DeleteUser(id uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository Users) GetUserByEmail(email string) (model.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, nickname, email, password, createdAt FROM users WHERE email = ?",
		email,
	)

	if err != nil {
		return model.User{}, err
	}

	defer rows.Close()

	var user model.User

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

func (repository Users) FollowUser(userId, followerId uint64) error {
	stmt, err := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository Users) UnfollowUser(userId, followerId uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository Users) GetUserFollowers(userId uint64) ([]model.User, error) {
	rows, err := repository.db.Query(
		"SELECT u.id, u.name, u.nickname, u.email FROM followers as f JOIN users as u ON f.follower_id = u.id WHERE f.user_id = ?",
		userId,
	)

	if err != nil {
		return nil, err
	}

	var users []model.User

	for rows.Next() {
		var user model.User

		if err = rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetUserFollowing(userId uint64) ([]model.User, error) {
	rows, err := repository.db.Query(
		"SELECT u.id, u.name, u.nickname, u.email FROM followers as f JOIN users as u ON f.user_id = u.id WHERE f.follower_id = ?",
		userId,
	)

	if err != nil {
		return nil, err
	}

	var users []model.User

	for rows.Next() {
		var user model.User

		if err = rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) UpdatePassword(passwordHash string, userId uint64) error {
	stmt, err := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(passwordHash, userId); err != nil {
		return err
	}

	return nil
}

func (repository Users) GetUserPasswordHash(id uint64) (string, error) {
	rows, err := repository.db.Query("SELECT password FROM users WHERE id = ?", id)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var user model.User

	if rows.Next() {
		if err = rows.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}
