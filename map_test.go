package rusty

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {

	target := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"c": "2",
			"d": map[string]interface{}{
				"e": 3,
				"f": true,
			},
		},
	}

	Convey("Find value at path", t, func() {
		value := Get[string](target, "a.b.c").Unwrap()

		So(value, ShouldEqual, "2")
	})

	Convey("Find int value at path", t, func() {
		value := Get[int](target, "a.b.c.d.e").Unwrap()

		So(value, ShouldEqual, 3)
	})

	Convey("Default value if path not found", t, func() {
		value := Get[string](target, "a.b.c.z").UnwrapOr("test")

		So(value, ShouldEqual, "test")
	})

	Convey("Unwrap panics if path not found", t, func() {
		So(func() { Get[string](target, "a.b.c.z").Unwrap() }, ShouldPanic)
	})

	Convey("Unwrap panics if type assertion fails", t, func() {
		So(func() { Get[int](target, "a.b.c").Unwrap() }, ShouldPanic)
	})

}
