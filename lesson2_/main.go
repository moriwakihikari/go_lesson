package main

import (
	"fmt"
	"strconv"
)

const Pi = 3.14

const (
	URL = "http://xxx.co.jp"
	siteName = "test"
)

const (
	a = 1
	b
	c
	d = "A"
	e
	f
)

//iota 連番
const (
	c0 = iota
	c1
	c2
)


//整数型

func main() {
	// var i int = 100

	// var i2 int64 = 200

	// fmt.Println(i + 50)

	// //型を出力
	// fmt.Printf("%T\n", i2)

	// //型変換
	// fmt.Printf("%T\n", int32(i2))


	// //浮動小数点型
	// var fl64 float64 = 2.4
	// fmt.Println(fl64)

	// fl := 3.2
	// fmt.Println(fl64 * fl)
	// fmt.Printf("%T, %T\n", fl64, fl)

	// var fl32 float32 = 1.2
	// fmt.Printf("%T\n", fl32)

	// fmt.Printf("%T\n", float64(fl32))

	// zero := 0.0
	// pinf := 1.0 / zero
	// fmt.Println(pinf)

	// ninf := -1.0 / zero
	// fmt.Println(ninf)

	// nan := zero / zero
	// fmt.Println(nan)


	// var t, f bool = true, false
	// fmt.Println(t, f)


	// //文字列
	// var s string = "Hello Golang"
	// fmt.Println(s)
	// fmt.Printf("%T\n", s)

	// var si string = "300"
	// fmt.Printf(si)
	// fmt.Printf("%T\n", si)

	// fmt.Println(`test
	// test
	//     test`)

	// 	fmt.Println("\"")
	// 	fmt.Println(`"`)

	// 	//バイト型
	// 	fmt.Println(s[0])

	// 	//文字数1文字目（”H”）
	// 	fmt.Println(string(s[0]))

	// byteA := []byte{72, 73}
	// fmt.Println(byteA)
	// fmt.Println(string(byteA))

	// c := []byte("HI")
	// fmt.Println(c)

	// fmt.Println(string(c))

	// //配列型
	// var arr1 [3]int
	// fmt.Println((arr1))
	// fmt.Printf("%T\n", arr1)

	// var arr2 [3]string = [3]string{"A", "B", "C"}
	// fmt.Println(arr2)

	// arr3 := [3]int{1,2,3}
	// fmt.Println(arr3)

	// //...配列の数を自動で数える
	// arr4 := [...]string{"C", "D"}
	// fmt.Println(arr4)
	// fmt.Printf("%T\n", arr4)

	// fmt.Println(arr2[0])

	// arr2[2] = "C"
	// fmt.Println(arr2)
	
	// //interface & nil
	// var x interface{}
	// fmt.Println(x)

	// x = 1
	// fmt.Println(x)

	// x=3.14
	// fmt.Println(x)

	// x="A"
	// fmt.Println(x)

	// //計算できない
	// // x = 2
	// // fmt.Println(x + 3)

	//型変換
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	// var s string = "100"
	// fmt.Printf("s = %T\n", s)

	// i, err := strconv.Atoi("A")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)

	//定数 関数の外で定義
	fmt.Println(Pi)

	fmt.Println(URL, siteName)

	fmt.Println(a,b,c,d,e,f)

	fmt.Println(c0,c1,c2)
}