package main

import (
	"fmt"

	"slf4go"
	"slf4go/slog"
)

func main() {
	err := slog.InitSLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = slf4go.InitLog(slf4go.LL_Debug, slog.SLF{})
	if err != nil {
		fmt.Println(err)
	}

	slf4go.Warn("this is a test string")
	slf4go.Warn("this is a test string for %s.", "11111")

	//tb := []byte("this is a test string.")
	//sl.Logout(1, &tb)
}
