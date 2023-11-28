// validation.go
package validation

// CredentialsValidator is an interface for validation of username and password.
type CredentialsValidator interface {
	ValidateCredentials(username, password string) bool
}

// DefaultCredentialsValidator is a default implementation of CredentialsValidator.
type DefaultCredentialsValidator struct{}

// ValidateCredentials checks if the username and password are not empty.
func (v *DefaultCredentialsValidator) ValidateCredentials(username, password string) bool {
	return username != "" && password != ""
}
