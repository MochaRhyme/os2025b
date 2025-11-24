package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2006)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(4)
	// err=event.SetMonth(19)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(30)
	if err != nil {
		log.Fatal(err)
	}
	event.SetTitle("생년월일")
	fmt.Println(event.Title())
	fmt.Println(event.Year(), "년", event.Month(), "월", event.Day(), "일")
}
