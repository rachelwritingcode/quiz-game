package questions

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
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
func OutputQuestions(questionsAndAnswers map[string]string, timeout int, stdin io.Reader) {

	correctAnswers := 0
	// Set the timer
	timer := time.NewTimer(time.Duration(timeout) * time.Second)

questionsloop:
	// Display all the questions
	for question, answer := range questionsAndAnswers {

		// Display the question
		fmt.Println("What is the answer to this question?")
		fmt.Println(question)
		answerChannel := make(chan string)

		go func() {
			userResponse, err := ReadInput(stdin)
			if err != nil {
				log.Fatal(err)
			}
			userResponse = strings.TrimSpace(string(userResponse))
			answerChannel <- userResponse
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break questionsloop
		case userResponse := <-answerChannel:
			if userResponse == answer {
				correctAnswers++
			}
		}
	}
	fmt.Println("Score is " + strconv.Itoa(correctAnswers) + " Correct Answers out of " + strconv.Itoa((len(questionsAndAnswers))) + " Questions")

}

// Set the timer
func SetTimeout() int {
	timeout := flag.Int("timeout", 30, "length of time per question")
	flag.Parse()
	return *timeout
}
