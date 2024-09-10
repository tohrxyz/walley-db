package lib

import (
	"fmt"
)

func Table(subject string, name string, method string, args []string) error {
	switch method {
	case "CREATE":
		err := DeclareTableConf(name, args)
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func DeclareTableConf(tableName string, args []string) error {
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

	confFilepath := filepathDir + "/" + tableName + ".conf"
	err = CreateFileIfNotExists(confFilepath)
	if err != nil {
		return err
	}
	fmt.Printf("LOG: Conf file for table %v successfully created at %v\n", tableName, confFilepath)

	var dataToWrite []byte
	for _, arg := range args {
		dataToWrite = append(dataToWrite, []byte("\n"+arg)...)
	}
	err = WriteToFile(confFilepath, dataToWrite)
	if err != nil {
		return err
	}
	fmt.Printf("LOG: Write of table conf at %v was successfull.", confFilepath)

	return nil
}
