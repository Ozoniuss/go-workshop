package main

import (
	"errors"
	"fmt"
	"slices"
)

/*
Practice exercise. Let's implement something using everything we learnt
so far.

Assume you want to start your own company in Cluj. So you're obviously doing
outsourcing. Which means you need to keep track of the number of programmers
you have. Basically we're looking for the following operations:

- Adding a new programmer
- Listing all programmers in an array (for tax purposes)
- Send the top programmer to work (we can't pay bench forever)

You want to make two implementations for this. In practice, at Snyk we typically
do a "real" implementation and a "fake" implementation of something. We'll
explain why later, if that won't be clear from the exercise.
*/

// Let's define the struct.

type Programmer struct {
	Name     string
	Language string
	YoE      int
}

// Alright, so any ideas on what to do next?
// ** ask them about ideas and discuss

/*
Okay, let's first define our behaviour of our storage. We want an interface.
We're defining the behaviour we are looking for.
*/
type BenchPool interface {
	AddProgrammer(p Programmer)
	ListAllProgrammers() []Programmer
	// SendToWork(language string, minYears int) Programmer

	// this will be changed during the workshop instead
	SendToWork(language string, minYears int) (Programmer, error)
}

// Now let's discuss a bit about potential implementations. Probably there isn't
// going to be an implementation that is best suited for all operations. So let's
// see the options we have.

// An array bench pool only stores programmers in an array, and nothing more.
type ArrayBenchPool struct {
	// programmers should be private. no need to access it directly
	programmers []Programmer
}

// This is ideal for the first and second operation. Why?

// Now, the very first thing: we're going to need to initialize this pool
// somehow right? Remember that arrays must be initialzied. Let's define
// an idiomatic constructor.
// Return a pointer because that's how the object is meant to be used so other
// programmers should be aware of that.
func NewArrayBenchPool() *ArrayBenchPool {
	return &ArrayBenchPool{
		// 16 is completely random
		programmers: make([]Programmer, 0, 16),
	}
}

// Use a pointer receiver!! This is going to be an object that we'll be sharing
// around probably. There's no point in initiating two pools. And we want to
// mutate it therefore.
func (pool *ArrayBenchPool) AddProgrammer(p Programmer) {
	pool.programmers = append(pool.programmers, p)
}

// Clone here. We don't want to modify the original array do we?
func (pool *ArrayBenchPool) ListAllProgrammers() []Programmer {
	return slices.Clone(pool.programmers)
}

// Tests are passing. Let's do the second implementation now. What is another
// data structure that we could use here?
type MapBenchPool struct {
	// yes there is duplication, and we could remove it to save memory.
	// but that doesn't matter rn
	programmers map[string][]Programmer
}

// Now let's follow the exact same steps as before.
func NewMapBenchPool() *MapBenchPool {
	return &MapBenchPool{
		programmers: make(map[string][]Programmer),
	}
}

// And let's implement the first two methods
func (pool *MapBenchPool) AddProgrammer(p Programmer) {
	if _, ok := pool.programmers[p.Language]; ok {
		pool.programmers[p.Language] = append(pool.programmers[p.Language], p)
	} else {
		pool.programmers[p.Language] = []Programmer{p} // you can also initialize array like this
	}
}

// And now the second method.
func (pool *MapBenchPool) ListAllProgrammers() []Programmer {
	allProgrammers := make([]Programmer, 0)
	for _, listOfProgrammers := range pool.programmers {
		// add all the programmers of that language
		allProgrammers = append(allProgrammers, listOfProgrammers...)
	}
	return allProgrammers
}

// Nice job! Time to write the tests.
func (pool *MapBenchPool) SendToWork(language string, minYears int) (Programmer, error) {
	return Programmer{}, nil
}

// Ok this seems alright. But let's set up some tests now so we can test
// our functionality.
// ** write the test first
// func (pool *ArrayBenchPool) SendToWork(language string, minyears int) (Programmer, error) {
// 	var goesToWork Programmer // the return
// 	var location int = -1
// 	for idx, programmer := range pool.programmers {
// 		if programmer.Language != language {
// 			continue
// 		}
// 		if programmer.YoE >= minyears {
// 			goesToWork = programmer
// 			location = idx
// 			break
// 		}
// 	}
// 	if location == -1 {
// 		return Programmer{}, nil
// 	}
// 	pool.programmers = slices.Delete(pool.programmers, location, location+1)
// 	return goesToWork, nil
// }

// Test should be green. Once test is green, time to introduce errors.

// After going through the test and adding the signature, let's add a new error.

// Explain a bit about errors and error handling in Go, in general.

// Sentinel error.
var ErrNoProgrammerAvailable = errors.New("no programmer is available")

func (pool *ArrayBenchPool) SendToWork(language string, minyears int) (Programmer, error) {
	var goesToWork Programmer // the return
	var location int = -1

	// after struct error is defined, add it here. Normally you want to use
	// all the langugaes.
	availableLanguages := []string{"Go", "Java"}
	if !slices.Contains(availableLanguages, language) {
		return Programmer{}, NewInvalidLanguageError(availableLanguages, language)
	}

	for idx, programmer := range pool.programmers {
		if programmer.Language != language {
			continue
		}
		if programmer.YoE >= minyears {
			goesToWork = programmer
			location = idx
			break
		}
	}
	if location == -1 {
		// Add the error here.
		return Programmer{}, ErrNoProgrammerAvailable
	}
	pool.programmers = slices.Delete(pool.programmers, location, location+1)
	return goesToWork, nil
}

// Finally, since Error is an interface, errors can also be types. This is in
// case you want to store contextual information.
type InvalidLanguageError struct {
	availableLanguages []string
	language           string
}

func NewInvalidLanguageError(available []string, language string) InvalidLanguageError {
	return InvalidLanguageError{
		availableLanguages: available,
		language:           language,
	}
}

// Implement the error interface.
func (err InvalidLanguageError) Error() string {
	return fmt.Sprintf("language %s is not available. we only offer %v", err.language, err.availableLanguages)
}

// Note how this would be very nice to do with the map implementation.
