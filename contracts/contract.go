package contracts

import (
	"errors"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	signedTx, err := cfg.GetNodeKeyStore().SignTx(*cfg.GetNodeAccount(), tx, big.NewInt(int64(cfg.ChainId)))
	if err != nil {
		panic(err.Error())
	}

	tx1, err := signedTx.MarshalJSON()
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

	tx := types.NewTransaction(
		cfg.GetNonce(),
		contract.Address,
		big.NewInt(0),
		big.NewInt(int64(cfg.GasLimit)),
		big.NewInt(int64(cfg.Gasprice)),
		methodBytes,
	)

	signedTx, err := cfg.GetNodeKeyStore().SignTx(*cfg.GetNodeAccount(), tx, big.NewInt(int64(cfg.ChainId)))
	if err != nil {
		logger.Error("SignTx error", "err", err)
		panic(err.Error())
	}

	tx1, err := signedTx.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	return tx1
}

func Deploy(data []byte) []byte {

	tx := types.NewContractCreation(
		cfg.GetNonce(),
		big.NewInt(0),
		big.NewInt(int64(cfg.GasLimit)),
		big.NewInt(int64(cfg.Gasprice)),
		data,
	)

	signedTx, err := cfg.GetNodeKeyStore().SignTx(*cfg.GetNodeAccount(), tx, big.NewInt(int64(cfg.ChainId)))
	if err != nil {
		panic(err.Error())
	}

	tx1, err := signedTx.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	return tx1
}
