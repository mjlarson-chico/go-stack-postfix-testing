package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"stack-string/stackString"
	"strings"
)

var operatorPrecedence = map[string]int{
	"^": 1,
	"*": 5,
	"/": 5,
	"+": 10,
	"-": 10,
}

func inFixToPostFix(inFix []string) []string {
	var postFix []string
	s := stackString.NewStack()

	for i := 0; i < len(inFix); i++ {
		if val, ok := operatorPrecedence[inFix[i]]; ok {
			if s.IsEmpty() || val < operatorPrecedence[s.Top()] {
				s.Push(inFix[i])
			} else if val == operatorPrecedence[s.Top()] {
				postFix = append(postFix, s.Top())
				s.Pop()
				s.Push(inFix[i])
			} else {
				postFix = append(postFix, s.Top())
				s.Pop()
				postFix = append(postFix, s.Top())
				s.Pop()
				s.Push(inFix[i])
			}

		} else {
			postFix = append(postFix, inFix[i])
		}
	}

	for !s.IsEmpty() {
		postFix = append(postFix, s.Top())
		s.Pop()
	}

	return postFix
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your space seperated infix formula: ")
	inFix, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inFix = strings.TrimSuffix(inFix, "\n")
	fmt.Println(inFix)
	splitFormula := strings.Split(inFix, " ")
	fmt.Println(splitFormula)

	postFix := inFixToPostFix(splitFormula)
	fmt.Println(postFix)
}
