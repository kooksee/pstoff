package cmd

import (
	"github.com/urfave/cli"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func AccountCmd() cli.Command {
	return cli.Command{
		Name:    "newAccount",
		Aliases: []string{"nc"},
		Usage:   "create account",
		Flags: []cli.Flag{
			passwdFlag(),
		},
		Action: func(c *cli.Context) error {

			a, err := keystore.StoreKey(cfg.KeystoreDir, cfg.PassWD, keystore.LightScryptN, keystore.LightScryptP)
			if err != nil {
				panic(err.Error())
			}

			logger.Info("account", "address", a.Hash().String())

			return nil
		},
	}
}
