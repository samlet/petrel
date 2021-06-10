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

// BasicAuthABI is the input ABI used to generate the binding from.
const BasicAuthABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BasicAuthFuncSigs maps the 4-byte function signature to its string representation.
var BasicAuthFuncSigs = map[string]string{
	"b2bdfa7b": "_owner()",
	"13af4035": "setOwner(address)",
}

// BasicAuthBin is the compiled bytecode used for deploying new contracts.
var BasicAuthBin = "0x608060405234801561001057600080fd5b5060008054600160a060020a031916331790556101ed806100326000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af40358114610050578063b2bdfa7b14610080575b600080fd5b34801561005c57600080fd5b5061007e73ffffffffffffffffffffffffffffffffffffffff600435166100be565b005b34801561008c57600080fd5b506100956101a5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461016957604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4261736963417574683a206f6e6c79206f776e657220697320617574686f726960448201527f7a65642e00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820c97d30411dab1cfa43dbb9be9c400c042800eb6499fd3d52e9a1db51aaa1d6180029"

// DeployBasicAuth deploys a new Ethereum contract, binding an instance of BasicAuth to it.
func DeployBasicAuth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicAuth, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicAuthABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicAuthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicAuth{BasicAuthCaller: BasicAuthCaller{contract: contract}, BasicAuthTransactor: BasicAuthTransactor{contract: contract}, BasicAuthFilterer: BasicAuthFilterer{contract: contract}}, nil
}

// BasicAuth is an auto generated Go binding around an Ethereum contract.
type BasicAuth struct {
	BasicAuthCaller     // Read-only binding to the contract
	BasicAuthTransactor // Write-only binding to the contract
	BasicAuthFilterer   // Log filterer for contract events
}

// BasicAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicAuthSession struct {
	Contract     *BasicAuth        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicAuthCallerSession struct {
	Contract *BasicAuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BasicAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicAuthTransactorSession struct {
	Contract     *BasicAuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BasicAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicAuthRaw struct {
	Contract *BasicAuth // Generic contract binding to access the raw methods on
}

// BasicAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicAuthCallerRaw struct {
	Contract *BasicAuthCaller // Generic read-only contract binding to access the raw methods on
}

// BasicAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicAuthTransactorRaw struct {
	Contract *BasicAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicAuth creates a new instance of BasicAuth, bound to a specific deployed contract.
func NewBasicAuth(address common.Address, backend bind.ContractBackend) (*BasicAuth, error) {
	contract, err := bindBasicAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicAuth{BasicAuthCaller: BasicAuthCaller{contract: contract}, BasicAuthTransactor: BasicAuthTransactor{contract: contract}, BasicAuthFilterer: BasicAuthFilterer{contract: contract}}, nil
}

// NewBasicAuthCaller creates a new read-only instance of BasicAuth, bound to a specific deployed contract.
func NewBasicAuthCaller(address common.Address, caller bind.ContractCaller) (*BasicAuthCaller, error) {
	contract, err := bindBasicAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicAuthCaller{contract: contract}, nil
}

// NewBasicAuthTransactor creates a new write-only instance of BasicAuth, bound to a specific deployed contract.
func NewBasicAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicAuthTransactor, error) {
	contract, err := bindBasicAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicAuthTransactor{contract: contract}, nil
}

// NewBasicAuthFilterer creates a new log filterer instance of BasicAuth, bound to a specific deployed contract.
func NewBasicAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicAuthFilterer, error) {
	contract, err := bindBasicAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicAuthFilterer{contract: contract}, nil
}

// bindBasicAuth binds a generic wrapper to an already deployed contract.
func bindBasicAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicAuth *BasicAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicAuth.Contract.BasicAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicAuth *BasicAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicAuth.Contract.BasicAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicAuth *BasicAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicAuth.Contract.BasicAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicAuth *BasicAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicAuth *BasicAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicAuth *BasicAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicAuth.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_BasicAuth *BasicAuthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasicAuth.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_BasicAuth *BasicAuthSession) Owner() (common.Address, error) {
	return _BasicAuth.Contract.Owner(&_BasicAuth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_BasicAuth *BasicAuthCallerSession) Owner() (common.Address, error) {
	return _BasicAuth.Contract.Owner(&_BasicAuth.CallOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_BasicAuth *BasicAuthTransactor) SetOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _BasicAuth.contract.Transact(opts, "setOwner", owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_BasicAuth *BasicAuthSession) SetOwner(owner common.Address) (*types.Transaction, error) {
	return _BasicAuth.Contract.SetOwner(&_BasicAuth.TransactOpts, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_BasicAuth *BasicAuthTransactorSession) SetOwner(owner common.Address) (*types.Transaction, error) {
	return _BasicAuth.Contract.SetOwner(&_BasicAuth.TransactOpts, owner)
}

// FruitStore2ABI is the input ABI used to generate the binding from.
const FruitStore2ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"fruit\",\"type\":\"bytes\"}],\"name\":\"getStock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fruitName\",\"type\":\"bytes\"},{\"name\":\"stock\",\"type\":\"uint256\"}],\"name\":\"setFruitStock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FruitStore2FuncSigs maps the 4-byte function signature to its string representation.
var FruitStore2FuncSigs = map[string]string{
	"b2bdfa7b": "_owner()",
	"54f5782d": "getStock(bytes)",
	"5ef23072": "setFruitStock(bytes,uint256)",
	"13af4035": "setOwner(address)",
}

// FruitStore2Bin is the compiled bytecode used for deploying new contracts.
var FruitStore2Bin = "0x608060405260008054600160a060020a03191633179055610405806100256000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af4035811461006657806354f5782d146100965780635ef23072146100c8578063b2bdfa7b146100ec575b600080fd5b34801561007257600080fd5b5061009473ffffffffffffffffffffffffffffffffffffffff6004351661012a565b005b3480156100a257600080fd5b506100b66004803560248101910135610211565b60408051918252519081900360200190f35b3480156100d457600080fd5b50610094602460048035828101929101359035610240565b3480156100f857600080fd5b506101016103bd565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1633146101d557604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4261736963417574683a206f6e6c79206f776e657220697320617574686f726960448201527f7a65642e00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600060018383604051808383808284379091019485525050604051928390036020019092205495945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146102eb57604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4261736963417574683a206f6e6c79206f776e657220697320617574686f726960448201527f7a65642e00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b82828080601f01602080910402602001604051908101604052809392919081815260200183838082843750508451600010935061038d9250505057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f667275697465206e616d6520697320696e76616c696421000000000000000000604482015290519081900360640190fd5b81600185856040518083838082843782019150509250505090815260200160405180910390208190555050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a7230582001859b09faee5d8ac895c0f0abb948e6db93a80e809bba54d45226beb2dab1590029"

// DeployFruitStore2 deploys a new Ethereum contract, binding an instance of FruitStore2 to it.
func DeployFruitStore2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FruitStore2, error) {
	parsed, err := abi.JSON(strings.NewReader(FruitStore2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FruitStore2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FruitStore2{FruitStore2Caller: FruitStore2Caller{contract: contract}, FruitStore2Transactor: FruitStore2Transactor{contract: contract}, FruitStore2Filterer: FruitStore2Filterer{contract: contract}}, nil
}

// FruitStore2 is an auto generated Go binding around an Ethereum contract.
type FruitStore2 struct {
	FruitStore2Caller     // Read-only binding to the contract
	FruitStore2Transactor // Write-only binding to the contract
	FruitStore2Filterer   // Log filterer for contract events
}

// FruitStore2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FruitStore2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FruitStore2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FruitStore2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FruitStore2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FruitStore2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FruitStore2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FruitStore2Session struct {
	Contract     *FruitStore2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FruitStore2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FruitStore2CallerSession struct {
	Contract *FruitStore2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FruitStore2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FruitStore2TransactorSession struct {
	Contract     *FruitStore2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FruitStore2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FruitStore2Raw struct {
	Contract *FruitStore2 // Generic contract binding to access the raw methods on
}

// FruitStore2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FruitStore2CallerRaw struct {
	Contract *FruitStore2Caller // Generic read-only contract binding to access the raw methods on
}

// FruitStore2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FruitStore2TransactorRaw struct {
	Contract *FruitStore2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFruitStore2 creates a new instance of FruitStore2, bound to a specific deployed contract.
func NewFruitStore2(address common.Address, backend bind.ContractBackend) (*FruitStore2, error) {
	contract, err := bindFruitStore2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FruitStore2{FruitStore2Caller: FruitStore2Caller{contract: contract}, FruitStore2Transactor: FruitStore2Transactor{contract: contract}, FruitStore2Filterer: FruitStore2Filterer{contract: contract}}, nil
}

// NewFruitStore2Caller creates a new read-only instance of FruitStore2, bound to a specific deployed contract.
func NewFruitStore2Caller(address common.Address, caller bind.ContractCaller) (*FruitStore2Caller, error) {
	contract, err := bindFruitStore2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FruitStore2Caller{contract: contract}, nil
}

// NewFruitStore2Transactor creates a new write-only instance of FruitStore2, bound to a specific deployed contract.
func NewFruitStore2Transactor(address common.Address, transactor bind.ContractTransactor) (*FruitStore2Transactor, error) {
	contract, err := bindFruitStore2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FruitStore2Transactor{contract: contract}, nil
}

// NewFruitStore2Filterer creates a new log filterer instance of FruitStore2, bound to a specific deployed contract.
func NewFruitStore2Filterer(address common.Address, filterer bind.ContractFilterer) (*FruitStore2Filterer, error) {
	contract, err := bindFruitStore2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FruitStore2Filterer{contract: contract}, nil
}

// bindFruitStore2 binds a generic wrapper to an already deployed contract.
func bindFruitStore2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FruitStore2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FruitStore2 *FruitStore2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FruitStore2.Contract.FruitStore2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FruitStore2 *FruitStore2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FruitStore2.Contract.FruitStore2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FruitStore2 *FruitStore2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FruitStore2.Contract.FruitStore2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FruitStore2 *FruitStore2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FruitStore2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FruitStore2 *FruitStore2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FruitStore2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FruitStore2 *FruitStore2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FruitStore2.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FruitStore2 *FruitStore2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FruitStore2.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FruitStore2 *FruitStore2Session) Owner() (common.Address, error) {
	return _FruitStore2.Contract.Owner(&_FruitStore2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FruitStore2 *FruitStore2CallerSession) Owner() (common.Address, error) {
	return _FruitStore2.Contract.Owner(&_FruitStore2.CallOpts)
}

// GetStock is a free data retrieval call binding the contract method 0x54f5782d.
//
// Solidity: function getStock(bytes fruit) view returns(uint256)
func (_FruitStore2 *FruitStore2Caller) GetStock(opts *bind.CallOpts, fruit []byte) (*big.Int, error) {
	var out []interface{}
	err := _FruitStore2.contract.Call(opts, &out, "getStock", fruit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStock is a free data retrieval call binding the contract method 0x54f5782d.
//
// Solidity: function getStock(bytes fruit) view returns(uint256)
func (_FruitStore2 *FruitStore2Session) GetStock(fruit []byte) (*big.Int, error) {
	return _FruitStore2.Contract.GetStock(&_FruitStore2.CallOpts, fruit)
}

// GetStock is a free data retrieval call binding the contract method 0x54f5782d.
//
// Solidity: function getStock(bytes fruit) view returns(uint256)
func (_FruitStore2 *FruitStore2CallerSession) GetStock(fruit []byte) (*big.Int, error) {
	return _FruitStore2.Contract.GetStock(&_FruitStore2.CallOpts, fruit)
}

// SetFruitStock is a paid mutator transaction binding the contract method 0x5ef23072.
//
// Solidity: function setFruitStock(bytes fruitName, uint256 stock) returns()
func (_FruitStore2 *FruitStore2Transactor) SetFruitStock(opts *bind.TransactOpts, fruitName []byte, stock *big.Int) (*types.Transaction, error) {
	return _FruitStore2.contract.Transact(opts, "setFruitStock", fruitName, stock)
}

// SetFruitStock is a paid mutator transaction binding the contract method 0x5ef23072.
//
// Solidity: function setFruitStock(bytes fruitName, uint256 stock) returns()
func (_FruitStore2 *FruitStore2Session) SetFruitStock(fruitName []byte, stock *big.Int) (*types.Transaction, error) {
	return _FruitStore2.Contract.SetFruitStock(&_FruitStore2.TransactOpts, fruitName, stock)
}

// SetFruitStock is a paid mutator transaction binding the contract method 0x5ef23072.
//
// Solidity: function setFruitStock(bytes fruitName, uint256 stock) returns()
func (_FruitStore2 *FruitStore2TransactorSession) SetFruitStock(fruitName []byte, stock *big.Int) (*types.Transaction, error) {
	return _FruitStore2.Contract.SetFruitStock(&_FruitStore2.TransactOpts, fruitName, stock)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_FruitStore2 *FruitStore2Transactor) SetOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _FruitStore2.contract.Transact(opts, "setOwner", owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_FruitStore2 *FruitStore2Session) SetOwner(owner common.Address) (*types.Transaction, error) {
	return _FruitStore2.Contract.SetOwner(&_FruitStore2.TransactOpts, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_FruitStore2 *FruitStore2TransactorSession) SetOwner(owner common.Address) (*types.Transaction, error) {
	return _FruitStore2.Contract.SetOwner(&_FruitStore2.TransactOpts, owner)
}
