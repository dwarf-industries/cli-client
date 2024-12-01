package interfaces

type PasswordManager interface {
	Encrypt(plaintext, key []byte) (*[]byte, error)
	LoadFromFile(filename string, password []byte) ([]byte, error)
	Decrypt(data []byte, key []byte) ([]byte, error)
	SetupPassword(password *string) bool
	Match(password *string) bool
	LoadHash() (bool, error)
	Input() *string
}
