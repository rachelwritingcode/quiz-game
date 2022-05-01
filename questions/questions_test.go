package questions

import (
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

// func TestUserContinue(t *testing.T) {

// 	var stdin bytes.Buffer
// 	stdin.Write([]byte("Y\n"))
// }

// func TestOutputQuestions() {

// }
