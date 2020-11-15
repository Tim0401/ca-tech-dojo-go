package util

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type db struct {
	DbDriver string `json:"dbDriver" yaml:"dbDriver"`
	Dsn      string `json:"dsn" yaml:"dsn"`
}

type redis struct {
	Address  string `json:"address" yaml:"address"`
	Port     string `json:"port" yaml:"port"`
	Protocol string `json:"protocol" yaml:"protocol"`
}

// Config アプリケーション設定
type Config struct {
	Db    db    `json:"db" yaml:"db"`
	Redis redis `json:"redis" yaml:"redis"`
}

func LoadConfigForYaml() (*Config, error) {
	f, err := os.Open("../../config.yaml")
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}
