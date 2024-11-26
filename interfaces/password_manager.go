package interfaces

type PasswordManager interface {
	Encrypt(plaintext, key []byte) (bool, error)
	LoadFromFile(filename string, password []byte) ([]byte, error)
}
