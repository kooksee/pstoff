package contracts

import (
	"github.com/kooksee/pstoff/config"
	"fmt"
)

var cfg *config.Config

func Init() {
	cfg = config.GetCfg()
	if err := initContracts(); err != nil {
		panic(fmt.Sprintf("初始化合约错误\n%s", err.Error()))
	}
}
