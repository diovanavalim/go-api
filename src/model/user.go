package model

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user User) Validate(stage string) error {
	v := reflect.ValueOf(user)

	typeOfUser := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if strings.TrimSpace(fmt.Sprint(v.Field(i).Interface())) == "" && typeOfUser.Field(i).Name != "Password" {
			return errors.New(fmt.Sprintf("Field %s can not be empty", typeOfUser.Field(i).Name))
		}
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Invalid email format")
	}

	if stage == "create" {
		if strings.TrimSpace(user.Password) == "" {
			return errors.New(fmt.Sprintf("Field %s can not be empty", "password"))
		}
	}

	return nil
}

func (user *User) Format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nickname = strings.TrimSpace(user.Nickname)
}
