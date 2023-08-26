package quoter

import (
	"errors"
	"math/big"
)

var (
	bigZero = big.NewInt(0)
	big997  = big.NewInt(997)
	big1000 = big.NewInt(1000)
)

func getAmountOut(amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	if amountIn.Cmp(bigZero) <= 0 {
		return bigZero, errors.New("insufficient input amount")
	}

	if reserveIn.Cmp(bigZero) <= 0 || reserveOut.Cmp(bigZero) <= 0 {
		return bigZero, errors.New("insufficient liquidity")
	}

	inputAmountWithFee, numerator, denominator := new(big.Int), new(big.Int), new(big.Int)

	// uint amountInWithFee = amountIn.mul(997);
	inputAmountWithFee.Mul(amountIn, big997)

	// uint numerator = amountInWithFee.mul(reserveOut);
	numerator.Mul(inputAmountWithFee, reserveOut)

	// uint denominator = reserveIn.mul(1000).add(amountInWithFee);
	denominator.Mul(reserveIn, big1000)
	denominator.Add(denominator, inputAmountWithFee)

	// amountOut = numerator / denominator;
	return numerator.Div(numerator, denominator), nil
}
