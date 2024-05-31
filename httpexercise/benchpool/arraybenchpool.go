package benchpool

import (
	"errors"
	"fmt"
	"httpexercise/domain"
	"slices"
)

type ArrayBenchPool struct {
	programmers []domain.Programmer
}

func NewArrayBenchPool() *ArrayBenchPool {
	return &ArrayBenchPool{
		programmers: make([]domain.Programmer, 0, 16),
	}
}

func (pool *ArrayBenchPool) AddProgrammer(p domain.Programmer) {
	pool.programmers = append(pool.programmers, p)
}

func (pool *ArrayBenchPool) ListAllProgrammers() []domain.Programmer {
	return slices.Clone(pool.programmers)
}

func (pool *ArrayBenchPool) SendToWork(language string, minyears int) (domain.Programmer, error) {
	var goesToWork domain.Programmer // the return
	var location int = -1

	// after struct error is defined, add it here. Normally you want to use
	// all the langugaes.
	availableLanguages := []string{"Go", "Java"}
	if !slices.Contains(availableLanguages, language) {
		return domain.Programmer{}, NewInvalidLanguageError(availableLanguages, language)
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
		return domain.Programmer{}, ErrNoProgrammerAvailable
	}
	pool.programmers = slices.Delete(pool.programmers, location, location+1)
	return goesToWork, nil
}

var ErrNoProgrammerAvailable = errors.New("no programmer is available")

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
