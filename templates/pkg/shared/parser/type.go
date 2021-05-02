package parser

import "strconv"

func StringToUint64(value string) (uint64, error) {
	parsedValue := uint64(0)
	if parsedValue, err := strconv.ParseUint(value, 10, 10); err != nil {
		return parsedValue, err
	}
	return parsedValue, nil
}
