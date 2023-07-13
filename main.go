package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var filename = flag.String("file", "problems.csv", "filename with the quiz questions")

func main() {
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	answers, questions, err := startQuiz(file)
	if err != nil {
		panic(fmt.Errorf("error in startQuiz: %v", err))
	}

	fmt.Printf("Вы ответили правильно на %d вопросов из %d\n", answers, questions)
}

func startQuiz(file *os.File) (uint, uint, error) {
	fmt.Println("Начало викторины. Ваша задача - ввести правильное число.")
	fmt.Println("Для начала нажмите enter...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	reader := csv.NewReader(file)
	reader.Comma = ','

	var rightAnswers, questions uint
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		if _, err := strconv.Atoi(record[1]); err != nil {
			return 0, 0, fmt.Errorf("error with answer in file: %v", err)
		}
		fmt.Println(record[0])

		input.Scan()
		if input.Text() == record[1] {
			rightAnswers++
		}
		questions++
	}

	return uint(rightAnswers), questions, nil
}
