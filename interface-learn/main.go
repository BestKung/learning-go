package main

import "fmt"

type movie struct {
	name        string
	year        float32
	rating      float32
	votes       []float32
	movieType   []string
	isSuperHero bool
}

type voter interface {
	addVotes(rating float32)
}

func (m *movie) addVotes(rating float32) {
	m.votes = append(m.votes, rating)
}

func vote(v voter, rating float32) {
	v.addVotes(rating)
}

func main() {

	m1 := &movie{name: "Aventures: Endgame", year: 2019, rating: 8.4, movieType: []string{"Action", "Drama"}, isSuperHero: true, votes: []float32{7, 8, 9, 10, 6, 9, 9, 10, 8}}
	m1.addVotes(8)

	fmt.Printf("movie : %#v \n", m1)
	fmt.Printf("votes : %v \n", m1.votes)
}
