package main

import (
	"fmt"
	"strings"
)

func wordCount(s []string) map[string]int {
	m := map[string]int{}
	for _, val := range s {
		v, ok := m[val]
		if ok {
			m[val] = v + 1
		} else {
			m[val] = 1
		}

	}
	return m
}

func main() {
	s := "If it looks like a duck swims like a duck and quacks like a duck then it probably is a duck"
	str := strings.Fields(s)
	result := wordCount(str)
	fmt.Printf("%#v", result)
}
