package utils

import "strconv"

func ParseInt(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return id
}
