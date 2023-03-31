package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

// init関数は特別な関数でmain関数よりも先に呼ばれる
func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

// 変数宣言だけ (varを用いた関数宣言は関数外でも使用できる、スコープが違えば同じ変数名は宣言できる)
// var (
// 	i    int
// 	f64  float64
// 	s    string
// 	t, f bool
// )

// constで定義する時は型定義をしなくていい
const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

// varだとint型の最大値を宣言できない
// var big int = 9223372036854775807 + 1
// constではできる
const big = 9223372036854775807 + 1

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func convert(price int) float64 {
	return float64(price)
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	// import文

	bazz()
	fmt.Println("Hello world", time.Now())
	fmt.Println(user.Current())

	// 変数宣言 (明示的に型を宣言したいときはvarを用いる。varや:=を用いるのは定義時のみ。値を変えたいときは=を用いる)

	// var i int = 1
	// var f64 float64 = 1.2
	// var s string = "test"
	// var t, f bool = true, false

	// 変数宣言だけ
	// var (
	// 	i    int
	// 	f64  float64
	// 	s    string
	// 	t, f bool
	// )

	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// 関数内でしかショートの変数宣言は出来ない
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)

	// const

	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
	// constで定義されたものは出力時にコンパイルエラーが起きる
	// fmt.Println(big)

	// 数値型

	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v", u8, u8)

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 / 3 =", 10.0+3)
	fmt.Println("10 / 3.0 =", 10/3.0)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)

	x := 0
	fmt.Println(x)
	// x = x + 1
	x++
	fmt.Println(x)
	// x = x - 1
	x--
	fmt.Println(x)

	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000

	// 文字列型

	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello World"[0]))

	var ss string = "Hello World"
	ss = strings.Replace(ss, "H", "X", 1)
	fmt.Println(ss)
	fmt.Println(strings.Contains(ss, "World"))

	fmt.Println(`Test
											Test
Test`)
	fmt.Println("\"")
	fmt.Println(`"`)

	// 論理値型
	tt, ff := true, false
	fmt.Printf("%T %v %t\n", tt, 1, tt)
	fmt.Printf("%T %v %t\n", ff, 1, ff)

	// fmt.Println(true && true)
	// fmt.Println(true && false)
	// fmt.Println(false && false)

	// fmt.Println(true || true)
	// fmt.Println(true || false)
	// fmt.Println(false || false)

	// fmt.Println(!true)
	// fmt.Println(!false)

	// 型変換

	var xx int = 1
	xxx := float64(xx)
	fmt.Printf("%T %v %f\n", xxx, xxx, xxx)

	var yy float64 = 1.2
	yyy := int(yy)
	fmt.Printf("%T %v %d\n", yyy, yyy, yyy)

	var sss string = "14"
	i, err := strconv.Atoi(sss)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))

	// 配列

	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)
	fmt.Println(a[1])

	// var b [2]int = [2]int{100, 200}
	// fmt.Println(b)

	b := [2]int{100, 200}
	fmt.Println(b)

	// スライス

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])

	n[2] = 100
	fmt.Println(n)
	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)

	// スライスのmakeとcap

	nn := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn)
	nn = append(nn, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn)
	nn = append(nn, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn)

	aa := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(aa), cap(aa), aa)

	// メモリに確保する
	bb := make([]int, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(bb), cap(bb), bb)
	// メモリに確保はしない (nil)
	var cc []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(cc), cap(cc), cc)

	// cc = make([]int, 5)
	cc = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		cc = append(cc, i)
		fmt.Println(cc)
	}
	fmt.Println(cc)

	// map

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// メモリに確保する
	m2 := make(map[string]int)
	fmt.Println(m2)
	if m2 == nil {
		fmt.Println("Nill")
	} else {
		fmt.Println("not nil")
	}
	m2["pc"] = 5000
	fmt.Println(m2)

	// メモリに確保しない
	var m3 map[string]int
	fmt.Println(m3)

	var ssss []int
	if ssss == nil {
		fmt.Println("Nill")
	}

	// バイト型

	bbb := []byte{72, 73}
	fmt.Println(bbb)
	fmt.Println(string(bbb))

	ccc := []byte("HI")
	fmt.Println(ccc)
	fmt.Println(string(ccc))

	// 関数

	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	fff := func(x int) {
		fmt.Println("inner func", x)
	}
	fff(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)

	// クロージャー

	counter := incrementGenerator()
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	c2 := circleArea(3)
	fmt.Println(c2(2))

	// 可変長引数

	foo()
	foo(10, 20, 30)

	sssss := []int{1, 2, 3}
	fmt.Println(sssss)

	foo(sssss...)

	// Q1
	ffff := 1.11
	fnew := int(ffff)
	fmt.Printf("%T %v\n", fnew, fnew)

	// Q2
	// コードを書かずに以下の出力結果を答えてください。
	// s := []int{1, 2, 5, 6, 2, 3, 1}
	// fmt.Println(s[2:4])
	// 5,6

	// Q3
	// 以下のコードを実行した時に

	// fmt.Printf("%T %v", m, m)

	// 以下のような出力結果となるmを作成してください。

	// map[string]int map[Mike:20 Nancy:24 Messi:30]

	mm := map[string]int{}
	mm["Mike"] = 20
	mm["Nancy"] = 24
	mm["Messi"] = 30
	fmt.Printf("%T %v", mm, mm)

}
