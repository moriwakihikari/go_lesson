package main

import "fmt"

//panic 強制終了
//recover panicを復帰させる

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Start")
}