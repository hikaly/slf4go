package slog

import (
	"errors"
	"fmt"
	//"log"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

const (
	LL_Trace = 1 << iota // log level
	LL_Debug
	LL_Info
	LL_Warn
	LL_Error
)

const (
	Color_Red    = 31
	Color_Green  = 32
	Color_Yellow = 33
	Color_Blue   = 34
	Color_Purple = 35
)

type SLF struct{}

func (slf SLF) Logout(l int, ts string, ls string, b *[]byte) error {
	return logger.Logout(l, ts, ls, b)
}

type SLogger struct {
	mu      sync.Mutex
	FileMap map[int]*os.File
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

func (sl SLogger) Logout(level int, ts string, ls string, b *[]byte) error {
	var wb []byte
	wb = append(wb, ts...)
	wb = append(wb, ' ')
	wb = append(wb, *b...)
	wb = append(wb, ' ')
	itoa(&wb, level, -1)
	wb = append(wb, ' ')
	wb = append(wb, ls...)
	if len(ls) == 0 || ls[len(ls)-1] != '\n' {
		wb = append(wb, '\n')
	}

	if sl.FileMap[level] == nil {
		return errors.New("log file not exist.")
	}

	sl.mu.Lock()
	defer sl.mu.Unlock()
	_, err := sl.FileMap[level].Write(wb)

	//OutputColorString(level, string(wb))
	return err
}

func (sl *SLogger) SetOutput(k int, f *os.File) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.FileMap[k] = f
}

var logPath = map[int]string{LL_Trace: "/logs/Trace/", LL_Debug: "/logs/Debug/", LL_Warn: "/logs/Warn/", LL_Error: "/logs/Error/", LL_Info: "/logs/Info/"}

var logger = &SLogger{
	FileMap: make(map[int]*os.File),
}

func InitSLogger() error {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return errors.New("runtime caller fail.")
	}

	file = path.Join(file, "..")
	fmt.Println(file)
	for k, v := range logPath {
		f, err := getLogIoWriter(file + v)
		if err != nil {
			return err
		}

		logPath[k] = file + v
		logger.FileMap[k] = f
	}

	// 每天更换输出文件
	go outputWriterCronCheck()
	return nil
}

func getLogIoWriter(v string) (*os.File, error) {
	filename := time.Now().Format("2006-01-02")
	fn := v + filename + ".log"
	return os.OpenFile(fn, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
}

func outputWriterCronCheck() {
	var sleepTime int
	for {
		sleepTime = 86400 - time.Now().Hour()*3600 - time.Now().Minute()*60 - time.Now().Second()
		time.Sleep(time.Duration(sleepTime) * time.Second)
		setLoggerWriter()
	}
}

//重定向输出文件
func setLoggerWriter() {
	for k, v := range logPath {
		f, err := getLogIoWriter(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if f == nil {
			continue
		}

		if logger.FileMap[k] != nil {
			logger.FileMap[k].Close()
		}

		fmt.Println("current file name: ", f.Name())
		logger.SetOutput(k, f)
	}
}

// 控制太输出字体染色
func OutputColorString(level int, s string) {
	baseFormat := "\x1b[1;%dm%s\x1b[0m"

	switch level {
	case LL_Warn:
		fmt.Println(fmt.Sprintf(baseFormat, Color_Yellow, s))
	case LL_Error:
		fmt.Println(fmt.Sprintf(baseFormat, Color_Red, s))
	case LL_Info:
		fmt.Println(fmt.Sprintf(baseFormat, Color_Green, s))
	default:
		fmt.Println(s)
	}
}
