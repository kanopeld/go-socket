package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewCaller(t *testing.T) {
	convey.Convey("test normal caller", t, func() {
		c, err := NewCaller(func(c Client) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 0)
		convey.So(c.NeedSocket, convey.ShouldBeTrue)
	})

	convey.Convey("test normal caller with args", t, func() {
		c, err := NewCaller(func(c Client, msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 1)
		convey.So(c.NeedSocket, convey.ShouldBeTrue)
	})

	convey.Convey("test normal caller without client ", t, func() {
		c, err := NewCaller(func(msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 1)
		convey.So(c.NeedSocket, convey.ShouldBeFalse)
	})

	convey.Convey("test normal caller call", t, func() {
		c, err := NewCaller(func(c Client) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 0)
		convey.So(c.NeedSocket, convey.ShouldBeTrue)

		retV := c.Call(&FakeClient{}, nil)
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call with msg", t, func() {
		c, err := NewCaller(func(c Client, msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 1)
		convey.So(c.NeedSocket, convey.ShouldBeTrue)

		retV := c.Call(&FakeClient{}, []byte{0x00})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client", t, func() {
		c, err := NewCaller(func(msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 1)
		convey.So(c.NeedSocket, convey.ShouldBeFalse)

		retV := c.Call(&FakeClient{}, []byte{0x00})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})
}