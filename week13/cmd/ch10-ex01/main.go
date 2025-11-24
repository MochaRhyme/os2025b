package main

import (
	"fmt"
	"log"
	"week13/pkg/calendar"
)

func main() {
	today := calendar.Date{}
	err:=today.SetYear(2025)
	if err!=nil{
		log.Fatal(err)
	}
	err=today.SetMonth(11)
	// err=today.SetMonth(19)
	if err!=nil{
		log.Fatal(err)
	}
	err=today.SetDay(24)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(today.Year(),"년",today.Month(),"월",today.Day(),"일")
}
