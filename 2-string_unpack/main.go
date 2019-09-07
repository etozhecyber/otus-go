package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func unpack(str string) string {
	result_str := ""
	for i := 0; i < len(str); i++ {
		if (49 < str[i]) && (str[i] < 57) {
			count, _ := strconv.Atoi(string(str[i]))
			if i != 0 { //workaround if first digit
				if (str[i-1] < 49) || (57 < str[i-1]) { // check if 2 digit in a row
					result_str = result_str + strings.Repeat(string(str[i-1]), count-1)
				}
			}
		} else {
			result_str = result_str + string(str[i])
		}
	}
	return result_str
}

func main() {
	fmt.Println(unpack(strings.Join(os.Args[1:], " ")))
}
