package dial

import (
	"fmt"
	"github.com/kanopeld/go-socket/core"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewCaller(t *testing.T) {
	convey.Convey("test normal caller", t, func() {
		c, err := NewCaller(func(c core.SClient) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.GetArgs(), convey.ShouldHaveLength, 0)
		convey.So(c.Socket(), convey.ShouldBeTrue)
	})

	convey.Convey("test normal caller with args", t, func() {
		c, err := NewCaller(func(c core.SClient, msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.GetArgs(), convey.ShouldHaveLength, 1)
		convey.So(c.Socket(), convey.ShouldBeTrue)
	})

	convey.Convey("test normal caller without client ", t, func() {
		c, err := NewCaller(func(msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.GetArgs(), convey.ShouldHaveLength, 1)
		convey.So(c.Socket(), convey.ShouldBeFalse)
	})

	convey.Convey("test normal caller call", t, func() {
		c, err := NewCaller(func(c core.SClient) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.GetArgs(), convey.ShouldHaveLength, 0)
		convey.So(c.Socket(), convey.ShouldBeTrue)

		retV := c.Call(&core.FakeServerClient{}, nil)
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call with msg", t, func() {
		c, err := NewCaller(func(c core.SClient, msg []byte) {})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.GetArgs(), convey.ShouldHaveLength, 1)
		convey.So(c.Socket(), convey.ShouldBeTrue)

		retV := c.Call(&core.FakeServerClient{}, []byte{0x00})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client", t, func() {
		c, err := NewCaller(func(msg []byte) {
			fmt.Println(string(msg))
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.GetArgs(), convey.ShouldHaveLength, 1)
		convey.So(c.Socket(), convey.ShouldBeFalse)

		retV := c.Call(&core.FakeServerClient{}, []byte("hello bytes"))
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client", t, func() {
		c, err := NewCaller(func(msg string) {
			fmt.Println(msg)
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.GetArgs(), convey.ShouldHaveLength, 1)
		convey.So(c.Socket(), convey.ShouldBeFalse)

		retV := c.Call(&core.FakeServerClient{}, []byte("hello string"))
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test normal caller call without client. empty callback", t, func() {
		c, err := NewCaller(func() {
			fmt.Println("empty")
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.GetArgs(), convey.ShouldHaveLength, 0)
		convey.So(c.Socket(), convey.ShouldBeFalse)

		retV := c.Call(&core.FakeServerClient{}, nil)
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test error too many args in callback", t, func() {
		c, err := NewCaller(func(c core.SClient, msg string, tested string) {})
		convey.So(err, convey.ShouldEqual, ErrTooManyArgsForCaller)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test call caller with bytes slice", t, func() {
		c, err := NewCaller(func(msg []byte) {})
		convey.So(err, convey.ShouldBeNil)

		retV := c.Call(&core.FakeServerClient{}, []byte{0x0})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test call caller with string", t, func() {
		c, err := NewCaller(func(msg string) {})
		convey.So(err, convey.ShouldBeNil)

		retV := c.Call(&core.FakeServerClient{}, []byte{0x0})
		convey.So(retV, convey.ShouldHaveLength, 0)
	})

	convey.Convey("test incorrect arg type uint", t, func() {
		c, err := NewCaller(func(c core.SClient, id uint) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type int", t, func() {
		c, err := NewCaller(func(c core.SClient, id int) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type byte", t, func() {
		c, err := NewCaller(func(c core.SClient, id byte) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type interface (not Client)", t, func() {
		c, err := NewCaller(func(c core.SClient, id interface{}) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type uintptr", t, func() {
		c, err := NewCaller(func(c core.SClient, id uintptr) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})

	convey.Convey("test incorrect arg type float", t, func() {
		c, err := NewCaller(func(c core.SClient, id float64) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
		convey.So(c, convey.ShouldBeNil)
	})
}
