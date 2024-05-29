package main

import "fmt"

/*

Interfaces. Part 1. Interface basics. Boxed type casting.

This is an important part of the workshop. After this chapter, we should probably
take a break.const
- An interface is one of the most important cooncepts in Go.
- Let's start by writing the most basic one, hello wordable, that says hello
world.
*/

type HelloWorldable interface {
	// These are the interface set of methods. When implementing them, you
	// have to use the exact same signature.
	HelloWorld()

	// ** add this later to your code during the internship.
	// GetHelloMessage() string
}

// Now let's define a new type which implements that interface.
type RomanianHelloWorlder struct {
}

// And let's implement the hello world method
func (r RomanianHelloWorlder) HelloWorld() {
	fmt.Println("salut lume!")
}

// Add this a bit later, once you explain boxing
func (r RomanianHelloWorlder) CeMaiFaci() {
	fmt.Println("ce mai faci?")
}

// Okay, let's see this in action
func main() {
	var helloworder HelloWorldable
	helloworder = RomanianHelloWorlder{}

	// Show them that this works.
	// At this point, it's extremely important to explain them that the runtime
	// type of helloworder is actually HelloWordable, and not RomanianHelloWorlder !!
	// The value is boxed inside the interface, and is of that specific type.
	// Which can change!!
	helloworder.HelloWorld()

	// ** next example is to add a method to Romanian and see that the method
	// does not show up
	// helloworder.CeMaiFaci() // THIS DOES NOT WORK!

	// We can change the underlying value to get a new output! Keep in mind that
	// helloworder is still of the interface type.
	helloworder = EnglishHelloWorlder{}
	helloworder.HelloWorld()

	// ** This is essentially Go's way of doing polymorphism. It's also the
	// only way to store values of different "types" (they're the same type,
	// just the underlying type is different). However, you can cast to the
	// original type if you want to, by doing a type assertion.

	ehw, ok := helloworder.(EnglishHelloWorlder)
	fmt.Println(ehw, ok)

	// Let's do the same for romanian, and see what happens
	rhw, ok := helloworder.(RomanianHelloWorlder)
	fmt.Println(rhw, ok)

	// Here's a way in which we can do it, based on all types that implement
	// it.

	// ** this is the point where you should probably introduce switch
	// ** at this point, add something in the struct so that they print out
	// nicely.
	switch helloworder.(type) {
	case EnglishHelloWorlder:
		fmt.Println("english")
	case RomanianHelloWorlder:
		fmt.Println("romanian")
	}

}

// Now let's define a new type which implements that interface.
type EnglishHelloWorlder struct {
}

// And let's implement the hello world method
func (r EnglishHelloWorlder) HelloWorld() {
	fmt.Println("hello world!")
}
