package main

import (
	"fmt"
)

// first discuss arrays, which are of fixed length.

// then discuss slices, which build on top of arrays. len vs capacity. introduce
// make.

// discuss maps. (maybe say Keys have to be comparable? they would discover
// this anyway, and you don't want information overload. Explain perhaps if
// they ask only.)

// structs. embedded structs

/*
Let's introduce arrays.
*/

func example3a() {
	// array is initialized, with 0 values.
	var arr [5]int
	// arrays are values in Go! they're like string, int, etc.
	fmt.Println(arr)

	// however, they are a mutable container type. Even if they are values,
	// you can change them.

	arr[1] = 10
	fmt.Println(arr)

	// Arrays have a fixed sized length, which you can't change. you can use
	// these functions to get length and capacity.
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

}

/*
Now, let's introduce slices, which are based on arrays. A slice is a dynamic
array, that is backed up by a real array.
This is visible in the runtime package, slice.go
*/

func example3b() {
	// make an array, with length and capacity. lenght is the number of (zeroed)
	// elements, capacity is more or less an internal detail used to resize
	// the array.
	arr := make([]int, 0, 10)
	fmt.Println(arr, len(arr), cap(arr))

	// array with 5 zeroes
	arr2 := make([]int, 5, 10)
	fmt.Println(arr2, len(arr2), cap(arr2))

	// shorthand initialization. This is a common error source!!!!
	arr3 := make([]int, 6)
	fmt.Println(arr3, len(arr3), cap(arr3))

	// shorthand initialization v2
	arr4 := []int{1, 2, 3, 4}
	fmt.Println(arr4, len(arr4), cap(arr4))
}

/* Array basics. Append */
func example3c() {
	arr := make([]int, 0, 10)

	// Runtime panic!!
	// arr[1] = 10

	// This is how elements are appended to the array.
	arr = append(arr, 10, 10, 10, 10)
	arr[2] = 17
	fmt.Println(arr)

	// uninitialized array, nil check.
	var arr2 []int
	fmt.Println(arr2 == nil)
	arr2 = append(arr2, 10)
	fmt.Println(arr2)

	// empty array
	// https://stackoverflow.com/questions/44305170/nil-slices-vs-non-nil-slices-vs-empty-slices-in-go-language
	var arr3 = make([]int, 0)
	fmt.Println(arr3 == nil)
}

/*
Let's do some very basic maps now.
*/

func example3d() {
	// maps are useful for fast element access and insertions.
	var mymap = make(map[int]int, 0)
	mymap[10] = 4
	fmt.Println(mymap)

	mymap[9] = 3
	mymap[6] = 4

	// This is what you will likely always use
	var mymap2 = make(map[int]int)
	mymap2[10] = 100

	// Keys need to be comparable. That is, they can't be slices, maps, functions
	// or any struct containing an uncomparable field. For example, this can
	// be used to store coordinates.
	var mymap3 = make(map[[2]int]string)
	mymap3[[2]int{0, 0}] = "O"
}

/*
Custom types (or, rather named types) and structs
*/

// type definition (supports defining new types over it, we'll see later)
type LocationDefinition [2]int

// type alias (reference the exact type)
type LocationAlias = [2]int

func example3e() {
	var loc LocationDefinition = [2]int{1, 2}
	fmt.Println(loc)
}

// Add SET example

/*
Structs. struct basics. Introduce struct. We'll be using this a lot.
Very similar with C struct
*/

type SomeStruct struct {
	Name string
	Age  int
}

func example3f() {
	s1 := SomeStruct{}
	fmt.Println(s1)

	s2 := SomeStruct{
		Name: "Alex",
		Age:  23,
	}
	fmt.Println(s2)
}

/*
Structs within structs
*/

type SomeOtherStruct struct {
	Person SomeStruct
	Job    string
}

// embeddding is essentially inheritance
type SomeOtherStructEmbedded struct {
	SomeStruct
	Job string
}

func example3g() {
	alex := SomeStruct{
		Name: "Alex",
		Age:  23,
	}
	s1 := SomeOtherStruct{
		Job: "programmer",
		Person: SomeStruct{
			Name: "Alex",
			Age:  23,
		},
	}
	fmt.Println(s1)
	fmt.Printf("%v %+v", s1, s1)

	s2 := SomeOtherStruct{
		Job:    "programmer",
		Person: alex,
	}
	fmt.Printf("%v %+v", s2, s2)

	s3 := SomeOtherStructEmbedded{
		Job:        "programmer",
		SomeStruct: alex,
	}
	fmt.Printf("%v %+v", s3, s3)
	fmt.Println(s3.SomeStruct.Age, s3.Age)
}

func main() {
	example3a()
	example3b()
	example3c()
	example3d()
	example3e()
	example3f()
}
