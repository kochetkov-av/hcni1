package main

import (
	"github.com/kochetkov-av/hcni1/cli"
	"github.com/kochetkov-av/hcni1/quoter"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	c := cli.New(logger.Named("cli"), quoter.New(logger.Named("quoter")))
	if err := c.RootCmd.Execute(); err != nil {
		logger.Fatal("command execution error", zap.Error(err))
	}
}
