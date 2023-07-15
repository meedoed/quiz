package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	problems, err := readFile(*csvFilename)
	if err != nil {
		exit(err.Error())
	}

	correct := startQuiz(problems, *timeLimit)
	fmt.Printf("Вы ответили правильно на %d вопросов из %d\n", correct, len(problems))
}

func startQuiz(problems []problem, timeLimit int) uint {
	fmt.Println("Начало викторины. Ваша задача - ввести правильное число.")

	for {
		fmt.Println("Для начала нажмине enter...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if input.Text() == "" {
			break
		}
	}

	var correct uint
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

problemloop:
	for i, p := range problems {
		fmt.Printf("Задача #%d: %s = \n", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("done")
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	return uint(correct)
}

func readFile(filename string) ([]problem, error) {
	file, err := os.Open(filename)
	if err != nil {
		exit(err.Error())
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to open the CSV file: %v", err)
	}
	return parseLines(lines), nil
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
