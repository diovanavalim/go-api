package model

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user User) Validate() error {
	v := reflect.ValueOf(user)

	typeOfUser := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if strings.TrimSpace(fmt.Sprint(v.Field(i).Interface())) == "" {
			return errors.New(fmt.Sprintf("Field %s can not be empty", typeOfUser.Field(i).Name))
		}
	}

	return nil
}

func (user *User) Format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nickname = strings.TrimSpace(user.Nickname)
}
