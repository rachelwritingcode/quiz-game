package main

import (
	"fmt"
	"os"
	"quiz-game/csv"
	q "quiz-game/questions"
	"strconv"
)

func main() {

	userStart := q.StartQuiz(os.Stdin)
	if userStart {
		csvFile := csv.SetCSV()
		records := csv.ReadCSV(csvFile)
		questionsAndAnswers := q.ParseQuestions(records)
		userAnswers := q.OutputQuestions(questionsAndAnswers)
		answers := q.CheckAnswers(userAnswers)

		fmt.Println("Total Questions: " + strconv.FormatInt(answers[0]+answers[1], 10))
		fmt.Println("Total Correct Answers: " + strconv.FormatInt(answers[0], 10))
		fmt.Println("Total Incorrect Answers: " + strconv.FormatInt(answers[1], 10))

	} else {
		os.Exit(3)
	}
}
