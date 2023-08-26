package cli

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type (
	Quoter interface {
		Quote(ethereumURL string, poolID, fromToken, toToken common.Address, amountIn *big.Int) (*big.Int, error)
	}

	Cli struct {
		logger  *zap.Logger
		RootCmd *cobra.Command

		quoter Quoter
	}

	Flags struct {
		EthereumURL string

		PoolID      string
		FromToken   string
		ToToken     string
		InputAmount string
	}
)

func New(logger *zap.Logger, quoter Quoter) *Cli {
	c := &Cli{
		logger: logger,
		quoter: quoter,
	}

	flags := new(Flags)

	cmd := &cobra.Command{
		Use:   "quoter",
		Short: "Uniswap V2 pool quoter",
		RunE: func(cmd *cobra.Command, r []string) error {
			return c.quote(flags)
		},
		Example: "quoter -e https://mainnet.infura.io/v3/d1c035ea43f4494c951b3d216e412691 -p 0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852 -f 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2 -t 0xdac17f958d2ee523a2206206994597c13d831ec7 -a 1e18",
	}

	cmd.Flags().StringVarP(&flags.EthereumURL, "ethereum_url", "e", "", "Ethereum provider URL")

	cmd.Flags().StringVarP(&flags.PoolID, "pool_id", "p", "", "Pool address")
	cmd.Flags().StringVarP(&flags.FromToken, "from_token", "f", "", "From token (address)")
	cmd.Flags().StringVarP(&flags.ToToken, "to_token", "t", "", "To token (address)")
	cmd.Flags().StringVarP(&flags.InputAmount, "input_amount", "a", "", "Imput amount")

	cmd.MarkFlagRequired("ethereum_url")

	cmd.MarkFlagRequired("pool_id")
	cmd.MarkFlagRequired("from_token")
	cmd.MarkFlagRequired("to_token")
	cmd.MarkFlagRequired("input_amount")

	c.RootCmd = cmd

	return c
}

func (c *Cli) quote(flags *Flags) error {
	if !common.IsHexAddress(flags.PoolID) {
		return errors.New("pool_id is not valid eth address")
	}
	if !common.IsHexAddress(flags.FromToken) {
		return errors.New("from_token is not valid eth address")
	}
	if !common.IsHexAddress(flags.ToToken) {
		return errors.New("to_token is not valid eth address")
	}

	f, _, err := big.ParseFloat(flags.InputAmount, 10, 0, big.ToNearestEven)
	if err != nil {
		return err
	}
	amountIn, _ := f.Int(nil)

	result, err := c.quoter.Quote(flags.EthereumURL, common.HexToAddress(flags.PoolID), common.HexToAddress(flags.FromToken), common.HexToAddress(flags.ToToken), amountIn)
	if err != nil {
		return err
	}

	c.logger.Info("RESULT", zap.String("result", result.String()))

	return nil
}
