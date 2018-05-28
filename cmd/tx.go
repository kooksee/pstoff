package cmd

import (
	"github.com/urfave/cli"
	"context"
	"time"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"encoding/json"
)

func TxCmd() cli.Command {
	return cli.Command{
		Name:    "sentTx",
		Aliases: []string{"st"},
		Usage:   "sent tx",
		Flags: []cli.Flag{
			inputFileFlag(),
			outputFileFlag(),
		},
		Action: func(c *cli.Context) error {
			cfg.InitEthClient()

			client := cfg.GetEthClient()
			d, err := ioutil.ReadFile(cfg.IFile)
			if err != nil {
				panic(err.Error())
			}

			ds := make([]string, 0)
			if err := json.Unmarshal(d, &ds); err != nil {
				panic(err.Error())
			}

			for _, t := range ds {
				tx := &types.Transaction{}
				if err := tx.UnmarshalJSON(common.FromHex(t)); err != nil {
					panic(err.Error())
				}

				ctx2, _ := context.WithTimeout(context.Background(), time.Minute)
				if err := client.SendTransaction(ctx2, tx); err != nil {
					panic(err.Error())
				}

				logger.Info("SendTransaction", "hash", tx.Hash().String())
			}
			return nil
		},
	}
}

/*
func(c *cli.Context) error {
			cfg.InitEthClient()

			var testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")


			client := cfg.GetEthClient()
			d, err := ioutil.ReadFile(cfg.IFile)
			if err != nil {
				panic(err.Error())
			}

			ds := make([]string, 0)
			if err := json.Unmarshal(d, &ds); err != nil {
				panic(err.Error())
			}

			//backend := backends.NewSimulatedBackend(core.GenesisAlloc{
			//	crypto.PubkeyToAddress(testKey.PublicKey): {Balance: big.NewInt(10000000000)},
			//})
			//
			//backend.SendTransaction(ctx, tx)
			//backend.Commit()

			return nil
		}
 */
