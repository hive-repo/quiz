package helper

import (
	"fmt"
	"os"
	"os/exec"
)

// Option for quiz
type Option string

// QuizStat records current statistics of Quiz
type QuizStat struct {
	Total       int
	Staged      int
	Mastered    int
	Masked      int
	StageCursor int
}

type QuizConfig struct {
	PerStage int
}

// Quiz struct
type Quiz struct {
	Next *Quiz
	Prev *Quiz

	ID            int
	Question      string
	Options       []Option
	CorrectOption int

	Stat *QuizStat

	all *[]Quiz

	Config *QuizConfig
}

// Display displays quiz
func (q *Quiz) Display() {

	fmt.Println(q.Question)
	fmt.Println()

	// print options
	for k, v := range q.Options {
		fmt.Printf("%d. %s\n", k+1, v)
	}

	fmt.Println()
}

// PromptAns asks ans
func (q Quiz) PromptAns() int {
	var ans int
	fmt.Print("Answer: ")
	fmt.Scan(&ans)

	return ans
}

// IsCorrect verifies given answer with correct option
func (q Quiz) IsCorrect(ans int) bool {
	return q.CorrectOption == ans-1
}

// PromptNext prompts next action
func (q *Quiz) PromptNext() string {
	var input string
	fmt.Printf("NEXT(n)\t\tMASTER(m)\tMASK(u)\t\tVIEW(v)\t\tQUIT(q): ")
	fmt.Scan(&input)

	return input
}

// Master masters the Quiz
func (q *Quiz) Master() *Quiz {
	q.Stat.Mastered++

	// no new quiz
	// remove current node from the chain
	if q.Stat.Total == q.Stat.StageCursor {
		q.Prev.Next = q.Next
		q.Next.Prev = q.Prev

		// node removed from staged chain
		// but no new node added
		q.Stat.Staged--

		return q
	}

	quiz := (*q.all)[q.Stat.StageCursor]

	nq := Quiz{
		Next:          q.Next,
		Prev:          q.Prev,
		ID:            q.Stat.StageCursor,
		Question:      quiz.Question,
		Options:       quiz.Options,
		CorrectOption: quiz.CorrectOption,
		Stat:          q.Stat,
		all:           q.all,
	}

	// last node
	if q.Next.ID == q.ID {
		q = &nq
		q.Next = &nq
		q.Prev = &nq

	} else {
		q.Prev.Next = &nq
		q.Next.Prev = &nq
	}

	q.Stat.StageCursor++

	return q
}

// Mask masks the Quiz
func (q *Quiz) Mask() {
	q.Stat.Masked++
	q.Prev.Next = q.Next
	q.Next.Prev = q.Prev
}

// DisplayStat displays stats
func (q *Quiz) DisplayStat() {
	clear()

	fmt.Printf("Total: %d\tStaged: %d\tMastered: %d [%.2f%s]\tMasked: %d\n",
		q.Stat.Total,
		q.Stat.Staged,
		q.Stat.Mastered,
		float64(q.Stat.Mastered)/float64(q.Stat.Total)*100,
		"%",
		q.Stat.Masked)

	fmt.Println()
}

// Build builds the quiz
func (q *Quiz) Build() Quiz {

	config := QuizConfig{
		PerStage: 3,
	}

	stat := QuizStat{}

	quizes := []Quiz{
		Quiz{
			Question: "1. What is the Capital city of Nepal?",
			Options: []Option{
				"Delhi",
				"Dhaka",
				"Kathmandu",
				"Nepalgunj",
			},
			CorrectOption: 2,
		},
		Quiz{
			Question: "2. Which is the biggest lake of Nepal?",
			Options: []Option{
				"Rara",
				"Foksundo",
				"Fewatal",
				"Se Foksundo",
			},
			CorrectOption: 0,
		},
		Quiz{
			Question: "3. National bird of Nepal",
			Options: []Option{
				"Maina",
				"Danfe",
				"Suga",
				"Gothri",
			},
			CorrectOption: 1,
		},
		Quiz{
			Question: "4. 2 x 2",
			Options: []Option{
				"6",
				"4",
				"3",
				"8",
			},
			CorrectOption: 1,
		},
		Quiz{
			Question: "5. 2-2",
			Options: []Option{
				"0",
				"1",
				"-1",
				"8",
			},
			CorrectOption: 0,
		},
		Quiz{
			Question: "6. 2+2",
			Options: []Option{
				"0",
				"1",
				"-1",
				"4",
			},
			CorrectOption: 3,
		},
	}

	for i := 0; i < config.PerStage; i++ {
		var prev, next int

		prev = i - 1
		// prev first first node will be last node
		if i == 0 {
			prev = config.PerStage - 1
		}

		next = i + 1
		// next for last node will be first node
		if i == config.PerStage-1 {
			next = 0
		}

		quizes[i].ID = i
		quizes[i].Prev = &quizes[prev]
		quizes[i].Next = &quizes[next]
		quizes[i].Stat = &stat
		quizes[i].all = &quizes
		quizes[i].Config = &config
	}

	stat.Total = len(quizes)
	stat.Staged = config.PerStage
	stat.StageCursor = stat.Staged

	return quizes[0]
}

// Advance advances the quiz to next
func (q *Quiz) Advance() *Quiz {
	return q.Next
}

func clear() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
