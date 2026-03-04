package main

import "fmt"

func main() {
	cal := &calc{}
	cal.SetCalc("Divisiton", 289, 17)
	format, left, right, div_result := cal.calc_div()
	fmt.Printf("\n %s: %d / %d = %d", format, left, right, div_result)
}
