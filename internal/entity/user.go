package entity

import "fmt"

type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

func (u *User) ValidateUserByRegistration() error {
	if u.FirstName == "" || u.LastName == "" || u.Email == "" || u.Password == "" {
		return fmt.Errorf("Необходимо заполнить все обязательные поля")
	}

	return nil
}

func (u *User) ValidateUserByLogin() error {
	if u.Email == "" || u.Password == "" {
		return fmt.Errorf("Необходимо заполнить все обязательные поля")
	}

	return nil
}
