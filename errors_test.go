// -------------------------------------------
// @file      : errors_test.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 下午2:19
// -------------------------------------------

package cberrors

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNew(t *testing.T) {
	Convey("创建错误", t, func() {
		Convey(" 正常创建", func() {
			err := New("这是一个新的错误 code:%d", 119)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestPanic(t *testing.T) {
	Convey("Panic", t, func() {
		Convey("输入格式化错误信息并panic", func() {
			defer func() {
				r := recover()
				So(r, ShouldNotBeNil)
				err, ok := r.(CBError)
				So(ok, ShouldBeTrue)
				So(err, ShouldNotBeNil)
			}()
			Panic("这是一个panic错误 code:%d", 119)
		})
	})
}

func TestPanicWrap(t *testing.T) {
	Convey("PanicWrap", t, func() {
		Convey("使用已有错误并调用到panic", func() {
			defer func() {
				r := recover()
				So(r, ShouldNotBeNil)
				err, ok := r.(error)
				So(ok, ShouldBeTrue)
				So(err, ShouldNotBeNil)
			}()
			err := fmt.Errorf("这是一个新的错误 code:%d", 119)
			PanicWrap(err)
		})
	})
}

func TestWrap(t *testing.T) {
	Convey("包装已有错误", t, func() {
		Convey("正常返回", func() {
			err := fmt.Errorf("这是一个新的错误 code:%d", 119)
			ret := Wrap(err)
			So(ret, ShouldNotBeNil)
		})
	})
}

func Test_errorHost_Error(t *testing.T) {
	Convey("返回输出错误信息", t, func() {
		Convey("正常返回", func() {
			Convey("无错误信息", func() {
				err := New("")
				ret := err.Error()
				So(len(ret) > 0, ShouldBeTrue)
			})
			Convey("有错误信息", func() {
				err := New("这是一个新的错误 code:%d", 119)
				ret := err.Error()
				So(len(ret) > 0, ShouldBeTrue)
			})
		})
	})
}

func Test_errorHost_Stack(t *testing.T) {
	Convey("返回栈信息", t, func() {
		Convey("正常返回", func() {
			err := New("这是一个新的错误 code:%d", 119)
			ret := err.Stack()
			So(len(ret) > 0, ShouldBeTrue)
		})
	})
}

func Test_errorHost_String(t *testing.T) {
	Convey("返回错误输出", t, func() {
		Convey("正常返回", func() {
			err := New("这是一个新的错误 code:%d", 119)
			e, ok := err.(fmt.Stringer)
			So(ok, ShouldBeTrue)
			ret := e.String()
			So(len(ret) > 0, ShouldBeTrue)
		})
	})
}

func Test_stack(t *testing.T) {
	Convey("获取堆栈信息", t, func() {
		Convey("正常获取", func() {
			ret := stack()
			So(len(ret) > 0, ShouldBeTrue)
		})
	})
}
