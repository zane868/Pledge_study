// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// DebtTokenMetaData contains all meta data concerning the DebtToken contract.
var DebtTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"multiSignature\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addMinter\",\"inputs\":[{\"name\":\"_addMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delMinter\",\"inputs\":[{\"name\":\"_delMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMinter\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinterLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiSignatureAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMinter\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// DebtTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use DebtTokenMetaData.ABI instead.
var DebtTokenABI = DebtTokenMetaData.ABI

// DebtToken is an auto generated Go binding around an Ethereum contract.
type DebtToken struct {
	DebtTokenCaller     // Read-only binding to the contract
	DebtTokenTransactor // Write-only binding to the contract
	DebtTokenFilterer   // Log filterer for contract events
}

// DebtTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebtTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebtTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebtTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebtTokenSession struct {
	Contract     *DebtToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebtTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebtTokenCallerSession struct {
	Contract *DebtTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DebtTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebtTokenTransactorSession struct {
	Contract     *DebtTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DebtTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebtTokenRaw struct {
	Contract *DebtToken // Generic contract binding to access the raw methods on
}

// DebtTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebtTokenCallerRaw struct {
	Contract *DebtTokenCaller // Generic read-only contract binding to access the raw methods on
}

// DebtTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebtTokenTransactorRaw struct {
	Contract *DebtTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebtToken creates a new instance of DebtToken, bound to a specific deployed contract.
func NewDebtToken(address common.Address, backend bind.ContractBackend) (*DebtToken, error) {
	contract, err := bindDebtToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebtToken{DebtTokenCaller: DebtTokenCaller{contract: contract}, DebtTokenTransactor: DebtTokenTransactor{contract: contract}, DebtTokenFilterer: DebtTokenFilterer{contract: contract}}, nil
}

// NewDebtTokenCaller creates a new read-only instance of DebtToken, bound to a specific deployed contract.
func NewDebtTokenCaller(address common.Address, caller bind.ContractCaller) (*DebtTokenCaller, error) {
	contract, err := bindDebtToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebtTokenCaller{contract: contract}, nil
}

// NewDebtTokenTransactor creates a new write-only instance of DebtToken, bound to a specific deployed contract.
func NewDebtTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*DebtTokenTransactor, error) {
	contract, err := bindDebtToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebtTokenTransactor{contract: contract}, nil
}

// NewDebtTokenFilterer creates a new log filterer instance of DebtToken, bound to a specific deployed contract.
func NewDebtTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*DebtTokenFilterer, error) {
	contract, err := bindDebtToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebtTokenFilterer{contract: contract}, nil
}

// bindDebtToken binds a generic wrapper to an already deployed contract.
func bindDebtToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DebtTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebtToken *DebtTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebtToken.Contract.DebtTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebtToken *DebtTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebtToken.Contract.DebtTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebtToken *DebtTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebtToken.Contract.DebtTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebtToken *DebtTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebtToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebtToken *DebtTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebtToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebtToken *DebtTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebtToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DebtToken *DebtTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DebtToken *DebtTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DebtToken.Contract.Allowance(&_DebtToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DebtToken *DebtTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DebtToken.Contract.Allowance(&_DebtToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DebtToken *DebtTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DebtToken *DebtTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DebtToken.Contract.BalanceOf(&_DebtToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DebtToken *DebtTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DebtToken.Contract.BalanceOf(&_DebtToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DebtToken *DebtTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DebtToken *DebtTokenSession) Decimals() (uint8, error) {
	return _DebtToken.Contract.Decimals(&_DebtToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DebtToken *DebtTokenCallerSession) Decimals() (uint8, error) {
	return _DebtToken.Contract.Decimals(&_DebtToken.CallOpts)
}

// GetMinter is a free data retrieval call binding the contract method 0x5b7121f8.
//
// Solidity: function getMinter(uint256 _index) view returns(address)
func (_DebtToken *DebtTokenCaller) GetMinter(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "getMinter", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMinter is a free data retrieval call binding the contract method 0x5b7121f8.
//
// Solidity: function getMinter(uint256 _index) view returns(address)
func (_DebtToken *DebtTokenSession) GetMinter(_index *big.Int) (common.Address, error) {
	return _DebtToken.Contract.GetMinter(&_DebtToken.CallOpts, _index)
}

// GetMinter is a free data retrieval call binding the contract method 0x5b7121f8.
//
// Solidity: function getMinter(uint256 _index) view returns(address)
func (_DebtToken *DebtTokenCallerSession) GetMinter(_index *big.Int) (common.Address, error) {
	return _DebtToken.Contract.GetMinter(&_DebtToken.CallOpts, _index)
}

// GetMinterLength is a free data retrieval call binding the contract method 0x0323aac7.
//
// Solidity: function getMinterLength() view returns(uint256)
func (_DebtToken *DebtTokenCaller) GetMinterLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "getMinterLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinterLength is a free data retrieval call binding the contract method 0x0323aac7.
//
// Solidity: function getMinterLength() view returns(uint256)
func (_DebtToken *DebtTokenSession) GetMinterLength() (*big.Int, error) {
	return _DebtToken.Contract.GetMinterLength(&_DebtToken.CallOpts)
}

// GetMinterLength is a free data retrieval call binding the contract method 0x0323aac7.
//
// Solidity: function getMinterLength() view returns(uint256)
func (_DebtToken *DebtTokenCallerSession) GetMinterLength() (*big.Int, error) {
	return _DebtToken.Contract.GetMinterLength(&_DebtToken.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_DebtToken *DebtTokenCaller) GetMultiSignatureAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "getMultiSignatureAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_DebtToken *DebtTokenSession) GetMultiSignatureAddress() (common.Address, error) {
	return _DebtToken.Contract.GetMultiSignatureAddress(&_DebtToken.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_DebtToken *DebtTokenCallerSession) GetMultiSignatureAddress() (common.Address, error) {
	return _DebtToken.Contract.GetMultiSignatureAddress(&_DebtToken.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_DebtToken *DebtTokenCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_DebtToken *DebtTokenSession) IsMinter(account common.Address) (bool, error) {
	return _DebtToken.Contract.IsMinter(&_DebtToken.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_DebtToken *DebtTokenCallerSession) IsMinter(account common.Address) (bool, error) {
	return _DebtToken.Contract.IsMinter(&_DebtToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DebtToken *DebtTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DebtToken *DebtTokenSession) Name() (string, error) {
	return _DebtToken.Contract.Name(&_DebtToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DebtToken *DebtTokenCallerSession) Name() (string, error) {
	return _DebtToken.Contract.Name(&_DebtToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DebtToken *DebtTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DebtToken *DebtTokenSession) Symbol() (string, error) {
	return _DebtToken.Contract.Symbol(&_DebtToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DebtToken *DebtTokenCallerSession) Symbol() (string, error) {
	return _DebtToken.Contract.Symbol(&_DebtToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DebtToken *DebtTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DebtToken *DebtTokenSession) TotalSupply() (*big.Int, error) {
	return _DebtToken.Contract.TotalSupply(&_DebtToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DebtToken *DebtTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _DebtToken.Contract.TotalSupply(&_DebtToken.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _addMinter) returns(bool)
func (_DebtToken *DebtTokenTransactor) AddMinter(opts *bind.TransactOpts, _addMinter common.Address) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "addMinter", _addMinter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _addMinter) returns(bool)
func (_DebtToken *DebtTokenSession) AddMinter(_addMinter common.Address) (*types.Transaction, error) {
	return _DebtToken.Contract.AddMinter(&_DebtToken.TransactOpts, _addMinter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _addMinter) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) AddMinter(_addMinter common.Address) (*types.Transaction, error) {
	return _DebtToken.Contract.AddMinter(&_DebtToken.TransactOpts, _addMinter)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Approve(&_DebtToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Approve(&_DebtToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns(bool)
func (_DebtToken *DebtTokenTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns(bool)
func (_DebtToken *DebtTokenSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Burn(&_DebtToken.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Burn(&_DebtToken.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DebtToken *DebtTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DebtToken *DebtTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.DecreaseAllowance(&_DebtToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.DecreaseAllowance(&_DebtToken.TransactOpts, spender, subtractedValue)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address _delMinter) returns(bool)
func (_DebtToken *DebtTokenTransactor) DelMinter(opts *bind.TransactOpts, _delMinter common.Address) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "delMinter", _delMinter)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address _delMinter) returns(bool)
func (_DebtToken *DebtTokenSession) DelMinter(_delMinter common.Address) (*types.Transaction, error) {
	return _DebtToken.Contract.DelMinter(&_DebtToken.TransactOpts, _delMinter)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address _delMinter) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) DelMinter(_delMinter common.Address) (*types.Transaction, error) {
	return _DebtToken.Contract.DelMinter(&_DebtToken.TransactOpts, _delMinter)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DebtToken *DebtTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DebtToken *DebtTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.IncreaseAllowance(&_DebtToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.IncreaseAllowance(&_DebtToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_DebtToken *DebtTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_DebtToken *DebtTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Mint(&_DebtToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Mint(&_DebtToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Transfer(&_DebtToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Transfer(&_DebtToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.TransferFrom(&_DebtToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_DebtToken *DebtTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.TransferFrom(&_DebtToken.TransactOpts, from, to, amount)
}

// DebtTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DebtToken contract.
type DebtTokenApprovalIterator struct {
	Event *DebtTokenApproval // Event containing the contract specifics and raw log

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
func (it *DebtTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtTokenApproval)
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
		it.Event = new(DebtTokenApproval)
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
func (it *DebtTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtTokenApproval represents a Approval event raised by the DebtToken contract.
type DebtTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DebtToken *DebtTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DebtTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DebtToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DebtTokenApprovalIterator{contract: _DebtToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DebtToken *DebtTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DebtTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DebtToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtTokenApproval)
				if err := _DebtToken.contract.UnpackLog(event, "Approval", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DebtToken *DebtTokenFilterer) ParseApproval(log types.Log) (*DebtTokenApproval, error) {
	event := new(DebtTokenApproval)
	if err := _DebtToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebtTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DebtToken contract.
type DebtTokenTransferIterator struct {
	Event *DebtTokenTransfer // Event containing the contract specifics and raw log

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
func (it *DebtTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtTokenTransfer)
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
		it.Event = new(DebtTokenTransfer)
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
func (it *DebtTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtTokenTransfer represents a Transfer event raised by the DebtToken contract.
type DebtTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DebtToken *DebtTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DebtTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DebtToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DebtTokenTransferIterator{contract: _DebtToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DebtToken *DebtTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DebtTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DebtToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtTokenTransfer)
				if err := _DebtToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DebtToken *DebtTokenFilterer) ParseTransfer(log types.Log) (*DebtTokenTransfer, error) {
	event := new(DebtTokenTransfer)
	if err := _DebtToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
