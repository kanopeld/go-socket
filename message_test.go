package socket

import (
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMessage_MarshalBinary(t *testing.T) {
	convey.Convey("test encode/decode message", t, func() {
		msg := Message{
			EventName: "test",
			Data:      []byte("hello test"),
		}

		b := msg.MarshalBinary()
		m := DecodeMessage(b)
		convey.So(m.EventName, convey.ShouldNotBeEmpty)
		assert.So(m.EventName, convey.ShouldEqual, "test")
		convey.So(m.Data, convey.ShouldNotBeEmpty)
		assert.So(string(m.Data), convey.ShouldEqual, "hello test")
	})
}
