package helper

import "fmt"

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

// Quiz struct
type Quiz struct {
	Next *Quiz
	Prev *Quiz

	Question      string
	Options       []Option
	CorrectOption int

	Stat *QuizStat

	all *[]Quiz
}

// Display displays quiz
func (q *Quiz) Display() {

	fmt.Println(q.Question)
	fmt.Println()

	// print options
	for k, v := range q.Options {
		fmt.Printf("%d. %s\n", k+1, v)
	}
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
func (q *Quiz) Master() {
	q.Stat.Mastered++

	quiz := (*q.all)[q.Stat.StageCursor]

	nq := Quiz{
		Next:          q.Next,
		Prev:          q.Prev,
		Question:      quiz.Question,
		Options:       quiz.Options,
		CorrectOption: quiz.CorrectOption,
		Stat:          q.Stat,
		all:           q.all,
	}

	q.Prev.Next = &nq
	q.Next.Prev = &nq

	q.Stat.StageCursor++
}

// Mask masks the Quiz
func (q *Quiz) Mask() {
	q.Stat.Masked++
	q.Prev.Next = q.Next
	q.Next.Prev = q.Prev
}

// DisplayStat displays stats
func (q *Quiz) DisplayStat() {
	fmt.Printf("Total Words: %d\tStaged: %d\tMastered: %d[%.2f]\tMasked: %d\n",
		q.Stat.Total,
		q.Stat.Staged,
		q.Stat.Mastered,
		float64(q.Stat.Mastered)/float64(q.Stat.Total)*100,
		q.Stat.Masked)
}

// Build builds the quiz
func (q *Quiz) Build() Quiz {

	perStage := 3

	qs := QuizStat{}

	quizes := []Quiz{
		Quiz{
			Question: "What is the Capital city of Nepal?",
			Options: []Option{
				"Delhi",
				"Dhaka",
				"Kathmandu",
				"Nepalgunj",
			},
			CorrectOption: 2,
		},
		Quiz{
			Question: "Which is the biggest lake of Nepal?",
			Options: []Option{
				"Rara",
				"Foksundo",
				"Fewatal",
				"Se Foksundo",
			},
			CorrectOption: 0,
		},
		Quiz{
			Question: "National bird of Nepal",
			Options: []Option{
				"Maina",
				"Danfe",
				"Suga",
				"Gothri",
			},
			CorrectOption: 1,
		},
		Quiz{
			Question: "2 x 2",
			Options: []Option{
				"6",
				"4",
				"3",
				"8",
			},
			CorrectOption: 1,
		},
		Quiz{
			Question: "2-2",
			Options: []Option{
				"0",
				"1",
				"-1",
				"8",
			},
			CorrectOption: 0,
		},
	}

	for i := 0; i < perStage; i++ {
		var prev, next int

		prev = i - 1
		// prev first first node will be last node
		if i == 0 {
			prev = perStage - 1
		}

		next = i + 1
		// next for last node will be first node
		if i == perStage-1 {
			next = 0
		}

		quizes[i].Prev = &quizes[prev]
		quizes[i].Next = &quizes[next]
		quizes[i].Stat = &qs
		quizes[i].all = &quizes
	}

	qs.Total = len(quizes)
	qs.Staged = perStage
	qs.StageCursor = qs.Staged

	return quizes[0]
}
