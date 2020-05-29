package main

import (
	"fmt"

	"github.com/hive-repo/quiz/helper"
)

func main() {

	q := (helper.Quiz{}).Build()

	for {

		q.DisplayStat()

		q.Display()

		ans := q.PromptAns()

		if q.IsCorrect(ans) {
			fmt.Printf("\nCorrect!\n\n")
		} else {
			fmt.Printf("\nIncorrect! Correct answer is: %s\n\n", q.Options[q.CorrectOption])
		}

		input := q.PromptNext()

		// if switch is used requires labled break
		// testing if the request is to quit
		if input == "q" {
			break
		}

		// default case doesn't require any action
		// ommiting 'n'
		switch input {
		case "m":
			// mastering last node requires replacing
			// the quiz pointer
			q = q.Master()
		case "u":
			q = q.Mask()
			fmt.Printf("%T\n", q)
		}

		// mastering all quizes should be checked
		// before staged and masked is compared
		// if no new node available mastering a quiz
		// decreases Stat.Staged
		if len(q.Stat.Mastered) == q.Stat.Total {
			q.DisplayStat()
			fmt.Println("All quizes are mastered")
			break
		}

		// mask + master exceeds the quiz
		if q.Stat.Staged < q.Config.PerStage && q.AllMaskedOrMastered() {
			q.DisplayStat()
			fmt.Println("All quiezes are either mastered or masked")
			break
		}

		if q.ReachedMaskLimit() {
			q.DisplayStat()
			fmt.Println("Mask limit reached!")
			break
		}

		q = q.Advance()
	}
}
