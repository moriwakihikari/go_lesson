package main //packageは一つ

//必要なライブラリをimport
import (
	"fmt"
	"time"
)

// import "time"

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())

	//明示的な定義
	//変数(var 変数名 型 = 値)
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int    //0
	var s3 string //""
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"

	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	//暗黙的な定義(型が決まる,:=再代入出来ないconstの様なもの,関数内でしか使えない)
	//変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 := 500
	// fmt.Println(i4)

	// i4 = "Hello"
	// fmt.Println(i4)

	outer()

	// fmt.Println(s4)

	var s5 string = "Not Use"
	fmt.Println(s5)
}