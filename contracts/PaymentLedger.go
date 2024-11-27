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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeSetter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"paymentIDs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataSize\",\"type\":\"uint256\"}],\"name\":\"calculatePayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"paymentIDsList\",\"type\":\"bytes32[]\"}],\"name\":\"recordPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161108e38038061108e833981810160405281019061003191906100d4565b805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f5ffd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b610f828061010c5f395ff3fe60806040526004361061003e575f3560e01c80634e71d92d1461004257806363fcacb8146100585780638c67f63a14610074578063b69ef8a8146100b1575b5f5ffd5b34801561004d575f5ffd5b506100566100db565b005b610072600480360381019061006d9190610826565b610291565b005b34801561007f575f5ffd5b5061009a600480360381019061009591906108a4565b6104f7565b6040516100a89291906108de565b60405180910390f35b3480156100bc575f5ffd5b506100c561066d565b6040516100d29190610905565b60405180910390f35b5f5f90505f5f90505b600280549050811015610207575f600282815481106101065761010561091e565b5b905f5260205f20015490505f60015f8381526020019081526020015f20905061012f8233610786565b156101f8575f815f015411610179576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610170906109cb565b60405180910390fd5b806001015f9054906101000a900460ff16156101ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101c190610a33565b60405180910390fd5b6001816001015f6101000a81548160ff021916908315150217905550805f0154846101f59190610a7e565b93505b505080806001019150506100e4565b505f811161024a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024190610b21565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f1935050505015801561028d573d5f5f3e3d5ffd5b5050565b5f82829050116102d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102cd90610b89565b60405180910390fd5b5f3411610318576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030f90610c17565b60405180910390fd5b5f3490505f5f610327836104f7565b80935081945050505f828461033c9190610c35565b9050858590508161034d9190610c95565b915080868690508361035f9190610cc5565b1461039f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039690610d50565b60405180910390fd5b5f5f90505b868690508110156104b3575f8787838181106103c3576103c261091e565b5b90506020020135905060015f8281526020019081526020015f206001015f9054906101000a900460ff161561042d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042490610db8565b60405180910390fd5b60405180604001604052808581526020015f151581525060015f8381526020019081526020015f205f820151815f01556020820151816001015f6101000a81548160ff021916908315150217905550905050600281908060018154018082558091505060019003905f5260205f20015f90919091909150555080806001019150506103a4565b507f0d5cd0de3f5c7e52d21bc9b61c65d7eeb33051519983ede24ed680242ecb5f6d8686866040516104e793929190610e4e565b60405180910390a1505050505050565b5f5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663822fcca56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610563573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105879190610e92565b90505f816104006103ff8761059c9190610a7e565b6105a69190610c95565b6105b09190610cc5565b905060645f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc0438306040518163ffffffff1660e01b8152600401602060405180830381865afa15801561061d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106419190610e92565b8261064c9190610cc5565b6106569190610c95565b925082816106649190610a7e565b93505050915091565b5f5f5f90505f5f90505b60028054905081101561077e575f600282815481106106995761069861091e565b5b905f5260205f20015490505f60015f8381526020019081526020015f2090506106c28233610786565b1561076f575f815f01541161070c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610703906109cb565b60405180910390fd5b806001015f9054906101000a900460ff161561075d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075490610a33565b60405180910390fd5b805f01548461076c9190610a7e565b93505b50508080600101915050610677565b508091505090565b5f5f826040516020016107999190610f32565b60405160208183030381529060405280519060200120905080841491505092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126107e6576107e56107c5565b5b8235905067ffffffffffffffff811115610803576108026107c9565b5b60208301915083602082028301111561081f5761081e6107cd565b5b9250929050565b5f5f6020838503121561083c5761083b6107bd565b5b5f83013567ffffffffffffffff811115610859576108586107c1565b5b610865858286016107d1565b92509250509250929050565b5f819050919050565b61088381610871565b811461088d575f5ffd5b50565b5f8135905061089e8161087a565b92915050565b5f602082840312156108b9576108b86107bd565b5b5f6108c684828501610890565b91505092915050565b6108d881610871565b82525050565b5f6040820190506108f15f8301856108cf565b6108fe60208301846108cf565b9392505050565b5f6020820190506109185f8301846108cf565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b7f4e6f207061796d656e7420617661696c61626c6520666f7220746869732070615f8201527f796d656e74204944000000000000000000000000000000000000000000000000602082015250565b5f6109b560288361094b565b91506109c08261095b565b604082019050919050565b5f6020820190508181035f8301526109e2816109a9565b9050919050565b7f5061796d656e7420616c726561647920636c61696d65640000000000000000005f82015250565b5f610a1d60178361094b565b9150610a28826109e9565b602082019050919050565b5f6020820190508181035f830152610a4a81610a11565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610a8882610871565b9150610a9383610871565b9250828201905080821115610aab57610aaa610a51565b5b92915050565b7f4e6f207061796d656e747320617661696c61626c6520666f7220636c61696d695f8201527f6e67000000000000000000000000000000000000000000000000000000000000602082015250565b5f610b0b60228361094b565b9150610b1682610ab1565b604082019050919050565b5f6020820190508181035f830152610b3881610aff565b9050919050565b7f4174206c65617374206f6e65207061796d656e742049442072657175697265645f82015250565b5f610b7360208361094b565b9150610b7e82610b3f565b602082019050919050565b5f6020820190508181035f830152610ba081610b67565b9050919050565b7f5061796d656e74206d7573742062652067726561746572207468616e207a65725f8201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b5f610c0160218361094b565b9150610c0c82610ba7565b604082019050919050565b5f6020820190508181035f830152610c2e81610bf5565b9050919050565b5f610c3f82610871565b9150610c4a83610871565b9250828203905081811115610c6257610c61610a51565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610c9f82610871565b9150610caa83610871565b925082610cba57610cb9610c68565b5b828204905092915050565b5f610ccf82610871565b9150610cda83610871565b9250828202610ce881610871565b91508282048414831517610cff57610cfe610a51565b5b5092915050565b7f496e636f727265637420616d6f756e7420646973747269627574696f6e0000005f82015250565b5f610d3a601d8361094b565b9150610d4582610d06565b602082019050919050565b5f6020820190508181035f830152610d6781610d2e565b9050919050565b7f5061796d656e7420494420616c72656164792070726f636573736564000000005f82015250565b5f610da2601c8361094b565b9150610dad82610d6e565b602082019050919050565b5f6020820190508181035f830152610dcf81610d96565b9050919050565b5f82825260208201905092915050565b5f5ffd5b82818337505050565b5f610dfe8385610dd6565b93507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610e3157610e30610de6565b5b602083029250610e42838584610dea565b82840190509392505050565b5f6040820190508181035f830152610e67818587610df3565b9050610e7660208301846108cf565b949350505050565b5f81519050610e8c8161087a565b92915050565b5f60208284031215610ea757610ea66107bd565b5b5f610eb484828501610e7e565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ee682610ebd565b9050919050565b5f8160601b9050919050565b5f610f0382610eed565b9050919050565b5f610f1482610ef9565b9050919050565b610f2c610f2782610edc565b610f0a565b82525050565b5f610f3d8284610f1b565b6014820191508190509291505056fea26469706673582212202af506be821235a950c44129046d6b1c9ebf9514de79c58cf4616899c2a4467c64736f6c634300081c0033",
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

// RecordPayment is a paid mutator transaction binding the contract method 0x63fcacb8.
//
// Solidity: function recordPayment(bytes32[] paymentIDsList) payable returns()
func (_PaymentLedger *PaymentLedgerTransactor) RecordPayment(opts *bind.TransactOpts, paymentIDsList [][32]byte) (*types.Transaction, error) {
	return _PaymentLedger.contract.Transact(opts, "recordPayment", paymentIDsList)
}

// RecordPayment is a paid mutator transaction binding the contract method 0x63fcacb8.
//
// Solidity: function recordPayment(bytes32[] paymentIDsList) payable returns()
func (_PaymentLedger *PaymentLedgerSession) RecordPayment(paymentIDsList [][32]byte) (*types.Transaction, error) {
	return _PaymentLedger.Contract.RecordPayment(&_PaymentLedger.TransactOpts, paymentIDsList)
}

// RecordPayment is a paid mutator transaction binding the contract method 0x63fcacb8.
//
// Solidity: function recordPayment(bytes32[] paymentIDsList) payable returns()
func (_PaymentLedger *PaymentLedgerTransactorSession) RecordPayment(paymentIDsList [][32]byte) (*types.Transaction, error) {
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
	PaymentIDs  [][32]byte
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x0d5cd0de3f5c7e52d21bc9b61c65d7eeb33051519983ede24ed680242ecb5f6d.
//
// Solidity: event PaymentReceived(bytes32[] paymentIDs, uint256 totalAmount)
func (_PaymentLedger *PaymentLedgerFilterer) FilterPaymentReceived(opts *bind.FilterOpts) (*PaymentLedgerPaymentReceivedIterator, error) {

	logs, sub, err := _PaymentLedger.contract.FilterLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return &PaymentLedgerPaymentReceivedIterator{contract: _PaymentLedger.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x0d5cd0de3f5c7e52d21bc9b61c65d7eeb33051519983ede24ed680242ecb5f6d.
//
// Solidity: event PaymentReceived(bytes32[] paymentIDs, uint256 totalAmount)
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

// ParsePaymentReceived is a log parse operation binding the contract event 0x0d5cd0de3f5c7e52d21bc9b61c65d7eeb33051519983ede24ed680242ecb5f6d.
//
// Solidity: event PaymentReceived(bytes32[] paymentIDs, uint256 totalAmount)
func (_PaymentLedger *PaymentLedgerFilterer) ParsePaymentReceived(log types.Log) (*PaymentLedgerPaymentReceived, error) {
	event := new(PaymentLedgerPaymentReceived)
	if err := _PaymentLedger.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
