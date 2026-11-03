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

// MultiSignatureMetaData contains all meta data concerning the MultiSignature contract.
var MultiSignatureMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"limitedSignNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createApplication\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApplicationCount\",\"inputs\":[{\"name\":\"msghash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApplicationHash\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getApplicationInfo\",\"inputs\":[{\"name\":\"msghash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiSignatureAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidSignature\",\"inputs\":[{\"name\":\"msghash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lastIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeSignApplication\",\"inputs\":[{\"name\":\"msghash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signApplication\",\"inputs\":[{\"name\":\"msghash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signatureMap\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"applicant\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"signatureOwners\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"threshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwner\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CreateApplication\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevokeApplication\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignApplication\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// MultiSignatureABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiSignatureMetaData.ABI instead.
var MultiSignatureABI = MultiSignatureMetaData.ABI

// MultiSignature is an auto generated Go binding around an Ethereum contract.
type MultiSignature struct {
	MultiSignatureCaller     // Read-only binding to the contract
	MultiSignatureTransactor // Write-only binding to the contract
	MultiSignatureFilterer   // Log filterer for contract events
}

// MultiSignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSignatureSession struct {
	Contract     *MultiSignature   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiSignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSignatureCallerSession struct {
	Contract *MultiSignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MultiSignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSignatureTransactorSession struct {
	Contract     *MultiSignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MultiSignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSignatureRaw struct {
	Contract *MultiSignature // Generic contract binding to access the raw methods on
}

// MultiSignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSignatureCallerRaw struct {
	Contract *MultiSignatureCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSignatureTransactorRaw struct {
	Contract *MultiSignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSignature creates a new instance of MultiSignature, bound to a specific deployed contract.
func NewMultiSignature(address common.Address, backend bind.ContractBackend) (*MultiSignature, error) {
	contract, err := bindMultiSignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSignature{MultiSignatureCaller: MultiSignatureCaller{contract: contract}, MultiSignatureTransactor: MultiSignatureTransactor{contract: contract}, MultiSignatureFilterer: MultiSignatureFilterer{contract: contract}}, nil
}

// NewMultiSignatureCaller creates a new read-only instance of MultiSignature, bound to a specific deployed contract.
func NewMultiSignatureCaller(address common.Address, caller bind.ContractCaller) (*MultiSignatureCaller, error) {
	contract, err := bindMultiSignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSignatureCaller{contract: contract}, nil
}

// NewMultiSignatureTransactor creates a new write-only instance of MultiSignature, bound to a specific deployed contract.
func NewMultiSignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSignatureTransactor, error) {
	contract, err := bindMultiSignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSignatureTransactor{contract: contract}, nil
}

// NewMultiSignatureFilterer creates a new log filterer instance of MultiSignature, bound to a specific deployed contract.
func NewMultiSignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSignatureFilterer, error) {
	contract, err := bindMultiSignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSignatureFilterer{contract: contract}, nil
}

// bindMultiSignature binds a generic wrapper to an already deployed contract.
func bindMultiSignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiSignatureMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSignature *MultiSignatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSignature.Contract.MultiSignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSignature *MultiSignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSignature.Contract.MultiSignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSignature *MultiSignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSignature.Contract.MultiSignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSignature *MultiSignatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSignature *MultiSignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSignature *MultiSignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSignature.Contract.contract.Transact(opts, method, params...)
}

// GetApplicationCount is a free data retrieval call binding the contract method 0x7000823d.
//
// Solidity: function getApplicationCount(bytes32 msghash) view returns(uint256)
func (_MultiSignature *MultiSignatureCaller) GetApplicationCount(opts *bind.CallOpts, msghash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MultiSignature.contract.Call(opts, &out, "getApplicationCount", msghash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetApplicationCount is a free data retrieval call binding the contract method 0x7000823d.
//
// Solidity: function getApplicationCount(bytes32 msghash) view returns(uint256)
func (_MultiSignature *MultiSignatureSession) GetApplicationCount(msghash [32]byte) (*big.Int, error) {
	return _MultiSignature.Contract.GetApplicationCount(&_MultiSignature.CallOpts, msghash)
}

// GetApplicationCount is a free data retrieval call binding the contract method 0x7000823d.
//
// Solidity: function getApplicationCount(bytes32 msghash) view returns(uint256)
func (_MultiSignature *MultiSignatureCallerSession) GetApplicationCount(msghash [32]byte) (*big.Int, error) {
	return _MultiSignature.Contract.GetApplicationCount(&_MultiSignature.CallOpts, msghash)
}

// GetApplicationHash is a free data retrieval call binding the contract method 0x39270196.
//
// Solidity: function getApplicationHash(address from, address to) pure returns(bytes32)
func (_MultiSignature *MultiSignatureCaller) GetApplicationHash(opts *bind.CallOpts, from common.Address, to common.Address) ([32]byte, error) {
	var out []interface{}
	err := _MultiSignature.contract.Call(opts, &out, "getApplicationHash", from, to)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetApplicationHash is a free data retrieval call binding the contract method 0x39270196.
//
// Solidity: function getApplicationHash(address from, address to) pure returns(bytes32)
func (_MultiSignature *MultiSignatureSession) GetApplicationHash(from common.Address, to common.Address) ([32]byte, error) {
	return _MultiSignature.Contract.GetApplicationHash(&_MultiSignature.CallOpts, from, to)
}

// GetApplicationHash is a free data retrieval call binding the contract method 0x39270196.
//
// Solidity: function getApplicationHash(address from, address to) pure returns(bytes32)
func (_MultiSignature *MultiSignatureCallerSession) GetApplicationHash(from common.Address, to common.Address) ([32]byte, error) {
	return _MultiSignature.Contract.GetApplicationHash(&_MultiSignature.CallOpts, from, to)
}

// GetApplicationInfo is a free data retrieval call binding the contract method 0xdf18ec8f.
//
// Solidity: function getApplicationInfo(bytes32 msghash, uint256 index) view returns(address, address[])
func (_MultiSignature *MultiSignatureCaller) GetApplicationInfo(opts *bind.CallOpts, msghash [32]byte, index *big.Int) (common.Address, []common.Address, error) {
	var out []interface{}
	err := _MultiSignature.contract.Call(opts, &out, "getApplicationInfo", msghash, index)

	if err != nil {
		return *new(common.Address), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetApplicationInfo is a free data retrieval call binding the contract method 0xdf18ec8f.
//
// Solidity: function getApplicationInfo(bytes32 msghash, uint256 index) view returns(address, address[])
func (_MultiSignature *MultiSignatureSession) GetApplicationInfo(msghash [32]byte, index *big.Int) (common.Address, []common.Address, error) {
	return _MultiSignature.Contract.GetApplicationInfo(&_MultiSignature.CallOpts, msghash, index)
}

// GetApplicationInfo is a free data retrieval call binding the contract method 0xdf18ec8f.
//
// Solidity: function getApplicationInfo(bytes32 msghash, uint256 index) view returns(address, address[])
func (_MultiSignature *MultiSignatureCallerSession) GetApplicationInfo(msghash [32]byte, index *big.Int) (common.Address, []common.Address, error) {
	return _MultiSignature.Contract.GetApplicationInfo(&_MultiSignature.CallOpts, msghash, index)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_MultiSignature *MultiSignatureCaller) GetMultiSignatureAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiSignature.contract.Call(opts, &out, "getMultiSignatureAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_MultiSignature *MultiSignatureSession) GetMultiSignatureAddress() (common.Address, error) {
	return _MultiSignature.Contract.GetMultiSignatureAddress(&_MultiSignature.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_MultiSignature *MultiSignatureCallerSession) GetMultiSignatureAddress() (common.Address, error) {
	return _MultiSignature.Contract.GetMultiSignatureAddress(&_MultiSignature.CallOpts)
}

// GetValidSignature is a free data retrieval call binding the contract method 0x1ebaa166.
//
// Solidity: function getValidSignature(bytes32 msghash, uint256 lastIndex) view returns(uint256)
func (_MultiSignature *MultiSignatureCaller) GetValidSignature(opts *bind.CallOpts, msghash [32]byte, lastIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MultiSignature.contract.Call(opts, &out, "getValidSignature", msghash, lastIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidSignature is a free data retrieval call binding the contract method 0x1ebaa166.
//
// Solidity: function getValidSignature(bytes32 msghash, uint256 lastIndex) view returns(uint256)
func (_MultiSignature *MultiSignatureSession) GetValidSignature(msghash [32]byte, lastIndex *big.Int) (*big.Int, error) {
	return _MultiSignature.Contract.GetValidSignature(&_MultiSignature.CallOpts, msghash, lastIndex)
}

// GetValidSignature is a free data retrieval call binding the contract method 0x1ebaa166.
//
// Solidity: function getValidSignature(bytes32 msghash, uint256 lastIndex) view returns(uint256)
func (_MultiSignature *MultiSignatureCallerSession) GetValidSignature(msghash [32]byte, lastIndex *big.Int) (*big.Int, error) {
	return _MultiSignature.Contract.GetValidSignature(&_MultiSignature.CallOpts, msghash, lastIndex)
}

// SignatureMap is a free data retrieval call binding the contract method 0xcf195165.
//
// Solidity: function signatureMap(bytes32 , uint256 ) view returns(address applicant)
func (_MultiSignature *MultiSignatureCaller) SignatureMap(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MultiSignature.contract.Call(opts, &out, "signatureMap", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignatureMap is a free data retrieval call binding the contract method 0xcf195165.
//
// Solidity: function signatureMap(bytes32 , uint256 ) view returns(address applicant)
func (_MultiSignature *MultiSignatureSession) SignatureMap(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _MultiSignature.Contract.SignatureMap(&_MultiSignature.CallOpts, arg0, arg1)
}

// SignatureMap is a free data retrieval call binding the contract method 0xcf195165.
//
// Solidity: function signatureMap(bytes32 , uint256 ) view returns(address applicant)
func (_MultiSignature *MultiSignatureCallerSession) SignatureMap(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _MultiSignature.Contract.SignatureMap(&_MultiSignature.CallOpts, arg0, arg1)
}

// SignatureOwners is a free data retrieval call binding the contract method 0x5e63ff3f.
//
// Solidity: function signatureOwners(uint256 ) view returns(address)
func (_MultiSignature *MultiSignatureCaller) SignatureOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MultiSignature.contract.Call(opts, &out, "signatureOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignatureOwners is a free data retrieval call binding the contract method 0x5e63ff3f.
//
// Solidity: function signatureOwners(uint256 ) view returns(address)
func (_MultiSignature *MultiSignatureSession) SignatureOwners(arg0 *big.Int) (common.Address, error) {
	return _MultiSignature.Contract.SignatureOwners(&_MultiSignature.CallOpts, arg0)
}

// SignatureOwners is a free data retrieval call binding the contract method 0x5e63ff3f.
//
// Solidity: function signatureOwners(uint256 ) view returns(address)
func (_MultiSignature *MultiSignatureCallerSession) SignatureOwners(arg0 *big.Int) (common.Address, error) {
	return _MultiSignature.Contract.SignatureOwners(&_MultiSignature.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_MultiSignature *MultiSignatureCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSignature.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_MultiSignature *MultiSignatureSession) Threshold() (*big.Int, error) {
	return _MultiSignature.Contract.Threshold(&_MultiSignature.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_MultiSignature *MultiSignatureCallerSession) Threshold() (*big.Int, error) {
	return _MultiSignature.Contract.Threshold(&_MultiSignature.CallOpts)
}

// CreateApplication is a paid mutator transaction binding the contract method 0x35157be1.
//
// Solidity: function createApplication(address to) returns(uint256)
func (_MultiSignature *MultiSignatureTransactor) CreateApplication(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MultiSignature.contract.Transact(opts, "createApplication", to)
}

// CreateApplication is a paid mutator transaction binding the contract method 0x35157be1.
//
// Solidity: function createApplication(address to) returns(uint256)
func (_MultiSignature *MultiSignatureSession) CreateApplication(to common.Address) (*types.Transaction, error) {
	return _MultiSignature.Contract.CreateApplication(&_MultiSignature.TransactOpts, to)
}

// CreateApplication is a paid mutator transaction binding the contract method 0x35157be1.
//
// Solidity: function createApplication(address to) returns(uint256)
func (_MultiSignature *MultiSignatureTransactorSession) CreateApplication(to common.Address) (*types.Transaction, error) {
	return _MultiSignature.Contract.CreateApplication(&_MultiSignature.TransactOpts, to)
}

// RevokeSignApplication is a paid mutator transaction binding the contract method 0x25675059.
//
// Solidity: function revokeSignApplication(bytes32 msghash) returns()
func (_MultiSignature *MultiSignatureTransactor) RevokeSignApplication(opts *bind.TransactOpts, msghash [32]byte) (*types.Transaction, error) {
	return _MultiSignature.contract.Transact(opts, "revokeSignApplication", msghash)
}

// RevokeSignApplication is a paid mutator transaction binding the contract method 0x25675059.
//
// Solidity: function revokeSignApplication(bytes32 msghash) returns()
func (_MultiSignature *MultiSignatureSession) RevokeSignApplication(msghash [32]byte) (*types.Transaction, error) {
	return _MultiSignature.Contract.RevokeSignApplication(&_MultiSignature.TransactOpts, msghash)
}

// RevokeSignApplication is a paid mutator transaction binding the contract method 0x25675059.
//
// Solidity: function revokeSignApplication(bytes32 msghash) returns()
func (_MultiSignature *MultiSignatureTransactorSession) RevokeSignApplication(msghash [32]byte) (*types.Transaction, error) {
	return _MultiSignature.Contract.RevokeSignApplication(&_MultiSignature.TransactOpts, msghash)
}

// SignApplication is a paid mutator transaction binding the contract method 0x665bbfde.
//
// Solidity: function signApplication(bytes32 msghash) returns()
func (_MultiSignature *MultiSignatureTransactor) SignApplication(opts *bind.TransactOpts, msghash [32]byte) (*types.Transaction, error) {
	return _MultiSignature.contract.Transact(opts, "signApplication", msghash)
}

// SignApplication is a paid mutator transaction binding the contract method 0x665bbfde.
//
// Solidity: function signApplication(bytes32 msghash) returns()
func (_MultiSignature *MultiSignatureSession) SignApplication(msghash [32]byte) (*types.Transaction, error) {
	return _MultiSignature.Contract.SignApplication(&_MultiSignature.TransactOpts, msghash)
}

// SignApplication is a paid mutator transaction binding the contract method 0x665bbfde.
//
// Solidity: function signApplication(bytes32 msghash) returns()
func (_MultiSignature *MultiSignatureTransactorSession) SignApplication(msghash [32]byte) (*types.Transaction, error) {
	return _MultiSignature.Contract.SignApplication(&_MultiSignature.TransactOpts, msghash)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x1ebe85ba.
//
// Solidity: function transferOwner(uint256 index, address newOwner) returns()
func (_MultiSignature *MultiSignatureTransactor) TransferOwner(opts *bind.TransactOpts, index *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSignature.contract.Transact(opts, "transferOwner", index, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x1ebe85ba.
//
// Solidity: function transferOwner(uint256 index, address newOwner) returns()
func (_MultiSignature *MultiSignatureSession) TransferOwner(index *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSignature.Contract.TransferOwner(&_MultiSignature.TransactOpts, index, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x1ebe85ba.
//
// Solidity: function transferOwner(uint256 index, address newOwner) returns()
func (_MultiSignature *MultiSignatureTransactorSession) TransferOwner(index *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSignature.Contract.TransferOwner(&_MultiSignature.TransactOpts, index, newOwner)
}

// MultiSignatureCreateApplicationIterator is returned from FilterCreateApplication and is used to iterate over the raw logs and unpacked data for CreateApplication events raised by the MultiSignature contract.
type MultiSignatureCreateApplicationIterator struct {
	Event *MultiSignatureCreateApplication // Event containing the contract specifics and raw log

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
func (it *MultiSignatureCreateApplicationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSignatureCreateApplication)
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
		it.Event = new(MultiSignatureCreateApplication)
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
func (it *MultiSignatureCreateApplicationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSignatureCreateApplicationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSignatureCreateApplication represents a CreateApplication event raised by the MultiSignature contract.
type MultiSignatureCreateApplication struct {
	From    common.Address
	To      common.Address
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateApplication is a free log retrieval operation binding the contract event 0x313df56710bfec73d53fcb8e8d48ff821d5b2e9a113aad1e9bfc7997ef5a0d7f.
//
// Solidity: event CreateApplication(address indexed from, address indexed to, bytes32 indexed msgHash)
func (_MultiSignature *MultiSignatureFilterer) FilterCreateApplication(opts *bind.FilterOpts, from []common.Address, to []common.Address, msgHash [][32]byte) (*MultiSignatureCreateApplicationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _MultiSignature.contract.FilterLogs(opts, "CreateApplication", fromRule, toRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return &MultiSignatureCreateApplicationIterator{contract: _MultiSignature.contract, event: "CreateApplication", logs: logs, sub: sub}, nil
}

// WatchCreateApplication is a free log subscription operation binding the contract event 0x313df56710bfec73d53fcb8e8d48ff821d5b2e9a113aad1e9bfc7997ef5a0d7f.
//
// Solidity: event CreateApplication(address indexed from, address indexed to, bytes32 indexed msgHash)
func (_MultiSignature *MultiSignatureFilterer) WatchCreateApplication(opts *bind.WatchOpts, sink chan<- *MultiSignatureCreateApplication, from []common.Address, to []common.Address, msgHash [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _MultiSignature.contract.WatchLogs(opts, "CreateApplication", fromRule, toRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSignatureCreateApplication)
				if err := _MultiSignature.contract.UnpackLog(event, "CreateApplication", log); err != nil {
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

// ParseCreateApplication is a log parse operation binding the contract event 0x313df56710bfec73d53fcb8e8d48ff821d5b2e9a113aad1e9bfc7997ef5a0d7f.
//
// Solidity: event CreateApplication(address indexed from, address indexed to, bytes32 indexed msgHash)
func (_MultiSignature *MultiSignatureFilterer) ParseCreateApplication(log types.Log) (*MultiSignatureCreateApplication, error) {
	event := new(MultiSignatureCreateApplication)
	if err := _MultiSignature.contract.UnpackLog(event, "CreateApplication", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSignatureRevokeApplicationIterator is returned from FilterRevokeApplication and is used to iterate over the raw logs and unpacked data for RevokeApplication events raised by the MultiSignature contract.
type MultiSignatureRevokeApplicationIterator struct {
	Event *MultiSignatureRevokeApplication // Event containing the contract specifics and raw log

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
func (it *MultiSignatureRevokeApplicationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSignatureRevokeApplication)
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
		it.Event = new(MultiSignatureRevokeApplication)
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
func (it *MultiSignatureRevokeApplicationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSignatureRevokeApplicationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSignatureRevokeApplication represents a RevokeApplication event raised by the MultiSignature contract.
type MultiSignatureRevokeApplication struct {
	From    common.Address
	MsgHash [32]byte
	Index   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevokeApplication is a free log retrieval operation binding the contract event 0xfb7251ba5d43bdf2193d4616dad278fd04102e9fd2291bb9f95510a03f10f8ce.
//
// Solidity: event RevokeApplication(address indexed from, bytes32 indexed msgHash, uint256 index)
func (_MultiSignature *MultiSignatureFilterer) FilterRevokeApplication(opts *bind.FilterOpts, from []common.Address, msgHash [][32]byte) (*MultiSignatureRevokeApplicationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _MultiSignature.contract.FilterLogs(opts, "RevokeApplication", fromRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return &MultiSignatureRevokeApplicationIterator{contract: _MultiSignature.contract, event: "RevokeApplication", logs: logs, sub: sub}, nil
}

// WatchRevokeApplication is a free log subscription operation binding the contract event 0xfb7251ba5d43bdf2193d4616dad278fd04102e9fd2291bb9f95510a03f10f8ce.
//
// Solidity: event RevokeApplication(address indexed from, bytes32 indexed msgHash, uint256 index)
func (_MultiSignature *MultiSignatureFilterer) WatchRevokeApplication(opts *bind.WatchOpts, sink chan<- *MultiSignatureRevokeApplication, from []common.Address, msgHash [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _MultiSignature.contract.WatchLogs(opts, "RevokeApplication", fromRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSignatureRevokeApplication)
				if err := _MultiSignature.contract.UnpackLog(event, "RevokeApplication", log); err != nil {
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

// ParseRevokeApplication is a log parse operation binding the contract event 0xfb7251ba5d43bdf2193d4616dad278fd04102e9fd2291bb9f95510a03f10f8ce.
//
// Solidity: event RevokeApplication(address indexed from, bytes32 indexed msgHash, uint256 index)
func (_MultiSignature *MultiSignatureFilterer) ParseRevokeApplication(log types.Log) (*MultiSignatureRevokeApplication, error) {
	event := new(MultiSignatureRevokeApplication)
	if err := _MultiSignature.contract.UnpackLog(event, "RevokeApplication", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSignatureSignApplicationIterator is returned from FilterSignApplication and is used to iterate over the raw logs and unpacked data for SignApplication events raised by the MultiSignature contract.
type MultiSignatureSignApplicationIterator struct {
	Event *MultiSignatureSignApplication // Event containing the contract specifics and raw log

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
func (it *MultiSignatureSignApplicationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSignatureSignApplication)
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
		it.Event = new(MultiSignatureSignApplication)
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
func (it *MultiSignatureSignApplicationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSignatureSignApplicationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSignatureSignApplication represents a SignApplication event raised by the MultiSignature contract.
type MultiSignatureSignApplication struct {
	From    common.Address
	MsgHash [32]byte
	Index   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignApplication is a free log retrieval operation binding the contract event 0x2d579b8389bf07e7ce434f3bd88b036a33d28787d19af73a5aa2003a5dcb7d36.
//
// Solidity: event SignApplication(address indexed from, bytes32 indexed msgHash, uint256 index)
func (_MultiSignature *MultiSignatureFilterer) FilterSignApplication(opts *bind.FilterOpts, from []common.Address, msgHash [][32]byte) (*MultiSignatureSignApplicationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _MultiSignature.contract.FilterLogs(opts, "SignApplication", fromRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return &MultiSignatureSignApplicationIterator{contract: _MultiSignature.contract, event: "SignApplication", logs: logs, sub: sub}, nil
}

// WatchSignApplication is a free log subscription operation binding the contract event 0x2d579b8389bf07e7ce434f3bd88b036a33d28787d19af73a5aa2003a5dcb7d36.
//
// Solidity: event SignApplication(address indexed from, bytes32 indexed msgHash, uint256 index)
func (_MultiSignature *MultiSignatureFilterer) WatchSignApplication(opts *bind.WatchOpts, sink chan<- *MultiSignatureSignApplication, from []common.Address, msgHash [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _MultiSignature.contract.WatchLogs(opts, "SignApplication", fromRule, msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSignatureSignApplication)
				if err := _MultiSignature.contract.UnpackLog(event, "SignApplication", log); err != nil {
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

// ParseSignApplication is a log parse operation binding the contract event 0x2d579b8389bf07e7ce434f3bd88b036a33d28787d19af73a5aa2003a5dcb7d36.
//
// Solidity: event SignApplication(address indexed from, bytes32 indexed msgHash, uint256 index)
func (_MultiSignature *MultiSignatureFilterer) ParseSignApplication(log types.Log) (*MultiSignatureSignApplication, error) {
	event := new(MultiSignatureSignApplication)
	if err := _MultiSignature.contract.UnpackLog(event, "SignApplication", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSignatureTransferOwnerIterator is returned from FilterTransferOwner and is used to iterate over the raw logs and unpacked data for TransferOwner events raised by the MultiSignature contract.
type MultiSignatureTransferOwnerIterator struct {
	Event *MultiSignatureTransferOwner // Event containing the contract specifics and raw log

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
func (it *MultiSignatureTransferOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSignatureTransferOwner)
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
		it.Event = new(MultiSignatureTransferOwner)
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
func (it *MultiSignatureTransferOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSignatureTransferOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSignatureTransferOwner represents a TransferOwner event raised by the MultiSignature contract.
type MultiSignatureTransferOwner struct {
	Sender   common.Address
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferOwner is a free log retrieval operation binding the contract event 0xf48dd64d794d3623209edddaa975cfeae5b0ef403fbce64d7b377ceb50396c74.
//
// Solidity: event TransferOwner(address indexed sender, address indexed oldOwner, address indexed newOwner)
func (_MultiSignature *MultiSignatureFilterer) FilterTransferOwner(opts *bind.FilterOpts, sender []common.Address, oldOwner []common.Address, newOwner []common.Address) (*MultiSignatureTransferOwnerIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultiSignature.contract.FilterLogs(opts, "TransferOwner", senderRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSignatureTransferOwnerIterator{contract: _MultiSignature.contract, event: "TransferOwner", logs: logs, sub: sub}, nil
}

// WatchTransferOwner is a free log subscription operation binding the contract event 0xf48dd64d794d3623209edddaa975cfeae5b0ef403fbce64d7b377ceb50396c74.
//
// Solidity: event TransferOwner(address indexed sender, address indexed oldOwner, address indexed newOwner)
func (_MultiSignature *MultiSignatureFilterer) WatchTransferOwner(opts *bind.WatchOpts, sink chan<- *MultiSignatureTransferOwner, sender []common.Address, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultiSignature.contract.WatchLogs(opts, "TransferOwner", senderRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSignatureTransferOwner)
				if err := _MultiSignature.contract.UnpackLog(event, "TransferOwner", log); err != nil {
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

// ParseTransferOwner is a log parse operation binding the contract event 0xf48dd64d794d3623209edddaa975cfeae5b0ef403fbce64d7b377ceb50396c74.
//
// Solidity: event TransferOwner(address indexed sender, address indexed oldOwner, address indexed newOwner)
func (_MultiSignature *MultiSignatureFilterer) ParseTransferOwner(log types.Log) (*MultiSignatureTransferOwner, error) {
	event := new(MultiSignatureTransferOwner)
	if err := _MultiSignature.contract.UnpackLog(event, "TransferOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
