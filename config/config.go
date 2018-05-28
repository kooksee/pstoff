package config

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
	log "github.com/inconshreveable/log15"
	"path"
	"github.com/kooksee/pstoff/cmn"
)

func (c *Config) GetNonce() int {
	return c.Nonce
}

func (c *Config) GetNodeAccount() *accounts.Account {
	if c.nodeAccount == nil {
		panic("please init node account")
	}
	return c.nodeAccount
}

func (c *Config) GetNodeKeyStore() *keystore.KeyStore {
	if c.nodeKeystore == nil {
		panic("please init nodeKeystore")
	}
	return c.nodeKeystore
}

func GetCfg() *Config {
	if instance == nil {
		panic("please init config")
	}
	return instance
}

func GetHomeDir(defaultHome string) string {
	if len(os.Args) > 2 && os.Args[len(os.Args)-2] == "--home" {
		defaultHome = os.Args[len(os.Args)-1]
		os.Args = os.Args[:len(os.Args)-2]
	}
	return defaultHome
}

func Log() log.Logger {
	cfg := GetCfg()
	if cfg.l == nil {
		panic("please init log")
	}
	return cfg.l
}

func NewCfg(defaultHomeDir string) *Config {
	defaultHomeDir = GetHomeDir(defaultHomeDir)
	instance = &Config{}

	instance.home = defaultHomeDir
	instance.LogLevel = "debug"
	instance.cfgFile = path.Join(defaultHomeDir, "kdata.yaml")
	instance.LogLevel = "debug"
	instance.LogPath = path.Join(defaultHomeDir, "log")
	instance.KeystoreDir = path.Join(defaultHomeDir, "keystore")
	instance.IFile = "input.json"
	instance.OFile = "output.json"
	instance.ChainId = 0
	instance.Passphrase = "Test123:::"
	instance.GasLimit = 750000
	instance.Gasprice = 100000000000
	instance.Nonce = 1
	instance.EthAddr = "http://localhost:8545"

	cmn.EnsureDir(instance.home, os.FileMode(0755))
	cmn.EnsureDir(instance.LogPath, os.FileMode(0755))
	cmn.EnsureDir(instance.KeystoreDir, os.FileMode(0755))

	return instance
}
