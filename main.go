package main

import (
	"GoAnime/searcher"
	"GoAnime/sources"
)

func main() {
	searcher.Search("danganronpa", sources.GoGoAnime)
}
