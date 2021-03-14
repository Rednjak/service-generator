package app

import "strconv"

func stringToUint64(value string) (uint64, error) {
	userID := uint64(0)
	if userID, err := strconv.ParseUint(value, 10, 10); err != nil {
		return userID, err
	}
	return userID, nil
}
