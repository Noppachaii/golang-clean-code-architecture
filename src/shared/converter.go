package shared

import "strconv"

func ConvertStringToInteger(stringValue string) int {
	integerValue, err := strconv.Atoi(stringValue)
	if err != nil {
		// show log
		return 0
	}
	return integerValue
}
