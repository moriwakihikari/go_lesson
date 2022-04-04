package main

import (
	"fmt"
	"math/rand"
	"time"
)

//4.rand
//疑似乱数を生成する昨日がまとめられた機能。
//ランタイム全体で共有される疑似乱数生成と、任意のシード値を元にした疑似乱数生成などができる。
//最も単純な使い方は、rand.Seedへ任意のint64の数値を渡してデフォルトの疑似乱数生成器のシードを設定する方法
//rand.Float64はデフォルトの疑似乱数生成器を使って、0.0<=n<1.0　の条件を満たす疑似乱数を生成する。


func main() {
	//デフォルトの疑似乱数生成器のシードを設定
	//シードに設定された数値が256に固定されているため、何度実行しても同じ内容の擬似乱数が生成される。
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	//現在の時刻をシードに使った疑似乱数の生成
	//プログラムの中で毎回異なった疑似乱数生成器のシードを設定するには、現在の時刻を利用する方法が最も手軽。
	//timeパッケージを利用して、現在時刻をもとにしたシード値を設定する。
	//例
	//1970/1/1からの累積ナノ秒をシードに設定
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	//0~99の間の疑似乱数
	fmt.Println(rand.Intn(100))
	//int型の疑似乱数
	fmt.Println(rand.Int())
	//int32型の疑似乱数
	fmt.Println(rand.Int31())
	//int64型の疑似乱数
	fmt.Println(rand.Int63())
	//uint32型の疑似乱数
	fmt.Println(rand.Uint32())
	//特にrand.Intn(n)を使って0以上でnより小さい乱数を生成するパターンが使われる。

}