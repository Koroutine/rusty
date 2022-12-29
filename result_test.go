package rusty

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResult(t *testing.T) {
	Convey("Unwrap returns value if Ok", t, func() {
		value := ToResult(strconv.Atoi("10")).Unwrap()

		So(value, ShouldEqual, 10)
	})

	Convey("Unwrap panics if value Err", t, func() {
		So(func() { ToResult(strconv.Atoi("test")).Unwrap() }, ShouldPanic)
	})

	Convey("Expect panics with message if value Err", t, func() {
		So(func() { ToResult(strconv.Atoi("test")).Expect("woops") }, ShouldPanic)
	})

	Convey("UnwrapOr returns fallback value if Err", t, func() {
		value := ToResult(strconv.Atoi("test")).UnwrapOr(0)

		So(value, ShouldEqual, 0)
	})

	Convey("UnwrapOrElse calls fallback function if Err", t, func() {
		value := ToResult(strconv.Atoi("test")).UnwrapOrElse(func() int { return 1 })

		So(value, ShouldEqual, 1)
	})

	Convey("MapOption convers Option to another value", t, func() {
		res := ToResult(strconv.Atoi("10"))
		value := MapResult(res, func(v int) string { return fmt.Sprint(v) }).Unwrap()

		So(value, ShouldEqual, "10")
	})

	Convey("MapOption with UnwrapOr returns fallback value if None", t, func() {
		res := ToResult(strconv.Atoi("-"))
		value := MapResult(res, func(v int) string { return fmt.Sprint(v) }).UnwrapOr("1")

		So(value, ShouldEqual, "1")
	})

	Convey("IsOk returns true if Ok", t, func() {
		value := ToResult(strconv.Atoi("10")).IsOk()

		So(value, ShouldEqual, true)
	})

	Convey("Some is Ok", t, func() {
		value := ToOk(10).IsOk()

		So(value, ShouldEqual, true)
	})

	Convey("IsErr returns true if Err", t, func() {
		value := ToResult(strconv.Atoi("test")).IsErr()

		So(value, ShouldEqual, true)
	})

	Convey("Error is Err", t, func() {
		value := ToError(fmt.Errorf("Error")).IsErr()

		So(value, ShouldEqual, true)
	})

	Convey("Error is Err with check", t, func() {
		value := ToError(fmt.Errorf("Error")).IsErrWith(func(err error) bool { return err.Error() == "Error" })

		So(value, ShouldEqual, true)
	})

	Convey("Error is Err with bad check", t, func() {
		value := ToError(fmt.Errorf("test")).IsErrWith(func(err error) bool { return err.Error() == "Error" })

		So(value, ShouldEqual, false)
	})

	Convey("Ok returns result as Option", t, func() {
		value := ToResult(strconv.Atoi("10")).Ok().Unwrap()

		So(value, ShouldEqual, 10)
	})

	Convey("Ok returns result as Option - handles bad result", t, func() {
		So(func() { ToResult(strconv.Atoi("test")).Ok().Expect("woops") }, ShouldPanic)
	})
}
