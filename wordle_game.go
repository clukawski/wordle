package wordle

import "strings"

// WordleGame represents a wordle game state, keep of:
// - Game Status
// - The correct result, represented as a rune slice
// - Maximum number of Attempts
// - A list of previous attempts
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

// WordleError represents an error response used by
// [wordle.WordleGame], containing the current game status.
type WordleError struct {
	Status WordleGameStatus
}

// Error implements the requirements the error interface
func (e *WordleError) Error() string {
	return e.Status.String()
}

// GetAnswer returns a string representation of the correct
// answer of the current game
func (wg *WordleGame) GetAnswer() string {
	return string(wg.Answer)
}

// Attempt validates and records an attempt using the provided word,
// and returns whether or not the attempt was correct.
//
// The game status is updated when a win (correct guess) or game over
// condition (last remaining guess is incorrect) is met.
//
// An error is returned if the game state is [wordle.WordleStateGameOver], or
// the last remaining guess was incorrect.
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
		if wa.numberOfFinds(char) < strings.Count(string(wg.Answer), string(char)) {
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
