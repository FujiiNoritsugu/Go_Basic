package main // mainパッケージであることを宣言

import (
	"errors"
	"fmt"
	"hello/local/mypkg"
	"os"
	"time"

	"golang.org/x/example/hello/reverse"
)
func main() {		// 最初に実行されるmain()関数を定義
	fmt.Println(reverse.String("Hello, Go!"))
	
	mypkg.MypkgFuncA()
	// mypkg.mypkgFuncB()
	
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
	p2_1 := Person2{"Yamada", 26, 1, "day of dream"}			// 順序通りに初期化
	p2_2 := Person2{Name: "Tanaka", Age: 32}		// 名前で初期化
	fmt.Println(p2_1, p2_2)
    PrintOut(p1)
    PrintOut(p2_1)
	funcC(3)
	funcD(false)

	p4 := map[string]interface{} {
		"name": "Yamada",
		"age": 26,
	}
	PrintOut(p4)
	p5 := dict {
		"name": "Yamada",
		"age": 26,
		"address": dict {
			"zip": "123-4567",
			"tel": "012-3456-7890",
		},
	}
	p5_name := p5["name"]
	p5_tel := p5["address"].(dict)["tel"]	// anyをdictに変換してから参照
	fmt.Println(p5_name)
	fmt.Println(p5_tel)

	var norm_a int		// int型変数a1を定義
	var pointer_a *int;		// intへのポインタ変数p1を定義

	pointer_a = &norm_a		// p1にa1のポインタを設定
	*pointer_a = 123		// ポインタの中身(つまりa1)に123を代入
	fmt.Println(norm_a)	// => 123
	*pointer_a = 456		// ポインタの中身(つまりa1)に123を代入
	fmt.Println(norm_a)	// => 123

	var a7 int = 123
    var a8 int = 123
    fn(a7, &a8)		// a1は値渡し、a2は参照渡し
    fmt.Println(a7, a8)	// => 123 456

	a9 := Person{"Tanaka", 26}	// 構造体Personのオブジェクトa1を確保して初期化
	p9 := &a9			// 構造体a1へのポインタをp1に格納
	fmt.Println(a9.name)		// メンバ変数には左記のようにアクセス
	fmt.Println((*p9).name)		// ポインタpの中身(後続体)のメンバ変数には左記のようにアクセス
	fmt.Println(p9.name)		// ただし、Go言語ではこれを、左記のようにも記述できる
	
    bookList := []*Book{}

    for i := 0; i < 10; i++ {
        book := new(Book)
        book.title = fmt.Sprintf("Title#%d", i)
        bookList = append(bookList, book)
    }
    for _, book := range bookList {
        fmt.Println(book.title)
    }
	funcE()

	go funcF()
    for i := 0; i < 10; i++ {
        fmt.Print("M")
        time.Sleep(20 * time.Millisecond)
    }

	chA := make(chan string)	// チャネルを作成する
    defer close(chA)		// 使い終わったらクローズする
    go funcG(chA)		// チャネルをゴルーチンに渡す
    msg := <- chA		// チャネルからメッセージを受信する
    fmt.Println(msg)

	chC := make(chan string)
    chD := make(chan string)
    defer close(chC)
    defer close(chD)
    finflagC := false
    finflagD := false
    go funcH(chC)
    go funcI(chD)
    for {
        select {
        case msg := <- chC:
            finflagC = true
            fmt.Println(msg)
        case msg := <- chD:
            finflagD = true
            fmt.Println(msg)
        }
        if finflagC && finflagD {
            break
        }
    }

}

func funcH(chA chan <- string) {
    time.Sleep(1 * time.Second)
    chA <- "funcH Finished"
}

func funcI(chB chan <- string) {
    time.Sleep(2 * time.Second)
    chB <- "funcI Finished"
}

func funcG(chA chan <- string) {
    time.Sleep(3 * time.Second)
    chA <- "Finished"		// チャネルにメッセージを送信する
}

func funcF() {
    for i := 0; i < 10; i++ {
        fmt.Print("A")
        time.Sleep(10 * time.Millisecond)
    }
}


func funcE() {
    fp, err := os.Open("/home/fujii/Go_Basic/sample.txt")
    if err != nil {
        return
    }
    defer fp.Close()

    for {
		var temp []byte
        n, err := fp.Read(temp)
		if err != nil {
			fmt.Printf("エラーが発生しました。%v", err)
			break
		}
		fmt.Printf("n=%d", n)
		if n == 0{
			break
		}
		fmt.Println(temp)
    }
}

type Book struct {
    title string
}

func fn(b1 int, b2 *int) {
    b1 = 456
    *b2 = 456
}

type any interface{}
type dict map[string]any

func funcD(a interface{}) {
    switch a.(type) {
    case bool:
        fmt.Printf("%t\n", a)
    case int:
        fmt.Printf("%d\n", a)
    case string:
        fmt.Printf("%s\n", a)
    }
}

func funcC(a interface{}) {
    fmt.Printf("%d\n", a.(int))
}

type Printable interface {
    ToString() string
}

func PrintOut(a interface{}) {
    // aをPrintableインタフェースを実装したオブジェクトに変換してみる
    q, ok := a.(Printable)
    if ok {
        // 変換できたらそのインタフェースを呼び出す
        fmt.Println(q.ToString())
    } else {
        fmt.Println("Not printable.")
    }
}

func (p Person) ToString() string {
    return p.name
}

func (p2 Person2) ToString() string {
    return p2.title
}

type Person2 struct {
    Name string	// 外部からアクセス可能
    Age int		// 外部からアクセス可能
    status int		// 外部からアクセス不可
	title string
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
