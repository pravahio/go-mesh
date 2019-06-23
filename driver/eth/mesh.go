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

// EthBin is the compiled bytecode used for deploying new contracts.
const EthBin = `608060405234801561001057600080fd5b506108df806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633c659743146100515780633d00ba091461012c578063ba68a2bf1461021f578063cd75072b146102fa575b600080fd5b61012a6004803603604081101561006757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156100a457600080fd5b8201836020820111156100b657600080fd5b803590602001918460018302840111640100000000831117156100d857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506103ed565b005b6102056004803603604081101561014257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561017f57600080fd5b82018360208201111561019157600080fd5b803590602001918460018302840111640100000000831117156101b357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610515565b604051808215151515815260200191505060405180910390f35b6102f86004803603604081101561023557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561027257600080fd5b82018360208201111561028457600080fd5b803590602001918460018302840111640100000000831117156102a657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061064b565b005b6103d36004803603604081101561031057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561034d57600080fd5b82018360208201111561035f57600080fd5b8035906020019184600183028401116401000000008311171561038157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610773565b604051808215151515815260200191505060405180910390f35b60016000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff0219169083600281111561044b57fe5b021790555060016000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101826040518082805190602001908083835b602083106104c757805182526020820191506020810190506020830392506104a4565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff0219169083151502179055505050565b600060028081111561052357fe5b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16600281111561057d57fe5b1480156106435750600115156000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101836040518082805190602001908083835b602083106105fe57805182526020820191506020810190506020830392506105db565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff161515145b905092915050565b60026000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908360028111156106a957fe5b021790555060016000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101826040518082805190602001908083835b602083106107255780518252602082019150602081019050602083039250610702565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff0219169083151502179055505050565b60006001600281111561078257fe5b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1660028111156107dc57fe5b1480156108a25750600115156000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101836040518082805190602001908083835b6020831061085d578051825260208201915060208101905060208303925061083a565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff161515145b90509291505056fea265627a7a72305820755ff9e9d3a8ac24af66944dbe95958afb9b34b394f6a4f1b5e77a583daa6ff264736f6c63430005090032`

// DeployEth deploys a new Ethereum contract, binding an instance of Eth to it.
func DeployEth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Eth, error) {
	parsed, err := abi.JSON(strings.NewReader(EthABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eth{EthCaller: EthCaller{contract: contract}, EthTransactor: EthTransactor{contract: contract}, EthFilterer: EthFilterer{contract: contract}}, nil
}

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
