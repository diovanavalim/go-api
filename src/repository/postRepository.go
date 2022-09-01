package repository

import (
	"api/src/dto"
	"api/src/model"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func CreatePostRepository(db *sql.DB) *Posts {
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

func (repository Posts) GetPost(id uint64) (dto.PostDto, error) {
	row, err := repository.db.Query(
		"SELECT p.*, u.nickname from posts as p INNER JOIN users as u ON p.author_id = u.id WHERE p.id = ?", id)

	if err != nil {
		return dto.PostDto{}, err
	}

	defer row.Close()

	var post dto.PostDto

	if row.Next() {
		if err := row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return dto.PostDto{}, err
		}
	}

	return post, nil
}

func (repository Posts) UpdatePost(id uint64) error {
	return nil
}

func (repository Posts) DeletePost(id uint64) error {
	return nil
}
