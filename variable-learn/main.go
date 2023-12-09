package main

import "fmt"

// movie := "Aventure Endgame" cannot use declare global
func main() {
	movie := "Aventure Endgame"
	var year = 2019
	var scroll float32 = 8.5
	movieType := "Sci-Fi"
	hero := true

	fmt.Println("เรื่อง : ", movie)
	fmt.Println("ปี : ", year)
	fmt.Println("เรตติ้ง : ", scroll)
	fmt.Println("ประเภท : ", movieType)
	fmt.Println("ซุปเปอร์ฮีโร่ : ", hero)
}
