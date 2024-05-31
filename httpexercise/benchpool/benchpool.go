package benchpool

import (
	"httpexercise/domain"
)

// Alright, so any ideas on what to do next?
// ** ask them about ideas and discuss

/*
Okay, let's first define our behaviour of our storage. We want an interface.
We're defining the behaviour we are looking for.
*/
type BenchPool interface {
	AddProgrammer(p domain.Programmer)
	ListAllProgrammers() []domain.Programmer
	// SendToWork(language string, minYears int) Programmer

	// this will be changed during the workshop instead
	SendToWork(language string, minYears int) (domain.Programmer, error)
}
