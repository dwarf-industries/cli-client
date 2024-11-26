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

// LedgerMetaData contains all meta data concerning the Ledger contract.
var LedgerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"paymentID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"paymentId\",\"type\":\"bytes32\"}],\"name\":\"paymentProcessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processedPayments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"paymentID\",\"type\":\"bytes32\"}],\"name\":\"recordPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506104458061001c5f395ff3fe608060405260043610610033575f3560e01c80639641413814610037578063e434a1be14610073578063f3f247eb1461008f575b5f5ffd5b348015610042575f5ffd5b5061005d6004803603810190610058919061025d565b6100cb565b60405161006a91906102a2565b60405180910390f35b61008d6004803603810190610088919061025d565b6100f0565b005b34801561009a575f5ffd5b506100b560048036038101906100b0919061025d565b61020a565b6040516100c291906102a2565b60405180910390f35b5f5f5f8381526020019081526020015f205f9054906101000a900460ff169050919050565b5f3411610132576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101299061033b565b60405180910390fd5b5f5f8281526020019081526020015f205f9054906101000a900460ff161561018f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610186906103a3565b60405180910390fd5b60015f5f8381526020019081526020015f205f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f4c3e182d9e8a50f18c7d26c444315e47590087919a9e619901c6cae494e2e24d82346040516101ff9291906103e8565b60405180910390a250565b5f602052805f5260405f205f915054906101000a900460ff1681565b5f5ffd5b5f819050919050565b61023c8161022a565b8114610246575f5ffd5b50565b5f8135905061025781610233565b92915050565b5f6020828403121561027257610271610226565b5b5f61027f84828501610249565b91505092915050565b5f8115159050919050565b61029c81610288565b82525050565b5f6020820190506102b55f830184610293565b92915050565b5f82825260208201905092915050565b7f5061796d656e74206d7573742062652067726561746572207468616e207a65725f8201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b5f6103256021836102bb565b9150610330826102cb565b604082019050919050565b5f6020820190508181035f83015261035281610319565b9050919050565b7f5061796d656e7420616c72656164792070726f636573736564000000000000005f82015250565b5f61038d6019836102bb565b915061039882610359565b602082019050919050565b5f6020820190508181035f8301526103ba81610381565b9050919050565b6103ca8161022a565b82525050565b5f819050919050565b6103e2816103d0565b82525050565b5f6040820190506103fb5f8301856103c1565b61040860208301846103d9565b939250505056fea2646970667358221220979adc13985c4c507daceaa61adf8bdc8a1cd6238998708929cd5466f2d35cf364736f6c634300081c0033",
}

// LedgerABI is the input ABI used to generate the binding from.
// Deprecated: Use LedgerMetaData.ABI instead.
var LedgerABI = LedgerMetaData.ABI

// LedgerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LedgerMetaData.Bin instead.
var LedgerBin = LedgerMetaData.Bin

// DeployLedger deploys a new Ethereum contract, binding an instance of Ledger to it.
func DeployLedger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ledger, error) {
	parsed, err := LedgerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LedgerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ledger{LedgerCaller: LedgerCaller{contract: contract}, LedgerTransactor: LedgerTransactor{contract: contract}, LedgerFilterer: LedgerFilterer{contract: contract}}, nil
}

// Ledger is an auto generated Go binding around an Ethereum contract.
type Ledger struct {
	LedgerCaller     // Read-only binding to the contract
	LedgerTransactor // Write-only binding to the contract
	LedgerFilterer   // Log filterer for contract events
}

// LedgerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LedgerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LedgerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LedgerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LedgerSession struct {
	Contract     *Ledger           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LedgerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LedgerCallerSession struct {
	Contract *LedgerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LedgerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LedgerTransactorSession struct {
	Contract     *LedgerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LedgerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LedgerRaw struct {
	Contract *Ledger // Generic contract binding to access the raw methods on
}

// LedgerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LedgerCallerRaw struct {
	Contract *LedgerCaller // Generic read-only contract binding to access the raw methods on
}

// LedgerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LedgerTransactorRaw struct {
	Contract *LedgerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLedger creates a new instance of Ledger, bound to a specific deployed contract.
func NewLedger(address common.Address, backend bind.ContractBackend) (*Ledger, error) {
	contract, err := bindLedger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ledger{LedgerCaller: LedgerCaller{contract: contract}, LedgerTransactor: LedgerTransactor{contract: contract}, LedgerFilterer: LedgerFilterer{contract: contract}}, nil
}

// NewLedgerCaller creates a new read-only instance of Ledger, bound to a specific deployed contract.
func NewLedgerCaller(address common.Address, caller bind.ContractCaller) (*LedgerCaller, error) {
	contract, err := bindLedger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LedgerCaller{contract: contract}, nil
}

// NewLedgerTransactor creates a new write-only instance of Ledger, bound to a specific deployed contract.
func NewLedgerTransactor(address common.Address, transactor bind.ContractTransactor) (*LedgerTransactor, error) {
	contract, err := bindLedger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LedgerTransactor{contract: contract}, nil
}

// NewLedgerFilterer creates a new log filterer instance of Ledger, bound to a specific deployed contract.
func NewLedgerFilterer(address common.Address, filterer bind.ContractFilterer) (*LedgerFilterer, error) {
	contract, err := bindLedger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LedgerFilterer{contract: contract}, nil
}

// bindLedger binds a generic wrapper to an already deployed contract.
func bindLedger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LedgerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ledger *LedgerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ledger.Contract.LedgerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ledger *LedgerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ledger.Contract.LedgerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ledger *LedgerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ledger.Contract.LedgerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ledger *LedgerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ledger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ledger *LedgerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ledger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ledger *LedgerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ledger.Contract.contract.Transact(opts, method, params...)
}

// PaymentProcessed is a free data retrieval call binding the contract method 0x96414138.
//
// Solidity: function paymentProcessed(bytes32 paymentId) view returns(bool)
func (_Ledger *LedgerCaller) PaymentProcessed(opts *bind.CallOpts, paymentId [32]byte) (bool, error) {
	var out []interface{}
	err := _Ledger.contract.Call(opts, &out, "paymentProcessed", paymentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PaymentProcessed is a free data retrieval call binding the contract method 0x96414138.
//
// Solidity: function paymentProcessed(bytes32 paymentId) view returns(bool)
func (_Ledger *LedgerSession) PaymentProcessed(paymentId [32]byte) (bool, error) {
	return _Ledger.Contract.PaymentProcessed(&_Ledger.CallOpts, paymentId)
}

// PaymentProcessed is a free data retrieval call binding the contract method 0x96414138.
//
// Solidity: function paymentProcessed(bytes32 paymentId) view returns(bool)
func (_Ledger *LedgerCallerSession) PaymentProcessed(paymentId [32]byte) (bool, error) {
	return _Ledger.Contract.PaymentProcessed(&_Ledger.CallOpts, paymentId)
}

// ProcessedPayments is a free data retrieval call binding the contract method 0xf3f247eb.
//
// Solidity: function processedPayments(bytes32 ) view returns(bool)
func (_Ledger *LedgerCaller) ProcessedPayments(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Ledger.contract.Call(opts, &out, "processedPayments", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedPayments is a free data retrieval call binding the contract method 0xf3f247eb.
//
// Solidity: function processedPayments(bytes32 ) view returns(bool)
func (_Ledger *LedgerSession) ProcessedPayments(arg0 [32]byte) (bool, error) {
	return _Ledger.Contract.ProcessedPayments(&_Ledger.CallOpts, arg0)
}

// ProcessedPayments is a free data retrieval call binding the contract method 0xf3f247eb.
//
// Solidity: function processedPayments(bytes32 ) view returns(bool)
func (_Ledger *LedgerCallerSession) ProcessedPayments(arg0 [32]byte) (bool, error) {
	return _Ledger.Contract.ProcessedPayments(&_Ledger.CallOpts, arg0)
}

// RecordPayment is a paid mutator transaction binding the contract method 0xe434a1be.
//
// Solidity: function recordPayment(bytes32 paymentID) payable returns()
func (_Ledger *LedgerTransactor) RecordPayment(opts *bind.TransactOpts, paymentID [32]byte) (*types.Transaction, error) {
	return _Ledger.contract.Transact(opts, "recordPayment", paymentID)
}

// RecordPayment is a paid mutator transaction binding the contract method 0xe434a1be.
//
// Solidity: function recordPayment(bytes32 paymentID) payable returns()
func (_Ledger *LedgerSession) RecordPayment(paymentID [32]byte) (*types.Transaction, error) {
	return _Ledger.Contract.RecordPayment(&_Ledger.TransactOpts, paymentID)
}

// RecordPayment is a paid mutator transaction binding the contract method 0xe434a1be.
//
// Solidity: function recordPayment(bytes32 paymentID) payable returns()
func (_Ledger *LedgerTransactorSession) RecordPayment(paymentID [32]byte) (*types.Transaction, error) {
	return _Ledger.Contract.RecordPayment(&_Ledger.TransactOpts, paymentID)
}

// LedgerPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the Ledger contract.
type LedgerPaymentReceivedIterator struct {
	Event *LedgerPaymentReceived // Event containing the contract specifics and raw log

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
func (it *LedgerPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LedgerPaymentReceived)
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
		it.Event = new(LedgerPaymentReceived)
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
func (it *LedgerPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LedgerPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LedgerPaymentReceived represents a PaymentReceived event raised by the Ledger contract.
type LedgerPaymentReceived struct {
	Payer     common.Address
	PaymentID [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x4c3e182d9e8a50f18c7d26c444315e47590087919a9e619901c6cae494e2e24d.
//
// Solidity: event PaymentReceived(address indexed payer, bytes32 paymentID, uint256 amount)
func (_Ledger *LedgerFilterer) FilterPaymentReceived(opts *bind.FilterOpts, payer []common.Address) (*LedgerPaymentReceivedIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _Ledger.contract.FilterLogs(opts, "PaymentReceived", payerRule)
	if err != nil {
		return nil, err
	}
	return &LedgerPaymentReceivedIterator{contract: _Ledger.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x4c3e182d9e8a50f18c7d26c444315e47590087919a9e619901c6cae494e2e24d.
//
// Solidity: event PaymentReceived(address indexed payer, bytes32 paymentID, uint256 amount)
func (_Ledger *LedgerFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *LedgerPaymentReceived, payer []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _Ledger.contract.WatchLogs(opts, "PaymentReceived", payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LedgerPaymentReceived)
				if err := _Ledger.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
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

// ParsePaymentReceived is a log parse operation binding the contract event 0x4c3e182d9e8a50f18c7d26c444315e47590087919a9e619901c6cae494e2e24d.
//
// Solidity: event PaymentReceived(address indexed payer, bytes32 paymentID, uint256 amount)
func (_Ledger *LedgerFilterer) ParsePaymentReceived(log types.Log) (*LedgerPaymentReceived, error) {
	event := new(LedgerPaymentReceived)
	if err := _Ledger.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
