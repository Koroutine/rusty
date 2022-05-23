package json

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJson(t *testing.T) {

	Convey("Should return Result[string] when Marshalling", t, func() {
		value := ToString(map[string]interface{}{
			"a": 1,
			"b": "test",
		}).Unwrap()

		So(value, ShouldEqual, "{\"a\":1,\"b\":\"test\"}")
	})

	Convey("Should handle error when Marshalling", t, func() {

		So(func() { ToString(func() {}).Unwrap() }, ShouldPanic)
	})

}
