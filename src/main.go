package main

import (
	"fmt"
	"os"
	"walley-db/main/lib"
)

var POSSIBLE_ACTIONS = []string{"CREATE", "UPDATE", "INSERT", "DELETE"}
var POSSIBLE_SUBJECTS = []string{"TABLE"}

type Command struct {
	Action  string
	Subject string
	Name    string
}

func main() {
	// get home dir and create directory for db files, if it doesn't exist yet
	dbPath := lib.GetDbPath()
	err := lib.CreateDirIfNotExists(dbPath, true)
	if err != nil {
		panic(err)
	}

	// get cli args, expect ACTION SUBJECT name
	args := os.Args[1:]
	fmt.Println(args)

	if len(args) < 3 {
		panic("minimum num of args is 3 (action, subject, name)")
	}

	action := args[0]
	isValidAction := lib.IsWithinPossibilities(action, POSSIBLE_ACTIONS)
	fmt.Printf("Is %s valid action: %t\n", action, isValidAction)
	if !isValidAction {
		panic(action + " " + "is not a valid action")
	}

	subject := args[1]
	isValidSubject := lib.IsWithinPossibilities(subject, POSSIBLE_SUBJECTS)
	fmt.Printf("Is %s valid subject: %t\n", subject, isValidSubject)
	if !isValidAction {
		panic(subject + " " + "is not a valid subject")
	}

	name := args[2]

	switch action {
	case POSSIBLE_ACTIONS[0]:
		lib.Create(subject, name)
		break
	case POSSIBLE_ACTIONS[1]:
		fmt.Println("updating...")
		lib.Update(subject, name)
		break
	case POSSIBLE_SUBJECTS[2]:
		fmt.Println("inserting...")
		lib.Insert(subject, name)
		break
	case POSSIBLE_ACTIONS[3]:
		fmt.Println("deleting...")
		lib.Delete(subject, name)
		break
	}
}
