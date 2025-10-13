package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "reflect"
)

func main() {
	// var length float64 = 1.2
	// var width int = 2
	// fmt.Println("면적은", int(length)*width)
	// fmt.Println("length > width?", int(length) > width)
	// fmt.Println("형변환", reflect.TypeOf(int(length)))
	// fmt.Println("원본", reflect.TypeOf(length))

	// var now time.Time = time.Now()
	// // var year int = now.Year()
	// var month time.Month = now.Month()
	// var day int = now.Day()
	// fmt.Println(month, day)
	// univ := "Go$ Inha$"
	// changer := strings.NewReplacer("$", "!")
	// changed := changer.Replace(univ)
	// fmt.Println(changed)

	// r := bufio.NewReader(os.Stdin)
	// i, err := r.ReadString('\n')
	// // fmt.Println(err)
	// if err != nil {
	// 	log.Fatal(err) //report the error, exit the program
	// }
	// fmt.Println(i)

	//엉망진창 변수선언 - 무조건 에러!
	// var int int = 99
	// var b int = 8
	// var fmt string="inha"
	// fmt.Println(int,b)
	// fmt.Println()

	fmt.Print("Enter a score: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // 문자열 주위 띄어쓰기 탭키 등등 제거
	grade, err := strconv.ParseFloat(input, 64) // 정리된 문자열을 실수타입으로 변환
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
