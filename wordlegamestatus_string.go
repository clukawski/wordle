// Code generated by "stringer -type=WordleGameStatus"; DO NOT EDIT.

package wordle

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WordleGameStatusPlaying-0]
	_ = x[WordleGameStatusGameOver-1]
	_ = x[WordleGameStatusWon-2]
}

const _WordleGameStatus_name = "WordleGameStatusPlayingWordleGameStatusGameOverWordleGameStatusWon"

var _WordleGameStatus_index = [...]uint8{0, 23, 47, 66}

func (i WordleGameStatus) String() string {
	if i < 0 || i >= WordleGameStatus(len(_WordleGameStatus_index)-1) {
		return "WordleGameStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _WordleGameStatus_name[_WordleGameStatus_index[i]:_WordleGameStatus_index[i+1]]
}