package rusty

import (
	"fmt"
	"strconv"
	"strings"
)

type Map map[string]interface{}

func Get[T any](d Map, keypath string) *Result[T] {

	var segs []string = strings.Split(keypath, ".")

	var obj interface{}

	obj = d

	for fieldIndex, field := range segs {

		if fieldIndex == len(segs)-1 {

			var value interface{}

			switch o := obj.(type) {
			case []interface{}:
				value = o[ToString(field).ParseInt().Unwrap()]
			case []string:
				value = o[ToString(field).ParseInt().Unwrap()]
			case []int:
				value = o[ToString(field).ParseInt().Unwrap()]
			case Map:
				value = o[field]
			}

			v, ok := value.(T)

			if ok {
				return ToResult(v, nil)
			} else {
				return ToResult(v, fmt.Errorf("type assertion failed: %v", value))
			}
		}

		switch o := obj.(type) {
		case []interface{}:
			value := o[ToString(field).ParseInt().Unwrap()]
			switch v := value.(type) {
			case Map:
				obj = v
			case map[string]interface{}:
				obj = Map(v)
			case []any, []string, []int:
				obj = v
			}
		case Map:
			value := o[field]
			switch v := value.(type) {
			case Map:
				obj = v
			case map[string]any:
				obj = Map(v)
			case []any, []string, []int:
				obj = v
			}
		}

	}

	var v T

	return ToResult(v, fmt.Errorf("path not found: %s", keypath))

}

func GetInt(d Map, keypath string) *Result[int] {
	res := Get[string](d, keypath)

	if res.IsErr() {
		return ToResult(0, res.err)
	}

	v, err := strconv.ParseInt(res.data, 10, 64)

	return ToResult(int(v), err)
}
