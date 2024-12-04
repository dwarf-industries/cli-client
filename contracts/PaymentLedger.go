// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PaymentLedgerMetaData contains all meta data concerning the PaymentLedger contract.
var PaymentLedgerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeSetter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"paymentIDs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataSize\",\"type\":\"uint256\"}],\"name\":\"calculatePayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"paymentId\",\"type\":\"string\"}],\"name\":\"paymentProcessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"paymentIDsList\",\"type\":\"string[]\"}],\"name\":\"recordPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611cff380380611cff833981810160405281019061003191906100d4565b805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f5ffd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b611bf38061010c5f395ff3fe608060405260043610610049575f3560e01c80630dfd1e091461004d5780634e71d92d146100895780638c67f63a1461009f578063a07a1aca146100dc578063b69ef8a8146100f8575b5f5ffd5b348015610058575f5ffd5b50610073600480360381019061006e9190610e9b565b610122565b6040516100809190610efc565b60405180910390f35b348015610094575f5ffd5b5061009d610180565b005b3480156100aa575f5ffd5b506100c560048036038101906100c09190610f48565b6103c9565b6040516100d3929190610f82565b60405180910390f35b6100f660048036038101906100f1919061108b565b61053f565b005b348015610103575f5ffd5b5061010c6107c1565b60405161011991906110d2565b60405180910390f35b5f5f600183604051610134919061113d565b90815260200160405180910390206040518060400160405290815f8201548152602001600182015f9054906101000a900460ff16151515158152505090505f815f015111915050919050565b5f5f90505f5f90505b60028054905081101561033f575f600282815481106101ab576101aa611153565b5b905f5260205f200180546101be906111ad565b80601f01602080910402602001604051908101604052809291908181526020018280546101ea906111ad565b80156102355780601f1061020c57610100808354040283529160200191610235565b820191905f5260205f20905b81548152906001019060200180831161021857829003601f168201915b505050505090505f60018260405161024d919061113d565b90815260200160405180910390209050610267823361096d565b15610330575f815f0154116102b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a89061125d565b60405180910390fd5b806001015f9054906101000a900460ff1615610302576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f9906112c5565b60405180910390fd5b6001816001015f6101000a81548160ff021916908315150217905550805f01548461032d9190611310565b93505b50508080600101915050610189565b505f8111610382576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610379906113b3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156103c5573d5f5f3e3d5ffd5b5050565b5f5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663822fcca56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610435573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061045991906113e5565b90505f816104006103ff8761046e9190611310565b610478919061143d565b610482919061146d565b905060645f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc0438306040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104ef573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061051391906113e5565b8261051e919061146d565b610528919061143d565b925082816105369190611310565b93505050915091565b5f815111610582576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610579906114f8565b60405180910390fd5b5f34116105c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105bb90611586565b60405180910390fd5b5f3490505f5f6105d3836103c9565b80935081945050505f82846105e891906115a4565b90508451816105f7919061143d565b915080855183610607919061146d565b14610647576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063e90611621565b60405180910390fd5b5f5f90505b8551811015610780575f86828151811061066957610668611153565b5b60200260200101519050600181604051610683919061113d565b90815260200160405180910390206001015f9054906101000a900460ff16156106e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d890611689565b60405180910390fd5b60405180604001604052808581526020015f1515815250600182604051610708919061113d565b90815260200160405180910390205f820151815f01556020820151816001015f6101000a81548160ff021916908315150217905550905050600281908060018154018082558091505060019003905f5260205f20015f9091909190915090816107719190611847565b5050808060010191505061064c565b507f1e49930a386aa7603def05c5509bece3d51a5de6b6401b4aeac72261ed2b966785856040516107b2929190611a19565b60405180910390a15050505050565b5f5f5f90505f5f90505b600280549050811015610965575f600282815481106107ed576107ec611153565b5b905f5260205f20018054610800906111ad565b80601f016020809104026020016040519081016040528092919081815260200182805461082c906111ad565b80156108775780601f1061084e57610100808354040283529160200191610877565b820191905f5260205f20905b81548152906001019060200180831161085a57829003601f168201915b505050505090505f60018260405161088f919061113d565b908152602001604051809103902090506108a9823361096d565b15610956575f815f0154116108f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ea9061125d565b60405180910390fd5b806001015f9054906101000a900460ff1615610944576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093b906112c5565b60405180910390fd5b805f0154846109539190611310565b93505b505080806001019150506107cb565b508091505090565b5f5f61097883610992565b90505f6109858583610bf4565b9050809250505092915050565b60605f826040516020016109a69190611abc565b60405160208183030381529060405290505f602a67ffffffffffffffff8111156109d3576109d2610d77565b5b6040519080825280601f01601f191660200182016040528015610a055781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610a3c57610a3b611153565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610a9f57610a9e611153565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5f90505b6014811015610be9575f838281518110610aef57610aee611153565b5b602001015160f81c60f81b60f81c9050610b14601082610b0f9190611ae2565b610d12565b83600284610b22919061146d565b6002610b2e9190611310565b81518110610b3f57610b3e611153565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350610b82601082610b7d9190611b12565b610d12565b83600284610b90919061146d565b6003610b9c9190611310565b81518110610bad57610bac611153565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350508080600101915050610ad2565b508092505050919050565b5f5f8390505f839050815181511015610c0b575f5ffd5b5f5f90505f5f90505b83518351610c2291906115a4565b8111610d05575f600190505f5f90505b8551811015610ce057858181518110610c4e57610c4d611153565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916858285610c889190611310565b81518110610c9957610c98611153565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610cd3575f9150610ce0565b8080600101915050610c32565b508015610cf1576001925050610d05565b508080610cfd90611b42565b915050610c14565b5080935050505092915050565b5f600a8260ff161015610d3657603082610d2c9190611b89565b60f81b9050610d49565b605782610d439190611b89565b60f81b90505b919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610dad82610d67565b810181811067ffffffffffffffff82111715610dcc57610dcb610d77565b5b80604052505050565b5f610dde610d4e565b9050610dea8282610da4565b919050565b5f67ffffffffffffffff821115610e0957610e08610d77565b5b610e1282610d67565b9050602081019050919050565b828183375f83830152505050565b5f610e3f610e3a84610def565b610dd5565b905082815260208101848484011115610e5b57610e5a610d63565b5b610e66848285610e1f565b509392505050565b5f82601f830112610e8257610e81610d5f565b5b8135610e92848260208601610e2d565b91505092915050565b5f60208284031215610eb057610eaf610d57565b5b5f82013567ffffffffffffffff811115610ecd57610ecc610d5b565b5b610ed984828501610e6e565b91505092915050565b5f8115159050919050565b610ef681610ee2565b82525050565b5f602082019050610f0f5f830184610eed565b92915050565b5f819050919050565b610f2781610f15565b8114610f31575f5ffd5b50565b5f81359050610f4281610f1e565b92915050565b5f60208284031215610f5d57610f5c610d57565b5b5f610f6a84828501610f34565b91505092915050565b610f7c81610f15565b82525050565b5f604082019050610f955f830185610f73565b610fa26020830184610f73565b9392505050565b5f67ffffffffffffffff821115610fc357610fc2610d77565b5b602082029050602081019050919050565b5f5ffd5b5f610fea610fe584610fa9565b610dd5565b9050808382526020820190506020840283018581111561100d5761100c610fd4565b5b835b8181101561105457803567ffffffffffffffff81111561103257611031610d5f565b5b80860161103f8982610e6e565b8552602085019450505060208101905061100f565b5050509392505050565b5f82601f83011261107257611071610d5f565b5b8135611082848260208601610fd8565b91505092915050565b5f602082840312156110a05761109f610d57565b5b5f82013567ffffffffffffffff8111156110bd576110bc610d5b565b5b6110c98482850161105e565b91505092915050565b5f6020820190506110e55f830184610f73565b92915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f611117826110eb565b61112181856110f5565b93506111318185602086016110ff565b80840191505092915050565b5f611148828461110d565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111c457607f821691505b6020821081036111d7576111d6611180565b5b50919050565b5f82825260208201905092915050565b7f4e6f207061796d656e7420617661696c61626c6520666f7220746869732070615f8201527f796d656e74204944000000000000000000000000000000000000000000000000602082015250565b5f6112476028836111dd565b9150611252826111ed565b604082019050919050565b5f6020820190508181035f8301526112748161123b565b9050919050565b7f5061796d656e7420616c726561647920636c61696d65640000000000000000005f82015250565b5f6112af6017836111dd565b91506112ba8261127b565b602082019050919050565b5f6020820190508181035f8301526112dc816112a3565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61131a82610f15565b915061132583610f15565b925082820190508082111561133d5761133c6112e3565b5b92915050565b7f4e6f207061796d656e747320617661696c61626c6520666f7220636c61696d695f8201527f6e67000000000000000000000000000000000000000000000000000000000000602082015250565b5f61139d6022836111dd565b91506113a882611343565b604082019050919050565b5f6020820190508181035f8301526113ca81611391565b9050919050565b5f815190506113df81610f1e565b92915050565b5f602082840312156113fa576113f9610d57565b5b5f611407848285016113d1565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61144782610f15565b915061145283610f15565b92508261146257611461611410565b5b828204905092915050565b5f61147782610f15565b915061148283610f15565b925082820261149081610f15565b915082820484148315176114a7576114a66112e3565b5b5092915050565b7f4174206c65617374206f6e65207061796d656e742049442072657175697265645f82015250565b5f6114e26020836111dd565b91506114ed826114ae565b602082019050919050565b5f6020820190508181035f83015261150f816114d6565b9050919050565b7f5061796d656e74206d7573742062652067726561746572207468616e207a65725f8201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b5f6115706021836111dd565b915061157b82611516565b604082019050919050565b5f6020820190508181035f83015261159d81611564565b9050919050565b5f6115ae82610f15565b91506115b983610f15565b92508282039050818111156115d1576115d06112e3565b5b92915050565b7f496e636f727265637420616d6f756e7420646973747269627574696f6e0000005f82015250565b5f61160b601d836111dd565b9150611616826115d7565b602082019050919050565b5f6020820190508181035f830152611638816115ff565b9050919050565b7f5061796d656e7420494420616c72656164792070726f636573736564000000005f82015250565b5f611673601c836111dd565b915061167e8261163f565b602082019050919050565b5f6020820190508181035f8301526116a081611667565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026117037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826116c8565b61170d86836116c8565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61174861174361173e84610f15565b611725565b610f15565b9050919050565b5f819050919050565b6117618361172e565b61177561176d8261174f565b8484546116d4565b825550505050565b5f5f905090565b61178c61177d565b611797818484611758565b505050565b5b818110156117ba576117af5f82611784565b60018101905061179d565b5050565b601f8211156117ff576117d0816116a7565b6117d9846116b9565b810160208510156117e8578190505b6117fc6117f4856116b9565b83018261179c565b50505b505050565b5f82821c905092915050565b5f61181f5f1984600802611804565b1980831691505092915050565b5f6118378383611810565b9150826002028217905092915050565b611850826110eb565b67ffffffffffffffff81111561186957611868610d77565b5b61187382546111ad565b61187e8282856117be565b5f60209050601f8311600181146118af575f841561189d578287015190505b6118a7858261182c565b86555061190e565b601f1984166118bd866116a7565b5f5b828110156118e4578489015182556001820191506020850194506020810190506118bf565b8683101561190157848901516118fd601f891682611810565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f611959826110eb565b611963818561193f565b93506119738185602086016110ff565b61197c81610d67565b840191505092915050565b5f611992838361194f565b905092915050565b5f602082019050919050565b5f6119b082611916565b6119ba8185611920565b9350836020820285016119cc85611930565b805f5b85811015611a0757848403895281516119e88582611987565b94506119f38361199a565b925060208a019950506001810190506119cf565b50829750879550505050505092915050565b5f6040820190508181035f830152611a3181856119a6565b9050611a406020830184610f73565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611a7082611a47565b9050919050565b5f8160601b9050919050565b5f611a8d82611a77565b9050919050565b5f611a9e82611a83565b9050919050565b611ab6611ab182611a66565b611a94565b82525050565b5f611ac78284611aa5565b60148201915081905092915050565b5f60ff82169050919050565b5f611aec82611ad6565b9150611af783611ad6565b925082611b0757611b06611410565b5b828204905092915050565b5f611b1c82611ad6565b9150611b2783611ad6565b925082611b3757611b36611410565b5b828206905092915050565b5f611b4c82610f15565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b7e57611b7d6112e3565b5b600182019050919050565b5f611b9382611ad6565b9150611b9e83611ad6565b9250828201905060ff811115611bb757611bb66112e3565b5b9291505056fea2646970667358221220efdc270856439bc00e53e445e1063209eb2ef1d5193914f256c2535315607d4564736f6c634300081c0033",
}

// PaymentLedgerABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentLedgerMetaData.ABI instead.
var PaymentLedgerABI = PaymentLedgerMetaData.ABI

// PaymentLedgerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PaymentLedgerMetaData.Bin instead.
var PaymentLedgerBin = PaymentLedgerMetaData.Bin

// DeployPaymentLedger deploys a new Ethereum contract, binding an instance of PaymentLedger to it.
func DeployPaymentLedger(auth *bind.TransactOpts, backend bind.ContractBackend, _feeSetter common.Address) (common.Address, *types.Transaction, *PaymentLedger, error) {
	parsed, err := PaymentLedgerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymentLedgerBin), backend, _feeSetter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaymentLedger{PaymentLedgerCaller: PaymentLedgerCaller{contract: contract}, PaymentLedgerTransactor: PaymentLedgerTransactor{contract: contract}, PaymentLedgerFilterer: PaymentLedgerFilterer{contract: contract}}, nil
}

// PaymentLedger is an auto generated Go binding around an Ethereum contract.
type PaymentLedger struct {
	PaymentLedgerCaller     // Read-only binding to the contract
	PaymentLedgerTransactor // Write-only binding to the contract
	PaymentLedgerFilterer   // Log filterer for contract events
}

// PaymentLedgerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentLedgerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentLedgerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentLedgerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentLedgerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentLedgerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentLedgerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentLedgerSession struct {
	Contract     *PaymentLedger    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentLedgerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentLedgerCallerSession struct {
	Contract *PaymentLedgerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PaymentLedgerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentLedgerTransactorSession struct {
	Contract     *PaymentLedgerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PaymentLedgerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentLedgerRaw struct {
	Contract *PaymentLedger // Generic contract binding to access the raw methods on
}

// PaymentLedgerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentLedgerCallerRaw struct {
	Contract *PaymentLedgerCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentLedgerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentLedgerTransactorRaw struct {
	Contract *PaymentLedgerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentLedger creates a new instance of PaymentLedger, bound to a specific deployed contract.
func NewPaymentLedger(address common.Address, backend bind.ContractBackend) (*PaymentLedger, error) {
	contract, err := bindPaymentLedger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentLedger{PaymentLedgerCaller: PaymentLedgerCaller{contract: contract}, PaymentLedgerTransactor: PaymentLedgerTransactor{contract: contract}, PaymentLedgerFilterer: PaymentLedgerFilterer{contract: contract}}, nil
}

// NewPaymentLedgerCaller creates a new read-only instance of PaymentLedger, bound to a specific deployed contract.
func NewPaymentLedgerCaller(address common.Address, caller bind.ContractCaller) (*PaymentLedgerCaller, error) {
	contract, err := bindPaymentLedger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentLedgerCaller{contract: contract}, nil
}

// NewPaymentLedgerTransactor creates a new write-only instance of PaymentLedger, bound to a specific deployed contract.
func NewPaymentLedgerTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentLedgerTransactor, error) {
	contract, err := bindPaymentLedger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentLedgerTransactor{contract: contract}, nil
}

// NewPaymentLedgerFilterer creates a new log filterer instance of PaymentLedger, bound to a specific deployed contract.
func NewPaymentLedgerFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentLedgerFilterer, error) {
	contract, err := bindPaymentLedger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentLedgerFilterer{contract: contract}, nil
}

// bindPaymentLedger binds a generic wrapper to an already deployed contract.
func bindPaymentLedger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentLedgerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentLedger *PaymentLedgerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentLedger.Contract.PaymentLedgerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentLedger *PaymentLedgerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentLedger.Contract.PaymentLedgerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentLedger *PaymentLedgerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentLedger.Contract.PaymentLedgerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentLedger *PaymentLedgerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentLedger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentLedger *PaymentLedgerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentLedger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentLedger *PaymentLedgerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentLedger.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_PaymentLedger *PaymentLedgerCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PaymentLedger.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_PaymentLedger *PaymentLedgerSession) Balance() (*big.Int, error) {
	return _PaymentLedger.Contract.Balance(&_PaymentLedger.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_PaymentLedger *PaymentLedgerCallerSession) Balance() (*big.Int, error) {
	return _PaymentLedger.Contract.Balance(&_PaymentLedger.CallOpts)
}

// CalculatePayment is a free data retrieval call binding the contract method 0x8c67f63a.
//
// Solidity: function calculatePayment(uint256 dataSize) view returns(uint256 totalAmount, uint256 networkFee)
func (_PaymentLedger *PaymentLedgerCaller) CalculatePayment(opts *bind.CallOpts, dataSize *big.Int) (struct {
	TotalAmount *big.Int
	NetworkFee  *big.Int
}, error) {
	var out []interface{}
	err := _PaymentLedger.contract.Call(opts, &out, "calculatePayment", dataSize)

	outstruct := new(struct {
		TotalAmount *big.Int
		NetworkFee  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NetworkFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculatePayment is a free data retrieval call binding the contract method 0x8c67f63a.
//
// Solidity: function calculatePayment(uint256 dataSize) view returns(uint256 totalAmount, uint256 networkFee)
func (_PaymentLedger *PaymentLedgerSession) CalculatePayment(dataSize *big.Int) (struct {
	TotalAmount *big.Int
	NetworkFee  *big.Int
}, error) {
	return _PaymentLedger.Contract.CalculatePayment(&_PaymentLedger.CallOpts, dataSize)
}

// CalculatePayment is a free data retrieval call binding the contract method 0x8c67f63a.
//
// Solidity: function calculatePayment(uint256 dataSize) view returns(uint256 totalAmount, uint256 networkFee)
func (_PaymentLedger *PaymentLedgerCallerSession) CalculatePayment(dataSize *big.Int) (struct {
	TotalAmount *big.Int
	NetworkFee  *big.Int
}, error) {
	return _PaymentLedger.Contract.CalculatePayment(&_PaymentLedger.CallOpts, dataSize)
}

// PaymentProcessed is a free data retrieval call binding the contract method 0x0dfd1e09.
//
// Solidity: function paymentProcessed(string paymentId) view returns(bool)
func (_PaymentLedger *PaymentLedgerCaller) PaymentProcessed(opts *bind.CallOpts, paymentId string) (bool, error) {
	var out []interface{}
	err := _PaymentLedger.contract.Call(opts, &out, "paymentProcessed", paymentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PaymentProcessed is a free data retrieval call binding the contract method 0x0dfd1e09.
//
// Solidity: function paymentProcessed(string paymentId) view returns(bool)
func (_PaymentLedger *PaymentLedgerSession) PaymentProcessed(paymentId string) (bool, error) {
	return _PaymentLedger.Contract.PaymentProcessed(&_PaymentLedger.CallOpts, paymentId)
}

// PaymentProcessed is a free data retrieval call binding the contract method 0x0dfd1e09.
//
// Solidity: function paymentProcessed(string paymentId) view returns(bool)
func (_PaymentLedger *PaymentLedgerCallerSession) PaymentProcessed(paymentId string) (bool, error) {
	return _PaymentLedger.Contract.PaymentProcessed(&_PaymentLedger.CallOpts, paymentId)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PaymentLedger *PaymentLedgerTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentLedger.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PaymentLedger *PaymentLedgerSession) Claim() (*types.Transaction, error) {
	return _PaymentLedger.Contract.Claim(&_PaymentLedger.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PaymentLedger *PaymentLedgerTransactorSession) Claim() (*types.Transaction, error) {
	return _PaymentLedger.Contract.Claim(&_PaymentLedger.TransactOpts)
}

// RecordPayment is a paid mutator transaction binding the contract method 0xa07a1aca.
//
// Solidity: function recordPayment(string[] paymentIDsList) payable returns()
func (_PaymentLedger *PaymentLedgerTransactor) RecordPayment(opts *bind.TransactOpts, paymentIDsList []string) (*types.Transaction, error) {
	return _PaymentLedger.contract.Transact(opts, "recordPayment", paymentIDsList)
}

// RecordPayment is a paid mutator transaction binding the contract method 0xa07a1aca.
//
// Solidity: function recordPayment(string[] paymentIDsList) payable returns()
func (_PaymentLedger *PaymentLedgerSession) RecordPayment(paymentIDsList []string) (*types.Transaction, error) {
	return _PaymentLedger.Contract.RecordPayment(&_PaymentLedger.TransactOpts, paymentIDsList)
}

// RecordPayment is a paid mutator transaction binding the contract method 0xa07a1aca.
//
// Solidity: function recordPayment(string[] paymentIDsList) payable returns()
func (_PaymentLedger *PaymentLedgerTransactorSession) RecordPayment(paymentIDsList []string) (*types.Transaction, error) {
	return _PaymentLedger.Contract.RecordPayment(&_PaymentLedger.TransactOpts, paymentIDsList)
}

// PaymentLedgerPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the PaymentLedger contract.
type PaymentLedgerPaymentReceivedIterator struct {
	Event *PaymentLedgerPaymentReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentLedgerPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentLedgerPaymentReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentLedgerPaymentReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentLedgerPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentLedgerPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentLedgerPaymentReceived represents a PaymentReceived event raised by the PaymentLedger contract.
type PaymentLedgerPaymentReceived struct {
	PaymentIDs  []string
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x1e49930a386aa7603def05c5509bece3d51a5de6b6401b4aeac72261ed2b9667.
//
// Solidity: event PaymentReceived(string[] paymentIDs, uint256 totalAmount)
func (_PaymentLedger *PaymentLedgerFilterer) FilterPaymentReceived(opts *bind.FilterOpts) (*PaymentLedgerPaymentReceivedIterator, error) {

	logs, sub, err := _PaymentLedger.contract.FilterLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return &PaymentLedgerPaymentReceivedIterator{contract: _PaymentLedger.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x1e49930a386aa7603def05c5509bece3d51a5de6b6401b4aeac72261ed2b9667.
//
// Solidity: event PaymentReceived(string[] paymentIDs, uint256 totalAmount)
func (_PaymentLedger *PaymentLedgerFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *PaymentLedgerPaymentReceived) (event.Subscription, error) {

	logs, sub, err := _PaymentLedger.contract.WatchLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentLedgerPaymentReceived)
				if err := _PaymentLedger.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentReceived is a log parse operation binding the contract event 0x1e49930a386aa7603def05c5509bece3d51a5de6b6401b4aeac72261ed2b9667.
//
// Solidity: event PaymentReceived(string[] paymentIDs, uint256 totalAmount)
func (_PaymentLedger *PaymentLedgerFilterer) ParsePaymentReceived(log types.Log) (*PaymentLedgerPaymentReceived, error) {
	event := new(PaymentLedgerPaymentReceived)
	if err := _PaymentLedger.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
