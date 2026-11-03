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

// PledgePoolTokenMetaData contains all meta data concerning the PledgePoolToken contract.
var PledgePoolTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_swapRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_multiSignature\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"borrowFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkoutFinish\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkoutLiquidate\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkoutSettle\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimBorrow\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimLend\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createPoolInfo\",\"inputs\":[{\"name\":\"_settleTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_interestRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_martgageRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_lendToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_borrowToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_jpToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_autoLiquidateThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositBorrow\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositLend\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"emergencyBorrowWithdrawal\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emergencyLendWithdrawal\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finish\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMultiSignatureAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolState\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnderlyingPriceView\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lendFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lendScale\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lendingAsset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBscPledgeOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolBaseInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"settleTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interestRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lendSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"borrowSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"martgageRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lendToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumPledgePool.PoolState\"},{\"name\":\"spCoin\",\"type\":\"address\",\"internalType\":\"contractIDebtToken\"},{\"name\":\"jpCoin\",\"type\":\"address\",\"internalType\":\"contractIDebtToken\"},{\"name\":\"autoLiquidateThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolDataInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"settleAmountLend\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settleAmountBorrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finishAmountLend\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finishAmountBorrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidationAmounLend\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidationAmounBorrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundBorrow\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"refundLend\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFee\",\"inputs\":[{\"name\":\"_lendFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_borrowFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeAddress\",\"inputs\":[{\"name\":\"_feeAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinAmount\",\"inputs\":[{\"name\":\"_minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSwapRouterAddress\",\"inputs\":[{\"name\":\"_swapRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settle\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userBorrowInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"refundAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hasNoRefund\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"hasNoClaim\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userLendInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"refundAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hasNoRefund\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"hasNoClaim\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawBorrow\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_jpAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLend\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_spAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ClaimBorrow\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimLend\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositBorrow\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositLend\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyBorrowWithdrawal\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyLendWithdrawal\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Redeem\",\"inputs\":[{\"name\":\"recieptor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RefundBorrow\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"refund\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RefundLend\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"refund\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetFee\",\"inputs\":[{\"name\":\"newLendFee\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newBorrowFee\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetFeeAddress\",\"inputs\":[{\"name\":\"oldFeeAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newFeeAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMinAmount\",\"inputs\":[{\"name\":\"oldMinAmount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newMinAmount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetSwapRouterAddress\",\"inputs\":[{\"name\":\"oldSwapAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newSwapAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StateChange\",\"inputs\":[{\"name\":\"pid\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"beforeState\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"afterState\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"fromCoin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toCoin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fromValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawBorrow\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawLend\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// PledgePoolTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PledgePoolTokenMetaData.ABI instead.
var PledgePoolTokenABI = PledgePoolTokenMetaData.ABI

// PledgePoolToken is an auto generated Go binding around an Ethereum contract.
type PledgePoolToken struct {
	PledgePoolTokenCaller     // Read-only binding to the contract
	PledgePoolTokenTransactor // Write-only binding to the contract
	PledgePoolTokenFilterer   // Log filterer for contract events
}

// PledgePoolTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgePoolTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgePoolTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgePoolTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgePoolTokenSession struct {
	Contract     *PledgePoolToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgePoolTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgePoolTokenCallerSession struct {
	Contract *PledgePoolTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PledgePoolTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgePoolTokenTransactorSession struct {
	Contract     *PledgePoolTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PledgePoolTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgePoolTokenRaw struct {
	Contract *PledgePoolToken // Generic contract binding to access the raw methods on
}

// PledgePoolTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgePoolTokenCallerRaw struct {
	Contract *PledgePoolTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PledgePoolTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgePoolTokenTransactorRaw struct {
	Contract *PledgePoolTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledgePoolToken creates a new instance of PledgePoolToken, bound to a specific deployed contract.
func NewPledgePoolToken(address common.Address, backend bind.ContractBackend) (*PledgePoolToken, error) {
	contract, err := bindPledgePoolToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PledgePoolToken{PledgePoolTokenCaller: PledgePoolTokenCaller{contract: contract}, PledgePoolTokenTransactor: PledgePoolTokenTransactor{contract: contract}, PledgePoolTokenFilterer: PledgePoolTokenFilterer{contract: contract}}, nil
}

// NewPledgePoolTokenCaller creates a new read-only instance of PledgePoolToken, bound to a specific deployed contract.
func NewPledgePoolTokenCaller(address common.Address, caller bind.ContractCaller) (*PledgePoolTokenCaller, error) {
	contract, err := bindPledgePoolToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenCaller{contract: contract}, nil
}

// NewPledgePoolTokenTransactor creates a new write-only instance of PledgePoolToken, bound to a specific deployed contract.
func NewPledgePoolTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgePoolTokenTransactor, error) {
	contract, err := bindPledgePoolToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenTransactor{contract: contract}, nil
}

// NewPledgePoolTokenFilterer creates a new log filterer instance of PledgePoolToken, bound to a specific deployed contract.
func NewPledgePoolTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgePoolTokenFilterer, error) {
	contract, err := bindPledgePoolToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenFilterer{contract: contract}, nil
}

// bindPledgePoolToken binds a generic wrapper to an already deployed contract.
func bindPledgePoolToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PledgePoolTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgePoolToken *PledgePoolTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgePoolToken.Contract.PledgePoolTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgePoolToken *PledgePoolTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.PledgePoolTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgePoolToken *PledgePoolTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.PledgePoolTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgePoolToken *PledgePoolTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgePoolToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgePoolToken *PledgePoolTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgePoolToken *PledgePoolTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.contract.Transact(opts, method, params...)
}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCaller) BorrowFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "borrowFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenSession) BorrowFee() (*big.Int, error) {
	return _PledgePoolToken.Contract.BorrowFee(&_PledgePoolToken.CallOpts)
}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCallerSession) BorrowFee() (*big.Int, error) {
	return _PledgePoolToken.Contract.BorrowFee(&_PledgePoolToken.CallOpts)
}

// CheckoutFinish is a free data retrieval call binding the contract method 0x6abd7f29.
//
// Solidity: function checkoutFinish(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenCaller) CheckoutFinish(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "checkoutFinish", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckoutFinish is a free data retrieval call binding the contract method 0x6abd7f29.
//
// Solidity: function checkoutFinish(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenSession) CheckoutFinish(_pid *big.Int) (bool, error) {
	return _PledgePoolToken.Contract.CheckoutFinish(&_PledgePoolToken.CallOpts, _pid)
}

// CheckoutFinish is a free data retrieval call binding the contract method 0x6abd7f29.
//
// Solidity: function checkoutFinish(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenCallerSession) CheckoutFinish(_pid *big.Int) (bool, error) {
	return _PledgePoolToken.Contract.CheckoutFinish(&_PledgePoolToken.CallOpts, _pid)
}

// CheckoutLiquidate is a free data retrieval call binding the contract method 0x08e7305f.
//
// Solidity: function checkoutLiquidate(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenCaller) CheckoutLiquidate(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "checkoutLiquidate", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckoutLiquidate is a free data retrieval call binding the contract method 0x08e7305f.
//
// Solidity: function checkoutLiquidate(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenSession) CheckoutLiquidate(_pid *big.Int) (bool, error) {
	return _PledgePoolToken.Contract.CheckoutLiquidate(&_PledgePoolToken.CallOpts, _pid)
}

// CheckoutLiquidate is a free data retrieval call binding the contract method 0x08e7305f.
//
// Solidity: function checkoutLiquidate(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenCallerSession) CheckoutLiquidate(_pid *big.Int) (bool, error) {
	return _PledgePoolToken.Contract.CheckoutLiquidate(&_PledgePoolToken.CallOpts, _pid)
}

// CheckoutSettle is a free data retrieval call binding the contract method 0x14c090cc.
//
// Solidity: function checkoutSettle(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenCaller) CheckoutSettle(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "checkoutSettle", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckoutSettle is a free data retrieval call binding the contract method 0x14c090cc.
//
// Solidity: function checkoutSettle(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenSession) CheckoutSettle(_pid *big.Int) (bool, error) {
	return _PledgePoolToken.Contract.CheckoutSettle(&_PledgePoolToken.CallOpts, _pid)
}

// CheckoutSettle is a free data retrieval call binding the contract method 0x14c090cc.
//
// Solidity: function checkoutSettle(uint256 _pid) view returns(bool)
func (_PledgePoolToken *PledgePoolTokenCallerSession) CheckoutSettle(_pid *big.Int) (bool, error) {
	return _PledgePoolToken.Contract.CheckoutSettle(&_PledgePoolToken.CallOpts, _pid)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCaller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_PledgePoolToken *PledgePoolTokenSession) FeeAddress() (common.Address, error) {
	return _PledgePoolToken.Contract.FeeAddress(&_PledgePoolToken.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCallerSession) FeeAddress() (common.Address, error) {
	return _PledgePoolToken.Contract.FeeAddress(&_PledgePoolToken.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCaller) GetMultiSignatureAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "getMultiSignatureAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_PledgePoolToken *PledgePoolTokenSession) GetMultiSignatureAddress() (common.Address, error) {
	return _PledgePoolToken.Contract.GetMultiSignatureAddress(&_PledgePoolToken.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCallerSession) GetMultiSignatureAddress() (common.Address, error) {
	return _PledgePoolToken.Contract.GetMultiSignatureAddress(&_PledgePoolToken.CallOpts)
}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCaller) GetPoolState(opts *bind.CallOpts, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "getPoolState", _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenSession) GetPoolState(_pid *big.Int) (*big.Int, error) {
	return _PledgePoolToken.Contract.GetPoolState(&_PledgePoolToken.CallOpts, _pid)
}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCallerSession) GetPoolState(_pid *big.Int) (*big.Int, error) {
	return _PledgePoolToken.Contract.GetPoolState(&_PledgePoolToken.CallOpts, _pid)
}

// GetUnderlyingPriceView is a free data retrieval call binding the contract method 0xc9333756.
//
// Solidity: function getUnderlyingPriceView(uint256 _pid) view returns(uint256[2])
func (_PledgePoolToken *PledgePoolTokenCaller) GetUnderlyingPriceView(opts *bind.CallOpts, _pid *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "getUnderlyingPriceView", _pid)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetUnderlyingPriceView is a free data retrieval call binding the contract method 0xc9333756.
//
// Solidity: function getUnderlyingPriceView(uint256 _pid) view returns(uint256[2])
func (_PledgePoolToken *PledgePoolTokenSession) GetUnderlyingPriceView(_pid *big.Int) ([2]*big.Int, error) {
	return _PledgePoolToken.Contract.GetUnderlyingPriceView(&_PledgePoolToken.CallOpts, _pid)
}

// GetUnderlyingPriceView is a free data retrieval call binding the contract method 0xc9333756.
//
// Solidity: function getUnderlyingPriceView(uint256 _pid) view returns(uint256[2])
func (_PledgePoolToken *PledgePoolTokenCallerSession) GetUnderlyingPriceView(_pid *big.Int) ([2]*big.Int, error) {
	return _PledgePoolToken.Contract.GetUnderlyingPriceView(&_PledgePoolToken.CallOpts, _pid)
}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_PledgePoolToken *PledgePoolTokenCaller) GlobalPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "globalPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_PledgePoolToken *PledgePoolTokenSession) GlobalPaused() (bool, error) {
	return _PledgePoolToken.Contract.GlobalPaused(&_PledgePoolToken.CallOpts)
}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_PledgePoolToken *PledgePoolTokenCallerSession) GlobalPaused() (bool, error) {
	return _PledgePoolToken.Contract.GlobalPaused(&_PledgePoolToken.CallOpts)
}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCaller) LendFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "lendFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenSession) LendFee() (*big.Int, error) {
	return _PledgePoolToken.Contract.LendFee(&_PledgePoolToken.CallOpts)
}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCallerSession) LendFee() (*big.Int, error) {
	return _PledgePoolToken.Contract.LendFee(&_PledgePoolToken.CallOpts)
}

// LendScale is a free data retrieval call binding the contract method 0x8919a105.
//
// Solidity: function lendScale(uint256 ) view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCaller) LendScale(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "lendScale", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendScale is a free data retrieval call binding the contract method 0x8919a105.
//
// Solidity: function lendScale(uint256 ) view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenSession) LendScale(arg0 *big.Int) (*big.Int, error) {
	return _PledgePoolToken.Contract.LendScale(&_PledgePoolToken.CallOpts, arg0)
}

// LendScale is a free data retrieval call binding the contract method 0x8919a105.
//
// Solidity: function lendScale(uint256 ) view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCallerSession) LendScale(arg0 *big.Int) (*big.Int, error) {
	return _PledgePoolToken.Contract.LendScale(&_PledgePoolToken.CallOpts, arg0)
}

// LendingAsset is a free data retrieval call binding the contract method 0x1ffce3f4.
//
// Solidity: function lendingAsset() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCaller) LendingAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "lendingAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LendingAsset is a free data retrieval call binding the contract method 0x1ffce3f4.
//
// Solidity: function lendingAsset() view returns(address)
func (_PledgePoolToken *PledgePoolTokenSession) LendingAsset() (common.Address, error) {
	return _PledgePoolToken.Contract.LendingAsset(&_PledgePoolToken.CallOpts)
}

// LendingAsset is a free data retrieval call binding the contract method 0x1ffce3f4.
//
// Solidity: function lendingAsset() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCallerSession) LendingAsset() (common.Address, error) {
	return _PledgePoolToken.Contract.LendingAsset(&_PledgePoolToken.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenSession) MinAmount() (*big.Int, error) {
	return _PledgePoolToken.Contract.MinAmount(&_PledgePoolToken.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCallerSession) MinAmount() (*big.Int, error) {
	return _PledgePoolToken.Contract.MinAmount(&_PledgePoolToken.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PledgePoolToken *PledgePoolTokenSession) Oracle() (common.Address, error) {
	return _PledgePoolToken.Contract.Oracle(&_PledgePoolToken.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCallerSession) Oracle() (common.Address, error) {
	return _PledgePoolToken.Contract.Oracle(&_PledgePoolToken.CallOpts)
}

// PoolBaseInfo is a free data retrieval call binding the contract method 0x5a5a971e.
//
// Solidity: function poolBaseInfo(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 martgageRate, address lendToken, address borrowToken, uint8 state, address spCoin, address jpCoin, uint256 autoLiquidateThreshold)
func (_PledgePoolToken *PledgePoolTokenCaller) PoolBaseInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MartgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	SpCoin                 common.Address
	JpCoin                 common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "poolBaseInfo", arg0)

	outstruct := new(struct {
		SettleTime             *big.Int
		EndTime                *big.Int
		InterestRate           *big.Int
		MaxSupply              *big.Int
		LendSupply             *big.Int
		BorrowSupply           *big.Int
		MartgageRate           *big.Int
		LendToken              common.Address
		BorrowToken            common.Address
		State                  uint8
		SpCoin                 common.Address
		JpCoin                 common.Address
		AutoLiquidateThreshold *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SettleTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InterestRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LendSupply = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BorrowSupply = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MartgageRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LendToken = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.BorrowToken = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.State = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.SpCoin = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.JpCoin = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)
	outstruct.AutoLiquidateThreshold = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolBaseInfo is a free data retrieval call binding the contract method 0x5a5a971e.
//
// Solidity: function poolBaseInfo(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 martgageRate, address lendToken, address borrowToken, uint8 state, address spCoin, address jpCoin, uint256 autoLiquidateThreshold)
func (_PledgePoolToken *PledgePoolTokenSession) PoolBaseInfo(arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MartgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	SpCoin                 common.Address
	JpCoin                 common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	return _PledgePoolToken.Contract.PoolBaseInfo(&_PledgePoolToken.CallOpts, arg0)
}

// PoolBaseInfo is a free data retrieval call binding the contract method 0x5a5a971e.
//
// Solidity: function poolBaseInfo(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 martgageRate, address lendToken, address borrowToken, uint8 state, address spCoin, address jpCoin, uint256 autoLiquidateThreshold)
func (_PledgePoolToken *PledgePoolTokenCallerSession) PoolBaseInfo(arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MartgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	SpCoin                 common.Address
	JpCoin                 common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	return _PledgePoolToken.Contract.PoolBaseInfo(&_PledgePoolToken.CallOpts, arg0)
}

// PoolDataInfo is a free data retrieval call binding the contract method 0x0177b68c.
//
// Solidity: function poolDataInfo(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmounLend, uint256 liquidationAmounBorrow)
func (_PledgePoolToken *PledgePoolTokenCaller) PoolDataInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SettleAmountLend       *big.Int
	SettleAmountBorrow     *big.Int
	FinishAmountLend       *big.Int
	FinishAmountBorrow     *big.Int
	LiquidationAmounLend   *big.Int
	LiquidationAmounBorrow *big.Int
}, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "poolDataInfo", arg0)

	outstruct := new(struct {
		SettleAmountLend       *big.Int
		SettleAmountBorrow     *big.Int
		FinishAmountLend       *big.Int
		FinishAmountBorrow     *big.Int
		LiquidationAmounLend   *big.Int
		LiquidationAmounBorrow *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SettleAmountLend = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SettleAmountBorrow = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FinishAmountLend = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FinishAmountBorrow = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LiquidationAmounLend = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidationAmounBorrow = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolDataInfo is a free data retrieval call binding the contract method 0x0177b68c.
//
// Solidity: function poolDataInfo(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmounLend, uint256 liquidationAmounBorrow)
func (_PledgePoolToken *PledgePoolTokenSession) PoolDataInfo(arg0 *big.Int) (struct {
	SettleAmountLend       *big.Int
	SettleAmountBorrow     *big.Int
	FinishAmountLend       *big.Int
	FinishAmountBorrow     *big.Int
	LiquidationAmounLend   *big.Int
	LiquidationAmounBorrow *big.Int
}, error) {
	return _PledgePoolToken.Contract.PoolDataInfo(&_PledgePoolToken.CallOpts, arg0)
}

// PoolDataInfo is a free data retrieval call binding the contract method 0x0177b68c.
//
// Solidity: function poolDataInfo(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmounLend, uint256 liquidationAmounBorrow)
func (_PledgePoolToken *PledgePoolTokenCallerSession) PoolDataInfo(arg0 *big.Int) (struct {
	SettleAmountLend       *big.Int
	SettleAmountBorrow     *big.Int
	FinishAmountLend       *big.Int
	FinishAmountBorrow     *big.Int
	LiquidationAmounLend   *big.Int
	LiquidationAmounBorrow *big.Int
}, error) {
	return _PledgePoolToken.Contract.PoolDataInfo(&_PledgePoolToken.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenSession) PoolLength() (*big.Int, error) {
	return _PledgePoolToken.Contract.PoolLength(&_PledgePoolToken.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PledgePoolToken *PledgePoolTokenCallerSession) PoolLength() (*big.Int, error) {
	return _PledgePoolToken.Contract.PoolLength(&_PledgePoolToken.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_PledgePoolToken *PledgePoolTokenSession) SwapRouter() (common.Address, error) {
	return _PledgePoolToken.Contract.SwapRouter(&_PledgePoolToken.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_PledgePoolToken *PledgePoolTokenCallerSession) SwapRouter() (common.Address, error) {
	return _PledgePoolToken.Contract.SwapRouter(&_PledgePoolToken.CallOpts)
}

// UserBorrowInfo is a free data retrieval call binding the contract method 0x3c9fadc3.
//
// Solidity: function userBorrowInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePoolToken *PledgePoolTokenCaller) UserBorrowInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "userBorrowInfo", arg0, arg1)

	outstruct := new(struct {
		StakeAmount  *big.Int
		RefundAmount *big.Int
		HasNoRefund  bool
		HasNoClaim   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RefundAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HasNoRefund = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.HasNoClaim = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// UserBorrowInfo is a free data retrieval call binding the contract method 0x3c9fadc3.
//
// Solidity: function userBorrowInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePoolToken *PledgePoolTokenSession) UserBorrowInfo(arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	return _PledgePoolToken.Contract.UserBorrowInfo(&_PledgePoolToken.CallOpts, arg0, arg1)
}

// UserBorrowInfo is a free data retrieval call binding the contract method 0x3c9fadc3.
//
// Solidity: function userBorrowInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePoolToken *PledgePoolTokenCallerSession) UserBorrowInfo(arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	return _PledgePoolToken.Contract.UserBorrowInfo(&_PledgePoolToken.CallOpts, arg0, arg1)
}

// UserLendInfo is a free data retrieval call binding the contract method 0xbb176a64.
//
// Solidity: function userLendInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePoolToken *PledgePoolTokenCaller) UserLendInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	var out []interface{}
	err := _PledgePoolToken.contract.Call(opts, &out, "userLendInfo", arg0, arg1)

	outstruct := new(struct {
		StakeAmount  *big.Int
		RefundAmount *big.Int
		HasNoRefund  bool
		HasNoClaim   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RefundAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HasNoRefund = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.HasNoClaim = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// UserLendInfo is a free data retrieval call binding the contract method 0xbb176a64.
//
// Solidity: function userLendInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePoolToken *PledgePoolTokenSession) UserLendInfo(arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	return _PledgePoolToken.Contract.UserLendInfo(&_PledgePoolToken.CallOpts, arg0, arg1)
}

// UserLendInfo is a free data retrieval call binding the contract method 0xbb176a64.
//
// Solidity: function userLendInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePoolToken *PledgePoolTokenCallerSession) UserLendInfo(arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	return _PledgePoolToken.Contract.UserLendInfo(&_PledgePoolToken.CallOpts, arg0, arg1)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) ClaimBorrow(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "claimBorrow", _pid)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) ClaimBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.ClaimBorrow(&_PledgePoolToken.TransactOpts, _pid)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) ClaimBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.ClaimBorrow(&_PledgePoolToken.TransactOpts, _pid)
}

// ClaimLend is a paid mutator transaction binding the contract method 0x6c42fed2.
//
// Solidity: function claimLend(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) ClaimLend(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "claimLend", _pid)
}

// ClaimLend is a paid mutator transaction binding the contract method 0x6c42fed2.
//
// Solidity: function claimLend(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) ClaimLend(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.ClaimLend(&_PledgePoolToken.TransactOpts, _pid)
}

// ClaimLend is a paid mutator transaction binding the contract method 0x6c42fed2.
//
// Solidity: function claimLend(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) ClaimLend(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.ClaimLend(&_PledgePoolToken.TransactOpts, _pid)
}

// CreatePoolInfo is a paid mutator transaction binding the contract method 0xebded017.
//
// Solidity: function createPoolInfo(uint256 _settleTime, uint256 _endTime, uint64 _interestRate, uint256 _maxSupply, uint256 _martgageRate, address _lendToken, address _borrowToken, address _spToken, address _jpToken, uint256 _autoLiquidateThreshold) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) CreatePoolInfo(opts *bind.TransactOpts, _settleTime *big.Int, _endTime *big.Int, _interestRate uint64, _maxSupply *big.Int, _martgageRate *big.Int, _lendToken common.Address, _borrowToken common.Address, _spToken common.Address, _jpToken common.Address, _autoLiquidateThreshold *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "createPoolInfo", _settleTime, _endTime, _interestRate, _maxSupply, _martgageRate, _lendToken, _borrowToken, _spToken, _jpToken, _autoLiquidateThreshold)
}

// CreatePoolInfo is a paid mutator transaction binding the contract method 0xebded017.
//
// Solidity: function createPoolInfo(uint256 _settleTime, uint256 _endTime, uint64 _interestRate, uint256 _maxSupply, uint256 _martgageRate, address _lendToken, address _borrowToken, address _spToken, address _jpToken, uint256 _autoLiquidateThreshold) returns()
func (_PledgePoolToken *PledgePoolTokenSession) CreatePoolInfo(_settleTime *big.Int, _endTime *big.Int, _interestRate uint64, _maxSupply *big.Int, _martgageRate *big.Int, _lendToken common.Address, _borrowToken common.Address, _spToken common.Address, _jpToken common.Address, _autoLiquidateThreshold *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.CreatePoolInfo(&_PledgePoolToken.TransactOpts, _settleTime, _endTime, _interestRate, _maxSupply, _martgageRate, _lendToken, _borrowToken, _spToken, _jpToken, _autoLiquidateThreshold)
}

// CreatePoolInfo is a paid mutator transaction binding the contract method 0xebded017.
//
// Solidity: function createPoolInfo(uint256 _settleTime, uint256 _endTime, uint64 _interestRate, uint256 _maxSupply, uint256 _martgageRate, address _lendToken, address _borrowToken, address _spToken, address _jpToken, uint256 _autoLiquidateThreshold) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) CreatePoolInfo(_settleTime *big.Int, _endTime *big.Int, _interestRate uint64, _maxSupply *big.Int, _martgageRate *big.Int, _lendToken common.Address, _borrowToken common.Address, _spToken common.Address, _jpToken common.Address, _autoLiquidateThreshold *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.CreatePoolInfo(&_PledgePoolToken.TransactOpts, _settleTime, _endTime, _interestRate, _maxSupply, _martgageRate, _lendToken, _borrowToken, _spToken, _jpToken, _autoLiquidateThreshold)
}

// DepositBorrow is a paid mutator transaction binding the contract method 0x16f941b5.
//
// Solidity: function depositBorrow(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) DepositBorrow(opts *bind.TransactOpts, _pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "depositBorrow", _pid, _stakeAmount)
}

// DepositBorrow is a paid mutator transaction binding the contract method 0x16f941b5.
//
// Solidity: function depositBorrow(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_PledgePoolToken *PledgePoolTokenSession) DepositBorrow(_pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.DepositBorrow(&_PledgePoolToken.TransactOpts, _pid, _stakeAmount)
}

// DepositBorrow is a paid mutator transaction binding the contract method 0x16f941b5.
//
// Solidity: function depositBorrow(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) DepositBorrow(_pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.DepositBorrow(&_PledgePoolToken.TransactOpts, _pid, _stakeAmount)
}

// DepositLend is a paid mutator transaction binding the contract method 0x90590da0.
//
// Solidity: function depositLend(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) DepositLend(opts *bind.TransactOpts, _pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "depositLend", _pid, _stakeAmount)
}

// DepositLend is a paid mutator transaction binding the contract method 0x90590da0.
//
// Solidity: function depositLend(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_PledgePoolToken *PledgePoolTokenSession) DepositLend(_pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.DepositLend(&_PledgePoolToken.TransactOpts, _pid, _stakeAmount)
}

// DepositLend is a paid mutator transaction binding the contract method 0x90590da0.
//
// Solidity: function depositLend(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) DepositLend(_pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.DepositLend(&_PledgePoolToken.TransactOpts, _pid, _stakeAmount)
}

// EmergencyBorrowWithdrawal is a paid mutator transaction binding the contract method 0xe271fa0c.
//
// Solidity: function emergencyBorrowWithdrawal(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) EmergencyBorrowWithdrawal(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "emergencyBorrowWithdrawal", _pid)
}

// EmergencyBorrowWithdrawal is a paid mutator transaction binding the contract method 0xe271fa0c.
//
// Solidity: function emergencyBorrowWithdrawal(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) EmergencyBorrowWithdrawal(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.EmergencyBorrowWithdrawal(&_PledgePoolToken.TransactOpts, _pid)
}

// EmergencyBorrowWithdrawal is a paid mutator transaction binding the contract method 0xe271fa0c.
//
// Solidity: function emergencyBorrowWithdrawal(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) EmergencyBorrowWithdrawal(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.EmergencyBorrowWithdrawal(&_PledgePoolToken.TransactOpts, _pid)
}

// EmergencyLendWithdrawal is a paid mutator transaction binding the contract method 0xbf38b8f6.
//
// Solidity: function emergencyLendWithdrawal(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) EmergencyLendWithdrawal(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "emergencyLendWithdrawal", _pid)
}

// EmergencyLendWithdrawal is a paid mutator transaction binding the contract method 0xbf38b8f6.
//
// Solidity: function emergencyLendWithdrawal(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) EmergencyLendWithdrawal(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.EmergencyLendWithdrawal(&_PledgePoolToken.TransactOpts, _pid)
}

// EmergencyLendWithdrawal is a paid mutator transaction binding the contract method 0xbf38b8f6.
//
// Solidity: function emergencyLendWithdrawal(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) EmergencyLendWithdrawal(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.EmergencyLendWithdrawal(&_PledgePoolToken.TransactOpts, _pid)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) Finish(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "finish", _pid)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) Finish(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.Finish(&_PledgePoolToken.TransactOpts, _pid)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) Finish(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.Finish(&_PledgePoolToken.TransactOpts, _pid)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) Liquidate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "liquidate", _pid)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) Liquidate(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.Liquidate(&_PledgePoolToken.TransactOpts, _pid)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) Liquidate(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.Liquidate(&_PledgePoolToken.TransactOpts, _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) RefundBorrow(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "refundBorrow", _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) RefundBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.RefundBorrow(&_PledgePoolToken.TransactOpts, _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) RefundBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.RefundBorrow(&_PledgePoolToken.TransactOpts, _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) RefundLend(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "refundLend", _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) RefundLend(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.RefundLend(&_PledgePoolToken.TransactOpts, _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) RefundLend(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.RefundLend(&_PledgePoolToken.TransactOpts, _pid)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) SetFee(opts *bind.TransactOpts, _lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "setFee", _lendFee, _borrowFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_PledgePoolToken *PledgePoolTokenSession) SetFee(_lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetFee(&_PledgePoolToken.TransactOpts, _lendFee, _borrowFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) SetFee(_lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetFee(&_PledgePoolToken.TransactOpts, _lendFee, _borrowFee)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) SetFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "setFeeAddress", _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_PledgePoolToken *PledgePoolTokenSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetFeeAddress(&_PledgePoolToken.TransactOpts, _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetFeeAddress(&_PledgePoolToken.TransactOpts, _feeAddress)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) SetMinAmount(opts *bind.TransactOpts, _minAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "setMinAmount", _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_PledgePoolToken *PledgePoolTokenSession) SetMinAmount(_minAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetMinAmount(&_PledgePoolToken.TransactOpts, _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) SetMinAmount(_minAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetMinAmount(&_PledgePoolToken.TransactOpts, _minAmount)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) SetPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "setPause")
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_PledgePoolToken *PledgePoolTokenSession) SetPause() (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetPause(&_PledgePoolToken.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) SetPause() (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetPause(&_PledgePoolToken.TransactOpts)
}

// SetSwapRouterAddress is a paid mutator transaction binding the contract method 0x5249961b.
//
// Solidity: function setSwapRouterAddress(address _swapRouter) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) SetSwapRouterAddress(opts *bind.TransactOpts, _swapRouter common.Address) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "setSwapRouterAddress", _swapRouter)
}

// SetSwapRouterAddress is a paid mutator transaction binding the contract method 0x5249961b.
//
// Solidity: function setSwapRouterAddress(address _swapRouter) returns()
func (_PledgePoolToken *PledgePoolTokenSession) SetSwapRouterAddress(_swapRouter common.Address) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetSwapRouterAddress(&_PledgePoolToken.TransactOpts, _swapRouter)
}

// SetSwapRouterAddress is a paid mutator transaction binding the contract method 0x5249961b.
//
// Solidity: function setSwapRouterAddress(address _swapRouter) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) SetSwapRouterAddress(_swapRouter common.Address) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.SetSwapRouterAddress(&_PledgePoolToken.TransactOpts, _swapRouter)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) Settle(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "settle", _pid)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenSession) Settle(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.Settle(&_PledgePoolToken.TransactOpts, _pid)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _pid) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) Settle(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.Settle(&_PledgePoolToken.TransactOpts, _pid)
}

// WithdrawBorrow is a paid mutator transaction binding the contract method 0x1e107979.
//
// Solidity: function withdrawBorrow(uint256 _pid, uint256 _jpAmount) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) WithdrawBorrow(opts *bind.TransactOpts, _pid *big.Int, _jpAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "withdrawBorrow", _pid, _jpAmount)
}

// WithdrawBorrow is a paid mutator transaction binding the contract method 0x1e107979.
//
// Solidity: function withdrawBorrow(uint256 _pid, uint256 _jpAmount) returns()
func (_PledgePoolToken *PledgePoolTokenSession) WithdrawBorrow(_pid *big.Int, _jpAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.WithdrawBorrow(&_PledgePoolToken.TransactOpts, _pid, _jpAmount)
}

// WithdrawBorrow is a paid mutator transaction binding the contract method 0x1e107979.
//
// Solidity: function withdrawBorrow(uint256 _pid, uint256 _jpAmount) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) WithdrawBorrow(_pid *big.Int, _jpAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.WithdrawBorrow(&_PledgePoolToken.TransactOpts, _pid, _jpAmount)
}

// WithdrawLend is a paid mutator transaction binding the contract method 0x38f2aa76.
//
// Solidity: function withdrawLend(uint256 _pid, uint256 _spAmount) returns()
func (_PledgePoolToken *PledgePoolTokenTransactor) WithdrawLend(opts *bind.TransactOpts, _pid *big.Int, _spAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.contract.Transact(opts, "withdrawLend", _pid, _spAmount)
}

// WithdrawLend is a paid mutator transaction binding the contract method 0x38f2aa76.
//
// Solidity: function withdrawLend(uint256 _pid, uint256 _spAmount) returns()
func (_PledgePoolToken *PledgePoolTokenSession) WithdrawLend(_pid *big.Int, _spAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.WithdrawLend(&_PledgePoolToken.TransactOpts, _pid, _spAmount)
}

// WithdrawLend is a paid mutator transaction binding the contract method 0x38f2aa76.
//
// Solidity: function withdrawLend(uint256 _pid, uint256 _spAmount) returns()
func (_PledgePoolToken *PledgePoolTokenTransactorSession) WithdrawLend(_pid *big.Int, _spAmount *big.Int) (*types.Transaction, error) {
	return _PledgePoolToken.Contract.WithdrawLend(&_PledgePoolToken.TransactOpts, _pid, _spAmount)
}

// PledgePoolTokenClaimBorrowIterator is returned from FilterClaimBorrow and is used to iterate over the raw logs and unpacked data for ClaimBorrow events raised by the PledgePoolToken contract.
type PledgePoolTokenClaimBorrowIterator struct {
	Event *PledgePoolTokenClaimBorrow // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenClaimBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenClaimBorrow)
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
		it.Event = new(PledgePoolTokenClaimBorrow)
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
func (it *PledgePoolTokenClaimBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenClaimBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenClaimBorrow represents a ClaimBorrow event raised by the PledgePoolToken contract.
type PledgePoolTokenClaimBorrow struct {
	From   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimBorrow is a free log retrieval operation binding the contract event 0x3ddafe3ebb4d0c818317027aabfa82dc9983942ceeb80523167e2de047b17fbd.
//
// Solidity: event ClaimBorrow(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterClaimBorrow(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenClaimBorrowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "ClaimBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenClaimBorrowIterator{contract: _PledgePoolToken.contract, event: "ClaimBorrow", logs: logs, sub: sub}, nil
}

// WatchClaimBorrow is a free log subscription operation binding the contract event 0x3ddafe3ebb4d0c818317027aabfa82dc9983942ceeb80523167e2de047b17fbd.
//
// Solidity: event ClaimBorrow(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchClaimBorrow(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenClaimBorrow, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "ClaimBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenClaimBorrow)
				if err := _PledgePoolToken.contract.UnpackLog(event, "ClaimBorrow", log); err != nil {
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

// ParseClaimBorrow is a log parse operation binding the contract event 0x3ddafe3ebb4d0c818317027aabfa82dc9983942ceeb80523167e2de047b17fbd.
//
// Solidity: event ClaimBorrow(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseClaimBorrow(log types.Log) (*PledgePoolTokenClaimBorrow, error) {
	event := new(PledgePoolTokenClaimBorrow)
	if err := _PledgePoolToken.contract.UnpackLog(event, "ClaimBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenClaimLendIterator is returned from FilterClaimLend and is used to iterate over the raw logs and unpacked data for ClaimLend events raised by the PledgePoolToken contract.
type PledgePoolTokenClaimLendIterator struct {
	Event *PledgePoolTokenClaimLend // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenClaimLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenClaimLend)
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
		it.Event = new(PledgePoolTokenClaimLend)
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
func (it *PledgePoolTokenClaimLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenClaimLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenClaimLend represents a ClaimLend event raised by the PledgePoolToken contract.
type PledgePoolTokenClaimLend struct {
	From   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimLend is a free log retrieval operation binding the contract event 0x6f4dd2687b3c3bfa99d39742b01d6e0ad9604c48559791d5df4ff5df44b41dfd.
//
// Solidity: event ClaimLend(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterClaimLend(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenClaimLendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "ClaimLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenClaimLendIterator{contract: _PledgePoolToken.contract, event: "ClaimLend", logs: logs, sub: sub}, nil
}

// WatchClaimLend is a free log subscription operation binding the contract event 0x6f4dd2687b3c3bfa99d39742b01d6e0ad9604c48559791d5df4ff5df44b41dfd.
//
// Solidity: event ClaimLend(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchClaimLend(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenClaimLend, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "ClaimLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenClaimLend)
				if err := _PledgePoolToken.contract.UnpackLog(event, "ClaimLend", log); err != nil {
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

// ParseClaimLend is a log parse operation binding the contract event 0x6f4dd2687b3c3bfa99d39742b01d6e0ad9604c48559791d5df4ff5df44b41dfd.
//
// Solidity: event ClaimLend(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseClaimLend(log types.Log) (*PledgePoolTokenClaimLend, error) {
	event := new(PledgePoolTokenClaimLend)
	if err := _PledgePoolToken.contract.UnpackLog(event, "ClaimLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenDepositBorrowIterator is returned from FilterDepositBorrow and is used to iterate over the raw logs and unpacked data for DepositBorrow events raised by the PledgePoolToken contract.
type PledgePoolTokenDepositBorrowIterator struct {
	Event *PledgePoolTokenDepositBorrow // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenDepositBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenDepositBorrow)
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
		it.Event = new(PledgePoolTokenDepositBorrow)
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
func (it *PledgePoolTokenDepositBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenDepositBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenDepositBorrow represents a DepositBorrow event raised by the PledgePoolToken contract.
type PledgePoolTokenDepositBorrow struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositBorrow is a free log retrieval operation binding the contract event 0x1d7b72e666a0b6217efe7cfa1b604ea5c7b39219563ce48b30c9da77045247a5.
//
// Solidity: event DepositBorrow(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterDepositBorrow(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenDepositBorrowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "DepositBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenDepositBorrowIterator{contract: _PledgePoolToken.contract, event: "DepositBorrow", logs: logs, sub: sub}, nil
}

// WatchDepositBorrow is a free log subscription operation binding the contract event 0x1d7b72e666a0b6217efe7cfa1b604ea5c7b39219563ce48b30c9da77045247a5.
//
// Solidity: event DepositBorrow(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchDepositBorrow(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenDepositBorrow, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "DepositBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenDepositBorrow)
				if err := _PledgePoolToken.contract.UnpackLog(event, "DepositBorrow", log); err != nil {
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

// ParseDepositBorrow is a log parse operation binding the contract event 0x1d7b72e666a0b6217efe7cfa1b604ea5c7b39219563ce48b30c9da77045247a5.
//
// Solidity: event DepositBorrow(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseDepositBorrow(log types.Log) (*PledgePoolTokenDepositBorrow, error) {
	event := new(PledgePoolTokenDepositBorrow)
	if err := _PledgePoolToken.contract.UnpackLog(event, "DepositBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenDepositLendIterator is returned from FilterDepositLend and is used to iterate over the raw logs and unpacked data for DepositLend events raised by the PledgePoolToken contract.
type PledgePoolTokenDepositLendIterator struct {
	Event *PledgePoolTokenDepositLend // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenDepositLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenDepositLend)
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
		it.Event = new(PledgePoolTokenDepositLend)
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
func (it *PledgePoolTokenDepositLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenDepositLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenDepositLend represents a DepositLend event raised by the PledgePoolToken contract.
type PledgePoolTokenDepositLend struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositLend is a free log retrieval operation binding the contract event 0x129e8c18c2f7baf99c7eb257934c21f038c72412803512dcf0a942a4562a82ea.
//
// Solidity: event DepositLend(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterDepositLend(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenDepositLendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "DepositLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenDepositLendIterator{contract: _PledgePoolToken.contract, event: "DepositLend", logs: logs, sub: sub}, nil
}

// WatchDepositLend is a free log subscription operation binding the contract event 0x129e8c18c2f7baf99c7eb257934c21f038c72412803512dcf0a942a4562a82ea.
//
// Solidity: event DepositLend(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchDepositLend(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenDepositLend, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "DepositLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenDepositLend)
				if err := _PledgePoolToken.contract.UnpackLog(event, "DepositLend", log); err != nil {
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

// ParseDepositLend is a log parse operation binding the contract event 0x129e8c18c2f7baf99c7eb257934c21f038c72412803512dcf0a942a4562a82ea.
//
// Solidity: event DepositLend(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseDepositLend(log types.Log) (*PledgePoolTokenDepositLend, error) {
	event := new(PledgePoolTokenDepositLend)
	if err := _PledgePoolToken.contract.UnpackLog(event, "DepositLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenEmergencyBorrowWithdrawalIterator is returned from FilterEmergencyBorrowWithdrawal and is used to iterate over the raw logs and unpacked data for EmergencyBorrowWithdrawal events raised by the PledgePoolToken contract.
type PledgePoolTokenEmergencyBorrowWithdrawalIterator struct {
	Event *PledgePoolTokenEmergencyBorrowWithdrawal // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenEmergencyBorrowWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenEmergencyBorrowWithdrawal)
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
		it.Event = new(PledgePoolTokenEmergencyBorrowWithdrawal)
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
func (it *PledgePoolTokenEmergencyBorrowWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenEmergencyBorrowWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenEmergencyBorrowWithdrawal represents a EmergencyBorrowWithdrawal event raised by the PledgePoolToken contract.
type PledgePoolTokenEmergencyBorrowWithdrawal struct {
	From   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyBorrowWithdrawal is a free log retrieval operation binding the contract event 0x5a06c7de92f1dc59e8cba872927d016c80ce5f0fb2295c898dfb7a2f08e43fb1.
//
// Solidity: event EmergencyBorrowWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterEmergencyBorrowWithdrawal(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenEmergencyBorrowWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "EmergencyBorrowWithdrawal", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenEmergencyBorrowWithdrawalIterator{contract: _PledgePoolToken.contract, event: "EmergencyBorrowWithdrawal", logs: logs, sub: sub}, nil
}

// WatchEmergencyBorrowWithdrawal is a free log subscription operation binding the contract event 0x5a06c7de92f1dc59e8cba872927d016c80ce5f0fb2295c898dfb7a2f08e43fb1.
//
// Solidity: event EmergencyBorrowWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchEmergencyBorrowWithdrawal(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenEmergencyBorrowWithdrawal, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "EmergencyBorrowWithdrawal", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenEmergencyBorrowWithdrawal)
				if err := _PledgePoolToken.contract.UnpackLog(event, "EmergencyBorrowWithdrawal", log); err != nil {
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

// ParseEmergencyBorrowWithdrawal is a log parse operation binding the contract event 0x5a06c7de92f1dc59e8cba872927d016c80ce5f0fb2295c898dfb7a2f08e43fb1.
//
// Solidity: event EmergencyBorrowWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseEmergencyBorrowWithdrawal(log types.Log) (*PledgePoolTokenEmergencyBorrowWithdrawal, error) {
	event := new(PledgePoolTokenEmergencyBorrowWithdrawal)
	if err := _PledgePoolToken.contract.UnpackLog(event, "EmergencyBorrowWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenEmergencyLendWithdrawalIterator is returned from FilterEmergencyLendWithdrawal and is used to iterate over the raw logs and unpacked data for EmergencyLendWithdrawal events raised by the PledgePoolToken contract.
type PledgePoolTokenEmergencyLendWithdrawalIterator struct {
	Event *PledgePoolTokenEmergencyLendWithdrawal // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenEmergencyLendWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenEmergencyLendWithdrawal)
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
		it.Event = new(PledgePoolTokenEmergencyLendWithdrawal)
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
func (it *PledgePoolTokenEmergencyLendWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenEmergencyLendWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenEmergencyLendWithdrawal represents a EmergencyLendWithdrawal event raised by the PledgePoolToken contract.
type PledgePoolTokenEmergencyLendWithdrawal struct {
	From   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyLendWithdrawal is a free log retrieval operation binding the contract event 0x71d14c5f08cb34cbfb59c06ea5151aafbf742d0b6ed00fdb83addd9afb5c0fd0.
//
// Solidity: event EmergencyLendWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterEmergencyLendWithdrawal(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenEmergencyLendWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "EmergencyLendWithdrawal", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenEmergencyLendWithdrawalIterator{contract: _PledgePoolToken.contract, event: "EmergencyLendWithdrawal", logs: logs, sub: sub}, nil
}

// WatchEmergencyLendWithdrawal is a free log subscription operation binding the contract event 0x71d14c5f08cb34cbfb59c06ea5151aafbf742d0b6ed00fdb83addd9afb5c0fd0.
//
// Solidity: event EmergencyLendWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchEmergencyLendWithdrawal(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenEmergencyLendWithdrawal, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "EmergencyLendWithdrawal", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenEmergencyLendWithdrawal)
				if err := _PledgePoolToken.contract.UnpackLog(event, "EmergencyLendWithdrawal", log); err != nil {
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

// ParseEmergencyLendWithdrawal is a log parse operation binding the contract event 0x71d14c5f08cb34cbfb59c06ea5151aafbf742d0b6ed00fdb83addd9afb5c0fd0.
//
// Solidity: event EmergencyLendWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseEmergencyLendWithdrawal(log types.Log) (*PledgePoolTokenEmergencyLendWithdrawal, error) {
	event := new(PledgePoolTokenEmergencyLendWithdrawal)
	if err := _PledgePoolToken.contract.UnpackLog(event, "EmergencyLendWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the PledgePoolToken contract.
type PledgePoolTokenRedeemIterator struct {
	Event *PledgePoolTokenRedeem // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenRedeem)
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
		it.Event = new(PledgePoolTokenRedeem)
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
func (it *PledgePoolTokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenRedeem represents a Redeem event raised by the PledgePoolToken contract.
type PledgePoolTokenRedeem struct {
	Recieptor common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed recieptor, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterRedeem(opts *bind.FilterOpts, recieptor []common.Address, token []common.Address) (*PledgePoolTokenRedeemIterator, error) {

	var recieptorRule []interface{}
	for _, recieptorItem := range recieptor {
		recieptorRule = append(recieptorRule, recieptorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "Redeem", recieptorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenRedeemIterator{contract: _PledgePoolToken.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed recieptor, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenRedeem, recieptor []common.Address, token []common.Address) (event.Subscription, error) {

	var recieptorRule []interface{}
	for _, recieptorItem := range recieptor {
		recieptorRule = append(recieptorRule, recieptorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "Redeem", recieptorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenRedeem)
				if err := _PledgePoolToken.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed recieptor, address indexed token, uint256 amount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseRedeem(log types.Log) (*PledgePoolTokenRedeem, error) {
	event := new(PledgePoolTokenRedeem)
	if err := _PledgePoolToken.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenRefundBorrowIterator is returned from FilterRefundBorrow and is used to iterate over the raw logs and unpacked data for RefundBorrow events raised by the PledgePoolToken contract.
type PledgePoolTokenRefundBorrowIterator struct {
	Event *PledgePoolTokenRefundBorrow // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenRefundBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenRefundBorrow)
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
		it.Event = new(PledgePoolTokenRefundBorrow)
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
func (it *PledgePoolTokenRefundBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenRefundBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenRefundBorrow represents a RefundBorrow event raised by the PledgePoolToken contract.
type PledgePoolTokenRefundBorrow struct {
	From   common.Address
	Token  common.Address
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundBorrow is a free log retrieval operation binding the contract event 0x732816f48de550f238bd0d4f5b79819c7b24a49d6132928978e3cd36568dd5db.
//
// Solidity: event RefundBorrow(address indexed from, address indexed token, uint256 refund)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterRefundBorrow(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenRefundBorrowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "RefundBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenRefundBorrowIterator{contract: _PledgePoolToken.contract, event: "RefundBorrow", logs: logs, sub: sub}, nil
}

// WatchRefundBorrow is a free log subscription operation binding the contract event 0x732816f48de550f238bd0d4f5b79819c7b24a49d6132928978e3cd36568dd5db.
//
// Solidity: event RefundBorrow(address indexed from, address indexed token, uint256 refund)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchRefundBorrow(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenRefundBorrow, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "RefundBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenRefundBorrow)
				if err := _PledgePoolToken.contract.UnpackLog(event, "RefundBorrow", log); err != nil {
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

// ParseRefundBorrow is a log parse operation binding the contract event 0x732816f48de550f238bd0d4f5b79819c7b24a49d6132928978e3cd36568dd5db.
//
// Solidity: event RefundBorrow(address indexed from, address indexed token, uint256 refund)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseRefundBorrow(log types.Log) (*PledgePoolTokenRefundBorrow, error) {
	event := new(PledgePoolTokenRefundBorrow)
	if err := _PledgePoolToken.contract.UnpackLog(event, "RefundBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenRefundLendIterator is returned from FilterRefundLend and is used to iterate over the raw logs and unpacked data for RefundLend events raised by the PledgePoolToken contract.
type PledgePoolTokenRefundLendIterator struct {
	Event *PledgePoolTokenRefundLend // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenRefundLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenRefundLend)
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
		it.Event = new(PledgePoolTokenRefundLend)
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
func (it *PledgePoolTokenRefundLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenRefundLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenRefundLend represents a RefundLend event raised by the PledgePoolToken contract.
type PledgePoolTokenRefundLend struct {
	From   common.Address
	Token  common.Address
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundLend is a free log retrieval operation binding the contract event 0xc3e20279d41b3ed21d277920877e5e5c6665bf6aca607046a3fe0fd2bd6bda7d.
//
// Solidity: event RefundLend(address indexed from, address indexed token, uint256 refund)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterRefundLend(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenRefundLendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "RefundLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenRefundLendIterator{contract: _PledgePoolToken.contract, event: "RefundLend", logs: logs, sub: sub}, nil
}

// WatchRefundLend is a free log subscription operation binding the contract event 0xc3e20279d41b3ed21d277920877e5e5c6665bf6aca607046a3fe0fd2bd6bda7d.
//
// Solidity: event RefundLend(address indexed from, address indexed token, uint256 refund)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchRefundLend(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenRefundLend, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "RefundLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenRefundLend)
				if err := _PledgePoolToken.contract.UnpackLog(event, "RefundLend", log); err != nil {
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

// ParseRefundLend is a log parse operation binding the contract event 0xc3e20279d41b3ed21d277920877e5e5c6665bf6aca607046a3fe0fd2bd6bda7d.
//
// Solidity: event RefundLend(address indexed from, address indexed token, uint256 refund)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseRefundLend(log types.Log) (*PledgePoolTokenRefundLend, error) {
	event := new(PledgePoolTokenRefundLend)
	if err := _PledgePoolToken.contract.UnpackLog(event, "RefundLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the PledgePoolToken contract.
type PledgePoolTokenSetFeeIterator struct {
	Event *PledgePoolTokenSetFee // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenSetFee)
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
		it.Event = new(PledgePoolTokenSetFee)
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
func (it *PledgePoolTokenSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenSetFee represents a SetFee event raised by the PledgePoolToken contract.
type PledgePoolTokenSetFee struct {
	NewLendFee   *big.Int
	NewBorrowFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterSetFee(opts *bind.FilterOpts, newLendFee []*big.Int, newBorrowFee []*big.Int) (*PledgePoolTokenSetFeeIterator, error) {

	var newLendFeeRule []interface{}
	for _, newLendFeeItem := range newLendFee {
		newLendFeeRule = append(newLendFeeRule, newLendFeeItem)
	}
	var newBorrowFeeRule []interface{}
	for _, newBorrowFeeItem := range newBorrowFee {
		newBorrowFeeRule = append(newBorrowFeeRule, newBorrowFeeItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "SetFee", newLendFeeRule, newBorrowFeeRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenSetFeeIterator{contract: _PledgePoolToken.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenSetFee, newLendFee []*big.Int, newBorrowFee []*big.Int) (event.Subscription, error) {

	var newLendFeeRule []interface{}
	for _, newLendFeeItem := range newLendFee {
		newLendFeeRule = append(newLendFeeRule, newLendFeeItem)
	}
	var newBorrowFeeRule []interface{}
	for _, newBorrowFeeItem := range newBorrowFee {
		newBorrowFeeRule = append(newBorrowFeeRule, newBorrowFeeItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "SetFee", newLendFeeRule, newBorrowFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenSetFee)
				if err := _PledgePoolToken.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseSetFee(log types.Log) (*PledgePoolTokenSetFee, error) {
	event := new(PledgePoolTokenSetFee)
	if err := _PledgePoolToken.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenSetFeeAddressIterator is returned from FilterSetFeeAddress and is used to iterate over the raw logs and unpacked data for SetFeeAddress events raised by the PledgePoolToken contract.
type PledgePoolTokenSetFeeAddressIterator struct {
	Event *PledgePoolTokenSetFeeAddress // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenSetFeeAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenSetFeeAddress)
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
		it.Event = new(PledgePoolTokenSetFeeAddress)
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
func (it *PledgePoolTokenSetFeeAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenSetFeeAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenSetFeeAddress represents a SetFeeAddress event raised by the PledgePoolToken contract.
type PledgePoolTokenSetFeeAddress struct {
	OldFeeAddress common.Address
	NewFeeAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetFeeAddress is a free log retrieval operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterSetFeeAddress(opts *bind.FilterOpts, oldFeeAddress []common.Address, newFeeAddress []common.Address) (*PledgePoolTokenSetFeeAddressIterator, error) {

	var oldFeeAddressRule []interface{}
	for _, oldFeeAddressItem := range oldFeeAddress {
		oldFeeAddressRule = append(oldFeeAddressRule, oldFeeAddressItem)
	}
	var newFeeAddressRule []interface{}
	for _, newFeeAddressItem := range newFeeAddress {
		newFeeAddressRule = append(newFeeAddressRule, newFeeAddressItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "SetFeeAddress", oldFeeAddressRule, newFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenSetFeeAddressIterator{contract: _PledgePoolToken.contract, event: "SetFeeAddress", logs: logs, sub: sub}, nil
}

// WatchSetFeeAddress is a free log subscription operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchSetFeeAddress(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenSetFeeAddress, oldFeeAddress []common.Address, newFeeAddress []common.Address) (event.Subscription, error) {

	var oldFeeAddressRule []interface{}
	for _, oldFeeAddressItem := range oldFeeAddress {
		oldFeeAddressRule = append(oldFeeAddressRule, oldFeeAddressItem)
	}
	var newFeeAddressRule []interface{}
	for _, newFeeAddressItem := range newFeeAddress {
		newFeeAddressRule = append(newFeeAddressRule, newFeeAddressItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "SetFeeAddress", oldFeeAddressRule, newFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenSetFeeAddress)
				if err := _PledgePoolToken.contract.UnpackLog(event, "SetFeeAddress", log); err != nil {
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

// ParseSetFeeAddress is a log parse operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseSetFeeAddress(log types.Log) (*PledgePoolTokenSetFeeAddress, error) {
	event := new(PledgePoolTokenSetFeeAddress)
	if err := _PledgePoolToken.contract.UnpackLog(event, "SetFeeAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenSetMinAmountIterator is returned from FilterSetMinAmount and is used to iterate over the raw logs and unpacked data for SetMinAmount events raised by the PledgePoolToken contract.
type PledgePoolTokenSetMinAmountIterator struct {
	Event *PledgePoolTokenSetMinAmount // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenSetMinAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenSetMinAmount)
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
		it.Event = new(PledgePoolTokenSetMinAmount)
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
func (it *PledgePoolTokenSetMinAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenSetMinAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenSetMinAmount represents a SetMinAmount event raised by the PledgePoolToken contract.
type PledgePoolTokenSetMinAmount struct {
	OldMinAmount *big.Int
	NewMinAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetMinAmount is a free log retrieval operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterSetMinAmount(opts *bind.FilterOpts, oldMinAmount []*big.Int, newMinAmount []*big.Int) (*PledgePoolTokenSetMinAmountIterator, error) {

	var oldMinAmountRule []interface{}
	for _, oldMinAmountItem := range oldMinAmount {
		oldMinAmountRule = append(oldMinAmountRule, oldMinAmountItem)
	}
	var newMinAmountRule []interface{}
	for _, newMinAmountItem := range newMinAmount {
		newMinAmountRule = append(newMinAmountRule, newMinAmountItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "SetMinAmount", oldMinAmountRule, newMinAmountRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenSetMinAmountIterator{contract: _PledgePoolToken.contract, event: "SetMinAmount", logs: logs, sub: sub}, nil
}

// WatchSetMinAmount is a free log subscription operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchSetMinAmount(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenSetMinAmount, oldMinAmount []*big.Int, newMinAmount []*big.Int) (event.Subscription, error) {

	var oldMinAmountRule []interface{}
	for _, oldMinAmountItem := range oldMinAmount {
		oldMinAmountRule = append(oldMinAmountRule, oldMinAmountItem)
	}
	var newMinAmountRule []interface{}
	for _, newMinAmountItem := range newMinAmount {
		newMinAmountRule = append(newMinAmountRule, newMinAmountItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "SetMinAmount", oldMinAmountRule, newMinAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenSetMinAmount)
				if err := _PledgePoolToken.contract.UnpackLog(event, "SetMinAmount", log); err != nil {
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

// ParseSetMinAmount is a log parse operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseSetMinAmount(log types.Log) (*PledgePoolTokenSetMinAmount, error) {
	event := new(PledgePoolTokenSetMinAmount)
	if err := _PledgePoolToken.contract.UnpackLog(event, "SetMinAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenSetSwapRouterAddressIterator is returned from FilterSetSwapRouterAddress and is used to iterate over the raw logs and unpacked data for SetSwapRouterAddress events raised by the PledgePoolToken contract.
type PledgePoolTokenSetSwapRouterAddressIterator struct {
	Event *PledgePoolTokenSetSwapRouterAddress // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenSetSwapRouterAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenSetSwapRouterAddress)
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
		it.Event = new(PledgePoolTokenSetSwapRouterAddress)
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
func (it *PledgePoolTokenSetSwapRouterAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenSetSwapRouterAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenSetSwapRouterAddress represents a SetSwapRouterAddress event raised by the PledgePoolToken contract.
type PledgePoolTokenSetSwapRouterAddress struct {
	OldSwapAddress common.Address
	NewSwapAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetSwapRouterAddress is a free log retrieval operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterSetSwapRouterAddress(opts *bind.FilterOpts, oldSwapAddress []common.Address, newSwapAddress []common.Address) (*PledgePoolTokenSetSwapRouterAddressIterator, error) {

	var oldSwapAddressRule []interface{}
	for _, oldSwapAddressItem := range oldSwapAddress {
		oldSwapAddressRule = append(oldSwapAddressRule, oldSwapAddressItem)
	}
	var newSwapAddressRule []interface{}
	for _, newSwapAddressItem := range newSwapAddress {
		newSwapAddressRule = append(newSwapAddressRule, newSwapAddressItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "SetSwapRouterAddress", oldSwapAddressRule, newSwapAddressRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenSetSwapRouterAddressIterator{contract: _PledgePoolToken.contract, event: "SetSwapRouterAddress", logs: logs, sub: sub}, nil
}

// WatchSetSwapRouterAddress is a free log subscription operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchSetSwapRouterAddress(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenSetSwapRouterAddress, oldSwapAddress []common.Address, newSwapAddress []common.Address) (event.Subscription, error) {

	var oldSwapAddressRule []interface{}
	for _, oldSwapAddressItem := range oldSwapAddress {
		oldSwapAddressRule = append(oldSwapAddressRule, oldSwapAddressItem)
	}
	var newSwapAddressRule []interface{}
	for _, newSwapAddressItem := range newSwapAddress {
		newSwapAddressRule = append(newSwapAddressRule, newSwapAddressItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "SetSwapRouterAddress", oldSwapAddressRule, newSwapAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenSetSwapRouterAddress)
				if err := _PledgePoolToken.contract.UnpackLog(event, "SetSwapRouterAddress", log); err != nil {
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

// ParseSetSwapRouterAddress is a log parse operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseSetSwapRouterAddress(log types.Log) (*PledgePoolTokenSetSwapRouterAddress, error) {
	event := new(PledgePoolTokenSetSwapRouterAddress)
	if err := _PledgePoolToken.contract.UnpackLog(event, "SetSwapRouterAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenStateChangeIterator is returned from FilterStateChange and is used to iterate over the raw logs and unpacked data for StateChange events raised by the PledgePoolToken contract.
type PledgePoolTokenStateChangeIterator struct {
	Event *PledgePoolTokenStateChange // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenStateChange)
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
		it.Event = new(PledgePoolTokenStateChange)
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
func (it *PledgePoolTokenStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenStateChange represents a StateChange event raised by the PledgePoolToken contract.
type PledgePoolTokenStateChange struct {
	Pid         *big.Int
	BeforeState *big.Int
	AfterState  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStateChange is a free log retrieval operation binding the contract event 0x516112f3bf06e373fcea44db364769c04cc7ef4392e6de95d2b250720bcacefb.
//
// Solidity: event StateChange(uint256 indexed pid, uint256 indexed beforeState, uint256 indexed afterState)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterStateChange(opts *bind.FilterOpts, pid []*big.Int, beforeState []*big.Int, afterState []*big.Int) (*PledgePoolTokenStateChangeIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var beforeStateRule []interface{}
	for _, beforeStateItem := range beforeState {
		beforeStateRule = append(beforeStateRule, beforeStateItem)
	}
	var afterStateRule []interface{}
	for _, afterStateItem := range afterState {
		afterStateRule = append(afterStateRule, afterStateItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "StateChange", pidRule, beforeStateRule, afterStateRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenStateChangeIterator{contract: _PledgePoolToken.contract, event: "StateChange", logs: logs, sub: sub}, nil
}

// WatchStateChange is a free log subscription operation binding the contract event 0x516112f3bf06e373fcea44db364769c04cc7ef4392e6de95d2b250720bcacefb.
//
// Solidity: event StateChange(uint256 indexed pid, uint256 indexed beforeState, uint256 indexed afterState)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchStateChange(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenStateChange, pid []*big.Int, beforeState []*big.Int, afterState []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var beforeStateRule []interface{}
	for _, beforeStateItem := range beforeState {
		beforeStateRule = append(beforeStateRule, beforeStateItem)
	}
	var afterStateRule []interface{}
	for _, afterStateItem := range afterState {
		afterStateRule = append(afterStateRule, afterStateItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "StateChange", pidRule, beforeStateRule, afterStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenStateChange)
				if err := _PledgePoolToken.contract.UnpackLog(event, "StateChange", log); err != nil {
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

// ParseStateChange is a log parse operation binding the contract event 0x516112f3bf06e373fcea44db364769c04cc7ef4392e6de95d2b250720bcacefb.
//
// Solidity: event StateChange(uint256 indexed pid, uint256 indexed beforeState, uint256 indexed afterState)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseStateChange(log types.Log) (*PledgePoolTokenStateChange, error) {
	event := new(PledgePoolTokenStateChange)
	if err := _PledgePoolToken.contract.UnpackLog(event, "StateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the PledgePoolToken contract.
type PledgePoolTokenSwapIterator struct {
	Event *PledgePoolTokenSwap // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenSwap)
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
		it.Event = new(PledgePoolTokenSwap)
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
func (it *PledgePoolTokenSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenSwap represents a Swap event raised by the PledgePoolToken contract.
type PledgePoolTokenSwap struct {
	FromCoin  common.Address
	ToCoin    common.Address
	FromValue *big.Int
	ToValue   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xfa2dda1cc1b86e41239702756b13effbc1a092b5c57e3ad320fbe4f3b13fe235.
//
// Solidity: event Swap(address indexed fromCoin, address indexed toCoin, uint256 fromValue, uint256 toValue)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterSwap(opts *bind.FilterOpts, fromCoin []common.Address, toCoin []common.Address) (*PledgePoolTokenSwapIterator, error) {

	var fromCoinRule []interface{}
	for _, fromCoinItem := range fromCoin {
		fromCoinRule = append(fromCoinRule, fromCoinItem)
	}
	var toCoinRule []interface{}
	for _, toCoinItem := range toCoin {
		toCoinRule = append(toCoinRule, toCoinItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "Swap", fromCoinRule, toCoinRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenSwapIterator{contract: _PledgePoolToken.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xfa2dda1cc1b86e41239702756b13effbc1a092b5c57e3ad320fbe4f3b13fe235.
//
// Solidity: event Swap(address indexed fromCoin, address indexed toCoin, uint256 fromValue, uint256 toValue)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenSwap, fromCoin []common.Address, toCoin []common.Address) (event.Subscription, error) {

	var fromCoinRule []interface{}
	for _, fromCoinItem := range fromCoin {
		fromCoinRule = append(fromCoinRule, fromCoinItem)
	}
	var toCoinRule []interface{}
	for _, toCoinItem := range toCoin {
		toCoinRule = append(toCoinRule, toCoinItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "Swap", fromCoinRule, toCoinRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenSwap)
				if err := _PledgePoolToken.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xfa2dda1cc1b86e41239702756b13effbc1a092b5c57e3ad320fbe4f3b13fe235.
//
// Solidity: event Swap(address indexed fromCoin, address indexed toCoin, uint256 fromValue, uint256 toValue)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseSwap(log types.Log) (*PledgePoolTokenSwap, error) {
	event := new(PledgePoolTokenSwap)
	if err := _PledgePoolToken.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenWithdrawBorrowIterator is returned from FilterWithdrawBorrow and is used to iterate over the raw logs and unpacked data for WithdrawBorrow events raised by the PledgePoolToken contract.
type PledgePoolTokenWithdrawBorrowIterator struct {
	Event *PledgePoolTokenWithdrawBorrow // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenWithdrawBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenWithdrawBorrow)
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
		it.Event = new(PledgePoolTokenWithdrawBorrow)
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
func (it *PledgePoolTokenWithdrawBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenWithdrawBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenWithdrawBorrow represents a WithdrawBorrow event raised by the PledgePoolToken contract.
type PledgePoolTokenWithdrawBorrow struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	BurnAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBorrow is a free log retrieval operation binding the contract event 0x0f5e74952c2f9259a748f3aa9a6c4534a6f46a5966e5baabdb6bd337f05234a8.
//
// Solidity: event WithdrawBorrow(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterWithdrawBorrow(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenWithdrawBorrowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "WithdrawBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenWithdrawBorrowIterator{contract: _PledgePoolToken.contract, event: "WithdrawBorrow", logs: logs, sub: sub}, nil
}

// WatchWithdrawBorrow is a free log subscription operation binding the contract event 0x0f5e74952c2f9259a748f3aa9a6c4534a6f46a5966e5baabdb6bd337f05234a8.
//
// Solidity: event WithdrawBorrow(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchWithdrawBorrow(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenWithdrawBorrow, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "WithdrawBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenWithdrawBorrow)
				if err := _PledgePoolToken.contract.UnpackLog(event, "WithdrawBorrow", log); err != nil {
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

// ParseWithdrawBorrow is a log parse operation binding the contract event 0x0f5e74952c2f9259a748f3aa9a6c4534a6f46a5966e5baabdb6bd337f05234a8.
//
// Solidity: event WithdrawBorrow(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseWithdrawBorrow(log types.Log) (*PledgePoolTokenWithdrawBorrow, error) {
	event := new(PledgePoolTokenWithdrawBorrow)
	if err := _PledgePoolToken.contract.UnpackLog(event, "WithdrawBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolTokenWithdrawLendIterator is returned from FilterWithdrawLend and is used to iterate over the raw logs and unpacked data for WithdrawLend events raised by the PledgePoolToken contract.
type PledgePoolTokenWithdrawLendIterator struct {
	Event *PledgePoolTokenWithdrawLend // Event containing the contract specifics and raw log

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
func (it *PledgePoolTokenWithdrawLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolTokenWithdrawLend)
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
		it.Event = new(PledgePoolTokenWithdrawLend)
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
func (it *PledgePoolTokenWithdrawLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolTokenWithdrawLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolTokenWithdrawLend represents a WithdrawLend event raised by the PledgePoolToken contract.
type PledgePoolTokenWithdrawLend struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	BurnAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawLend is a free log retrieval operation binding the contract event 0x690f32ccf3e832d5ff975d781039bc2affebee9c973939c9b710091b87954c9d.
//
// Solidity: event WithdrawLend(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) FilterWithdrawLend(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgePoolTokenWithdrawLendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.FilterLogs(opts, "WithdrawLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTokenWithdrawLendIterator{contract: _PledgePoolToken.contract, event: "WithdrawLend", logs: logs, sub: sub}, nil
}

// WatchWithdrawLend is a free log subscription operation binding the contract event 0x690f32ccf3e832d5ff975d781039bc2affebee9c973939c9b710091b87954c9d.
//
// Solidity: event WithdrawLend(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) WatchWithdrawLend(opts *bind.WatchOpts, sink chan<- *PledgePoolTokenWithdrawLend, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PledgePoolToken.contract.WatchLogs(opts, "WithdrawLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolTokenWithdrawLend)
				if err := _PledgePoolToken.contract.UnpackLog(event, "WithdrawLend", log); err != nil {
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

// ParseWithdrawLend is a log parse operation binding the contract event 0x690f32ccf3e832d5ff975d781039bc2affebee9c973939c9b710091b87954c9d.
//
// Solidity: event WithdrawLend(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_PledgePoolToken *PledgePoolTokenFilterer) ParseWithdrawLend(log types.Log) (*PledgePoolTokenWithdrawLend, error) {
	event := new(PledgePoolTokenWithdrawLend)
	if err := _PledgePoolToken.contract.UnpackLog(event, "WithdrawLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
