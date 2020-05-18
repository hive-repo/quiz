package helper

type Manager struct {
	Total    int
	Staged   int
	Mastered int
	Masked   int
}

func (m Manager) BuildQuiz() []Quiz {
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

	return quizes
}
