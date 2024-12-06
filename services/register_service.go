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

func (r *RegisterService) Oracles() ([]models.Node, error) {
	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := contracts.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return []models.Node{}, err
	}

	oracleResult, err := contract.GetOracles(&bind.CallOpts{})
	if err != nil {
		return []models.Node{}, nil
	}

	var oracles []models.Node
	for _, o := range oracleResult {
		verified := r.VerificationService.Verify(o.Ip, o.Name.Hex())
		if !verified {
			continue
		}

		oracles = append(oracles, models.Node{
			Name:       o.Name.Hex(),
			Ip:         o.Ip,
			Port:       o.Port,
			Reputation: *o.Reputation,
		})
	}
	return oracles, nil
}
