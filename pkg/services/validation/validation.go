// validation.go
package validation

// CredentialsValidationService is an interface for validation of username and password.
type CredentialsValidationService interface {
	ValidateCredentials(username, password string) bool
}

// DefaultCredentialsValidationService is a default implementation of CredentialsValidationService.
type DefaultCredentialsValidationService struct{}

// ValidateCredentials checks if the username and password are not empty.
func (v *DefaultCredentialsValidationService) ValidateCredentials(username, password string) bool {
	return username != "" && password != ""
}
