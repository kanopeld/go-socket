package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDecodePackage(t *testing.T) {
	convey.Convey("test decode packet", t, func() {
		msgByte := Message{Data: []byte("hello"), EventName: "test"}.MarshalBinary()
		p := Package{PT: _PACKET_TYPE_EVENT, Payload: msgByte}.MarshalBinary()

		//Remove \n char. In real code this will remove by method ReadBytes("\n")!
		comlPack := p[:len(p)-1]

		pack, err := DecodePackage(comlPack)
		convey.So(err, convey.ShouldBeNil)
		convey.So(pack.PT, convey.ShouldEqual, _PACKET_TYPE_EVENT)
		convey.So(pack.Payload, convey.ShouldHaveLength, len(msgByte))

		msg := DecodeMessage(pack.Payload)
		convey.So(msg.EventName, convey.ShouldEqual, "test")
		convey.So(string(msg.Data), convey.ShouldEqual, "hello")
	})
}
