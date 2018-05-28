package main

import (
	"github.com/kooksee/pstoff/config"
	"github.com/kooksee/pstoff/cmd"
	"github.com/kooksee/pstoff/contracts"
)

func main() {
	cfg := config.NewCfg("kdata")
	cfg.LoadConfig()
	cfg.InitLog()
	cfg.InitNode()

	contracts.Init()

	cmd.Init()
	cmd.RunCmd()

	cfg.Dumps()
}
