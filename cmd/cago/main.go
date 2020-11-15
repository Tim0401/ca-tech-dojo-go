package main

import (
	"ca-tech-dojo-go/pkg/cago"
	"ca-tech-dojo-go/pkg/util"
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"time"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
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

	// 乱数シード
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}

func main() {
	cago.Serve(gConfig)
}
