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
	subjectSlice := subjects[:3]                                 // slicing
	subjects[0] = "Java"
	subjectSlice = append(subjectSlice, "Go") // 길이가 원본 배열 길이를 배반하지 않은 경우 원본에 덮어씌워짐.
	// 배반한 경우 덮어씌워지지 않고, 원본은 그대로 따로 감.
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("================")
	for i := 0; i < len(subjectSlice); i++ {
		fmt.Println(subjectSlice[i])
	}
}
