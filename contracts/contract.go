package contracts

import (
	"errors"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	kts "github.com/kooksee/pstoff/types"
)

type Contract struct {
	Address          common.Address
	ABI              abi.ABI
	eventNameHashMap map[string]string
}

func NewContract(address, abi string) (*Contract, error) {

	if address == "" || abi == "" {
		return nil, errors.New("contract address and abi cannot be blank")
	}

	contract := new(Contract)

	contract.Address = common.HexToAddress(address)

	if err := contract.InitABI(abi); err != nil {
		return nil, err
	}

	return contract, nil
}

func (contract *Contract) InitABI(ABIJson string) error {
	abiInstance, err := abi.JSON(strings.NewReader(ABIJson))

	if err != nil {
		return err
	}

	contract.ABI = abiInstance

	return nil
}

func (contract *Contract) Execute(method string, args ...interface{}) []byte {

	methodBytes, err := contract.ABI.Pack(method, args...)
	if err != nil {
		panic(err.Error())
	}

	tx := types.NewTransaction(
		cfg.GetNonce(),
		contract.Address,
		big.NewInt(0),
		big.NewInt(int64(cfg.GasLimit)),
		big.NewInt(int64(cfg.Gasprice)),
		methodBytes,
	)

	tx1, err := tx.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	return tx1
}

func (contract *Contract) AddRule(userAddress, roleType string) []byte {
	methodBytes, err := contract.ABI.Pack(
		"addRole",
		common.StringToAddress(userAddress),
		roleType,
	)
	if err != nil {
		logger.Error("AddRule error", "err", err)
		panic(err.Error())
	}

	tx := &kts.Tx{
		Nonce:    cfg.GetNonce(),
		To:       contract.Address.Hex(),
		Amount:   0,
		GasLimit: int64(cfg.GasLimit),
		GasPrice: int64(cfg.Gasprice),
		Data:     methodBytes,
	}

	return tx.Encode()
}

func Deploy(data []byte) []byte {

	tx := &kts.Tx{
		IsCreateContract: true,
		Nonce:            cfg.GetNonce(),
		To:               "",
		Amount:           0,
		GasLimit:         int64(cfg.GasLimit),
		GasPrice:         int64(cfg.Gasprice),
		Data:             data,
	}

	return tx.Encode()
}
