package helper

import "fmt"

// Option for quiz
type Option string

// Quiz struct
type Quiz struct {
	Question      string
	Options       []Option
	CorrectOption int
	ans           int
}

// Display displays quiz
func (q Quiz) Display() {
	fmt.Println(q.Question)
	fmt.Println()

	// print options
	for k, v := range q.Options {
		fmt.Printf("%d. %s\n", k+1, v)
	}
}

// PromptAns asks ans
func (q *Quiz) PromptAns() {
	fmt.Print("Answer: ")
	fmt.Scan(&q.ans)
}

// IsCorrect verifies given answer with correct option
func (q Quiz) IsCorrect() bool {
	return q.ans == q.CorrectOption+1
}

// PromptNext prompts next action
func (q Quiz) PromptNext() string {
	var next string
	fmt.Printf("NEXT(n)\t\tMASTER(m)\tMASK(u)\t\tVIEW(v)\t\tQUIT(q): ")
	fmt.Scan(&next)

	return next
}
