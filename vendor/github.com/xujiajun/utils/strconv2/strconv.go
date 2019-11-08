package strconv2

import "strconv"

// StrToInt is shorthand for `strconv.Atoi(s)`.
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// IntToStr is shorthand for `strconv.Itoa(i)`.
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func IntToFloat64(i int) (float64, error) {
	return StrToFloat64(IntToStr(i))
}

func Int64ToFloat64(i int64) (float64, error) {
	return StrToFloat64(Int64ToStr(i))
}

func StrToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func StrToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
