package core

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBroadcast_Join(t *testing.T) {
	bc := newDefaultBroadcast()

	convey.Convey("testing broadcast join", t, func() {
		err := bc.Join(DefaultBroadcastRoomName, &FakeClient{Id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)
	})

	convey.Convey("testing broadcast join, duplicate client id", t, func() {
		err := bc.Join(DefaultBroadcastRoomName, &FakeClient{"tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)
	})

	convey.Convey("testing broadcast join", t, func() {
		err := bc.Join(DefaultBroadcastRoomName, &FakeClient{"tests1"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len(DefaultBroadcastRoomName), convey.ShouldEqual, 2)
	})

	convey.Convey("testing broadcast join", t, func() {
		err := bc.Join("test", &FakeClient{Id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len("test"), convey.ShouldEqual, 1)
	})

	convey.Convey("testing broadcast join, duplicate client id", t, func() {
		err := bc.Join("test", &FakeClient{"tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len("test"), convey.ShouldEqual, 1)
	})

	convey.Convey("testing broadcast join", t, func() {
		err := bc.Join("test", &FakeClient{"tests1"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len("test"), convey.ShouldEqual, 2)
	})
}

func TestBroadcast_Leave(t *testing.T) {
	bc := newDefaultBroadcast()

	convey.Convey("testing broadcast leave", t, func() {
		err := bc.Join(DefaultBroadcastRoomName, &FakeClient{Id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)

		err = bc.Leave(DefaultBroadcastRoomName, &FakeClient{Id: "error"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)

		err = bc.Leave(DefaultBroadcastRoomName, &FakeClient{Id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len(DefaultBroadcastRoomName), convey.ShouldEqual, 0)
	})

	convey.Convey("testing broadcast leave", t, func() {
		err := bc.Join(DefaultBroadcastRoomName, &FakeClient{Id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)

		err = bc.Leave("test", &FakeClient{Id: "error"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len("test"), convey.ShouldEqual, -1)

		err = bc.Leave(DefaultBroadcastRoomName, &FakeClient{Id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.Len(DefaultBroadcastRoomName), convey.ShouldEqual, 0)
	})
}
