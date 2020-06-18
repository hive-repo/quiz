package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"time"

	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"
)

// Option for quiz
type Option string

// QuizStat records current statistics of Quiz
type QuizStat struct {
	Total    int   `yaml:"total"`
	Mastered []int `yaml:"mastered"`
	Cursor   int   `yaml:"cursor"`
	Masked   int
	Staged   int
}

// QuizConfig holds configs
type QuizConfig struct {
	PerStage int `yaml:"perStage"`
}

// Quiz struct
type Quiz struct {
	Prev *Quiz
	Next *Quiz

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

	fmt.Printf("Question: %s\n", chalk.Bold.TextStyle(q.Question))
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
	ans, _, _ = getChar()
	ans, _ = strconv.Atoi(string(ans))

	return ans
}

// IsCorrect verifies given answer with correct option
func (q Quiz) IsCorrect(ans int) bool {
	return q.CorrectOption == ans-1
}

// PromptNext prompts next action
func (q *Quiz) PromptNext() string {
	var input string
	fmt.Printf("MASTER[m]\tMASK[u]\t\tQUIT[q]\t\tNEXT[n]: ")
	ans, _, _ := getChar()
	input = string(ans)
	fmt.Println(input)

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
		q.Stat.Staged--
	} else {
		n := &(*q.all)[q.Stat.Cursor]

		n.Next = q.Next
		n.Prev = q.Prev
		n.Stat = q.Stat
		n.Config = q.Config
		n.all = q.all

		q.Prev.Next = n
		q.Next.Prev = n

		q.Stat.Cursor++
	}

	q.saveStat()
}

// SaveStat saves current stats
func (q *Quiz) saveStat() {
	// var data []byte
	data, _ := yaml.Marshal(q.Stat)

	user, _ := user.Current()
	ioutil.WriteFile(user.HomeDir+"/.quiz/stat.yaml", data, 0640)

	if _, err := os.Stat(user.HomeDir + "/.quiz/mastered-stats"); os.IsNotExist(err) {
		os.Mkdir(user.HomeDir+"/.quiz/mastered-stats", 755)
	}

	ioutil.WriteFile(user.HomeDir+"/.quiz/mastered-stats/"+time.Now().Format("2006-01-02"), []byte(strconv.Itoa(len(q.Stat.Mastered))+"\n"), 0640)
}

// Mask masks the Quiz
func (q *Quiz) Mask() {
	q.Stat.Masked++
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
	for _, i := range stat.Mastered {
		quizes[i].isMastered = true
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

		q.Prev = &staged[prev]
		q.Next = &staged[next]
		q.Stat = &stat
		q.all = &quizes
		q.Config = &config

		staged[i] = q

		i++
		stat.Staged++
	}

	stat.Total = len(quizes)
	stat.Cursor = len(stat.Mastered) + stat.Staged

	return &staged[0]
}

// Advance advances the quiz to next
func (q *Quiz) Advance() *Quiz {
	// Next pointer needs to be returned
	// Receiver recives the Reference Pointer in it's own Pointer
	return q.Next
}

// AllMaskedOrMastered checks if all ethier masked or mastered
func (q *Quiz) AllMaskedOrMastered() bool {
	return q.Stat.Staged < q.Config.PerStage &&
		q.Stat.Masked+len(q.Stat.Mastered) == q.Stat.Total
}

// ReachedMaskLimit checks if mask limit is reached
func (q *Quiz) ReachedMaskLimit() bool {
	return q.Stat.Masked == q.Stat.Staged
}

func clear() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
