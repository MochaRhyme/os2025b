package main

import "fmt"

type Kilometers float64
type Meters float64
type Miles float64

func (km Kilometers) ToMiles() Miles {
	return Miles(km * 0.621371)
}
func (m Meters) ToMiles() Miles {
	return Miles(m * 0.000621371)
}

func main() {
	kmph := Kilometers(160.1)
	fmt.Printf("%0.3f km/h equals %0.3f mile/h\n", kmph, kmph.ToMiles())
	meters := Meters(160100)
	fmt.Printf("%0.3f meters equals %0.3f miles\n", meters, meters.ToMiles())
}
