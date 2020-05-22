package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hive-repo/quiz/helper"
)

func main() {

	qm := helper.Manager{}

	qm.BuildQuiz()

	q := helper.Quiz{}

	q = q.Build()

	for {

		clear()
		q.DisplayStat()
		fmt.Println()

		if q.Stat.Masked == q.Stat.Staged {
			fmt.Println("All staged quizes masked or mastered!")
			break
		}

		q.Display()
		fmt.Println()

		ans := q.PromptAns()

		if q.IsCorrect(ans) {
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect! Correct answer is:", q.Options[q.CorrectOption])
		}

		input := q.PromptNext()

		// switch is used requires labled break
		// testing if the request is to quit
		if input == "q" {
			break
		}

		// default case doesn't require any action
		// ommiting 'n'
		switch input {
		case "m":
			q.Master()
		case "u":
			q.Mask()
		}

		q = *q.Next
	}
}

func clear() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
