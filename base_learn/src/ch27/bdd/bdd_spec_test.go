package bdd

import (
	"testing"
)

func TestSpec(t *testing.T) {
	//Only pass t into top-level Convey calls
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
