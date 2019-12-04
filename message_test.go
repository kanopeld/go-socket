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
		m := decodeMessage(b)
		convey.So(m.EventName, convey.ShouldNotBeEmpty)
		assert.So(m.EventName, convey.ShouldEqual, "test")
		convey.So(m.Data, convey.ShouldNotBeEmpty)
		assert.So(string(m.Data), convey.ShouldEqual, "hello test")
	})

	convey.Convey("test decode bytes msg data", t, func() {
		msg := Message{
			EventName: "test",
			Data:      []byte{0x1, 0x2, 0x03},
		}

		b := msg.MarshalBinary()
		m := decodeMessage(b)
		convey.So(m.EventName, convey.ShouldNotBeEmpty)
		assert.So(m.EventName, convey.ShouldEqual, "test")
		convey.So(m.Data, convey.ShouldNotBeEmpty)
		convey.So(m.Data, convey.ShouldHaveLength, 3)
		convey.So(m.Data[0], convey.ShouldEqual, 0x1)
		convey.So(m.Data[1], convey.ShouldEqual, 0x2)
	})
}
