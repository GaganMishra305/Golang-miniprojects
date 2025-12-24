package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"time"
	"flag"
)

func main() {
	questionTime := flag.Int("qtime", 5, "time limit per question in seconds")
    flag.Parse()

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

		// adding a timer
		answerCh := make(chan string, 1)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			answerCh <- ans
		}()

		select {
		case ans := <-answerCh:
			if ans == record[1]{
				correct++
			}
		case <-time.After(time.Duration(*questionTime) * time.Second):
			fmt.Println("Time's up for this question")
		}
	}

	fmt.Printf("SCORE: %v / %v \n", correct, len(records))
}