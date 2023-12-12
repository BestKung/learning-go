package main

import "fmt"

func main() {
	genres := [...]string{"Action", "Advanture", "Fantasy", "Sci-fi"}

	for i := 0; i < len(genres); i++ {
		fmt.Println("movie : ", genres[i])
	}
	fmt.Println("----------------------------")
	for _, val := range genres {
		fmt.Println("movie : ", val)
	}
}
