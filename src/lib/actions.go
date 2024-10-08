package lib

import "fmt"

func Create(subject string, name string, args []string) error {
	switch subject {
	case "TABLE":
		err := Table(subject, name, "CREATE", args)
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func Update(subject string, name string) error {
	fmt.Println("update")

	// switch subject {
	// case "TABLE":
	// 	err := Table(subject, name, "UPDATE")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	break
	// }
	return nil
}

func Insert(subject string, name string, args []string) error {

	switch subject {
	case "TABLE":
		err := Table(subject, name, "INSERT", args)
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func Delete(subject string, name string) error {
	fmt.Println("delete")

	// switch subject {
	// case "TABLE":
	// 	err := Table(subject, name, "DELETE")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	break
	// }
	return nil
}
