package utils

import (
	"NutritionCalculator/data/models"
	"encoding/json"
	"errors"
	"os"
)

type JSONFileOperations interface {
	ReadUserCredFromJSONFile(filename string) ([]models.UserCredentials, error)
	WriteUserCredInJSONFile(newUser models.UserCredentials, filename string) error
	ReadUserDataFromJSONFile(u *models.UserData, userDataPath string) error
	WriteUserDataToJSONFile(u *models.UserData, userDataPath string) error
}

type DefaultJSONFileOperations struct{}

// ==================== UserCredentials data operations ====================
func ReadUserCredFromJSONFile(filename string) ([]models.UserCredentials, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []models.UserCredentials

	if len(data) == 0 {
		// The JSON file is empty, so initialize 'users' as an empty slice
		users = []models.UserCredentials{}
	} else {
		if err := json.Unmarshal(data, &users); err != nil {
			return nil, err
		}
	}

	return users, nil
}

func WriteUserCredInJSONFile(newUser models.UserCredentials, filename string) error {

	if newUser.Username == "" || newUser.PasswordHash == "" {
		return errors.New("username and password cannot be empty")
	}

	// Read existing users from the file
	existingUsers, err := ReadUserCredFromJSONFile(filename)
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

// ==================== UserData data operations ====================

func ReadUserDataFromJSONFile(u *models.UserData, userDataPath string) error {
	filepath := userDataPath + u.Username + ".json"
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		// The JSON file is empty, so initialize 'u' as an empty struct
		*u = models.UserData{}
		return nil
	}

	return json.Unmarshal(data, u)
}
func WriteUserDataToJSONFile(u *models.UserData, userDataPath string) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	filepath := userDataPath + u.Username + ".json"
	return os.WriteFile(filepath, data, 0644)
}
