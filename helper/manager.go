package helper

import "fmt"

// Manager is the Quiz Manager
type Manager struct {
	Total    int
	Staged   int
	Mastered int
	Masked   int
}

// BuildQuiz builds quiz
func (m *Manager) BuildQuiz() []Quiz {
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
	}

	m.Total = len(quizes)

	return quizes
}

// DisplayStat displays stats
func (m Manager) DisplayStat() {
	fmt.Printf("Total Words: %d\tStaged: %d\tMastered: %d[%.2f]\tMasked: %d\n", m.Total, m.Staged, m.Mastered, float64(m.Mastered)/float64(m.Total)*100, m.Masked)
}
