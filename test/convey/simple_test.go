package convey

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// convey 测试框架
func TestSpec(t *testing.T) {
	Convey("Start ", t, func() {
		x := 1
		Convey("when the var incremented", func() {
			x++
			Convey("first output", func() {
				So(x, ShouldEqual, 2)
			})
		})

		Convey("when the another var incremented", func() {
			x++
			Convey("seoncd output", func() {
				So(x, ShouldEqual, 3)
			})
		})
	})
}
