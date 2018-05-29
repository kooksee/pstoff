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

			if cfg.PassWD == "" {
				logger.Error("请输入传入密码参数")
				panic("请输入传入密码参数")
			}

			logger.Info("key store file", "file", cfg.KeystoreDir)

			a, err := keystore.StoreKey(cfg.KeystoreDir, cfg.PassWD, keystore.LightScryptN, keystore.LightScryptP)
			if err != nil {
				logger.Error("keystore.StoreKey error", "err", err)
				panic(err.Error())
			}

			logger.Info("new account", "account", a)
			logger.Info("account", "address", a.Hash().String())

			return nil
		},
	}
}
