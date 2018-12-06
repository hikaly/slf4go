package slf4go

import (
	"fmt"
	"runtime"
	"time"
)

func GetCurrentTimeString() string {
	tn := time.Now()
	return tn.Format("2006/01/02-15:04:05.0000")
}

func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func GetLogFileNameAndLine() *[]byte {
	var rb []byte
	_, f, l, ok := runtime.Caller(2)
	if !ok {
		f = "???"
		l = 0
	}

	rb = append(rb, f...)
	rb = append(rb, ':')
	itoa(&rb, l, -1)
	return &rb
}

//
// log facade
func Trace(format string, data ...interface{}) {
	Logout(LL_Trace, GetCurrentTimeString(), fmt.Sprintf(format, data...), GetLogFileNameAndLine())
}

func Debug(format string, data ...interface{}) {
	Logout(LL_Debug, GetCurrentTimeString(), fmt.Sprintf(format, data...), GetLogFileNameAndLine())
}

func Info(format string, data ...interface{}) {
	Logout(LL_Info, GetCurrentTimeString(), fmt.Sprintf(format, data...), GetLogFileNameAndLine())
}

func Warn(format string, data ...interface{}) {
	Logout(LL_Warn, GetCurrentTimeString(), fmt.Sprintf(format, data...), GetLogFileNameAndLine())
}

func Error(format string, data ...interface{}) {
	Logout(LL_Error, GetCurrentTimeString(), fmt.Sprintf(format, data...), GetLogFileNameAndLine())
}

func Logout(lv int, ts string, ls string, nal *[]byte) {
	// if have logger, forward to logger
	// else, discard
	if logger.Logger != nil {
		if logger.Level > lv {
			return
		}

		err := logger.Logger.Logout(lv, ts, ls, nal)
		if err != nil {
			// todo:
			fmt.Println(err)
		}
	}
}
