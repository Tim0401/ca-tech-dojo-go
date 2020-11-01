package main

import (
	"ca-tech-dojo-go/pkg/cago"
	"ca-tech-dojo-go/pkg/util"
	"time"
)

const location = "Asia/Tokyo"

var gConfig *util.Config

func init() {
	// timezone設定
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	// config読み込み
	config, err := util.LoadConfigForYaml()
	if err != nil {
		panic(err.Error())
	}
	gConfig = config
}

func main() {
	cago.Serve(gConfig)
}
