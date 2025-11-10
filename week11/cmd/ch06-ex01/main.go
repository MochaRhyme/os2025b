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
	subjects := []string{"Go", "", "Python"} // slice literal
	for _, subject := range subjects {
		fmt.Println(subject)
	}
}
