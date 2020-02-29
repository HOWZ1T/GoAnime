package main

import (
	"GoAnime/requests"
	"log"
)

func main() {
	// make http get request
	out, err := requests.Get("https://gogoanime.video//search.html?keyword=danganronpa")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(out)
}
