package socket

import (
	"errors"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	testError = errors.New("test error")
)

func TestClientHandler_Call(t *testing.T) {
	convey.Convey("testing call client handler", t, func() {
		ch := &clientHandler{client: &fakeServerClient{}, baseHandler: newHandler(nil, getCaller(""))}
		err := ch.On("test", func() {})
		convey.So(err, convey.ShouldBeNil)
		err = ch.call("test", nil)
		convey.So(err, convey.ShouldBeNil)
	})

	convey.Convey("testing call client handler. room not exist", t, func() {
		ch := &clientHandler{client: &fakeServerClient{}, baseHandler: newHandler(nil, getCaller(""))}
		err := ch.call("test", nil)
		convey.So(err, convey.ShouldEqual, ErrEventNotExist)
	})

	convey.Convey("testing call client handler", t, func() {
		ch := &clientHandler{client: &fakeServerClient{}, baseHandler: newHandler(nil, getCaller(""))}
		err := ch.On("test", func(i string) error {
			return testError
		})
		convey.So(err, convey.ShouldBeNil)
		err = ch.call("test", nil)
		convey.So(err, convey.ShouldEqual, testError)
	})
}
