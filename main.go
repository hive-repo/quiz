package main

import "fmt"

// define quiz
// ask input
// compare result

type Option string

type Quiz struct {
	Question      string
	Options       []Option
	CorrectOption int
}

func main() {

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

	for _, q := range quizes {
		fmt.Println(q.Question)
		fmt.Println()

		// print options
		for k, v := range q.Options {
			fmt.Printf("%d. %s\n", k+1, v)
		}

		fmt.Println()
		fmt.Print("Answer: ")

		var ans int
		// ask ans
		fmt.Scan(&ans)

		// check ans

		if ans == q.CorrectOption {
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}
	}

}
