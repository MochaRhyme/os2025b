package main

import (
	"fmt"
)

func main() {
	// 스택 형태 - 즉 후입선출(LIFO)로 동작함.
	defer fmt.Println("1st defer")
	defer fmt.Println("2st defer")
	defer fmt.Println("3st defer")
	fmt.Println("main!")
}
