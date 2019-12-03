package server

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewServer(t *testing.T) {
	convey.Convey("testing get new server", t, func() {
		_, err := NewServer(":81881")
		convey.So(err, convey.ShouldNotBeNil)
	})
}
