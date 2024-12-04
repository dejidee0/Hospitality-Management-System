package models

import (
	"errors"
	"fmt"
	"hms/database"

	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
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

	// hash password here
	hashed_password := HashPassword(u.Password)
	fmt.Printf("Pass: %s AND hashed: %s\n", u.Password, hashed_password)

	query := `INSERT INTO users (id, email, password) VALUES (?,?,?);`

	_, err := db.Exec(query, id, u.Email, hashed_password)
	if err != nil {
		return err
	}
	return nil
}
