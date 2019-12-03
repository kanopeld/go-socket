package core

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetEmitter(t *testing.T) {
	convey.Convey("testing get emitter", t, func() {
		em := GetEmitter(&FakeNetComm{})
		convey.So(em, convey.ShouldNotBeNil)
	})
}

func TestDefaultEmitter_Emit(t *testing.T) {
	convey.Convey("testing emitter send sting message", t, func() {
		em := GetEmitter(&FakeNetComm{})
		convey.So(em, convey.ShouldNotBeNil)
		err := em.Emit("test", "test")
		convey.So(err, convey.ShouldBeNil)
	})

	convey.Convey("testing emitter send bytes slice message", t, func() {
		em := GetEmitter(&FakeNetComm{})
		convey.So(em, convey.ShouldNotBeNil)
		err := em.Emit("test", []byte("test"))
		convey.So(err, convey.ShouldBeNil)
	})

	convey.Convey("testing emitter send not support message type", t, func() {
		em := GetEmitter(&FakeNetComm{})
		convey.So(em, convey.ShouldNotBeNil)
		err := em.Emit("test", []uint{})
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
	})

	convey.Convey("testing emitter send not support message type", t, func() {
		em := GetEmitter(&FakeNetComm{})
		convey.So(em, convey.ShouldNotBeNil)
		err := em.Emit("test", 124142131)
		convey.So(err, convey.ShouldEqual, ErrUnsupportedArgType)
	})
}
