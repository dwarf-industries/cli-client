package services

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"client/contracts"
	"client/interfaces"
	"client/models"
)

type RegisterService struct {
	WalletService       interfaces.WalletService
	RpcService          interfaces.RpcService
	ContractAddr        string
	VerificationService interfaces.IdentityVerificationService
}

func (r *RegisterService) Oracles() ([]models.Oracle, error) {
	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := contracts.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return []models.Oracle{}, err
	}

	oracleResult, err := contract.GetOracles(&bind.CallOpts{})
	if err != nil {
		return []models.Oracle{}, nil
	}

	wallet, err := r.WalletService.ActiveWallet()
	if err != nil {
		fmt.Println("Failed to retrive credentails")
		return []models.Oracle{}, err
	}
	auth, err := r.WalletService.NewTransactor(wallet)
	if err != nil {
		fmt.Println("Failed to Assign credentails")
		return []models.Oracle{}, err
	}

	var oracles []models.Oracle
	for _, o := range oracleResult {

		if o.Name == auth.From {
			oracles = append(oracles, models.Oracle{
				Name:       o.Name.Hex(),
				Ip:         o.Ip,
				Port:       o.Port,
				Reputation: *o.Reputation,
			})
			continue
		}

		//Verify that the node is running with the expected public address
		verified := r.VerificationService.Verify(o.Ip, o.Name.Hex())
		if !verified {
			continue
		}

		oracles = append(oracles, models.Oracle{
			Name:       o.Name.Hex(),
			Ip:         o.Ip,
			Port:       o.Port,
			Reputation: *o.Reputation,
		})
	}
	return oracles, nil
}
