package quoter

import (
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kochetkov-av/hcni1/contract/pair"
	"go.uber.org/zap"
)

type Quoter struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Quoter {
	return &Quoter{
		logger: logger,
	}
}

func (q *Quoter) Quote(ethereumURL string, poolID, fromToken, toToken common.Address, amountIn *big.Int) (*big.Int, error) {
	q.logger.Info("dialing ethereum endpoint")
	client, err := ethclient.Dial(ethereumURL)
	if err != nil {
		return nil, err
	}

	pool, err := pair.NewPair(poolID, client)
	if err != nil {
		log.Fatal(err)
	}

	q.logger.Info("fetching token0")
	t0, err := pool.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Token0: %w", err)
	}

	q.logger.Info("fetching token1")
	t1, err := pool.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Token1: %w", err)
	}

	if !(t0 == fromToken && t1 == toToken) && !(t1 == fromToken && t0 == toToken) {
		return nil, errors.New("provided tokens doesn't match pool")
	}

	q.logger.Info("fetching reserves")
	reserves, err := pool.GetReserves(nil)
	if err != nil {
		return nil, errors.New("failed to fetch reserves")
	}

	q.logger.Info("calculating quotes")
	var reserve0, reserve1 *big.Int
	if t0 == fromToken {
		reserve0, reserve1 = reserves.Reserve0, reserves.Reserve1
	} else {
		reserve0, reserve1 = reserves.Reserve1, reserves.Reserve0
	}

	return getAmountOut(amountIn, reserve0, reserve1)
}
