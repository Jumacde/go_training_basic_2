package main

import "fmt"

func main() {
	cal := &calc{}
	cal.SetCalc("Divisiton", 289, 17)
	format, left, right, div_result := cal.calc_div()
	fmt.Printf("\n %s: %d / %d = %d", format, left, right, div_result)

	// score lists
	alice := &school_score{}
	alice.SetSchool_score("Alice", 70, 80, 0)

	bob := &school_score{}
	bob.SetSchool_score("Bob", 90, 85, 0)

	fmt.Println("\n", alice.name, alice.score_list)
	fmt.Println("\n", bob.name, bob.score_list)

	// math.go
	mathe := &mathe{}
	mathe.SetMathe(12, 17, 0)
	x, y, result := mathe.addtion()
	fmt.Printf("\n%d + %d = %d", x, y, result)

}
