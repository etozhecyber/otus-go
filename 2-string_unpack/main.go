package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(char byte) bool {
	if (48 <= char) && (char <= 57) {
		return true
	} else {
		return false
	}
}

func unpack(str string) string {

	if len(str) == 1 && !isDigit(str[0]) {
		return str
	}

	result_str := ""
	for i := 1; i < len(str); i++ {
		if isDigit(str[i]) {
			count, _ := strconv.Atoi(string(str[i]))
			if !isDigit(str[i-1]) {
				result_str = result_str + strings.Repeat(string(str[i-1]), count)
			}
		} else {

			if !isDigit(str[i-1]) {
				result_str = result_str + string(str[i-1])
			}
			if i == len(str)-1 { //always add last
				result_str = result_str + string(str[i])
			}
		}
	}
	return result_str
}

func main() {
	fmt.Println(unpack(strings.Join(os.Args[1:], " ")))
}
