package interfaces

type AuthenticationService interface {
	Authenticate(url string, certificate *[]byte) (*[]byte, error)
	GenerateSessionToken(url *string) (*string, error)
	GetSessionToken(url *string) *string
	Init()
}
