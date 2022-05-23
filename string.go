package rusty

import (
	"strconv"
	"strings"
)

type String struct {
	value string
}

func ToString(value string) *String {
	return &String{
		value,
	}
}

func (s *String) Replace(old, new string) *String {
	s.value = strings.ReplaceAll(s.value, old, new)

	return s
}

func (s *String) Split(sep string) []string {
	return strings.Split(s.value, sep)
}

func (s *String) ParseInt() *Result[int64] {
	return ToResult(strconv.ParseInt(s.value, 10, 64))
}

func (s *String) ParseFloat() *Result[float64] {
	return ToResult(strconv.ParseFloat(s.value, 64))
}
