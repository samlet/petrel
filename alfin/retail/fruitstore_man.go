// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package retail

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FruitStoreABI is the input ABI used to generate the binding from.
const FruitStoreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"fruit\",\"type\":\"bytes\"}],\"name\":\"getStock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fruitName\",\"type\":\"bytes\"},{\"name\":\"stock\",\"type\":\"uint256\"}],\"name\":\"setFruitStock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FruitStoreBin is the compiled bytecode used for deploying new contracts.
var FruitStoreBin = "0x608060405234801561001057600080fd5b506102c0806100206000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166354f5782d811461005b5780635ef230721461008d578063b2bdfa7b146100b3575b600080fd5b34801561006757600080fd5b5061007b60048035602481019101356100f1565b60408051918252519081900360200190f35b34801561009957600080fd5b506100b1602460048035828101929101359035610120565b005b3480156100bf57600080fd5b506100c8610278565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b600060018383604051808383808284379091019485525050604051928390036020019092205495945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101a657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f417574683a206f6e6c79206f776e657220697320617574686f72697a65642e00604482015290519081900360640190fd5b82828080601f0160208091040260200160405190810160405280939291908181526020018383808284375050845160001093506102489250505057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f667275697465206e616d6520697320696e76616c696421000000000000000000604482015290519081900360640190fd5b81600185856040518083838082843782019150509250505090815260200160405180910390208190555050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a7230582029723b4d96839cd6f4561ff79ad4bb8a73db7f458328674f6333e70f29407d790029"

// DeployFruitStore deploys a new Ethereum contract, binding an instance of FruitStore to it.
func DeployFruitStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FruitStore, error) {
	parsed, err := abi.JSON(strings.NewReader(FruitStoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FruitStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FruitStore{FruitStoreCaller: FruitStoreCaller{contract: contract}, FruitStoreTransactor: FruitStoreTransactor{contract: contract}, FruitStoreFilterer: FruitStoreFilterer{contract: contract}}, nil
}

// FruitStore is an auto generated Go binding around an Ethereum contract.
type FruitStore struct {
	FruitStoreCaller     // Read-only binding to the contract
	FruitStoreTransactor // Write-only binding to the contract
	FruitStoreFilterer   // Log filterer for contract events
}

// FruitStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type FruitStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FruitStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FruitStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FruitStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FruitStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FruitStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FruitStoreSession struct {
	Contract     *FruitStore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FruitStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FruitStoreCallerSession struct {
	Contract *FruitStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FruitStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FruitStoreTransactorSession struct {
	Contract     *FruitStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FruitStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type FruitStoreRaw struct {
	Contract *FruitStore // Generic contract binding to access the raw methods on
}

// FruitStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FruitStoreCallerRaw struct {
	Contract *FruitStoreCaller // Generic read-only contract binding to access the raw methods on
}

// FruitStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FruitStoreTransactorRaw struct {
	Contract *FruitStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFruitStore creates a new instance of FruitStore, bound to a specific deployed contract.
func NewFruitStore(address common.Address, backend bind.ContractBackend) (*FruitStore, error) {
	contract, err := bindFruitStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FruitStore{FruitStoreCaller: FruitStoreCaller{contract: contract}, FruitStoreTransactor: FruitStoreTransactor{contract: contract}, FruitStoreFilterer: FruitStoreFilterer{contract: contract}}, nil
}

// NewFruitStoreCaller creates a new read-only instance of FruitStore, bound to a specific deployed contract.
func NewFruitStoreCaller(address common.Address, caller bind.ContractCaller) (*FruitStoreCaller, error) {
	contract, err := bindFruitStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FruitStoreCaller{contract: contract}, nil
}

// NewFruitStoreTransactor creates a new write-only instance of FruitStore, bound to a specific deployed contract.
func NewFruitStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*FruitStoreTransactor, error) {
	contract, err := bindFruitStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FruitStoreTransactor{contract: contract}, nil
}

// NewFruitStoreFilterer creates a new log filterer instance of FruitStore, bound to a specific deployed contract.
func NewFruitStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*FruitStoreFilterer, error) {
	contract, err := bindFruitStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FruitStoreFilterer{contract: contract}, nil
}

// bindFruitStore binds a generic wrapper to an already deployed contract.
func bindFruitStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FruitStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FruitStore *FruitStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FruitStore.Contract.FruitStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FruitStore *FruitStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FruitStore.Contract.FruitStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FruitStore *FruitStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FruitStore.Contract.FruitStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FruitStore *FruitStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FruitStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FruitStore *FruitStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FruitStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FruitStore *FruitStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FruitStore.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FruitStore *FruitStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FruitStore.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FruitStore *FruitStoreSession) Owner() (common.Address, error) {
	return _FruitStore.Contract.Owner(&_FruitStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FruitStore *FruitStoreCallerSession) Owner() (common.Address, error) {
	return _FruitStore.Contract.Owner(&_FruitStore.CallOpts)
}

// GetStock is a free data retrieval call binding the contract method 0x54f5782d.
//
// Solidity: function getStock(bytes fruit) view returns(uint256)
func (_FruitStore *FruitStoreCaller) GetStock(opts *bind.CallOpts, fruit []byte) (*big.Int, error) {
	var out []interface{}
	err := _FruitStore.contract.Call(opts, &out, "getStock", fruit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStock is a free data retrieval call binding the contract method 0x54f5782d.
//
// Solidity: function getStock(bytes fruit) view returns(uint256)
func (_FruitStore *FruitStoreSession) GetStock(fruit []byte) (*big.Int, error) {
	return _FruitStore.Contract.GetStock(&_FruitStore.CallOpts, fruit)
}

// GetStock is a free data retrieval call binding the contract method 0x54f5782d.
//
// Solidity: function getStock(bytes fruit) view returns(uint256)
func (_FruitStore *FruitStoreCallerSession) GetStock(fruit []byte) (*big.Int, error) {
	return _FruitStore.Contract.GetStock(&_FruitStore.CallOpts, fruit)
}

// SetFruitStock is a paid mutator transaction binding the contract method 0x5ef23072.
//
// Solidity: function setFruitStock(bytes fruitName, uint256 stock) returns()
func (_FruitStore *FruitStoreTransactor) SetFruitStock(opts *bind.TransactOpts, fruitName []byte, stock *big.Int) (*types.Transaction, error) {
	return _FruitStore.contract.Transact(opts, "setFruitStock", fruitName, stock)
}

// SetFruitStock is a paid mutator transaction binding the contract method 0x5ef23072.
//
// Solidity: function setFruitStock(bytes fruitName, uint256 stock) returns()
func (_FruitStore *FruitStoreSession) SetFruitStock(fruitName []byte, stock *big.Int) (*types.Transaction, error) {
	return _FruitStore.Contract.SetFruitStock(&_FruitStore.TransactOpts, fruitName, stock)
}

// SetFruitStock is a paid mutator transaction binding the contract method 0x5ef23072.
//
// Solidity: function setFruitStock(bytes fruitName, uint256 stock) returns()
func (_FruitStore *FruitStoreTransactorSession) SetFruitStock(fruitName []byte, stock *big.Int) (*types.Transaction, error) {
	return _FruitStore.Contract.SetFruitStock(&_FruitStore.TransactOpts, fruitName, stock)
}
