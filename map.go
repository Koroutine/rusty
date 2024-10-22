package rusty

import (
	"fmt"
	"strconv"
	"strings"
)

type Map map[string]interface{}

func Set(d Map, keypath string, value any) *Result[Map] {
	// Update Map with new value set at keypath (through cloning)

	// Ignore any escaped dots in the keypath
	keypath = strings.ReplaceAll(keypath, "\\.", "_")

	var segs []string = strings.Split(keypath, ".")

	var obj interface{}

	obj = d

	for fieldIndex, field := range segs {

		if fieldIndex == len(segs)-1 {

			switch o := obj.(type) {
			case []interface{}:
				o[ToString(field).ParseInt().Unwrap()] = value
			case []string:

				v, ok := value.(string)
				if !ok {
					return ToResult(d, fmt.Errorf("type assertion failed: %v", value))
				}

				o[ToString(field).ParseInt().Unwrap()] = v
			case []int:

				v, ok := value.(int)
				if !ok {
					return ToResult(d, fmt.Errorf("type assertion failed: %v", value))
				}

				o[ToString(field).ParseInt().Unwrap()] = v
			case Map:
				o[field] = value
			}

			return ToResult(d, nil)
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

			if value == nil {
				value = make(map[string]interface{})
				o[field] = value
			}

			switch v := value.(type) {
			case Map:
				obj = v
			case map[string]interface{}:
				obj = Map(v)
			case []any, []string, []int:
				obj = v
			}
		}

	}

	return ToResult(d, fmt.Errorf("path not found: %s", keypath))
}

func Get[T any](d Map, keypath string) *Result[T] {
	// Ignore any escaped dots in the keypath
	keypath = strings.ReplaceAll(keypath, "\\.", "_")

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

func GetInt64(d Map, keypath string) *Result[int64] {
	res := Get[string](d, keypath)

	if res.IsErr() {
		return ToResult(int64(0), res.err)
	}

	v, err := strconv.ParseInt(res.data, 10, 64)

	return ToResult(v, err)
}
