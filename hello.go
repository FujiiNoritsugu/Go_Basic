package main // mainパッケージであることを宣言

import "fmt" // fmtモジュールをインポート

func main() {		// 最初に実行されるmain()関数を定義
    fmt.Println("hello, world")
	// 変数の定義
	var a int			// 通常の定義
	var b int = 123		// 初期値の指定
	var name = "456"		// 初期値の型が明白な場合varを省略できる
	age := 789			// 初期値を指定する際「:=」でvarも省略
	var (			// まとめて定義できる
	e int = 987
	f int = 654
	)
	a = 321			// = で右辺の値を左辺に代入
	name, age = "Fujii", 54// 複数の値を同時に代入
	const dummy = 100
	const (			// 複数の値をまとめて定義できる
		pom = 100
		qom = 200
	)
	fmt.Println(a, b, name, age, e, f, dummy, pom, qom)
	// コメントの書き方
	// これはコメントです。
	/*
	これもコメントです。
	*/
	num := 123; str := "ABC"
	fmt.Println(num, str)
	// 型の定義
	type A string
	type B string
	var c A = "ABC"
	var d B = "ABC"
	// c = d
	fmt.Println(c, d)
	AToB := A(d) 
	fmt.Println(AToB)
	// 配列
	a1 := [3]string{}
	a1[0] = "Red"
	a1[1] = "Green"
	a1[2] = "Blue"
	fmt.Println(a1[0], a1[1], a1[2])
	a2 := [3]int{1, 2, 3}
	a3 := [...]bool{true, false, true}
	fmt.Println(a2[0], a2[1], a2[2])
	fmt.Println(a3[0], a3[1], a3[2])
}
