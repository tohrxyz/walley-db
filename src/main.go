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

func create(subject string, name string) error {

}
func insert(subject string, name string) error {

}
func update(subject string, name string) error {

}
func delete(subject string, name string) error {

}

func table(subject string, name string, method string) error {

}

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	if len(args) < 3 {
		panic("minimum num of args is 3 (action, subject, name)")
	}

	action := args[0]
	isValidAction := isWithinPossibilities(action, POSSIBLE_ACTIONS)
	fmt.Printf("Is %s valid action: %t\n", action, isValidAction)
	if !isValidAction {
		panic(action + " " + "is not a valid action")
	}

	subject := args[1]
	isValidSubject := isWithinPossibilities(subject, POSSIBLE_SUBJECTS)
	fmt.Printf("Is %s valid subject: %t\n", subject, isValidSubject)
	if !isValidAction {
		panic(subject + " " + "is not a valid subject")
	}

	name := args[2]

	switch action {
	case POSSIBLE_ACTIONS[0]:
		fmt.Println("creating...")
		create(subject, name)
		break
	case POSSIBLE_ACTIONS[1]:
		fmt.Println("updating...")
		update(subject, name)
		break
	case POSSIBLE_SUBJECTS[2]:
		fmt.Println("inserting...")
		insert(subject, name)
		break
	case POSSIBLE_ACTIONS[3]:
		fmt.Println("deleting...")
		delete(subject, name)
		break
	}
}
