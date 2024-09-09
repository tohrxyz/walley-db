package main

import (
	"fmt"
	"os"
)

var POSSIBLE_ACTIONS = []string{"CREATE", "UPDATE", "INSERT", "DELETE"}
var POSSIBLE_SUBJECTS = []string{"TABLE"}

type Command struct {
	Action  string
	Subject string
	Name    string
}

func isWithinPossibilities(input string, possibilities []string) bool {
	isWithin := false
	for _, in := range possibilities {
		if input == in {
			isWithin = true
			break
		}
	}
	return isWithin
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	if len(args) < 3 {
		panic("minimum num of args is 3 (action, subject, name)")
	}

	isValidAction := isWithinPossibilities(args[0], POSSIBLE_ACTIONS)
	fmt.Printf("Is %s valid action: %t\n", args[0], isValidAction)
	if !isValidAction {
		panic("%s is not a valid action")
	}

	isValidSubject := isWithinPossibilities(args[1], POSSIBLE_SUBJECTS)
	fmt.Printf("Is %s valid subject: %t\n", args[1], isValidSubject)
	if !isValidAction {
		panic("%s is not a valid subject")
	}
}
