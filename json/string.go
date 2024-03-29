package json

import (
	"encoding/json"

	"github.com/koroutine/rusty"
)

func ToString(value any) *rusty.Result[string] {

	res, err := json.Marshal(value)

	return rusty.ToResult(string(res), err)

}

func FromString[T any](value string) *rusty.Result[T] {

	var data T
	err := json.Unmarshal([]byte(value), &data)

	return rusty.ToResult(data, err)
}
