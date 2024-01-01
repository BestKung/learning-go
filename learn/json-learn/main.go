package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	ImdbId      string   `json:"imdbId"`
	Title       string   `json:"title"`
	Year        int      `json:"year"`
	Rating      float32  `json:"rating"`
	Genres      []string `json:"genres"`
	IsSuperHero bool     `json:"isSuperHero"`
}

func main() {

	body := `[{
		"imdbId":"AA0081",
		"title":"Aventure Infinity War",
		"year":2018,
		"rating":8.4,
		"genres":["Action","Sci-fi"],
		"isSuperHero":true
	},{
		"imdbId":"BG0667",
		"title":"Aventure EndGame",
		"year":2019,
		"rating":8.4,
		"genres":["Action","Drama"],
		"isSuperHero":true
	}]`

	mv := []movie{}

	err := json.Unmarshal([]byte(body), &mv)
	if err != nil {
		fmt.Println("is error ", err)
	} else {
		fmt.Printf("movie : %#v", mv)
	}

}
