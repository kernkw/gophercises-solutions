package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	csvFile   = flag.String("csv", "problems.csv", "the location of the csv in format `question,answer`")
	timeLimit = flag.Int("timeLimit", 30, "the time limit of the quiz in seconds")
)

func main() {
	flag.Parse()

	fmt.Println(fmt.Sprintf("\nWelcome, you have %d seconds to finish this quiz.\n", *timeLimit))

	lines, err := readCSVLines(*csvFile)
	if err != nil {
		exitErr(err.Error())
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	problems := parseQuiz(lines)
	var correct int
	answerCh := make(chan string)
	for i, p := range problems {
		fmt.Println(fmt.Sprintf("Problem %d: %s", i+1, p.question))
		go func() {
			getAnswer(answerCh)
		}()

		select {
		case <-timer.C:
			fmt.Println("\nYou ran out of time!")
			fmt.Printf("\nYou got %d out of %d correct!\n", correct, len(lines))
			return
		case answer := <-answerCh:
			checkAnswer(answer, p, &correct)
		}
	}
	fmt.Printf("\nYou got %d out of %d correct!\n", correct, len(lines))
}

type problem struct {
	question string
	answer   string
}

func readCSVLines(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", file)
	}

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv")
	}
	return lines, nil
}

func checkAnswer(a string, p problem, correct *int) {
	if a == p.answer {
		*correct++
	}
}

func getAnswer(answerCh chan string) {
	var answer string
	fmt.Scanf("%s", &answer)
	answerCh <- answer
}

func parseQuiz(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]), // To prevent spaces in csv answer data
		}
	}
	return problems
}

func exitErr(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
