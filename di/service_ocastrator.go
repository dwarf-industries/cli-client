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
var registerService interfaces.RegisterService
var authenticationService interfaces.AuthenticationService
var socketService interfaces.SocketConnection

func SetupServices() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpc := getRpc()
	storage = setupDatabase()

	rpcService = &services.RpcService{}
	rpcService.SetClient(rpc)
	passwordManager = &services.PasswordManager{
		Storage: storage,
	}
	walletService = &services.WalletService{
		Storage:         storage,
		PasswordManager: passwordManager,
		RpcService:      rpcService,
	}
	IdentityVerificationService = &services.IdentityService{
		WalletService: walletService,
	}
	configured, err := passwordManager.LoadHash()
	if !configured || err != nil {
		fmt.Println("It appears that your account is not setup, please use 'client setup --help for more information'")
	}
	certificateService = &services.CertificateService{}
	keyService = &services.KeyService{}
	registerService = &services.RegisterService{
		WalletService:       walletService,
		RpcService:          rpcService,
		ContractAddr:        os.Getenv("CONTRACT_ADDRESS"),
		VerificationService: IdentityVerificationService,
	}
	authenticationService = &services.AuthenticationService{}
	authenticationService.Init()
	socketService = &services.SocketConnection{}
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

func GetRegisterService() interfaces.RegisterService {
	return registerService
}

func GetAuthenticationService() interfaces.AuthenticationService {
	return authenticationService
}

func GetSocketService() interfaces.SocketConnection {
	return socketService
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
