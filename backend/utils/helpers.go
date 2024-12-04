package utils

import "errors"

func ValidateSignupData(data *SignupData) (err error) {
	if data.Email == "" {
		err = errors.New("email is required")
		return
	} else if data.Password == "" {
		err = errors.New("password is required")
		return
	}
	return nil
}

type SignupData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
