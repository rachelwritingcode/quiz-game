package main

import (
	"os"
	"quiz-game/csv"
	q "quiz-game/questions"
)

func main() {

	timeout := q.SetTimeout()
	userStart := q.StartQuiz(os.Stdin)

	if userStart {
		csvFile := csv.SetCSV()
		records := csv.ReadCSV(csvFile)
		questionsAndAnswers := q.ParseQuestions(records)
		q.OutputQuestions(questionsAndAnswers, timeout, os.Stdin)
	} else {
		os.Exit(3)
	}
}
