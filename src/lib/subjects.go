package lib

import (
	"bytes"
	"fmt"
	"strings"
)

func Table(subject string, name string, method string, args []string) error {
	switch method {
	case "CREATE":
		err := DeclareTableConf(name, args)
		if err != nil {
			return err
		}
		break
	case "INSERT":
		err := InsertIntoTable(name, args)
		if err != nil {
			return err
		}
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

func InsertIntoTable(name string, args []string) error {
	// 1. load conf
	data, err := LoadConfForTable(name)
	if err != nil {
		panic(err)
	}
	rawStringData := string(data)
	columnArgs, err := RawStringToStruct(rawStringData)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", columnArgs)

	// cli: id=144 name=Peter
	// 2. check if args are valid based on conf
	rawStringDataArray := strings.Split(rawStringData, "\n")[1:] //1: because 0th is blank
	areArgsValid := true
	for i, arg := range args {
		if !isArgValidFromConf(arg, rawStringDataArray[i]) {
			areArgsValid = false
		}
		fmt.Printf("LOG: %v is valid arg\n", arg)
	}
	fmt.Println("Are args from cli valid based on conf?: ", areArgsValid)
	if !areArgsValid {
		panic("invalid args from cli compared to conf")
	}
	// 3. write new record !!! only up to conf column lenght
	recordLen, eachRecordLen, err := GetRecordByteLength(name)
	if err != nil {
		panic(err)
	}
	var dataToWrite []byte
	for j, arg := range args {
		extractedVal := strings.Split(arg, "=")[1]
		exValPadded := make([]byte, eachRecordLen[j])
		// everything is saved as UTF-8 string, even if it's e.g. int64 (8B)
		// this is wip for saving int64 as int64 directly - not utf8 string
		// but it needs to be dynamic based on conf file, not purely guesswork with 8B as default
		// exValNumber, err := strconv.Atoi(extractedVal)
		// if err == nil {
		// 	b := make([]byte, 8)
		// 	binary.LittleEndian.PutUint64(b, uint64(exValNumber))
		// 	fmt.Printf("b: %v\n", b)
		// 	copy(exValPadded, b)
		// } else {
		// 	copy(exValPadded, []byte(extractedVal))
		// }
		copy(exValPadded, []byte(extractedVal))
		exValPaddedShiftedRight := shiftToRight(bytes.Clone(exValPadded), len([]byte(extractedVal)))
		dataToWrite = append(dataToWrite, exValPaddedShiftedRight...)
	}
	err = WriteToFile(FilepathFromTableName(name, false), dataToWrite[:recordLen])
	if err != nil {
		fmt.Errorf("ERROR: Problem with inserting into table, unable to save: %v\n", err)
		panic(err)
	}
	fmt.Printf("LOG: Successfully inserted to db table %v\n", name)
	return nil
}

func LoadConfForTable(tableName string) ([]byte, error) {
	confPath := GetDbPath() + "/" + tableName + "/" + tableName + ".conf"
	data, err := ReadFromFile(confPath)
	if err != nil {
		// panic(err)
		fmt.Errorf("ERROR: Can't load conf file: %v\n", err)
		return nil, err
	}

	return data, nil
}

// from [1, 2, 3, 0, 0]
// to   [0, 0, 1, 2, 3]
// in O(n)
func shiftToRight(currentArray []byte, nonZeroCount int) []byte {
	currlen := len(currentArray)
	newArray := make([]byte, currlen)
	for i := 0; i < nonZeroCount; i++ {
		currElement := currentArray[i]
		newPos := currlen - nonZeroCount + (i * 1)
		newArray[newPos] = currElement
	}

	return newArray
}

func Int64ToBytes(i int64) []byte {
	b := make([]byte, 8)
	for j := uint8(0); j < 8; j++ {
		b[j] = byte(i >> (j * 8))
	}
	return b
}
