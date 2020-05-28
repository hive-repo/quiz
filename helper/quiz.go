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
	Total       int   `yaml:"total"`
	Staged      int   `yaml:"staged"`
	Mastered    []int `yaml:"mastered"`
	Masked      int   `yaml:"masked"`
	StageCursor int   `yaml:"stageCursor"`
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
		q.Stat.Staged,
		q.Stat.Total,
		q.Stat.Masked,
		q.Stat.Staged,
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
func (q *Quiz) Master() *Quiz {

	q.Stat.Mastered = append(q.Stat.Mastered, q.ID)

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
		Config:        q.Config,
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

	q.saveStat()
	return q
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
	q.Stat.Masked++
	q.Prev.Next = q.Next
	q.Next.Prev = q.Prev
}

// Build builds the quiz
func (q *Quiz) Build() Quiz {

	config := QuizConfig{}
	user, _ := user.Current()

	configFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/config.yaml")

	yaml.Unmarshal([]byte(configFile), &config)

	stat := QuizStat{}

	quizes := []Quiz{}
	staged := make([]Quiz, 0, 50)

	// fill out staged
	for i := 0; i < config.PerStage; i++ {
		staged = append(staged, Quiz{})
	}

	quizFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/quizes.yaml")

	yaml.Unmarshal([]byte(quizFile), &quizes)

	statFile, _ := ioutil.ReadFile(user.HomeDir + "/.quiz/stat.yaml")

	yaml.Unmarshal([]byte(statFile), &stat)

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
	}

	stat.Total = len(quizes)
	stat.Staged = config.PerStage
	stat.StageCursor = stat.Staged

	return staged[0]
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
