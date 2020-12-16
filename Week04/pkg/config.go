package pkg

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf = &Config{}

type Config struct {
	Db         DbConfig         `yaml:"db"`
	HttpServer HttpServerConfig `yaml:"http_server"`
}

func NewConfig(path string) *Config {

	dataFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Errorf("get fail err:%v", err))
	}
	err = yaml.Unmarshal(dataFile, Conf)
	if err != nil {
		log.Fatal(fmt.Errorf("Unmarshal fail err:%v", err))
	}
	fmt.Printf("%+v\n", Conf)
	return Conf
}

//数据库
type DbConfig struct {
	DNS string `yaml:"dns"`
}

//http服务端口
type HttpServerConfig struct {
	Addr string `yaml:"addr"`
}
