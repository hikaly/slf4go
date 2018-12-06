package slog

/*
func toString(arg interface{}) string {
	switch arg.(type) {
	case bool:
		i, _ := arg.(bool)
		if i {
			return StrTrue
		} else {
			return StrFalse
		}

	case int8:
		i, _ := arg.(int8)
		return strconv.Itoa(int(i))

	case int16:
		i, _ := arg.(int16)
		return strconv.Itoa(int(i))

	case int32:
		i, _ := arg.(int32)
		return strconv.Itoa(int(i))

	case int64:
		i, _ := arg.(int64)
		return strconv.Itoa(int(i))

	case int:
		i, _ := arg.(int)
		return strconv.Itoa(int(i))

	case uint8:
		i, _ := arg.(uint8)
		return strconv.Itoa(int(i))

	case uint16:
		i, _ := arg.(uint16)
		return strconv.Itoa(int(i))

	case uint32:
		i, _ := arg.(uint32)
		return strconv.Itoa(int(i))

	case uint64:
		i, _ := arg.(uint64)
		return strconv.Itoa(int(i))

	case uint:
		i, _ := arg.(uint)
		return strconv.Itoa(int(i))

	case []byte:
		i, _ := arg.([]byte)
		return string(i)

	case string:
		s, _ := arg.(string)
		return s

	default:
		//b := bytes.Buffer
		tp := reflect.TypeOf(arg)
		v := reflect.ValueOf(arg)

		var b bytes.Buffer
		b.WriteString("{")
		for i := 0; i < tp.NumField(); i++ {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(tp.Field(i).Name)
			b.WriteString(":")
			b.WriteString(toString(v.Field(i).Interface()))
		}

		b.WriteString("}")
		return b.String()
	}

	return ""
}

func formatAndLog(format string, data ...interface{}) string {

}
*/
