package main

import "fmt"

// var i float32 = 42
// var i = 42.
// var i float32 = "foo" -> "WRONG FORMAT"!
// var i int = 27

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
	// var a int = 10
	// var b int8 = 3
	// fmt.Println(a + int(b))

	// OPERATORS AND, OR, NOT, XOR
	// a := 10
	// b := 3

	// fmt.Println(a & b)
	// fmt.Println(a | b)
	// fmt.Println(a ^ b)
	// fmt.Println(a &^ b)

	// HIGH NUMBERS
	// a := 8              //2^3
	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0

	// FLOAT POINTS
	// n := 3.14
	// n = 13.7e72
	// n = 2.1E14
	// fmt.Printf("%v, %T", n, n)

	// COMPLEX NUMBERS
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T", n, n)

	// var n complex128 = complex(5, 12)
	// fmt.Printf("%v, %T", n, n)

}
