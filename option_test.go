package rusty

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOption(t *testing.T) {

	some := 10
	var none *int
	none = nil

	Convey("Unwrap returns value if Some", t, func() {
		value := ToOption(&some).Unwrap()

		So(value, ShouldEqual, 10)
	})

	Convey("Unwrap panics if value None", t, func() {
		So(func() { ToOption(none).Unwrap() }, ShouldPanic)
	})

	Convey("Expect panics with message if value None", t, func() {
		So(func() { ToOption(none).Expect("woops") }, ShouldPanic)
	})

	Convey("UnwrapOr returns fallback value if None", t, func() {
		value := ToOption(none).UnwrapOr(0)

		So(value, ShouldEqual, 0)
	})

	Convey("UnwrapOrElse calls fallback function if None", t, func() {
		value := ToOption(none).UnwrapOrElse(func() int { return 1 })

		So(value, ShouldEqual, 1)
	})

	Convey("IsSome returns true if Some", t, func() {
		value := ToOption(&some).IsSome()

		So(value, ShouldEqual, true)
	})

	Convey("IsNone returns true if None", t, func() {
		value := ToOption(none).IsNone()

		So(value, ShouldEqual, true)
	})

	Convey("None IsNone", t, func() {
		value := ToNone[int]().IsNone()

		So(value, ShouldEqual, true)
	})
}
