package interfaces

type AuthenticationService interface {
	Authenticate(url string, certificate *[]byte) (*string, error)
	GenerateSessionToken(url *string) (*string, error)
	GetSessionToken(url *string) *string
}
