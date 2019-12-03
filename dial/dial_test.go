package dial

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewDial(t *testing.T) {
	convey.Convey("testing get dial", t, func() {
		_, err := NewDial(":81")
		convey.So(err, convey.ShouldNotBeNil)
	})
}
