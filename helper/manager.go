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
	StageCount int
	staged     []Quiz
	mastered   []int
	masked     []int
	curIndex   int
	quizes     []Quiz
}

// StageQuiz stages quiz
func (m *Manager) StageQuiz() Node {
	m.StageCount = 3

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

	nodes := make([]Node, m.StageCount)

	for i := 0; i < m.StageCount; i++ {
		var prev, next int

		prev = i - 1
		// prev first first node will be last node
		if i == 0 {
			prev = m.StageCount - 1
		}

		next = i + 1
		// next for last node will be first node
		if i == m.StageCount-1 {
			next = 0
		}

		nodes[i] = Node{
			Prev: &nodes[prev],
			Next: &nodes[next],
			Cur:  quizes[i],
		}
	}

	// Total
	m.quizes = quizes
	m.Stat.Total = len(quizes)

	// Staged
	m.staged = quizes[0:m.StageCount]

	m.Stat.Staged = len(m.staged)

	return nodes[0]
}

// BuildQuiz builds quiz
func (m *Manager) BuildQuiz() {
	m.StageCount = 3

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
	m.staged = quizes[0:m.StageCount]

	m.Stat.Staged = len(m.staged)
}

// GetQuiz returns a quiz
func (m *Manager) GetQuiz() Quiz {
	//fmt.Println("staged", m.staged)
	fmt.Println("masked", m.masked)
	fmt.Println("cur index:", m.curIndex)
	// staged - masked

	i := m.curIndex % m.StageCount
	fmt.Println("eci:", i)

	for _, v := range m.masked {
		fmt.Println(i, v)
		if v == i {
			m.curIndex++
		}
	}

	fmt.Println(m.StageCount)
	fmt.Println(len(m.staged))
	if m.StageCount == len(m.staged) {
		fmt.Println("All masked")
	}

	fmt.Println("after mask check:", m.curIndex)

	q := m.staged[m.curIndex%m.Stat.Staged]

	m.curIndex++

	return q
}

// Mask maskes the quiz
func (m *Manager) Mask(qn *Node) {

	qn.Prev.Next = qn.Next
	qn.Next.Prev = qn.Prev

	m.Stat.Masked++
}
