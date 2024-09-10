package lib

import (
	"os"
	"path"
)

func FilepathFromTableName(tableName string, isDir bool) string {
	dbPath := GetDbPath()
	if isDir {
		return dbPath + "/" + tableName
	} else {
		return dbPath + "/" + tableName + "/" + tableName + ".wdb"
	}
}

func CreateFileIfNotExists(filepath string) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		_, err := os.Create(filepath)
		if err != nil {
			return err
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

func CreateDirIfNotExists(filepath string, isAbsolutePath bool) error {
	if !CheckIfFileOrDirExists(filepath) {
		var dirname string
		if isAbsolutePath {
			dirname = filepath
		} else {
			dbPath := GetDbPath()
			dirname = dbPath + "/" + path.Base(filepath)
		}
		err := os.Mkdir(dirname, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteToFile(filepath string, data []byte) error {
	if !CheckIfFileOrDirExists(filepath) {
		err := CreateFileIfNotExists(filepath)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
