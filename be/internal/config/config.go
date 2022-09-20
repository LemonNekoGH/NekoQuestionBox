package config

import (
	"io/ioutil"
	"neko-question-box-be/internal/logger"
	"os"

	"gopkg.in/yaml.v3"

	_ "embed"
)

type Config struct {
	Port     int `yaml:"port"`
	Database *struct {
		Host     string `yaml:"host"`     // 主机
		Port     int    `yaml:"port"`     // 端口号
		Username string `yaml:"username"` // 用户名
		Password string `yaml:"password"` // 密码
		Database string `yaml:"database"` // 使用的数据库
	} `yaml:"database"`
}

//go:embed config.test.yaml
var testCfg []byte
var Conf *Config // 配置

// 初始化配置
func InitConfig(forTest bool) {
	var confContent []byte
	var err error = nil
	// 读取环境变量
	configFile := os.Getenv("QBOX_CONFIG_PATH")
	if forTest {
		// 尝试读取根目录的配置文件
		confContent = testCfg
	} else {
		var file *os.File
		// 读取环境变量中指定的文件
		file, err = os.Open(configFile)
		if err != nil {
			panic(err)
		}
		confContent, err = ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
	}

	// 读取完成，映射成对象
	Conf = new(Config)
	err = yaml.Unmarshal(confContent, Conf)
	if err != nil {
		panic(err)
	}
	logger.Infof("配置文件加载成功")
}
