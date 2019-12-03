package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewCaller(t *testing.T) {
	convey.Convey("test normal caller", t, func() {
		c, err := GetCaller("SClient")(123)
		convey.So(err, convey.ShouldEqual, ErrFIsNotFunc)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test normal caller with args", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.argsLen(), convey.ShouldEqual, 1)
		convey.So(c.socket(), convey.ShouldBeTrue)
	})

	convey.Convey("test normal caller without client ", t, func() {
		c, err := GetCaller("SClient")(func(msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.argsLen(), convey.ShouldEqual, 1)
		convey.So(c.socket(), convey.ShouldBeFalse)
	})

	convey.Convey("test normal caller call", t, func() {
		c, err := GetCaller("SClient")(func(c SClient) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.argsLen(), convey.ShouldEqual, 0)
		convey.So(c.socket(), convey.ShouldBeTrue)

		retV := c.call(&FakeServerClient{}, nil)
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call with msg", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.argsLen(), convey.ShouldEqual, 1)
		convey.So(c.socket(), convey.ShouldBeTrue)

		retV := c.call(&FakeServerClient{}, []byte{0x00})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client", t, func() {
		c, err := GetCaller("SClient")(func(msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.argsLen(), convey.ShouldEqual, 1)
		convey.So(c.socket(), convey.ShouldBeFalse)

		retV := c.call(&FakeServerClient{}, []byte("hello bytes"))
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client", t, func() {
		c, err := GetCaller("SClient")(func(msg string) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.argsLen(), convey.ShouldEqual, 1)
		convey.So(c.socket(), convey.ShouldBeFalse)

		retV := c.call(&FakeServerClient{}, []byte("hello string"))
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client. empty callback", t, func() {
		c, err := GetCaller("SClient")(func() {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.argsLen(), convey.ShouldEqual, 0)
		convey.So(c.socket(), convey.ShouldBeFalse)

		retV := c.call(&FakeServerClient{}, nil)
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test error too many args in callback", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, msg string, tested string) {})
		convey.So(err, convey.ShouldEqual, ErrTooManyArgsForCaller)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test call caller with bytes slice", t, func() {
		c, err := GetCaller("SClient")(func(msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		retV := c.call(&FakeServerClient{}, []byte{0x0})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test call caller with string", t, func() {
		c, err := GetCaller("SClient")(func(msg string) {})
		convey.So(err, convey.ShouldBeNil)

		retV := c.call(&FakeServerClient{}, []byte{0x0})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test incorrect arg type uint", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, id uint) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type int", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, id int) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type byte", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, id byte) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type interface (not Client)", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, id interface{}) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type uintptr", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, id uintptr) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type float", t, func() {
		c, err := GetCaller("SClient")(func(c SClient, id float64) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})
}
