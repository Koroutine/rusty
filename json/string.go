package json

import (
	"encoding/json"

	"github.com/koroutine/rusty"
)

func ToString(value any) *rusty.Result[string] {

	res, err := json.Marshal(value)

	return rusty.ToResult(string(res), err)

}
