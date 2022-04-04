package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//入力全体を読み込む
	f, _ := os.Open("foo.txt")
	bs1, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs1))

	//ファイルに書き込む
	bs2, _ := ioutil.ReadFile("foo.txt")
	fmt.Println(string(bs2))

}
