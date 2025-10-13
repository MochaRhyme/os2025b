package main

import (
	"fmt"
	"strings"
	// "reflect"
	"time"
)

func main() {
	// var length float64 = 1.2
	// var width int = 2
	// fmt.Println("면적은", int(length)*width)
	// fmt.Println("length > width?", int(length) > width)
	// fmt.Println("형변환", reflect.TypeOf(int(length)))
	// fmt.Println("원본", reflect.TypeOf(length))

	var now time.Time = time.Now()
	// var year int = now.Year()
	var month time.Month = now.Month()
	var day int = now.Day()
	fmt.Println(month, day)
	univ := "Go$ Inha$"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)
	fmt.Println(changed)
}
