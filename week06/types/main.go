package main

import (
	"fmt"
	"reflect"
)

func main() {
	//선언 그리고 할당
	//var name string
	//var id int
	//name="Kim Inha"
	//id=1000

	//선언과 할당
	//var name="Kim Inha"
	//var id=1000

	//선언과 할당 (가장 짧은)
	//이녀석도 엄연한 선언이므로, 두번 선언시 에러남.
	name := "Kim Inha"
	id := 1000

	//평점은 소숫점 둘째짜리까지 잘리므로, float64(double)보단 float32
	var gpa float32=3.99

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
	fmt.Println(gpa,reflect.TypeOf(gpa))

	//Zero values
	var f64 float64
	var i16 int16
	var t bool
	var i int
	
	fmt.Println(f64)
	fmt.Println(i16)
	fmt.Println(t)
	fmt.Println(i)
}
