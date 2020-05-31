package helper

import (
	"testing"
)

func TestBuild(t *testing.T) {

	q := (Quiz{}).Build(
		QuizConfig{
			PerStage: 3,
		},
		QuizStat{},
		[]Quiz{
			Quiz{
				ID:            1,
				Question:      "2*3",
				Options:       []Option{"4", "5", "6", "7"},
				CorrectOption: 2,
			},
			Quiz{
				ID:            2,
				Question:      "2+3",
				Options:       []Option{"4", "5", "6", "7"},
				CorrectOption: 1,
			},
			Quiz{
				ID:            3,
				Question:      "2-3",
				Options:       []Option{"1", "-1", "0", "5"},
				CorrectOption: 1,
			},
			Quiz{
				ID:            4,
				Question:      "6/2",
				Options:       []Option{"1", "2", "3", "4"},
				CorrectOption: 2,
			},
			Quiz{
				ID:            5,
				Question:      "123-1",
				Options:       []Option{"122", "123", "125", "129"},
				CorrectOption: 0,
			},
		})

	if q.IsCorrect(3) != true {
		t.Errorf("Failed to verify ans. got %T, want true", q.IsCorrect(3))
	}

	t.Run("Chain", func(t *testing.T) {

		if q.Next.Prev != q {
			t.Errorf("Next's Prev is not Current")
		}

		if q.Prev.Next != q {
			t.Errorf("Prev's Next is not Currect")
		}
	})

	t.Run("Advancing", func(t *testing.T) {
		n := q.Advance()

		if q.Next != n {
			t.Errorf("Currect quiz is not Next of old quiz")
		}

	})

	t.Run("Masking", func(t *testing.T) {
		q.Mask()

		// removes current quiz from chain
		if q.Prev.Next != q.Next {
			t.Errorf("Prev's Next is not Next")
		}

		if q.Next.Prev != q.Prev {
			t.Errorf("Next's Prev is not Prev")
		}
	})

	t.Run("Mastering", func(t *testing.T) {

		t.Run("Removes current Quiz from the chain", func(t *testing.T) {

			q.Master()

			if q.Next.Prev == q || q.Prev.Next == q {
				t.Errorf("Current quiz is not removed from the chain")
			}

			if q.Next.Prev != q.Prev.Next {
				t.Errorf("Quiz chain is broken")
			}

		})
	})
}
