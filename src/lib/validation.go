package lib

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
