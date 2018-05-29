package cmd

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
)

func TestAddrule(t *testing.T) {
	t.Logf(common.HexToAddress("0x46e97077b5efd97c5f7316e955052489772dc382").String())
	t.Logf(common.HexToAddress("0x0174aDd69d9102d223f6E61Ff82DC0b36B8f3ae8").String())
	fmt.Sprintln(common.StringToAddress("46e97077b5efd97c5f7316e955052489772dc382").Hex())

}
