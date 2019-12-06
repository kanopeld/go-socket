package socket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetEmitter(t *testing.T) {
	convey.Convey("testing get emitter", t, func() {
		em := getEmitter(&fakeNetConn{})
		convey.So(em, convey.ShouldNotBeNil)
	})
}

func TestDefaultEmitter_Emit(t *testing.T) {
	convey.Convey("testing emitter send bytes slice message", t, func() {
		em := getEmitter(&fakeNetConn{})
		convey.So(em, convey.ShouldNotBeNil)
		err := em.Emit("test", []byte("test"))
		convey.So(err, convey.ShouldBeNil)
	})
}
