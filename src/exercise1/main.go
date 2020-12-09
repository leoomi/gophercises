package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := flag.String("file", "problems.csv", "The CSV file to be read. The default is problems.csv")
	duration := flag.Int("time", 30, "The time limit in seconds for the test")
	flag.Parse()
	fmt.Printf("Loading %s...\n", *filePath)

	file, err := os.Open(*filePath)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	csvFile := csv.NewReader(file)
	lines, err := csvFile.ReadAll()

	fmt.Println(lines)

	problems := makeProblems(lines)
	score := 0

	fmt.Println("Press enter to start asnwering")
	fmt.Scanf("\n")

	ch := make(chan interface{})

	go timer(*duration, ch)
	go questionLoop(problems, &score, ch)

	<-ch

	fmt.Printf("Your score: %d of %d", score, len(problems))
}

func questionLoop(problems []problem, score *int, ch chan interface{}) {
	for _, problem := range problems {
		fmt.Println(problem.q + "?")

		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == problem.a {
			*score++
		}
	}
	close(ch)
}

func timer(duration int, ch chan interface{}) {
	time.Sleep(time.Duration(duration) * time.Second)
	close(ch)
}

type problem struct {
	q string
	a string
}

func makeProblems(lines [][]string) []problem {
	var problems = make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return problems
}
