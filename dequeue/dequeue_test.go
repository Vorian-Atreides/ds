package dequeue

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestBorderCases(t *testing.T) {

	convey.Convey("With an empty dequeue", t, func() {
		d := New()

		convey.Convey("Should be empty", func() {
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Front should be nil", func() {
			convey.So(d.Front(), convey.ShouldBeNil)
		})

		convey.Convey("Back should be nil", func() {
			convey.So(d.Back(), convey.ShouldBeNil)
		})

		convey.Convey("PopFront should be nil", func() {
			convey.So(d.PopFront(), convey.ShouldBeNil)
		})

		convey.Convey("PopBack should be nil", func() {
			convey.So(d.PopBack(), convey.ShouldBeNil)
		})

		convey.Convey("Remove with top should be nil", func() {
			convey.So(d.remove(d.top), convey.ShouldBeNil)
		})

		convey.Convey("Remove with bottom should be nil", func() {
			convey.So(d.remove(d.bottom), convey.ShouldBeNil)
		})

		convey.Convey("Should be able to insert in top", func() {
			d.insertFront(new(Element))
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Back(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to insert in bottom", func() {
			d.insertBack(new(Element))
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Back(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to PushFront", func() {
			d.PushFront(14)
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Front().Value, convey.ShouldEqual, 14)
			convey.So(d.top, convey.ShouldEqual, d.bottom)
		})
	})
} // With an empty Dequeue

func TestNormalCase(t *testing.T) {

	convey.Convey("With one item (10) in the Dequeue", t, func() {
		d := New()
		d.PushFront(10)

		convey.Convey("Should not be empty", func() {
			convey.So(d.IsEmpty(), convey.ShouldBeFalse)
		})

		convey.Convey("Front should be 10", func() {
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Front().Value, convey.ShouldEqual, 10)
		})

		convey.Convey("Back should be 10", func() {
			convey.So(d.Back(), convey.ShouldNotBeNil)
			convey.So(d.Back().Value, convey.ShouldEqual, 10)
		})

		convey.Convey("PopFront should not be nil", func() {
			convey.So(d.PopFront(), convey.ShouldNotBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("PopBack should not be nil", func() {
			convey.So(d.PopBack(), convey.ShouldNotBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("PopFront should be 10", func() {
			convey.So(d.PopFront().Value, convey.ShouldEqual, 10)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("PopBack should be 10", func() {
			convey.So(d.PopBack().Value, convey.ShouldEqual, 10)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Remove from top should not be nil", func() {
			convey.So(d.remove(d.top), convey.ShouldNotBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Remove from bottom should not be nil", func() {
			convey.So(d.remove(d.bottom), convey.ShouldNotBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Remove from top should be 10", func() {
			convey.So(d.remove(d.top).Value, convey.ShouldEqual, 10)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Remove from bottom should be 10", func() {
			convey.So(d.remove(d.bottom).Value, convey.ShouldEqual, 10)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
			convey.So(d.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Should be able to insert in top", func() {
			d.insertFront(new(Element))
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Back(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to insert in bottom", func() {
			d.insertBack(new(Element))
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Back(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to PushFront (14)", func() {
			d.PushFront(14)
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Front().Value, convey.ShouldEqual, 14)
			convey.So(d.Back().Value, convey.ShouldEqual, 10)
		})

		convey.Convey("Should be able to PushBack (14)", func() {
			d.PushBack(14)
			convey.So(d.Back(), convey.ShouldNotBeNil)
			convey.So(d.Back().Value, convey.ShouldEqual, 14)
			convey.So(d.Front().Value, convey.ShouldEqual, 10)
		})

	}) // With one item (10) in the Dequeue

	convey.Convey("With three items (1, 2, 3) in the Dequeue", t, func() {
		d := New()
		d.PushFront(1)
		d.PushFront(2)
		d.PushFront(3)

		convey.Convey("Should not be empty", func() {
			convey.So(d.IsEmpty(), convey.ShouldBeFalse)
		})

		convey.Convey("Should be able to PopFront", func() {
			convey.So(d.PopFront(), convey.ShouldNotBeNil)
			convey.So(d.PopFront(), convey.ShouldNotBeNil)
			convey.So(d.PopFront(), convey.ShouldNotBeNil)
			convey.So(d.PopFront(), convey.ShouldBeNil)
		})

		convey.Convey("Should be able to PopBack", func() {
			convey.So(d.PopBack(), convey.ShouldNotBeNil)
			convey.So(d.PopBack(), convey.ShouldNotBeNil)
			convey.So(d.PopBack(), convey.ShouldNotBeNil)
			convey.So(d.PopBack(), convey.ShouldBeNil)
		})

		convey.Convey("PopFront should have the values 3, 2, 1", func() {
			convey.So(d.PopFront().Value, convey.ShouldEqual, 3)
			convey.So(d.PopFront().Value, convey.ShouldEqual, 2)
			convey.So(d.PopFront().Value, convey.ShouldEqual, 1)
			convey.So(d.PopFront(), convey.ShouldBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
		})

		convey.Convey("PopBack should have the values 1, 2, 3", func() {
			convey.So(d.PopBack().Value, convey.ShouldEqual, 1)
			convey.So(d.PopBack().Value, convey.ShouldEqual, 2)
			convey.So(d.PopBack().Value, convey.ShouldEqual, 3)
			convey.So(d.PopBack(), convey.ShouldBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
		})

		convey.Convey("Should be able to Remove from top", func() {
			convey.So(d.remove(d.top), convey.ShouldNotBeNil)
			convey.So(d.remove(d.top), convey.ShouldNotBeNil)
			convey.So(d.remove(d.top), convey.ShouldNotBeNil)
			convey.So(d.remove(d.top), convey.ShouldBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
		})

		convey.Convey("Should be able to Remove from bottom", func() {
			convey.So(d.remove(d.bottom), convey.ShouldNotBeNil)
			convey.So(d.remove(d.bottom), convey.ShouldNotBeNil)
			convey.So(d.remove(d.bottom), convey.ShouldNotBeNil)
			convey.So(d.remove(d.bottom), convey.ShouldBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
		})

		convey.Convey("Remove from top should have the values 3, 2, 1", func() {
			convey.So(d.remove(d.top).Value, convey.ShouldEqual, 3)
			convey.So(d.remove(d.top).Value, convey.ShouldEqual, 2)
			convey.So(d.remove(d.top).Value, convey.ShouldEqual, 1)
			convey.So(d.remove(d.top), convey.ShouldBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
		})

		convey.Convey("Remove from top should have the values 1, 2, 3", func() {
			convey.So(d.remove(d.bottom).Value, convey.ShouldEqual, 1)
			convey.So(d.remove(d.bottom).Value, convey.ShouldEqual, 2)
			convey.So(d.remove(d.bottom).Value, convey.ShouldEqual, 3)
			convey.So(d.remove(d.bottom), convey.ShouldBeNil)
			convey.So(d.bottom, convey.ShouldBeNil)
			convey.So(d.top, convey.ShouldBeNil)
		})

		convey.Convey("Should be able to insert from top", func() {
			d.insertFront(new(Element))
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Front().Value, convey.ShouldNotEqual, 3)
			convey.So(d.bottom, convey.ShouldNotBeNil)
			convey.So(d.top, convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to insert from bottom", func() {
			d.insertBack(new(Element))
			convey.So(d.Back(), convey.ShouldNotBeNil)
			convey.So(d.Back().Value, convey.ShouldNotEqual, 1)
			convey.So(d.bottom, convey.ShouldNotBeNil)
			convey.So(d.top, convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to PushFront (14)", func() {
			d.PushFront(14)
			convey.So(d.Front(), convey.ShouldNotBeNil)
			convey.So(d.Front().Value, convey.ShouldEqual, 14)
			convey.So(d.Back().Value, convey.ShouldEqual, 1)
		})

		convey.Convey("Should be able to PushBack (14)", func() {
			d.PushBack(14)
			convey.So(d.Back(), convey.ShouldNotBeNil)
			convey.So(d.Back().Value, convey.ShouldEqual, 14)
			convey.So(d.Front().Value, convey.ShouldEqual, 3)
		})

	}) // With three items (1, 2, 3) in the Dequeue
}
