package main

import (
	"fmt"
)

func swap(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp
	fmt.Println(*first, *second)
}

func main() {
	// fmt.Println(math.Sqrt(-144)) 허수 그딴거 없음. NaN이 출력됨.
	a, b := 10, 20
	fmt.Println(a, b)
	swap(&a, &b) // pass by address
	fmt.Println(a, b)
}
