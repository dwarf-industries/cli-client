package services

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AuthenticationService struct {
	tokens map[string]string
}

func (a *AuthenticationService) Authenticate(url string, certificate *[]byte) (*[]byte, error) {
	requestBody := bytes.NewBuffer(*certificate)

	resp, err := http.Post(url+"/v1/identity/certificate", "application/octet-stream", requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned error: %s", string(body))
	}

	body, _ := io.ReadAll(resp.Body)
	var str string
	if err := json.Unmarshal(body, &str); err != nil {
		fmt.Println("Failed to unmarshal challenge")
		return nil, err
	}

	bytes, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println("Failed to extract bytes from hex")
		return nil, err
	}

	return &bytes, nil
}

func (a *AuthenticationService) GenerateSessionToken(url *string) (*string, error) {
	tokenLength := 32
	bytes := make([]byte, tokenLength)

	if _, err := rand.Read(bytes); err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	token := hex.EncodeToString(bytes)
	a.tokens[*url] = token
	return &token, nil
}

func (a *AuthenticationService) GetSessionToken(url *string) *string {
	token := a.tokens[*url]
	return &token
}

func (a *AuthenticationService) Init() {
	a.tokens = make(map[string]string)
}
