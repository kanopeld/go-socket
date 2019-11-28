package socket

import (
	"fmt"
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
		c, err := NewCaller(func(msg []byte) {
			fmt.Println(string(msg))
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 1)
		convey.So(c.NeedSocket, convey.ShouldBeFalse)

		retV := c.Call(&FakeClient{}, []byte("hello bytes"))
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client", t, func() {
		c, err := NewCaller(func(msg string) {
			fmt.Println(msg)
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 1)
		convey.So(c.NeedSocket, convey.ShouldBeFalse)

		retV := c.Call(&FakeClient{}, []byte("hello string"))
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client. empty callback", t, func() {
		c, err := NewCaller(func() {
			fmt.Println("empty")
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Args, convey.ShouldHaveLength, 0)
		convey.So(c.NeedSocket, convey.ShouldBeFalse)

		retV := c.Call(&FakeClient{}, nil)
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test error too many args in callback", t, func() {
		c, err := NewCaller(func(c Client, msg string, tested string) {})
		convey.So(err, convey.ShouldEqual, ErrorTooManyArgumnetsForCaller)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test call caller with bytes slice", t, func() {
		c, err := NewCaller(func(msg []byte) {})
		convey.So(err, convey.ShouldBeNil)

		retV := c.Call(&FakeClient{}, []byte{0x0})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test call caller with string", t, func() {
		c, err := NewCaller(func(msg string) {})
		convey.So(err, convey.ShouldBeNil)

		retV := c.Call(&FakeClient{}, []byte{0x0})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test incorrect arg type uint", t, func() {
		c, err := NewCaller(func(c Client, id uint) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type int", t, func() {
		c, err := NewCaller(func(c Client, id int) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type byte", t, func() {
		c, err := NewCaller(func(c Client, id byte) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type interface (not Client)", t, func() {
		c, err := NewCaller(func(c Client, id interface{}) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type uintptr", t, func() {
		c, err := NewCaller(func(c Client, id uintptr) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type float", t, func() {
		c, err := NewCaller(func(c Client, id float64) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})
}
