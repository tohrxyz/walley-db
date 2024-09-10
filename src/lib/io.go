package lib

import (
	"os"
	"path"
)

func Check(e error) error {
	if e != nil {
		return e
	}
	return nil
}

func FilepathFromTableName(tableName string, isDir bool) string {
	if isDir {
		return "./db/" + tableName
	} else {
		return "./db/" + tableName + "/" + tableName + ".wdb"
	}
}

func CreateFileIfNotExists(filepath string) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		_, err := os.Create(filepath)
		if err != nil {
			return Check(err)
		}
	}
	return nil
}

func CheckIfFileOrDirExists(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDirIfNotExists(filepath string) error {
	if !CheckIfFileOrDirExists(filepath) {
		dirname := path.Base(filepath)
		err := os.Mkdir(dirname, 0755)
		return Check(err)
	} else {
		return nil
	}
}

func WriteToFile(filepath string, data []byte) error {
	if !CheckIfFileOrDirExists(filepath) {
		err := CreateFileIfNotExists(filepath)
		return Check(err)
	}

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return Check(err)
	}

	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return Check(err)
	}

	return nil
}

func DeclareTableConf(tableName string) error {
	filepathDir := FilepathFromTableName(tableName, true)
	err := CreateDirIfNotExists(filepathDir)
	if err != nil {
		return Check(err)
	}

	filepath := FilepathFromTableName(tableName, false)
	err = CreateFileIfNotExists(filepath)
	if err != nil {
		return Check(err)
	}

	return nil
}
