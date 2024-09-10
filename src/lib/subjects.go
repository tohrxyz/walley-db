package lib

import "fmt"

func Table(subject string, name string, method string) error {
	switch method {
	case "CREATE":
		err := DeclareTableConf(name)
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func DeclareTableConf(tableName string) error {
	filepathDir := FilepathFromTableName(tableName, true)
	fmt.Printf("LOG: filepathDir: %v\n", filepathDir)
	err := CreateDirIfNotExists(filepathDir, false)
	if err != nil {
		return err
	}
	fmt.Println("LOG: Directory for table " + tableName + " successfully created.")

	filepath := FilepathFromTableName(tableName, false)
	fmt.Printf("LOG: filepath for table file: %v\n", filepath)
	err = CreateFileIfNotExists(filepath)
	if err != nil {
		return err
	}
	fmt.Println("LOG: File for table " + tableName + " successfully created.")

	return nil
}
