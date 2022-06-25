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

func findIndex(s [][]string, str string) int {
	for i, n := range s {
		for _, v := range n {
			if str == v {
				return i
			}
		}
	}

	return -1
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

func FindPrimeByRange(start int, end int) []int {
	answer := []int{}
	for i := start; i <= end; i++ {

		var prime bool = true

		for p := 2; p < i; p++ {
			if i != p && i%p == 0 {
				prime = false
			}
		}

		if prime {
			answer = append(answer, i)
		}
	}

	return answer
}

func Group(arr []string) [][]string {

	answer := [][]string{}
	for _, n := range arr {
		index := findIndex(answer, n)
		if index > -1 {
			answer[index] = append(answer[index], n)
		} else {
			answer = append(answer, []string{n})
		}
	}

	return answer
}
