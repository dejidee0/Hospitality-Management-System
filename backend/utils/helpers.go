package utils

type SignupData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordData struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}
