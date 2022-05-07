package questions

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQuestions(t *testing.T) {

	expected := map[string]string{"1+1": "2", "2+2": "4"}
	testSlice := []string{"1+1", "2", "2+2", "4"}
	assert.Equal(t, expected, ParseQuestions(testSlice), "Output should be equal")

	expected = map[string]string{"3+1": "4", "5+5": "10"}
	testSlice = []string{"1+1", "2", "2+2", "4"}
	assert.NotEqual(t, expected, ParseQuestions(testSlice), "Output should not be equal")
}

func TestStartQuiz(t *testing.T) {

	var stdin bytes.Buffer
	stdin.Write([]byte("Y\n"))
	result := StartQuiz(&stdin)
	assert.Equal(t, true, result)

	stdin.Write([]byte("N\n"))
	result = StartQuiz(&stdin)
	assert.NotEqual(t, true, result)

	stdin.Write([]byte("hello\n"))
	result = StartQuiz(&stdin)
	assert.Equal(t, false, result)
}

// func TestOutputQuestions(t *testing.T) {
// }

func TestCheckAnswers(t *testing.T) {

	testMap := map[string]string{"3+1": "correct", "5+5": "correct"}
	expected := []int64{2, 0}
	assert.Equal(t, expected, CheckAnswers(testMap), "Output should be equal")

	testMap = map[string]string{"3+1": "correct", "5+5": "wrong", "6+1": "correct", "6+5": "wrong"}
	expected = []int64{2, 2}
	assert.Equal(t, expected, CheckAnswers(testMap), "Output should be equal")

	testMap = map[string]string{"3+1": "wrong", "5+5": "wrong", "6+1": "wrong", "6+5": "wrong"}
	expected = []int64{0, 4}
	assert.Equal(t, expected, CheckAnswers(testMap), "Output should be equal")
}
