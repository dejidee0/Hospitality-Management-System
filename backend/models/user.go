package models

import (
	"errors"
	"fmt"
	"hms/database"

	"github.com/google/uuid"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
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

func (u *User) GetUserByEmail(email string) error {
	db := database.GetDB()
	defer db.Close()

	query := `SELECT id FROM users WHERE email = ?;`

	row := db.QueryRow(query, email)
	// var user models.User
	err := row.Scan(&u.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateResetPasswordToken(token string) error {
	db := database.GetDB()
	defer db.Close()

	query := `UPDATE users SET change_password_token = ? WHERE id = ?;`

	_, err := db.Exec(query, token, u.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdatePassword(email, password string) error {
	db := database.GetDB()
	defer db.Close()

	query := `UPDATE users SET password = ?, change_password_token = "" WHERE email = ?;`

	_, err := db.Exec(query, password, email)
	if err != nil {
		return err
	}
	return nil
}
