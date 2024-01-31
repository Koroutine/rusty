package rusty

import (
	"encoding/json"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSet(t *testing.T) {
	Convey("Set value at existing path", t, func() {
		target := map[string]interface{}{
			"a": 1,
			"b": map[string]interface{}{
				"c": "2",
				"d": map[string]interface{}{
					"e": 3,
					"f": true,
					"g": "hello",
				},
			},
		}

		Set(target, "b.d.g", "test").Unwrap()

		updated := Get[string](target, "b.d.g").Unwrap()

		So(updated, ShouldEqual, "test")

		snaps.MatchJSON(t, target)
	})

	Convey("Set value at new path", t, func() {
		target := map[string]interface{}{
			"a": 1,
			"b": map[string]interface{}{
				"c": "2",
				"d": map[string]interface{}{
					"e": 3,
					"f": true,
					"g": "hello",
				},
			},
		}

		Set(target, "b.d.g.r", "test").Unwrap()

		updated := Get[string](target, "b.d.g.r").Unwrap()

		So(updated, ShouldEqual, "test")

		snaps.MatchJSON(t, target)
	})

	Convey("Set value on object inside array", t, func() {
		target := map[string]interface{}{
			"a": 1,
			"arr": []interface{}{
				map[string]interface{}{
					"key": "value",
				},
			},
		}

		Set(target, "arr.0.key", "test").Unwrap()

		updated := Get[string](target, "arr.0.key").Unwrap()

		So(updated, ShouldEqual, "test")

		snaps.MatchJSON(t, target)
	})

	Convey("Set realworld", t, func() {
		// String JSON multiline
		input := `{
			"Records": [
				{
					"Record_details": {
						"id": "test"
					},
					"Outcome_details": [
						{
							"id": "test"
						}
					]
				}
			]
		}`

		var inputMap Map
		err := json.Unmarshal([]byte(input), &inputMap)
		So(err, ShouldBeNil)

		Set(inputMap, "Records.0.Record_details.obj.a", "test1").Unwrap()
		Set(inputMap, "Records.0.Record_details.obj.b", "test2").Unwrap()

		Set(inputMap, "Records.0.Outcome_details.0.obj.a", "test3").Unwrap()
		Set(inputMap, "Records.0.Outcome_details.0.obj.b", "test4").Unwrap()

		updated := Get[string](inputMap, "Records.0.Record_details.obj.a").Unwrap()

		So(updated, ShouldEqual, "test1")

		snaps.MatchJSON(t, inputMap)
	})
}

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

	arrayTest := Map{
		"a": []string{
			"1",
		},
	}

	Convey("Find value at array path", t, func() {
		value := Get[string](arrayTest, "a.0").Unwrap()

		So(value, ShouldEqual, "1")
	})

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
