package stack

import (
	"testing"

	"github.com/Vorian-Atreides/ds/src/stack"
	"github.com/smartystreets/goconvey/convey"
)

func TestBorderCases(t *testing.T) {

	convey.Convey("With an empty Stack", t, func() {
		s := stack.New()

		convey.Convey("Should be empty", func() {
			So(s.IsEmpty(), ShouldEqual, true)
		})
	})
}
