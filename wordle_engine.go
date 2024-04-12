package wordle

import (
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

type WordleEngine struct {
	Dictionary    []string
	DictionaryMap map[string]struct{}
	MaxAttempts   int
	CurrentGame   *WordleGame
}

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

func (we *WordleEngine) RandomDictionaryWord() (string, error) {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(we.Dictionary))))
	if err != nil {
		return "", nil
	}
	return we.Dictionary[i.Int64()], nil
}

func openDictionary(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dictionary := strings.Split(string(file), "\n")

	return dictionary, nil
}
