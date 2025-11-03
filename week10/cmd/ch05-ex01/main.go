package main

import (
	"fmt"
	// "reflect"
)

func main() {
	// arrayInt := [3]int{-9, 11, 7}
	// for i := 0; i < len(arrayInt); i++ {
	// 	fmt.Println(i, arrayInt[i])
	// }
	arrayInt := [3]int{-9, 11, 7}
	for i,n := range arrayInt {
		fmt.Println(i, n)
	}
}
