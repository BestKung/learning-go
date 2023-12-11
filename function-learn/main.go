package main

import "fmt"

func getRatings(rating float32) (float32, string) {

	if rating < 5 {
		return rating, "Disappointed 😥"
	} else if rating >= 5 && rating < 7 {
		return rating, "Normal 😐"
	} else if rating >= 7 && rating < 10 {
		return rating, "Good 😁"
	} else {
		return rating, "Fail 😵"
	}
}

func main() {
	r1, v1 := getRatings(4.7)
	r2, v2 := getRatings(6.1)
	r3, v3 := getRatings(8)
	r4, v4 := getRatings(11)

	fmt.Printf("rating : %.2f , result : %s\n", r1, v1)
	fmt.Printf("rating : %.2f , result : %s\n", r2, v2)
	fmt.Printf("rating : %.2f , result : %s\n", r3, v3)
	fmt.Printf("rating : %.2f , result : %s\n", r4, v4)
}
