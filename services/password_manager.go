package services

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"syscall"

	"golang.org/x/term"

	"client/interfaces"
)

type PasswordManager struct {
	Storage  interfaces.Storage
	password string
}

func (p *PasswordManager) Encrypt(plaintext, key []byte) (*[]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher block: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES-GCM: %w", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return &ciphertext, nil
}

func (p *PasswordManager) LoadFromFile(filename string, password []byte) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return p.Decrypt(file, password)
}

func (p *PasswordManager) Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher block: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES-GCM: %w", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("data is too short to contain a nonce")
	}

	nonce, ciphertextData := data[:nonceSize], data[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertextData, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return plaintext, nil
}

func (p *PasswordManager) SetupPassword(password *string) bool {
	hash := sha512.Sum512([]byte(*password))
	passwordHash := hex.EncodeToString(hash[:])

	saveSql := `
		INSERT INTO Accounts (password)
        VALUES ($1)
        ON CONFLICT DO NOTHING
	`
	err := p.Storage.Exec(&saveSql, &[]interface{}{
		&passwordHash,
	})

	return err == nil
}

func (p *PasswordManager) Match(password *string) bool {
	hash := sha512.Sum512([]byte(*password))
	if len(p.password) == 0 {
		return true
	}

	pHash := hex.EncodeToString(hash[:])

	return pHash == p.password
}

func (p *PasswordManager) LoadHash() (bool, error) {
	p.Storage.Open()
	defer p.Storage.Close()

	sql := `
		select password from Accounts
	`

	result := p.Storage.QuerySingle(&sql, &[]interface{}{})

	var dbHash string
	err := result.Scan(&dbHash)
	if err != nil {
		return false, err
	}

	p.password = dbHash
	return true, nil
}

func (p *PasswordManager) Input() *string {
	fmt.Print("Enter password: ")

	password := ""
	for {
		char, err := readSingleChar()
		if err != nil {
			fmt.Println("\nError reading password:", err)
			return nil
		}

		if char == '\n' || char == '\r' {
			fmt.Println()
			break
		}

		if char == 8 || char == 127 {
			if len(password) > 0 {
				password = password[:len(password)-1]
				fmt.Print("\b \b")
			}
			continue
		}

		password += string(char)
		fmt.Print("*")
	}

	fmt.Print("\r")
	return &password
}

func readSingleChar() (byte, error) {
	fd := int(syscall.Stdin)
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return 0, err
	}
	defer term.Restore(fd, oldState)

	reader := bufio.NewReader(os.Stdin)
	char, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	return char, nil
}
