package questions

import (
	"fmt"
	"strings"
)

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

func UserContinue() bool {

	var user string
	fmt.Println("Do you want to start the quiz? Type Y or N.")
	fmt.Scanln(&user)

	if strings.Contains(strings.ToUpper(user), "Y") {
		return true
	} else {
		return false
	}
}

// func OutputQuestions(questionsAndAnswers map[string]string) {

// 	var userResponse string
// 	userAnswers := make(map[string]string)

// 	// Loop through all Q & As
// 	for question, answer := range questionsAndAnswers {
// 		// Output the Question to the screen
// 		// Wait for user input
// 		// Check if the answer is correct
// 		// Record if the answer is correct or not
// 		// Once all questions are looped through, output the results of the correct and incorrent answers
// 		fmt.Println("What is the answer to this question?")
// 		fmt.Println(question)
// 		fmt.Scanln(&userResponse)
// 		fmt.Println("Your Answer: " + userResponse)

// 		if strings.Contains(userResponse, answer) {
// 			userAnswers[question] = "correct"
// 		} else {
// 			userAnswers[question] = "wrong"
// 		}
// 	}

// }
