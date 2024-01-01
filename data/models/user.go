package models

import (
	"encoding/json"
	"errors"
	"os"
)

type UserCredentials struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

func ReadUsersFromJSONFile(filename string) ([]UserCredentials, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []UserCredentials

	if len(data) == 0 {
		// The JSON file is empty, so initialize 'users' as an empty slice
		users = []UserCredentials{}
	} else {
		if err := json.Unmarshal(data, &users); err != nil {
			return nil, err
		}
	}

	return users, nil
}

func WriteUserInJSONFile(newUser UserCredentials, filename string) error {

	if newUser.Username == "" || newUser.PasswordHash == "" {
		return errors.New("username and password cannot be empty")
	}

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
