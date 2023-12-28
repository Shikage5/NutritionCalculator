// validation.go
package validation

// ValidationService is an interface for validation of username and password.
type ValidationService interface {
	ValidateCredentials(username, password string) bool
}

// CredentialsValidationService is a default implementation of ValidationService.
type CredentialsValidationService struct{}

// ValidateCredentials checks if the username and password are not empty.
func (v *CredentialsValidationService) ValidateCredentials(username, password string) bool {
	return username != "" && password != ""
}
