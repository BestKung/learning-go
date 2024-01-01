package main

import "fmt"

func main() {
	var movie string
	var year int32
	var scroll float32
	var movieType string
	var hero bool
	var like rune

	fmt.Printf("movie : %s \n", movie)
	fmt.Printf("movie : %q \n", movie)

	fmt.Printf("year : %d \n", year)
	fmt.Printf("scroll : %.2f \n", scroll)
	fmt.Printf("movieType : %s \n", movieType)
	fmt.Printf("hero : %t \n", hero)

	fmt.Printf("like : %c \n", like)
	fmt.Printf("like : %d \n", like)
}
