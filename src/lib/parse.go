package lib

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseCreateTableArgs(args []string) {

}

type ConfArg struct {
	name       string
	dataType   string
	byteLenght int64
}

func RawStringToStruct(raw string) ([]ConfArg, error) {
	confArgs := strings.Split(raw, "\n")

	if len(confArgs) == 0 {
		fmt.Errorf("ERROR: Zero parsed conf args.")
		return nil, nil
	}
	var columnArgs []ConfArg
	for _, line := range confArgs {
		if line == "" {
			continue
		}

		name := strings.Split(line, ":")[0]
		dataType := strings.Split(strings.Split(line, ":")[1], "=")[0]
		byteLength, err := strconv.ParseInt(strings.Split(strings.Split(line, ":")[1], "=")[1], 10, 8)
		if err != nil {
			byteLength = 0
		}
		currConfArg := ConfArg{
			name:       name,
			dataType:   dataType,
			byteLenght: byteLength,
		}
		columnArgs = append(columnArgs, currConfArg)
	}

	return columnArgs, nil
}

func StructToRawString(confArgs []ConfArg) (string, error) {
	if len(confArgs) == 0 {
		fmt.Errorf("ERROR: Zero passed conf args.")
		return "", nil
	}

	var raw string
	for i, arg := range confArgs {
		if i == 0 {
			raw = arg.name + ":" + arg.dataType + "=" + strconv.Itoa(int(arg.byteLenght)) + "\n"
		}
		raw = raw + arg.name + ":" + arg.dataType + "=" + strconv.Itoa(int(arg.byteLenght)) + "\n"
	}

	return raw, nil
}
