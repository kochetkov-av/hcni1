# Uniswap V2 quoter

Cli V2 pool quote calculator

For help
```shell
go run ./cmd/quoter/main.go --help
```

```shell
Uniswap V2 pool quoter

Usage:
  quoter [flags]

Examples:
quoter -e https://mainnet.infura.io/v3/d1c035ea43f4494c951b3d216e412691 -p 0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852 -f 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2 -t 0xdac17f958d2ee523a2206206994597c13d831ec7 -a 1e18

Flags:
  -e, --ethereum_url string   Ethereum provider URL
  -f, --from_token string     From token (address)
  -h, --help                  help for quoter
  -a, --input_amount string   Imput amount
  -p, --pool_id string        Pool address
  -t, --to_token string       To token (address)
```
