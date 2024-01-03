package models

type UserCredentials struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}
