package services

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"client/contracts"
	"client/interfaces"
)

type PaymentProcessor struct {
	WalletService interfaces.WalletService
	RpcService    interfaces.RpcService
}

func (p *PaymentProcessor) PayNetworkTax(nodes *[]string, tax *big.Int) bool {
	contractAddress := common.HexToAddress(os.Getenv("PAYMENT_LEDGER"))
	contract, err := contracts.NewPaymentLedger(contractAddress, p.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return false
	}
	wallet, err := p.WalletService.ActiveWallet()
	if err != nil {
		fmt.Println("Wallet locked, aborting!")
		return false
	}
	transactionOps, err := p.WalletService.NewTransactor(wallet)
	if err != nil {
		fmt.Println("Failed to generate transaction options for active wallet, aborting")
		return false
	}

	transactionOps.Value = tax
	_, err = contract.RecordPayment(transactionOps, *nodes)
	if err != nil {
		fmt.Println("Failed to execute payment, aborting!")
		return false
	}

	return true
}

func (p *PaymentProcessor) CalculatePayment(size int) *big.Int {
	contractAddress := common.HexToAddress(os.Getenv("PAYMENT_LEDGER"))
	contract, err := contracts.NewPaymentLedger(contractAddress, p.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
	}

	data, err := contract.CalculatePayment(&bind.CallOpts{}, big.NewInt(int64(size)))
	if err != nil {
		fmt.Println("Failed to retrive transfer amount")
		return &big.Int{}
	}

	return data.TotalAmount
}
