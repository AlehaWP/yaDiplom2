package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/AlehaWP/yaDiplom2.git/server/internal/logger"
	"github.com/caarlos0/env/v6"
)

const (
	PathSeparator = string(os.PathSeparator)
)

var cfg Config
var once sync.Once

type Config struct {
	ServAddr      string `env:"SERVER_ADDRESS" json:"server_address"`
	WorkDir       string `env:"WORK_DIR" json:"work_dir"`
	AppDir        string `json:"-"`
	DBConnStr     string `json:"-"` //`env:"DB_CONN_STR" json:"db_conn_str"`
	ConfigPath    string `env:"CONFIG_FOR_READ" json:"-"`
	TrustedSubnet string `env:"TRUSTED_SUBNET" json:"trusted_subnet"`
}

func (ic *Config) fillFromConf(oc *Config) {
	if len(oc.ServAddr) != 0 {
		ic.ServAddr = oc.ServAddr
	}

	if len(oc.WorkDir) != 0 {
		ic.WorkDir = oc.WorkDir
	}

	if len(oc.DBConnStr) != 0 {
		ic.DBConnStr = oc.DBConnStr
	}

	if len(oc.TrustedSubnet) != 0 {
		ic.TrustedSubnet = oc.TrustedSubnet
	}

	if len(oc.ConfigPath) != 0 {
		ic.readConfig(oc.ConfigPath)
	}
}

func (c *Config) readConfig(f string) {
	if f == "" {
		f = c.ConfigPath
	}

	if ok, _ := exists(f); !ok {
		logger.Info("Не найден файл конфигурации ", f)
		return
	}

	config := &Config{}

	configFile, err := os.Open(f)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.fillFromConf(config)
}

//checkEnv for get options from env to default application options.
func (c *Config) parseEnv() {
	e := &Config{}
	err := env.Parse(e)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.fillFromConf(e)

}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (c *Config) saveConfiguration() error {
	f := c.ConfigPath
	configFile, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer configFile.Close()
	if err != nil {
		return errors.New("не удалось найти файл " + f)
	}
	jsonParser := json.NewEncoder(configFile)
	jsonParser.Encode(c)
	return nil
}

//setFlags for get options from console to default application options.
func (c *Config) setFlags() {

	f := new(Config)

	flag.StringVar(&f.ServAddr, "s", c.ServAddr, "a server address string")
	flag.StringVar(&f.WorkDir, "w", c.WorkDir, "a work directory path for db and other files string")
	flag.StringVar(&f.DBConnStr, "d", c.DBConnStr, "a file storage name string")
	flag.StringVar(&f.ConfigPath, "c", "", "a config file path string")
	flag.StringVar(&f.TrustedSubnet, "t", c.TrustedSubnet, "a trusted ip CIDR xxx.xxx.xxx.xxx/32")

	flag.Parse()

	c.fillFromConf(f)

}

func createConfig() {

	AppDir, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	}
	cfg = Config{
		ServAddr:      "localhost:8080",
		WorkDir:       AppDir,
		AppDir:        AppDir,
		DBConnStr:     "user=ksei password=ksei dbname=yandex sslmode=disable",
		ConfigPath:    AppDir + PathSeparator + "gophepass.json",
		TrustedSubnet: "",
	}
	cfg.readConfig("")
	cfg.parseEnv()
	cfg.setFlags()
	cfg.saveConfiguration()

	logger.Info("Создан config")
}

// NewDefOptions return obj like Options interfase.
func NewConfig() Config {
	once.Do(createConfig)
	return cfg
}

// func (d defOptions) IsTrustedIp(ip string) bool {
// 	ip2 := net.ParseIP(ip)
// 	if ip2 == nil {
// 		return false
// 	}

// 	if len(d.trustedSubnet) == 0 {
// 		return false
// 	}

// 	_, n, err := net.ParseCIDR(d.trustedSubnet)
// 	if err != nil {
// 		return false
// 	}

// 	if ok := n.Contains(ip2); !ok {
// 		return false
// 	}
// 	return true
// }
