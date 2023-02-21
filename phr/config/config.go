package config

import (
	"encoding/json"
	"log"
	mrand "math/rand"
	"os"
	"path/filepath"
	"regexp"
	"sardines/err"
	"sardines/tool"
	"strconv"
	"strings"

	"github.com/libp2p/go-libp2p/core/crypto"
)

var (
	cpath     string
	kpath     string
	Ktab      string
	Dir       string
	Manifest  string
	Downloads string
	WD        string
)

func init() {
	wd, _ := os.Getwd()
	WD = wd
	Dir = filepath.Join(wd, "/data")
	kpath = filepath.Join(Dir, "/priv_key")
	Ktab = filepath.Join(Dir, "/key_tab.db")
	cpath = filepath.Join(Dir, "/config.json")
	Manifest = filepath.Join(Dir, "/manifest.json")
	Downloads = filepath.Join(wd, "/DownloadFiles")
}

type Config struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IP         string `json:"ip"`
	Port       string `json:"port"`
	RandomSeed int64  `json:"randomSeed"`
	// Private Key
	PrvKey crypto.PrivKey `json:"prvKey"`
	// Bootstrap Node
	BootstrapNode string `json:"bootstrapNode"`
}

func (c *Config) Save() bool {
	c.PrvKey = nil
	b, _ := json.Marshal(*c)
	err := tool.WriteFile(b, cpath)
	if err != nil {
		return false
	}
	return true
}

func (c *Config) LoadAll() error {
	data, err2 := tool.LoadFile(cpath)
	if err2 != nil {
		return err2
	}
	err2 = json.Unmarshal(data, c)
	if err2 != nil {
		return err2
	}
	pk, _ := tool.LoadFile(kpath)
	if len(pk) == 0 {
		// prvKey is empty
		log.Printf("you haven't configure your own private key, so that your Peer ID could be variable")
		c.PrvKey = nil
	} else {
		c.PrvKey, _ = crypto.UnmarshalPrivateKey(pk)
	}
	return nil
}

func (c *Config) Load() error {
	data, err2 := tool.LoadFile(cpath)
	if err2 != nil {
		return err2
	}
	err2 = json.Unmarshal(data, c)
	if err2 != nil {
		return err2
	}
	c.PrvKey = nil
	return nil
}

func (c *Config) AddrString() string {
	// "/ip4/127.0.0.1/tcp/2000"
	return "/ip4/" + c.IP + "/tcp/" + c.Port
}

// 获取私钥

func (c *Config) GenKey() error {
	if c.RandomSeed > 0 {
		r := mrand.New(mrand.NewSource(c.RandomSeed))
		prvKey, _, err := crypto.GenerateRSAKeyPair(2048, r)
		if err != nil {
			return err
		}
		c.PrvKey = prvKey
		return nil
	} else {
		return err.IllegalSeed
	}
}

func (c *Config) SaveKey() error {
	if c.PrvKey != nil {
		b, err := crypto.MarshalPrivateKey(c.PrvKey)
		if err != nil {
			return err
		}
		if err = tool.WriteFile(b, kpath); err != nil {
			return err
		}
	}
	return nil
}

// New Get a Configuration
func New(username, pwd, ipAddr string, rs int64, bn string) (*Config, error) {
	strs := strings.Split(ipAddr, ":")
	if strs[0] == "" {
		strs[0] = "127.0.0.1"
	}

	if b := ipFormatCheck(strs); b {
		return &Config{
			Username:      username,
			Password:      pwd,
			IP:            strs[0],
			Port:          strs[1],
			RandomSeed:    rs,
			PrvKey:        nil,
			BootstrapNode: bn,
		}, nil
	} else {
		return nil, err.IllFormedIP
	}
}

func ipFormatCheck(ipAddr []string) bool {
	compile, _ := regexp.Compile(`((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`)
	if b := compile.MatchString(ipAddr[0]); b {
		port, err := strconv.ParseInt(ipAddr[1], 10, 64)
		if err != nil || port < 0 || port > 65535 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
