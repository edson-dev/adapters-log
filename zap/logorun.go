package adapters

import (
	"go-pocketbase/pkg/adapters-log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strings"
	"sync"
)

type Log struct {
	instance *zap.Logger
	level    zap.AtomicLevel
}

var once sync.Once

func (z Log) New() adapters.Log {
	once.Do(func() {
		z.level = zap.NewAtomicLevel()
		z.instance = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.Lock(os.Stdout),
			z.level,
		))
		z.SetLevel("")
		defer z.instance.Sync()
	})
	return z
}

func (z Log) SetLevel(level string) {
	switch level {
	case "fatal":
		z.level.SetLevel(zap.FatalLevel)
	case "error":
		z.level.SetLevel(zap.ErrorLevel)
	case "warn":
		z.level.SetLevel(zap.WarnLevel)
	case "info":
		z.level.SetLevel(zap.InfoLevel)
	default:
		z.level.SetLevel(zap.DebugLevel)
	}
}
func (z Log) GetLevel() string {
	return z.level.String()
}
func (z Log) Throw(err error) {
	if err != nil {
		pc, _, line, _ := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		fileStructure := strings.Split(details.Name(), "/")
		z.instance.Error(err.Error(),
			zap.String("file", fileStructure[len(fileStructure)-1]),
			zap.Int("line", line),
			zap.String("type", "relevant"))
	}
}
func (z Log) Debug(message string) {
	z.instance.Debug(message)
}
func (z Log) Info(message string) {
	z.instance.Info(message)
}
func (z Log) Warn(message string) {
	z.instance.Warn(message)
}
func (z Log) Error(message string) {
	z.instance.Error(message)
}
func (z Log) Fatal(message string) {
	z.instance.Fatal(message)
}
