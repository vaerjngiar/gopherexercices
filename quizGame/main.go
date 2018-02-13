package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
)

func main()  {
	csvFilename := flag.String("csv", "problem.csv","a csv file in format of 'question,answer")
	flag.Parse()
	//_ = csvFilename

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	//_ = file

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to read CSV file")
	}
	//fmt.Println(lines)

	problems := parseLines(lines)
	//fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.a {
			//fmt.Println("Thats correct")
			correct++
		}
	}

	fmt.Printf("You scored %d from %d\n", correct, len(problems))


}

func parseLines(lines [][]string) []problem  {
	result := make([]problem, len(lines))

	for i, line := range lines{
		result[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return result


}

type problem struct {
	q string
	a string
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}
