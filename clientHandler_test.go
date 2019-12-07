package socket

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestClientHandler_Call(t *testing.T) {
	convey.Convey("testing call client handler", t, func() {
		ch := &clientHandler{client: &fakeServerClient{}, baseHandler: newHandler(nil)}
		ch.On("test", func(c Client, data []byte) error {
			return nil
		})
		err := ch.call("test", nil)
		convey.So(err, convey.ShouldBeNil)
	})

	convey.Convey("testing call client handler. room not exist", t, func() {
		ch := &clientHandler{client: &fakeServerClient{}, baseHandler: newHandler(nil)}
		err := ch.call("test", nil)
		convey.So(err, convey.ShouldEqual, ErrEventNotExist)
	})
}
