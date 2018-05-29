package contracts

import (
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
	Btc              []byte
	eventNameHashMap map[string]string
}

func NewContract(address, abi string, btc []byte) (*Contract, error) {

	contract := new(Contract)

	contract.Address = common.HexToAddress(address)
	contract.Btc = btc

	if err := contract.InitABI(abi); err != nil {
		logger.Error("contract.InitABI error", "err", err, "name", address)
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

func (c *Contract) JoinNode(method, address string) []byte {
	methodBytes, err := c.ABI.Pack(
		method,
		common.HexToAddress(address),
	)
	if err != nil {
		logger.Error("JoinNode error", "err", err)
		panic(err.Error())
	}

	tx := &kts.Tx{
		Nonce:    cfg.GetNonce(),
		To:       c.Address.Hex(),
		Amount:   0,
		GasLimit: int64(cfg.GasLimit),
		GasPrice: int64(cfg.Gasprice),
		Data:     methodBytes,
	}

	return tx.Encode()
}

func (contract *Contract) AddWhiteList(method string, address string) []byte {
	methodBytes, err := contract.ABI.Pack(
		method,
		common.HexToAddress(address),
	)
	if err != nil {
		logger.Error("AddWhiteList error", "err", err)
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

func (contract *Contract) AddRule(mthName, userAddress, roleType string) []byte {
	methodBytes, err := contract.ABI.Pack(
		mthName,
		common.HexToAddress(userAddress),
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
