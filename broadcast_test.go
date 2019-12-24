package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBroadcast_Join(t *testing.T) {
	bc := newDefaultBroadcast()

	convey.Convey("testing broadcast join", t, func() {
		err := bc.join(DefaultBroadcastRoomName, &fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)
	})

	convey.Convey("testing broadcast join, duplicate client id", t, func() {
		err := bc.join(DefaultBroadcastRoomName, &fakeServerClient{"tests"})
		convey.So(err, convey.ShouldEqual, ErrClientAlreadyInRoom)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)
	})

	convey.Convey("testing broadcast join", t, func() {
		err := bc.join(DefaultBroadcastRoomName, &fakeServerClient{"tests1"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, 2)
	})

	convey.Convey("testing broadcast join", t, func() {
		err := bc.join("test", &fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.len("test"), convey.ShouldEqual, 1)
	})
}

func TestBroadcast_Leave(t *testing.T) {
	bc := newDefaultBroadcast()

	convey.Convey("testing broadcast leave", t, func() {
		err := bc.join(DefaultBroadcastRoomName, &fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)

		err = bc.leave(DefaultBroadcastRoomName, &fakeServerClient{id: "error"})
		convey.So(err, convey.ShouldEqual, ErrClientNotInRoom)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)

		err = bc.leave(DefaultBroadcastRoomName, &fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, -1)

		err = bc.leave("qwe", &fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldEqual, ErrRoomNotExist)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, -1)
	})
}

func TestBroadcast_Send(t *testing.T) {
	bc := newDefaultBroadcast()

	convey.Convey("testing broadcast send", t, func() {
		err := bc.join(DefaultBroadcastRoomName, &fakeServerClient{id: "tests1"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, 1)

		err = bc.join(DefaultBroadcastRoomName, &fakeServerClient{id: "tests2"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, 2)

		err = bc.join(DefaultBroadcastRoomName, &fakeServerClient{id: "tests3"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(bc.len(DefaultBroadcastRoomName), convey.ShouldEqual, 3)

		err = bc.send(&fakeServerClient{id: "tests1"}, DefaultBroadcastRoomName, "test", nil)
		convey.So(err, convey.ShouldBeNil)

		err = bc.send(&fakeServerClient{id: "tests1"}, "test", "test", nil)
		convey.So(err, convey.ShouldEqual, ErrRoomNotExist)

		err = bc.send(nil, DefaultBroadcastRoomName, "test", nil)
		convey.So(err, convey.ShouldBeNil)
	})
}
