package adapters

import (
	"github.com/edson-dev/adapters-log"
	"sync"
)

type Log struct {
}

var once sync.Once

func (z Log) New() adapters.Log {
	return z
}

func (z Log) SetLevel(level string) {
}
func (z Log) GetLevel() string {
	return ""
}
func (z Log) Throw(err error) {
	println(err.Error())
}
func (z Log) Debug(message string) {
	println(message)
}
func (z Log) Info(message string) {
	println(message)
}
func (z Log) Warn(message string) {
	println(message)
}
func (z Log) Error(message string) {
	println(message)
}
func (z Log) Fatal(message string) {
	println(message)
}
