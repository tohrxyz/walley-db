package lib

import (
	"fmt"
	"strconv"
	"strings"
)

func IsWithinPossibilities(input string, possibilities []string) bool {
	isWithin := false
	for _, in := range possibilities {
		if input == in {
			isWithin = true
			break
		}
	}
	return isWithin
}

// cli: id=23
// conf id:int=8
func isArgValidFromConf(cliArg string, confArg string) bool {
	cliSplit := strings.Split(cliArg, "=")
	nameCli := cliSplit[0]
	valCli := cliSplit[1]

	nameConf := strings.Split(confArg, ":")[0]
	// typeConf := strings.Split((strings.Split(confArg, ":")[1]), "=")[0]
	lenConfStr := strings.Split(confArg, "=")[1]

	lenConf, err := strconv.Atoi(lenConfStr)
	if err != nil {
		panic("can't parse byte len from conf")
	}

	if nameCli != nameConf {
		return false
	}

	if len([]byte(valCli)) > lenConf {
		fmt.Errorf("ERROR: %v of %v is larger than %v required len.\n", valCli, nameCli, lenConf)
		return false
	}

	return true
}
