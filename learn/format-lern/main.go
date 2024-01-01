package main

import "fmt"

func main() {
	movie, year, scroll, movieType, hero := "Aventure Endgame", 2019, 8.5, "Sci-Fi", true
	var like rune = '⭐'

	fmt.Printf("Type -- %T , movie : %s \n", movie, movie)
	fmt.Printf("Type -- %T , year : %d \n", year, year)
	fmt.Printf("Type -- %T , scroll : %.2f \n", scroll, scroll)
	fmt.Printf("Type -- %T , movieType : %s \n", movieType, movieType)
	fmt.Printf("Type -- %T , hero : %t \n", hero, hero)
	fmt.Printf("Type -- %T , like : %c \n", like, like)

	fmt.Printf("-----------------------------\n")
	fmt.Printf("Type -- %T , movie : %#v \n", movie, movie)
	fmt.Printf("Type -- %T , year : %#v \n", year, year)
	fmt.Printf("Type -- %T , scroll : %#v \n", scroll, scroll)
	fmt.Printf("Type -- %T , movieType : %#v \n", movieType, movieType)
	fmt.Printf("Type -- %T , hero : %#v \n", hero, hero)
	fmt.Printf("Type -- %T , like : %c \n", like, like)
}
