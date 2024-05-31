package main

import "fmt"

/*
At this point we probably have enough to dive into functions. The following are
particularly interesting:

- basics of parameters and return values
- functions defined on structs
*/

/*
Write a very basic function, that takes in a few arguments and returns some
other arguments.
*/

func division(dividend int, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor

	return quotient, remainder
}

// Note that if consecutive arguments have the same type, we can leave the
// type at the end.
func division2(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor

	return quotient, remainder
}

// We can also have named return types, alongisde a naked return. The return
// values are initialized with the zero values.
// Named return types.
func division3(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor

	return
}

// Functions are first class citizens and can be stored in variables.
func example1() {
	div := division
	fmt.Println(div(7, 3))
}

// Functions can be passed within functions. A function type is essentially
// it's signature
func divideNumbers(a, b int, divisionFunction func(int, int) (int, int)) {
	fmt.Println(divisionFunction(a, b))
}

func example2() {
	divideNumbers(7, 3, division)
}

// Functions can be defined on new types. Very important for interfaces.
type Location [2]int

func (l Location) SumOfCoordinates() int {
	return l[0] + l[1]
}

type Myself struct {
	Name string
	Age  int
}

func (m Myself) PrintInfo() {
	fmt.Println(m.Name, m.Age)
}

func main() {
	example1()
	example2()
}
