package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSV(t *testing.T) {

	// Read test file
	expected := []string{"4*5", "20"}
	testFile := "test.csv"
	assert.Equal(t, expected, ReadCSV(testFile), "Output should be equal")

	// Files does not exist
	expected = []string{"error"}
	anotherFile := "anotherFile.csv"
	assert.Equal(t, expected, ReadCSV(anotherFile), "Output should be equal")

}
