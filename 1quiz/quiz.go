package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"time"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	correct := 0
	for i,record := range records {
		fmt.Printf("Question %v: %v \n", i, record[0])
		var ans string
		fmt.Scanf("%s", &ans)

		if ans == record[1] {
			correct ++
		}
	}

	fmt.Printf("SCORE: %v / %v \n", correct, len(records))
}