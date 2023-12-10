package main

import "fmt"

func main() {
	rating := 7.5
	disappointed := '😥'
	normal := '😐'
	good := '😁'
	fail := '😵'

	if rating < 5 {
		fmt.Printf("Disappointed %c", disappointed)
	} else if rating >= 5 && rating < 7 {
		fmt.Printf("Normal %c", normal)
	} else if rating >= 7 && rating < 10 {
		fmt.Printf("Good %c", good)
	} else {
		fmt.Printf("Fail %c", fail)
	}
}
