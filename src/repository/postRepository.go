package repository

import (
	"api/src/model"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func CreateRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) CreatePost(post model.Post, authorId uint64) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(post.Title, post.Content, authorId)

	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (repository Posts) GetPosts(userId uint64) ([]model.Post, error) {
	return nil, nil
}

func (repository Posts) GetPost(id uint64) (model.Post, error) {
	return model.Post{}, nil
}

func (repository Posts) UpdatePost(id uint64) error {
	return nil
}

func (repository Posts) DeletePost(id uint64) error {
	return nil
}
