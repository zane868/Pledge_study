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

// BscPledgeOracleMainnetTokenMetaData contains all meta data concerning the BscPledgeOracleMainnetToken contract.
var BscPledgeOracleMainnetTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"multiSignature\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAssetsAggregator\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiSignatureAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrices\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnderlyingAggregator\",\"inputs\":[{\"name\":\"underlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnderlyingPrice\",\"inputs\":[{\"name\":\"underlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAssetsAggregator\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"aggergator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDecimals\",\"inputs\":[{\"name\":\"newDecimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPrice\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPrices\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"prices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnderlyingAggregator\",\"inputs\":[{\"name\":\"underlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggergator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnderlyingPrice\",\"inputs\":[{\"name\":\"underlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// BscPledgeOracleMainnetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BscPledgeOracleMainnetTokenMetaData.ABI instead.
var BscPledgeOracleMainnetTokenABI = BscPledgeOracleMainnetTokenMetaData.ABI

// BscPledgeOracleMainnetToken is an auto generated Go binding around an Ethereum contract.
type BscPledgeOracleMainnetToken struct {
	BscPledgeOracleMainnetTokenCaller     // Read-only binding to the contract
	BscPledgeOracleMainnetTokenTransactor // Write-only binding to the contract
	BscPledgeOracleMainnetTokenFilterer   // Log filterer for contract events
}

// BscPledgeOracleMainnetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscPledgeOracleMainnetTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleMainnetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscPledgeOracleMainnetTokenSession struct {
	Contract     *BscPledgeOracleMainnetToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BscPledgeOracleMainnetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscPledgeOracleMainnetTokenCallerSession struct {
	Contract *BscPledgeOracleMainnetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// BscPledgeOracleMainnetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscPledgeOracleMainnetTokenTransactorSession struct {
	Contract     *BscPledgeOracleMainnetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// BscPledgeOracleMainnetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenRaw struct {
	Contract *BscPledgeOracleMainnetToken // Generic contract binding to access the raw methods on
}

// BscPledgeOracleMainnetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenCallerRaw struct {
	Contract *BscPledgeOracleMainnetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BscPledgeOracleMainnetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscPledgeOracleMainnetTokenTransactorRaw struct {
	Contract *BscPledgeOracleMainnetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscPledgeOracleMainnetToken creates a new instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetToken(address common.Address, backend bind.ContractBackend) (*BscPledgeOracleMainnetToken, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetToken{BscPledgeOracleMainnetTokenCaller: BscPledgeOracleMainnetTokenCaller{contract: contract}, BscPledgeOracleMainnetTokenTransactor: BscPledgeOracleMainnetTokenTransactor{contract: contract}, BscPledgeOracleMainnetTokenFilterer: BscPledgeOracleMainnetTokenFilterer{contract: contract}}, nil
}

// NewBscPledgeOracleMainnetTokenCaller creates a new read-only instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenCaller(address common.Address, caller bind.ContractCaller) (*BscPledgeOracleMainnetTokenCaller, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenCaller{contract: contract}, nil
}

// NewBscPledgeOracleMainnetTokenTransactor creates a new write-only instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BscPledgeOracleMainnetTokenTransactor, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenTransactor{contract: contract}, nil
}

// NewBscPledgeOracleMainnetTokenFilterer creates a new log filterer instance of BscPledgeOracleMainnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleMainnetTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BscPledgeOracleMainnetTokenFilterer, error) {
	contract, err := bindBscPledgeOracleMainnetToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleMainnetTokenFilterer{contract: contract}, nil
}

// bindBscPledgeOracleMainnetToken binds a generic wrapper to an already deployed contract.
func bindBscPledgeOracleMainnetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BscPledgeOracleMainnetTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.BscPledgeOracleMainnetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleMainnetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.contract.Transact(opts, method, params...)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetAssetsAggregator(opts *bind.CallOpts, asset common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getAssetsAggregator", asset)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetMultiSignatureAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getMultiSignatureAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetMultiSignatureAddress() (common.Address, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetMultiSignatureAddress(&_BscPledgeOracleMainnetToken.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetMultiSignatureAddress() (common.Address, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetMultiSignatureAddress(&_BscPledgeOracleMainnetToken.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrice(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrice(&_BscPledgeOracleMainnetToken.CallOpts, asset)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetPrices(opts *bind.CallOpts, assets []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrices(&_BscPledgeOracleMainnetToken.CallOpts, assets)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetPrices(&_BscPledgeOracleMainnetToken.CallOpts, assets)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetUnderlyingAggregator(opts *bind.CallOpts, underlying *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getUnderlyingAggregator", underlying)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCaller) GetUnderlyingPrice(opts *bind.CallOpts, underlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleMainnetToken.contract.Call(opts, &out, "getUnderlyingPrice", underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenCallerSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleMainnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleMainnetToken.CallOpts, underlying)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetAssetsAggregator(opts *bind.TransactOpts, asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setAssetsAggregator", asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetDecimals(opts *bind.TransactOpts, newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setDecimals", newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetDecimals(&_BscPledgeOracleMainnetToken.TransactOpts, newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetDecimals(&_BscPledgeOracleMainnetToken.TransactOpts, newDecimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setPrice", asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrice(&_BscPledgeOracleMainnetToken.TransactOpts, asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrice(&_BscPledgeOracleMainnetToken.TransactOpts, asset, price)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetPrices(opts *bind.TransactOpts, assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setPrices", assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrices(&_BscPledgeOracleMainnetToken.TransactOpts, assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetPrices(&_BscPledgeOracleMainnetToken.TransactOpts, assets, prices)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetUnderlyingAggregator(opts *bind.TransactOpts, underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setUnderlyingAggregator", underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.contract.Transact(opts, "setUnderlyingPrice", underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleMainnetToken *BscPledgeOracleMainnetTokenTransactorSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleMainnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleMainnetToken.TransactOpts, underlying, price)
}
