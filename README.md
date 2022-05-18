# Gophercises Exercise 1 - Quiz Game 🧑‍🏫

## How to run the program? 🏃
 
- Type in the following to run the quiz game: `go run main.go`


## What is this? 🤔

- I have built a quiz game, the user can insert or user the default csv file which includes questions and answers.
- This program parses those questions and prompts the user to answer questions 
- There is validation built to keep track of all correct answers
- There is timer built-in which is has the default of 30 seconds or you may pass in the parameter `-timeout=XXX` with the desired amount of time.
- If the timer runs out, the quiz automatically exits and will print out the user score.
- Go ahead and give it a try! Test out your math and/or quiz questions of your choice! 😉 


```
go run main.go -timeout=5
Do you want to start the quiz? Type Y or N.
y
What is the answer to this question?
8+6
14
What is the answer to this question?
1+4
5
What is the answer to this question?
5+1
6
What is the answer to this question?
2+4

Score is 3 Correct Answers out of 12 Questions

```

## How to run tests? 🧪

Go into the directory where you wish to run tests and enter the following: `go test`


---

## Original Exercise Requirements

This exercise is broken into two parts to help simplify the process of explaining it as well as to make it easier to solve. The second part is harder than the first, so if you get stuck feel free to move on to another problem then come back to part 2 later.

*Note: I didn't break this into multiple exercises like I do for some exercises because both of these combined should only take ~30m to cover in screencasts.*

### Part 1
 
Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to `problems.csv` (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

**NOTE:** *CSV files may have questions with commas in them. Eg: `"what 2+2, sir?",4` is a valid row in a CSV. I suggest you look into the CSV package in Go and don't try to write your own CSV parser.*

### Part 2

Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn't wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

## Bonus

As a bonus exercises you can also...

1. Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. *Hint: Check out the [strings](https://golang.org/pkg/strings/) package.*
2. Add an option (a new flag) to shuffle the quiz order each time it is run.
