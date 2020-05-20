package helper

import "fmt"

// Manager is the Quiz Manager
type Manager struct {
	Stat struct {
		Total    int
		Staged   int
		Mastered int
		Masked   int
	}
	stageCount int
	staged     []Quiz
	mastered   []int
	masked     []int
	curIndex   int
	quizes     []Quiz
}

// BuildQuiz builds quiz
func (m *Manager) BuildQuiz() {
	m.stageCount = 3

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

	// Total
	m.quizes = quizes
	m.Stat.Total = len(quizes)

	// Staged
	m.staged = quizes[0:m.stageCount]

	m.Stat.Staged = len(m.staged)
}

// DisplayStat displays stats
func (m Manager) DisplayStat() {
	fmt.Printf("Total Words: %d\tStaged: %d\tMastered: %d[%.2f]\tMasked: %d\n", m.Stat.Total, m.Stat.Staged, m.Stat.Mastered, float64(m.Stat.Mastered)/float64(m.Stat.Total)*100, m.Stat.Masked)
}

// GetQuiz returns a quiz
func (m *Manager) GetQuiz() Quiz {
	q := m.staged[m.curIndex%m.Stat.Staged]

	m.curIndex++

	return q
}
