package main

import (
	"os"
	csv "quiz-game/csv"
	q "quiz-game/questions"
)

func main() {

	userStart := q.StartQuiz()

	if userStart {
		csvFile := csv.SetCSV()
		records := csv.ReadCSV(csvFile)
		questionsAndAnswers := q.ParseQuestions(records)
		userAnswers := q.OutputQuestions(questionsAndAnswers)
		correctAnswers, incorrectAnswers := q.CheckAnswers(userAnswers)
		q.OutputResults(correctAnswers, incorrectAnswers)
	} else {
		os.Exit(3)
	}
}
