package main

import (
	"fmt"

	"github.com/hive-repo/quiz/helper"
)

func main() {

	qm := helper.Manager{}

	quizes := qm.BuildQuiz()

	for _, q := range quizes {

		q.Display()
		q.PromptAns()

		if q.IsCorrect() {
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}
	}
}
