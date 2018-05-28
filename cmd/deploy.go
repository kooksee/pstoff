package cmd

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"fmt"
	"github.com/kooksee/pstoff/contracts"
	"quorum/common"
	"encoding/json"
)

func DeployCmd() cli.Command {
	return cli.Command{
		Name:    "deploy",
		Aliases: []string{"dp"},
		Usage:   "deploy contract",
		Flags: []cli.Flag{
			inputFileFlag(),
			outputFileFlag(),
		},
		Action: func(c *cli.Context) error {

			d, err := ioutil.ReadFile(cfg.IFile)
			if err != nil {
				panic(err.Error())
			}

			iFile := make([]string, 0)
			if err := json.Unmarshal(d, &iFile); err != nil {
				panic(err.Error())
			}

			oFile := make([]string, 0)
			for _, ifile := range iFile {
				d1 := common.FromHex(ifile)
				oFile = append(oFile, common.ToHex(contracts.Deploy(d1)))
			}

			d1, err := json.Marshal(oFile)
			if err != nil {
				panic(err.Error())
			}
			if err := ioutil.WriteFile(cfg.OFile, d1, 0755); err != nil {
				panic(fmt.Sprintf("写入失败\n%s", err.Error()))
			}
			return nil
		},
	}
}
