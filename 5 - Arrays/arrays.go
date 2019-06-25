package main

import (
	"fmt"
)

func main() {
	// grade1 := 97
	// grade2 := 85
	// grade3 := 93
	// fmt.Printf("Grades: %v, %v, %v \n", grade1, grade2, grade3)

	// -----------------------------------------------------------

	// grades := [3]int{97, 85, 93}
	// fmt.Printf("Grades: %v \n", grades)

	// -----------------------------------------------------------

	// grades := [...]int{97, 85, 93}
	// fmt.Printf("Grades: %v", grades)

	// -----------------------------------------------------------
	// ARRAY INSERT
	var students [3]string
	fmt.Printf("Students: %v \n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v \n", students)

}
