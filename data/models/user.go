package models

import (
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
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

func WriteUserInJSONFile(newUser User, filename string) error {

	if newUser.Username == "" || newUser.PasswordHash == "" {
		return errors.New("username and password cannot be empty")
	}

	// Hash the password before saving it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.PasswordHash = string(hashedPassword)
	// Read existing users from the file
	existingUsers, err := ReadUsersFromJSONFile(filename)
	if err != nil {
		return err
	}

	// Check if the username already exists
	for _, user := range existingUsers {
		if user.Username == newUser.Username {
			return errors.New("username already exists")
		}
	}

	// Add the new user to the slice
	existingUsers = append(existingUsers, newUser)

	// Serialize the updated user list to JSON
	updatedUserJSON, err := json.Marshal(existingUsers)
	if err != nil {
		return err
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile(filename, updatedUserJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}
