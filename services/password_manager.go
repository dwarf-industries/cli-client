package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

type PasswordManager struct{}

func (p *PasswordManager) Encrypt(plaintext, key []byte) (bool, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return false, fmt.Errorf("failed to create cipher block: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return false, fmt.Errorf("failed to create AES-GCM: %w", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return false, fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	err = os.WriteFile("oracle_data", ciphertext, 0600)
	if err != nil {
		return false, fmt.Errorf("failed to write encrypted data to file: %w", err)
	}

	return true, nil
}

func (p *PasswordManager) LoadFromFile(filename string, password []byte) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return p.decrypt(file, password)
}

func (p *PasswordManager) decrypt(data []byte, key []byte) ([]byte, error) {
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
