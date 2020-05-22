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

	qn := qm.StageQuiz()

	var q helper.Quiz

mainLoop:
	for {

		clearScreen()
		qm.DisplayStat()
		fmt.Println()

		if qm.Stat.Masked == qm.StageCount {
			fmt.Println("All staged quizes masked or mastered!")
			break
		}

		q = qn.Cur.(helper.Quiz)

		q.Display()
		fmt.Println()
		q.PromptAns()

		if q.IsCorrect() {
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect! Correct answer is:", q.Options[q.CorrectOption])
		}

		switch q.PromptNext() {
		case "n":
			fmt.Println("n is pressed")
		case "m":
			// fmt.Println("word should be mastered")
			qm.Stat.Mastered++
		case "u":
			qm.Mask(&qn)
		case "v":
			fmt.Println("Correct answer should be displayed")
		case "q":
			break mainLoop
		default:
			fmt.Println("Invalid Option")
		}

		qn = *qn.Next
	}
}

func clearScreen() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
