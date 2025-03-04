package utils

import (
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Error to covert string to int: " + err.Error())
	}
	return i
}