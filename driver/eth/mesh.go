// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthABI is the input ABI used to generate the binding from.
const EthABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"subscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"isPeerAPublisher\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"publish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"isPeerSubscribed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Eth is an auto generated Go binding around an Ethereum contract.
type Eth struct {
	EthCaller     // Read-only binding to the contract
	EthTransactor // Write-only binding to the contract
	EthFilterer   // Log filterer for contract events
}

// EthCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthSession struct {
	Contract     *Eth              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCallerSession struct {
	Contract *EthCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthTransactorSession struct {
	Contract     *EthTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthRaw struct {
	Contract *Eth // Generic contract binding to access the raw methods on
}

// EthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCallerRaw struct {
	Contract *EthCaller // Generic read-only contract binding to access the raw methods on
}

// EthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthTransactorRaw struct {
	Contract *EthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEth creates a new instance of Eth, bound to a specific deployed contract.
func NewEth(address common.Address, backend bind.ContractBackend) (*Eth, error) {
	contract, err := bindEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eth{EthCaller: EthCaller{contract: contract}, EthTransactor: EthTransactor{contract: contract}, EthFilterer: EthFilterer{contract: contract}}, nil
}

// NewEthCaller creates a new read-only instance of Eth, bound to a specific deployed contract.
func NewEthCaller(address common.Address, caller bind.ContractCaller) (*EthCaller, error) {
	contract, err := bindEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCaller{contract: contract}, nil
}

// NewEthTransactor creates a new write-only instance of Eth, bound to a specific deployed contract.
func NewEthTransactor(address common.Address, transactor bind.ContractTransactor) (*EthTransactor, error) {
	contract, err := bindEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthTransactor{contract: contract}, nil
}

// NewEthFilterer creates a new log filterer instance of Eth, bound to a specific deployed contract.
func NewEthFilterer(address common.Address, filterer bind.ContractFilterer) (*EthFilterer, error) {
	contract, err := bindEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthFilterer{contract: contract}, nil
}

// bindEth binds a generic wrapper to an already deployed contract.
func bindEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth *EthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eth.Contract.EthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth *EthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.Contract.EthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth *EthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth.Contract.EthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth *EthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth *EthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth *EthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth.Contract.contract.Transact(opts, method, params...)
}

// IsPeerAPublisher is a free data retrieval call binding the contract method 0x3d00ba09.
//
// Solidity: function isPeerAPublisher(address a, string topic) constant returns(bool)
func (_Eth *EthCaller) IsPeerAPublisher(opts *bind.CallOpts, a common.Address, topic string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "isPeerAPublisher", a, topic)
	return *ret0, err
}

// IsPeerAPublisher is a free data retrieval call binding the contract method 0x3d00ba09.
//
// Solidity: function isPeerAPublisher(address a, string topic) constant returns(bool)
func (_Eth *EthSession) IsPeerAPublisher(a common.Address, topic string) (bool, error) {
	return _Eth.Contract.IsPeerAPublisher(&_Eth.CallOpts, a, topic)
}

// IsPeerAPublisher is a free data retrieval call binding the contract method 0x3d00ba09.
//
// Solidity: function isPeerAPublisher(address a, string topic) constant returns(bool)
func (_Eth *EthCallerSession) IsPeerAPublisher(a common.Address, topic string) (bool, error) {
	return _Eth.Contract.IsPeerAPublisher(&_Eth.CallOpts, a, topic)
}

// IsPeerSubscribed is a free data retrieval call binding the contract method 0xcd75072b.
//
// Solidity: function isPeerSubscribed(address a, string topic) constant returns(bool)
func (_Eth *EthCaller) IsPeerSubscribed(opts *bind.CallOpts, a common.Address, topic string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "isPeerSubscribed", a, topic)
	return *ret0, err
}

// IsPeerSubscribed is a free data retrieval call binding the contract method 0xcd75072b.
//
// Solidity: function isPeerSubscribed(address a, string topic) constant returns(bool)
func (_Eth *EthSession) IsPeerSubscribed(a common.Address, topic string) (bool, error) {
	return _Eth.Contract.IsPeerSubscribed(&_Eth.CallOpts, a, topic)
}

// IsPeerSubscribed is a free data retrieval call binding the contract method 0xcd75072b.
//
// Solidity: function isPeerSubscribed(address a, string topic) constant returns(bool)
func (_Eth *EthCallerSession) IsPeerSubscribed(a common.Address, topic string) (bool, error) {
	return _Eth.Contract.IsPeerSubscribed(&_Eth.CallOpts, a, topic)
}

// Publish is a paid mutator transaction binding the contract method 0xba68a2bf.
//
// Solidity: function publish(address a, string topic) returns()
func (_Eth *EthTransactor) Publish(opts *bind.TransactOpts, a common.Address, topic string) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "publish", a, topic)
}

// Publish is a paid mutator transaction binding the contract method 0xba68a2bf.
//
// Solidity: function publish(address a, string topic) returns()
func (_Eth *EthSession) Publish(a common.Address, topic string) (*types.Transaction, error) {
	return _Eth.Contract.Publish(&_Eth.TransactOpts, a, topic)
}

// Publish is a paid mutator transaction binding the contract method 0xba68a2bf.
//
// Solidity: function publish(address a, string topic) returns()
func (_Eth *EthTransactorSession) Publish(a common.Address, topic string) (*types.Transaction, error) {
	return _Eth.Contract.Publish(&_Eth.TransactOpts, a, topic)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3c659743.
//
// Solidity: function subscribe(address a, string topic) returns()
func (_Eth *EthTransactor) Subscribe(opts *bind.TransactOpts, a common.Address, topic string) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "subscribe", a, topic)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3c659743.
//
// Solidity: function subscribe(address a, string topic) returns()
func (_Eth *EthSession) Subscribe(a common.Address, topic string) (*types.Transaction, error) {
	return _Eth.Contract.Subscribe(&_Eth.TransactOpts, a, topic)
}

// Subscribe is a paid mutator transaction binding the contract method 0x3c659743.
//
// Solidity: function subscribe(address a, string topic) returns()
func (_Eth *EthTransactorSession) Subscribe(a common.Address, topic string) (*types.Transaction, error) {
	return _Eth.Contract.Subscribe(&_Eth.TransactOpts, a, topic)
}
