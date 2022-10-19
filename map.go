package rusty

import "strings"

type Map map[string]interface{}

func Get(d Map, keypath string) interface{} {

  var segs []string = strings.Split(keypath, ".")

  obj := d

  for fieldIndex, field := range segs {

    if fieldIndex == len(segs)-1 {
      return obj[field]
    }

    switch obj[field].(type) {
    case Map:
      obj = obj[field].(Map)
      case map[string]interface{}:
        obj = Map(obj[field].(map[string]interface{}))
    }

  }

  return obj

}