package mypkg

import "fmt"

func MypkgFuncA() {			// 大文字で始まるものは自動的にエクスポートされる
    fmt.Println("MyPkgFuncA()")
}

func mypkgFuncB() {			// 小文字で始まるものはエクスポートされない
    fmt.Println("mypkgFuncB()")
}
