package main

import (
	"bufio"
	"fmt"
	"os"
)

//ターミナルで打った文字が返ってくる

func main() {
	//標準入力をソースにしたスキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	//入力のスキャンが成功する限り繰り返す
	for scanner.Scan() {
		//スキャン出力
		fmt.Println(scanner.Text())
	}

	//エラー処理
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr)
	}

}
