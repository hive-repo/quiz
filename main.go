package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/ghodss/yaml"
	"github.com/hive-repo/quiz/helper"
	"github.com/ttacon/chalk"
)

func main() {

	app := app.New()

	w := app.NewWindow("Quiz")
	w.Resize(fyne.Size{
		Width:  800,
		Height: 300,
	})

	config := helper.QuizConfig{}
	// config.PerStage = 3
	user, _ := user.Current()

	configFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/config.yaml")

	yaml.Unmarshal([]byte(configFile), &config)

	quizes := []helper.Quiz{}
	// quizes := []helper.Quiz{
	// 	{
	// 		Question: "2x2",
	// 		Options: []helper.Option{
	// 			"1",
	// 			"2",
	// 			"3",
	// 			"4",
	// 		},
	// 		CorrectOption: 3,
	// 	},
	// 	{
	// 		Question: "2-2",
	// 		Options: []helper.Option{
	// 			"0",
	// 			"1",
	// 			"2",
	// 			"3",
	// 		},
	// 		CorrectOption: 0,
	// 	},
	// 	{
	// 		Question: "2+2",
	// 		Options: []helper.Option{
	// 			"4",
	// 			"1",
	// 			"2",
	// 			"3",
	// 		},
	// 		CorrectOption: 0,
	// 	},
	// 	{
	// 		Question: "2/2",
	// 		Options: []helper.Option{
	// 			"4",
	// 			"1",
	// 			"2",
	// 			"3",
	// 		},
	// 		CorrectOption: 1,
	// 	},
	// }

	quizFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/quizes.yaml")

	yaml.Unmarshal([]byte(quizFile), &quizes)

	stat := helper.QuizStat{}
	statFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/stat.yaml")

	yaml.Unmarshal([]byte(statFile), &stat)

	q := (helper.Quiz{}).Build(config, stat, quizes)

	l := struct {
		question *widget.Label
		options  *widget.Radio
		result   *widget.Label
		correct  *widget.Label
		Stat     struct {
			Total    *widget.Label
			Staged   *widget.Label
			Masked   *widget.Label
			Mastered *widget.Label
		}
	}{}

	l.question = widget.NewLabel("")
	l.options = widget.NewRadio([]string{
		"1. " + string(q.Options[0]),
		"2. " + string(q.Options[1]),
		"3. " + string(q.Options[2]),
		"4. " + string(q.Options[3])},
		func(o string) {
			fmt.Println(o)
			if o == strconv.Itoa(q.CorrectOption+1)+". "+string(q.Options[q.CorrectOption]) {
				l.result.Text = "Correct!"
				l.result.TextStyle.Bold = true

			} else {
				l.result.Text = "Incorrect!"
				l.result.TextStyle.Bold = true
				l.correct.Text = "Correct Ans: " + string(q.Options[q.CorrectOption])
			}

			l.correct.Refresh()
			l.result.Refresh()

		})
	l.result = widget.NewLabel("")
	l.correct = widget.NewLabel("")

	l.question.Text = q.Question
	l.question.TextStyle.Bold = true
	l.Stat.Total = widget.NewLabel(strconv.Itoa(q.Stat.Total))
	l.Stat.Staged = widget.NewLabel(strconv.Itoa(q.Stat.Staged))
	l.Stat.Masked = widget.NewLabel(strconv.Itoa(q.Stat.Masked))
	l.Stat.Mastered = widget.NewLabel(strconv.Itoa(len(q.Stat.Mastered)))

	quizStat := widget.NewHBox(
		widget.NewLabel("Total: "), l.Stat.Total,
		widget.NewLabel("Staged: "), l.Stat.Staged,
		widget.NewLabel("Masked: "), l.Stat.Masked,
		widget.NewLabel("Mastered: "), l.Stat.Mastered,
	)
	quizBox := widget.NewVBox(
		quizStat,
		widget.NewHBox(
			widget.NewLabel("Question: "),
			l.question,
		),
		widget.NewLabel("Options:   "),
		l.options,
		widget.NewHBox(
			l.result,
			l.correct,
		),
	)

	w.SetContent(widget.NewVBox(
		quizBox,

		widget.NewHBox(
			widget.NewButton("Master", func() {
				q.Master()
				q = q.Advance()
				l.question.Text = q.Question
				l.question.Refresh()
				l.result.Text = ""
				l.result.Refresh()
				l.Stat.Mastered.Text = strconv.Itoa(len(q.Stat.Mastered))
				l.Stat.Mastered.Refresh()

				l.options.Options[0] = "1. " + string(q.Options[0])
				l.options.Options[1] = "2. " + string(q.Options[1])
				l.options.Options[2] = "3. " + string(q.Options[2])
				l.options.Options[3] = "4. " + string(q.Options[3])

				l.options.Refresh()
			}),
			widget.NewButton("Mask", func() {
				q.Mask()
				q = q.Advance()
				l.question.Text = q.Question
				l.question.Refresh()
				l.correct.Text = ""
				l.correct.Refresh()
				l.result.Text = ""
				l.result.Refresh()
				l.Stat.Masked.Text = strconv.Itoa(q.Stat.Masked)
				l.Stat.Masked.Refresh()

				l.options.Options[0] = "1. " + string(q.Options[0])
				l.options.Options[1] = "2. " + string(q.Options[1])
				l.options.Options[2] = "3. " + string(q.Options[2])
				l.options.Options[3] = "4. " + string(q.Options[3])

				l.options.Refresh()

			}),
			widget.NewButton("Next", func() {
				q = q.Advance()
				l.question.Text = q.Question
				l.question.Refresh()
				l.result.Text = ""
				l.result.Refresh()
				l.correct.Text = ""
				l.correct.Refresh()

				l.options.Options[0] = "1. " + string(q.Options[0])
				l.options.Options[1] = "2. " + string(q.Options[1])
				l.options.Options[2] = "3. " + string(q.Options[2])
				l.options.Options[3] = "4. " + string(q.Options[3])

				l.options.Refresh()
			}),

			widget.NewButton("Quit", func() {
				app.Quit()
			}),
		),
	))

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		fmt.Println("KeyPress:", k)
		switch k.Name {
		case "Q":
			app.Quit()
		case "1":
			o, _ := strconv.Atoi(string(k.Name))
			l.options.SetSelected("1. " + string(q.Options[o-1]))
			if o-1 == q.CorrectOption {
				l.result.Text = "Correct!"
				l.result.TextStyle.Bold = true

			} else {
				l.result.Text = "Incorrect!"
				l.result.TextStyle.Bold = true
				l.correct.Text = string(q.Options[q.CorrectOption])
			}
			l.correct.Refresh()
			l.result.Refresh()

		case "2":
			o, _ := strconv.Atoi(string(k.Name))
			l.options.SetSelected("2. " + string(q.Options[o-1]))
			if o-1 == q.CorrectOption {
				l.result.Text = "Correct!"
				l.result.TextStyle.Bold = true

			} else {
				l.result.Text = "Incorrect!"
				l.result.TextStyle.Bold = true
				l.correct.Text = string(q.Options[q.CorrectOption])
			}
			l.correct.Refresh()
			l.result.Refresh()
		case "3":
			o, _ := strconv.Atoi(string(k.Name))
			l.options.SetSelected("3. " + string(q.Options[o-1]))
			if o-1 == q.CorrectOption {
				l.result.Text = "Correct!"
				l.result.TextStyle.Bold = true

			} else {
				l.result.Text = "Incorrect!"
				l.result.TextStyle.Bold = true
				l.correct.Text = string(q.Options[q.CorrectOption])
			}
			l.correct.Refresh()
			l.result.Refresh()
		case "4":
			o, _ := strconv.Atoi(string(k.Name))
			l.options.SetSelected("4. " + string(q.Options[o-1]))
			if o-1 == q.CorrectOption {
				l.result.Text = "Correct!"
				l.result.TextStyle.Bold = true

			} else {
				l.result.Text = "Incorrect!"
				l.result.TextStyle.Bold = true
				l.correct.Text = string(q.Options[q.CorrectOption])
			}
			l.correct.Refresh()
			l.result.Refresh()
		case "9":
			q.Mask()
			q = q.Advance()
			l.question.Text = q.Question
			l.question.Refresh()
			l.result.Text = ""
			l.result.Refresh()
			l.Stat.Masked.Text = strconv.Itoa(q.Stat.Masked)
			l.Stat.Masked.Refresh()

			l.options.Options[0] = "1. " + string(q.Options[0])
			l.options.Options[1] = "2. " + string(q.Options[1])
			l.options.Options[2] = "3. " + string(q.Options[2])
			l.options.Options[3] = "4. " + string(q.Options[3])

			l.options.Refresh()
		case "0":
			q.Master()
			q = q.Advance()
			l.question.Text = q.Question
			l.question.Refresh()
			l.result.Text = ""
			l.result.Refresh()
			l.Stat.Mastered.Text = strconv.Itoa(len(q.Stat.Mastered))
			l.Stat.Mastered.Refresh()

			l.options.Options[0] = "1. " + string(q.Options[0])
			l.options.Options[1] = "2. " + string(q.Options[1])
			l.options.Options[2] = "3. " + string(q.Options[2])
			l.options.Options[3] = "4. " + string(q.Options[3])

			l.options.Refresh()
		default:
			q = q.Advance()
			l.question.Text = q.Question
			l.question.Refresh()
			l.result.Text = ""
			l.result.Refresh()

			l.options.Options[0] = "1. " + string(q.Options[0])
			l.options.Options[1] = "2. " + string(q.Options[1])
			l.options.Options[2] = "3. " + string(q.Options[2])
			l.options.Options[3] = "4. " + string(q.Options[3])

			l.options.Refresh()
		}

	})

	w.ShowAndRun()
	return

	for {

		q.DisplayStat()

		q.Display()

		ans := q.PromptAns()

		if q.IsCorrect(ans) {
			fmt.Printf(" [%d] %s\n\n", ans, chalk.Green.Color("Correct!"))
		} else {
			fmt.Printf(" [%d] %s\nCorrect: [%d] %s\n\n", ans,
				chalk.Red.Color("Incorrect!"),
				q.CorrectOption+1,
				q.Options[q.CorrectOption])
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
