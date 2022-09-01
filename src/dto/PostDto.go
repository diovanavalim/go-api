package dto

import "time"

type PostDto struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"authorId,omitempty"`
	AuthorNickname string    `json:"nickname,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	Likes          uint64    `json:"likes"`
}
