package main

import "fmt"

// var i float32 = 42
// var i = 42.
// var i float32 = "foo" -> "WRONG FORMAT"!
var i int = 27

func main() {
	// var i int
	// i = 42
	// var j float32 = 27
	// k := 99.
	// fmt.Printf("%v, %T", i, i)
	// var theHttpRequest string = "https://google.com"

	// var i int = 42
	// fmt.Printf("%v, %T \n", i, i)

	// WORK WITH STRING
	// var j string
	// j = strconv.Itoa(i)
	// fmt.Printf("%v, %T \n", j, j)

	// OPERATORS
	// a := 10
	// b := 3
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)

	// DIFFERENT TYPE OF INT
	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

}
