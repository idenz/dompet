package helper

import (
	"strings"
)

func contains(s []bool, str bool) bool {
	for _, v := range s {
		if v == str {
			return str
		}
	}

	return !str
}

func IsPallindrome(v string) (bool, string) {
	splited := strings.Split(v, "")
	var length int = len(splited)

	answer := make([]bool, length)
	for i := 0; i < length; i++ {
		if splited[i] == splited[length-(i+1)] {
			answer[i] = true
		} else {
			answer[i] = false
		}
	}

	return contains(answer, false), v
}
