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

// BscPledgeOracleTestnetTokenMetaData contains all meta data concerning the BscPledgeOracleTestnetToken contract.
var BscPledgeOracleTestnetTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"multiSignature\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAssetsAggregator\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiSignatureAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrices\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnderlyingAggregator\",\"inputs\":[{\"name\":\"underlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnderlyingPrice\",\"inputs\":[{\"name\":\"underlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAssetsAggregator\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"aggergator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDecimals\",\"inputs\":[{\"name\":\"newDecimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPrice\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPrices\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"prices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnderlyingAggregator\",\"inputs\":[{\"name\":\"underlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggergator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnderlyingPrice\",\"inputs\":[{\"name\":\"underlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// BscPledgeOracleTestnetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BscPledgeOracleTestnetTokenMetaData.ABI instead.
var BscPledgeOracleTestnetTokenABI = BscPledgeOracleTestnetTokenMetaData.ABI

// BscPledgeOracleTestnetToken is an auto generated Go binding around an Ethereum contract.
type BscPledgeOracleTestnetToken struct {
	BscPledgeOracleTestnetTokenCaller     // Read-only binding to the contract
	BscPledgeOracleTestnetTokenTransactor // Write-only binding to the contract
	BscPledgeOracleTestnetTokenFilterer   // Log filterer for contract events
}

// BscPledgeOracleTestnetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleTestnetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleTestnetTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscPledgeOracleTestnetTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleTestnetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscPledgeOracleTestnetTokenSession struct {
	Contract     *BscPledgeOracleTestnetToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BscPledgeOracleTestnetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscPledgeOracleTestnetTokenCallerSession struct {
	Contract *BscPledgeOracleTestnetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// BscPledgeOracleTestnetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscPledgeOracleTestnetTokenTransactorSession struct {
	Contract     *BscPledgeOracleTestnetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// BscPledgeOracleTestnetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenRaw struct {
	Contract *BscPledgeOracleTestnetToken // Generic contract binding to access the raw methods on
}

// BscPledgeOracleTestnetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenCallerRaw struct {
	Contract *BscPledgeOracleTestnetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BscPledgeOracleTestnetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscPledgeOracleTestnetTokenTransactorRaw struct {
	Contract *BscPledgeOracleTestnetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscPledgeOracleTestnetToken creates a new instance of BscPledgeOracleTestnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleTestnetToken(address common.Address, backend bind.ContractBackend) (*BscPledgeOracleTestnetToken, error) {
	contract, err := bindBscPledgeOracleTestnetToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetToken{BscPledgeOracleTestnetTokenCaller: BscPledgeOracleTestnetTokenCaller{contract: contract}, BscPledgeOracleTestnetTokenTransactor: BscPledgeOracleTestnetTokenTransactor{contract: contract}, BscPledgeOracleTestnetTokenFilterer: BscPledgeOracleTestnetTokenFilterer{contract: contract}}, nil
}

// NewBscPledgeOracleTestnetTokenCaller creates a new read-only instance of BscPledgeOracleTestnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleTestnetTokenCaller(address common.Address, caller bind.ContractCaller) (*BscPledgeOracleTestnetTokenCaller, error) {
	contract, err := bindBscPledgeOracleTestnetToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetTokenCaller{contract: contract}, nil
}

// NewBscPledgeOracleTestnetTokenTransactor creates a new write-only instance of BscPledgeOracleTestnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleTestnetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BscPledgeOracleTestnetTokenTransactor, error) {
	contract, err := bindBscPledgeOracleTestnetToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetTokenTransactor{contract: contract}, nil
}

// NewBscPledgeOracleTestnetTokenFilterer creates a new log filterer instance of BscPledgeOracleTestnetToken, bound to a specific deployed contract.
func NewBscPledgeOracleTestnetTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BscPledgeOracleTestnetTokenFilterer, error) {
	contract, err := bindBscPledgeOracleTestnetToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTestnetTokenFilterer{contract: contract}, nil
}

// bindBscPledgeOracleTestnetToken binds a generic wrapper to an already deployed contract.
func bindBscPledgeOracleTestnetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BscPledgeOracleTestnetTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleTestnetToken.Contract.BscPledgeOracleTestnetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.BscPledgeOracleTestnetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.BscPledgeOracleTestnetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracleTestnetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.contract.Transact(opts, method, params...)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetAssetsAggregator(opts *bind.CallOpts, asset common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getAssetsAggregator", asset)

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
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleTestnetToken.CallOpts, asset)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetAssetsAggregator(&_BscPledgeOracleTestnetToken.CallOpts, asset)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetMultiSignatureAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getMultiSignatureAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetMultiSignatureAddress() (common.Address, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetMultiSignatureAddress(&_BscPledgeOracleTestnetToken.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetMultiSignatureAddress() (common.Address, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetMultiSignatureAddress(&_BscPledgeOracleTestnetToken.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetPrice(&_BscPledgeOracleTestnetToken.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetPrice(&_BscPledgeOracleTestnetToken.CallOpts, asset)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetPrices(opts *bind.CallOpts, assets []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetPrices(&_BscPledgeOracleTestnetToken.CallOpts, assets)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetPrices(&_BscPledgeOracleTestnetToken.CallOpts, assets)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetUnderlyingAggregator(opts *bind.CallOpts, underlying *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getUnderlyingAggregator", underlying)

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
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleTestnetToken.CallOpts, underlying)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetUnderlyingAggregator(&_BscPledgeOracleTestnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCaller) GetUnderlyingPrice(opts *bind.CallOpts, underlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracleTestnetToken.contract.Call(opts, &out, "getUnderlyingPrice", underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleTestnetToken.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenCallerSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracleTestnetToken.Contract.GetUnderlyingPrice(&_BscPledgeOracleTestnetToken.CallOpts, underlying)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetAssetsAggregator(opts *bind.TransactOpts, asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setAssetsAggregator", asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleTestnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xcd9ffa0b.
//
// Solidity: function setAssetsAggregator(address asset, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetAssetsAggregator(asset common.Address, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetAssetsAggregator(&_BscPledgeOracleTestnetToken.TransactOpts, asset, aggergator, _decimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetDecimals(opts *bind.TransactOpts, newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setDecimals", newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetDecimals(&_BscPledgeOracleTestnetToken.TransactOpts, newDecimals)
}

// SetDecimals is a paid mutator transaction binding the contract method 0x8c8885c8.
//
// Solidity: function setDecimals(uint256 newDecimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetDecimals(newDecimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetDecimals(&_BscPledgeOracleTestnetToken.TransactOpts, newDecimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setPrice", asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetPrice(&_BscPledgeOracleTestnetToken.TransactOpts, asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetPrice(&_BscPledgeOracleTestnetToken.TransactOpts, asset, price)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetPrices(opts *bind.TransactOpts, assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setPrices", assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetPrices(&_BscPledgeOracleTestnetToken.TransactOpts, assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetPrices(&_BscPledgeOracleTestnetToken.TransactOpts, assets, prices)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetUnderlyingAggregator(opts *bind.TransactOpts, underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setUnderlyingAggregator", underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleTestnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x434b33ac.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggergator, uint256 _decimals) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetUnderlyingAggregator(underlying *big.Int, aggergator common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetUnderlyingAggregator(&_BscPledgeOracleTestnetToken.TransactOpts, underlying, aggergator, _decimals)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.contract.Transact(opts, "setUnderlyingPrice", underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleTestnetToken.TransactOpts, underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracleTestnetToken *BscPledgeOracleTestnetTokenTransactorSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracleTestnetToken.Contract.SetUnderlyingPrice(&_BscPledgeOracleTestnetToken.TransactOpts, underlying, price)
}
