package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"

	"gopkg.in/yaml.v2"
)

// Option for quiz
type Option string

// QuizStat records current statistics of Quiz
type QuizStat struct {
	Total    int   `yaml:"total"`
	Mastered []int `yaml:"mastered"`
	Cursor   int   `yaml:"cursor"`
	masked   int
	staged   int
}

// QuizConfig holds configs
type QuizConfig struct {
	PerStage int `yaml:"perStage"`
}

// Quiz struct
type Quiz struct {
	Next *Quiz
	Prev *Quiz

	ID            int      `yaml:"id"`
	Question      string   `yaml:"question"`
	Options       []Option `yaml:"options"`
	CorrectOption int      `yaml:"correctOption"`

	Stat       *QuizStat
	isMastered bool

	all *[]Quiz

	Config *QuizConfig
}

// DisplayStat displays stats
func (q *Quiz) DisplayStat() {
	clear()

	fmt.Printf("TOTAL: %d\tSTAGED: %d/%d\t\tMASKED: %d/%d\tMASTERED: %d/%d, %.2f%s\n",
		q.Stat.Total,
		q.Stat.staged,
		q.Stat.Total,
		q.Stat.masked,
		q.Stat.staged,
		len(q.Stat.Mastered),
		q.Stat.Total,
		float64(len(q.Stat.Mastered))/float64(q.Stat.Total)*100,
		"%")

	fmt.Println()
}

// Display displays quiz
func (q *Quiz) Display() {

	fmt.Printf("Question: %s\n", q.Question)
	fmt.Println()

	// print options
	fmt.Printf("OPTIONS:\n\n")
	for k, v := range q.Options {
		fmt.Printf(" [%d] %s\n", k+1, v)
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
	fmt.Printf("NEXT[n]\t\tMASTER[m]\tMASK[u]\t\tVIEW[v]\t\tQUIT[q]: ")
	fmt.Scan(&input)

	return input
}

// Master masters the Quiz
func (q *Quiz) Master() {

	q.Stat.Mastered = append(q.Stat.Mastered, q.ID)

	// no new quiz
	// remove current node from the chain
	if q.Stat.Total == q.Stat.Cursor {
		q.Prev.Next = q.Next
		q.Next.Prev = q.Prev

		// node removed from staged chain
		// but no new node added
		q.Stat.staged--
	}

	nq := (*q.all)[q.Stat.Cursor]

	nq.Next = q.Next
	nq.Prev = q.Prev
	nq.Stat = q.Stat
	nq.Config = q.Config
	nq.all = q.all

	// last node
	if q.Next.ID == q.ID {
		q = &nq
		q.Next = &nq
		q.Prev = &nq

	} else {
		q.Prev.Next = &nq
		q.Next.Prev = &nq
	}

	q.Stat.Cursor++

	q.saveStat()
}

// SaveStat saves current stats
func (q *Quiz) saveStat() {
	// var data []byte
	data, _ := yaml.Marshal(q.Stat)

	user, _ := user.Current()
	ioutil.WriteFile(user.HomeDir+"/.quiz/stat.yaml", data, 0640)
}

// Mask masks the Quiz
func (q *Quiz) Mask() {
	q.Stat.masked++
	q.Prev.Next = q.Next
	q.Next.Prev = q.Prev
}

// Build builds the quiz
func (q Quiz) Build(config QuizConfig, stat QuizStat, quizes []Quiz) *Quiz {

	staged := make([]Quiz, 0, 50)

	// fill out staged
	for i := 0; i < config.PerStage; i++ {
		staged = append(staged, Quiz{})
	}
	// mark mastered
	for _, qi := range stat.Mastered {
		quizes[qi].isMastered = true
	}

	var prev, next int

	i := 0
	for _, q := range quizes {
		if i == config.PerStage {
			break
		}

		if q.isMastered {
			continue
		}

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

		staged[i] = q
		staged[i].ID = q.ID
		staged[i].Prev = &staged[prev]
		staged[i].Next = &staged[next]
		staged[i].Stat = &stat
		staged[i].all = &quizes
		staged[i].Config = &config

		i++
		stat.staged++
	}

	stat.Total = len(quizes)
	stat.Cursor += stat.staged
	fmt.Println(stat.Cursor)

	return &staged[0]
}

// Advance advances the quiz to next
func (q *Quiz) Advance() *Quiz {
	return q.Next
}

// AllMaskedOrMastered checks if all ethier masked or mastered
func (q *Quiz) AllMaskedOrMastered() bool {
	return q.Stat.staged < q.Config.PerStage &&
		q.Stat.masked+len(q.Stat.Mastered) == q.Stat.Total
}

// ReachedMaskLimit checks if mask limit is reached
func (q *Quiz) ReachedMaskLimit() bool {
	return q.Stat.masked == q.Stat.staged
}

func clear() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
