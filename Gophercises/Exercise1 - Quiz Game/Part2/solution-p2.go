package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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
	seconds   int
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
	var ready string
	fmt.Println("Press enter to start")
	fmt.Scanf("%s", ready)
	timer := time.NewTimer(time.Second * time.Duration(quiz.seconds))

	for i, q := range quiz.questions {
		fmt.Printf("Problem #%v: %v = ", i+1, q.question)
		chAnswer := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			chAnswer <- answer
		}()

		select {
		case <-timer.C:
			return
		case answer := <-chAnswer:
			if strings.EqualFold(answer, q.answer) {
				quiz.bingo++
			}
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
	var timer = flag.Int("t", 30, "set in seconds a timer to you quiz. Default 30s")
	flag.Parse()

	file, err := os.Open(*fileName)
	check(err, fmt.Sprintf("The file %v could not loaded.", *fileName))
	defer file.Close()

	freader := csv.NewReader(file)
	questions, err := freader.ReadAll()
	check(err, fmt.Sprintf("The file %v could not be parsed.", *fileName))
	quiz.readQuestions(questions)
	quiz.seconds = *timer
	quiz.start()
	quiz.printScore()

}
