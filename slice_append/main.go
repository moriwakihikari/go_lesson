package main

import "fmt"

//append 追加
//make ?
//len 要素カウント
//cap 容量カウント パフォーマンスの話
//copy

func main() {
// 	sl := []int{100, 200}
// 	fmt.Println(sl)

// 	sl = append(sl, 300)
// 	fmt.Println(sl)

// 	sl = append(sl , 400, 500, 600)
// 	fmt.Println(sl)

// 	sl2 := make([]int, 5)
// 	fmt.Println(sl2)

// 	fmt.Println(len(sl2))

// 	fmt.Println(cap(sl2))

// 	sl3 := make([]int, 5, 10)

// 	fmt.Println(len(sl3))

// 	fmt.Println(cap(sl3))

// 	sl3 = append(sl3, 1 ,2 ,3 ,4 ,5 ,6 ,7)
// 	fmt.Println(cap(sl3))


//同じになってしまう
// sl := []int{100, 200}
// sl2 := sl

// sl2[0] = 1000

// fmt.Println(sl)

// var i int = 10
// i2 := i
// i2 = 100
// fmt.Println(i, i2)

//copy

	// sl := []int{1, 2 ,3 ,4 ,5}
	// sl2 := make([]int , 5 ,10)
	// fmt.Println(sl2)

	// n := copy(sl2, sl)

	// fmt.Println(n, sl2)

	//スライス
	//for

	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// for i := range sl {
	// 	fmt.Println(i, sl[i])
	// }

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

	

}
