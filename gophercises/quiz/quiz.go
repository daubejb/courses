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

var quiz struct {
	questions  int
	numCorrect int
}

func main() {
	fmt.Println(SayHello())

	var fileFlag = flag.String("file", fyal, "provide path to quiz file")
	flag.Parse()
	// read in a quiz provided by a csv file
	var problems [][]string = openAndReadFile(*fileFlag)

	for _, problem := range problems {

		var answer string
		var correctAnswer = problem[1]
		// ask the question
		fmt.Printf("%s : ", problem[0])

		// catpure the answer
		fmt.Scan(&answer)

		// evaluate the answer
		if answer == correctAnswer {
			quiz.numCorrect += 1
		}
		quiz.questions += 1
	}

	fmt.Printf("You got %v correct out of %v questions.", quiz.numCorrect, quiz.questions)

}

func openAndReadFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s", fileName)
	}
	rows, err := readFile(f)
	if err != nil {
		fmt.Printf("Failed to read file: %s", fileName)
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
