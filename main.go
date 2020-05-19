package main

import (
	"fmt"

	"github.com/hive-repo/quiz/helper"
)

func main() {

	qm := helper.Manager{}

	quizes := qm.BuildQuiz()

	for _, q := range quizes {

		qm.DisplayStat()
		fmt.Println()
		q.Display()
		fmt.Println()
		q.PromptAns()

		switch q.PromptNext() {
		case "n":
			fmt.Println("n is pressed")
		case "m":
			fmt.Println("word should be mastered")
			qm.Mastered++
		case "u":
			fmt.Println("word should be masked")
		case "v":
			fmt.Println("Correct answer should be displayed")
		case "q":
			fmt.Println("Program should exit")
		default:
			fmt.Println("Invalid Option")
		}

		if q.IsCorrect() {
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}
	}
}
