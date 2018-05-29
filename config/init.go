package config

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	log "github.com/inconshreveable/log15"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
	"path"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	once     sync.Once
	instance *Config
)

type Contract struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Abi     string `yaml:"abi"`
}
type Config struct {
	isNonce      bool
	ethClient    *ethclient.Client
	l            log.Logger
	cfgFile      string
	nonce        chan int
	nodeAccount  *accounts.Account
	nodeKeystore *keystore.KeyStore
	home         string

	IFile  string `yaml:"-"`
	OFile  string `yaml:"-"`
	PassWD string `yaml:"-"`

	EthAddr     string     `yaml:"eth_addr"`
	LogPath     string     `yaml:"log_path"`
	LogLevel    string     `yaml:"log_level"`
	Nonce       uint64     `yaml:"nonce"`
	KeystoreDir string     `yaml:"keystore"`
	Passphrase  string     `yaml:"passphrase"`
	GasLimit    int        `yaml:"gas_limit"`
	Gasprice    int        `yaml:"gas_price"`
	ChainId     int        `yaml:"chain_id"`
	Contracts   []Contract `yaml:"contracts"`
}

func (c *Config) LoadConfig() {
	d, err := ioutil.ReadFile(c.cfgFile)
	if err != nil {
		panic(fmt.Sprintf("配置文件读取错误\n%s", err.Error()))
	}

	if err := yaml.Unmarshal(d, c); err != nil {
		panic(fmt.Sprintf("%s\n%s", "配置文件加载错误", err.Error()))
	}

	c.cfgFile = path.Join(c.home, "kdata.yaml")
	instance.LogPath = path.Join(c.home, "log")
	instance.KeystoreDir = path.Join(c.home, "keystore")
}

func (c *Config) Dumps() {
	d, err := yaml.Marshal(c)
	if err != nil {
		panic(err.Error())
	}

	if err := ioutil.WriteFile(c.cfgFile, d, 0755); err != nil {
		panic(fmt.Sprintf("写入配置文件\n%s", err.Error()))
	}
}

func (c *Config) InitNode() {
	nodeKeystore := keystore.NewKeyStore(c.KeystoreDir, keystore.LightScryptN, keystore.LightScryptP)
	if len(nodeKeystore.Accounts()) == 0 {
		panic("node account not found")
	}

	nodeAccount := &nodeKeystore.Accounts()[0]
	if err := nodeKeystore.Unlock(*nodeAccount, c.Passphrase); err != nil {
		panic(fmt.Sprintf("%s\n%s", "账号解锁失败", err.Error()))
	}

	c.nodeKeystore = nodeKeystore
	c.nodeAccount = nodeAccount
}

func (t *Config) InitLog() {
	t.l = log.New()
	if t.LogLevel != "error" {
		ll, err := log.LvlFromString(t.LogLevel)
		if err != nil {
			panic(err.Error())
		}
		t.l.SetHandler(log.LvlFilterHandler(ll, log.StreamHandler(os.Stdout, log.TerminalFormat())))
	} else {
		h, err := log.FileHandler(t.LogPath, log.LogfmtFormat())
		if err != nil {
			t.l.Error(err.Error())
			panic(err.Error())
		}
		log.MultiHandler(
			log.LvlFilterHandler(log.LvlError, log.StreamHandler(os.Stderr, log.LogfmtFormat())),
			h,
		)
	}
}

func (c *Config) InitEthClient() {
	//client, err := ethclient.Dial(c.EthAddr)
	client, err := ethclient.Dial("ws://101.132.139.155:8867")
	if err != nil {
		panic(err.Error())
	}

	c.ethClient = client
}

func (c *Config) GetEthClient() *ethclient.Client {
	if c.ethClient == nil {
		panic("请初始化以太坊客户端")
	}
	return c.ethClient
}
