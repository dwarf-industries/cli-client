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

func (a *AuthenticationService) Authenticate(url string, certificate *[]byte) (*string, error) {
	requestBody := bytes.NewBuffer(*certificate)

	resp, err := http.Post(url+"/identity/certificate", "application/octet-stream", requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned error: %s", string(body))
	}

	var challenge map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&challenge); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	challengeValue, ok := challenge["Challenge"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid response structure")
	}

	return &challengeValue, nil
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
