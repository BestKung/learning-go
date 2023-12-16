package main

import "fmt"

type movie struct {
	name        string
	year        float32
	rating      float32
	movieType   []string
	isSuperHero bool
}

func (m movie) getMovie() {
	fmt.Println("name : ", m.name)
	fmt.Println("year : ", m.year)
	fmt.Println("rating : ", m.rating)
	fmt.Println("type : ", m.movieType)
	fmt.Println("supper hero : ", m.isSuperHero)
}

func main() {

	m1 := movie{name: "Aventures: Endgame", year: 2019, rating: 8.4, movieType: []string{"Action", "Drama"}, isSuperHero: true}

	m1.getMovie()

}
