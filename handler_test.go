package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewHandler(t *testing.T) {
	convey.Convey("testing get new handler", t, func() {
		h := newHandler(nil)
		convey.So(h, convey.ShouldNotBeNil)
	})
}

func TestBaseHandler_On(t *testing.T) {
	convey.Convey("testing On action", t, func() {
		h := newHandler(nil)
		convey.So(h, convey.ShouldNotBeNil)
		h.On("test", func(c Client, data []byte) error {
			return nil
		})
	})
}

func TestBaseHandler_Off(t *testing.T) {
	convey.Convey("testing OFF action", t, func() {
		h := newHandler(nil)
		h.On("test", func(c Client, data []byte) error {
			return nil
		})
		ok := h.Off("test")
		convey.So(ok, convey.ShouldBeTrue)
	})
}
