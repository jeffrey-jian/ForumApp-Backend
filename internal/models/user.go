package models

import (
	"fmt"
)

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	AvatarColor string `json:"avatarColor"`
}

func (user *User) Greet() string {
	return fmt.Sprintf("Hello, I am %s", user.Username)
}
