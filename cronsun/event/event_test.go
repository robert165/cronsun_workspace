package event

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEvent(t *testing.T) {
	i := []int{}
	f := func(s interface{}) {
		i = append(i, 1)
	}
	f2 := func(s interface{}) {
		i = append(i, 2)
		i = append(i, 3)
	}

	Convey("events package test", t, func() {
		Convey("init events package should be success", func() {
			So(len(i), ShouldEqual, 0)
			So(len(Events[EXIT]), ShouldEqual, 0)
		})

		Convey("empty events execute Off function should not be success", func() {
			So(Off(EXIT, f), ShouldNotBeNil)
		})

		Convey("multi execute On function for a function should not be success", func() {
			So(On(EXIT, f), ShouldBeNil)
			So(On(EXIT, f), ShouldNotBeNil)
		})

		Convey("execute Emit function should be success", func() {
			Emit(EXIT, nil)
			So(len(i), ShouldEqual, 1)
		})

		Convey("events package should be work", func() {
			So(On(EXIT, f2), ShouldBeNil)
			So(len(Events[EXIT]), ShouldEqual, 2)
			So(len(i), ShouldEqual, 1)

			So(Off(EXIT, f), ShouldBeNil)
			So(len(Events[EXIT]), ShouldEqual, 1)
		})
	})
}
