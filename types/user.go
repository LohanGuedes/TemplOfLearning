package types

import (
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Something() string {
	return fmt.Sprintf("Meu nome é: %s", u.Name)
}
