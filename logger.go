package slf4go

import (
	"errors"
)

type LoggerBase interface {
	Logout(int, string, string, *[]byte) error
}

type Logger struct {
	Level  int
	Logger LoggerBase
}

const (
	LL_Trace = 1 << iota // log level
	LL_Debug
	LL_Info
	LL_Warn
	LL_Error
)

var logger = &Logger{Level: LL_Trace}

func InitLog(lv int, lb LoggerBase) error {
	if lv < LL_Trace || lv > LL_Error {
		return errors.New("invalid log level.")
	}

	logger.Level = lv
	logger.Logger = lb
	return nil
}
