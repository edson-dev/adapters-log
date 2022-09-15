package adapters

type Log interface {
	New() Log
	SetLevel(level string)
	GetLevel() string
	Throw(err error)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}
