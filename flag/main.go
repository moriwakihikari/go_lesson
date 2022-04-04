package main

import (
	"flag"
	"fmt"
)

func main() {
	//コマンドラインのオプション処理
	//コマンドラインを処理するサンプル
	//go run main.go -n 20 -m message -x
	var (
		max int
		msg string
		x   bool
	)
	//整数
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	//文字列
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	//bool
	flag.BoolVar(&x, "x", false, "拡張オプション")
	flag.Parse()

	fmt.Println("処理数の最大値 =", max)
	fmt.Println("処理メッセージ =", msg)
	fmt.Println("拡張オプション =", x)

}