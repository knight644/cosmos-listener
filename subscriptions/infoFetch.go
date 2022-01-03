package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cosmos-listener/configs"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type TokenSupply struct {
	Amount TokenAmount `json:"amount"`
}

type TokenAmount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type StakingPool struct {
	Pool BondedPool `json:"pool"`
}

type BondedPool struct {
	NotBondedTokens string `json:"not_bonded_tokens"`
	BondedTokens    string `json:"bonded_tokens"`
}

func FetchTokenSupply(denom string) uint64 {
	uri := configs.HttpAddr()
	path := "/cosmos/bank/v1beta1/supply/" + denom
	url := uri + path

	rsp, err := http.Get(url)
	var supply TokenSupply
	if err == nil {
		json.NewDecoder(rsp.Body).Decode(&supply)
	}

	val, _ := strconv.ParseUint(supply.Amount.Amount, 10, 64)
	return val
}

func FetchStakingPool() uint64 {
	uri := configs.HttpAddr()
	path := "/cosmos/staking/v1beta1/pool"
	url := uri + path

	rsp, err := http.Get(url)
	var pool StakingPool
	if err == nil {
		json.NewDecoder(rsp.Body).Decode(&pool)
	}

	val, _ := strconv.ParseUint(pool.Pool.BondedTokens, 10, 64)
	return val
}

func GetTransactionDetails(txhash []byte) (*ctypes.ResultTx, error) {
	// Create blank context (Top level)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	txResult, err := Client.Tx(ctx, txhash, false)
	// log.Println(txResult, err)
	return txResult, err
}

func GetPrevBlockDetails(height int64) (*ctypes.ResultBlock, error) {
	height = height - 1

	// Create blank context (Top level)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	blockDetails, err := Client.Block(ctx, &height)
	return blockDetails, err
}
