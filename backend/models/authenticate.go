package models

import (
	"database/sql"
	"errors"
	"hms/database"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Authenticate(email, password string) (*User, error) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	// hash password
	// hash_pass := HashPassword(password)
	// fmt.Printf("Pass: %s AND hashed: %s\n", password, hash_pass)

	// query := `SELECT id, email, password FROM users WHERE email = ?;`
	query := `SELECT id, email, password FROM users WHERE email = @email;`
	row := db.QueryRow(query, sql.Named("email", email))

	var user User
	err = row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no such email")
		}
		return nil, err
	}
	ok := VerifyPassword(user.Password, password)
	if !ok {
		return nil, errors.New("incorrect password")
	}
	// sending a new user object without the password
	return &User{
		Id:    user.Id,
		Email: user.Email,
	}, nil
}

func HashPassword(password string) string {
	hash_bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash_bytes)
}

func VerifyPassword(hash, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
