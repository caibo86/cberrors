// -------------------------------------------
// @file      : example.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 上午11:36
// -------------------------------------------

package main

import (
	"fmt"
	"github.com/caibo86/cberrors"
)

func main() {
	example1()
	example2()
	example3()
}

func example1() {
	// 创建一个新的错误
	s := "hello world"
	err := cberrors.New("这是的一个新的错误:%s", s)
	fmt.Println(err)
	// 包装一个已有的错误
	err1 := fmt.Errorf("原生golang的错误类型")
	err = cberrors.Wrap(err1)
	fmt.Println(err)
}

func example2() {
	// 输入格式化信息并panic
	s := "hello world"
	cberrors.Panic("这是一个panic错误:%s", s)
}

func example3() {
	// 利用一个已有的系统错误包装并panic
	s := "hello world"
	err := fmt.Errorf("这是一个panic错误:%s", s)
	cberrors.PanicWrap(err)
}
