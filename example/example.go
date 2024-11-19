// -------------------------------------------
// @file      : example.go
// @author    : 蔡波
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 上午11:36
// -------------------------------------------

package main

import (
	"fmt"
	"github.com/caibo86/cberrors"
)

func main() {
	err := cberrors.New("test error")
	fmt.Println(err)
}
