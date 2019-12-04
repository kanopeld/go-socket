package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewHandler(t *testing.T) {
	convey.Convey("testing get new handler", t, func() {
		h := NewHandler(nil, func(f interface{}) (c caller, err error) {
			return &call{}, nil
		})
		convey.So(h, convey.ShouldNotBeNil)
	})
}

func TestBaseHandler_On(t *testing.T) {
	convey.Convey("testing On action", t, func() {
		h := NewHandler(nil, func(f interface{}) (c caller, err error) {
			return &call{}, nil
		})
		convey.So(h, convey.ShouldNotBeNil)

		err := h.On("test", func() {})
		convey.So(err, convey.ShouldBeNil)
	})

	convey.Convey("testing ON action with error", t, func() {
		h := NewHandler(nil, getCaller("test"))
		convey.So(h, convey.ShouldNotBeNil)

		err := h.On("test", func(i int) {})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
	})
}

func TestBaseHandler_Off(t *testing.T) {
	convey.Convey("testing OFF action", t, func() {
		h := NewHandler(nil, func(f interface{}) (c caller, err error) {
			return &call{}, nil
		})

		err := h.On("test", func() {})
		convey.So(err, convey.ShouldBeNil)
		ok := h.Off("test")
		convey.So(ok, convey.ShouldBeTrue)
	})
}

func TestBaseHandler_GetEvents(t *testing.T) {
	convey.Convey("testing handler events getting", t, func() {
		h := NewHandler(nil, getCaller("test"))
		convey.So(h, convey.ShouldNotBeNil)
		convey.So(h.events, convey.ShouldHaveLength, 0)
	})
}

func TestBaseHandler_GetBroadcast(t *testing.T) {
	convey.Convey("testing get broadcast", t, func() {
		h := NewHandler(nil, getCaller("test"))
		convey.So(h, convey.ShouldNotBeNil)
		b := h.GetBroadcast()
		convey.So(b, convey.ShouldBeNil)
	})
}
