package models

import (
	"database/sql"
	"errors"
	"fmt"
	"hms/database"
	"log"

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
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	id := uuid.New().String()

	// hash password here
	hashed_password := HashPassword(u.Password)
	fmt.Printf("Pass: %s AND hashed: %s\n", u.Password, hashed_password)

	query := `INSERT INTO users (id, email, password) VALUES (@id, @Email, @Password);`

	_, err = db.Exec(query,
		sql.Named("id", id),
		sql.Named("Email", u.Email),
		sql.Named("Password", hashed_password),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserByEmail(email string) error {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	query := `SELECT id FROM users WHERE email = @email;`

	row := db.QueryRow(query, sql.Named("email", email))
	// var user models.User
	err = row.Scan(&u.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateResetPasswordToken(token string) error {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	query := `UPDATE users SET change_password_token = @reset_token WHERE id = @id;`

	_, err = db.Exec(query, sql.Named("reset_token", token), sql.Named("id", u.Id))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdatePassword(email, password string) error {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	query := `UPDATE users SET password = @password, change_password_token = NULL WHERE email = @email;`

	_, err = db.Exec(query, sql.Named("password", password), sql.Named("email", email))
	if err != nil {
		return err
	}
	return nil
}
