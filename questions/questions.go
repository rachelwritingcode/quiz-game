package questions

import (
	"fmt"
	"strconv"
	"strings"
)

// Retrieve the questions and answers and store in a map
func ParseQuestions(records []string) map[string]string {

	var question string
	var answer string
	questionsAndAnswers := make(map[string]string)
	// First index
	questionsAndAnswers[records[0]] = records[1]

	for i := 2; i < len(records); i++ {
		if i%2 == 0 {
			question = records[i]
			answer = records[i+1]
			questionsAndAnswers[question] = answer
		}
	}
	return questionsAndAnswers
}

// Check if the user wants to start the quiz
func StartQuiz() bool {

	var user string
	fmt.Println("Do you want to start the quiz? Type Y or N.")
	fmt.Scanln(&user)

	if strings.Contains(strings.ToUpper(user), "Y") {
		return true
	} else {
		return false
	}
}

// Output the questions for the user to read and collect user input
func OutputQuestions(questionsAndAnswers map[string]string) map[string]string {

	var userResponse string
	userAnswers := make(map[string]string)

	for question, answer := range questionsAndAnswers {
		fmt.Println("What is the answer to this question?")
		fmt.Println(question)
		fmt.Scanln(&userResponse)
		fmt.Println("Your Answer: " + userResponse + "\n")

		if strings.Contains(userResponse, answer) {
			userAnswers[question] = "correct"
		} else {
			userAnswers[question] = "wrong"
		}
	}
	return userAnswers
}

// Check the user answers
func CheckAnswers(userAnswers map[string]string) []int64 {

	var answers []int64
	var correctAnswers int64 = 0
	var incorrectAnswers int64 = 0
	for _, answer := range userAnswers {
		if strings.Contains(answer, "correct") {
			correctAnswers += 1
		} else {
			incorrectAnswers += 1
		}
	}
	answers = append(answers, correctAnswers, incorrectAnswers)
	return answers
}

// Output the amount of correct and incorrect answers
func OutputResults(answers []int64) {

	totalQuestions := answers[0] + answers[1]
	correctAnswers := answers[0]
	incorrectAnswers := answers[1]

	fmt.Println("Total Questions: " + strconv.FormatInt(totalQuestions, 10))
	fmt.Println("Total Correct Answers: " + strconv.FormatInt(correctAnswers, 10))
	fmt.Println("Total Incorrect Answers: " + strconv.FormatInt(incorrectAnswers, 10))
}
