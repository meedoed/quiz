package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var csvFilename = flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	problems, err := readFile(*csvFilename)
	if err != nil {
		exit(err.Error())
	}

	correct := startQuiz(problems)
	fmt.Printf("Вы ответили правильно на %d вопросов из %d\n", correct, len(problems))
}

func startQuiz(problems []problem) uint {
	fmt.Println("Начало викторины. Ваша задача - ввести правильное число.")
	fmt.Println("Для начала нажмите enter...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	var correct uint
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	return uint(correct)
}

func readFile(filename string) ([]problem, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
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
