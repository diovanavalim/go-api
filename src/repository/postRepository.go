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

func (repository Posts) GetPosts(userId uint64) ([]dto.PostDto, error) {
	rows, err := repository.db.Query(`SELECT DISTINCT
    				p.*, u.nickname FROM posts p
        			INNER JOIN
    				users u ON u.id = p.author_id
        			LEFT JOIN
    				followers f ON p.author_id = f.user_id
					WHERE u.id = ? OR f.follower_id = ?
					ORDER BY 1 DESC;`, userId, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []dto.PostDto

	for rows.Next() {
		var post dto.PostDto

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
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

func (repository Posts) UpdatePost(id uint64, post model.Post) error {
	stmt, err := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(post.Title, post.Content, id); err != nil {
		return err
	}

	return nil
}

func (repository Posts) DeletePost(id uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM posts WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository Posts) GetUserPosts(id uint64) ([]dto.PostDto, error) {
	rows, err := repository.db.Query(
		"SELECT p.*, u.nickname from posts as p INNER JOIN users as u ON p.author_id = u.id WHERE p.author_id = ?", id,
	)

	if err != nil {
		return nil, err
	}

	var posts []dto.PostDto

	for rows.Next() {
		var post dto.PostDto

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) LikePost(id uint64) error {
	return nil
}

func (repository Posts) UnlikePost(id uint64) error {
	return nil
}
