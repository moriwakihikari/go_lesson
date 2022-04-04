package main

import "fmt"

//init

//packageより先に走る
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Main")
}