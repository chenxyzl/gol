package version_impl

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestName2(t *testing.T) {
	Convey("test version", t, func() {
		Convey("1", func() {
			v1, err := NewVersion("v2.5")
			if err != nil {
				fmt.Println(err.Error())
				So(err, ShouldNotBeNil)
			}
			v2, err := NewVersion("v2.5-beta")
			if err != nil {
				fmt.Println(err.Error())
				So(err, ShouldNotBeNil)
			}
			So(v1.GT(*v2), ShouldEqual, false)
		})

		Convey("2", func() {
			v1, err := NewVersion("v2.-5")
			if err != nil {
				fmt.Println(err.Error())
				So(err, ShouldNotBeNil)
			}
			v2, err := NewVersion("v2.5-beta")
			if err != nil {
				fmt.Println(err.Error())
				So(err, ShouldNotBeNil)
			}
			_ = v1
			_ = v2
		})

		Convey("3", func() {
			v1, _ := NewVersion("v2.5.2.4")
			v2, _ := NewVersion("v2.2")
			So(v1.LE(*v2), ShouldEqual, true)
		})

		Convey("4", func() {
			v1, _ := NewVersion("v2.5.2.4")
			v2, _ := NewVersion("v2.2.2.2.2")
			So(v1.LE(*v2), ShouldEqual, true)
			So(v1.GE(*v2), ShouldEqual, false)
		})

		Convey("func 5", func() {
			v1, _ := NewVersion("v2.5.2.4")
			v2, _ := NewVersion("v2.5.2.2.2")
			So(v1.EQ(*v2), ShouldEqual, true)
		})

		Convey("func 6", func() {
			v1, _ := NewVersion("v2.5.2.2")
			v2, _ := NewVersion("v2.5.2.2.2")
			So(v1.EQ(*v2), ShouldEqual, true)
		})

		Convey("func 7", func() {
			v1, _ := NewVersion("ver2.5.2-beta")
			v2, _ := NewVersion("version2.5.2-public")
			So(v1.EQ(*v2), ShouldEqual, true)
		})

	})
}
