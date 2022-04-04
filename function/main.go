package main

import (
	"fmt"
	"strconv"
)

func Plus(x int, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

//1行目で引数と返り値を設定することも
func Double(price int) (result int) {
	result = price * 2
	return 
}

//関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

//関数を引数にとる関数
func CallFunction(k func()) {
	k()
}

//クロージャー
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

//ジェネレーター
func integers() func() int {
	b := 0
	return func() int {
		b++
		return b
	}
}



func main() {
	i := Plus(1,2)
	fmt.Println(i)

	// i2, i3 := Div(9, 3)
	// fmt.Println(i2, i3)

	//返り値全て使いたくない場合_で破棄
	i2, _ := Div(9, 3)
	fmt.Println(i2)

	i4 := Double(1000)
	fmt.Println(i4)

	
	//無名関数
	f := func(x, y int) int {
		return x + y
	}
	j := f(1, 2)
	fmt.Println(j)

	j2 := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(j2)

	f2 := ReturnFunc()
	f2() 

	CallFunction(func() {
		fmt.Println("Im a function")
	})

	a := Later()
	fmt.Println(a("Hello"))
	fmt.Println(a("My"))
	fmt.Println(a("name"))
	fmt.Println(a("is"))
	fmt.Println(a("Golang"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	oeherints := integers()

	fmt.Println(oeherints())
	fmt.Println(oeherints())
	fmt.Println(oeherints())
	fmt.Println(oeherints())
	fmt.Println(oeherints())
	fmt.Println(oeherints())

	//エラーハンドリング
	var c string = "100"
	// var c string = "100"

	d, err := strconv.Atoi(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", d)

	//for
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v:= range m {
		fmt.Println(k, v)
	}


	//switch
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't Know")
	}

	//型スイッチ
}