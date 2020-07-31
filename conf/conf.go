package conf

import (
	"demo/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type conf struct {
	Server server `yaml:"server"`
	Db db `yaml:"db"`
	MyLog myLog `yaml:"myLog"`
}

type server struct {
	Address string `yaml:"address"'`
}

type db struct {
	Dialects string `yaml:"dialects"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Db string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset string `yaml:"charset"`
	MaxIdle int `yaml:"maxIdle"`
	MaxOpen int `yaml:"maxOpen"`
}

type myLog struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

var Conf *conf

func init() {
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)

	if err != nil {
		panic(err)
	}

	util.SetLogFile(filepath.Join(Conf.MyLog.Path, Conf.MyLog.Name))

	logger := util.Log()
	logger.Info("config file read success")


}
