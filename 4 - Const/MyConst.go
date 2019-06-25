package main

import "fmt"

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// const (
// 	a2 = iota
// )
const (
	catSpecialist = iota + 5
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int = catSpecialist
	fmt.Printf("%v \n", specialistType == catSpecialist)

	fmt.Printf("%v", dogSpecialist)

	// fmt.Printf("%v \n", a)
	// fmt.Printf("%v \n", b)
	// fmt.Printf("%v \n", c)
	// fmt.Printf("%v \n", a2)

	// const myConst int = 42
	// myConst int = 27

	// const myConst float64 = math.Sin(1.57) // IS NOT A CONSTANT
	// fmt.Printf("%v, %T", myConst, myConst)

	// const a int = 14
	// const b string = "foo"
	// const c float32 = 3.14
	// const d bool = true

	// fmt.Printf("%v, %T \n", a, a)
	// fmt.Printf("%v, %T \n", b, b)
	// fmt.Printf("%v, %T \n", c, c)
	// fmt.Printf("%v, %T \n", d, d)

	// const a int = 42
	// var b int = 27
	// fmt.Printf("%v, %T \n", (a + b), (a + b))
}
