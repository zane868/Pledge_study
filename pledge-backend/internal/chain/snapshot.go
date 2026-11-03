package chain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"pledge-backend/config"
	"pledge-backend/contract/bindings"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Token struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
	Price    string `json:"priceRaw,omitempty"`
}

type Pool struct {
	PoolIndex               uint64 `json:"poolIndex"`
	DisplayPoolID           uint64 `json:"displayPoolId"`
	SettleTime              string `json:"settleTime"`
	EndTime                 string `json:"endTime"`
	InterestRate            string `json:"interestRate"`
	MaxSupply               string `json:"maxSupply"`
	LendSupply              string `json:"lendSupply"`
	BorrowSupply            string `json:"borrowSupply"`
	MartgageRate            string `json:"martgageRate"`
	LendToken               string `json:"lendToken"`
	BorrowToken             string `json:"borrowToken"`
	State                   uint8  `json:"state"`
	SP                      string `json:"sp"`
	JP                      string `json:"jp"`
	AutoLiquidateThreshold  string `json:"autoLiquidateThreshold"`
	SettleAmountLend        string `json:"settleAmountLend"`
	SettleAmountBorrow      string `json:"settleAmountBorrow"`
	FinishAmountLend        string `json:"finishAmountLend"`
	FinishAmountBorrow      string `json:"finishAmountBorrow"`
	LiquidationAmountLend   string `json:"liquidationAmountLend"`
	LiquidationAmountBorrow string `json:"liquidationAmountBorrow"`
}

type Snapshot struct {
	ChainID      string           `json:"chainId"`
	PoolContract string           `json:"poolContract"`
	BlockNumber  string           `json:"blockNumber"`
	BlockHash    string           `json:"blockHash"`
	LendFee      string           `json:"lendFee"`
	BorrowFee    string           `json:"borrowFee"`
	Pools        []Pool           `json:"pools"`
	Tokens       map[string]Token `json:"tokens"`
}

type Reader interface {
	Read(context.Context) (*Snapshot, error)
}

type Client struct {
	URL           string
	Deployment    config.Deployment
	Confirmations uint64
}

func (c Client) Read(ctx context.Context) (snapshot *Snapshot, err error) {
	defer func() {
		if err != nil && c.URL != "" {
			err = errors.New(strings.ReplaceAll(err.Error(), c.URL, "[RPC]"))
		}
	}()
	rpc, err := ethclient.DialContext(ctx, c.URL)
	if err != nil {
		return nil, fmt.Errorf("connect RPC failed")
	}
	defer rpc.Close()
	actual, err := rpc.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("read chain ID failed")
	}
	if !actual.IsUint64() || actual.Uint64() != c.Deployment.ChainID {
		return nil, fmt.Errorf("RPC chain ID does not match deployment")
	}
	head, err := rpc.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("read latest block failed")
	}
	if !head.Number.IsUint64() || head.Number.Uint64() < c.Confirmations {
		return nil, fmt.Errorf("insufficient confirmed blocks")
	}
	number := new(big.Int).SetUint64(head.Number.Uint64() - c.Confirmations)
	if number.Uint64() < c.Deployment.DeploymentStartBlock {
		return nil, fmt.Errorf("deployment is not confirmed yet")
	}
	header, err := rpc.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("read confirmed block failed")
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: number}
	contract, err := bindings.NewPledgePoolToken(common.HexToAddress(c.Deployment.Pool), rpc)
	if err != nil {
		return nil, err
	}
	oracle, err := bindings.NewBscPledgeOracleTestnetToken(common.HexToAddress(c.Deployment.Oracle), rpc)
	if err != nil {
		return nil, err
	}
	lendFee, err := contract.LendFee(opts)
	if err != nil {
		return nil, fmt.Errorf("read lend fee: %w", err)
	}
	borrowFee, err := contract.BorrowFee(opts)
	if err != nil {
		return nil, fmt.Errorf("read borrow fee: %w", err)
	}
	length, err := contract.PoolLength(opts)
	if err != nil {
		return nil, fmt.Errorf("read pool length: %w", err)
	}
	if !length.IsUint64() || length.Uint64() > 10000 {
		return nil, fmt.Errorf("pool count exceeds supported snapshot size")
	}
	result := &Snapshot{ChainID: actual.String(), PoolContract: common.HexToAddress(c.Deployment.Pool).Hex(), BlockNumber: number.String(), BlockHash: header.Hash().Hex(), LendFee: lendFee.String(), BorrowFee: borrowFee.String(), Pools: []Pool{}, Tokens: map[string]Token{}}
	token := func(address common.Address, priced bool) error {
		key := address.Hex()
		if _, ok := result.Tokens[key]; ok {
			return nil
		}
		value := Token{Address: key, Symbol: "ETH", Decimals: 18}
		if address != (common.Address{}) {
			erc, err := bindings.NewDebtToken(address, rpc)
			if err != nil {
				return err
			}
			value.Symbol, err = erc.Symbol(opts)
			if err != nil {
				return fmt.Errorf("read token symbol: %w", err)
			}
			value.Decimals, err = erc.Decimals(opts)
			if err != nil {
				return fmt.Errorf("read token decimals: %w", err)
			}
		}
		if address == common.HexToAddress(c.Deployment.Stable) && c.Deployment.StableDecimals != 0 && value.Decimals != c.Deployment.StableDecimals {
			return fmt.Errorf("stablecoin decimals do not match deployment")
		}
		if priced {
			price, err := oracle.GetPrice(opts, address)
			if err != nil {
				return fmt.Errorf("read oracle price: %w", err)
			}
			if price.Sign() <= 0 {
				return fmt.Errorf("oracle returned a nonpositive price for %s", key)
			}
			value.Price = price.String()
		}
		result.Tokens[key] = value
		return nil
	}
	for i := uint64(0); i < length.Uint64(); i++ {
		index := new(big.Int).SetUint64(i)
		base, err := contract.PoolBaseInfo(opts, index)
		if err != nil {
			return nil, fmt.Errorf("read pool %d base: %w", i, err)
		}
		data, err := contract.PoolDataInfo(opts, index)
		if err != nil {
			return nil, fmt.Errorf("read pool %d data: %w", i, err)
		}
		for _, a := range []common.Address{base.LendToken, base.BorrowToken} {
			if err = token(a, true); err != nil {
				return nil, err
			}
		}
		for _, a := range []common.Address{base.SpCoin, base.JpCoin} {
			if err = token(a, false); err != nil {
				return nil, err
			}
		}
		result.Pools = append(result.Pools, Pool{PoolIndex: i, DisplayPoolID: i + 1, SettleTime: base.SettleTime.String(), EndTime: base.EndTime.String(), InterestRate: base.InterestRate.String(), MaxSupply: base.MaxSupply.String(), LendSupply: base.LendSupply.String(), BorrowSupply: base.BorrowSupply.String(), MartgageRate: base.MartgageRate.String(), LendToken: base.LendToken.Hex(), BorrowToken: base.BorrowToken.Hex(), State: base.State, SP: base.SpCoin.Hex(), JP: base.JpCoin.Hex(), AutoLiquidateThreshold: base.AutoLiquidateThreshold.String(), SettleAmountLend: data.SettleAmountLend.String(), SettleAmountBorrow: data.SettleAmountBorrow.String(), FinishAmountLend: data.FinishAmountLend.String(), FinishAmountBorrow: data.FinishAmountBorrow.String(), LiquidationAmountLend: data.LiquidationAmounLend.String(), LiquidationAmountBorrow: data.LiquidationAmounBorrow.String()})
	}
	// Reject a reorg during the read; do not publish a mixed or orphaned snapshot.
	end, err := rpc.HeaderByNumber(ctx, number)
	if err != nil || end.Hash() != header.Hash() {
		return nil, fmt.Errorf("snapshot block changed while reading")
	}
	return result, nil
}
