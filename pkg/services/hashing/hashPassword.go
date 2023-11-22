package hashing

import "golang.org/x/crypto/bcrypt"

type HashingService interface {
	HashPassword(password string) (string, error)
}

type DefaultHashingService struct{}

func (h *DefaultHashingService) HashPassword(password string) (string, error) {
	// Implement hashing logic using a secure algorithm (e.g., bcrypt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
