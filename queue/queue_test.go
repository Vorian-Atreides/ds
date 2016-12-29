package queue

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestBorderCases(t *testing.T) {

	convey.Convey("With an empty Queue", t, func() {
		q := New()

		convey.Convey("Should be empty", func() {
			convey.So(q.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Peek should be nil", func() {
			convey.So(q.Peek(), convey.ShouldBeNil)
		})

		convey.Convey("Dequeue should be nil", func() {
			convey.So(q.Dequeue(), convey.ShouldBeNil)
		})

		convey.Convey("Remove should be nil", func() {
			convey.So(q.remove(), convey.ShouldBeNil)
		})

		convey.Convey("Should be able to insert", func() {
			q.insert(new(Element))
			convey.So(q.Peek(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to enqueue", func() {
			q.Enqueue(14)
			convey.So(q.Peek(), convey.ShouldNotBeNil)
			convey.So(q.Peek().Value, convey.ShouldEqual, 14)
			convey.So(q.top, convey.ShouldEqual, q.bottom)
		})
	})
} // With an empty Queue

func TestNormalCase(t *testing.T) {

	convey.Convey("With one item (10) in the Queue", t, func() {
		q := New()
		q.Enqueue(10)

		convey.Convey("Should not be empty", func() {
			convey.So(q.IsEmpty(), convey.ShouldBeFalse)
		})

		convey.Convey("Peek should be 10", func() {
			convey.So(q.Peek(), convey.ShouldNotBeNil)
			convey.So(q.Peek().Value, convey.ShouldEqual, 10)
		})

		convey.Convey("Dequeue should not be nil", func() {
			convey.So(q.Dequeue(), convey.ShouldNotBeNil)
			convey.So(q.bottom, convey.ShouldBeNil)
			convey.So(q.top, convey.ShouldBeNil)
			convey.So(q.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Dequeue should be 10", func() {
			convey.So(q.Dequeue().Value, convey.ShouldEqual, 10)
			convey.So(q.bottom, convey.ShouldBeNil)
			convey.So(q.top, convey.ShouldBeNil)
			convey.So(q.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Remove should not be nil", func() {
			convey.So(q.remove(), convey.ShouldNotBeNil)
			convey.So(q.bottom, convey.ShouldBeNil)
			convey.So(q.top, convey.ShouldBeNil)
			convey.So(q.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Remove should be 10", func() {
			convey.So(q.remove().Value, convey.ShouldEqual, 10)
			convey.So(q.bottom, convey.ShouldBeNil)
			convey.So(q.top, convey.ShouldBeNil)
			convey.So(q.IsEmpty(), convey.ShouldBeTrue)
		})

		convey.Convey("Should be able to insert", func() {
			q.insert(new(Element))
			convey.So(q.Peek(), convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to enqueue (14)", func() {
			q.Enqueue(14)
			convey.So(q.Peek(), convey.ShouldNotBeNil)
			convey.So(q.Peek().Value, convey.ShouldEqual, 10)
			convey.So(q.bottom.Value, convey.ShouldEqual, 10)
			convey.So(q.top.Value, convey.ShouldEqual, 14)
		})

	}) // With one item (10) in the Queue

	convey.Convey("With three items (1, 2, 3) in the Queue", t, func() {
		q := New()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		convey.Convey("Should not be empty", func() {
			convey.So(q.IsEmpty(), convey.ShouldBeFalse)
		})

		convey.Convey("Should be able to Dequeue", func() {
			convey.So(q.Dequeue(), convey.ShouldNotBeNil)
			convey.So(q.Dequeue(), convey.ShouldNotBeNil)
			convey.So(q.Dequeue(), convey.ShouldNotBeNil)
			convey.So(q.Dequeue(), convey.ShouldBeNil)
		})

		convey.Convey("Dequeue should have the values 1, 2, 3", func() {
			convey.So(q.Dequeue().Value, convey.ShouldEqual, 1)
			convey.So(q.Dequeue().Value, convey.ShouldEqual, 2)
			convey.So(q.Dequeue().Value, convey.ShouldEqual, 3)
			convey.So(q.Dequeue(), convey.ShouldBeNil)
			convey.So(q.bottom, convey.ShouldBeNil)
			convey.So(q.top, convey.ShouldBeNil)
		})

		convey.Convey("Should be able to Remove", func() {
			convey.So(q.remove(), convey.ShouldNotBeNil)
			convey.So(q.remove(), convey.ShouldNotBeNil)
			convey.So(q.remove(), convey.ShouldNotBeNil)
			convey.So(q.remove(), convey.ShouldBeNil)
			convey.So(q.bottom, convey.ShouldBeNil)
			convey.So(q.top, convey.ShouldBeNil)
		})

		convey.Convey("Remove should have the values 1, 2, 3", func() {
			convey.So(q.remove().Value, convey.ShouldEqual, 1)
			convey.So(q.remove().Value, convey.ShouldEqual, 2)
			convey.So(q.remove().Value, convey.ShouldEqual, 3)
			convey.So(q.remove(), convey.ShouldBeNil)
			convey.So(q.bottom, convey.ShouldBeNil)
			convey.So(q.top, convey.ShouldBeNil)
		})

		convey.Convey("Should be able to enqueue", func() {
			q.insert(new(Element))
			convey.So(q.Peek(), convey.ShouldNotBeNil)
			convey.So(q.bottom, convey.ShouldNotBeNil)
			convey.So(q.top, convey.ShouldNotBeNil)
		})

		convey.Convey("Should be able to enqueue (14)", func() {
			q.Enqueue(14)
			convey.So(q.Peek(), convey.ShouldNotBeNil)
			convey.So(q.Peek().Value, convey.ShouldEqual, 1)
			convey.So(q.top.Value, convey.ShouldEqual, 14)
		})

	}) // With three items (1, 2, 3) in the Queue
}
