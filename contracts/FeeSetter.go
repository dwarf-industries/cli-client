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

// FeeSetterMetaData contains all meta data concerning the FeeSetter contract.
var FeeSetterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_networkFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_networkFeeCollector\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"changeNetworkFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCostPerKilobyte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNetworkFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNetworkFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161062f38038061062f83398181016040528101906100319190610158565b836001819055508160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600281905550805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506101bc565b5f5ffd5b5f819050919050565b6100dd816100cb565b81146100e7575f5ffd5b50565b5f815190506100f8816100d4565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610127826100fe565b9050919050565b6101378161011d565b8114610141575f5ffd5b50565b5f815190506101528161012e565b92915050565b5f5f5f5f608085870312156101705761016f6100c7565b5b5f61017d878288016100ea565b945050602061018e878288016100ea565b935050604061019f87828801610144565b92505060606101b087828801610144565b91505092959194509250565b610466806101c95f395ff3fe608060405260043610610049575f3560e01c806318a9055c1461004d5780636a1db1bf14610077578063822fcca5146100a75780638e5aa1a3146100d1578063fc04383014610101575b5f5ffd5b348015610058575f5ffd5b5061006161012b565b60405161006e91906102e2565b60405180910390f35b610091600480360381019061008c9190610332565b610153565b60405161009e9190610377565b60405180910390f35b3480156100b2575f5ffd5b506100bb6101f2565b6040516100c8919061039f565b60405180910390f35b6100eb60048036038101906100e69190610332565b6101fb565b6040516100f89190610377565b60405180910390f35b34801561010c575f5ffd5b5061011561029a565b604051610122919061039f565b60405180910390f35b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d990610412565b60405180910390fd5b8160018190555060019050919050565b5f600154905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028190610412565b60405180910390fd5b8160028190555060019050919050565b5f600254905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102cc826102a3565b9050919050565b6102dc816102c2565b82525050565b5f6020820190506102f55f8301846102d3565b92915050565b5f5ffd5b5f819050919050565b610311816102ff565b811461031b575f5ffd5b50565b5f8135905061032c81610308565b92915050565b5f60208284031215610347576103466102fb565b5b5f6103548482850161031e565b91505092915050565b5f8115159050919050565b6103718161035d565b82525050565b5f60208201905061038a5f830184610368565b92915050565b610399816102ff565b82525050565b5f6020820190506103b25f830184610390565b92915050565b5f82825260208201905092915050565b7f4e6f7420617574686f72697a65640000000000000000000000000000000000005f82015250565b5f6103fc600e836103b8565b9150610407826103c8565b602082019050919050565b5f6020820190508181035f830152610429816103f0565b905091905056fea2646970667358221220918cee82528b2c27b71d9a077f03bfe263dab70e5c4550481ec6797e75688a9f64736f6c634300081c0033",
}

// FeeSetterABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeSetterMetaData.ABI instead.
var FeeSetterABI = FeeSetterMetaData.ABI

// FeeSetterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeSetterMetaData.Bin instead.
var FeeSetterBin = FeeSetterMetaData.Bin

// DeployFeeSetter deploys a new Ethereum contract, binding an instance of FeeSetter to it.
func DeployFeeSetter(auth *bind.TransactOpts, backend bind.ContractBackend, _fee *big.Int, _networkFee *big.Int, _networkFeeCollector common.Address, dao common.Address) (common.Address, *types.Transaction, *FeeSetter, error) {
	parsed, err := FeeSetterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeSetterBin), backend, _fee, _networkFee, _networkFeeCollector, dao)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeSetter{FeeSetterCaller: FeeSetterCaller{contract: contract}, FeeSetterTransactor: FeeSetterTransactor{contract: contract}, FeeSetterFilterer: FeeSetterFilterer{contract: contract}}, nil
}

// FeeSetter is an auto generated Go binding around an Ethereum contract.
type FeeSetter struct {
	FeeSetterCaller     // Read-only binding to the contract
	FeeSetterTransactor // Write-only binding to the contract
	FeeSetterFilterer   // Log filterer for contract events
}

// FeeSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeSetterSession struct {
	Contract     *FeeSetter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeSetterCallerSession struct {
	Contract *FeeSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FeeSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeSetterTransactorSession struct {
	Contract     *FeeSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FeeSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeSetterRaw struct {
	Contract *FeeSetter // Generic contract binding to access the raw methods on
}

// FeeSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeSetterCallerRaw struct {
	Contract *FeeSetterCaller // Generic read-only contract binding to access the raw methods on
}

// FeeSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeSetterTransactorRaw struct {
	Contract *FeeSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeSetter creates a new instance of FeeSetter, bound to a specific deployed contract.
func NewFeeSetter(address common.Address, backend bind.ContractBackend) (*FeeSetter, error) {
	contract, err := bindFeeSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeSetter{FeeSetterCaller: FeeSetterCaller{contract: contract}, FeeSetterTransactor: FeeSetterTransactor{contract: contract}, FeeSetterFilterer: FeeSetterFilterer{contract: contract}}, nil
}

// NewFeeSetterCaller creates a new read-only instance of FeeSetter, bound to a specific deployed contract.
func NewFeeSetterCaller(address common.Address, caller bind.ContractCaller) (*FeeSetterCaller, error) {
	contract, err := bindFeeSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeSetterCaller{contract: contract}, nil
}

// NewFeeSetterTransactor creates a new write-only instance of FeeSetter, bound to a specific deployed contract.
func NewFeeSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeSetterTransactor, error) {
	contract, err := bindFeeSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeSetterTransactor{contract: contract}, nil
}

// NewFeeSetterFilterer creates a new log filterer instance of FeeSetter, bound to a specific deployed contract.
func NewFeeSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeSetterFilterer, error) {
	contract, err := bindFeeSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeSetterFilterer{contract: contract}, nil
}

// bindFeeSetter binds a generic wrapper to an already deployed contract.
func bindFeeSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeSetterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeSetter *FeeSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeSetter.Contract.FeeSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeSetter *FeeSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeSetter.Contract.FeeSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeSetter *FeeSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeSetter.Contract.FeeSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeSetter *FeeSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeSetter *FeeSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeSetter *FeeSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeSetter.Contract.contract.Transact(opts, method, params...)
}

// GetCostPerKilobyte is a free data retrieval call binding the contract method 0x822fcca5.
//
// Solidity: function getCostPerKilobyte() view returns(uint256)
func (_FeeSetter *FeeSetterCaller) GetCostPerKilobyte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeSetter.contract.Call(opts, &out, "getCostPerKilobyte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCostPerKilobyte is a free data retrieval call binding the contract method 0x822fcca5.
//
// Solidity: function getCostPerKilobyte() view returns(uint256)
func (_FeeSetter *FeeSetterSession) GetCostPerKilobyte() (*big.Int, error) {
	return _FeeSetter.Contract.GetCostPerKilobyte(&_FeeSetter.CallOpts)
}

// GetCostPerKilobyte is a free data retrieval call binding the contract method 0x822fcca5.
//
// Solidity: function getCostPerKilobyte() view returns(uint256)
func (_FeeSetter *FeeSetterCallerSession) GetCostPerKilobyte() (*big.Int, error) {
	return _FeeSetter.Contract.GetCostPerKilobyte(&_FeeSetter.CallOpts)
}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_FeeSetter *FeeSetterCaller) GetNetworkFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeSetter.contract.Call(opts, &out, "getNetworkFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_FeeSetter *FeeSetterSession) GetNetworkFee() (*big.Int, error) {
	return _FeeSetter.Contract.GetNetworkFee(&_FeeSetter.CallOpts)
}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_FeeSetter *FeeSetterCallerSession) GetNetworkFee() (*big.Int, error) {
	return _FeeSetter.Contract.GetNetworkFee(&_FeeSetter.CallOpts)
}

// GetNetworkFeeCollector is a free data retrieval call binding the contract method 0x18a9055c.
//
// Solidity: function getNetworkFeeCollector() view returns(address)
func (_FeeSetter *FeeSetterCaller) GetNetworkFeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeSetter.contract.Call(opts, &out, "getNetworkFeeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNetworkFeeCollector is a free data retrieval call binding the contract method 0x18a9055c.
//
// Solidity: function getNetworkFeeCollector() view returns(address)
func (_FeeSetter *FeeSetterSession) GetNetworkFeeCollector() (common.Address, error) {
	return _FeeSetter.Contract.GetNetworkFeeCollector(&_FeeSetter.CallOpts)
}

// GetNetworkFeeCollector is a free data retrieval call binding the contract method 0x18a9055c.
//
// Solidity: function getNetworkFeeCollector() view returns(address)
func (_FeeSetter *FeeSetterCallerSession) GetNetworkFeeCollector() (common.Address, error) {
	return _FeeSetter.Contract.GetNetworkFeeCollector(&_FeeSetter.CallOpts)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 fee) payable returns(bool)
func (_FeeSetter *FeeSetterTransactor) ChangeFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _FeeSetter.contract.Transact(opts, "changeFee", fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 fee) payable returns(bool)
func (_FeeSetter *FeeSetterSession) ChangeFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeSetter.Contract.ChangeFee(&_FeeSetter.TransactOpts, fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 fee) payable returns(bool)
func (_FeeSetter *FeeSetterTransactorSession) ChangeFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeSetter.Contract.ChangeFee(&_FeeSetter.TransactOpts, fee)
}

// ChangeNetworkFee is a paid mutator transaction binding the contract method 0x8e5aa1a3.
//
// Solidity: function changeNetworkFee(uint256 fee) payable returns(bool)
func (_FeeSetter *FeeSetterTransactor) ChangeNetworkFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _FeeSetter.contract.Transact(opts, "changeNetworkFee", fee)
}

// ChangeNetworkFee is a paid mutator transaction binding the contract method 0x8e5aa1a3.
//
// Solidity: function changeNetworkFee(uint256 fee) payable returns(bool)
func (_FeeSetter *FeeSetterSession) ChangeNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeSetter.Contract.ChangeNetworkFee(&_FeeSetter.TransactOpts, fee)
}

// ChangeNetworkFee is a paid mutator transaction binding the contract method 0x8e5aa1a3.
//
// Solidity: function changeNetworkFee(uint256 fee) payable returns(bool)
func (_FeeSetter *FeeSetterTransactorSession) ChangeNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeSetter.Contract.ChangeNetworkFee(&_FeeSetter.TransactOpts, fee)
}
