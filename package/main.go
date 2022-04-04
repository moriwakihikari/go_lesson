package main

import (
	"fmt"
	"os"
)



func main() {
	//Exit
	// os.Exit(1)
	// fmt.Println("Start")

	// defer func() {
	// 	fmt.Println("defer")
	// } ()
	// os.Exit(0)

	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// fmt.Printf("length=%d\n", len(os.Args))
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }

	// f, err := os.Open("test.log")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	//ファイル作成
	// f, _ := os.Create("foo.txt")
	// f.Write([]byte("Hello\n"))
	// f.WriteAt([]byte("Golang"), 7)
	// f.Seek(0, os.SEEK_END)
	// f.WriteString("Yaah")
	// f, _ = os.Create("foo.txt")

	f, err := os.Open("foo.txt")
	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n, err)
	fmt.Println(string(bs))

	nn, err := f.ReadAt(bs, 10)
	fmt.Println(nn)
	fmt.Println(string(bs))


}