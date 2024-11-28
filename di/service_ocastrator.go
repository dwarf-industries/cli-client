package di

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"client/interfaces"
	"client/services"
)

var walletService interfaces.WalletService
var rpcService interfaces.RpcService
var storage interfaces.Storage
var IdentityVerificationService interfaces.IdentityVerificationService
var passwordManager interfaces.PasswordManager
var certificateService interfaces.CertificateService
var keyService interfaces.KeyService

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
	storage = setupDatabase()
	IdentityVerificationService = &services.IdentityService{}
	passwordManager = &services.PasswordManager{}
	certificateService = &services.CertificateService{}
	keyService = &services.KeyService{}
}

func setupDatabase() interfaces.Storage {
	storage := &services.Storage{}

	dbPath := os.Getenv("DbPath")
	dbName := os.Getenv("DbName")
	dbFile := "./db.sql"
	dbData := FileFromExecutable(&dbFile)
	tablesData, err := os.ReadFile(*dbData)

	if err != nil {
		log.Fatal("Couldn't find database file, aborting!")
	}

	tables := string(tablesData)
	storage.New(&dbPath, &dbName, &tables)
	storage.Open()
	storage.Initialize()
	return storage
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

func DatabaseService() interfaces.Storage {
	return storage
}

func GetIdentityVerificationService() interfaces.IdentityVerificationService {
	return IdentityVerificationService
}

func GetPasswordManager() interfaces.PasswordManager {
	return passwordManager
}

func GetCertificateService() interfaces.CertificateService {
	return certificateService
}

func GetKeyService() interfaces.KeyService {
	return keyService
}

func getExecutablePath() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		os.Exit(1)
	}
	return filepath.Dir(exePath)
}

func FileFromExecutable(fileName *string) *string {
	path := filepath.Join(getExecutablePath(), *fileName)
	return &path
}
