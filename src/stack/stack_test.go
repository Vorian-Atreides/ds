package stack

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestBorderCases(t *testing.T) {

	convey.Convey("With an empty Stack", t, func() {
		s := New()

		convey.Convey("Should be empty", func() {
			convey.So(s.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Peek should be nil", func() {
			convey.So(s.Peek(), convey.ShouldBeNil)
		})

		convey.Convey("Pop should be nil", func() {
			convey.So(s.Pop(), convey.ShouldBeNil)
		})

		convey.Convey("Remove should be nil", func() {
			convey.So(s.remove(), convey.ShouldBeNil)
		})

		convey.Convey("Should be able to insert", func() {
			s.insert(new(Element))
			convey.So(s.Peek(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to push", func() {
			s.Push(14)
			convey.So(s.Peek(), convey.ShouldNotBeNil)
			convey.So(s.Peek().Value, convey.ShouldEqual, 14)
		})

	}) // With an empty Stack
}

func TestNormalCase(t *testing.T) {

	convey.Convey("With one item (10) in the Stack", t, func() {
		s := New()
		s.Push(10)

		convey.Convey("Should not be empty", func() {
			convey.So(s.IsEmpty(), convey.ShouldBeFalse)
		})

		convey.Convey("Peek should be 10", func() {
			convey.So(s.Peek(), convey.ShouldNotBeNil)
			convey.So(s.Peek().Value, convey.ShouldEqual, 10)
		})

		convey.Convey("Pop should not be nil", func() {
			convey.So(s.Pop(), convey.ShouldNotBeNil)
			convey.So(s.top, convey.ShouldBeNil)
			convey.So(s.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Pop should be 10", func() {
			convey.So(s.Pop().Value, convey.ShouldEqual, 10)
			convey.So(s.top, convey.ShouldBeNil)
			convey.So(s.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Remove should not be nil", func() {
			convey.So(s.remove(), convey.ShouldNotBeNil)
			convey.So(s.top, convey.ShouldBeNil)
			convey.So(s.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Remove should be 10", func() {
			convey.So(s.remove().Value, convey.ShouldEqual, 10)
			convey.So(s.top, convey.ShouldBeNil)
			convey.So(s.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Should be able to insert", func() {
			s.insert(new(Element))
			convey.So(s.Peek(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to push (14)", func() {
			s.Push(14)
			convey.So(s.Peek(), convey.ShouldNotBeNil)
			convey.So(s.Peek().Value, convey.ShouldEqual, 14)
		})

	}) // With one item (10) in the Stack

	convey.Convey("With three items (1, 2, 3) in the Stack", t, func() {
		s := New()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		convey.Convey("Should not be empty", func() {
			convey.So(s.IsEmpty(), convey.ShouldBeFalse)
		})

		convey.Convey("Should be able to Pop", func() {
			convey.So(s.Pop(), convey.ShouldNotBeNil)
			convey.So(s.Pop(), convey.ShouldNotBeNil)
			convey.So(s.Pop(), convey.ShouldNotBeNil)
			convey.So(s.Pop(), convey.ShouldBeNil)
		})

		convey.Convey("Pop should have the values 3, 2, 1", func() {
			convey.So(s.Pop().Value, convey.ShouldEqual, 3)
			convey.So(s.Pop().Value, convey.ShouldEqual, 2)
			convey.So(s.Pop().Value, convey.ShouldEqual, 1)
			convey.So(s.Pop(), convey.ShouldBeNil)
		})

		convey.Convey("Should be able to Remove", func() {
			convey.So(s.remove(), convey.ShouldNotBeNil)
			convey.So(s.remove(), convey.ShouldNotBeNil)
			convey.So(s.remove(), convey.ShouldNotBeNil)
			convey.So(s.remove(), convey.ShouldBeNil)
		})

		convey.Convey("Remove should have the values 3, 2, 1", func() {
			convey.So(s.remove().Value, convey.ShouldEqual, 3)
			convey.So(s.remove().Value, convey.ShouldEqual, 2)
			convey.So(s.remove().Value, convey.ShouldEqual, 1)
			convey.So(s.remove(), convey.ShouldBeNil)
		})

		convey.Convey("Should be able to insert", func() {
			s.insert(new(Element))
			convey.So(s.Peek(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to push (14)", func() {
			s.Push(14)
			convey.So(s.Peek(), convey.ShouldNotBeNil)
			convey.So(s.Peek().Value, convey.ShouldEqual, 14)
		})

	}) // With three items (1, 2, 3) in the Stack
}
