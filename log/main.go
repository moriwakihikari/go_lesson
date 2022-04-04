package main

import (
	"log"
	"os"
)

func main() {
	//ログの出力先を変更
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	// log.Fatal("Log\n")
	// log.Fatalln("Log2\n")
	// log.Fatalf("Log3\n")

	// log.Panic("Log\n")
	// log.Panicln("Log2\n")
	// log.Panicf("Log3\n")

	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("ファイルに書き込む")

	log.SetFlags(log.LstdFlags)
	log.Println("A")

	//マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	//時刻とファイルの行番号
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	log.SetFlags(log.Ltime | log.Llongfile)
	log.Println("D")

	log.SetFlags(log.LstdFlags)

	log.SetPrefix("[LOG]")
	log.Println("E")

	//ロガーの生成
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Fatalln("message")

	_, err = os.Open("fdafdsafa")
	if err != nil {
		log.Fatalln("Exit", err)
	}

}
