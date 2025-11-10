package main

import "fmt"

func main() {
	//v1
	// var subjects []string
	// subjects = make([]string, 3)

	//v2
	// subjects := make([]string, 3)
	// subjects[0] = "Go"
	// subjects[2] = "Python"

	//v3
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // slice literal
	subjectSlice:=subjects[1:3] // slicing
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("================")
	for i:=0;i<len(subjectSlice);i++{
		fmt.Println(subjectSlice[i])
	}
}
