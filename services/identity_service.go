package services

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"client/interfaces"
)

type IdentityService struct {
	WalletService interfaces.WalletService
}

func (i *IdentityService) Verify(ip string, expected string) bool {
	challange, err := i.generateRandomBytes(512)
	if err != nil {
		fmt.Print("failed to generate challange aborting")
		return false
	}
	jsonData, err := json.Marshal(challange)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return false
	}

	req, err := http.NewRequest("POST", ip, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return false
	}

	fmt.Printf("Response status: %s\n", resp.Status)

	i.WalletService.VerifySignature(challange, []byte(body), expected)
	return false
}

func (i *IdentityService) generateRandomBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return bytes, nil
}
