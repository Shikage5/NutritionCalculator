package basicAuth

type Authenticator interface {
	Authenticate(user, password string) bool
}
type AuthenticatorFunc func(user, password string) bool

func (af AuthenticatorFunc) Authenticate(user, password string) bool {
	return af(user, password)
}
