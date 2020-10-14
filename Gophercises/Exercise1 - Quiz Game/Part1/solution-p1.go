package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Check if some error happened
func check(e error, m string) {
	if e != nil {
		println(m)
		os.Exit(1)
	}
}

// Question sruct for manageing questions
type Question struct {
	question string
	answer   string
}

//Quiz is the mains struct
type Quiz struct {
	questions []Question
	bingo     int
	nQuestion int
}

func (quiz *Quiz) readQuestions(lines [][]string) {
	for _, l := range lines {
		q := Question{
			l[0],
			strings.TrimSpace(l[1]),
		}
		quiz.questions = append(quiz.questions, q)
	}
	quiz.nQuestion = len(lines)
}

func (quiz *Quiz) start() {
	var answer string
	for i, q := range quiz.questions {
		fmt.Printf("Problem #%v: %v = ", i+1, q.question)
		fmt.Scanf("%s", &answer)
		if strings.EqualFold(answer, q.answer) {
			//fmt.Printf("\tBingo!\n")
			quiz.bingo++
		}
	}
}

func (quiz *Quiz) printScore() {
	fmt.Printf(`
	Number of questions:	%v
	Correct answer:		%v
	`, quiz.nQuestion, quiz.bingo)
}

var quiz Quiz

func main() {

	var fileName = flag.String("f", "problems.csv", "a csv file. The format is 'question, answer'")
	flag.Parse()

	file, err := os.Open(*fileName)
	check(err, fmt.Sprintf("The file %v could not loaded.", *fileName))
	defer file.Close()

	freader := csv.NewReader(file)
	questions, err := freader.ReadAll()
	check(err, fmt.Sprintf("The file %v could not be parsed.", *fileName))
	quiz.readQuestions(questions)
	quiz.start()
	quiz.printScore()

}
