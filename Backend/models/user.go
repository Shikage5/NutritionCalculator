package models

import (
	"encoding/json"
	"os"
)

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

func ReadUsersFromJSONFile(filename string) ([]User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}
