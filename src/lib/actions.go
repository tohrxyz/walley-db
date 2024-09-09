package lib

import "fmt"

func Create(subject string, name string) error {
	fmt.Println("create")
	return nil
}

func Update(subject string, name string) error {
	fmt.Println("update")
	return nil
}

func Insert(subject string, name string) error {
	fmt.Println("insert")
	return nil
}

func Delete(subject string, name string) error {
	fmt.Println("delete")
	return nil
}
