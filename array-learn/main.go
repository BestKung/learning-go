package main

import "fmt"

func main() {

	genres := [...]string{"Action", "Advanture", "Fantasy"}
	fmt.Println(len(genres))

	fmt.Println("genres[0]:", genres[0])
	fmt.Println("genres[1]:", genres[1])
	fmt.Println("genres[2]:", genres[2])

	genres[1] = "Sci-fi"
	fmt.Println("genres[1]:", genres[1])

}
