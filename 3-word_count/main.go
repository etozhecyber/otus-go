package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func topwords(str string, top int) map[string]int {

	//remove all unused characters
	punctuation_list := ".,;'\"\t\n"
	count := make(map[string]int)
	for _, value := range punctuation_list {
		str = strings.ReplaceAll(str, string(value), " ")
	}
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, "  ", " ")
	str = strings.ReplaceAll(str, "  ", " ")

	for _, value := range strings.Split(str, " ") {
		count[value] += 1
	}

	if len(count) > top {

		//create slice of keys for sort
		keys := make([]string, 0, len(count))
		for key := range count {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool { return count[keys[i]] > count[keys[j]] })

		//remove unnecessary items
		for _, key := range keys[top:] {
			delete(count, key)
		}
	}

	return count
}

func main() {
	fmt.Println(topwords(strings.Join(os.Args[1:], " "), 10))
}
