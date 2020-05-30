package main

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/hive-repo/quiz/helper"
	"gopkg.in/yaml.v2"
)

func main() {

	config := helper.QuizConfig{}
	user, _ := user.Current()

	configFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/config.yaml")

	yaml.Unmarshal([]byte(configFile), &config)

	quizes := []helper.Quiz{}

	quizFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/quizes.yaml")

	yaml.Unmarshal([]byte(quizFile), &quizes)

	stat := helper.QuizStat{}
	statFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/stat.yaml")

	yaml.Unmarshal([]byte(statFile), &stat)

	q := (helper.Quiz{}).Build(config, stat, quizes)

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
			fmt.Println()
			break
		}

		// default case doesn't require any action
		// ommiting 'n'
		switch input {
		case "m":
			// mastering last node requires replacing
			// the quiz pointer
			q.Master()
		case "u":
			q.Mask()
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
		if q.AllMaskedOrMastered() {
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
