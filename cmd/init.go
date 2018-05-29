package cmd

import (
	"github.com/inconshreveable/log15"
	"github.com/kooksee/pstoff/config"
	"github.com/urfave/cli"
)

var (
	logger log15.Logger
	cfg    *config.Config
)

func Init() {
	cfg = config.GetCfg()
	logger = config.Log().New("package", "cmd")
}

func logLevelFlag() cli.StringFlag   { return cli.StringFlag{Name: "ll", Value: cfg.LogLevel, Destination: &cfg.LogLevel, Usage: "log level"} }
func inputFileFlag() cli.StringFlag  { return cli.StringFlag{Name: "i", Value: cfg.IFile, Destination: &cfg.IFile, Usage: "input file"} }
func outputFileFlag() cli.StringFlag { return cli.StringFlag{Name: "o", Value: cfg.OFile, Destination: &cfg.OFile, Usage: "output file"} }
func passwdFlag() cli.StringFlag     { return cli.StringFlag{Name: "p", Value: cfg.PassWD, Destination: &cfg.PassWD, Usage: "password"} }
