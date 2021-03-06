package create

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/cli"

	"github.com/hiromaily/go-crypto-wallet/pkg/wallet/key"
	"github.com/hiromaily/go-crypto-wallet/pkg/wallet/wallets"
)

// TODO: this code is almost same to keygen wallet

// SeedCommand seed subcommand
type SeedCommand struct {
	name     string
	synopsis string
	ui       cli.Ui
	wallet   wallets.Signer
}

// Synopsis is explanation for this subcommand
func (c *SeedCommand) Synopsis() string {
	return c.synopsis
}

// Help returns usage for this subcommand
func (c *SeedCommand) Help() string {
	return `Usage: sign create seed [options...]
Options:
  -seed  given seed is used to store in database instead of generating new seed (development use)
`
}

// Run executes this subcommand
func (c *SeedCommand) Run(args []string) int {
	c.ui.Info(c.Synopsis())

	var (
		seed  string
		bSeed []byte
		err   error
	)

	flags := flag.NewFlagSet(c.name, flag.ContinueOnError)
	flags.StringVar(&seed, "seed", "", "given seed is used to store in database instead of generating new seed (development use)")
	if err := flags.Parse(args); err != nil {
		return 1
	}

	if seed == "" {
		seed = os.Getenv("SIGN_SEED")
		c.ui.Info("seed is found from environment variable")
	}

	if seed != "" {
		// store seed into database, not generate seed
		bSeed, err = c.wallet.StoreSeed(seed)
		if err != nil {
			c.ui.Error(fmt.Sprintf("fail to call StoreSeed() %+v", err))
			return 1
		}
	} else {
		// create seed
		bSeed, err = c.wallet.GenerateSeed()
		if err != nil {
			c.ui.Error(fmt.Sprintf("fail to call GenerateSeed() %+v", err))
			return 1
		}
	}
	c.ui.Info(fmt.Sprintf("seed: %s", key.SeedToString(bSeed)))

	return 0
}
