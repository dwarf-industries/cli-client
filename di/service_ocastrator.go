package di

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"client/interfaces"
	"client/services"
)

var walletService interfaces.WalletService
var rpcService interfaces.RpcService

func SetupServices() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpc := getRpc()

	rpcService = &services.RpcService{}
	rpcService.SetClient(rpc)

	walletService = &services.WalletService{
		PasswordManager: &services.PasswordManager{},
		RpcService:      rpcService,
	}

}
func getRpc() *string {
	rpc := os.Getenv("RPC")

	overriden, err := os.ReadFile("oracle-rpc")
	if err != nil {
		return &rpc
	}

	converted := string(overriden)
	return &converted
}

func WalletService() interfaces.WalletService {
	return walletService
}

func RpcService() interfaces.RpcService {
	return rpcService
}
