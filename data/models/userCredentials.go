package models

type UserCredentials struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
