package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	
}

func main() {
	start := time.Now()
	go say("고루틴")
	say("메인")
	fmt.Println("전체 실행 시간 : ",time.Since(start))
}
