package wordle

import "strings"

type WordleGame struct {
	Status      WordleGameStatus
	Answer      []rune
	MaxAttempts int
	Attempts    []*WordleAttempt
}

// WordleGameStatus is an enum for game status
//
//go:generate stringer -type=WordleGameStatus
type WordleGameStatus int

const (
	WordleGameStatusPlaying WordleGameStatus = iota
	WordleGameStatusGameOver
	WordleGameStatusWon
)

type WordleError struct {
	Status WordleGameStatus
}

func (e *WordleError) Error() string {
	return e.Status.String()
}

func (wg *WordleGame) GetAnswer() string {
	return string(wg.Answer)
}

func (wg *WordleGame) Attempt(word string) (bool, error) {
	if wg.Status == WordleGameStatusGameOver {
		return false, &WordleError{
			Status: WordleGameStatusGameOver,
		}
	}

	wa := &WordleAttempt{
		Guess: []rune(word),
	}
	wa.Result = make([]CharacterStatus, 5)

	for i, char := range wa.Guess {
		charStr := string(char)
		if !strings.Contains(string(wg.Answer), charStr) {
			wa.Result[i] = CharacterStatusNotPresent
			continue
		}
		if wa.Guess[i] == wg.Answer[i] {
			wa.Result[i] = CharacterStatusCorrectLocation
			continue
		}
	}

	for i, char := range wa.Guess {
		if wa.NumberOfFinds(char) < strings.Count(string(wg.Answer), string(char)) {
			if wa.Result[i] != CharacterStatusCorrectLocation {
				wa.Result[i] = CharacterStatusWrongLocation
			}
		}
	}

	wg.Attempts = append(wg.Attempts, wa)

	for _, status := range wa.Result {
		if status != CharacterStatusCorrectLocation {
			if len(wg.Attempts) == wg.MaxAttempts {
				wg.Status = WordleGameStatusGameOver
				return false, &WordleError{
					Status: WordleGameStatusGameOver,
				}
			}
			return false, nil
		}
	}
	wg.Status = WordleGameStatusWon

	return true, nil
}
