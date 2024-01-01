package main

import "fmt"

type movie struct {
	name        string
	year        float32
	rating      float32
	movieType   []string
	isSuperHero bool
}

func main() {
	m1 := movie{name: "Aventures: Endgame", year: 2019, rating: 8.4, movieType: []string{"Action", "Drama"}, isSuperHero: true}
	m2 := movie{name: "Infinity War", year: 2018, rating: 8.4, movieType: []string{"Action", "Sci-fi"}, isSuperHero: true}

	//append(mvs,m1,m2)
	mvs := []movie{m1, m2}

	for _, val := range mvs {
		fmt.Printf("movie : %#v \n", val)
	}

}
