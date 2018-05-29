package main

import (
	"github.com/kooksee/pstoff/config"
	"github.com/kooksee/pstoff/cmd"
)

func main() {
	cfg := config.NewCfg("kdata")
	cfg.LoadConfig()
	cfg.InitLog()
	cfg.InitNode()

	cmd.Init()
	cmd.RunCmd()
}
