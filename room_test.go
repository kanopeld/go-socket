package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

var testRoom = getRoom()

func TestRoom_SetClient(t *testing.T) {
	convey.Convey("testing room set client. regular case", t, func() {
		err := testRoom.SetClient(&fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(testRoom.Len(), convey.ShouldEqual, 1)
	})

	convey.Convey("testing room set client, duplicate client id", t, func() {
		err := testRoom.SetClient(&fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldEqual, ErrClientAlreadyInRoom)
		convey.So(testRoom.Len(), convey.ShouldEqual, 1)
	})

	convey.Convey("testing broadcast join", t, func() {
		err := testRoom.SetClient(&fakeServerClient{id: "tests1"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(testRoom.Len(), convey.ShouldEqual, 2)
	})
}

func TestRoom_RemoveClient(t *testing.T) {
	_ = testRoom.SetClient(&fakeServerClient{id: "tests"})
	_ = testRoom.SetClient(&fakeServerClient{id: "tests1"})

	convey.Convey("testing room set client. regular case", t, func() {
		err := testRoom.RemoveClient(&fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(testRoom.Len(), convey.ShouldEqual, 1)
	})

	convey.Convey("testing room set client, duplicate client id", t, func() {
		err := testRoom.RemoveClient(&fakeServerClient{id: "tests"})
		convey.So(err, convey.ShouldEqual, ErrClientNotInRoom)
		convey.So(testRoom.Len(), convey.ShouldEqual, 1)
	})

	convey.Convey("testing broadcast join", t, func() {
		err := testRoom.RemoveClient(&fakeServerClient{id: "tests1"})
		convey.So(err, convey.ShouldBeNil)
		convey.So(testRoom.Len(), convey.ShouldEqual, 0)
	})
}
