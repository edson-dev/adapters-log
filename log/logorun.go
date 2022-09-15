package adapters

import (
	"go-pocketbase/pkg/adapters-log"
	"log"
	"sync"
)

type Log struct {
	instance *log.Logger
	level    string
}

var once sync.Once

func (l Log) New() adapters.Log {
	once.Do(func() {
		l.level = "debug"
		l.instance = log.Default()
	})
	return l
}
func (l Log) SetLevel(level string) {
	l.level = level
}
func (l Log) GetLevel() string {
	return l.level
}

func (l Log) Throw(err error) {
	l.instance.Print(err.Error())
}
func (l Log) Debug(message string) {
	l.instance.Print(message)
}
func (l Log) Info(message string) {
	l.instance.Print(message)
}
func (l Log) Warn(message string) {
	l.instance.Print(message)
}
func (l Log) Error(message string) {
	l.instance.Print(message)
}
func (l Log) Fatal(message string) {
	l.instance.Print(message)
}
