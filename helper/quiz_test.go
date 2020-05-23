package helper

import "testing"

func TestBuild(t *testing.T) {

	q := Quiz{
		Question: "A Question",
		Options: []Option{
			"First",
			"Second",
			"Third",
			"Correct",
		},
		CorrectOption: 3,
	}

	if q.IsCorrect(4) != true {
		t.Errorf("Failed to verify ans. got %T, want true", q.IsCorrect(3))
	}
}
