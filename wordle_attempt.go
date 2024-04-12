package wordle

// WordleAttempt represents an attempt to guess the
// current game answer, and the status of each character
// as represented by the [wordle.CharacterStatus] enum.
type WordleAttempt struct {
	Guess  []rune
	Result []CharacterStatus
}

// CharacterStatus is an enum for character status in a guessed word
type CharacterStatus int

const (
	CharacterStatusNotPresent = iota
	CharacterStatusWrongLocation
	CharacterStatusCorrectLocation
)

func (wa *WordleAttempt) numberOfFinds(letter rune) int {
	count := 0
	for i, r := range wa.Guess {
		if r == letter {
			if wa.Result[i] == CharacterStatusCorrectLocation || wa.Result[i] == CharacterStatusWrongLocation {
				count++
			}
		}
	}
	return count
}
