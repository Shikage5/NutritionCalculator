package validation

// CredentialsValidator is a service that provides validation for username and password.
type CredentialsValidator struct{}

// NewCredentialsValidator creates a new instance of CredentialsValidator.
func NewCredentialsValidator() *CredentialsValidator {
	return &CredentialsValidator{}
}

// ValidateCredentials checks if the username and password are not empty.
func (v *CredentialsValidator) ValidateCredentials(username, password string) bool {
	return username != "" && password != ""
}
