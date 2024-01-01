package main

import "fmt"

func main() {

	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}

	//ys... => ys[0],ys[1],ys[2],ys[3],ys[4]
	votes := append(xs, ys...)

	fmt.Printf("result : %v", votes[5:9])

}
