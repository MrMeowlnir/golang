package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExampleCleanup(t *testing.T) {
	x := 0
	Convey("Given some integer with a starting value", t, func() {
		x++ // setup

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be 2", func() {
				So(x, ShouldEqual, 2)
			})
			Convey("The value should be 4", func() {
				So(x, ShouldEqual, 4)
			})
			Convey("The value should be 6", func() {
				So(x, ShouldNotEqual, 2)
			})
		})
	})
}
