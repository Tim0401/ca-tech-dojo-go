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

type config struct {
	Db db `json:"db" yaml:"db"`
}

func LoadConfigForYaml() (*config, error) {
	f, err := os.Open("../../config.yaml")
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}
