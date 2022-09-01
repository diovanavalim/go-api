package model

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Post struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	AuthorID  uint64    `json:"authorId,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (post Post) Validate(stage string) error {
	v := reflect.ValueOf(post)

	typeOfPost := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if strings.TrimSpace(fmt.Sprint(v.Field(i).Interface())) == "" {
			return errors.New(fmt.Sprintf("Field %s can not be empty", typeOfPost.Field(i).Name))
		}
	}

	return nil
}
