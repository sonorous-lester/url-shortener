package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Db struct {
		UserName     string `yaml:"userName"`
		Password     string `yaml:"password"`
		Addr         string `yaml:"addr"`
		Port         int    `yaml:"port"`
		Database     string `yaml:"database"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
	} `yaml:"db"`
	NoSql struct {
		Network   string `yaml:"network"`
		Addr      string `yaml:"addr"`
		MaxIdle   int    `yaml:"maxIdle"`
		MaxActive int    `yaml:"maxActive"`
	} `yaml:"nosql"`
}

func Read() Config {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	return cfg
}
