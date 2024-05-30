package main

import (
	"errors"
	"testing"
)

func TestAddProgrammer(t *testing.T) {
	p := Programmer{
		Name:     "Alex",
		Language: "Go",
		YoE:      10, // definitely not lying on the resume here
	}

	// create a pool and add the programmer
	pool := NewArrayBenchPool()
	pool.AddProgrammer(p)

	// let's get the programmer
	programmers := pool.ListAllProgrammers()

	if len(programmers) != 1 {
		// the other test would panic if the programmer is not there, so
		// stop here
		t.Fatalf("expected %d programmers, got %d\n", 1, len(programmers))
	}

	actualProgrammer := programmers[0]
	if p != actualProgrammer {
		t.Errorf("expected to contain %v, got %v\n", p, actualProgrammer)
	}
}

/*
I want you to notice that we dind't really make use of the actual Pool type.
We only used exposed methods. This is generally nice, because in this unit
test, we don't particualrly care about implementation details. We want to
make sure that using the public methods, we get what we wanted.

This is however, not always possible, as there are certain cases where you
can't really test this way. For example, if we didn't expose a List method,
we'd need to define some custom method for testing. Just keep this in mind.

But, since we only used the interface methods, we can also do the following:
*/

// Test_AddProgrammer_ProgrammerIsAdded
// test names are not a Go convention btw
func TestAddProgrammerInterface(t *testing.T) {
	p := Programmer{
		Name:     "Alex",
		Language: "Go",
		YoE:      10, // definitely not lying on the resume here
	}

	// create a pool and add the programmer, AS INTERFACE now. this will
	// be useful. Continue by implemnenting the other approach.
	// ** go in the other file and implement a dummy method so it
	// safisfies the interface
	//
	// ** then rename test to Test_AddProgrammer_ProgrammerIsAdded
	var pool BenchPool
	pool = NewArrayBenchPool()
	pool.AddProgrammer(p)

	// let's get the programmer
	programmers := pool.ListAllProgrammers()

	if len(programmers) != 1 {
		// the other test would panic if the programmer is not there, so
		// stop here
		t.Fatalf("expected %d programmers, got %d\n", 1, len(programmers))
	}

	actualProgrammer := programmers[0]
	if p != actualProgrammer {
		t.Errorf("expected to contain %v, got %v\n", p, actualProgrammer)
	}
}

// Once you write this with interface, go back and implement the second method.

// The other implementation is now complete. Let's write the tests.
// Now, we're going to notice the interface coming into handy. Let's copy
// the first test.

// Notice we only had to change the constructor!!!!

func TestAddProgrammerMap(t *testing.T) {
	p := Programmer{
		Name:     "Alex",
		Language: "Go",
		YoE:      10, // definitely not lying on the resume here
	}

	// create a pool and add the programmer, AS INTERFACE now. this will
	// be useful. Continue by implemnenting the other approach.
	// ** go in the other file and implement a dummy method so it
	// safisfies the interface
	//
	// ** then rename test to Test_AddProgrammer_ProgrammerIsAdded
	var pool BenchPool
	pool = NewMapBenchPool()
	pool.AddProgrammer(p)

	// let's get the programmer
	programmers := pool.ListAllProgrammers()

	if len(programmers) != 1 {
		// the other test would panic if the programmer is not there, so
		// stop here
		t.Fatalf("expected %d programmers, got %d\n", 1, len(programmers))
	}

	actualProgrammer := programmers[0]
	if p != actualProgrammer {
		t.Errorf("expected to contain %v, got %v\n", p, actualProgrammer)
	}
}

// So now, let's make use of the two implementations to make a very nice
// test. It's called table driven tests.
// https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

// Remove both tests, and add this new one. Name it properly.
func Test_AddProgrammer_ProgrammerIsAdded(t *testing.T) {

	// Define the data taht each tests holds.
	type testcase struct {
		pool BenchPool
	}
	testcases := []testcase{
		{
			pool: NewArrayBenchPool(),
		},
		{
			pool: NewMapBenchPool(),
		},
	}
	// And now let's test each individual pool.
	for _, tt := range testcases {

		p := Programmer{
			Name:     "Alex",
			Language: "Go",
			YoE:      10, // definitely not lying on the resume here
		}
		tt.pool.AddProgrammer(p)

		programmers := tt.pool.ListAllProgrammers()

		if len(programmers) != 1 {

			t.Fatalf("expected %d programmers, got %d\n", 1, len(programmers))
		}

		actualProgrammer := programmers[0]
		if p != actualProgrammer {
			t.Errorf("expected to contain %v, got %v\n", p, actualProgrammer)
		}
	}
}

// This is the power of table driven tests.

// Now, let's implement the last method in a TDD fashion. When I got hired
// at Snyk, one of the first things they taught me was this approach. It's
// called Red, Green, Refactor.
//
// First, write a test that's red, make it green in the simplest way possible,
// maybe even wrong way, then refactor it to be nicely. That may mean making it
// red again, green again, then refactor. Let's put that in practice.
//
// Write all 3 tests at the beginning, to set the scene of how test names
// should look like.

func Test_SendToWork_HasProgrammer_ThatProgrammerIsSentToWork(t *testing.T) {
	alex := Programmer{
		Name:     "Alex",
		Language: "Go",
		YoE:      10, // definitely not lying on the resume here
	}
	armand := Programmer{
		Name:     "Armand",
		Language: "Go",
		YoE:      1,
	}
	// ask from the crowd
	ioana := Programmer{
		Name:     "Ioana",
		Language: "Java", // :(
		YoE:      2,      // definitely not lying on the resume here
	}

	var pool BenchPool
	pool = NewArrayBenchPool()

	// Now we're adding all programmers
	pool.AddProgrammer(armand) // armand first
	pool.AddProgrammer(alex)
	pool.AddProgrammer(ioana)

	// let's keep the original test, just to have an extra guard
	programmers := pool.ListAllProgrammers()

	if len(programmers) != 3 {
		t.Fatalf("expected %d programmers, got %d\n", 3, len(programmers))
	}

	// expecting to send alex
	working, _ := pool.SendToWork("Go", 5)
	if working != alex {
		t.Errorf("expected to send %v, got %v\n", alex, working)
	}

	// the programmer should be out of the pool
	programmers = pool.ListAllProgrammers()
	if len(programmers) != 2 {
		t.Fatalf("expected %d programmers, got %d\n", 2, len(programmers))
	}
	// Now run the test. It should fail!!
	// Run only this test. explain go help test and go help testflag
	// go test -run=HasProgrammer 8.go 8_test.go

	// Now the test is failed, let's do the implementation
}

// Now, you may notice that testing is a bit awkward for the following tests.

func Test_SendToWork_NotEnoughYoe_NoProgrammerIsSentToWork(t *testing.T) {
	// Let's remove alex from the test.

	armand := Programmer{
		Name:     "Armand",
		Language: "Go",
		YoE:      1,
	}
	ioana := Programmer{
		Name:     "Ioana",
		Language: "Java",
		YoE:      2,
	}

	var pool BenchPool
	pool = NewArrayBenchPool()

	// Remove alex
	pool.AddProgrammer(armand)
	pool.AddProgrammer(ioana)

	// So now let's send a Go programmer with 5 years of experience. But
	// we don't have one:(
	_, err := pool.SendToWork("Go", 5)
	// How do we test?
	// Discussion about ways to implement this. But, to make it work nicely,
	// introdce errors. Change the signature to accept an error.
	if !errors.Is(err, ErrNoProgrammerAvailable) {
		t.Errorf("expected error %v, got %v instead\n", ErrNoProgrammerAvailable, err)
	}
}

// Let's write a quick additional test. We want an error if there is no
// language.
func Test_SendToWork_NoExistingLanguage_NoProgrammerIsSentToWork(t *testing.T) {
	// Let's remove alex from the test.

	armand := Programmer{
		Name:     "Armand",
		Language: "Go",
		YoE:      1,
	}
	ioana := Programmer{
		Name:     "Ioana",
		Language: "Java",
		YoE:      2,
	}

	var pool BenchPool
	pool = NewArrayBenchPool()

	pool.AddProgrammer(armand)
	pool.AddProgrammer(ioana)

	_, err := pool.SendToWork(".NET", 5)
	// How do we test this error?.

	// Now that we have the new language error, we can use a different
	// way of checking errors.
	var langerr InvalidLanguageError
	if !errors.As(err, &langerr) {
		t.Errorf("expected language error %v, got %v instead\n", ErrNoProgrammerAvailable, err)
	}
}
