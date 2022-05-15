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

func TestSetTimeout(t *testing.T) {
	expected := 30
	assert.Equal(t, expected, SetTimeout(), "Output should be equal")
}
