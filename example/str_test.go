package example

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var w = []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}

func BenchmarkAddStringWithOperator(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for _, v := range w {
			s += v
		}
		s = ""
	}
	//fmt.Println(s)
}

func BenchmarkAddStringWithSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", w[0], w[1], w[2], w[3], w[4], w[5], w[6], w[7], w[8], w[9])
	}
}

func BenchmarkAddStringWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(w, ",")
	}
}

func BenchmarkAddStringWithBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for _, v := range w {
			buffer.WriteString(v)
		}
		_ = buffer.String()
	}
}

func BenchmarkFormatString(b *testing.B) {
	f := "data:{}1,{}2,{}3,{}4,{}5,{}6,{}7,{}8,{}9,{}0"
	sl := strings.Split(f, "{}")
	for i := 0; i < b.N; i++ {
		if len(sl)-1 != len(w) {
			fmt.Println("!!!!!!", len(sl), sl)
		}

		var s string
		for k, v := range sl {
			s = s + v + w[k]
			if k+2 >= len(sl) {
				break
			}
		}
	}
}

/*
func BenchmarkMakeLogString(b *testing.B) {
	var rb []byte
	tb := []byte(GetCurrentTimeString())
	b := []byte(fmt.Sprintf(format, data...))
	_, f, l, ok := runtime.Caller(1)
	if !ok {
		f = "???"
		l = 0
	}

	rb = append(rb, tb)
	rb = append(rb, ' ')
	rb = append(rb, b)
	rb = append(rb, f)
	itoa(rb, l, -1)

	return string(rb)
}
*/
