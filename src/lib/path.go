package lib

import "os"

func GetDbPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return homeDir + "/.walleydb"
}
