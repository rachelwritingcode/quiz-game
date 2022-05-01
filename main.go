package main

import (
	"fmt"
	csv "quiz-game/csv"
	q "quiz-game/questions"
)

func main() {

	csvFile := csv.SetCSV()
	records := csv.ReadCSV(csvFile)
	questionsAndAnswers := q.ParseQuestions(records)
	fmt.Println(questionsAndAnswers)
	userInput := q.UserContinue()
	fmt.Println(userInput)
}
