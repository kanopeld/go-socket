package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDecodePackage(t *testing.T) {
	convey.Convey("test decode packet", t, func() {
		msgByte := Message{Data: []byte("hello"), EventName: "test"}.MarshalBinary()
		p := Package{PT: PackTypeEvent, Payload: msgByte}.MarshalBinary()

		//Remove \n char. In real code this will remove by method ReadBytes("\n")!
		comlPack := p[:len(p)-1]

		pack, err := decodePackage(comlPack)
		convey.So(err, convey.ShouldBeNil)
		convey.So(pack.PT, convey.ShouldEqual, PackTypeEvent)
		convey.So(pack.Payload, convey.ShouldHaveLength, len(msgByte))

		msg := decodeMessage(pack.Payload)
		convey.So(msg.EventName, convey.ShouldEqual, "test")
		convey.So(string(msg.Data), convey.ShouldEqual, "hello")
	})

	convey.Convey("test decode packet", t, func() {
		msgByte := Message{Data: append([]byte("hello"), []byte("\n")...), EventName: "test"}.MarshalBinary()
		p := Package{PT: PackTypeEvent, Payload: msgByte}.MarshalBinary()

		//Remove \n char. In real code this will remove by method ReadBytes("\n")!
		comlPack := p[:len(p)-1]

		pack, err := decodePackage(comlPack)
		convey.So(err, convey.ShouldBeNil)
		convey.So(pack.PT, convey.ShouldEqual, PackTypeEvent)
		convey.So(pack.Payload, convey.ShouldHaveLength, len(msgByte)-1)

		msg := decodeMessage(pack.Payload)
		convey.So(msg.EventName, convey.ShouldEqual, "test")
		convey.So(string(msg.Data), convey.ShouldEqual, "hello")
	})

	convey.Convey("test decode packet with bytes message data", t, func() {
		msgByte := Message{
			EventName: "test",
			Data:      []byte{0x1, 0x2, 0x03},
		}.MarshalBinary()

		p := Package{PT: PackTypeEvent, Payload: msgByte}.MarshalBinary()
		//Remove \n char. In real code this will remove by method ReadBytes("\n")!
		comlPack := p[:len(p)-1]

		pack, err := decodePackage(comlPack)
		convey.So(err, convey.ShouldBeNil)
		convey.So(pack.PT, convey.ShouldEqual, PackTypeEvent)
		convey.So(pack.Payload, convey.ShouldHaveLength, len(msgByte))

		msg := decodeMessage(pack.Payload)
		convey.So(msg.EventName, convey.ShouldEqual, "test")
		convey.So(msg.Data, convey.ShouldHaveLength, 3)

		convey.So(msg.Data[0], convey.ShouldEqual, byte(0x1))
		convey.So(msg.Data[1], convey.ShouldEqual, byte(0x2))
		convey.So(msg.Data[2], convey.ShouldEqual, byte(0x03))
	})

	convey.Convey("test get packet byte", t, func() {
		p1 := PackTypeEvent
		convey.So(p1.Byte(), convey.ShouldEqual, 0x02)

		p2 := PackTypeConnect
		convey.So(p2.Byte(), convey.ShouldEqual, 0x00)

		p3 := PackTypeDisconnect
		convey.So(p3.Byte(), convey.ShouldEqual, 0x01)
	})
}
