package models

import (
	"errors"
	"hms/database"

	"github.com/google/uuid"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(email, password string) (*User, error) {
	// basic validation. robust validation to be done later
	if email == "" || password == "" {
		err := errors.New("email and password are required")
		return nil, err
	}
	user := User{
		Email:    email,
		Password: password,
	}
	return &user, nil
}

func (u *User) Save() error {
	db := database.GetDB()
	defer db.Close()

	id := uuid.New().String()

	query := `INSERT INTO users (id, email, password) VALUES (?,?,?);`

	_, err := db.Exec(query, id, u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}
