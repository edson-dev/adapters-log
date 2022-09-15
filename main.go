package adapters

import (
	log2 "github.com/edson-dev/adapters-log/log"
	"github.com/edson-dev/adapters-log/zap"
)

func print(i Log) {
	println(i.GetLevel())
}

func main() {
	var log Log
	log = adapters.Log{}.New()
	log2 := log2.Log{}.New()
	print(log)
	log.Debug("erro")
	log2.Debug("erro")
}
