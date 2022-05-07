package questions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

// Retrieve the questions and answers and store in a map
func ParseQuestions(records []string) map[string]string {

	var question string
	var answer string
	questionsAndAnswers := make(map[string]string)

	// Assign first pair of questions
	questionsAndAnswers[records[0]] = records[1]

	// Assign remaining pairs of questions
	for i := 2; i < len(records); i++ {
		if i%2 == 0 {
			question = records[i]
			answer = records[i+1]
			questionsAndAnswers[question] = answer
		}
	}
	return questionsAndAnswers
}

// Read the user input
func ReadInput(stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	text, err := reader.ReadString('\n')
	return text, err
}

// Prompt user if they wish to start the quiz
func StartQuiz(stdin io.Reader) bool {

	var userInput string
	fmt.Println("Do you want to start the quiz? Type Y or N.")

	str, err := ReadInput(stdin)
	if err != nil {
		log.Fatal()
	}
	userInput = string(str)

	if strings.Contains(strings.ToUpper(userInput), "Y") {
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
