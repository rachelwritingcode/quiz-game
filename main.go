package main

import (
	"fmt"
	csv "quiz-game/csv"
)

func main() {

	csvFile := csv.SetCSV()
	records := csv.ReadCSV(csvFile)
	fmt.Println(records)
}
