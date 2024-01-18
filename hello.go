package main // mainパッケージであることを宣言

import (
	"errors"
	"fmt"
)

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
	// スライス
	a4 := []string{}			// スライス。個数不定
	a4 = append(a4, "Sweet")
	a4 = append(a4, "Bitter")
	a4 = append(a4, "Spicy")
	fmt.Println(a4[0], a4[1], a4[2])
	a5 := []int{}
	for i := 0; i < 10; i++ {
		a5 = append(a5, i)
		fmt.Println(len(a5), cap(a5))
	}
	bufa := make([]byte, 0, 1024)
	fmt.Println(len(bufa), cap(bufa))
	// マップ
	// マップを定義する
	a6 := map[string]int{
		"x": 100,
		"y": 200,			// 改行する場合はカンマ必須
	}

	// マップを参照する
	fmt.Println(a6["x"])

	// マップに要素を加える
	a6["z"] = 300

	// マップの要素を削除する
	delete(a6, "z")

	// マップの長さを求める
	fmt.Println(len(a6))

	// マップに要素が存在するかどうかを調べる
	_, ok := a6["z"]
	if ok {
		fmt.Println("Exist")
	} else {
		fmt.Println("Not exist")
	}

	// マップをループ処理する
	for key, value := range a6 {
		fmt.Printf("%s=%d\n", key, value)
	}

	// If文
	x := 1
	y := 2

	if x > y {
		fmt.Println("x is greater then y")
	} else {
		fmt.Println("x is less then y")
	}
	
	if x > y {
		fmt.Println("Big")
	} else if x < y {
		fmt.Println("Small")
	} else {
		fmt.Println("Equal")
	}

	dayOfWeek := "Thu"
	switch dayOfWeek {
	case "sat":
		fallthrough
	case "sun":
		fmt.Println("Horiday")
	default:
		fmt.Println("Weekday")
	}
	
	// for文
	for x < y {  // x が y よりも小さい間、処理を繰り返す
		x++
	}
	
	for i := 0; i < 10; i++ {  // for 開始処理; 条件; 後処理 { 処理 }
		fmt.Println(i)
	}
	
	n := 0
	for {
		n++
		if n > 10 {
			break
		} else if n % 2 == 1 {
			continue
		} else {
			fmt.Println(n)
		}
	}

	colors := [...]string{"Red", "Green", "Blue"}

	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}
	
    funcA()
	fmt.Println(add(5, 3))
	add, min := addMinus(8, 5)
    fmt.Println(add, min)

	funcB(1, 2, 3, 4, 5)

	var p1 Person
    p1.SetPerson("Yamada", 26)
    set_name, set_age := p1.GetPerson()
    fmt.Printf("%s(%d)\n", set_name, set_age)
	p2_1 := Person2{"Yamada", 26, 1}			// 順序通りに初期化
	p2_2 := Person2{Name: "Tanaka", Age: 32}		// 名前で初期化
	fmt.Println(p2_1, p2_2)

}

type Person2 struct {
    Name string	// 外部からアクセス可能
    Age int		// 外部からアクセス可能
    status int		// 外部からアクセス不可
}

func (p *Person) SetPerson(name string, age int) {
    p.name = name
    p.age = age
}

func (p Person) GetPerson() (string, int) {
    return p.name, p.age
}

type Person struct {
    name string
    age int
}

func funcB(a int, b ... int) {
    fmt.Printf("a=%d\n", a)
    for i, num := range b {
        fmt.Printf("b[%d]=%d\n", i, num)
    }
}

func addMinus(x int, y int) (int, int) {
    return x + y, x - y
}


func add(x int, y int) int {
    return x + y
}

func funcA() (string, error) {
    var err error
    filename := ""
    data := ""

    filename, err = GetFileName()
    if err != nil {
        fmt.Println(err)
        goto Done
    }

    data, err = ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        goto Done
    }

    fmt.Println(data)

Done:
    return data, err
}

func GetFileName() (string, error) {
    return "sample.txt", nil
}

func ReadFile(filename string) (string, error) {
    return "Hello world!", errors.New("Can't read file")
}
