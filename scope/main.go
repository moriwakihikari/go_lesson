package main

import (
	"fmt"
	"go_lesson/scope/foo"
	// f "fmt"
)

//スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	// fmt.Println(foo.min)

	fmt.Println(foo.ReturnMin())
	fmt.Println(foo.Max)

	fmt.Println(appName())
	// fmt.Println(AppName, Version)
	fmt.Println(Do("aaa"))
}