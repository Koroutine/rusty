package rusty

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEither(t *testing.T) {

	Convey("Left should return first value", t, func() {
		value := ToEither(10, "test").Left()

		So(value, ShouldEqual, 10)
	})

	Convey("Right should return second value", t, func() {
		value := ToEither(10, "test").Right()

		So(value, ShouldEqual, "test")
	})

}
