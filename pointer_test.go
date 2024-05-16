package rusty

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPointer(t *testing.T) {
	some := 10

	Convey("Convert value to pointer", t, func() {
		value := ToPointer(some)

		So(*value, ShouldEqual, 10)
	})

	Convey("Convert pointer to value and access with ToOption", t, func() {
		value := ToPointer(some)
		option := ToOption(value)

		So(option.Unwrap(), ShouldEqual, 10)
	})
}
