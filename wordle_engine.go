package wordle

import (
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

// WordleEngine represents a Wordle game engine and
// the dictionary and settings for running Wordle games.
type WordleEngine struct {
	Dictionary    []string
	DictionaryMap map[string]struct{}
	MaxAttempts   int
	CurrentGame   *WordleGame
}

// NewWordleEngine constructs a *WordleEngine using the
// dictionary at the path provided, and a mx number of attempts.
//
// The `dictionaryPath` parameter should point to a newline
// terminated list of equal-`rune`-length word strings.
func NewWordleEngine(dictionaryPath string, maxAttempts int) (*WordleEngine, error) {
	dictionary, err := openDictionary(dictionaryPath)
	if err != nil {
		return nil, err
	}

	dictionaryMap := make(map[string]struct{})
	for _, word := range dictionary {
		dictionaryMap[word] = struct{}{}
	}

	return &WordleEngine{
		Dictionary:    dictionary,
		DictionaryMap: dictionaryMap,
		MaxAttempts:   maxAttempts,
	}, nil
}

// NewGame starts instantiates a new *WordleGame when there is no
// previous game in `we.CurrentGame`, or the previous game is over.
func (we *WordleEngine) NewGame() error {
	if we.CurrentGame == nil || we.CurrentGame.Status == WordleGameStatusGameOver {
		word, err := we.RandomDictionaryWord()
		if err != nil {
			return err
		}
		we.CurrentGame = &WordleGame{
			Answer:      []rune(word),
			MaxAttempts: we.MaxAttempts,
			Status:      WordleGameStatusPlaying,
		}
	} else {
		return &WordleError{
			Status: WordleGameStatusPlaying,
		}
	}
	return nil
}

// Fetch a random word from the dictionary in our engine.
//
// Used to start a new game, or can be used to test valid words.
func (we *WordleEngine) RandomDictionaryWord() (string, error) {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(we.Dictionary))))
	if err != nil {
		return "", nil
	}
	return we.Dictionary[i.Int64()], nil
}

// openDictionary reads a newline terminated dictionary file
// containing equal-length words, and returns the words as a
// slice of `string` values.
func openDictionary(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dictionary := strings.Split(string(file), "\n")

	return dictionary, nil
}
