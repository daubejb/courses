package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func SayHello() string {
	return "Hello Github"
}

const fyal = "problems.csv"

type problem struct {
	q string
	a string
}

var quiz struct {
	questions  float64
	numCorrect float64
}

func main() {
	fmt.Println(SayHello())
	var fileFlag = flag.String("file", fyal, "provide path to quiz file")
	flag.Parse()
	// read in a quiz provided by a csv file
	var lines [][]string = openAndReadFile(*fileFlag)
	// parse file values into a struct of problems
	var problems []problem = parseLines(lines)
	for _, problem := range problems {
		var answer string
		var correctAnswer = problem.a
		// ask the question
		fmt.Printf("%s : ", problem.q)
		// catpure the answer
		fmt.Scan(&answer)
		// evaluate the answer
		if answer == correctAnswer {
			quiz.numCorrect += 1.00
		}
		quiz.questions += 1.00
	}
	var percentCorrect float64 = quiz.numCorrect / quiz.questions * 100.00
	fmt.Printf("You got %v correct out of %v questions. %.2f %%", quiz.numCorrect, quiz.questions, percentCorrect)
}

func openAndReadFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s", fileName))
	}
	rows, err := readFile(f)
	if err != nil {
		exit(fmt.Sprintf("Failed to read file: %s", fileName))
	}
	return rows
}

func readFile(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return rows, err
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
