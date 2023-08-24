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

func getAmountIn(amountOut, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	if amountOut.Cmp(bigZero) <= 0 {
		return bigZero, errors.New("insufficient output amount")
	}

	if reserveIn.Cmp(bigZero) <= 0 || reserveOut.Cmp(bigZero) <= 0 {
		return bigZero, errors.New("insufficient liquidity")
	}

	numerator, denominator := new(big.Int), new(big.Int)

	// uint numerator = reserveIn.mul(amountOut).mul(1000);
	numerator.Mul(reserveIn, amountOut)
	numerator.Mul(numerator, big1000)

	// uint denominator = reserveOut.sub(amountOut).mul(997);
	denominator.Sub(reserveOut, amountOut)
	denominator.Mul(denominator, big997)

	// amountIn = (numerator / denominator).add(1);
	amountIn := numerator.Div(numerator, denominator)
	amountIn.Add(amountIn, big.NewInt(1))
	return amountIn, nil
}
