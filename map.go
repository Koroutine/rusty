package rusty

import (
	"fmt"
	"strings"
)

type Map map[string]interface{}

func Get[T any](d Map, keypath string) *Result[T] {

	var segs []string = strings.Split(keypath, ".")

	obj := d

	for fieldIndex, field := range segs {

		if fieldIndex == len(segs)-1 {
			v, ok := obj[field].(T)

			if ok {
				return ToResult(v, nil)
			} else {
				return ToResult(v, fmt.Errorf("type assertion failed: %v", obj[field]))
			}
		}

		switch obj[field].(type) {
		case Map:
			obj = obj[field].(Map)
		case map[string]interface{}:
			obj = Map(obj[field].(map[string]interface{}))
		}

	}

	var v T

	return ToResult(v, fmt.Errorf("path not found: %s", keypath))

}
